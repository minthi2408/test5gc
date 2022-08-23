package nas_security

//	"encoding/hex"

/*
func Encode(ue *context.AmfUe, msg *nas.Message, accessType models.AccessType) ([]byte, error) {
	if ue == nil {
		return nil, fmt.Errorf("amfUe is nil")
	}
	if msg == nil {
		return nil, fmt.Errorf("Nas Message is empty")
	}

	secinfo := ue.AusfClient().SecInfo()
	// Plain NAS message
	if !secinfo.SecurityContextAvailable {
		return msg.PlainNasEncode()
	} else {
		// Security protected NAS Message
		// a security protected NAS message must be integrity protected, and ciphering is optional
		needCiphering := false
		switch msg.SecurityHeader.SecurityHeaderType {
		case nas.SecurityHeaderTypeIntegrityProtected:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected")
		case nas.SecurityHeaderTypeIntegrityProtectedAndCiphered:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected And Ciphered")
			needCiphering = true
		case nas.SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected With New 5G Security Context")
			secinfo.ULCount.Set(0, 0)
			secinfo.DLCount.Set(0, 0)
		default:
			return nil, fmt.Errorf("Wrong security header type: 0x%0x", msg.SecurityHeader.SecurityHeaderType)
		}

		// encode plain nas first
		payload, err := msg.PlainNasEncode()
		if err != nil {
			return nil, fmt.Errorf("Plain NAS encode error: %+v", err)
		}

		//secinfo.NASLog.Tracef("plain payload:\n%+v", hex.Dump(payload))
		if needCiphering {
			//secinfo.NASLog.Debugf("Encrypt NAS message (algorithm: %+v, DLCount: 0x%0x)", secinfo.CipheringAlg, secinfo.DLCount.Get())
			//secinfo.NASLog.Tracef("NAS ciphering key: %0x", secinfo.KnasEnc)
			if err = security.NASEncrypt(secinfo.CipheringAlg, secinfo.KnasEnc, secinfo.DLCount.Get(),
				GetBearerType(accessType), security.DirectionDownlink, payload); err != nil {
				return nil, fmt.Errorf("Encrypt error: %+v", err)
			}
		}

		// add sequece number
		payload = append([]byte{secinfo.DLCount.SQN()}, payload[:]...)

		//secinfo.NASLog.Debugf("Calculate NAS MAC (algorithm: %+v, DLCount: 0x%0x)", secinfo.IntegrityAlg, secinfo.DLCount.Get())
		//secinfo.NASLog.Tracef("NAS integrity key: %0x", secinfo.KnasInt)
		mac32, err := security.NASMacCalculate(secinfo.IntegrityAlg, secinfo.KnasInt, secinfo.DLCount.Get(),
			GetBearerType(accessType), security.DirectionDownlink, payload)
		if err != nil {
			return nil, fmt.Errorf("MAC calcuate error: %+v", err)
		}
		// Add mac value
		//secinfo.NASLog.Tracef("MAC: 0x%08x", mac32)
		payload = append(mac32, payload[:]...)

		// Add EPD and Security Type
		msgSecurityHeader := []byte{msg.SecurityHeader.ProtocolDiscriminator, msg.SecurityHeader.SecurityHeaderType}
		payload = append(msgSecurityHeader, payload[:]...)

		// Increase DL Count
		secinfo.DLCount.AddOne()
		return payload, nil
	}
}

func Decode(ue *context.AmfUe, accessType models.AccessType, payload []byte) (*nas.Message, error) {
	if ue == nil {
		return nil, fmt.Errorf("amfUe is nil")
	}
	if payload == nil {
		return nil, fmt.Errorf("Nas payload is empty")
	}

	secinfo := ue.AusfClient().SecInfo()
	msg := new(nas.Message)
	msg.SecurityHeaderType = nas.GetSecurityHeaderType(payload) & 0x0f
	//secinfo.NASLog.Traceln("securityHeaderType is ", msg.SecurityHeaderType)
	if msg.SecurityHeaderType == nas.SecurityHeaderTypePlainNas {
		// RRCEstablishmentCause 0 is for emergency service
		if secinfo.SecurityContextAvailable && ue.RanUe[accessType].RRCEstablishmentCause != "0" {
			//secinfo.NASLog.Warnln("Received Plain NAS message")
			secinfo.MacFailed = false
			if err := msg.PlainNasDecode(&payload); err != nil {
				return nil, err
			}

			if msg.GmmMessage == nil {
				return nil, fmt.Errorf("Gmm Message is nil")
			}

			// TS 24.501 4.4.4.3: Except the messages listed below, no NAS signalling messages shall be processed
			// by the receiving 5GMM entity in the AMF or forwarded to the 5GSM entity, unless the secure exchange
			// of NAS messages has been established for the NAS signalling connection
			switch msg.GmmHeader.GetMessageType() {
			case nas.MsgTypeRegistrationRequest:
				return msg, nil
			case nas.MsgTypeIdentityResponse:
				return msg, nil
			case nas.MsgTypeAuthenticationResponse:
				return msg, nil
			case nas.MsgTypeAuthenticationFailure:
				return msg, nil
			case nas.MsgTypeSecurityModeReject:
				return msg, nil
			case nas.MsgTypeDeregistrationRequestUEOriginatingDeregistration:
				return msg, nil
			case nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration:
				return msg, nil
			default:
				return nil, fmt.Errorf(
					"UE can not send plain nas for non-emergency service when there is a valid security context")
			}
		} else {
			secinfo.MacFailed = false
			err := msg.PlainNasDecode(&payload)
			return msg, err
		}
	} else { // Security protected NAS message
		securityHeader := payload[0:6]
		//secinfo.NASLog.Traceln("securityHeader is ", securityHeader)
		sequenceNumber := payload[6]
		//secinfo.NASLog.Traceln("sequenceNumber", sequenceNumber)

		receivedMac32 := securityHeader[2:]
		// remove security Header except for sequece Number
		payload = payload[6:]

		// a security protected NAS message must be integrity protected, and ciphering is optional
		ciphered := false
		switch msg.SecurityHeaderType {
		case nas.SecurityHeaderTypeIntegrityProtected:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected")
		case nas.SecurityHeaderTypeIntegrityProtectedAndCiphered:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected And Ciphered")
			ciphered = true
		case nas.SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext:
			//secinfo.NASLog.Debugln("Security header type: Integrity Protected And Ciphered With New 5G Security Context")
			ciphered = true
			secinfo.ULCount.Set(0, 0)
		default:
			return nil, fmt.Errorf("Wrong security header type: 0x%0x", msg.SecurityHeader.SecurityHeaderType)
		}

		if secinfo.ULCount.SQN() > sequenceNumber {
			//secinfo.NASLog.Debugf("set ULCount overflow")
			secinfo.ULCount.SetOverflow(secinfo.ULCount.Overflow() + 1)
		}
		secinfo.ULCount.SetSQN(sequenceNumber)

		//secinfo.NASLog.Debugf("Calculate NAS MAC (algorithm: %+v, ULCount: 0x%0x)", secinfo.IntegrityAlg, secinfo.ULCount.Get())
		//secinfo.NASLog.Tracef("NAS integrity key0x: %0x", secinfo.KnasInt)
		mac32, err := security.NASMacCalculate(secinfo.IntegrityAlg, secinfo.KnasInt, secinfo.ULCount.Get(),
			GetBearerType(accessType), security.DirectionUplink, payload)
		if err != nil {
			return nil, fmt.Errorf("MAC calcuate error: %+v", err)
		}

		if !reflect.DeepEqual(mac32, receivedMac32) {
			//secinfo.NASLog.Warnf("NAS MAC verification failed(received: 0x%08x, expected: 0x%08x)", receivedMac32, mac32)
			secinfo.MacFailed = true
		} else {
			//secinfo.NASLog.Tracef("cmac value: 0x%08x", mac32)
			secinfo.MacFailed = false
		}

		if ciphered {
			//secinfo.NASLog.Debugf("Decrypt NAS message (algorithm: %+v, ULCount: 0x%0x)", secinfo.CipheringAlg, secinfo.ULCount.Get())
			//secinfo.NASLog.Tracef("NAS ciphering key: %0x", secinfo.KnasEnc)
			// decrypt payload without sequence number (payload[1])
			if err = security.NASEncrypt(secinfo.CipheringAlg, secinfo.KnasEnc, secinfo.ULCount.Get(), GetBearerType(accessType),
				security.DirectionUplink, payload[1:]); err != nil {
				return nil, fmt.Errorf("Encrypt error: %+v", err)
			}
		}

		// remove sequece Number
		payload = payload[1:]
		err = msg.PlainNasDecode(&payload)
		return msg, err
	}
}

func GetBearerType(accessType models.AccessType) uint8 {
	if accessType == models.AccessType__3_GPP_ACCESS {
		return security.Bearer3GPP
	} else if accessType == models.AccessType_NON_3_GPP_ACCESS {
		return security.BearerNon3GPP
	} else {
		return security.OnlyOneBearer
	}
}
*/
