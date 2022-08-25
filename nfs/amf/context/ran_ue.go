package context

import (
	"encoding/hex"
	"fmt"
	"time"

	"etri5gc/openapi/models"
	"etri5gc/openapi/utils/ngapConvert"

	"github.com/free5gc/ngap/ngapType"
	"github.com/mohae/deepcopy"
)

type RelAction int

const (
	RanUeNgapIdUnspecified int64 = 0xffffffff
)

const (
	UeContextN2NormalRelease RelAction = iota
	UeContextReleaseHandover
	UeContextReleaseUeContext
)

type HandoverInfo struct {
	HandOverType        ngapType.HandoverType
	SuccessPduSessionId []int32
	SourceUe            *RanUe
	TargetUe            *RanUe
}

type RanUe struct {
	/* UE identity*/
	ranUeNgapId int64
	amfUeNgapId int64

	/* HandOver Info*/
	handover HandoverInfo

	/* UserLocation*/
	tai      models.Tai
	location models.UserLocation

	/* context about udm */
	supportVoPSn3gpp  bool
	supportVoPS       bool
	supportedFeatures string
	lastActTime       *time.Time

	/* Related Context*/
	amfUe *AmfUe
	ran   *AmfRan

	/* Routing ID */
	routingId string
	/* Trace Recording Session Reference */
	trsr string
	/* Ue Context Release Action */
	relAct RelAction

	/* context used for AMF Re-allocation procedure */
	OldAmfName            string
	InitialUEMessage      []byte
	RRCEstablishmentCause string // Received from initial ue message; pattern: ^[0-9a-fA-F]+$
	ueContextRequest      bool

	/* send initial context setup request or not*/
	SentInitialContextSetupRequest bool
}

/*
func (ranUe *RanUe) SetUeNgapId(v int64) {
	ranUe.RanUeNgapId = v
}
*/
func (ranUe *RanUe) Ran() *AmfRan {
	return ranUe.ran
}

func (ranUe *RanUe) AmfUe() *AmfUe {
	return ranUe.amfUe
}
func (ranUe *RanUe) GetHandoverInfo() *HandoverInfo {
	return &ranUe.handover
}

func (ranUe *RanUe) UeContextRequest() bool {
	return ranUe.ueContextRequest
}

func (ranUe *RanUe) SetUeContextRequest(f bool) {
	ranUe.ueContextRequest = f
}

func (ranUe *RanUe) ReleaseAction() RelAction {
	return ranUe.relAct
}
func (ranUe *RanUe) SetReleaseAction(a RelAction) {
	ranUe.relAct = a
}

func (ranUe *RanUe) RanUeNgapId() int64 {
	return ranUe.ranUeNgapId
}

func (ranUe *RanUe) SetRanUeNgapId(id int64) {
	ranUe.ranUeNgapId = id
}
func (ranUe *RanUe) AmfUeNgapId() int64 {
	return ranUe.amfUeNgapId
}

func (ranUe *RanUe) AmfRanUeNgapId(id int64) {
	ranUe.amfUeNgapId = id
}

func (ranUe *RanUe) RoutingId() string {
	return ranUe.routingId
}
func (ranUe *RanUe) SetRoutingId(id string) {
	ranUe.routingId = id
}

func (ranUe *RanUe) Location() models.UserLocation {
	return ranUe.location
}

func (ranUe *RanUe) SetLocation(loc models.UserLocation) {
	ranUe.location = loc
}
func (ranUe *RanUe) Tai() models.Tai {
	return ranUe.tai
}

func (ranUe *RanUe) Trsr() string {
	return ranUe.trsr
}
func (ranUe *RanUe) SetTrsr(s string) {
	ranUe.trsr = s
}
func (ranUe *RanUe) Remove() error {
	ran := ranUe.ran
	if ranUe.amfUe != nil {
		ranUe.amfUe.DetachRanUe(ran.atype)
		ranUe.DetachAmfUe()
	}

	for index, ranUe1 := range ran.uelist {
		if ranUe1 == ranUe {
			ran.uelist = append(ran.uelist[:index], ran.uelist[index+1:]...)
			break
		}
	}
	ran.amf.ruepool.Delete(ranUe.amfUeNgapId)
	ran.amf.ngapIdGen.FreeID(ranUe.amfUeNgapId)
	return nil
}

func (ranUe *RanUe) DetachAmfUe() {
	ranUe.amfUe = nil
}

func (ranUe *RanUe) SwitchToRan(newRan *AmfRan, ranUeNgapId int64) error {

	if newRan == nil {
		return fmt.Errorf("newRan is nil")
	}

	oldRan := ranUe.ran

	// remove ranUe from oldRan
	for index, ranUe1 := range oldRan.uelist {
		if ranUe1 == ranUe {
			oldRan.uelist = append(oldRan.uelist[:index], oldRan.uelist[index+1:]...)
			break
		}
	}

	// add ranUe to newRan
	newRan.uelist = append(newRan.uelist, ranUe)

	// switch to newRan
	ranUe.ran = newRan
	ranUe.ranUeNgapId = ranUeNgapId

	//	logger.ContextLog.Infof("RanUe[RanUeNgapID: %d] Switch to new Ran[Name: %s]", ranUe.RanUeNgapId, ranUe.Ran.Name)
	return nil
}

func (ranUe *RanUe) UpdateLocation(userLocationInformation *ngapType.UserLocationInformation) {
	if userLocationInformation == nil {
		return
	}

	amf := ranUe.ran.amf
	curTime := time.Now().UTC()
	switch userLocationInformation.Present {
	case ngapType.UserLocationInformationPresentUserLocationInformationEUTRA:
		locationInfoEUTRA := userLocationInformation.UserLocationInformationEUTRA
		if ranUe.location.EutraLocation == nil {
			ranUe.location.EutraLocation = new(models.EutraLocation)
		}

		tAI := locationInfoEUTRA.TAI
		plmnID := ngapConvert.PlmnIdToModels(tAI.PLMNIdentity)
		tac := hex.EncodeToString(tAI.TAC.Value)

		if ranUe.location.EutraLocation.Tai == nil {
			ranUe.location.EutraLocation.Tai = new(models.Tai)
		}
		ranUe.location.EutraLocation.Tai.PlmnId = &plmnID
		ranUe.location.EutraLocation.Tai.Tac = tac
		ranUe.tai = *ranUe.location.EutraLocation.Tai

		eUTRACGI := locationInfoEUTRA.EUTRACGI
		ePlmnID := ngapConvert.PlmnIdToModels(eUTRACGI.PLMNIdentity)
		eutraCellID := ngapConvert.BitStringToHex(&eUTRACGI.EUTRACellIdentity.Value)

		if ranUe.location.EutraLocation.Ecgi == nil {
			ranUe.location.EutraLocation.Ecgi = new(models.Ecgi)
		}
		ranUe.location.EutraLocation.Ecgi.PlmnId = &ePlmnID
		ranUe.location.EutraLocation.Ecgi.EutraCellId = eutraCellID
		ranUe.location.EutraLocation.UeLocationTimestamp = &curTime
		if locationInfoEUTRA.TimeStamp != nil {
			ranUe.location.EutraLocation.AgeOfLocationInformation = ngapConvert.TimeStampToInt32(
				locationInfoEUTRA.TimeStamp.Value)
		}
		if ranUe.amfUe != nil {
			//TODO: tungtq - create a update location method for AmfRan.loc
			//ranUe.amfUe.updateloc(ranUe)
			if ranUe.amfUe.loc.Tai != ranUe.tai {
				ranUe.amfUe.loc.LocationChanged = true
			}
			ranUe.amfUe.loc.Location = deepcopy.Copy(ranUe.location).(models.UserLocation)
			ranUe.amfUe.loc.Tai = deepcopy.Copy(*ranUe.amfUe.loc.Location.EutraLocation.Tai).(models.Tai)
		}
	case ngapType.UserLocationInformationPresentUserLocationInformationNR:
		locationInfoNR := userLocationInformation.UserLocationInformationNR
		if ranUe.location.NrLocation == nil {
			ranUe.location.NrLocation = new(models.NrLocation)
		}

		tAI := locationInfoNR.TAI
		plmnID := ngapConvert.PlmnIdToModels(tAI.PLMNIdentity)
		tac := hex.EncodeToString(tAI.TAC.Value)

		if ranUe.location.NrLocation.Tai == nil {
			ranUe.location.NrLocation.Tai = new(models.Tai)
		}
		ranUe.location.NrLocation.Tai.PlmnId = &plmnID
		ranUe.location.NrLocation.Tai.Tac = tac
		ranUe.tai = deepcopy.Copy(*ranUe.location.NrLocation.Tai).(models.Tai)

		nRCGI := locationInfoNR.NRCGI
		nRPlmnID := ngapConvert.PlmnIdToModels(nRCGI.PLMNIdentity)
		nRCellID := ngapConvert.BitStringToHex(&nRCGI.NRCellIdentity.Value)

		if ranUe.location.NrLocation.Ncgi == nil {
			ranUe.location.NrLocation.Ncgi = new(models.Ncgi)
		}
		ranUe.location.NrLocation.Ncgi.PlmnId = &nRPlmnID
		ranUe.location.NrLocation.Ncgi.NrCellId = nRCellID
		ranUe.location.NrLocation.UeLocationTimestamp = &curTime
		if locationInfoNR.TimeStamp != nil {
			ranUe.location.NrLocation.AgeOfLocationInformation = ngapConvert.TimeStampToInt32(locationInfoNR.TimeStamp.Value)
		}
		if ranUe.amfUe != nil {
			if ranUe.amfUe.loc.Tai != ranUe.tai {
				ranUe.amfUe.loc.LocationChanged = true
			}
			ranUe.amfUe.loc.Location = deepcopy.Copy(ranUe.location).(models.UserLocation)
			ranUe.amfUe.loc.Tai = deepcopy.Copy(*ranUe.amfUe.loc.Location.NrLocation.Tai).(models.Tai)
		}
	case ngapType.UserLocationInformationPresentUserLocationInformationN3IWF:
		locationInfoN3IWF := userLocationInformation.UserLocationInformationN3IWF
		if ranUe.location.N3gaLocation == nil {
			ranUe.location.N3gaLocation = new(models.N3gaLocation)
		}

		ip := locationInfoN3IWF.IPAddress
		port := locationInfoN3IWF.PortNumber

		ipv4Addr, ipv6Addr := ngapConvert.IPAddressToString(ip)

		ranUe.location.N3gaLocation.UeIpv4Addr = ipv4Addr
		ranUe.location.N3gaLocation.UeIpv6Addr = ipv6Addr
		ranUe.location.N3gaLocation.PortNumber = ngapConvert.PortNumberToInt(port)
		// N3GPP TAI is operator-specific
		// TODO: define N3GPP TAI
		ranUe.location.N3gaLocation.N3gppTai = &models.Tai{
			PlmnId: amf.cfg.Configuration.SupportTAIList[0].PlmnId,
			Tac:    amf.cfg.Configuration.SupportTAIList[0].Tac,
		}
		ranUe.tai = deepcopy.Copy(*ranUe.location.N3gaLocation.N3gppTai).(models.Tai)

		if ranUe.amfUe != nil {
			ranUe.amfUe.loc.Location = deepcopy.Copy(ranUe.location).(models.UserLocation)
			ranUe.amfUe.loc.Tai = *ranUe.location.N3gaLocation.N3gppTai
		}
	case ngapType.UserLocationInformationPresentNothing:
	}
}
