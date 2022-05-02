package nas

import (
	"etri5gc/nfs/amf/context"

//	libnas "github.com/free5gc/nas"
//	"github.com/free5gc/nas/nasConvert"
//	"github.com/free5gc/nas/nasMessage"
//	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/fsm"
)

const (
	GmmMessageEvent           fsm.EventType = "Gmm Message"
	StartAuthEvent            fsm.EventType = "Start Authentication"
	AuthSuccessEvent          fsm.EventType = "Authentication Success"
	AuthErrorEvent            fsm.EventType = "Authentication Error"
	AuthRestartEvent          fsm.EventType = "Authentication Restart"
	AuthFailEvent             fsm.EventType = "Authentication Fail"
	SecurityModeSuccessEvent  fsm.EventType = "SecurityMode Success"
	SecurityModeFailEvent     fsm.EventType = "SecurityMode Fail"
	ContextSetupSuccessEvent  fsm.EventType = "ContextSetup Success"
	ContextSetupFailEvent     fsm.EventType = "ContextSetup Fail"
	InitDeregistrationEvent   fsm.EventType = "Initialize Deregistration"
	DeregistrationAcceptEvent fsm.EventType = "Deregistration Accept"
)

const (
	ArgAmfUe               string = "AMF Ue"
	ArgNASMessage          string = "NAS Message"
	ArgProcedureCode       string = "Procedure Code"
	ArgAccessType          string = "Access Type"
	ArgEAPSuccess          string = "EAP Success"
	ArgEAPMessage          string = "EAP Message"
	Arg3GPPDeregistered    string = "3GPP Deregistered"
	ArgNon3GPPDeregistered string = "Non3GPP Deregistered"
)

var transitions = fsm.Transitions{
	{Event: GmmMessageEvent, From: context.Deregistered, To: context.Deregistered},
	{Event: GmmMessageEvent, From: context.Authentication, To: context.Authentication},
	{Event: GmmMessageEvent, From: context.SecurityMode, To: context.SecurityMode},
	{Event: GmmMessageEvent, From: context.ContextSetup, To: context.ContextSetup},
	{Event: GmmMessageEvent, From: context.Registered, To: context.Registered},
	{Event: StartAuthEvent, From: context.Deregistered, To: context.Authentication},
	{Event: StartAuthEvent, From: context.Registered, To: context.Authentication},
	{Event: AuthRestartEvent, From: context.Authentication, To: context.Authentication},
	{Event: AuthSuccessEvent, From: context.Authentication, To: context.SecurityMode},
	{Event: AuthFailEvent, From: context.Authentication, To: context.Deregistered},
	{Event: AuthErrorEvent, From: context.Authentication, To: context.Deregistered},
	{Event: SecurityModeSuccessEvent, From: context.SecurityMode, To: context.ContextSetup},
	{Event: SecurityModeFailEvent, From: context.SecurityMode, To: context.Deregistered},
	{Event: ContextSetupSuccessEvent, From: context.ContextSetup, To: context.Registered},
	{Event: ContextSetupFailEvent, From: context.ContextSetup, To: context.Deregistered},
	{Event: InitDeregistrationEvent, From: context.Registered, To: context.DeregistrationInitiated},
	{Event: DeregistrationAcceptEvent, From: context.DeregistrationInitiated, To: context.Deregistered},
}


type GmmFsm struct {
	sm *fsm.FSM
	nas *Nas
}

func newGmmFsm(nas *Nas) *GmmFsm {
	gmm := &GmmFsm {
		nas:	nas,
	}
	callbacks := fsm.Callbacks {
		context.Deregistered:            gmm.DeRegistered,
		context.Authentication:          gmm.Authentication,
		context.SecurityMode:            gmm.SecurityMode,
		context.ContextSetup:            gmm.ContextSetup,
		context.Registered:              gmm.Registered,
		context.DeregistrationInitiated: gmm.DeregisteredInitiated,
	}
	//Note: make sure that transitions and states are well-defined. The program
	//will panic if an error is returned 
	gmm.sm, _ = fsm.NewFSM(transitions, callbacks)
	return gmm
}

func (gmm *GmmFsm) DeRegistered(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.ClearRegistrationRequestData(accessType)
		amfUe.GmmLog.Debugln("EntryEvent at GMM State[DeRegistered]")
	case GmmMessageEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		procedureCode := args[ArgProcedureCode].(int64)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("GmmMessageEvent at GMM State[DeRegistered]")
		switch gmmMessage.GetMessageType() {
		case nas.MsgTypeRegistrationRequest:
			if err := HandleRegistrationRequest(amfUe, accessType, procedureCode, gmmMessage.RegistrationRequest); err != nil {
				logger.GmmLog.Errorln(err)
			} else {
				if err := gmm.sm.SendEvent(state, StartAuthEvent, fsm.ArgsType{
					ArgAmfUe:         amfUe,
					ArgAccessType:    accessType,
					ArgProcedureCode: procedureCode,
				}); err != nil {
					logger.GmmLog.Errorln(err)
				}
			}
		// If UE that considers itself Registared and CM-IDLE throws a ServiceRequest
		case nas.MsgTypeServiceRequest:
			if err := HandleServiceRequest(amfUe, accessType, gmmMessage.ServiceRequest); err != nil {
				logger.GmmLog.Errorln(err)
			}
		default:
			amfUe.GmmLog.Errorf("state mismatch: receieve gmm message[message type 0x%0x] at %s state",
				gmmMessage.GetMessageType(), state.Current())
		}
	case StartAuthEvent:
		logger.GmmLog.Debugln(event)
	case fsm.ExitEvent:
		logger.GmmLog.Debugln(event)
	default:
		logger.GmmLog.Errorf("Unknown event [%+v]", event)
	}
}

func (gmm *GmmFsm) Registered(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	switch event {
	case fsm.EntryEvent:
		// clear stored registration request data for this registration
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.ClearRegistrationRequestData(accessType)
		log.Debugln("EntryEvent at GMM State[Registered]")
	case GmmMessageEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		procedureCode := args[ArgProcedureCode].(int64)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		log.Debugln("GmmMessageEvent at GMM State[Registered]")
		switch gmmMessage.GetMessageType() {
		// Mobility Registration update / Periodic Registration update
		case nas.MsgTypeRegistrationRequest:
			if err := gmm.handleRegistrationRequest(amfUe, accessType, procedureCode, gmmMessage.RegistrationRequest); err != nil {
				log.Errorln(err)
			} else {
				if err := gmm.sm.SendEvent(state, StartAuthEvent, fsm.ArgsType{
					ArgAmfUe:         amfUe,
					ArgAccessType:    accessType,
					ArgProcedureCode: procedureCode,
				}); err != nil {
					log.Errorln(err)
				}
			}
		case nas.MsgTypeULNASTransport:
			if err := gmm.handleULNASTransport(amfUe, accessType, gmmMessage.ULNASTransport); err != nil {
				log.Errorln(err)
			}
		case nas.MsgTypeConfigurationUpdateComplete:
			if err := gmm.handleConfigurationUpdateComplete(amfUe, gmmMessage.ConfigurationUpdateComplete); err != nil {
				log.Errorln(err)
			}
		case nas.MsgTypeServiceRequest:
			if err := gmm.handleServiceRequest(amfUe, accessType, gmmMessage.ServiceRequest); err != nil {
				log.Errorln(err)
			}
		case nas.MsgTypeNotificationResponse:
			if err := gmm.handleNotificationResponse(amfUe, gmmMessage.NotificationResponse); err != nil {
				log.Errorln(err)
			}
		case nas.MsgTypeDeregistrationRequestUEOriginatingDeregistration:
			if err := gmm.sm.SendEvent(state, InitDeregistrationEvent, fsm.ArgsType{
				ArgAmfUe:      amfUe,
				ArgAccessType: accessType,
				ArgNASMessage: gmmMessage,
			}); err != nil {
				log.Errorln(err)
			}
		case nas.MsgTypeStatus5GMM:
			if err := gmm.handleStatus5GMM(amfUe, accessType, gmmMessage.Status5GMM); err != nil {
				log.Errorln(err)
			}
		default:
			log.Errorf("state mismatch: receieve gmm message[message type 0x%0x] at %s state",
				gmmMessage.GetMessageType(), state.Current())
		}
	case StartAuthEvent:
		log.Debugln(event)
	case InitDeregistrationEvent:
		log.Debugln(event)
	case fsm.ExitEvent:
		log.Debugln(event)
	default:
		log.Errorf("Unknown event [%+v]", event)
	}
}

func (gmm *GmmFsm) Authentication(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	var amfUe *context.AmfUe
	switch event {
	case fsm.EntryEvent:
		amfUe = args[ArgAmfUe].(*context.AmfUe)
		log.Debugln("EntryEvent at GMM State[Authentication]")
		fallthrough
	case AuthRestartEvent:
		amfUe = args[ArgAmfUe].(*context.AmfUe)
		accessType := args[ArgAccessType].(models.AccessType)
		log.Debugln("AuthRestartEvent at GMM State[Authentication]")

		pass, err := AuthenticationProcedure(amfUe, accessType)
		if err != nil {
			if err := gmm.sm.SendEvent(state, AuthErrorEvent, fsm.ArgsType{
				ArgAmfUe:      amfUe,
				ArgAccessType: accessType,
			}); err != nil {
				log.Errorln(err)
			}
		}
		if pass {
			if err := gmm.sm.SendEvent(state, AuthSuccessEvent, fsm.ArgsType{
				ArgAmfUe:      amfUe,
				ArgAccessType: accessType,
			}); err != nil {
				logger.GmmLog.Errorln(err)
			}
		}
	case GmmMessageEvent:
		amfUe = args[ArgAmfUe].(*context.AmfUe)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("GmmMessageEvent at GMM State[Authentication]")

		switch gmmMessage.GetMessageType() {
		case nas.MsgTypeIdentityResponse:
			if err := HandleIdentityResponse(amfUe, gmmMessage.IdentityResponse); err != nil {
				logger.GmmLog.Errorln(err)
			} else {
				// update identity type used for reauthentication
				mobileIdentityContents := gmmMessage.IdentityResponse.MobileIdentity.GetMobileIdentityContents()
				amfUe.IdentityTypeUsedForRegistration = nasConvert.GetTypeOfIdentity(mobileIdentityContents[0])

				err := gmm.sm.SendEvent(state, AuthRestartEvent, fsm.ArgsType{ArgAmfUe: amfUe, ArgAccessType: accessType})
				if err != nil {
					logger.GmmLog.Errorln(err)
				}
			}
		case nas.MsgTypeAuthenticationResponse:
			if err := HandleAuthenticationResponse(amfUe, accessType, gmmMessage.AuthenticationResponse); err != nil {
				logger.GmmLog.Errorln(err)
			}
		case nas.MsgTypeAuthenticationFailure:
			if err := HandleAuthenticationFailure(amfUe, accessType, gmmMessage.AuthenticationFailure); err != nil {
				logger.GmmLog.Errorln(err)
			}
		case nas.MsgTypeStatus5GMM:
			if err := HandleStatus5GMM(amfUe, accessType, gmmMessage.Status5GMM); err != nil {
				logger.GmmLog.Errorln(err)
			}
		default:
			logger.GmmLog.Errorf("UE state mismatch: receieve gmm message[message type 0x%0x] at %s state",
				gmmMessage.GetMessageType(), state.Current())
		}
	case AuthSuccessEvent:
		logger.GmmLog.Debugln(event)
	case AuthErrorEvent:
		amfUe = args[ArgAmfUe].(*context.AmfUe)
		accessType := args[ArgAccessType].(models.AccessType)
		logger.GmmLog.Debugln(event)
		if err := HandleAuthenticationError(amfUe, accessType); err != nil {
			logger.GmmLog.Errorln(err)
		}
	case AuthFailEvent:
		logger.GmmLog.Debugln(event)
		logger.GmmLog.Warnln("Reject authentication")
	case fsm.ExitEvent:
		// clear authentication related data at exit
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		amfUe.GmmLog.Debugln(event)
		amfUe.AuthenticationCtx = nil
		amfUe.AuthFailureCauseSynchFailureTimes = 0
		amfUe.IdentityRequestSendTimes = 0
	default:
		logger.GmmLog.Errorf("Unknown event [%+v]", event)
	}
	*/
}

func (gmm *GmmFsm) SecurityMode(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	/*
	switch event {
	case fsm.EntryEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		accessType := args[ArgAccessType].(models.AccessType)
		// set log information
		amfUe.UpdateLogFields()

		amfUe.GmmLog.Debugln("EntryEvent at GMM State[SecurityMode]")
		if amfUe.SecurityContextIsValid() {
			amfUe.GmmLog.Debugln("UE has a valid security context - skip security mode control procedure")
			if err := gmm.sm.SendEvent(state, SecurityModeSuccessEvent, fsm.ArgsType{
				ArgAmfUe:      amfUe,
				ArgAccessType: accessType,
				ArgNASMessage: amfUe.RegistrationRequest,
			}); err != nil {
				logger.GmmLog.Errorln(err)
			}
		} else {
			eapSuccess := args[ArgEAPSuccess].(bool)
			eapMessage := args[ArgEAPMessage].(string)
			// Select enc/int algorithm based on ue security capability & amf's policy,
			amfSelf := context.AMF_Self()
			amfUe.SelectSecurityAlg(amfSelf.SecurityAlgorithm.IntegrityOrder, amfSelf.SecurityAlgorithm.CipheringOrder)
			// Generate KnasEnc, KnasInt
			amfUe.DerivateAlgKey()
			gmm_message.SendSecurityModeCommand(amfUe.RanUe[accessType], accessType, eapSuccess, eapMessage)
		}
	case GmmMessageEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		procedureCode := args[ArgProcedureCode].(int64)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("GmmMessageEvent to GMM State[SecurityMode]")
		switch gmmMessage.GetMessageType() {
		case nas.MsgTypeSecurityModeComplete:
			if err := HandleSecurityModeComplete(amfUe, accessType, procedureCode, gmmMessage.SecurityModeComplete); err != nil {
				logger.GmmLog.Errorln(err)
			}
		case nas.MsgTypeSecurityModeReject:
			if err := HandleSecurityModeReject(amfUe, accessType, gmmMessage.SecurityModeReject); err != nil {
				logger.GmmLog.Errorln(err)
			}
			err := gmm.sm.SendEvent(state, SecurityModeFailEvent, fsm.ArgsType{
				ArgAmfUe:      amfUe,
				ArgAccessType: accessType,
			})
			if err != nil {
				logger.GmmLog.Errorln(err)
			}
		case nas.MsgTypeStatus5GMM:
			if err := HandleStatus5GMM(amfUe, accessType, gmmMessage.Status5GMM); err != nil {
				logger.GmmLog.Errorln(err)
			}
		default:
			amfUe.GmmLog.Errorf("state mismatch: receieve gmm message[message type 0x%0x] at %s state",
				gmmMessage.GetMessageType(), state.Current())
		}
	case SecurityModeSuccessEvent:
		logger.GmmLog.Debugln(event)
	case SecurityModeFailEvent:
		logger.GmmLog.Debugln(event)
	case fsm.ExitEvent:
		logger.GmmLog.Debugln(event)
		return
	default:
		logger.GmmLog.Errorf("Unknown event [%+v]", event)
	}
	*/
}

func (gmm *GmmFsm) ContextSetup(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	/*
	switch event {
	case fsm.EntryEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		gmmMessage := args[ArgNASMessage]
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("EntryEvent at GMM State[ContextSetup]")

		switch message := gmmMessage.(type) {
		case *nasMessage.RegistrationRequest:
			amfUe.RegistrationRequest = message
			switch amfUe.RegistrationType5GS {
			case nasMessage.RegistrationType5GSInitialRegistration:
				if err := HandleInitialRegistration(amfUe, accessType); err != nil {
					logger.GmmLog.Errorln(err)
					err = gmm.sm.SendEvent(state, ContextSetupFailEvent, fsm.ArgsType{
						ArgAmfUe:      amfUe,
						ArgAccessType: accessType,
					})
					if err != nil {
						logger.GmmLog.Errorln(err)
					}
				}
			case nasMessage.RegistrationType5GSMobilityRegistrationUpdating:
				fallthrough
			case nasMessage.RegistrationType5GSPeriodicRegistrationUpdating:
				if err := HandleMobilityAndPeriodicRegistrationUpdating(amfUe, accessType); err != nil {
					logger.GmmLog.Errorln(err)
					err = gmm.sm.SendEvent(state, ContextSetupFailEvent, fsm.ArgsType{
						ArgAmfUe:      amfUe,
						ArgAccessType: accessType,
					})
					if err != nil {
						logger.GmmLog.Errorln(err)
					}
				}
			}
		case *nasMessage.ServiceRequest:
			if err := HandleServiceRequest(amfUe, accessType, message); err != nil {
				logger.GmmLog.Errorln(err)
			}
		default:
			logger.GmmLog.Errorf("UE state mismatch: receieve wrong gmm message")
		}
	case GmmMessageEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("GmmMessageEvent at GMM State[ContextSetup]")
		switch gmmMessage.GetMessageType() {
		case nas.MsgTypeIdentityResponse:
			if err := HandleIdentityResponse(amfUe, gmmMessage.IdentityResponse); err != nil {
				logger.GmmLog.Errorln(err)
			} else {
				switch amfUe.RegistrationType5GS {
				case nasMessage.RegistrationType5GSInitialRegistration:
					if err := HandleInitialRegistration(amfUe, accessType); err != nil {
						logger.GmmLog.Errorln(err)
						err = gmm.sm.SendEvent(state, ContextSetupFailEvent, fsm.ArgsType{
							ArgAmfUe:      amfUe,
							ArgAccessType: accessType,
						})
						if err != nil {
							logger.GmmLog.Errorln(err)
						}
					}
				case nasMessage.RegistrationType5GSMobilityRegistrationUpdating:
					fallthrough
				case nasMessage.RegistrationType5GSPeriodicRegistrationUpdating:
					if err := HandleMobilityAndPeriodicRegistrationUpdating(amfUe, accessType); err != nil {
						logger.GmmLog.Errorln(err)
						err = gmm.sm.SendEvent(state, ContextSetupFailEvent, fsm.ArgsType{
							ArgAmfUe:      amfUe,
							ArgAccessType: accessType,
						})
						if err != nil {
							logger.GmmLog.Errorln(err)
						}
					}
				}
			}
		case nas.MsgTypeRegistrationComplete:
			if err := HandleRegistrationComplete(amfUe, accessType, gmmMessage.RegistrationComplete); err != nil {
				logger.GmmLog.Errorln(err)
			}
		case nas.MsgTypeStatus5GMM:
			if err := HandleStatus5GMM(amfUe, accessType, gmmMessage.Status5GMM); err != nil {
				logger.GmmLog.Errorln(err)
			}
		default:
			amfUe.GmmLog.Errorf("state mismatch: receieve gmm message[message type 0x%0x] at %s state",
				gmmMessage.GetMessageType(), state.Current())
		}
	case ContextSetupSuccessEvent:
		logger.GmmLog.Debugln(event)
	case ContextSetupFailEvent:
		logger.GmmLog.Debugln(event)
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		accessType := args[ArgAccessType].(models.AccessType)
		if amfUe.UeCmRegistered {
			problemDetails, err := consumer.UeCmDeregistration(amfUe, accessType)
			if problemDetails != nil {
				if problemDetails.Cause != "CONTEXT_NOT_FOUND" {
					amfUe.GmmLog.Errorf("UECM_Registration Failed Problem[%+v]", problemDetails)
				}
			} else if err != nil {
				amfUe.GmmLog.Errorf("UECM_Registration Error[%+v]", err)
			}
		}
	case fsm.ExitEvent:
		logger.GmmLog.Debugln(event)
	default:
		logger.GmmLog.Errorf("Unknown event [%+v]", event)
	}
	*/
}

func (gmm *GmmFsm) DeregisteredInitiated(state *fsm.State, event fsm.EventType, args fsm.ArgsType) {
	/*
	switch event {
	case fsm.EntryEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("EntryEvent at GMM State[DeregisteredInitiated]")
		if err := HandleDeregistrationRequest(amfUe, accessType,
			gmmMessage.DeregistrationRequestUEOriginatingDeregistration); err != nil {
			logger.GmmLog.Errorln(err)
		}
	case GmmMessageEvent:
		amfUe := args[ArgAmfUe].(*context.AmfUe)
		gmmMessage := args[ArgNASMessage].(*nas.GmmMessage)
		accessType := args[ArgAccessType].(models.AccessType)
		amfUe.GmmLog.Debugln("GmmMessageEvent at GMM State[DeregisteredInitiated]")
		switch gmmMessage.GetMessageType() {
		case nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration:
			if err := HandleDeregistrationAccept(amfUe, accessType,
				gmmMessage.DeregistrationAcceptUETerminatedDeregistration); err != nil {
				logger.GmmLog.Errorln(err)
			}
		default:
			amfUe.GmmLog.Errorf("state mismatch: receieve gmm message[message type 0x%0x] at %s state",
				gmmMessage.GetMessageType(), state.Current())
		}
	case DeregistrationAcceptEvent:
		logger.GmmLog.Debugln(event)
	case fsm.ExitEvent:
		logger.GmmLog.Debugln(event)
	default:
		logger.GmmLog.Errorf("Unknown event [%+v]", event)
	}
	*/
}
