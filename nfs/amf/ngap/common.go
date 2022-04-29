package ngap

import (
//	"encoding/hex"
//	"strconv"
	"etri5gc/nfs/amf/context"
//	"etri5gc/nfs/amf/nas"

	"github.com/free5gc/aper"
//	"github.com/free5gc/nas/nasMessage"
//	libngap "github.com/free5gc/ngap"
//	"github.com/free5gc/ngap/ngapConvert"
	"github.com/free5gc/ngap/ngapType"
//	"github.com/free5gc/openapi/models"
)


func printAndGetCause(ran *context.AmfRan, cause *ngapType.Cause) (present int, value aper.Enumerated) {
	present = cause.Present
	switch cause.Present {
	case ngapType.CausePresentRadioNetwork:
	//	ran.Log.Warnf("Cause RadioNetwork[%d]", cause.RadioNetwork.Value)
		value = cause.RadioNetwork.Value
	case ngapType.CausePresentTransport:
	//	ran.Log.Warnf("Cause Transport[%d]", cause.Transport.Value)
		value = cause.Transport.Value
	case ngapType.CausePresentProtocol:
	//	ran.Log.Warnf("Cause Protocol[%d]", cause.Protocol.Value)
		value = cause.Protocol.Value
	case ngapType.CausePresentNas:
	//	ran.Log.Warnf("Cause Nas[%d]", cause.Nas.Value)
		value = cause.Nas.Value
	case ngapType.CausePresentMisc:
	//	ran.Log.Warnf("Cause Misc[%d]", cause.Misc.Value)
		value = cause.Misc.Value
	default:
	//	ran.Log.Errorf("Invalid Cause group[%d]", cause.Present)
	}
	return
}

func printCriticalityDiagnostics(ran *context.AmfRan, criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	/*
	ran.Log.Trace("Criticality Diagnostics")

	if criticalityDiagnostics.ProcedureCriticality != nil {
		switch criticalityDiagnostics.ProcedureCriticality.Value {
		case ngapType.CriticalityPresentReject:
			ran.Log.Trace("Procedure Criticality: Reject")
		case ngapType.CriticalityPresentIgnore:
			ran.Log.Trace("Procedure Criticality: Ignore")
		case ngapType.CriticalityPresentNotify:
			ran.Log.Trace("Procedure Criticality: Notify")
		}
	}

	if criticalityDiagnostics.IEsCriticalityDiagnostics != nil {
		for _, ieCriticalityDiagnostics := range criticalityDiagnostics.IEsCriticalityDiagnostics.List {
			ran.Log.Tracef("IE ID: %d", ieCriticalityDiagnostics.IEID.Value)

			switch ieCriticalityDiagnostics.IECriticality.Value {
			case ngapType.CriticalityPresentReject:
				ran.Log.Trace("Criticality Reject")
			case ngapType.CriticalityPresentNotify:
				ran.Log.Trace("Criticality Notify")
			}

			switch ieCriticalityDiagnostics.TypeOfError.Value {
			case ngapType.TypeOfErrorPresentNotUnderstood:
				ran.Log.Trace("Type of error: Not understood")
			case ngapType.TypeOfErrorPresentMissing:
				ran.Log.Trace("Type of error: Missing")
			}
		}
	}
	*/
}

func buildCriticalityDiagnostics(
	procedureCode *int64,
	triggeringMessage *aper.Enumerated,
	procedureCriticality *aper.Enumerated,
	iesCriticalityDiagnostics *ngapType.CriticalityDiagnosticsIEList) (
	criticalityDiagnostics ngapType.CriticalityDiagnostics) {
	if procedureCode != nil {
		criticalityDiagnostics.ProcedureCode = new(ngapType.ProcedureCode)
		criticalityDiagnostics.ProcedureCode.Value = *procedureCode
	}

	if triggeringMessage != nil {
		criticalityDiagnostics.TriggeringMessage = new(ngapType.TriggeringMessage)
		criticalityDiagnostics.TriggeringMessage.Value = *triggeringMessage
	}

	if procedureCriticality != nil {
		criticalityDiagnostics.ProcedureCriticality = new(ngapType.Criticality)
		criticalityDiagnostics.ProcedureCriticality.Value = *procedureCriticality
	}

	if iesCriticalityDiagnostics != nil {
		criticalityDiagnostics.IEsCriticalityDiagnostics = iesCriticalityDiagnostics
	}

	return criticalityDiagnostics
}

func buildCriticalityDiagnosticsIEItem(ieCriticality aper.Enumerated, ieID int64, typeOfErr aper.Enumerated) (
	item ngapType.CriticalityDiagnosticsIEItem) {
	item = ngapType.CriticalityDiagnosticsIEItem{
		IECriticality: ngapType.Criticality{
			Value: ieCriticality,
		},
		IEID: ngapType.ProtocolIEID{
			Value: ieID,
		},
		TypeOfError: ngapType.TypeOfError{
			Value: typeOfErr,
		},
	}

	return item
}
func (h *ngapHandler) isLatestAmfUe(amfUe *context.AmfUe) bool {
	if latestAmfUe, ok := h.backend.Context().AmfUeFindByUeContextID(amfUe.Supi); ok {
		if amfUe == latestAmfUe {
			return true
		}
	}
	return false
}
