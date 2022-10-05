package context

import (
	"etri5gc/sbi/models"

	//libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

func (ue *AmfUe) BuildConfigurationUpdateCommand(msg *nasMessage.ConfigurationUpdateCommand, anType models.AccessType) {
	/*
		if ue.Guti != "" {
			gutiNas := nasConvert.GutiToNas(ue.Guti)
			msg.GUTI5G = &gutiNas
			msg.GUTI5G.SetIei(nasMessage.ConfigurationUpdateCommandGUTI5GType)
		}

		if len(ue.RegistrationArea[anType]) > 0 {
			msg.TAIList = nasType.NewTAIList(nasMessage.ConfigurationUpdateCommandTAIListType)
			taiListNas := nasConvert.TaiListToNas(ue.RegistrationArea[anType])
			msg.TAIList.SetLen(uint8(len(taiListNas)))
			msg.TAIList.SetPartialTrackingAreaIdentityList(taiListNas)
		}
		if len(ue.nssf.AllowedNssai[anType]) > 0 {
			msg.AllowedNSSAI =
				nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
			var buf []uint8
			for _, allowedSnssai := range ue.nssf.AllowedNssai[anType] {
				buf = append(buf, nasConvert.SnssaiToNas(*allowedSnssai.AllowedSnssai)...)
			}
			msg.AllowedNSSAI.SetLen(uint8(len(buf)))
			msg.AllowedNSSAI.SetSNSSAIValue(buf)
		}

		if len(ue.nssf.ConfiguredNssai) > 0 {
			msg.ConfiguredNSSAI =
				nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
			var buf []uint8
			for _, snssai := range ue.nssf.ConfiguredNssai {
				buf = append(buf, nasConvert.SnssaiToNas(*snssai.ConfiguredSnssai)...)
			}
			msg.ConfiguredNSSAI.SetLen(uint8(len(buf)))
			msg.ConfiguredNSSAI.SetSNSSAIValue(buf)
		}

		if ue.nssf.NetworkSliceInfo != nil {
			if len(ue.nssf.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 || len(ue.nssf.NetworkSliceInfo.RejectedNssaiInTa) != 0 {
				rejectedNssaiNas := nasConvert.RejectedNssaiToNas(
					ue.nssf.NetworkSliceInfo.RejectedNssaiInPlmn, ue.nssf.NetworkSliceInfo.RejectedNssaiInTa)
				msg.RejectedNSSAI = &rejectedNssaiNas
				msg.RejectedNSSAI.SetIei(nasMessage.ConfigurationUpdateCommandRejectedNSSAIType)
			}
		}
		pcfinfo := ue.pcfcli.Info()
		// TODO: UniversalTimeAndLocalTimeZone
		if anType == models.AccessType__3_GPP_ACCESS && pcfinfo.AmPolicyAssociation != nil &&
			pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
			msg.ServiceAreaList =
				nasType.NewServiceAreaList(nasMessage.ConfigurationUpdateCommandServiceAreaListType)
			partialServiceAreaList := nasConvert.PartialServiceAreaListToNas(ue.PlmnId, *pcfinfo.AmPolicyAssociation.ServAreaRes)
			msg.ServiceAreaList.SetLen(uint8(len(partialServiceAreaList)))
			msg.ServiceAreaList.SetPartialServiceAreaList(partialServiceAreaList)
		}

		netname := ue.amf.NetworkName()
		if netname.Full != "" {
			fullNetworkName := nasConvert.FullNetworkNameToNas(netname.Full)
			msg.FullNameForNetwork = &fullNetworkName
			msg.FullNameForNetwork.SetIei(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
		}

		if netname.Short != "" {
			shortNetworkName := nasConvert.ShortNetworkNameToNas(netname.Short)
			msg.ShortNameForNetwork = &shortNetworkName
			msg.ShortNameForNetwork.SetIei(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
		}
		if ue.loc.TimeZone != "" {
			localTimeZone := nasConvert.LocalTimeZoneToNas(ue.loc.TimeZone)
			localTimeZone.SetIei(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
			msg.LocalTimeZone =
				nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
			msg.LocalTimeZone = &localTimeZone
		}

		if ue.loc.TimeZone != "" {
			daylightSavingTime := nasConvert.DaylightSavingTimeToNas(ue.loc.TimeZone)
			daylightSavingTime.SetIei(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
			msg.NetworkDaylightSavingTime =
				nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
			msg.NetworkDaylightSavingTime = &daylightSavingTime
		}

		if len(ue.LadnInfo) > 0 {
			msg.LADNInformation =
				nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
			var buf []uint8
			for _, ladn := range ue.LadnInfo {
				ladnNas := nasConvert.LadnToNas(ladn.Dnn, ladn.TaiLists)
				buf = append(buf, ladnNas...)
			}
			msg.LADNInformation.SetLen(uint16(len(buf)))
			msg.LADNInformation.SetLADND(buf)
		}
	*/
	return
}

func (ue *AmfUe) BuildRegistrationAccept(msg *nasMessage.RegistrationAccept, anType models.AccessType) {
	/*
		msg.RegistrationResult5GS.SetLen(1)

		regresult := uint8(0)
		if anType == models.AccessType__3_GPP_ACCESS {
			regresult |= nasMessage.AccessType3GPP
			if ue.State[models.AccessType_NON_3_GPP_ACCESS].Is(Registered) {
				regresult |= nasMessage.AccessTypeNon3GPP
			}
		} else {
			regresult |= nasMessage.AccessTypeNon3GPP
			if ue.State[models.AccessType__3_GPP_ACCESS].Is(Registered) {
				regresult |= nasMessage.AccessType3GPP
			}
		}
		msg.RegistrationResult5GS.SetRegistrationResultValue5GS(regresult)
		// TODO: set smsAllowed value of RegistrationResult5GS if need

		if ue.Guti != "" {
			gutiNas := nasConvert.GutiToNas(ue.Guti)
			msg.GUTI5G = &gutiNas
			msg.GUTI5G.SetIei(nasMessage.RegistrationAcceptGUTI5GType)
		}

		plmnlist := ue.amf.PlmnSupportList()
		if len(plmnlist) > 1 {
			msg.EquivalentPlmns = nasType.NewEquivalentPlmns(nasMessage.RegistrationAcceptEquivalentPlmnsType)
			var buf []uint8
			for _, plmnSupportItem := range plmnlist {
				buf = append(buf, nasConvert.PlmnIDToNas(*plmnSupportItem.PlmnId)...)
			}
			msg.EquivalentPlmns.SetLen(uint8(len(buf)))
			copy(msg.EquivalentPlmns.Octet[:], buf)
		}

		if len(ue.RegistrationArea[anType]) > 0 {
			msg.TAIList = nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
			taiListNas := nasConvert.TaiListToNas(ue.RegistrationArea[anType])
			msg.TAIList.SetLen(uint8(len(taiListNas)))
			msg.TAIList.SetPartialTrackingAreaIdentityList(taiListNas)
		}

		if len(ue.nssf.AllowedNssai[anType]) > 0 {
			msg.AllowedNSSAI = nasType.NewAllowedNSSAI(nasMessage.RegistrationAcceptAllowedNSSAIType)
			var buf []uint8
			for _, allowedSnssai := range ue.nssf.AllowedNssai[anType] {
				buf = append(buf, nasConvert.SnssaiToNas(*allowedSnssai.AllowedSnssai)...)
			}
			msg.AllowedNSSAI.SetLen(uint8(len(buf)))
			msg.AllowedNSSAI.SetSNSSAIValue(buf)
		}

		if ue.nssf.NetworkSliceInfo != nil {
			if len(ue.nssf.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 || len(ue.nssf.NetworkSliceInfo.RejectedNssaiInTa) != 0 {
				rejectedNssaiNas := nasConvert.RejectedNssaiToNas(
					ue.nssf.NetworkSliceInfo.RejectedNssaiInPlmn, ue.nssf.NetworkSliceInfo.RejectedNssaiInTa)
				msg.RejectedNSSAI = &rejectedNssaiNas
				msg.RejectedNSSAI.SetIei(nasMessage.RegistrationAcceptRejectedNSSAIType)
			}
		}

		if ue.includeConfiguredNssaiCheck() {
			msg.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(nasMessage.RegistrationAcceptConfiguredNSSAIType)
			var buf []uint8
			for _, snssai := range ue.nssf.ConfiguredNssai {
				buf = append(buf, nasConvert.SnssaiToNas(*snssai.ConfiguredSnssai)...)
			}
			msg.ConfiguredNSSAI.SetLen(uint8(len(buf)))
			msg.ConfiguredNSSAI.SetSNSSAIValue(buf)
		}

		if ue.LadnInfo != nil {
			msg.LADNInformation = nasType.NewLADNInformation(nasMessage.RegistrationAcceptLADNInformationType)
			buf := make([]uint8, 0)
			for _, ladn := range ue.LadnInfo {
				ladnNas := nasConvert.LadnToNas(ladn.Dnn, ladn.TaiLists)
				buf = append(buf, ladnNas...)
			}
			msg.LADNInformation.SetLen(uint16(len(buf)))
			msg.LADNInformation.SetLADND(buf)
		}

		if ue.nssf.NetworkSlicingSubscriptionChanged {
			msg.NetworkSlicingIndication =
				nasType.NewNetworkSlicingIndication(nasMessage.RegistrationAcceptNetworkSlicingIndicationType)
			msg.NetworkSlicingIndication.SetNSSCI(1)
			msg.NetworkSlicingIndication.SetDCNI(0)
			ue.nssf.NetworkSlicingSubscriptionChanged = false // reset the value
		}
		pcfinfo := ue.pcfcli.Info()
		if anType == models.AccessType__3_GPP_ACCESS && pcfinfo.AmPolicyAssociation != nil &&
			pcfinfo.AmPolicyAssociation.ServAreaRes != nil {
			msg.ServiceAreaList = nasType.NewServiceAreaList(nasMessage.RegistrationAcceptServiceAreaListType)
			partialServiceAreaList := nasConvert.PartialServiceAreaListToNas(ue.PlmnId, *pcfinfo.AmPolicyAssociation.ServAreaRes)
			msg.ServiceAreaList.SetLen(uint8(len(partialServiceAreaList)))
			msg.ServiceAreaList.SetPartialServiceAreaList(partialServiceAreaList)
		}

		if anType == models.AccessType__3_GPP_ACCESS && ue.T3512Value != 0 {
			msg.T3512Value = nasType.NewT3512Value(nasMessage.RegistrationAcceptT3512ValueType)
			msg.T3512Value.SetLen(1)
			t3512 := nasConvert.GPRSTimer3ToNas(ue.T3512Value)
			msg.T3512Value.Octet = t3512
		}

		if anType == models.AccessType_NON_3_GPP_ACCESS {
			msg.Non3GppDeregistrationTimerValue =
				nasType.NewNon3GppDeregistrationTimerValue(nasMessage.RegistrationAcceptNon3GppDeregistrationTimerValueType)
			msg.Non3GppDeregistrationTimerValue.SetLen(1)
			timerValue := nasConvert.GPRSTimer2ToNas(ue.Non3gppDeregistrationTimerValue)
			msg.Non3GppDeregistrationTimerValue.SetGPRSTimer2Value(timerValue)
		}

		if ue.T3502Value != 0 {
			msg.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationAcceptT3502ValueType)
			msg.T3502Value.SetLen(1)
			t3502 := nasConvert.GPRSTimer2ToNas(ue.T3502Value)
			msg.T3502Value.SetGPRSTimer2Value(t3502)
		}

		if ue.UESpecificDRX != nasMessage.DRXValueNotSpecified {
			msg.NegotiatedDRXParameters =
				nasType.NewNegotiatedDRXParameters(nasMessage.RegistrationAcceptNegotiatedDRXParametersType)
			msg.NegotiatedDRXParameters.SetLen(1)
			msg.NegotiatedDRXParameters.SetDRXValue(ue.UESpecificDRX)
		}
	*/
}

func (ue *AmfUe) includeConfiguredNssaiCheck() bool {
	/*
		if len(ue.nssf.ConfiguredNssai) == 0 {
			return false
		}

		registrationRequest := ue.RegistrationRequest
		if registrationRequest.RequestedNSSAI == nil {
			return true
		}
		if ue.nssf.NetworkSliceInfo != nil && len(ue.nssf.NetworkSliceInfo.RejectedNssaiInPlmn) != 0 {
			return true
		}
		if registrationRequest.NetworkSlicingIndication != nil && registrationRequest.NetworkSlicingIndication.GetDCNI() == 1 {
			return true
		}
	*/
	return false
}
