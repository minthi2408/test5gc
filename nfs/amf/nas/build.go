package nas

import (
	"encoding/base64"

	"etrib5gc/nfs/amf/context"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils/nasConvert"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
)

func (builder *Nas) BuildDLNASTransport(ue *context.AmfUe, accessType models.AccessType, payloadContainerType uint8, nasPdu []byte,
	pduSessionId uint8, cause *uint8, backoffTimerUint *uint8, backoffTimer uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeDLNASTransport)

	m.SecurityHeader = libnas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	dLNASTransport := nasMessage.NewDLNASTransport(0)
	dLNASTransport.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	dLNASTransport.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	dLNASTransport.SetMessageType(libnas.MsgTypeDLNASTransport)
	dLNASTransport.SpareHalfOctetAndPayloadContainerType.SetPayloadContainerType(payloadContainerType)
	dLNASTransport.PayloadContainer.SetLen(uint16(len(nasPdu)))
	dLNASTransport.PayloadContainer.SetPayloadContainerContents(nasPdu)

	if pduSessionId != 0 {
		dLNASTransport.PduSessionID2Value = new(nasType.PduSessionID2Value)
		dLNASTransport.PduSessionID2Value.SetIei(nasMessage.DLNASTransportPduSessionID2ValueType)
		dLNASTransport.PduSessionID2Value.SetPduSessionID2Value(pduSessionId)
	}
	if cause != nil {
		dLNASTransport.Cause5GMM = new(nasType.Cause5GMM)
		dLNASTransport.Cause5GMM.SetIei(nasMessage.DLNASTransportCause5GMMType)
		dLNASTransport.Cause5GMM.SetCauseValue(*cause)
	}
	if backoffTimerUint != nil {
		dLNASTransport.BackoffTimerValue = new(nasType.BackoffTimerValue)
		dLNASTransport.BackoffTimerValue.SetIei(nasMessage.DLNASTransportBackoffTimerValueType)
		dLNASTransport.BackoffTimerValue.SetLen(1)
		dLNASTransport.BackoffTimerValue.SetUnitTimerValue(*backoffTimerUint)
		dLNASTransport.BackoffTimerValue.SetTimerValue(backoffTimer)
	}

	m.GmmMessage.DLNASTransport = dLNASTransport

	return ue.AusfClient().NasEncode(m, accessType)
}

func (builder *Nas) BuildNotification(ue *context.AmfUe, accessType models.AccessType) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeNotification)

	m.SecurityHeader = libnas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	notification := nasMessage.NewNotification(0)
	notification.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	notification.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	notification.SetMessageType(libnas.MsgTypeNotification)
	if accessType == models.ACCESSTYPE__3_GPP_ACCESS {
		notification.SetAccessType(nasMessage.AccessType3GPP)
	} else {
		notification.SetAccessType(nasMessage.AccessTypeNon3GPP)
	}

	m.GmmMessage.Notification = notification

	return ue.AusfClient().NasEncode(m, accessType)
}

func (builder *Nas) BuildIdentityRequest(ue *context.AmfUe, accessType models.AccessType, typeOfIdentity uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeIdentityRequest)

	if ue.AusfClient().SecurityContextAvailable() {
		m.SecurityHeader = libnas.SecurityHeader{
			ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
			SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		}
	}

	identityRequest := nasMessage.NewIdentityRequest(0)
	identityRequest.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	identityRequest.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	identityRequest.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	identityRequest.IdentityRequestMessageIdentity.SetMessageType(libnas.MsgTypeIdentityRequest)
	identityRequest.SpareHalfOctetAndIdentityType.SetTypeOfIdentity(typeOfIdentity)

	m.GmmMessage.IdentityRequest = identityRequest

	return ue.AusfClient().NasEncode(m, accessType)
}

func (builder *Nas) BuildAuthenticationRequest(ue *context.AmfUe) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeAuthenticationRequest)

	req := nasMessage.NewAuthenticationRequest(0)
	if err := ue.AusfClient().BuildAuthRequest(req); err != nil {
		return nil, err
	}
	m.GmmMessage.AuthenticationRequest = req

	return m.PlainNasEncode()
}

func (builder *Nas) BuildServiceAccept(ue *context.AmfUe, accessType models.AccessType, pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool, errPduSessionId, errCause []uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeServiceAccept)

	m.SecurityHeader = libnas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	serviceAccept := nasMessage.NewServiceAccept(0)
	serviceAccept.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	serviceAccept.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	serviceAccept.SetMessageType(libnas.MsgTypeServiceAccept)
	if pDUSessionStatus != nil {
		serviceAccept.PDUSessionStatus = new(nasType.PDUSessionStatus)
		serviceAccept.PDUSessionStatus.SetIei(nasMessage.ServiceAcceptPDUSessionStatusType)
		serviceAccept.PDUSessionStatus.SetLen(2)
		serviceAccept.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pDUSessionStatus)
	}
	if reactivationResult != nil {
		serviceAccept.PDUSessionReactivationResult = new(nasType.PDUSessionReactivationResult)
		serviceAccept.PDUSessionReactivationResult.SetIei(nasMessage.ServiceAcceptPDUSessionReactivationResultType)
		serviceAccept.PDUSessionReactivationResult.SetLen(2)
		serviceAccept.PDUSessionReactivationResult.Buffer = nasConvert.PSIToBuf(*reactivationResult)
	}
	if errPduSessionId != nil {
		serviceAccept.PDUSessionReactivationResultErrorCause = new(nasType.PDUSessionReactivationResultErrorCause)
		serviceAccept.PDUSessionReactivationResultErrorCause.SetIei(
			nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType)
		buf := nasConvert.PDUSessionReactivationResultErrorCauseToBuf(errPduSessionId, errCause)
		serviceAccept.PDUSessionReactivationResultErrorCause.SetLen(uint16(len(buf)))
		serviceAccept.PDUSessionReactivationResultErrorCause.Buffer = buf
	}
	m.GmmMessage.ServiceAccept = serviceAccept

	return ue.AusfClient().NasEncode(m, accessType)
}

func (builder *Nas) BuildAuthenticationReject(ue *context.AmfUe, eapMsg string) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeAuthenticationReject)

	authenticationReject := nasMessage.NewAuthenticationReject(0)
	authenticationReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	authenticationReject.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	authenticationReject.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	authenticationReject.AuthenticationRejectMessageIdentity.SetMessageType(libnas.MsgTypeAuthenticationReject)

	if eapMsg != "" {
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMsg)
		if err != nil {
			return nil, err
		}
		authenticationReject.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRejectEAPMessageType)
		authenticationReject.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		authenticationReject.EAPMessage.SetEAPMessage(rawEapMsg)
	}

	m.GmmMessage.AuthenticationReject = authenticationReject

	return m.PlainNasEncode()
}

func (builder *Nas) BuildAuthenticationResult(ue *context.AmfUe, eapSuccess bool, eapMsg string) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeAuthenticationResult)

	result := nasMessage.NewAuthenticationResult(0)
	if err := ue.AusfClient().BuildAuthResult(result, eapSuccess, eapMsg); err != nil {
		return nil, err
	}

	m.GmmMessage.AuthenticationResult = result

	return m.PlainNasEncode()
}

// T3346 Timer and EAP are not Supported
func (builder *Nas) BuildServiceReject(pDUSessionStatus *[16]bool, cause uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeServiceReject)

	serviceReject := nasMessage.NewServiceReject(0)
	serviceReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	serviceReject.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	serviceReject.SetMessageType(libnas.MsgTypeServiceReject)
	serviceReject.SetCauseValue(cause)
	if pDUSessionStatus != nil {
		serviceReject.PDUSessionStatus = new(nasType.PDUSessionStatus)
		serviceReject.PDUSessionStatus.SetIei(nasMessage.ServiceAcceptPDUSessionStatusType)
		serviceReject.PDUSessionStatus.SetLen(2)
		serviceReject.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pDUSessionStatus)
	}

	m.GmmMessage.ServiceReject = serviceReject

	return m.PlainNasEncode()
}

// T3346 timer are not supported
func (builder *Nas) BuildRegistrationReject(ue *context.AmfUe, cause5GMM uint8, eapMessage string) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeRegistrationReject)

	registrationReject := nasMessage.NewRegistrationReject(0)
	registrationReject.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	registrationReject.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	registrationReject.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	registrationReject.RegistrationRejectMessageIdentity.SetMessageType(libnas.MsgTypeRegistrationReject)
	registrationReject.Cause5GMM.SetCauseValue(cause5GMM)

	if ue.T3502Value != 0 {
		registrationReject.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationRejectT3502ValueType)
		registrationReject.T3502Value.SetLen(1)
		t3502 := nasConvert.GPRSTimer2ToNas(ue.T3502Value)
		registrationReject.T3502Value.SetGPRSTimer2Value(t3502)
	}

	if eapMessage != "" {
		registrationReject.EAPMessage = nasType.NewEAPMessage(nasMessage.RegistrationRejectEAPMessageType)
		rawEapMsg, err := base64.StdEncoding.DecodeString(eapMessage)
		if err != nil {
			return nil, err
		}
		registrationReject.EAPMessage.SetLen(uint16(len(rawEapMsg)))
		registrationReject.EAPMessage.SetEAPMessage(rawEapMsg)
	}

	m.GmmMessage.RegistrationReject = registrationReject

	return m.PlainNasEncode()
}

// TS 24.501 8.2.25
func (builder *Nas) BuildSecurityModeCommand(ue *context.AmfUe, accessType models.AccessType, eapSuccess bool, eapMessage string) (
	[]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeSecurityModeCommand)
	ausf := ue.AusfClient()
	m.SecurityHeader = libnas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext,
	}

	cmd := nasMessage.NewSecurityModeCommand(0)
	if err := ausf.BuildSecModeCmd(cmd, eapSuccess, eapMessage); err != nil {
		return nil, err
	}

	m.GmmMessage.SecurityModeCommand = cmd
	payload, err := ue.AusfClient().NasEncode(m, accessType)
	if err != nil {
		ausf.SetSecurityContextAvailable(false)
		return nil, err
	} else {
		ausf.SetSecurityContextAvailable(true)
		return payload, nil
	}
}

// T3346 timer are not supported
func (builder *Nas) BuildDeregistrationRequest(ue *context.RanUe, accessType uint8, reRegistrationRequired bool,
	cause5GMM uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeDeregistrationRequestUETerminatedDeregistration)

	deregistrationRequest := nasMessage.NewDeregistrationRequestUETerminatedDeregistration(0)
	deregistrationRequest.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	deregistrationRequest.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	deregistrationRequest.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	deregistrationRequest.SetMessageType(libnas.MsgTypeDeregistrationRequestUETerminatedDeregistration)

	deregistrationRequest.SetAccessType(accessType)
	deregistrationRequest.SetSwitchOff(0)
	if reRegistrationRequired {
		deregistrationRequest.SetReRegistrationRequired(nasMessage.ReRegistrationRequired)
	} else {
		deregistrationRequest.SetReRegistrationRequired(nasMessage.ReRegistrationNotRequired)
	}

	if cause5GMM != 0 {
		deregistrationRequest.Cause5GMM = nasType.NewCause5GMM(
			nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
		deregistrationRequest.Cause5GMM.SetCauseValue(cause5GMM)
	}
	m.GmmMessage.DeregistrationRequestUETerminatedDeregistration = deregistrationRequest

	if ue != nil && ue.AmfUe != nil {
		m.SecurityHeader = libnas.SecurityHeader{
			ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
			SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		}
		var anType models.AccessType
		if accessType == 0x01 {
			anType = models.ACCESSTYPE__3_GPP_ACCESS
		} else if accessType == 0x02 {
			anType = models.ACCESSTYPE_NON_3_GPP_ACCESS
		}
		return ue.AmfUe().AusfClient().NasEncode(m, anType)
	}
	return m.PlainNasEncode()
}

func (builder *Nas) BuildDeregistrationAccept() ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeDeregistrationAcceptUEOriginatingDeregistration)

	deregistrationAccept := nasMessage.NewDeregistrationAcceptUEOriginatingDeregistration(0)
	deregistrationAccept.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	deregistrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	deregistrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	deregistrationAccept.SetMessageType(libnas.MsgTypeDeregistrationAcceptUEOriginatingDeregistration)

	m.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration = deregistrationAccept

	return m.PlainNasEncode()
}

func (builder *Nas) BuildRegistrationAccept(
	ue *context.AmfUe,
	anType models.AccessType,
	pDUSessionStatus *[16]bool,
	reactivationResult *[16]bool,
	errPduSessionId, errCause []uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeRegistrationAccept)

	m.SecurityHeader = libnas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    libnas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
	}

	registrationAccept := nasMessage.NewRegistrationAccept(0)
	registrationAccept.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	registrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	registrationAccept.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	registrationAccept.RegistrationAcceptMessageIdentity.SetMessageType(libnas.MsgTypeRegistrationAccept)

	// 5gs network feature support
	cfg := builder.backend.Config()
	if cfg.Get5gsNwFeatSuppEnable() {
		registrationAccept.NetworkFeatureSupport5GS =
			nasType.NewNetworkFeatureSupport5GS(nasMessage.RegistrationAcceptNetworkFeatureSupport5GSType)
		registrationAccept.NetworkFeatureSupport5GS.SetLen(2)
		if anType == models.ACCESSTYPE__3_GPP_ACCESS {
			registrationAccept.SetIMSVoPS3GPP(cfg.Get5gsNwFeatSuppImsVoPS())
		} else {
			registrationAccept.SetIMSVoPSN3GPP(cfg.Get5gsNwFeatSuppImsVoPS())
		}
		registrationAccept.SetEMC(cfg.Get5gsNwFeatSuppEmc())
		registrationAccept.SetEMF(cfg.Get5gsNwFeatSuppEmf())
		registrationAccept.SetIWKN26(cfg.Get5gsNwFeatSuppIwkN26())
		registrationAccept.SetMPSI(cfg.Get5gsNwFeatSuppMpsi())
		registrationAccept.SetEMCN(cfg.Get5gsNwFeatSuppEmcN3())
		registrationAccept.SetMCSI(cfg.Get5gsNwFeatSuppMcsi())
	}

	if pDUSessionStatus != nil {
		registrationAccept.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.RegistrationAcceptPDUSessionStatusType)
		registrationAccept.PDUSessionStatus.SetLen(2)
		registrationAccept.PDUSessionStatus.Buffer = nasConvert.PSIToBuf(*pDUSessionStatus)
	}

	if reactivationResult != nil {
		registrationAccept.PDUSessionReactivationResult =
			nasType.NewPDUSessionReactivationResult(nasMessage.RegistrationAcceptPDUSessionReactivationResultType)
		registrationAccept.PDUSessionReactivationResult.SetLen(2)
		registrationAccept.PDUSessionReactivationResult.Buffer = nasConvert.PSIToBuf(*reactivationResult)
	}

	if errPduSessionId != nil {
		registrationAccept.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(
			nasMessage.RegistrationAcceptPDUSessionReactivationResultErrorCauseType)
		buf := nasConvert.PDUSessionReactivationResultErrorCauseToBuf(errPduSessionId, errCause)
		registrationAccept.PDUSessionReactivationResultErrorCause.SetLen(uint16(len(buf)))
		registrationAccept.PDUSessionReactivationResultErrorCause.Buffer = buf
	}

	ue.BuildRegistrationAccept(registrationAccept, anType)

	m.GmmMessage.RegistrationAccept = registrationAccept

	return ue.AusfClient().NasEncode(m, anType)
}

func (builder *Nas) BuildStatus5GMM(cause uint8) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeStatus5GMM)

	status5GMM := nasMessage.NewStatus5GMM(0)
	status5GMM.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	status5GMM.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	status5GMM.SetMessageType(libnas.MsgTypeStatus5GMM)
	status5GMM.SetCauseValue(cause)

	m.GmmMessage.Status5GMM = status5GMM

	return m.PlainNasEncode()
}

func (builder *Nas) BuildConfigurationUpdateCommand(ue *context.AmfUe, anType models.AccessType,
	networkSlicingIndication *nasType.NetworkSlicingIndication) ([]byte, error) {
	m := libnas.NewMessage()
	m.GmmMessage = libnas.NewGmmMessage()
	m.GmmHeader.SetMessageType(libnas.MsgTypeConfigurationUpdateCommand)

	configurationUpdateCommand := nasMessage.NewConfigurationUpdateCommand(0)
	configurationUpdateCommand.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	configurationUpdateCommand.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(libnas.SecurityHeaderTypePlainNas)
	configurationUpdateCommand.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0)
	configurationUpdateCommand.SetMessageType(libnas.MsgTypeConfigurationUpdateCommand)

	if ue.ConfigurationUpdateIndication.Octet != 0 {
		configurationUpdateCommand.ConfigurationUpdateIndication =
			nasType.NewConfigurationUpdateIndication(nasMessage.ConfigurationUpdateCommandConfigurationUpdateIndicationType)
		configurationUpdateCommand.ConfigurationUpdateIndication = &ue.ConfigurationUpdateIndication
	}

	if networkSlicingIndication != nil {
		configurationUpdateCommand.NetworkSlicingIndication =
			nasType.NewNetworkSlicingIndication(nasMessage.ConfigurationUpdateCommandNetworkSlicingIndicationType)
		configurationUpdateCommand.NetworkSlicingIndication = networkSlicingIndication
	}

	ue.BuildConfigurationUpdateCommand(configurationUpdateCommand, anType)

	m.GmmMessage.ConfigurationUpdateCommand = configurationUpdateCommand

	return m.PlainNasEncode()
}
