package nas

import (
	//	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	//	"strings"
	//	"time"

	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/ngap/util"
	"etri5gc/openapi/models"
	"etri5gc/openapi/utils/nasConvert"

	//	"github.com/antihax/optional"
	"github.com/mitchellh/mapstructure"
	"github.com/mohae/deepcopy"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/nas/security"

	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/util/fsm"
)

func (gmm *GmmFsm) handleULNASTransport(ue *context.AmfUe, anType models.AccessType,
	ulNasTransport *nasMessage.ULNASTransport) error {
	log.Infoln("Handle UL NAS Transport")

	if ue.AusfClient().IsMacFailed() {
		return fmt.Errorf("NAS message integrity check failed")
	}

	switch ulNasTransport.GetPayloadContainerType() {
	// TS 24.501 5.4.5.2.3 case a)
	case nasMessage.PayloadContainerTypeN1SMInfo:
		return gmm.transport5GSMMessage(ue, anType, ulNasTransport)
	case nasMessage.PayloadContainerTypeSMS:
		return fmt.Errorf("PayloadContainerTypeSMS has not been implemented yet in UL NAS TRANSPORT")
	case nasMessage.PayloadContainerTypeLPP:
		return fmt.Errorf("PayloadContainerTypeLPP has not been implemented yet in UL NAS TRANSPORT")
	case nasMessage.PayloadContainerTypeSOR:
		return fmt.Errorf("PayloadContainerTypeSOR has not been implemented yet in UL NAS TRANSPORT")
	case nasMessage.PayloadContainerTypeUEPolicy:
		log.Infoln("AMF Transfer UEPolicy To PCF")
		ue.CallbackClient().SendN1MessageNotify(models.N1MessageClass_UPDP,
			ulNasTransport.PayloadContainer.GetPayloadContainerContents(), nil)
	case nasMessage.PayloadContainerTypeUEParameterUpdate:
		log.Infoln("AMF Transfer UEParameterUpdate To UDM")
		upuMac, err := nasConvert.UpuAckToModels(ulNasTransport.PayloadContainer.GetPayloadContainerContents())
		if err != nil {
			return err
		}
		err = ue.UdmClient().PutUpuAck(upuMac)
		if err != nil {
			return err
		}
		log.Debugf("UpuMac[%s] in UPU ACK NAS Msg", upuMac)
	case nasMessage.PayloadContainerTypeMultiplePayload:
		return fmt.Errorf("PayloadContainerTypeMultiplePayload has not been implemented yet in UL NAS TRANSPORT")
	}
	return nil
}

func (gmm *GmmFsm) transport5GSMMessage(ue *context.AmfUe, anType models.AccessType,
	ulNasTransport *nasMessage.ULNASTransport) error {
	var pduSessionID int32

	log.Info("Transport 5GSM Message to SMF")

	udminfo := ue.UdmClient().Info()

	smMessage := ulNasTransport.PayloadContainer.GetPayloadContainerContents()

	if id := ulNasTransport.PduSessionID2Value; id != nil {
		pduSessionID = int32(id.GetPduSessionID2Value())
	} else {
		return errors.New("PDU Session ID is nil")
	}

	// case 1): looks up a PDU session routing context for the UE and the PDU session ID IE in case the Old PDU
	// session ID IE is not included
	if ulNasTransport.OldPDUSessionID == nil {
		smContext, smContextExist := ue.SmContextFindByPDUSessionID(pduSessionID)
		requestType := ulNasTransport.RequestType

		if requestType != nil {
			switch requestType.GetRequestTypeValue() {
			case nasMessage.ULNASTransportRequestTypeInitialEmergencyRequest:
				fallthrough
			case nasMessage.ULNASTransportRequestTypeExistingEmergencyPduSession:
				log.Warnf("Emergency PDU Session is not supported")
				gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
					smMessage, pduSessionID, nasMessage.Cause5GMMPayloadWasNotForwarded, nil, 0)
				return nil
			}
		}

		// AMF has a PDU session routing context for the PDU session ID and the UE
		if smContextExist {
			// case i) Request type IE is either not included
			if requestType == nil {
				return gmm.forward5GSMMessageToSMF(ue, anType, pduSessionID, smContext, smMessage)
			}

			switch requestType.GetRequestTypeValue() {
			case nasMessage.ULNASTransportRequestTypeInitialRequest:
				smContext.StoreULNASTransport(ulNasTransport)
				//  perform a local release of the PDU session identified by the PDU session ID and shall request
				// the SMF to perform a local release of the PDU session
				updateData := models.SmContextUpdateData{
					Release: true,
					Cause:   models.Cause_REL_DUE_TO_DUPLICATE_SESSION_ID,
					SmContextStatusUri: fmt.Sprintf("%s/namf-callback/v1/smContextStatus/%s/%d",
						ue.ServingAMF().GetIPv4Uri(), ue.Guti, pduSessionID),
				}
				log.Warningf("Duplicated PDU session ID[%d]", pduSessionID)
				smContext.SetDuplicatedPduSessionID(true)
				response, _, _, err := smContext.SmfClient().SendUpdateSmContextRequest(updateData, nil, nil)
				if err != nil {
					return err
				} else if response == nil {
					err := fmt.Errorf("PDU Session ID[%d] can't be released in DUPLICATE_SESSION_ID case", pduSessionID)
					log.Errorln(err)
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
						smMessage, pduSessionID, nasMessage.Cause5GMMPayloadWasNotForwarded, nil, 0)
				} else if response != nil {
					smContext.SetUserLocation(ue.GetLocInfo().Location)
					responseData := response.JsonData
					n2Info := response.BinaryDataN2SmInformation
					if n2Info != nil {
						switch responseData.N2SmInfoType {
						case models.N2SmInfoType_PDU_RES_REL_CMD:
							log.Debugln("AMF Transfer NGAP PDU Session Resource Release Command from SMF")
							list := ngapType.PDUSessionResourceToReleaseListRelCmd{}
							util.AppendPDUSessionResourceToReleaseListRelCmd(&list, pduSessionID, n2Info)
							gmm.nas.ngap.SendPDUSessionResourceReleaseCommand(ue.RanUe[anType], nil, list)
						}
					}
				}

			// case ii) AMF has a PDU session routing context, and Request type is "existing PDU session"
			case nasMessage.ULNASTransportRequestTypeExistingPduSession:
				if ue.InAllowedNssai(smContext.Snssai(), anType) {
					return gmm.forward5GSMMessageToSMF(ue, anType, pduSessionID, smContext, smMessage)
				} else {
					log.Errorf("S-NSSAI[%v] is not allowed for access type[%s] (PDU Session ID: %d)",
						smContext.Snssai(), anType, pduSessionID)
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
						smMessage, pduSessionID, nasMessage.Cause5GMMPayloadWasNotForwarded, nil, 0)
				}
			// other requestType: AMF forward the 5GSM message, and the PDU session ID IE towards the SMF identified
			// by the SMF ID of the PDU session routing context
			default:
				return gmm.forward5GSMMessageToSMF(ue, anType, pduSessionID, smContext, smMessage)
			}
		} else { // AMF does not have a PDU session routing context for the PDU session ID and the UE
			if requestType == nil {
				log.Warnf("Request type is nil")
				gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
					smMessage, pduSessionID, nasMessage.Cause5GMMPayloadWasNotForwarded, nil, 0)
				return nil
			}
			switch requestType.GetRequestTypeValue() {
			// case iii) if the AMF does not have a PDU session routing context for the PDU session ID and the UE
			// and the Request type IE is included and is set to "initial request"
			case nasMessage.ULNASTransportRequestTypeInitialRequest:

				if newSmContext, err := ue.CreateSmContext(ulNasTransport, pduSessionID, anType); err != nil {
					//create a session context and prepare smf selection query: failed
					log.Errorf("Select SMF failed: %+v", err)
					cause := nasMessage.Cause5GMMPayloadWasNotForwarded
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
						smMessage, pduSessionID, cause, nil, 0)
				} else {
					smfcli := newSmContext.SmfClient()
					ue.Lock.Lock()
					defer ue.Lock.Unlock()
					_, smContextRef, errResponse, problemDetail, err := smfcli.SendCreateSmContextRequest(nil, smMessage)
					if err != nil {
						log.Errorf("CreateSmContextRequest Error: %+v", err)
						return nil
					} else if problemDetail != nil {
						// TODO: error handling
						return fmt.Errorf("Failed to Create smContext[pduSessionID: %d], Error[%v]", pduSessionID, problemDetail)
					} else if errResponse != nil {
						log.Warnf("PDU Session Establishment Request is rejected by SMF[pduSessionId:%d]",
							pduSessionID)
						gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
							errResponse.BinaryDataN1SmMessage, pduSessionID, 0, nil, 0)
					} else {
						//TODO: tqtung - the next few lines should be moved to the AmfUe
						newSmContext.SetSmContextRef(smContextRef)
						newSmContext.SetUserLocation(deepcopy.Copy(ue.GetLocInfo().Location).(models.UserLocation))
						ue.StoreSmContext(pduSessionID, newSmContext)
						log.Infof("create smContext[pduSessionID: %d] Success", pduSessionID)
						// TODO: handle response(response N2SmInfo to RAN if exists)
					}
				}

			case nasMessage.ULNASTransportRequestTypeModificationRequest:
				fallthrough
			case nasMessage.ULNASTransportRequestTypeExistingPduSession:
				if udminfo.UeContextInSmfData != nil {
					// TS 24.501 5.4.5.2.5 case a) 3)
					pduSessionIDStr := fmt.Sprintf("%d", pduSessionID)
					if ueContextInSmf, ok := udminfo.UeContextInSmfData.PduSessions[pduSessionIDStr]; !ok {
						gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
							smMessage, pduSessionID, nasMessage.Cause5GMMPayloadWasNotForwarded, nil, 0)
					} else {
						// TS 24.501 5.4.5.2.3 case a) 1) iv)
						smContext = context.NewSmContext(pduSessionID)
						smContext.SetAccessType(anType)
						smContext.SetSmfID(ueContextInSmf.SmfInstanceId)
						smContext.SetDnn(ueContextInSmf.Dnn)
						smContext.SetPlmnID(*ueContextInSmf.PlmnId)
						ue.StoreSmContext(pduSessionID, smContext)
						return gmm.forward5GSMMessageToSMF(ue, anType, pduSessionID, smContext, smMessage)
					}
				} else {
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
						smMessage, pduSessionID, nasMessage.Cause5GMMPayloadWasNotForwarded, nil, 0)
				}
			default:
			}
		}
	} else {
		// TODO: implement SSC mode3 Op
		return fmt.Errorf("SSC mode3 operation has not been implemented yet")
	}
	return nil
}

func (gmm *GmmFsm) forward5GSMMessageToSMF(
	ue *context.AmfUe,
	accessType models.AccessType,
	pduSessionID int32,
	smContext *context.SmContext,
	smMessage []byte) error {
	smContextUpdateData := models.SmContextUpdateData{
		N1SmMsg: &models.RefToBinaryData{
			ContentId: "N1SmMsg",
		},
	}
	smContextUpdateData.Pei = ue.Pei
	smContextUpdateData.Gpsi = ue.Gpsi
	if !context.CompareUserLocation(ue.GetLocInfo().Location, smContext.UserLocation()) {
		smContextUpdateData.UeLocation = &ue.GetLocInfo().Location
	}

	if accessType != smContext.AccessType() {
		smContextUpdateData.AnType = accessType
	}
	response, errResponse, problemDetail, err := smContext.SmfClient().SendUpdateSmContextRequest(smContextUpdateData, smMessage, nil)

	if err != nil {
		// TODO: error handling
		log.Errorf("Update SMContext error [pduSessionID: %d], Error[%v]", pduSessionID, err)
		return nil
	} else if problemDetail != nil {
		log.Errorf("Update SMContext failed [pduSessionID: %d], problem[%v]", pduSessionID, problemDetail)
		return nil
	} else if errResponse != nil {
		errJSON := errResponse.JsonData
		n1Msg := errResponse.BinaryDataN1SmMessage
		log.Warnf("PDU Session Modification Procedure is rejected by SMF[pduSessionId:%d], Error[%s]",
			pduSessionID, errJSON.Error.Cause)
		if n1Msg != nil {
			gmm.nas.SendDLNASTransport(ue.RanUe[accessType], nasMessage.PayloadContainerTypeN1SMInfo,
				errResponse.BinaryDataN1SmMessage, pduSessionID, 0, nil, 0)
		}
		// TODO: handle n2 info transfer
	} else if response != nil {
		// update SmContext in AMF
		smContext.SetAccessType(accessType)
		smContext.SetUserLocation(ue.GetLocInfo().Location)

		responseData := response.JsonData
		var n1Msg []byte
		n2SmInfo := response.BinaryDataN2SmInformation
		if response.BinaryDataN1SmMessage != nil {
			log.Debug("Receive N1 SM Message from SMF")
			n1Msg, err = gmm.nas.BuildDLNASTransport(ue, accessType, nasMessage.PayloadContainerTypeN1SMInfo,
				response.BinaryDataN1SmMessage, uint8(pduSessionID), nil, nil, 0)
			if err != nil {
				return err
			}
		}

		if response.BinaryDataN2SmInformation != nil {
			log.Debugf("Receive N2 SM Information[%s] from SMF", responseData.N2SmInfoType)
			switch responseData.N2SmInfoType {
			case models.N2SmInfoType_PDU_RES_MOD_REQ:
				list := ngapType.PDUSessionResourceModifyListModReq{}
				util.AppendPDUSessionResourceModifyListModReq(&list, pduSessionID, n1Msg, n2SmInfo)
				gmm.nas.ngap.SendPDUSessionResourceModifyRequest(ue.RanUe[accessType], list)
			case models.N2SmInfoType_PDU_RES_REL_CMD:
				list := ngapType.PDUSessionResourceToReleaseListRelCmd{}
				util.AppendPDUSessionResourceToReleaseListRelCmd(&list, pduSessionID, n2SmInfo)
				gmm.nas.ngap.SendPDUSessionResourceReleaseCommand(ue.RanUe[accessType], n1Msg, list)
			default:
				return fmt.Errorf("Error N2 SM information type[%s]", responseData.N2SmInfoType)
			}
		} else if n1Msg != nil {
			log.Debugf("AMF forward Only N1 SM Message to UE")
			gmm.nas.ngap.SendDownlinkNasTransport(ue.RanUe[accessType], n1Msg, nil)
		}
	}
	return nil
}

// Handle cleartext IEs of Registration Request, which cleattext IEs defined in TS 24.501 4.4.6
func (gmm *GmmFsm) handleRegistrationRequest(ue *context.AmfUe, anType models.AccessType, procedureCode int64,
	registrationRequest *nasMessage.RegistrationRequest) error {
	var guamiFromUeGuti models.Guami

	if ue == nil {
		return fmt.Errorf("AmfUe is nil")
	}

	log.Info("Handle Registration Request")

	if ue.RanUe[anType] == nil {
		return fmt.Errorf("RanUe is nil")
	}

	ue.SetOnGoing(anType, &context.OnGoing{
		Procedure: context.OnGoingProcedureRegistration,
	})

	if ue.T3513 != nil {
		ue.T3513.Stop()
		ue.T3513 = nil // clear the timer
	}
	if ue.T3565 != nil {
		ue.T3565.Stop()
		ue.T3565 = nil // clear the timer
	}

	// TS 24.501 8.2.6.21: if the UE is sending a REGISTRATION REQUEST message as an initial NAS message,
	// the UE has a valid 5G NAS security context and the UE needs to send non-cleartext IEs
	// TS 24.501 4.4.6: When the UE sends a REGISTRATION REQUEST or SERVICE REQUEST message that includes a NAS message
	// container IE, the UE shall set the security header type of the initial NAS message to "integrity protected"
	ausf := ue.AusfClient()
	secinfo := ausf.SecInfo()
	if registrationRequest.NASMessageContainer != nil && !ausf.IsMacFailed() {
		contents := registrationRequest.NASMessageContainer.GetNASMessageContainerContents()

		// TS 24.501 4.4.6: When the UE sends a REGISTRATION REQUEST or SERVICE REQUEST message that includes a NAS
		// message container IE, the UE shall set the security header type of the initial NAS message to
		// "integrity protected"; then the AMF shall decipher the value part of the NAS message container IE
		err := security.NASEncrypt(secinfo.CipheringAlg, secinfo.KnasEnc, secinfo.ULCount.Get(), security.Bearer3GPP,
			security.DirectionUplink, contents)
		if err != nil {
			secinfo.SecurityContextAvailable = false
		} else {
			m := libnas.NewMessage()
			if err := m.GmmMessageDecode(&contents); err != nil {
				return err
			}

			messageType := m.GmmMessage.GmmHeader.GetMessageType()
			if messageType != libnas.MsgTypeRegistrationRequest {
				return errors.New("The payload of NAS Message Container is not Registration Request")
			}
			// TS 24.501 4.4.6: The AMF shall consider the NAS message that is obtained from the NAS message container
			// IE as the initial NAS message that triggered the procedure
			registrationRequest = m.RegistrationRequest
		}
	}
	// TS 33.501 6.4.6 step 3: if the initial NAS message was protected but did not pass the integrity check
	ue.RetransmissionOfInitialNASMsg = ausf.IsMacFailed()

	ue.RegistrationRequest = registrationRequest
	ue.RegistrationType5GS = registrationRequest.NgksiAndRegistrationType5GS.GetRegistrationType5GS()
	switch ue.RegistrationType5GS {
	case nasMessage.RegistrationType5GSInitialRegistration:
		log.Debugf("RegistrationType: Initial Registration")
	case nasMessage.RegistrationType5GSMobilityRegistrationUpdating:
		log.Debugf("RegistrationType: Mobility Registration Updating")
		if ue.State[anType].Is(context.Deregistered) {
			gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMImplicitlyDeregistered, "")
			return fmt.Errorf("Mobility Registration Updating was sent when the UE state was Deregistered")
		}
	case nasMessage.RegistrationType5GSPeriodicRegistrationUpdating:
		log.Debugf("RegistrationType: Periodic Registration Updating")
		if ue.State[anType].Is(context.Deregistered) {
			gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMImplicitlyDeregistered, "")
			return fmt.Errorf("Periodic Registration Updating was sent when the UE state was Deregistered")
		}
	case nasMessage.RegistrationType5GSEmergencyRegistration:
		return fmt.Errorf("Not Supportted RegistrationType: Emergency Registration")
	case nasMessage.RegistrationType5GSReserved:
		ue.RegistrationType5GS = nasMessage.RegistrationType5GSInitialRegistration
		log.Debugf("RegistrationType: Reserved")
	default:
		log.Debugf("RegistrationType: %v, chage state to InitialRegistration", ue.RegistrationType5GS)
		ue.RegistrationType5GS = nasMessage.RegistrationType5GSInitialRegistration
	}

	mobileIdentity5GSContents := registrationRequest.MobileIdentity5GS.GetMobileIdentity5GSContents()
	ue.IdentityTypeUsedForRegistration = nasConvert.GetTypeOfIdentity(mobileIdentity5GSContents[0])
	switch ue.IdentityTypeUsedForRegistration { // get type of identity
	case nasMessage.MobileIdentity5GSTypeNoIdentity:
		log.Debugf("No Identity")
	case nasMessage.MobileIdentity5GSTypeSuci:
		var plmnId string
		ue.Suci, plmnId = nasConvert.SuciToString(mobileIdentity5GSContents)
		ue.PlmnId = context.PlmnIdStringToModels(plmnId)
		log.Debugf("SUCI: %s", ue.Suci)
	case nasMessage.MobileIdentity5GSType5gGuti:
		guamiFromUeGutiTmp, guti := nasConvert.GutiToString(mobileIdentity5GSContents)
		guamiFromUeGuti = guamiFromUeGutiTmp
		log.Debugf("GUTI: %s", guti)

		servedGuami := gmm.nas.amf().ServedGuamiList()[0]
		if reflect.DeepEqual(guamiFromUeGuti, servedGuami) {
			ue.ServingAmfChanged = false
			// refresh 5G-GUTI according to 6.12.3 Subscription temporary identifier, TS33.501
			if secinfo.SecurityContextAvailable {
				amf := gmm.nas.amf()
				amf.FreeTmsi(int64(ue.Tmsi))
				amf.AllocateGutiToUe(ue)
			}
		} else {
			log.Debugf("Serving AMF has changed")
			ue.ServingAmfChanged = true
			ue.Guti = guti
		}
	case nasMessage.MobileIdentity5GSTypeImei:
		imei := nasConvert.PeiToString(mobileIdentity5GSContents)
		ue.Pei = imei
		log.Debugf("PEI: %s", imei)
	case nasMessage.MobileIdentity5GSTypeImeisv:
		imeisv := nasConvert.PeiToString(mobileIdentity5GSContents)
		ue.Pei = imeisv
		log.Debugf("PEI: %s", imeisv)
	}

	// NgKsi: TS 24.501 9.11.3.32
	switch registrationRequest.NgksiAndRegistrationType5GS.GetTSC() {
	case nasMessage.TypeOfSecurityContextFlagNative:
		secinfo.NgKsi.Tsc = models.ScType_NATIVE
	case nasMessage.TypeOfSecurityContextFlagMapped:
		secinfo.NgKsi.Tsc = models.ScType_MAPPED
	}
	secinfo.NgKsi.Ksi = int32(registrationRequest.NgksiAndRegistrationType5GS.GetNasKeySetIdentifiler())
	if secinfo.NgKsi.Tsc == models.ScType_NATIVE && secinfo.NgKsi.Ksi != 7 {
	} else {
		secinfo.NgKsi.Tsc = models.ScType_NATIVE
		secinfo.NgKsi.Ksi = 0
	}

	// Copy UserLocation from ranUe
	ueloc := ue.GetLocInfo()
	ueloc.Location = ue.RanUe[anType].Location()
	ueloc.Tai = ue.RanUe[anType].Tai()

	// Check TAI
	if !context.InTaiList(ueloc.Tai, gmm.nas.amf().SupportTaiList()) {
		gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMTrackingAreaNotAllowed, "")
		return fmt.Errorf("Registration Reject[Tracking area not allowed]")
	}

	if registrationRequest.UESecurityCapability != nil {
		secinfo.UESecurityCapability = *registrationRequest.UESecurityCapability
	} else {
		gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMProtocolErrorUnspecified, "")
		return fmt.Errorf("UESecurityCapability is nil")
	}

	// TODO (TS 23.502 4.2.2.2 step 4): if UE's 5g-GUTI is included & serving AMF has changed
	// since last registration procedure, new AMF may invoke Namf_Communication_UEContextTransfer
	// to old AMF, including the complete registration request nas msg, to request UE's SUPI & UE Context

	if ue.ServingAmfChanged {
		var transferReason models.TransferReason
		switch ue.RegistrationType5GS {
		case nasMessage.RegistrationType5GSInitialRegistration:
			transferReason = models.TransferReason_INIT_REG
		case nasMessage.RegistrationType5GSMobilityRegistrationUpdating:
			fallthrough
		case nasMessage.RegistrationType5GSPeriodicRegistrationUpdating:
			transferReason = models.TransferReason_MOBI_REG
		}

		ueContextTransferRspData, problemDetails, err := ue.AmfClient().UEContextTransferRequest(anType, transferReason)
		if problemDetails != nil {
			if problemDetails.Cause == "INTEGRITY_CHECK_FAIL" || problemDetails.Cause == "CONTEXT_NOT_FOUND" {
				log.Warnf("Can not retrieve UE Context from old AMF[Cause: %s]", problemDetails.Cause)
			} else {
				log.Warnf("UE Context Transfer Request Failed Problem[%+v]", problemDetails)
			}
			secinfo.SecurityContextAvailable = false // need to start authentication procedure later
		} else if err != nil {
			log.Errorf("UE Context Transfer Request Error[%+v]", err)
			gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMUEIdentityCannotBeDerivedByTheNetwork, "")
		} else {
			//TODO: tqtung - should it be done within AmfUe?
			ue.CopyDataFromUeContextModel(*ueContextTransferRspData.UeContext)
		}

	}

	return nil
}

func IdentityVerification(ue *context.AmfUe) bool {
	return ue.Supi != "" || len(ue.Suci) != 0
}

func (gmm *GmmFsm) handleInitialRegistration(ue *context.AmfUe, anType models.AccessType) error {
	log.Infoln("Handle InitialRegistration")

	amf := gmm.nas.amf()
	udminfo := ue.UdmClient().Info()
	// update Kgnb/Kn3iwf
	ue.AusfClient().UpdateSecurityContext(anType)

	// Registration with AMF re-allocation (TS 23.502 4.2.2.2.3)
	if len(udminfo.SubscribedNssai) == 0 {
		gmm.getSubscribedNssai(ue)
	}

	if err := gmm.handleRequestedNssai(ue, anType); err != nil {
		return err
	}

	if ue.RegistrationRequest.Capability5GMM != nil {
		ue.Capability5GMM = *ue.RegistrationRequest.Capability5GMM
	} else {
		gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMProtocolErrorUnspecified, "")
		return fmt.Errorf("Capability5GMM is nil")
	}

	storeLastVisitedRegisteredTAI(ue, ue.RegistrationRequest.LastVisitedRegisteredTAI)

	if ue.RegistrationRequest.MICOIndication != nil {
		log.Warnf("Receive MICO Indication[RAAI: %d], Not Supported",
			ue.RegistrationRequest.MICOIndication.GetRAAI())
	}

	// TODO: Negotiate DRX value if need (TS 23.501 5.4.5)
	negotiateDRXParameters(ue, ue.RegistrationRequest.RequestedDRXParameters)

	// TODO (step 10 optional): send Namf_Communication_RegistrationCompleteNotify to old AMF if need
	if ue.ServingAmfChanged {
		// If the AMF has changed the new AMF notifies the old AMF that the registration of the UE in the new AMF is completed
		req := models.UeRegStatusUpdateReqData{
			TransferStatus: models.UeContextTransferStatus_TRANSFERRED,
		}
		// TODO: based on locol policy, decide if need to change serving PCF for UE
		regStatusTransferComplete, problemDetails, err := ue.AmfClient().RegistrationStatusUpdate(req)
		if problemDetails != nil {
			log.Errorf("Registration Status Update Failed Problem[%+v]", problemDetails)
		} else if err != nil {
			log.Errorf("Registration Status Update Error[%+v]", err)
		} else {
			if regStatusTransferComplete {
				log.Infof("Registration Status Transfer complete")
			}
		}
	}

	if len(ue.Pei) == 0 {
		gmm.nas.SendIdentityRequest(ue.RanUe[anType], anType, nasMessage.MobileIdentity5GSTypeImei)
		return nil
	}

	// TODO (step 12 optional): the new AMF initiates ME identity check by invoking the
	// N5g-eir_EquipmentIdentityCheck_Get service operation

	if ue.ServingAmfChanged || ue.State[models.AccessType_NON_3_GPP_ACCESS].Is(context.Registered) ||
		!udminfo.ContextValid {
		if err := gmm.communicateWithUDM(ue, anType); err != nil {
			gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMPLMNNotAllowed, "")
			return err
		}
	}
	//TODO: tqtung - these next few lines should be removed becuase we do not use NRF, but a check is needed
	pcfinfo := ue.PcfClient().Info()
	//pcfinfo.PcfId = snf.NfId()
	//pcfinfo.PcfUri = snf.NfUri()
	pcfcli := ue.PcfClient()
	pcfcli.Select(amf.Locality())
	problemDetails, err := pcfcli.AMPolicyControlCreate(anType)
	if problemDetails != nil {
		log.Errorf("AM Policy Control Create Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		log.Errorf("AM Policy Control Create Error[%+v]", err)
	}

	//TODO: tqtung: there next code block should be moved to the PcfClient
	// Service Area Restriction are applicable only to 3GPP access
	if anType == models.AccessType__3_GPP_ACCESS {
		if pcfinfo.AmPolicyAssociation != nil && pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
			servAreaRes := pcfinfo.AmPolicyAssociation.ServAreaRes
			if servAreaRes.RestrictionType == models.RestrictionType_ALLOWED_AREAS {
				numOfallowedTAs := 0
				for _, area := range servAreaRes.Areas {
					numOfallowedTAs += len(area.Tacs)
				}
				// if numOfallowedTAs < int(servAreaRes.MaxNumOfTAs) {
				// 	TODO: based on AMF Policy, assign additional allowed area for UE,
				// 	and the upper limit is servAreaRes.MaxNumOfTAs (TS 29.507 4.2.2.3)
				// }
			}
		}
	}

	// TODO (step 18 optional):
	// If the AMF has changed and the old AMF has indicated an existing NGAP UE association towards a N3IWF, the new AMF
	// creates an NGAP UE association towards the N3IWF to which the UE is connectedsend N2 AMF mobility request to N3IWF
	// if anType == models.AccessType_NON_3_GPP_ACCESS && ue.ServingAmfChanged {
	// 	TODO: send N2 AMF Mobility Request
	// }

	amf.AllocateRegistrationArea(ue, anType)
	log.Debugf("Use original GUTI[%s]", ue.Guti)
	gmm.assignLadnInfo(ue, anType)

	amf.AddAmfUeToUePool(ue, ue.Supi)
	/*
		//TODO: tungtq: temporarily commented
		ue.T3502Value = amf.T3502Value
		if anType == models.AccessType__3_GPP_ACCESS {
			ue.T3512Value = amf.T3512Value
		} else {
			ue.Non3gppDeregistrationTimerValue = amf.Non3gppDeregistrationTimerValue
		}

		if anType == models.AccessType__3_GPP_ACCESS {
			gmm.nas.SendRegistrationAccept(ue, anType, nil, nil, nil, nil, nil)
		} else {
			// TS 23.502 4.12.2.2 10a ~ 13: if non-3gpp, AMF should send initial context setup request to N3IWF first,
			// and send registration accept after receiving initial context setup response
			gmm.nas.ngap.SendInitialContextSetupRequest(ue, anType, nil, nil, nil, nil, nil)

			registrationAccept, err := gmm.nas.BuildRegistrationAccept(ue, anType, nil, nil, nil, nil)
			if err != nil {
				log.Errorf("Build Registration Accept: %+v", err)
				return nil
			}
			ue.RegistrationAcceptForNon3GPPAccess = registrationAccept
		}
	*/
	return nil
}

func (gmm *GmmFsm) handleMobilityAndPeriodicRegistrationUpdating(ue *context.AmfUe, anType models.AccessType) error {
	log.Infoln("Handle MobilityAndPeriodicRegistrationUpdating")

	udminfo := ue.UdmClient().Info()
	pcfinfo := ue.PcfClient().Info()
	ueloc := ue.GetLocInfo()

	if ue.RegistrationRequest.UpdateType5GS != nil {
		if ue.RegistrationRequest.UpdateType5GS.GetNGRanRcu() == nasMessage.NGRanRadioCapabilityUpdateNeeded {
			ue.UeRadioCapability = ""
			ue.UeRadioCapabilityForPaging = nil
		}
	}

	// Registration with AMF re-allocation (TS 23.502 4.2.2.2.3)
	if len(udminfo.SubscribedNssai) == 0 {
		gmm.getSubscribedNssai(ue)
	}

	if err := gmm.handleRequestedNssai(ue, anType); err != nil {
		return err
	}

	if ue.RegistrationRequest.Capability5GMM != nil {
		ue.Capability5GMM = *ue.RegistrationRequest.Capability5GMM
	} else {
		if ue.RegistrationType5GS != nasMessage.RegistrationType5GSPeriodicRegistrationUpdating {
			gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMProtocolErrorUnspecified, "")
			return fmt.Errorf("Capability5GMM is nil")
		}
	}

	storeLastVisitedRegisteredTAI(ue, ue.RegistrationRequest.LastVisitedRegisteredTAI)

	if ue.RegistrationRequest.MICOIndication != nil {
		log.Warnf("Receive MICO Indication[RAAI: %d], Not Supported",
			ue.RegistrationRequest.MICOIndication.GetRAAI())
	}

	// TODO: Negotiate DRX value if need (TS 23.501 5.4.5)
	negotiateDRXParameters(ue, ue.RegistrationRequest.RequestedDRXParameters)

	// TODO (step 10 optional): send Namf_Communication_RegistrationCompleteNotify to old AMF if need
	// if ue.ServingAmfChanged {
	// 	If the AMF has changed the new AMF notifies the old AMF that the registration of the UE in the new AMF is completed
	// }

	if len(ue.Pei) == 0 {
		gmm.nas.SendIdentityRequest(ue.RanUe[anType], anType, nasMessage.MobileIdentity5GSTypeImei)
		return nil
	}

	// TODO (step 12 optional): the new AMF initiates ME identity check by invoking the
	// N5g-eir_EquipmentIdentityCheck_Get service operation

	if ue.ServingAmfChanged || ue.State[models.AccessType_NON_3_GPP_ACCESS].Is(context.Registered) ||
		!udminfo.ContextValid {
		if err := gmm.communicateWithUDM(ue, anType); err != nil {
			gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMPLMNNotAllowed, "")
			return err
		}
	}

	var reactivationResult *[16]bool
	var errPduSessionId, errCause []uint8
	ctxList := ngapType.PDUSessionResourceSetupListCxtReq{}
	suList := ngapType.PDUSessionResourceSetupListSUReq{}
	if ue.RegistrationRequest.UplinkDataStatus != nil {
		uplinkDataPsi := nasConvert.PSIToBooleanArray(ue.RegistrationRequest.UplinkDataStatus.Buffer)
		reactivationResult = new([16]bool)
		allowReEstablishPduSession := true

		// determines that the UE is in non-allowed area or is not in allowed area
		if pcfinfo.AmPolicyAssociation != nil && pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
			switch pcfinfo.AmPolicyAssociation.ServAreaRes.RestrictionType {
			case models.RestrictionType_ALLOWED_AREAS:
				allowReEstablishPduSession = context.TacInAreas(ueloc.Tai.Tac, pcfinfo.AmPolicyAssociation.ServAreaRes.Areas)
			case models.RestrictionType_NOT_ALLOWED_AREAS:
				allowReEstablishPduSession = !context.TacInAreas(ueloc.Tai.Tac, pcfinfo.AmPolicyAssociation.ServAreaRes.Areas)
			}
		}

		if !allowReEstablishPduSession {
			for pduSessionId, hasUplinkData := range uplinkDataPsi {
				if hasUplinkData {
					errPduSessionId = append(errPduSessionId, uint8(pduSessionId))
					errCause = append(errCause, nasMessage.Cause5GMMRestrictedServiceArea)
				}
			}
		} else {
			for idx, hasUplinkData := range uplinkDataPsi {
				pduSessionId := int32(idx)
				if smContext, ok := ue.SmContextFindByPDUSessionID(pduSessionId); ok {
					// uplink data are pending for the corresponding PDU session identity
					if hasUplinkData && smContext.AccessType() == models.AccessType__3_GPP_ACCESS {
						response, errResponse, problemDetail, err := smContext.SmfClient().SendUpdateSmContextActivateUpCnxState(anType)
						if response == nil {
							reactivationResult[pduSessionId] = true
							errPduSessionId = append(errPduSessionId, uint8(pduSessionId))
							cause := nasMessage.Cause5GMMProtocolErrorUnspecified
							if errResponse != nil {
								switch errResponse.JsonData.Error.Cause {
								case "OUT_OF_LADN_SERVICE_AREA":
									cause = nasMessage.Cause5GMMLADNNotAvailable
								case "PRIORITIZED_SERVICES_ONLY":
									cause = nasMessage.Cause5GMMRestrictedServiceArea
								case "DNN_CONGESTION", "S-NSSAI_CONGESTION":
									cause = nasMessage.Cause5GMMInsufficientUserPlaneResourcesForThePDUSession
								}
							}
							errCause = append(errCause, cause)

							if problemDetail != nil {
								log.Errorf("Update SmContext Failed Problem[%+v]", problemDetail)
							} else if err != nil {
								log.Errorf("Update SmContext Error[%v]", err.Error())
							}
						} else {
							if ue.RanUe[anType].UeContextRequest() {
								util.AppendPDUSessionResourceSetupListCxtReq(&ctxList, pduSessionId,
									smContext.Snssai(), response.BinaryDataN1SmMessage, response.BinaryDataN2SmInformation)
							} else {
								util.AppendPDUSessionResourceSetupListSUReq(&suList, pduSessionId,
									smContext.Snssai(), response.BinaryDataN1SmMessage, response.BinaryDataN2SmInformation)
							}
						}
					}
				}
			}
		}
	}

	var pduSessionStatus *[16]bool
	if ue.RegistrationRequest.PDUSessionStatus != nil {
		pduSessionStatus = new([16]bool)
		psiArray := nasConvert.PSIToBooleanArray(ue.RegistrationRequest.PDUSessionStatus.Buffer)
		for psi := 1; psi <= 15; psi++ {
			pduSessionId := int32(psi)
			if smContext, ok := ue.SmContextFindByPDUSessionID(pduSessionId); ok {
				if !psiArray[psi] && smContext.AccessType() == anType {
					cause := models.Cause_PDU_SESSION_STATUS_MISMATCH
					causeAll := &context.CauseAll{
						Cause: &cause,
					}
					problemDetail, err := smContext.SmfClient().SendReleaseSmContextRequest(causeAll, "", nil)
					if problemDetail != nil {
						pduSessionStatus[psi] = true
						log.Errorf("Release SmContext Failed Problem[%+v]", problemDetail)
					} else if err != nil {
						pduSessionStatus[psi] = true
						log.Errorf("Release SmContext Error[%v]", err.Error())
					} else {
						pduSessionStatus[psi] = false
					}
				} else {
					pduSessionStatus[psi] = true
				}
			}
		}
	}

	if ue.RegistrationRequest.AllowedPDUSessionStatus != nil {
		allowedPsis := nasConvert.PSIToBooleanArray(ue.RegistrationRequest.AllowedPDUSessionStatus.Buffer)
		if ue.N1N2Message != nil {
			requestData := ue.N1N2Message.Request.JsonData
			n1Msg := ue.N1N2Message.Request.BinaryDataN1Message
			n2Info := ue.N1N2Message.Request.BinaryDataN2Information

			// downlink signalling
			if n2Info == nil {
				if len(suList.List) != 0 {
					nasPdu, err := gmm.nas.BuildRegistrationAccept(ue, anType, pduSessionStatus,
						reactivationResult, errPduSessionId, errCause)
					if err != nil {
						return err
					}
					gmm.nas.ngap.SendPDUSessionResourceSetupRequest(ue.RanUe[anType], nasPdu, suList)
				} else {
					gmm.nas.SendRegistrationAccept(ue, anType, pduSessionStatus,
						reactivationResult, errPduSessionId, errCause, &ctxList)
				}
				switch requestData.N1MessageContainer.N1MessageClass {
				case models.N1MessageClass_SM:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
						n1Msg, requestData.PduSessionId, 0, nil, 0)
				case models.N1MessageClass_LPP:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeLPP, n1Msg, 0, 0, nil, 0)
				case models.N1MessageClass_SMS:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeSMS, n1Msg, 0, 0, nil, 0)
				case models.N1MessageClass_UPDP:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeUEPolicy, n1Msg, 0, 0, nil, 0)
				}
				ue.N1N2Message = nil
				return nil
			}

			smInfo := requestData.N2InfoContainer.SmInfo
			smContext, exist := ue.SmContextFindByPDUSessionID(requestData.PduSessionId)
			if !exist {
				ue.N1N2Message = nil
				return fmt.Errorf("Pdu Session Id not Exists")
			}

			if smContext.AccessType() == models.AccessType_NON_3_GPP_ACCESS {
				if reactivationResult == nil {
					reactivationResult = new([16]bool)
				}
				if allowedPsis[requestData.PduSessionId] {
					// TODO: error handling
					response, errRes, _, err := smContext.SmfClient().SendUpdateSmContextChangeAccessType(true)
					if err != nil {
						return err
					} else if response == nil {
						reactivationResult[requestData.PduSessionId] = true
						errPduSessionId = append(errPduSessionId, uint8(requestData.PduSessionId))
						cause := nasMessage.Cause5GMMProtocolErrorUnspecified
						if errRes != nil {
							switch errRes.JsonData.Error.Cause {
							case "OUT_OF_LADN_SERVICE_AREA":
								cause = nasMessage.Cause5GMMLADNNotAvailable
							case "PRIORITIZED_SERVICES_ONLY":
								cause = nasMessage.Cause5GMMRestrictedServiceArea
							case "DNN_CONGESTION", "S-NSSAI_CONGESTION":
								cause = nasMessage.Cause5GMMInsufficientUserPlaneResourcesForThePDUSession
							}
						}
						errCause = append(errCause, cause)
					} else {
						smContext.SetUserLocation(deepcopy.Copy(ueloc.Location).(models.UserLocation))
						smContext.SetAccessType(models.AccessType__3_GPP_ACCESS)
						if response.BinaryDataN2SmInformation != nil &&
							response.JsonData.N2SmInfoType == models.N2SmInfoType_PDU_RES_SETUP_REQ {
							util.AppendPDUSessionResourceSetupListSUReq(&suList, requestData.PduSessionId,
								smContext.Snssai(), nil, response.BinaryDataN2SmInformation)
						}
					}
				} else {
					log.Warnf("UE was reachable but did not accept to re-activate the PDU Session[%d]",
						requestData.PduSessionId)
					ue.CallbackClient().SendN1N2TransferFailureNotification(models.N1N2MessageTransferCause_UE_NOT_REACHABLE_FOR_SESSION)
				}
			} else if smInfo.N2InfoContent.NgapIeType == models.NgapIeType_PDU_RES_SETUP_REQ {
				var nasPdu []byte
				var err error
				if n1Msg != nil {
					pduSessionId := uint8(smInfo.PduSessionId)
					nasPdu, err = gmm.nas.BuildDLNASTransport(ue, anType, nasMessage.PayloadContainerTypeN1SMInfo,
						n1Msg, pduSessionId, nil, nil, 0)
					if err != nil {
						return err
					}
				}
				util.AppendPDUSessionResourceSetupListSUReq(&suList, smInfo.PduSessionId,
					*smInfo.SNssai, nasPdu, n2Info)
			}
		}
	}

	if ueloc.LocationChanged && ue.PcfClient().Info().RequestTriggerLocationChange {
		updateReq := models.PolicyAssociationUpdateRequest{}
		updateReq.Triggers = append(updateReq.Triggers, models.RequestTrigger_LOC_CH)
		updateReq.UserLoc = &ueloc.Location
		problemDetails, err := ue.PcfClient().AMPolicyControlUpdate(updateReq)
		if problemDetails != nil {
			log.Errorf("AM Policy Control Update Failed Problem[%+v]", problemDetails)
		} else if err != nil {
			log.Errorf("AM Policy Control Update Error[%v]", err)
		}
		ueloc.LocationChanged = false
	}

	// TODO (step 18 optional):
	// If the AMF has changed and the old AMF has indicated an existing NGAP UE association towards a N3IWF, the new AMF
	// creates an NGAP UE association towards the N3IWF to which the UE is connectedsend N2 AMF mobility request to N3IWF
	// if anType == models.AccessType_NON_3_GPP_ACCESS && ue.ServingAmfChanged {
	// 	TODO: send N2 AMF Mobility Request
	// }

	gmm.nas.amf().AllocateRegistrationArea(ue, anType)
	gmm.assignLadnInfo(ue, anType)

	// TODO: GUTI reassignment if need (based on operator poilcy)
	// TODO: T3512/Non3GPP de-registration timer reassignment if need (based on operator policy)

	if ue.RanUe[anType].UeContextRequest() {
		// update Kgnb/Kn3iwf
		ue.AusfClient().UpdateSecurityContext(anType)

		if anType == models.AccessType__3_GPP_ACCESS {
			gmm.nas.SendRegistrationAccept(ue, anType, pduSessionStatus, reactivationResult,
				errPduSessionId, errCause, &ctxList)
		} else {
			gmm.nas.ngap.SendInitialContextSetupRequest(ue, anType, nil, &ctxList, nil, nil, nil)
			registrationAccept, err := gmm.nas.BuildRegistrationAccept(ue, anType,
				pduSessionStatus, reactivationResult, errPduSessionId, errCause)
			if err != nil {
				log.Errorf("Build Registration Accept: %+v", err)
				return nil
			}
			ue.RegistrationAcceptForNon3GPPAccess = registrationAccept
		}
		return nil
	} else {
		nasPdu, err := gmm.nas.BuildRegistrationAccept(ue, anType, pduSessionStatus, reactivationResult,
			errPduSessionId, errCause)
		if err != nil {
			log.Error(err.Error())
		}
		if len(suList.List) != 0 {
			gmm.nas.ngap.SendPDUSessionResourceSetupRequest(ue.RanUe[anType], nasPdu, suList)
		} else {
			gmm.nas.ngap.SendDownlinkNasTransport(ue.RanUe[anType], nasPdu, nil)
		}
		// TODO: when state machaine, remove it
		// ue.ClearRegistrationRequestData(anType)
		return nil
	}
}

// TS 23.502 4.2.2.2.2 step 1
// If available, the last visited TAI shall be included in order to help the AMF produce Registration Area for the UE
func storeLastVisitedRegisteredTAI(ue *context.AmfUe, lastVisitedRegisteredTAI *nasType.LastVisitedRegisteredTAI) {
	if lastVisitedRegisteredTAI != nil {
		plmnID := nasConvert.PlmnIDToString(lastVisitedRegisteredTAI.Octet[1:4])
		nasTac := lastVisitedRegisteredTAI.GetTAC()
		tac := hex.EncodeToString(nasTac[:])

		tai := models.Tai{
			PlmnId: &models.PlmnId{
				Mcc: plmnID[:3],
				Mnc: plmnID[3:],
			},
			Tac: tac,
		}

		ue.GetLocInfo().LastVisitedRegisteredTai = tai
		log.Debugf("Ue Last Visited Registered Tai; %v", ue.GetLocInfo().LastVisitedRegisteredTai)
	}
}

func negotiateDRXParameters(ue *context.AmfUe, requestedDRXParameters *nasType.RequestedDRXParameters) {
	if requestedDRXParameters != nil {
		switch requestedDRXParameters.GetDRXValue() {
		case nasMessage.DRXcycleParameterT32:
			log.Tracef("Requested DRX: T = 32")
			ue.UESpecificDRX = nasMessage.DRXcycleParameterT32
		case nasMessage.DRXcycleParameterT64:
			log.Tracef("Requested DRX: T = 64")
			ue.UESpecificDRX = nasMessage.DRXcycleParameterT64
		case nasMessage.DRXcycleParameterT128:
			log.Tracef("Requested DRX: T = 128")
			ue.UESpecificDRX = nasMessage.DRXcycleParameterT128
		case nasMessage.DRXcycleParameterT256:
			log.Tracef("Requested DRX: T = 256")
			ue.UESpecificDRX = nasMessage.DRXcycleParameterT256
		case nasMessage.DRXValueNotSpecified:
			fallthrough
		default:
			ue.UESpecificDRX = nasMessage.DRXValueNotSpecified
			log.Tracef("Requested DRX: Value not specified")
		}
	}
}

func (gmm *GmmFsm) communicateWithUDM(ue *context.AmfUe, accessType models.AccessType) error {
	log.Debugln("communicateWithUDM")

	udminfo := ue.UdmClient().Info()
	// UDM selection described in TS 23.501 6.3.8
	// TODO: consider udm group id, Routing ID part of SUCI, GPSI or External Group ID (e.g., by the NEF)

	udmc := ue.UdmClient()
	//make a query to find a UDM producer
	udmc.Select()

	problemDetails, err := udmc.UeCmRegistration(accessType, true)
	if problemDetails != nil {
		log.Errorf("UECM_Registration Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		log.Errorf("UECM_Registration Error[%+v]", err)
	}

	problemDetails, err = udmc.SDMGetAmData()
	if problemDetails != nil {
		log.Errorf("SDM_Get AmData Failed Problem[%+v]", problemDetails)
		return fmt.Errorf(problemDetails.Cause)
	} else if err != nil {
		return fmt.Errorf("SDM_Get AmData Error[%+v]", err)
	}

	problemDetails, err = udmc.SDMGetSmfSelectData()
	if problemDetails != nil {
		log.Errorf("SDM_Get SmfSelectData Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		return fmt.Errorf("SDM_Get SmfSelectData Error[%+v]", err)
	}

	problemDetails, err = udmc.SDMGetUeContextInSmfData()
	if problemDetails != nil {
		log.Errorf("SDM_Get UeContextInSmfData Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		return fmt.Errorf("SDM_Get UeContextInSmfData Error[%+v]", err)
	}

	problemDetails, err = udmc.SDMSubscribe()
	if problemDetails != nil {
		log.Errorf("SDM Subscribe Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		log.Errorf("SDM Subscribe Error[%+v]", err)
		return fmt.Errorf("SDM Subscribe Error[%+v]", err)
	}
	udminfo.ContextValid = true
	return nil
}

// NOTE: tungtq: a infinite loop for searching UDM has been replace by a one time search
func (gmm *GmmFsm) getSubscribedNssai(ue *context.AmfUe) error {
	//TODO: tungtq - looks like there a dupplicated UDM selection (free5gc), need further checked

	cli := ue.UdmClient()
	cli.Select()
	problemDetails, err := cli.SDMGetSliceSelectionSubscriptionData()

	if problemDetails != nil {
		log.Errorf("SDM_Get Slice Selection Subscription Data Failed Problem[%+v]", problemDetails)
	} else if err != nil {
		log.Errorf("SDM_Get Slice Selection Subscription Data Error[%+v]", err)
	}
	return err
}

// TS 23.502 4.2.2.2.3 Registration with AMF Re-allocation
func (gmm *GmmFsm) handleRequestedNssai(ue *context.AmfUe, anType models.AccessType) error {
	/* TODO: comment out temporarily by tungtq; will review later
	amf := gmm.nas.backend.Context()
	nssfinfo := ue.GetNssfInfo()

	if ue.RegistrationRequest.RequestedNSSAI != nil {
		requestedNssai, err := nasConvert.RequestedNssaiToModels(ue.RegistrationRequest.RequestedNSSAI)
		if err != nil {
			return fmt.Errorf("Decode failed at RequestedNSSAI[%s]", err)
		}

		needSliceSelection := false
		for _, requestedSnssai := range requestedNssai {
			log.Infof("RequestedNssai - ServingSnssai: %+v, HomeSnssai: %+v",
				requestedSnssai.ServingSnssai, requestedSnssai.HomeSnssai)
			if ue.InSubscribedNssai(*requestedSnssai.ServingSnssai) {
				allowedSnssai := models.AllowedSnssai{
					AllowedSnssai: &models.Snssai{
						Sst: requestedSnssai.ServingSnssai.Sst,
						Sd:  requestedSnssai.ServingSnssai.Sd,
					},
					MappedHomeSnssai: requestedSnssai.HomeSnssai,
				}
				if !ue.InAllowedNssai(*allowedSnssai.AllowedSnssai, anType) {
					nssfinfo.AllowedNssai[anType] = append(nssfinfo.AllowedNssai[anType], allowedSnssai)
				}
			} else {
				needSliceSelection = true
				break
			}
		}

		if needSliceSelection {
			if nssfinfo.NssfUri == "" {
				for {
					err := gmm.nas.backend.NfSelector().SearchNssfNSSelectionInstance(ue, models.NfType_NSSF, models.NfType_AMF, nil)
					if err != nil {
						log.Errorf("AMF can not select an NSSF Instance by NRF[Error: %+v]", err)
						time.Sleep(2 * time.Second)
					} else {
						break
					}
				}
			}

			// Step 4
			problemDetails, err := gmm.consumer.Nssf().NSSelectionGetForRegistration(ue, requestedNssai)
			if problemDetails != nil {
				log.Errorf("NSSelection Get Failed Problem[%+v]", problemDetails)
				gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMProtocolErrorUnspecified, "")
				return fmt.Errorf("Handle Requested Nssai of UE failed")
			} else if err != nil {
				log.Errorf("NSSelection Get Error[%+v]", err)
				gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMProtocolErrorUnspecified, "")
				return fmt.Errorf("Handle Requested Nssai of UE failed")
			}

			// Step 5: Initial AMF send Namf_Communication_RegistrationCompleteNotify to old AMF
			req := models.UeRegStatusUpdateReqData{
				TransferStatus: models.UeContextTransferStatus_NOT_TRANSFERRED,
			}
			_, problemDetails, err = ue.AmfClient().RegistrationStatusUpdate(req)
			if problemDetails != nil {
				log.Errorf("Registration Status Update Failed Problem[%+v]", problemDetails)
			} else if err != nil {
				log.Errorf("Registration Status Update Error[%+v]", err)
			}

			// Step 6
			searchTargetAmfQueryParam := Nnrf_NFDiscovery.SearchNFInstancesParamOpts{}
			if nssfinfo.NetworkSliceInfo != nil {
				netwotkSliceInfo := nssfinfo.NetworkSliceInfo
				if netwotkSliceInfo.TargetAmfSet != "" {
					// TS 29.531
					// TargetAmfSet format: ^[0-9]{3}-[0-9]{2-3}-[A-Fa-f0-9]{2}-[0-3][A-Fa-f0-9]{2}$
					// mcc-mnc-amfRegionId(8 bit)-AmfSetId(10 bit)
					targetAmfSetToken := strings.Split(netwotkSliceInfo.TargetAmfSet, "-")
					guami := amf.ServedGuamiList()[0]
					targetAmfPlmnId := models.PlmnId{
						Mcc: targetAmfSetToken[0],
						Mnc: targetAmfSetToken[1],
					}

					if !reflect.DeepEqual(*guami.PlmnId, targetAmfPlmnId) {
						searchTargetAmfQueryParam.TargetPlmnList =
							optional.NewInterface(openapi.MarshToJsonString([]models.PlmnId{targetAmfPlmnId}))
						searchTargetAmfQueryParam.RequesterPlmnList =
							optional.NewInterface(openapi.MarshToJsonString([]models.PlmnId{*guami.PlmnId}))
					}

					searchTargetAmfQueryParam.AmfRegionId = optional.NewString(targetAmfSetToken[2])
					searchTargetAmfQueryParam.AmfSetId = optional.NewString(targetAmfSetToken[3])
				} else if len(netwotkSliceInfo.CandidateAmfList) > 0 {
					// TODO: select candidate Amf based on local poilcy
					searchTargetAmfQueryParam.TargetNfInstanceId = optional.NewInterface(netwotkSliceInfo.CandidateAmfList[0])
				}
			}

			err = gmm.nas.backend.NfSelector().SearchAmfCommunicationInstance(ue,
				models.NfType_AMF, models.NfType_AMF, &searchTargetAmfQueryParam)
			if err == nil {
				// Condition (A) Step 7: initial AMF find Target AMF via NRF ->
				// Send Namf_Communication_N1MessageNotify to Target AMF
				//TungTQ: need to rewrite the BuildUeContextModel method
				ueContext := models.UeContext{}
				//ueContext := consumer.BuildUeContextModel(ue)
				registerContext := models.RegistrationContextContainer{
					UeContext:        &ueContext,
					AnType:           anType,
					AnN2ApId:         int32(ue.RanUe[anType].RanUeNgapId()),
					RanNodeId:        ue.RanUe[anType].Ran().Id(),
					InitialAmfName:   amf.Name(),
					UserLocation:     &ue.GetLocInfo().Location,
					RrcEstCause:      ue.RanUe[anType].RRCEstablishmentCause,
					UeContextRequest: ue.RanUe[anType].UeContextRequest(),
					AnN2IPv4Addr:     ue.RanUe[anType].Ran().Conn().RemoteAddr().String(),
					AllowedNssai: &models.AllowedNssai{
						AllowedSnssaiList: nssfinfo.AllowedNssai[anType],
						AccessType:        anType,
					},
				}
				if len(nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn) > 0 {
					registerContext.RejectedNssaiInPlmn = nssfinfo.NetworkSliceInfo.RejectedNssaiInPlmn
				}
				if len(nssfinfo.NetworkSliceInfo.RejectedNssaiInTa) > 0 {
					registerContext.RejectedNssaiInTa = nssfinfo.NetworkSliceInfo.RejectedNssaiInTa
				}

				var n1Message bytes.Buffer
				ue.RegistrationRequest.EncodeRegistrationRequest(&n1Message)
				ue.CallbackClient().SendN1MessageNotifyAtAMFReAllocation(n1Message.Bytes(), &registerContext)
			} else {
				// Condition (B) Step 7: initial AMF can not find Target AMF via NRF -> Send Reroute NAS Request to RAN
				allowedNssaiNgap := ngapConvert.AllowedNssaiToNgap(nssfinfo.AllowedNssai[anType])
				gmm.nas.ngap.SendRerouteNasRequest(ue, anType, nil, ue.RanUe[anType].InitialUEMessage, &allowedNssaiNgap)
			}
			return nil
		}
	}

	// if registration request has no requested nssai, or non of snssai in requested nssai is permitted by nssf
	// then use ue subscribed snssai which is marked as default as allowed nssai
	if len(nssfinfo.AllowedNssai[anType]) == 0 {
		for _, snssai := range ue.UdmClient().Info().SubscribedNssai {
			if snssai.DefaultIndication {
				if amf.InPlmnSupportList(*snssai.SubscribedSnssai) {
					allowedSnssai := models.AllowedSnssai{
						AllowedSnssai: snssai.SubscribedSnssai,
					}
					nssfinfo.AllowedNssai[anType] = append(nssfinfo.AllowedNssai[anType], allowedSnssai)
				}
			}
		}
	}
	*/
	return nil
}

func (gmm *GmmFsm) assignLadnInfo(ue *context.AmfUe, accessType models.AccessType) {
	amf := gmm.nas.backend.Context()

	ladnpool := amf.LadnPool()
	udminfo := ue.UdmClient().Info()

	ue.LadnInfo = nil
	if ue.RegistrationRequest.LADNIndication != nil {
		ue.LadnInfo = make([]context.LADN, 0)
		// request for LADN information
		if ue.RegistrationRequest.LADNIndication.GetLen() == 0 {
			if ue.HasWildCardSubscribedDNN() {
				for _, ladn := range ladnpool {
					if ue.TaiListInRegistrationArea(ladn.TaiLists, accessType) {
						ue.LadnInfo = append(ue.LadnInfo, *ladn)
					}
				}
			} else {
				for _, snssaiInfos := range udminfo.SmfSelectionData.SubscribedSnssaiInfos {
					for _, dnnInfo := range snssaiInfos.DnnInfos {
						if ladn, ok := ladnpool[dnnInfo.Dnn]; ok { // check if this dnn is a ladn
							if ue.TaiListInRegistrationArea(ladn.TaiLists, accessType) {
								ue.LadnInfo = append(ue.LadnInfo, *ladn)
							}
						}
					}
				}
			}
		} else {
			requestedLadnList := nasConvert.LadnToModels(ue.RegistrationRequest.LADNIndication.GetLADNDNNValue())
			for _, requestedLadn := range requestedLadnList {
				if ladn, ok := ladnpool[requestedLadn]; ok {
					if ue.TaiListInRegistrationArea(ladn.TaiLists, accessType) {
						ue.LadnInfo = append(ue.LadnInfo, *ladn)
					}
				}
			}
		}
	} else if udminfo.SmfSelectionData != nil {
		for _, snssaiInfos := range udminfo.SmfSelectionData.SubscribedSnssaiInfos {
			for _, dnnInfo := range snssaiInfos.DnnInfos {
				if dnnInfo.Dnn != "*" {
					if ladn, ok := ladnpool[dnnInfo.Dnn]; ok {
						if ue.TaiListInRegistrationArea(ladn.TaiLists, accessType) {
							ue.LadnInfo = append(ue.LadnInfo, *ladn)
						}
					}
				}
			}
		}
	}
}

func (gmm *GmmFsm) handleIdentityResponse(ue *context.AmfUe, identityResponse *nasMessage.IdentityResponse) error {

	log.Info("Handle Identity Response")
	ausf := ue.AusfClient()

	mobileIdentityContents := identityResponse.MobileIdentity.GetMobileIdentityContents()
	if nasConvert.GetTypeOfIdentity(mobileIdentityContents[0]) != ue.RequestIdentityType {
		return fmt.Errorf("Received identity type doesn't match request type")
	}

	if ue.T3570 != nil {
		ue.T3570.Stop()
		ue.T3570 = nil // clear the timer
	}

	switch nasConvert.GetTypeOfIdentity(mobileIdentityContents[0]) { // get type of identity
	case nasMessage.MobileIdentity5GSTypeSuci:
		var plmnId string
		ue.Suci, plmnId = nasConvert.SuciToString(mobileIdentityContents)
		ue.PlmnId = context.PlmnIdStringToModels(plmnId)
		log.Debugf("get SUCI: %s", ue.Suci)
	case nasMessage.MobileIdentity5GSType5gGuti:
		if ausf.IsMacFailed() {
			return fmt.Errorf("NAS message integrity check failed")
		}
		_, guti := nasConvert.GutiToString(mobileIdentityContents)
		ue.Guti = guti
		log.Debugf("get GUTI: %s", guti)
	case nasMessage.MobileIdentity5GSType5gSTmsi:
		if ausf.IsMacFailed() {
			return fmt.Errorf("NAS message integrity check failed")
		}
		sTmsi := hex.EncodeToString(mobileIdentityContents[1:])
		if tmp, err := strconv.ParseInt(sTmsi[4:], 10, 32); err != nil {
			return err
		} else {
			ue.Tmsi = int32(tmp)
		}
		log.Debugf("get 5G-S-TMSI: %s", sTmsi)
	case nasMessage.MobileIdentity5GSTypeImei:
		if ausf.IsMacFailed() {
			return fmt.Errorf("NAS message integrity check failed")
		}
		imei := nasConvert.PeiToString(mobileIdentityContents)
		ue.Pei = imei
		log.Debugf("get PEI: %s", imei)
	case nasMessage.MobileIdentity5GSTypeImeisv:
		if ausf.IsMacFailed() {
			return fmt.Errorf("NAS message integrity check failed")
		}
		imeisv := nasConvert.PeiToString(mobileIdentityContents)
		ue.Pei = imeisv
		log.Debugf("get PEI: %s", imeisv)
	}
	return nil
}

// TS 24501 5.6.3.2
func (gmm *GmmFsm) handleNotificationResponse(ue *context.AmfUe, notificationResponse *nasMessage.NotificationResponse) error {
	log.Info("Handle Notification Response")
	if ue.AusfClient().IsMacFailed() {
		return fmt.Errorf("NAS message integrity check failed")
	}

	if ue.T3565 != nil {
		ue.T3565.Stop()
		ue.T3565 = nil // clear the timer
	}

	if notificationResponse != nil && notificationResponse.PDUSessionStatus != nil {
		psiArray := nasConvert.PSIToBooleanArray(notificationResponse.PDUSessionStatus.Buffer)
		for psi := 1; psi <= 15; psi++ {
			pduSessionId := int32(psi)
			if smContext, ok := ue.SmContextFindByPDUSessionID(pduSessionId); ok {
				if !psiArray[psi] {
					cause := models.Cause_PDU_SESSION_STATUS_MISMATCH
					causeAll := &context.CauseAll{
						Cause: &cause,
					}
					problemDetail, err := smContext.SmfClient().SendReleaseSmContextRequest(causeAll, "", nil)
					if problemDetail != nil {
						log.Errorf("Release SmContext Failed Problem[%+v]", problemDetail)
					} else if err != nil {
						log.Errorf("Release SmContext Error[%v]", err.Error())
					}
				}
			}
		}
	}
	return nil
}

func (gmm *GmmFsm) handleConfigurationUpdateComplete(ue *context.AmfUe,
	configurationUpdateComplete *nasMessage.ConfigurationUpdateComplete) error {
	log.Info("Handle Configuration Update Complete")

	if ue.AusfClient().IsMacFailed() {
		return fmt.Errorf("NAS message integrity check failed")
	}

	// TODO: Stop timer T3555 in TS 24.501 Figure 5.4.4.1.1 in handler
	// TODO: Send acknowledgment by Nudm_SMD_Info_Service to UDM in handler
	//		import "github.com/free5gc/openapi/Nudm_SubscriberDataManagement" client.Info

	return nil
}

func (gmm *GmmFsm) AuthenticationProcedure(ue *context.AmfUe, accessType models.AccessType) (bool, error) {
	log.Info("Authentication procedure")

	ausf := ue.AusfClient()
	ausfinfo := ausf.Info()
	// Check whether UE has SUCI and SUPI
	if IdentityVerification(ue) {
		log.Debugln("UE has SUCI / SUPI")
		if ausf.SecurityContextIsValid() {
			log.Debugln("UE has a valid security context - skip the authentication procedure")
			return true, nil
		}
	} else {
		// Request UE's SUCI by sending identity request
		ausfinfo.IdentityRequestSendTimes++
		gmm.nas.SendIdentityRequest(ue.RanUe[accessType], accessType, nasMessage.MobileIdentity5GSTypeSuci)
		return false, nil
	}

	// TODO: consider ausf group id, Routing ID part of SUCI
	ausf.Select() //make a query to find an ausf producer

	response, problemDetails, err := ausf.SendUEAuthRequest(nil)

	if err != nil {
		log.Errorf("Nausf_UEAU Authenticate Request Error: %+v", err)
		return false, errors.New("Authentication procedure failed")
	} else if problemDetails != nil {
		log.Errorf("Nausf_UEAU Authenticate Request Failed: %+v", problemDetails)
		return false, nil
	}
	ausfinfo.AuthenticationCtx = response
	ausfinfo.ABBA = []uint8{0x00, 0x00} // set ABBA value as described at TS 33.501 Annex A.7.1

	gmm.nas.SendAuthenticationRequest(ue.RanUe[accessType])

	return false, nil
}

// TS 24501 5.6.1
func (gmm *GmmFsm) handleServiceRequest(ue *context.AmfUe, anType models.AccessType,
	serviceRequest *nasMessage.ServiceRequest) error {

	log.Info("Handle Service Request")
	ausf := ue.AusfClient()
	secinfo := ausf.SecInfo()

	if ue.T3513 != nil {
		ue.T3513.Stop()
		ue.T3513 = nil // clear the timer
	}
	if ue.T3565 != nil {
		ue.T3565.Stop()
		ue.T3565 = nil // clear the timer
	}

	pcfinfo := ue.PcfClient().Info()

	// Set No ongoing
	if procedure := ue.OnGoing(anType).Procedure; procedure == context.OnGoingProcedurePaging {
		ue.SetOnGoing(anType, &context.OnGoing{
			Procedure: context.OnGoingProcedureNothing,
		})
	} else if procedure != context.OnGoingProcedureNothing {
		log.Warnf("UE should not in OnGoing[%s]", procedure)
	}

	// Send Authtication / Security Procedure not support
	if !ue.AusfClient().SecurityContextIsValid() {
		log.Warnf("No Security Context : SUPI[%s]", ue.Supi)
		gmm.nas.SendServiceReject(ue.RanUe[anType], nil, nasMessage.Cause5GMMUEIdentityCannotBeDerivedByTheNetwork)
		gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[anType],
			context.UeContextN2NormalRelease, ngapType.CausePresentNas, ngapType.CauseNasPresentNormalRelease)
		return nil
	}

	// TS 24.501 8.2.6.21: if the UE is sending a REGISTRATION REQUEST message as an initial NAS message,
	// the UE has a valid 5G NAS security context and the UE needs to send non-cleartext IEs
	// TS 24.501 4.4.6: When the UE sends a REGISTRATION REQUEST or SERVICE REQUEST message that includes a NAS message
	// container IE, the UE shall set the security header type of the initial NAS message to "integrity protected"
	if serviceRequest.NASMessageContainer != nil {
		contents := serviceRequest.NASMessageContainer.GetNASMessageContainerContents()

		// TS 24.501 4.4.6: When the UE sends a REGISTRATION REQUEST or SERVICE REQUEST message that includes a NAS
		// message container IE, the UE shall set the security header type of the initial NAS message to
		// "integrity protected"; then the AMF shall decipher the value part of the NAS message container IE
		err := security.NASEncrypt(secinfo.CipheringAlg, secinfo.KnasEnc, secinfo.ULCount.Get(), security.Bearer3GPP,
			security.DirectionUplink, contents)

		if err != nil {
			secinfo.SecurityContextAvailable = false
		} else {
			m := libnas.NewMessage()
			if err := m.GmmMessageDecode(&contents); err != nil {
				return err
			}

			messageType := m.GmmMessage.GmmHeader.GetMessageType()
			if messageType != libnas.MsgTypeServiceRequest {
				return errors.New("The payload of NAS message Container is not service request")
			}
			// TS 24.501 4.4.6: The AMF shall consider the NAS message that is obtained from the NAS message container
			// IE as the initial NAS message that triggered the procedure
			serviceRequest = m.ServiceRequest
		}
		// TS 33.501 6.4.6 step 3: if the initial NAS message was protected but did not pass the integrity check
		ue.RetransmissionOfInitialNASMsg = ausf.IsMacFailed()
	}

	serviceType := serviceRequest.GetServiceTypeValue()
	var reactivationResult, acceptPduSessionPsi *[16]bool
	var errPduSessionId, errCause []uint8
	var targetPduSessionId int32
	suList := ngapType.PDUSessionResourceSetupListSUReq{}
	ctxList := ngapType.PDUSessionResourceSetupListCxtReq{}

	if serviceType == nasMessage.ServiceTypeEmergencyServices ||
		serviceType == nasMessage.ServiceTypeEmergencyServicesFallback {
		log.Warnf("emergency service is not supported")
	}

	if serviceType == nasMessage.ServiceTypeSignalling {
		err := gmm.sendServiceAccept(ue, anType, ctxList, suList, nil, nil, nil, nil)
		return err
	}
	if ue.N1N2Message != nil {
		requestData := ue.N1N2Message.Request.JsonData
		if ue.N1N2Message.Request.BinaryDataN2Information != nil {
			if requestData.N2InfoContainer.N2InformationClass == models.N2InformationClass_SM {
				targetPduSessionId = requestData.N2InfoContainer.SmInfo.PduSessionId
			} else {
				ue.N1N2Message = nil
				return fmt.Errorf("Service Request triggered by Network has not implemented about non SM N2Info")
			}
		}
	}

	if serviceRequest.UplinkDataStatus != nil {
		uplinkDataPsi := nasConvert.PSIToBooleanArray(serviceRequest.UplinkDataStatus.Buffer)
		reactivationResult = new([16]bool)
		ue.SmContextList.Range(func(key, value interface{}) bool {
			pduSessionID := key.(int32)
			smContext := value.(*context.SmContext)

			if pduSessionID == targetPduSessionId && serviceType == nasMessage.ServiceTypeMobileTerminatedServices {
				// Skipping SendUpdateSmContextActivateUpCnxState for the following reason:
				//   In Step 4 of 4.2.3.2 UE Triggered Service Request in TS23.502
				//   > This procedure is triggered by the SMF but the PDU Session(s) identified by the UE
				//   > correlates to other PDU Session ID(s) than the one triggering the procedure
				// However, in the case of Mo-data etc., it cannot be skipped because AMF need to know
				// latest N2SmInformation even if the UE has known the N2Information received at
				// previous N1N2MessageTransfer.
				return true
			}
			if uplinkDataPsi[pduSessionID] && smContext.AccessType() == models.AccessType__3_GPP_ACCESS {
				response, errRes, _, err := smContext.SmfClient().SendUpdateSmContextActivateUpCnxState(models.AccessType__3_GPP_ACCESS)
				if err != nil {
					log.Errorf("SendUpdateSmContextActivateUpCnxState[pduSessionID:%d] Error: %+v",
						pduSessionID, err)
				} else if response == nil {
					reactivationResult[pduSessionID] = true
					errPduSessionId = append(errPduSessionId, uint8(pduSessionID))
					cause := nasMessage.Cause5GMMProtocolErrorUnspecified
					if errRes != nil {
						switch errRes.JsonData.Error.Cause {
						case "OUT_OF_LADN_SERVICE_AREA":
							cause = nasMessage.Cause5GMMLADNNotAvailable
						case "PRIORITIZED_SERVICES_ONLY":
							cause = nasMessage.Cause5GMMRestrictedServiceArea
						case "DNN_CONGESTION", "S-NSSAI_CONGESTION":
							cause = nasMessage.Cause5GMMInsufficientUserPlaneResourcesForThePDUSession
						}
					}
					errCause = append(errCause, cause)
				} else if ue.RanUe[anType].UeContextRequest() {
					util.AppendPDUSessionResourceSetupListCxtReq(&ctxList,
						pduSessionID, smContext.Snssai(), nil, response.BinaryDataN2SmInformation)
				} else {
					util.AppendPDUSessionResourceSetupListSUReq(&suList,
						pduSessionID, smContext.Snssai(), nil, response.BinaryDataN2SmInformation)
				}
			}
			return true
		})
	}
	if serviceRequest.PDUSessionStatus != nil {
		acceptPduSessionPsi = new([16]bool)
		psiArray := nasConvert.PSIToBooleanArray(serviceRequest.PDUSessionStatus.Buffer)
		ue.SmContextList.Range(func(key, value interface{}) bool {
			pduSessionID := key.(int32)
			smContext := value.(*context.SmContext)
			if smContext.AccessType() == anType {
				if !psiArray[pduSessionID] {
					cause := models.Cause_PDU_SESSION_STATUS_MISMATCH
					causeAll := &context.CauseAll{
						Cause: &cause,
					}
					problemDetail, err := smContext.SmfClient().SendReleaseSmContextRequest(causeAll, "", nil)
					if problemDetail != nil {
						log.Errorf("Release SmContext Failed Problem[%+v]", problemDetail)
					} else if err != nil {
						log.Errorf("Release SmContext Error[%v]", err.Error())
					}
				} else {
					acceptPduSessionPsi[pduSessionID] = true
				}
			}
			return true
		})
	}
	switch serviceType {
	case nasMessage.ServiceTypeMobileTerminatedServices: // Trigger by Network
		if ue.N1N2Message != nil {
			requestData := ue.N1N2Message.Request.JsonData
			n1Msg := ue.N1N2Message.Request.BinaryDataN1Message
			n2Info := ue.N1N2Message.Request.BinaryDataN2Information

			// downlink signalling
			if n2Info == nil {
				err := gmm.sendServiceAccept(ue, anType, ctxList, suList, acceptPduSessionPsi,
					reactivationResult, errPduSessionId, errCause)
				if err != nil {
					return err
				}
				switch requestData.N1MessageContainer.N1MessageClass {
				case models.N1MessageClass_SM:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType],
						nasMessage.PayloadContainerTypeN1SMInfo, n1Msg, requestData.PduSessionId, 0, nil, 0)
				case models.N1MessageClass_LPP:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeLPP, n1Msg, 0, 0, nil, 0)
				case models.N1MessageClass_SMS:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeSMS, n1Msg, 0, 0, nil, 0)
				case models.N1MessageClass_UPDP:
					gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeUEPolicy, n1Msg, 0, 0, nil, 0)
				}
				ue.N1N2Message = nil
				return nil
			}
			// TODO: Area of validity for the N2 SM information
			smInfo := requestData.N2InfoContainer.SmInfo
			smContext, exist := ue.SmContextFindByPDUSessionID(requestData.PduSessionId)
			if !exist {
				ue.N1N2Message = nil
				return fmt.Errorf("Service Request triggered by Network error for pduSessionId does not exist")
			}

			if smContext.AccessType() == models.AccessType_NON_3_GPP_ACCESS {
				if serviceRequest.AllowedPDUSessionStatus != nil {
					allowPduSessionPsi := nasConvert.PSIToBooleanArray(serviceRequest.AllowedPDUSessionStatus.Buffer)
					if reactivationResult == nil {
						reactivationResult = new([16]bool)
					}
					if allowPduSessionPsi[requestData.PduSessionId] {
						response, errRes, _, err := smContext.SmfClient().SendUpdateSmContextChangeAccessType(true)
						if err != nil {
							return err
						} else if response == nil {
							reactivationResult[requestData.PduSessionId] = true
							errPduSessionId = append(errPduSessionId, uint8(requestData.PduSessionId))
							cause := nasMessage.Cause5GMMProtocolErrorUnspecified
							if errRes != nil {
								switch errRes.JsonData.Error.Cause {
								case "OUT_OF_LADN_SERVICE_AREA":
									cause = nasMessage.Cause5GMMLADNNotAvailable
								case "PRIORITIZED_SERVICES_ONLY":
									cause = nasMessage.Cause5GMMRestrictedServiceArea
								case "DNN_CONGESTION", "S-NSSAI_CONGESTION":
									cause = nasMessage.Cause5GMMInsufficientUserPlaneResourcesForThePDUSession
								}
							}
							errCause = append(errCause, cause)
						} else {
							smContext.SetUserLocation(deepcopy.Copy(ue.GetLocInfo().Location).(models.UserLocation))
							smContext.SetAccessType(models.AccessType__3_GPP_ACCESS)
							if response.BinaryDataN2SmInformation != nil &&
								response.JsonData.N2SmInfoType == models.N2SmInfoType_PDU_RES_SETUP_REQ {
								if ue.RanUe[anType].UeContextRequest() {
									util.AppendPDUSessionResourceSetupListCxtReq(&ctxList,
										requestData.PduSessionId, smContext.Snssai(), nil, response.BinaryDataN2SmInformation)
								} else {
									util.AppendPDUSessionResourceSetupListSUReq(&suList,
										requestData.PduSessionId, smContext.Snssai(), nil, response.BinaryDataN2SmInformation)
								}
							}
						}
					} else {
						log.Warnf("UE was reachable but did not accept to re-activate the PDU Session[%d]",
							requestData.PduSessionId)
						ue.CallbackClient().SendN1N2TransferFailureNotification(models.N1N2MessageTransferCause_UE_NOT_REACHABLE_FOR_SESSION)
					}
				}
			} else if smInfo.N2InfoContent.NgapIeType == models.NgapIeType_PDU_RES_SETUP_REQ {
				var nasPdu []byte
				var err error
				if n1Msg != nil {
					pduSessionId := uint8(smInfo.PduSessionId)
					nasPdu, err = gmm.nas.BuildDLNASTransport(ue, anType, nasMessage.PayloadContainerTypeN1SMInfo,
						n1Msg, pduSessionId, nil, nil, 0)
					if err != nil {
						return err
					}
				}
				if ue.RanUe[anType].UeContextRequest() {
					util.AppendPDUSessionResourceSetupListCxtReq(&ctxList, smInfo.PduSessionId, *smInfo.SNssai, nasPdu, n2Info)
				} else {
					util.AppendPDUSessionResourceSetupListSUReq(&suList, smInfo.PduSessionId, *smInfo.SNssai,
						nasPdu, n2Info)
				}
			}
			err := gmm.sendServiceAccept(ue, anType, ctxList, suList, acceptPduSessionPsi,
				reactivationResult, errPduSessionId, errCause)
			if err != nil {
				return err
			}
		}
		// downlink signaling
		if pcfinfo.ConfigurationUpdateMessage != nil {
			err := gmm.sendServiceAccept(ue, anType, ctxList, suList,
				acceptPduSessionPsi, reactivationResult, errPduSessionId, errCause)
			if err != nil {
				return err
			}
			mobilityRestrictionList := util.BuildIEMobilityRestrictionList(ue)
			gmm.nas.ngap.SendDownlinkNasTransport(ue.RanUe[models.AccessType__3_GPP_ACCESS],
				pcfinfo.ConfigurationUpdateMessage, &mobilityRestrictionList)
			pcfinfo.ConfigurationUpdateMessage = nil
		}
	case nasMessage.ServiceTypeData:
		if anType == models.AccessType__3_GPP_ACCESS {
			if pcfinfo.AmPolicyAssociation != nil && pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
				var accept bool
				switch pcfinfo.AmPolicyAssociation.ServAreaRes.RestrictionType {
				case models.RestrictionType_ALLOWED_AREAS:
					accept = context.TacInAreas(ue.GetLocInfo().Tai.Tac, pcfinfo.AmPolicyAssociation.ServAreaRes.Areas)
				case models.RestrictionType_NOT_ALLOWED_AREAS:
					accept = !context.TacInAreas(ue.GetLocInfo().Tai.Tac, pcfinfo.AmPolicyAssociation.ServAreaRes.Areas)
				}

				if !accept {
					gmm.nas.SendServiceReject(ue.RanUe[anType], nil, nasMessage.Cause5GMMRestrictedServiceArea)
					return nil
				}
			}
			err := gmm.sendServiceAccept(ue, anType, ctxList, suList, acceptPduSessionPsi,
				reactivationResult, errPduSessionId, errCause)
			if err != nil {
				return err
			}
		} else {
			err := gmm.sendServiceAccept(ue, anType, ctxList, suList, acceptPduSessionPsi,
				reactivationResult, errPduSessionId, errCause)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("Service Type[%d] is not supported", serviceType)
	}
	if len(errPduSessionId) != 0 {
		log.Info(errPduSessionId, errCause)
	}
	ue.N1N2Message = nil
	return nil
}

func (gmm *GmmFsm) sendServiceAccept(ue *context.AmfUe, anType models.AccessType, ctxList ngapType.PDUSessionResourceSetupListCxtReq,
	suList ngapType.PDUSessionResourceSetupListSUReq, pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool, errPduSessionId, errCause []uint8) error {
	if ue.RanUe[anType].UeContextRequest() {
		// update Kgnb/Kn3iwf
		ue.AusfClient().UpdateSecurityContext(anType)

		nasPdu, err := gmm.nas.BuildServiceAccept(ue, anType, pDUSessionStatus, reactivationResult,
			errPduSessionId, errCause)
		if err != nil {
			return err
		}
		if len(ctxList.List) != 0 {
			gmm.nas.ngap.SendInitialContextSetupRequest(ue, anType, nasPdu, &ctxList, nil, nil, nil)
		} else {
			gmm.nas.ngap.SendInitialContextSetupRequest(ue, anType, nasPdu, nil, nil, nil, nil)
		}
	} else if len(suList.List) != 0 {
		nasPdu, err := gmm.nas.BuildServiceAccept(ue, anType, pDUSessionStatus, reactivationResult,
			errPduSessionId, errCause)
		if err != nil {
			return err
		}
		gmm.nas.ngap.SendPDUSessionResourceSetupRequest(ue.RanUe[anType], nasPdu, suList)
	} else {
		gmm.nas.SendServiceAccept(ue.RanUe[anType], anType, pDUSessionStatus, reactivationResult,
			errPduSessionId, errCause)
	}
	return nil
}

// TS 24.501 5.4.1
func (gmm *GmmFsm) handleAuthenticationResponse(ue *context.AmfUe, accessType models.AccessType,
	authenticationResponse *nasMessage.AuthenticationResponse) error {
	log.Info("Handle Authentication Response")

	if ue.T3560 != nil {
		ue.T3560.Stop()
		ue.T3560 = nil // clear the timer
	}

	ausfinfo := ue.AusfClient().Info()
	if ausfinfo.AuthenticationCtx == nil {
		return fmt.Errorf("Ue Authentication Context is nil")
	}

	switch ausfinfo.AuthenticationCtx.AuthType {
	case models.AuthType__5_G_AKA:
		var av5gAka models.Av5gAka
		if err := mapstructure.Decode(ausfinfo.AuthenticationCtx.Var5gAuthData, &av5gAka); err != nil {
			return fmt.Errorf("Var5gAuthData Convert Type Error")
		}
		resStar := authenticationResponse.AuthenticationResponseParameter.GetRES()

		// Calculate HRES* (TS 33.501 Annex A.5)
		p0, err := hex.DecodeString(av5gAka.Rand)
		if err != nil {
			return err
		}
		p1 := resStar[:]
		concat := append(p0, p1...)
		hResStarBytes := sha256.Sum256(concat)
		hResStar := hex.EncodeToString(hResStarBytes[16:])

		if hResStar != av5gAka.HxresStar {
			log.Errorf("HRES* Validation Failure (received: %s, expected: %s)", hResStar, av5gAka.HxresStar)

			if ue.IdentityTypeUsedForRegistration == nasMessage.MobileIdentity5GSType5gGuti && ausfinfo.IdentityRequestSendTimes == 0 {
				ausfinfo.IdentityRequestSendTimes++
				gmm.nas.SendIdentityRequest(ue.RanUe[accessType], accessType, nasMessage.MobileIdentity5GSTypeSuci)
				return nil
			} else {
				gmm.nas.SendAuthenticationReject(ue.RanUe[accessType], "")
				return gmm.sm.SendEvent(ue.State[accessType], AuthFailEvent, fsm.ArgsType{
					ArgAmfUe:      ue,
					ArgAccessType: accessType,
				})
			}
		}

		response, problemDetails, err := ue.AusfClient().SendAuth5gAkaConfirmRequest(hex.EncodeToString(resStar[:]))
		if err != nil {
			return err
		} else if problemDetails != nil {
			log.Debugf("Auth5gAkaConfirm Error[Problem Detail: %+v]", problemDetails)
			return nil
		}
		switch response.AuthResult {
		case models.AuthResult_SUCCESS:
			ue.UnauthenticatedSupi = false
			ausfinfo.Kseaf = response.Kseaf
			ue.Supi = response.Supi
			ue.AusfClient().DerivateKamf()
			log.Debugln("ue.DerivateKamf()", ausfinfo.Kamf)
			return gmm.sm.SendEvent(ue.State[accessType], AuthSuccessEvent, fsm.ArgsType{
				ArgAmfUe:      ue,
				ArgAccessType: accessType,
				ArgEAPSuccess: false,
				ArgEAPMessage: "",
			})
		case models.AuthResult_FAILURE:
			if ue.IdentityTypeUsedForRegistration == nasMessage.MobileIdentity5GSType5gGuti && ausfinfo.IdentityRequestSendTimes == 0 {
				ausfinfo.IdentityRequestSendTimes++
				gmm.nas.SendIdentityRequest(ue.RanUe[accessType], accessType, nasMessage.MobileIdentity5GSTypeSuci)
				return nil
			} else {
				gmm.nas.SendAuthenticationReject(ue.RanUe[accessType], "")
				return gmm.sm.SendEvent(ue.State[accessType], AuthFailEvent, fsm.ArgsType{
					ArgAmfUe:      ue,
					ArgAccessType: accessType,
				})
			}
		}
	case models.AuthType_EAP_AKA_PRIME:
		response, problemDetails, err := ue.AusfClient().SendEapAuthConfirmRequest(*authenticationResponse.EAPMessage)
		if err != nil {
			return err
		} else if problemDetails != nil {
			log.Debugf("EapAuthConfirm Error[Problem Detail: %+v]", problemDetails)
			return nil
		}

		switch response.AuthResult {
		case models.AuthResult_SUCCESS:
			ue.UnauthenticatedSupi = false
			ausfinfo.Kseaf = response.KSeaf
			ue.Supi = response.Supi
			ue.AusfClient().DerivateKamf()
			// TODO: select enc/int algorithm based on ue security capability & amf's policy,
			// then generate KnasEnc, KnasInt
			return gmm.sm.SendEvent(ue.State[accessType], AuthSuccessEvent, fsm.ArgsType{
				ArgAmfUe:      ue,
				ArgAccessType: accessType,
				ArgEAPSuccess: true,
				ArgEAPMessage: response.EapPayload,
			})
		case models.AuthResult_FAILURE:
			if ue.IdentityTypeUsedForRegistration == nasMessage.MobileIdentity5GSType5gGuti && ausfinfo.IdentityRequestSendTimes == 0 {
				ausfinfo.IdentityRequestSendTimes++
				gmm.nas.SendAuthenticationResult(ue.RanUe[accessType], false, response.EapPayload)
				gmm.nas.SendIdentityRequest(ue.RanUe[accessType], accessType, nasMessage.MobileIdentity5GSTypeSuci)
				return nil
			} else {
				gmm.nas.SendAuthenticationReject(ue.RanUe[accessType], response.EapPayload)
				return gmm.sm.SendEvent(ue.State[accessType], AuthFailEvent, fsm.ArgsType{
					ArgAmfUe:      ue,
					ArgAccessType: accessType,
				})
			}
		case models.AuthResult_ONGOING:
			ausfinfo.AuthenticationCtx.Var5gAuthData = response.EapPayload
			if _, exists := response.Links["eap-session"]; exists {
				ausfinfo.AuthenticationCtx.Links = response.Links
			}
			gmm.nas.SendAuthenticationRequest(ue.RanUe[accessType])
		}
	}

	return nil
}

func (gmm *GmmFsm) handleAuthenticationError(ue *context.AmfUe, anType models.AccessType) error {
	log.Info("Handle Authentication Error")
	if ue.RegistrationRequest != nil {
		gmm.nas.SendRegistrationReject(ue.RanUe[anType], nasMessage.Cause5GMMTrackingAreaNotAllowed, "")
	}

	return nil
}

func (gmm *GmmFsm) handleAuthenticationFailure(ue *context.AmfUe, anType models.AccessType,
	authenticationFailure *nasMessage.AuthenticationFailure) error {
	log.Info("Handle Authentication Failure")

	ausfinfo := ue.AusfClient().Info()
	secinfo := ue.AusfClient().SecInfo()

	if ue.T3560 != nil {
		ue.T3560.Stop()
		ue.T3560 = nil // clear the timer
	}

	cause5GMM := authenticationFailure.Cause5GMM.GetCauseValue()

	if ausfinfo.AuthenticationCtx.AuthType == models.AuthType__5_G_AKA {
		switch cause5GMM {
		case nasMessage.Cause5GMMMACFailure:
			log.Warnln("Authentication Failure Cause: Mac Failure")
			gmm.nas.SendAuthenticationReject(ue.RanUe[anType], "")
			return gmm.sm.SendEvent(ue.State[anType], AuthFailEvent, fsm.ArgsType{ArgAmfUe: ue, ArgAccessType: anType})
		case nasMessage.Cause5GMMNon5GAuthenticationUnacceptable:
			log.Warnln("Authentication Failure Cause: Non-5G Authentication Unacceptable")
			gmm.nas.SendAuthenticationReject(ue.RanUe[anType], "")
			return gmm.sm.SendEvent(ue.State[anType], AuthFailEvent, fsm.ArgsType{ArgAmfUe: ue, ArgAccessType: anType})
		case nasMessage.Cause5GMMngKSIAlreadyInUse:
			log.Warnln("Authentication Failure Cause: NgKSI Already In Use")
			ausfinfo.AuthFailureCauseSynchFailureTimes = 0
			log.Warnln("Select new NgKsi")
			// select new ngksi
			if secinfo.NgKsi.Ksi < 6 { // ksi is range from 0 to 6
				secinfo.NgKsi.Ksi += 1
			} else {
				secinfo.NgKsi.Ksi = 0
			}
			gmm.nas.SendAuthenticationRequest(ue.RanUe[anType])
		case nasMessage.Cause5GMMSynchFailure: // TS 24.501 5.4.1.3.7 case f
			log.Warn("Authentication Failure 5GMM Cause: Synch Failure")

			ausfinfo.AuthFailureCauseSynchFailureTimes++
			if ausfinfo.AuthFailureCauseSynchFailureTimes >= 2 {
				log.Warnf("2 consecutive Synch Failure, terminate authentication procedure")
				gmm.nas.SendAuthenticationReject(ue.RanUe[anType], "")
				return gmm.sm.SendEvent(ue.State[anType], AuthFailEvent, fsm.ArgsType{ArgAmfUe: ue, ArgAccessType: anType})
			}

			auts := authenticationFailure.AuthenticationFailureParameter.GetAuthenticationFailureParameter()
			resynchronizationInfo := &models.ResynchronizationInfo{
				Auts: hex.EncodeToString(auts[:]),
			}

			var av5gAka models.Av5gAka
			if err := mapstructure.Decode(ausfinfo.AuthenticationCtx.Var5gAuthData, &av5gAka); err != nil {
				log.Error("Var5gAuthData Convert Type Error")
				return err
			}
			resynchronizationInfo.Rand = av5gAka.Rand

			response, problemDetails, err := ue.AusfClient().SendUEAuthRequest(resynchronizationInfo)
			if err != nil {
				return err
			} else if problemDetails != nil {
				log.Errorf("Nausf_UEAU Authenticate Request Error[Problem Detail: %+v]", problemDetails)
				return nil
			}
			ausfinfo.AuthenticationCtx = response
			ausfinfo.ABBA = []uint8{0x00, 0x00}

			gmm.nas.SendAuthenticationRequest(ue.RanUe[anType])
		}
	} else if ausfinfo.AuthenticationCtx.AuthType == models.AuthType_EAP_AKA_PRIME {
		switch cause5GMM {
		case nasMessage.Cause5GMMngKSIAlreadyInUse:
			log.Warn("Authentication Failure 5GMM Cause: NgKSI Already In Use")
			if secinfo.NgKsi.Ksi < 6 { // ksi is range from 0 to 6
				secinfo.NgKsi.Ksi += 1
			} else {
				secinfo.NgKsi.Ksi = 0
			}
			gmm.nas.SendAuthenticationRequest(ue.RanUe[anType])
		}
	}

	return nil
}

func (gmm *GmmFsm) handleRegistrationComplete(ue *context.AmfUe, accessType models.AccessType,
	registrationComplete *nasMessage.RegistrationComplete) error {
	log.Info("Handle Registration Complete")

	if ue.T3550 != nil {
		ue.T3550.Stop()
		ue.T3550 = nil // clear the timer
	}

	// if registrationComplete.SORTransparentContainer != nil {
	// 	TODO: if at regsitration procedure 14b, udm provide amf Steering of Roaming info & request an ack,
	// 	AMF provides the UE's ack with Nudm_SDM_Info (SOR not supportted in this stage)
	// }

	// TODO: if
	//	1. AMF has evaluated the support of IMS Voice over PS Sessions (TS 23.501 5.16.3.2)
	//	2. AMF determines that it needs to update the Homogeneous Support of IMS Voice over PS Sessions (TS 23.501 5.16.3.3)
	// Then invoke Nudm_UECM_Update to send "Homogeneous Support of IMS Voice over PS Sessions" indication to udm

	if ue.RegistrationRequest.UplinkDataStatus == nil &&
		ue.RegistrationRequest.GetFOR() == nasMessage.FollowOnRequestNoPending {
		gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[accessType], context.UeContextN2NormalRelease,
			ngapType.CausePresentNas, ngapType.CauseNasPresentNormalRelease)
	}
	return gmm.sm.SendEvent(ue.State[accessType], ContextSetupSuccessEvent, fsm.ArgsType{
		ArgAmfUe:      ue,
		ArgAccessType: accessType,
	})
}

// TS 33.501 6.7.2
func (gmm *GmmFsm) handleSecurityModeComplete(ue *context.AmfUe, anType models.AccessType, procedureCode int64,
	securityModeComplete *nasMessage.SecurityModeComplete) error {
	log.Info("Handle Security Mode Complete")

	if ue.AusfClient().IsMacFailed() {
		return fmt.Errorf("NAS message integrity check failed")
	}

	if ue.T3560 != nil {
		ue.T3560.Stop()
		ue.T3560 = nil // clear the timer
	}

	if ue.AusfClient().SecurityContextIsValid() {
		// update Kgnb/Kn3iwf
		ue.AusfClient().UpdateSecurityContext(anType)
	}

	if securityModeComplete.IMEISV != nil {
		log.Debugln("receieve IMEISV")
		ue.Pei = nasConvert.PeiToString(securityModeComplete.IMEISV.Octet[:])
	}

	// TODO: AMF shall set the NAS COUNTs to zero if horizontal derivation of KAMF is performed
	if securityModeComplete.NASMessageContainer != nil {
		contents := securityModeComplete.NASMessageContainer.GetNASMessageContainerContents()
		m := libnas.NewMessage()
		if err := m.GmmMessageDecode(&contents); err != nil {
			return err
		}

		messageType := m.GmmMessage.GmmHeader.GetMessageType()
		if messageType != libnas.MsgTypeRegistrationRequest && messageType != libnas.MsgTypeServiceRequest {
			log.Errorln("nas message container Iei type error")
			return errors.New("nas message container Iei type error")
		} else {
			return gmm.sm.SendEvent(ue.State[anType], SecurityModeSuccessEvent, fsm.ArgsType{
				ArgAmfUe:         ue,
				ArgAccessType:    anType,
				ArgProcedureCode: procedureCode,
				ArgNASMessage:    m.GmmMessage.RegistrationRequest,
			})
		}
	}
	return gmm.sm.SendEvent(ue.State[anType], SecurityModeSuccessEvent, fsm.ArgsType{
		ArgAmfUe:         ue,
		ArgAccessType:    anType,
		ArgProcedureCode: procedureCode,
		ArgNASMessage:    ue.RegistrationRequest,
	})
}

func (gmm *GmmFsm) handleSecurityModeReject(ue *context.AmfUe, anType models.AccessType,
	securityModeReject *nasMessage.SecurityModeReject) error {
	log.Info("Handle Security Mode Reject")

	if ue.T3560 != nil {
		ue.T3560.Stop()
		ue.T3560 = nil // clear the timer
	}

	cause := securityModeReject.Cause5GMM.GetCauseValue()
	log.Warnf("Reject Cause: %s", nasMessage.Cause5GMMToString(cause))
	log.Error("UE reject the security mode command, abort the ongoing procedure")
	return nil
}

// TS 23.502 4.2.2.3
func (gmm *GmmFsm) handleDeregistrationRequest(ue *context.AmfUe, anType models.AccessType,
	deregistrationRequest *nasMessage.DeregistrationRequestUEOriginatingDeregistration) error {
	log.Info("Handle Deregistration Request(UE Originating)")

	targetDeregistrationAccessType := deregistrationRequest.GetAccessType()
	ue.SmContextList.Range(func(key, value interface{}) bool {
		smContext := value.(*context.SmContext)

		if smContext.AccessType() == anType ||
			targetDeregistrationAccessType == nasMessage.AccessTypeBoth {
			problemDetail, err := smContext.SmfClient().SendReleaseSmContextRequest(nil, "", nil)
			if problemDetail != nil {
				log.Errorf("Release SmContext Failed Problem[%+v]", problemDetail)
			} else if err != nil {
				log.Errorf("Release SmContext Error[%v]", err.Error())
			}
		}
		return true
	})

	if ue.PcfClient().Info().AmPolicyAssociation != nil {
		terminateAmPolicyAssocaition := true
		switch anType {
		case models.AccessType__3_GPP_ACCESS:
			terminateAmPolicyAssocaition = ue.State[models.AccessType_NON_3_GPP_ACCESS].Is(context.Deregistered)
		case models.AccessType_NON_3_GPP_ACCESS:
			terminateAmPolicyAssocaition = ue.State[models.AccessType__3_GPP_ACCESS].Is(context.Deregistered)
		}

		if terminateAmPolicyAssocaition {
			problemDetails, err := ue.PcfClient().AMPolicyControlDelete()
			if problemDetails != nil {
				log.Errorf("AM Policy Control Delete Failed Problem[%+v]", problemDetails)
			} else if err != nil {
				log.Errorf("AM Policy Control Delete Error[%v]", err.Error())
			}
		}
	}

	// if Deregistration type is not switch-off, send Deregistration Accept
	if deregistrationRequest.GetSwitchOff() == 0 {
		gmm.nas.SendDeregistrationAccept(ue.RanUe[anType])
	}

	// TS 23.502 4.2.6, 4.12.3
	switch targetDeregistrationAccessType {
	case nasMessage.AccessType3GPP:
		if ue.RanUe[models.AccessType__3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
		return gmm.sm.SendEvent(ue.State[models.AccessType__3_GPP_ACCESS], DeregistrationAcceptEvent, fsm.ArgsType{
			ArgAmfUe:      ue,
			ArgAccessType: anType,
		})
	case nasMessage.AccessTypeNon3GPP:
		if ue.RanUe[models.AccessType_NON_3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType_NON_3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
		return gmm.sm.SendEvent(ue.State[models.AccessType_NON_3_GPP_ACCESS], DeregistrationAcceptEvent, fsm.ArgsType{
			ArgAmfUe:      ue,
			ArgAccessType: anType,
		})
	case nasMessage.AccessTypeBoth:
		if ue.RanUe[models.AccessType__3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
		if ue.RanUe[models.AccessType_NON_3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType_NON_3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}

		err := gmm.sm.SendEvent(ue.State[models.AccessType__3_GPP_ACCESS], DeregistrationAcceptEvent, fsm.ArgsType{
			ArgAmfUe:      ue,
			ArgAccessType: anType,
		})
		if err != nil {
			log.Errorln(err)
		}
		return gmm.sm.SendEvent(ue.State[models.AccessType_NON_3_GPP_ACCESS], DeregistrationAcceptEvent, fsm.ArgsType{
			ArgAmfUe:      ue,
			ArgAccessType: anType,
		})
	}

	return nil
}

// TS 23.502 4.2.2.3
func (gmm *GmmFsm) handleDeregistrationAccept(ue *context.AmfUe, anType models.AccessType,
	deregistrationAccept *nasMessage.DeregistrationAcceptUETerminatedDeregistration) error {
	log.Info("Handle Deregistration Accept(UE Terminated)")

	if ue.T3522 != nil {
		ue.T3522.Stop()
		ue.T3522 = nil // clear the timer
	}

	switch ue.DeregistrationTargetAccessType {
	case nasMessage.AccessType3GPP:
		if ue.RanUe[models.AccessType__3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
	case nasMessage.AccessTypeNon3GPP:
		if ue.RanUe[models.AccessType_NON_3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType_NON_3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
	case nasMessage.AccessTypeBoth:
		if ue.RanUe[models.AccessType__3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType__3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
		if ue.RanUe[models.AccessType_NON_3_GPP_ACCESS] != nil {
			gmm.nas.ngap.SendUEContextReleaseCommand(ue.RanUe[models.AccessType_NON_3_GPP_ACCESS],
				context.UeContextReleaseUeContext, ngapType.CausePresentNas, ngapType.CauseNasPresentDeregister)
		}
	}

	ue.DeregistrationTargetAccessType = 0
	return nil
}

func (gmm *GmmFsm) handleStatus5GMM(ue *context.AmfUe, anType models.AccessType, status5GMM *nasMessage.Status5GMM) error {
	log.Info("Handle Staus 5GMM")
	if ue.AusfClient().IsMacFailed() {
		return fmt.Errorf("NAS message integrity check failed")
	}

	cause := status5GMM.Cause5GMM.GetCauseValue()
	log.Errorf("Error condition [Cause Value: %s]", nasMessage.Cause5GMMToString(cause))
	return nil
}
