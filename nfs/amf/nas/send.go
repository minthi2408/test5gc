package nas

import (
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/ngap/util"
	"etri5gc/openapi/models"

	//	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/ngap/ngapType"
)

// backOffTimerUint = 7 means backoffTimer is null
func (sender *Nas) SendDLNASTransport(ue *context.RanUe, payloadContainerType uint8, nasPdu []byte,
	pduSessionId int32, cause uint8, backOffTimerUint *uint8, backOffTimer uint8) {

	amfUe := ue.AmfUe()
	log.Info("Send DL NAS Transport")

	var causePtr *uint8
	if cause != 0 {
		causePtr = &cause
	}
	nasMsg, err := sender.BuildDLNASTransport(amfUe, ue.Ran().AnType(), payloadContainerType, nasPdu,
		uint8(pduSessionId), causePtr, backOffTimerUint, backOffTimer)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

func (sender *Nas) SendNotification(ue *context.RanUe, nasMsg []byte) {

	//amfUe := ue.AmfUe
	//amfUe.GmmLog.Info("Send Notification")
	/*
		TODO: tungtq do it later
		if context.AMF_Self().T3565Cfg.Enable {
			cfg := context.AMF_Self().T3565Cfg
			amfUe.T3565 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
		log.Warnf("T3565 expires, retransmit Notification (retry: %d)", expireTimes)
				sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
			}, func() {
		log.Warnf("T3565 Expires %d times, abort notification procedure", cfg.MaxRetryTimes)
				if amfUe.OnGoing(models.AccessType__3_GPP_ACCESS).Procedure != context.OnGoingProcedureN2Handover {
					callback.SendN1N2TransferFailureNotification(amfUe, models.N1N2MessageTransferCause_UE_NOT_RESPONDING)
				}
				amfUe.T3565 = nil // clear the timer
			})
		}
	*/
}

func (sender *Nas) SendIdentityRequest(ue *context.RanUe, accessType models.AccessType, typeOfIdentity uint8) {

	amfUe := ue.AmfUe()
	log.Info("Send Identity Request")

	nasMsg, err := sender.BuildIdentityRequest(amfUe, accessType, typeOfIdentity)
	if err != nil {
		//amfUe.GmmLog.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)

	amfUe.RequestIdentityType = typeOfIdentity

	/*
		if context.AMF_Self().T3570Cfg.Enable {
			cfg := context.AMF_Self().T3570Cfg
			amfUe.T3570 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
				//amfUe.GmmLog.Warnf("T3570 expires, retransmit Identity Request (retry: %d)", expireTimes)
				sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
			}, func() {
				//amfUe.GmmLog.Warnf("T3570 Expires %d times, abort identification procedure & ongoing 5GMM procedure",
					//cfg.MaxRetryTimes)
					//TODO: tuntq - bad design
				//gmm_common.RemoveAmfUe(amfUe)
			})
		}
	*/
}

func (sender *Nas) SendAuthenticationRequest(ue *context.RanUe) {

	amfUe := ue.AmfUe()
	log.Infof("Send Authentication Request")

	if amfUe.AusfClient().Info().AuthenticationCtx == nil {
		log.Error("Authentication Context of UE is nil")
		return
	}

	nasMsg, err := sender.BuildAuthenticationRequest(amfUe)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)

	/*
		TODO: tungtq do it later
		if context.AMF_Self().T3560Cfg.Enable {
			cfg := context.AMF_Self().T3560Cfg
			amfUe.T3560 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
				log.Warnf("T3560 expires, retransmit Authentication Request (retry: %d)", expireTimes)
				sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
			}, func() {
				log.Warnf("T3560 Expires %d times, abort authentication procedure & ongoing 5GMM procedure",
			//		cfg.MaxRetryTimes)
			//TODO: tungtq - bad design
			//	gmm_common.RemoveAmfUe(amfUe)
			})
		}
	*/
}

func (sender *Nas) SendServiceAccept(ue *context.RanUe, anType models.AccessType, pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool, errPduSessionId, errCause []uint8) {

	amfUe := ue.AmfUe()
	//amfUe.GmmLog.Info("Send Service Accept")

	nasMsg, err := sender.BuildServiceAccept(amfUe, anType, pDUSessionStatus, reactivationResult,
		errPduSessionId, errCause)
	if err != nil {
		//amfUe.GmmLog.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

func (sender *Nas) SendConfigurationUpdateCommand(amfUe *context.AmfUe, accessType models.AccessType,
	networkSlicingIndication *nasType.NetworkSlicingIndication) {

	//amfUe.GmmLog.Info("Configuration Update Command")

	nasMsg, err := sender.BuildConfigurationUpdateCommand(amfUe, accessType, networkSlicingIndication)
	if err != nil {
		//amfUe.GmmLog.Error(err.Error())
		return
	}
	mobilityRestrictionList := util.BuildIEMobilityRestrictionList(amfUe)
	sender.ngap.SendDownlinkNasTransport(amfUe.RanUe[accessType], nasMsg, &mobilityRestrictionList)
	sender.ngap.SendDownlinkNasTransport(amfUe.RanUe[accessType], nasMsg, nil)
}

func (sender *Nas) SendAuthenticationReject(ue *context.RanUe, eapMsg string) {

	amfUe := ue.AmfUe()
	log.Info("Send Authentication Reject")

	nasMsg, err := sender.BuildAuthenticationReject(amfUe, eapMsg)
	if err != nil {
		//amfUe.GmmLog.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

func (sender *Nas) SendAuthenticationResult(ue *context.RanUe, eapSuccess bool, eapMsg string) {
	amfUe := ue.AmfUe()
	//amfUe.GmmLog.Info("Send Authentication Result")

	nasMsg, err := sender.BuildAuthenticationResult(amfUe, eapSuccess, eapMsg)
	if err != nil {
		//amfUe.GmmLog.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

func (sender *Nas) SendServiceReject(ue *context.RanUe, pDUSessionStatus *[16]bool, cause uint8) {

	//amfUe := ue.AmfUe
	log.Info("Send Service Reject")

	nasMsg, err := sender.BuildServiceReject(pDUSessionStatus, cause)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

// T3502: This IE may be included to indicate a value for timer T3502 during the initial registration
// eapMessage: if the REGISTRATION REJECT message is used to convey EAP-failure message
func (sender *Nas) SendRegistrationReject(ue *context.RanUe, cause5GMM uint8, eapMessage string) {

	amfUe := ue.AmfUe()
	log.Info("Send Registration Reject")

	nasMsg, err := sender.BuildRegistrationReject(amfUe, cause5GMM, eapMessage)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

// eapSuccess: only used when authType is EAP-AKA', set the value to false if authType is not EAP-AKA'
// eapMessage: only used when authType is EAP-AKA', set the value to "" if authType is not EAP-AKA'
func (sender *Nas) SendSecurityModeCommand(ue *context.RanUe, accessType models.AccessType, eapSuccess bool, eapMessage string) {

	amfUe := ue.AmfUe()
	//amfUe.GmmLog.Info("Send Security Mode Command")

	nasMsg, err := sender.BuildSecurityModeCommand(amfUe, accessType, eapSuccess, eapMessage)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)

	/*TODO tungtq - do it later
	if context.AMF_Self().T3560Cfg.Enable {
		cfg := context.AMF_Self().T3560Cfg
		amfUe.T3560 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
			amfUe.GmmLog.Warnf("T3560 expires, retransmit Security Mode Command (retry: %d)", expireTimes)
			sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
		}, func() {
			amfUe.GmmLog.Warnf("T3560 Expires %d times, abort security mode control procedure", cfg.MaxRetryTimes)
			gmm_common.RemoveAmfUe(amfUe)
		})
	}
	*/
}

func (sender *Nas) SendDeregistrationRequest(ue *context.RanUe, accessType uint8, reRegistrationRequired bool, cause5GMM uint8) {

	//amfUe := ue.AmfUe
	//amfUe.GmmLog.Info("Send Deregistration Request")

	nasMsg, err := sender.BuildDeregistrationRequest(ue, accessType, reRegistrationRequired, cause5GMM)
	if err != nil {
		//amfUe.GmmLog.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)

	/*
		TODO: tuntq - do it later
		if context.AMF_Self().T3522Cfg.Enable {
			cfg := context.AMF_Self().T3522Cfg
			amfUe.T3522 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
				amfUe.GmmLog.Warnf("T3522 expires, retransmit Deregistration Request (retry: %d)", expireTimes)
				sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
			}, func() {
				amfUe.GmmLog.Warnf("T3522 Expires %d times, abort deregistration procedure", cfg.MaxRetryTimes)
				amfUe.T3522 = nil // clear the timer
				if accessType == nasMessage.AccessType3GPP {
					amfUe.GmmLog.Warnln("UE accessType[3GPP] transfer to Deregistered state")
					amfUe.State[models.AccessType__3_GPP_ACCESS].Set(context.Deregistered)
				} else if accessType == nasMessage.AccessTypeNon3GPP {
					amfUe.GmmLog.Warnln("UE accessType[Non3GPP] transfer to Deregistered state")
					amfUe.State[models.AccessType_NON_3_GPP_ACCESS].Set(context.Deregistered)
				} else {
					amfUe.GmmLog.Warnln("UE accessType[3GPP] transfer to Deregistered state")
					amfUe.State[models.AccessType__3_GPP_ACCESS].Set(context.Deregistered)
					amfUe.GmmLog.Warnln("UE accessType[Non3GPP] transfer to Deregistered state")
					amfUe.State[models.AccessType_NON_3_GPP_ACCESS].Set(context.Deregistered)
				}
			})
		}
	*/
}

func (sender *Nas) SendDeregistrationAccept(ue *context.RanUe) {

	nasMsg, err := sender.BuildDeregistrationAccept()
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}

func (sender *Nas) SendRegistrationAccept(
	amfUe *context.AmfUe,
	anType models.AccessType,
	pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool,
	errPduSessionId, errCause []uint8,
	pduSessionResourceSetupList *ngapType.PDUSessionResourceSetupListCxtReq) {

	if amfUe.RanUe[anType] == nil {
		log.Error("SendRegistrationAccept: RanUe is nil")
		return
	}

	nasMsg, err := sender.BuildRegistrationAccept(amfUe, anType, pDUSessionStatus, reactivationResult, errPduSessionId, errCause)
	if err != nil {
		return
	}

	if amfUe.RanUe[anType].UeContextRequest() {
		sender.ngap.SendInitialContextSetupRequest(amfUe, anType, nasMsg, pduSessionResourceSetupList, nil, nil, nil)
	} else {
		sender.ngap.SendDownlinkNasTransport(amfUe.RanUe[models.AccessType__3_GPP_ACCESS], nasMsg, nil)
	}
	/*
		TODO: tungtq do it later
		if context.AMF_Self().T3550Cfg.Enable {
			cfg := context.AMF_Self().T3550Cfg
			amfUe.T3550 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
				if amfUe.RanUe[anType] == nil {
					amfUe.GmmLog.Warnf("[NAS] UE Context released, abort retransmission of Registration Accept")
					amfUe.T3550 = nil
				} else {
					amfUe.GmmLog.Warnf("T3550 expires, retransmit Registration Accept (retry: %d)", expireTimes)
					sender.ngap.SendDownlinkNasTransport(amfUe.RanUe[anType], nasMsg, nil)
				}
			}, func() {
				amfUe.GmmLog.Warnf("T3550 Expires %d times, abort retransmission of Registration Accept", cfg.MaxRetryTimes)
				amfUe.T3550 = nil // clear the timer
				// TS 24.501 5.5.1.2.8 case c, 5.5.1.3.8 case c
				amfUe.State[anType].Set(context.Registered)
				amfUe.ClearRegistrationRequestData(anType)
			})
		}
	*/
}

func (sender *Nas) SendStatus5GMM(ue *context.RanUe, cause uint8) {

	//amfUe := ue.AmfUe
	log.Info("Send Status 5GMM")

	nasMsg, err := sender.BuildStatus5GMM(cause)
	if err != nil {
		log.Error(err.Error())
		return
	}
	sender.ngap.SendDownlinkNasTransport(ue, nasMsg, nil)
}
