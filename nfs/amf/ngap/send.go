package ngap

import (
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf"
	//"etri5gc/nfs/amf/sbi/consumer"
	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

type ngapSender struct {
	backend		amf.Backend
}


func(s *ngapSender) SendToRan(ran *context.AmfRan, packet []byte) {

	if len(packet) == 0 {
		log.Error("packet len is 0")
		return
	}

	log.Debugf("Send NGAP message To Ran")

	if n, err := ran.Conn().Write(packet); err != nil {
		log.Errorf("Send error: %+v", err)
		return
	} else {
		log.Debugf("Write %d bytes", n)
	}
}

func(s *ngapSender) SendToRanUe(ue *context.RanUe, packet []byte) {
	s.SendToRan(ue.Ran(), packet)
}

func (s *ngapSender) NasSendToRan(ue *context.AmfUe, accessType models.AccessType, packet []byte) {
	ranUe := ue.RanUe[accessType]
	if ranUe == nil {
		log.Error("RanUe is nil")
		return
	}

	s.SendToRanUe(ranUe, packet)
}

func(s *ngapSender) SendNGSetupResponse(ran *context.AmfRan) {
	log.Info("Send NG-Setup response")

	pkt, err := s.buildNGSetupResponse()
	if err != nil {
		log.Errorf("Build NGSetupResponse failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

func(s *ngapSender) SendNGSetupFailure(ran *context.AmfRan, cause ngapType.Cause) {
	log.Info("Send NG-Setup failure")

	if cause.Present == ngapType.CausePresentNothing {
		log.Errorf("Cause present is nil")
		return
	}

	pkt, err := s.buildNGSetupFailure(cause)
	if err != nil {
		log.Errorf("Build NGSetupFailure failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// partOfNGInterface: if reset type is "reset all", set it to nil TS 38.413 9.2.6.11
func(s *ngapSender) SendNGReset(ran *context.AmfRan, cause ngapType.Cause,
	partOfNGInterface *ngapType.UEAssociatedLogicalNGConnectionList) {
	log.Info("Send NG Reset")

	pkt, err := s.buildNGReset(cause, partOfNGInterface)
	if err != nil {
		log.Errorf("Build NGReset failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

func(s *ngapSender) SendNGResetAcknowledge(ran *context.AmfRan, partOfNGInterface *ngapType.UEAssociatedLogicalNGConnectionList,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	log.Info("Send NG Reset Acknowledge")

	if partOfNGInterface != nil && len(partOfNGInterface.List) == 0 {
		log.Error("length of partOfNGInterface is 0")
		return
	}

	pkt, err := s.buildNGResetAcknowledge(partOfNGInterface, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build NGResetAcknowledge failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

func(s *ngapSender) SendDownlinkNasTransport(ue *context.RanUe, nasPdu []byte,
	mobilityRestrictionList *ngapType.MobilityRestrictionList) {

	log.Info("Send Downlink Nas Transport")

	if len(nasPdu) == 0 {
		log.Errorf("Send DownlinkNasTransport Error: nasPdu is nil")
	}

	pkt, err := s.buildDownlinkNasTransport(ue, nasPdu, mobilityRestrictionList)
	if err != nil {
		log.Errorf("Build DownlinkNasTransport failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendPDUSessionResourceReleaseCommand(ue *context.RanUe, nasPdu []byte,
	pduSessionResourceReleasedList ngapType.PDUSessionResourceToReleaseListRelCmd) {

	log.Info("Send PDU Session Resource Release Command")

	pkt, err := s.buildPDUSessionResourceReleaseCommand(ue, nasPdu, pduSessionResourceReleasedList)
	if err != nil {
		log.Errorf("Build PDUSessionResourceReleaseCommand failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendUEContextReleaseCommand(ue *context.RanUe, action context.RelAction, causePresent int, cause aper.Enumerated) {

	log.Info("Send UE Context Release Command")

	pkt, err := s.buildUEContextReleaseCommand(ue, causePresent, cause)
	if err != nil {
		log.Errorf("Build UEContextReleaseCommand failed : %s", err.Error())
		return
	}
	ue.SetReleaseAction(action)
	if ue.AmfUe != nil && ue.Ran != nil {
		ue.AmfUe().ReleaseCause[ue.Ran().AnType()] = &context.CauseAll{
			NgapCause: &models.NgApCause{
				Group: int32(causePresent),
				Value: int32(cause),
			},
		}
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendErrorIndication(ran *context.AmfRan, amfUeNgapId *ngapType.AMFUENGAPID, ranUeNgapId *ngapType.RANUENGAPID,
	cause *ngapType.Cause, criticalityDiagnostics *ngapType.CriticalityDiagnostics) {

	log.Info("Send Error Indication")

	var amfUeNgapIdValue *int64
	if amfUeNgapId != nil {
		amfUeNgapIdValue = &amfUeNgapId.Value
	}
	var ranUeNgapIdValue *int64
	if ranUeNgapId != nil {
		ranUeNgapIdValue = &ranUeNgapId.Value
	}

	pkt, err := s.buildErrorIndication(amfUeNgapIdValue, ranUeNgapIdValue, cause, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build ErrorIndication failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

func(s *ngapSender) SendUERadioCapabilityCheckRequest(ue *context.RanUe) {

	log.Info("Send UE Radio Capability Check Request")

	pkt, err := s.buildUERadioCapabilityCheckRequest(ue)
	if err != nil {
		log.Errorf("Build UERadioCapabilityCheckRequest failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendHandoverCancelAcknowledge(ue *context.RanUe, criticalityDiagnostics *ngapType.CriticalityDiagnostics) {

	log.Info("Send Handover Cancel Acknowledge")

	pkt, err := s.buildHandoverCancelAcknowledge(ue, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build HandoverCancelAcknowledge failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

// nasPDU: from nas layer
// pduSessionResourceSetupRequestList: provided by AMF, and transfer data is from SMF
func(s *ngapSender) SendPDUSessionResourceSetupRequest(ue *context.RanUe, nasPdu []byte,
	pduSessionResourceSetupRequestList ngapType.PDUSessionResourceSetupListSUReq) {

	log.Info("Send PDU Session Resource Setup Request")

	if len(pduSessionResourceSetupRequestList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildPDUSessionResourceSetupRequest(ue, nasPdu, pduSessionResourceSetupRequestList)
	if err != nil {
		log.Errorf("Build PDUSessionResourceSetupRequest failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

// pduSessionResourceModifyConfirmList: provided by AMF, and transfer data is return from SMF
// pduSessionResourceFailedToModifyList: provided by AMF, and transfer data is return from SMF
func(s *ngapSender) SendPDUSessionResourceModifyConfirm(
	ue *context.RanUe,
	pduSessionResourceModifyConfirmList ngapType.PDUSessionResourceModifyListModCfm,
	pduSessionResourceFailedToModifyList ngapType.PDUSessionResourceFailedToModifyListModCfm,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {

	log.Info("Send PDU Session Resource Modify Confirm")

	if len(pduSessionResourceModifyConfirmList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	if len(pduSessionResourceFailedToModifyList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildPDUSessionResourceModifyConfirm(ue, pduSessionResourceModifyConfirmList,
		pduSessionResourceFailedToModifyList, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build PDUSessionResourceModifyConfirm failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

// pduSessionResourceModifyRequestList: from SMF
func(s *ngapSender) SendPDUSessionResourceModifyRequest(ue *context.RanUe,
	pduSessionResourceModifyRequestList ngapType.PDUSessionResourceModifyListModReq) {

	log.Info("Send PDU Session Resource Modify Request")

	if len(pduSessionResourceModifyRequestList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildPDUSessionResourceModifyRequest(ue, pduSessionResourceModifyRequestList)
	if err != nil {
		log.Errorf("Build PDUSessionResourceModifyRequest failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendInitialContextSetupRequest(
	amfUe *context.AmfUe,
	anType models.AccessType,
	nasPdu []byte,
	pduSessionResourceSetupRequestList *ngapType.PDUSessionResourceSetupListCxtReq,
	rrcInactiveTransitionReportRequest *ngapType.RRCInactiveTransitionReportRequest,
	coreNetworkAssistanceInfo *ngapType.CoreNetworkAssistanceInformation,
	emergencyFallbackIndicator *ngapType.EmergencyFallbackIndicator) {

	log.Info("Send Initial Context Setup Request")

	if pduSessionResourceSetupRequestList != nil {
		if len(pduSessionResourceSetupRequestList.List) > context.MaxNumOfPDUSessions {
			//amfUe.RanUe[anType].Log.Error("Pdu List out of range")
			return
		}
	}

	pkt, err := s.buildInitialContextSetupRequest(amfUe, anType, nasPdu, pduSessionResourceSetupRequestList,
		rrcInactiveTransitionReportRequest, coreNetworkAssistanceInfo, emergencyFallbackIndicator)
	if err != nil {
		log.Errorf("Build InitialContextSetupRequest failed : %s", err.Error())
		return
	}
	amfUe.RanUe[anType].SetUeContextRequest(false)
	amfUe.RanUe[anType].SentInitialContextSetupRequest = true
	s.NasSendToRan(amfUe, anType, pkt)
}

func(s *ngapSender) SendUEContextModificationRequest(
	amfUe *context.AmfUe,
	anType models.AccessType,
	oldAmfUeNgapID *int64,
	rrcInactiveTransitionReportRequest *ngapType.RRCInactiveTransitionReportRequest,
	coreNetworkAssistanceInfo *ngapType.CoreNetworkAssistanceInformation,
	mobilityRestrictionList *ngapType.MobilityRestrictionList,
	emergencyFallbackIndicator *ngapType.EmergencyFallbackIndicator) {
	if amfUe == nil {
		log.Error("AmfUe is nil")
		return
	}

	log.Info("Send UE Context Modification Request")

	pkt, err := s.buildUEContextModificationRequest(amfUe, anType, oldAmfUeNgapID, rrcInactiveTransitionReportRequest,
		coreNetworkAssistanceInfo, mobilityRestrictionList, emergencyFallbackIndicator)
	if err != nil {
		log.Errorf("Build UEContextModificationRequest failed : %s", err.Error())
		return
	}
	s.NasSendToRan(amfUe, anType, pkt)
}

// pduSessionResourceHandoverList: provided by amf and transfer is return from smf
// pduSessionResourceToReleaseList: provided by amf and transfer is return from smf
// criticalityDiagnostics = criticalityDiagonstics IE in receiver node's error indication
// when received node can't comprehend the IE or missing IE
func(s *ngapSender) SendHandoverCommand(
	sourceUe *context.RanUe,
	pduSessionResourceHandoverList ngapType.PDUSessionResourceHandoverList,
	pduSessionResourceToReleaseList ngapType.PDUSessionResourceToReleaseListHOCmd,
	container ngapType.TargetToSourceTransparentContainer,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	if sourceUe == nil {
		log.Error("SourceUe is nil")
		return
	}

	log.Info("Send Handover Command")

	if len(pduSessionResourceHandoverList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	if len(pduSessionResourceToReleaseList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildHandoverCommand(sourceUe, pduSessionResourceHandoverList, pduSessionResourceToReleaseList,
		container, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build HandoverCommand failed : %s", err.Error())
		return
	}
	s.SendToRanUe(sourceUe, pkt)
}

// cause = initiate the Handover Cancel procedure with the appropriate value for the Cause IE.
// criticalityDiagnostics = criticalityDiagonstics IE in receiver node's error indication
// when received node can't comprehend the IE or missing IE
func(s *ngapSender) SendHandoverPreparationFailure(sourceUe *context.RanUe, cause ngapType.Cause,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	if sourceUe == nil {
		log.Error("SourceUe is nil")
		return
	}

	log.Info("Send Handover Preparation Failure")

	amfUe := sourceUe.AmfUe()
	if amfUe == nil {
		log.Error("amfUe is nil")
		return
	}
	amfUe.SetOnGoing(sourceUe.Ran().AnType(), &context.OnGoing{
		Procedure: context.OnGoingProcedureNothing,
	})
	pkt, err := s.buildHandoverPreparationFailure(sourceUe, cause, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build HandoverPreparationFailure failed : %s", err.Error())
		return
	}
	s.SendToRanUe(sourceUe, pkt)
}

/*The PGW-C+SMF (V-SMF in the case of home-routed roaming scenario only) sends
a Nsmf_PDUSession_CreateSMContext Response(N2 SM Information (PDU Session ID, cause code)) to the AMF.*/
// Cause is from SMF
// pduSessionResourceSetupList provided by AMF, and the transfer data is from SMF
// sourceToTargetTransparentContainer is received from S-RAN
// nsci: new security context indicator, if amfUe has updated security context, set nsci to true, otherwise set to false
// N2 handover in same AMF
func(s *ngapSender) SendHandoverRequest(sourceUe *context.RanUe, targetRan *context.AmfRan, cause ngapType.Cause,
	pduSessionResourceSetupListHOReq ngapType.PDUSessionResourceSetupListHOReq,
	sourceToTargetTransparentContainer ngapType.SourceToTargetTransparentContainer, nsci bool) {
	if sourceUe == nil {
		log.Error("sourceUe is nil")
		return
	}

	log.Info("Send Handover Request")

	amfUe := sourceUe.AmfUe
	if amfUe == nil {
		log.Error("amfUe is nil")
		return
	}
	if targetRan == nil {
		log.Error("targetRan is nil")
		return
	}

	if sourceUe.GetHandoverInfo().TargetUe != nil {
		log.Error("Handover Required Duplicated")
		return
	}

	if len(pduSessionResourceSetupListHOReq.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	if len(sourceToTargetTransparentContainer.Value) == 0 {
		log.Error("Source To Target TransparentContainer is nil")
		return
	}

	var targetUe *context.RanUe
	if targetUeTmp, err := targetRan.NewRanUe(context.RanUeNgapIdUnspecified); err != nil {
		log.Errorf("Create target UE error: %+v", err)
	} else {
		targetUe = targetUeTmp
	}

	log.Tracef("Source : AMF_UE_NGAP_ID[%d], RAN_UE_NGAP_ID[%d]", sourceUe.AmfUeNgapId, sourceUe.RanUeNgapId)
	log.Tracef("Target : AMF_UE_NGAP_ID[%d], RAN_UE_NGAP_ID[Unknown]", targetUe.AmfUeNgapId)
	context.AttachSourceUeTargetUe(sourceUe, targetUe)

	pkt, err := s.buildHandoverRequest(targetUe, cause, pduSessionResourceSetupListHOReq,
		sourceToTargetTransparentContainer, nsci)
	if err != nil {
		log.Errorf("Build HandoverRequest failed : %s", err.Error())
		return
	}
	s.SendToRanUe(targetUe, pkt)
}

// pduSessionResourceSwitchedList: provided by AMF, and the transfer data is from SMF
// pduSessionResourceReleasedList: provided by AMF, and the transfer data is from SMF
// newSecurityContextIndicator: if AMF has activated a new 5G NAS security context, set it to true,
// otherwise set to false
// coreNetworkAssistanceInformation: provided by AMF, based on collection of UE behaviour statistics
// and/or other available
// information about the expected UE behaviour. TS 23.501 5.4.6, 5.4.6.2
// rrcInactiveTransitionReportRequest: configured by amf
// criticalityDiagnostics: from received node when received not comprehended IE or missing IE
func(s *ngapSender) SendPathSwitchRequestAcknowledge(
	ue *context.RanUe,
	pduSessionResourceSwitchedList ngapType.PDUSessionResourceSwitchedList,
	pduSessionResourceReleasedList ngapType.PDUSessionResourceReleasedListPSAck,
	newSecurityContextIndicator bool,
	coreNetworkAssistanceInformation *ngapType.CoreNetworkAssistanceInformation,
	rrcInactiveTransitionReportRequest *ngapType.RRCInactiveTransitionReportRequest,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {

	log.Info("Send Path Switch Request Acknowledge")

	if len(pduSessionResourceSwitchedList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	if len(pduSessionResourceReleasedList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildPathSwitchRequestAcknowledge(ue, pduSessionResourceSwitchedList, pduSessionResourceReleasedList,
		newSecurityContextIndicator, coreNetworkAssistanceInformation, rrcInactiveTransitionReportRequest,
		criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build PathSwitchRequestAcknowledge failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

// pduSessionResourceReleasedList: provided by AMF, and the transfer data is from SMF
// criticalityDiagnostics: from received node when received not comprehended IE or missing IE
func(s *ngapSender) SendPathSwitchRequestFailure(
	ran *context.AmfRan,
	amfUeNgapId,
	ranUeNgapId int64,
	pduSessionResourceReleasedList *ngapType.PDUSessionResourceReleasedListPSFail,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	log.Info("Send Path Switch Request Failure")

	if pduSessionResourceReleasedList != nil && len(pduSessionResourceReleasedList.List) > context.MaxNumOfPDUSessions {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildPathSwitchRequestFailure(amfUeNgapId, ranUeNgapId, pduSessionResourceReleasedList,
		criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build PathSwitchRequestFailure failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// RanStatusTransferTransparentContainer from Uplink Ran Configuration Transfer
func(s *ngapSender) SendDownlinkRanStatusTransfer(ue *context.RanUe, container ngapType.RANStatusTransferTransparentContainer) {

	log.Info("Send Downlink Ran Status Transfer")

	if len(container.DRBsSubjectToStatusTransferList.List) > context.MaxNumOfDRBs {
		log.Error("Pdu List out of range")
		return
	}

	pkt, err := s.buildDownlinkRanStatusTransfer(ue, container)
	if err != nil {
		log.Errorf("Build DownlinkRanStatusTransfer failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

// anType indicate amfUe send this msg for which accessType
// Paging Priority: is included only if the AMF receives an Namf_Communication_N1N2MessageTransfer message
// with an ARP value associated with
// priority services (e.g., MPS, MCS), as configured by the operator. (TS 23.502 4.2.3.3, TS 23.501 5.22.3)
// pagingOriginNon3GPP: TS 23.502 4.2.3.3 step 4b: If the UE is simultaneously registered over 3GPP and non-3GPP
// accesses in the same PLMN,
// the UE is in CM-IDLE state in both 3GPP access and non-3GPP access, and the PDU Session ID in step 3a
// is associated with non-3GPP access, the AMF sends a Paging message with associated access "non-3GPP" to
// NG-RAN node(s) via 3GPP access.
// more paging policy with 3gpp/non-3gpp access is described in TS 23.501 5.6.8
func(s *ngapSender) SendPaging(ue *context.AmfUe, ngapBuf []byte) {
	// var pagingPriority *ngapType.PagingPriority
	
	// if ppi != nil {
	// pagingPriority = new(ngapType.PagingPriority)
	// pagingPriority.Value = aper.Enumerated(*ppi)
	// }
	// pkt, err := s.buildPaging(ue, pagingPriority, pagingOriginNon3GPP)
	// if err != nil {
	// 	ngaplog.Errorf("Build Paging failed : %s", err.Error())
	// }
	amf := s.backend.Context()
	ranpool := 	amf.AmfRanPool()
	callback := s.backend.Consumer().Callback()
	taiList := ue.RegistrationArea[models.AccessType__3_GPP_ACCESS]
	ranpool.Range(func(key, value interface{}) bool {
		ran := value.(*context.AmfRan)
		for _, item := range ran.SupportedTAList() {
			if context.InTaiList(item.Tai, taiList) {
				log.Infof("Send Paging to TAI(%+v, Tac:%+v)",
					item.Tai.PlmnId, item.Tai.Tac)
				s.SendToRan(ran, ngapBuf)
				break
			}
		}
		return true
	})

	cfg := amf.T3513Cfg()
	if cfg.Enable {
		ue.T3513 = context.NewTimer(cfg.ExpireTime, cfg.MaxRetryTimes, func(expireTimes int32) {
			log.Warnf("T3513 expires, retransmit Paging (retry: %d)", expireTimes)
			ranpool.Range(func(key, value interface{}) bool {
				ran := value.(*context.AmfRan)
				for _, item := range ran.SupportedTAList() {
					if context.InTaiList(item.Tai, taiList) {
						s.SendToRan(ran, ngapBuf)
						break
					}
				}
				return true
			})
		}, func() {
			log.Warnf("T3513 expires %d times, abort paging procedure", cfg.MaxRetryTimes)
			ue.T3513 = nil // clear the timer
			if ue.OnGoing(models.AccessType__3_GPP_ACCESS).Procedure != context.OnGoingProcedureN2Handover {
				callback.SendN1N2TransferFailureNotification(ue, models.N1N2MessageTransferCause_UE_NOT_RESPONDING)
			}
		})
	}
}

// TS 23.502 4.2.2.2.3
// anType: indicate amfUe send this msg for which accessType
// amfUeNgapID: initial AMF get it from target AMF
// ngapMessage: initial UE Message to reroute
// allowedNSSAI: provided by AMF, and AMF get it from NSSF (4.2.2.2.3 step 4b)
func(s *ngapSender) SendRerouteNasRequest(ue *context.AmfUe, anType models.AccessType, amfUeNgapID *int64, ngapMessage []byte,
	allowedNSSAI *ngapType.AllowedNSSAI) {
	log.Info("Send Reroute Nas Request")

	if len(ngapMessage) == 0 {
		log.Error("Ngap Message is nil")
		return
	}

	pkt, err := s.buildRerouteNasRequest(ue, anType, amfUeNgapID, ngapMessage, allowedNSSAI)
	if err != nil {
		log.Errorf("Build RerouteNasRequest failed : %s", err.Error())
		return
	}
	s.NasSendToRan(ue, anType, pkt)
}

// criticality ->from received node when received node can't comprehend the IE or missing IE
func(s *ngapSender) SendRanConfigurationUpdateAcknowledge(
	ran *context.AmfRan, criticalityDiagnostics *ngapType.CriticalityDiagnostics) {


	log.Info("Send Ran Configuration Update Acknowledge")

	pkt, err := s.buildRanConfigurationUpdateAcknowledge(criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build RanConfigurationUpdateAcknowledge failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// criticality ->from received node when received node can't comprehend the IE or missing IE
// If the AMF cannot accept the update,
// it shall respond with a RAN CONFIGURATION UPDATE FAILURE message and appropriate cause value.
func(s *ngapSender) SendRanConfigurationUpdateFailure(ran *context.AmfRan, cause ngapType.Cause,
	criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	log.Info("Send Ran Configuration Update Failure")

	pkt, err := s.buildRanConfigurationUpdateFailure(cause, criticalityDiagnostics)
	if err != nil {
		log.Errorf("Build RanConfigurationUpdateFailure failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// An AMF shall be able to instruct other peer CP NFs, subscribed to receive such a notification,
// that it will be unavailable on this AMF and its corresponding target AMF(s).
// If CP NF does not subscribe to receive AMF unavailable notification, the CP NF may attempt
// forwarding the transaction towards the old AMF and detect that the AMF is unavailable. When
// it detects unavailable, it marks the AMF and its associated GUAMI(s) as unavailable.
// Defined in 23.501 5.21.2.2.2
func(s *ngapSender) SendAMFStatusIndication(ran *context.AmfRan, unavailableGUAMIList ngapType.UnavailableGUAMIList) {
	log.Info("Send AMF Status Indication")

	if len(unavailableGUAMIList.List) > context.MaxNumOfServedGuamiList {
		log.Error("GUAMI List out of range")
		return
	}

	pkt, err := s.buildAMFStatusIndication(unavailableGUAMIList)
	if err != nil {
		log.Errorf("Build AMFStatusIndication failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// TS 23.501 5.19.5.2
// amfOverloadResponse: the required behaviour of NG-RAN, provided by AMF
// amfTrafficLoadReductionIndication(int 1~99): indicates the percentage of the type, set to 0 if does not need this ie
// of traffic relative to the instantaneous incoming rate at the NG-RAN node, provided by AMF
// overloadStartNSSAIList: overload slices, provide by AMF
func(s *ngapSender) SendOverloadStart(
	ran *context.AmfRan,
	amfOverloadResponse *ngapType.OverloadResponse,
	amfTrafficLoadReductionIndication int64,
	overloadStartNSSAIList *ngapType.OverloadStartNSSAIList) {
	
	log.Info("Send Overload Start")

	if amfTrafficLoadReductionIndication != 0 &&
		(amfTrafficLoadReductionIndication < 1 || amfTrafficLoadReductionIndication > 99) {
		log.Error("AmfTrafficLoadReductionIndication out of range (should be 1 ~ 99)")
		return
	}

	if overloadStartNSSAIList != nil && len(overloadStartNSSAIList.List) > context.MaxNumOfSlice {
		log.Error("NSSAI List out of range")
		return
	}

	pkt, err := s.buildOverloadStart(amfOverloadResponse, amfTrafficLoadReductionIndication, overloadStartNSSAIList)
	if err != nil {
		log.Errorf("Build OverloadStart failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

func(s *ngapSender) SendOverloadStop(ran *context.AmfRan) {
	
	log.Info("Send Overload Stop")

	pkt, err := s.buildOverloadStop()
	if err != nil {
		log.Errorf("Build OverloadStop failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// SONConfigurationTransfer = sONConfigurationTransfer from uplink Ran Configuration Transfer
func(s *ngapSender) SendDownlinkRanConfigurationTransfer(ran *context.AmfRan, transfer *ngapType.SONConfigurationTransfer) {
	log.Info("Send Downlink Ran Configuration Transfer")

	pkt, err := s.buildDownlinkRanConfigurationTransfer(transfer)
	if err != nil {
		log.Errorf("Build DownlinkRanConfigurationTransfer failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// NRPPa PDU is by pass
// NRPPa PDU is from LMF define in 4.13.5.6
func(s *ngapSender) SendDownlinkNonUEAssociatedNRPPATransport(ue *context.RanUe, nRPPaPDU ngapType.NRPPaPDU) {
	
	log.Info("Send Downlink Non UE Associated NRPPA Transport")

	if len(nRPPaPDU.Value) == 0 {
		log.Error("length of NRPPA-PDU is 0")
		return
	}

	pkt, err := s.buildDownlinkNonUEAssociatedNRPPATransport(ue, nRPPaPDU)
	if err != nil {
		log.Errorf("Build DownlinkNonUEAssociatedNRPPATransport failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendDeactivateTrace(amfUe *context.AmfUe, anType models.AccessType) {
	
	ranUe := amfUe.RanUe[anType]
	if ranUe == nil {
		log.Error("RanUe is nil")
		return
	}

	log.Info("Send Deactivate Trace")

	pkt, err := s.buildDeactivateTrace(amfUe, anType)
	if err != nil {
		log.Errorf("Build DeactivateTrace failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ranUe, pkt)
}

// AOI List is from SMF
// The SMF may subscribe to the UE mobility event notification from the AMF
// (e.g. location reporting, UE moving into or out of Area Of Interest) TS 23.502 4.3.2.2.1 Step.17
// The Location Reporting Control message shall identify the UE for which reports are requested and may include
// Reporting Type, Location Reporting Level, Area Of Interest and Request Reference ID
// TS 23.502 4.10 LocationReportingProcedure
// The AMF may request the NG-RAN location reporting with event reporting type (e.g. UE location or UE presence
// in Area of Interest), reporting mode and its related parameters (e.g. number of reporting) TS 23.501 5.4.7
// Location Reference ID To Be Cancelled IE shall be present if the Event Type IE is set to "Stop UE presence
// in the area of interest". otherwise set it to 0
func(s *ngapSender) SendLocationReportingControl(
	ue *context.RanUe,
	AOIList *ngapType.AreaOfInterestList,
	LocationReportingReferenceIDToBeCancelled int64,
	eventType ngapType.EventType) {
	log.Info("Send Location Reporting Control")

	if AOIList != nil && len(AOIList.List) > context.MaxNumOfAOI {
		log.Error("AOI List out of range")
		return
	}

	if eventType.Value == ngapType.EventTypePresentStopUePresenceInAreaOfInterest {
		if LocationReportingReferenceIDToBeCancelled < 1 || LocationReportingReferenceIDToBeCancelled > 64 {
			log.Error("LocationReportingReferenceIDToBeCancelled out of range (should be 1 ~ 64)")
			return
		}
	}

	pkt, err := s.buildLocationReportingControl(ue, AOIList, LocationReportingReferenceIDToBeCancelled, eventType)
	if err != nil {
		log.Errorf("Build LocationReportingControl failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

func(s *ngapSender) SendUETNLABindingReleaseRequest(ue *context.RanUe) {
	
	log.Info("Send UE TNLA Binging Release Request")

	pkt, err := s.buildUETNLABindingReleaseRequest(ue)
	if err != nil {
		log.Errorf("Build UETNLABindingReleaseRequest failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}

// Weight Factor associated with each of the TNL association within the AMF
func(s *ngapSender) SendAMFConfigurationUpdate(ran *context.AmfRan, usage ngapType.TNLAssociationUsage,
	weightfactor ngapType.TNLAddressWeightFactor) {
	
	log.Info("Send AMF Configuration Update")

	pkt, err := s.buildAMFConfigurationUpdate(usage, weightfactor)
	if err != nil {
		log.Errorf("Build AMFConfigurationUpdate failed : %s", err.Error())
		return
	}
	s.SendToRan(ran, pkt)
}

// NRPPa PDU is a pdu from LMF to RAN defined in TS 23.502 4.13.5.5 step 3
// NRPPa PDU is by pass
func(s *ngapSender) SendDownlinkUEAssociatedNRPPaTransport(ue *context.RanUe, nRPPaPDU ngapType.NRPPaPDU) {
	log.Info("Send Downlink UE Associated NRPPa Transport")

	if len(nRPPaPDU.Value) == 0 {
		log.Error("length of NRPPA-PDU is 0")
		return
	}

	pkt, err := s.buildDownlinkUEAssociatedNRPPaTransport(ue, nRPPaPDU)
	if err != nil {
		log.Errorf("Build DownlinkUEAssociatedNRPPaTransport failed : %s", err.Error())
		return
	}
	s.SendToRanUe(ue, pkt)
}
