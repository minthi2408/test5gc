package context

import (
	"fmt"
	"math"
	"net"
	"reflect"
	"strconv"
	"strings"
	"sync"

	//	"time"

	"etri5gc/fabric"
	"etri5gc/nfs/amf/config"
	"etri5gc/openapi/models"

	"github.com/free5gc/util/idgenerator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "context"})
}

type IdGenerator interface {
	Allocate() (int64, error)
	FreeID(int64)
}

type AmfContext struct {
	cfg           *config.AmfConfig //Amf configuration
	forwarder      fabric.Forwarder
	uepool        sync.Map //list of Amf-connected Ue
	ranpool       sync.Map //list of connected Ran
	ruepool       sync.Map //list of contacted Ue
	eventsubpool  sync.Map //list of event subscriptions
	statussubpool sync.Map //list of status subscriptions

	ladnpool map[string]*LADN // dnn as key

	eventsubIdGen  IdGenerator
	statussubIdGen IdGenerator
	tmsiIdGen      IdGenerator
	ngapIdGen      IdGenerator

	relcap   int64                                   //Relative Capacity
	id       string                                  //nf identity
	//services map[models.ServiceName]models.NfService // nfservice that amf support

	httpIPv6Address string
	tnlWeightFactor int64
	secAlgo         SecurityAlgorithm

	callback *amfCallback
}

type SecurityAlgorithm struct {
	IntegrityOrder []uint8 // slice of security.AlgIntegrityXXX
	CipheringOrder []uint8 // slice of security.AlgCipheringXXX
}

/*
func NewPlmnSupportItem() (item factory.PlmnSupportItem) {
	item.SNssaiList = make([]models.Snssai, 0, MaxNumOfSlice)
	return
}
*/

func NewAmfContext(cfg *config.AmfConfig, fw fabric.Forwarder) *AmfContext {
	ret := &AmfContext{
		cfg:            cfg,
		tmsiIdGen:      idgenerator.NewGenerator(1, math.MaxInt32),
		eventsubIdGen:  idgenerator.NewGenerator(1, math.MaxInt32),
		statussubIdGen: idgenerator.NewGenerator(1, math.MaxInt32),
		ngapIdGen:      idgenerator.NewGenerator(1, MaxValueOfAmfUeNgapId),
		ladnpool:       make(map[string]*LADN),
		//services:       make(map[models.ServiceName]models.NfService),
	}
	ret.init()
	return ret
}

func (amf *AmfContext) init() {
	amf.id = uuid.New().String()
	sec := amf.cfg.Security
	if sec != nil {
		amf.secAlgo.IntegrityOrder = getIntAlgOrder(sec.IntegrityOrder)
		amf.secAlgo.CipheringOrder = getEncAlgOrder(sec.CipheringOrder)
	}
}
func (amf *AmfContext) GetConfig() *config.AmfConfig {
	return amf.cfg
}

func (amf *AmfContext) CallbackCli() *amfCallback {
	return amf.callback
}

func (amf *AmfContext) NfId() string {
	return amf.id
}

func (amf *AmfContext) Locality() string {
	return amf.cfg.Locality
}

func (amf *AmfContext) SetNfId(id string) {
	amf.id = id
}

func (amf *AmfContext) Name() string {
	return amf.cfg.AmfName
}

func (amf *AmfContext) HttpIPv6Address() string {
	return amf.httpIPv6Address
}

func (amf *AmfContext) TNLWeightFactor() int64 {
	return amf.tnlWeightFactor
}

func (amf *AmfContext) SecurityAlgorithm() *SecurityAlgorithm {
	return &amf.secAlgo
}

func (amf *AmfContext) NetworkName() *config.NetworkName {
	return &amf.cfg.NetworkName
}

func (amf *AmfContext) RelativeCapacity() int64 {
	return amf.relcap
}

func (amf *AmfContext) SupportTaiList() []models.Tai {
	return amf.cfg.TaiList
}

func (amf *AmfContext) ServedGuamiList() []models.Guami {
	return amf.cfg.GuamiList
}

func (amf *AmfContext) PlmnSupportList() []config.PlmnItem {
	return amf.cfg.PlmnSupportList
}

func (amf *AmfContext) ServedGuami() *models.Guami {
	return &amf.cfg.GuamiList[0]
}

func (amf *AmfContext) LadnPool() map[string]*LADN {
	return amf.ladnpool
}
func (amf *AmfContext) Ladn(dnn string) (ladn *LADN, ok bool) {
	ladn, ok = amf.ladnpool[dnn]
	return
}

func (amf *AmfContext) UePool() sync.Map {
	return amf.uepool
}

func (amf *AmfContext) T3513Cfg() *config.TimerValue {
	return &amf.cfg.T3513
}

func (amf *AmfContext) AmfRanPool() sync.Map {
	return amf.ranpool
}

func (amf *AmfContext) TmsiAllocate() int32 {
	if tmsi, err := amf.tmsiIdGen.Allocate(); err != nil {
		return -1
	} else {
		return int32(tmsi)
	}
}

func (amf *AmfContext) RegisterIPv4() string {
//	return amf.cfg.Configuration.Sbi.RegisterIPv4
    return ""
}

func (amf *AmfContext) Port() int {
	//return amf.cfg.Configuration.Sbi.Port
    return 0
}
func (amf *AmfContext) FreeTmsi(tmsi int64) {
	amf.tmsiIdGen.FreeID(tmsi)
}

func (amf *AmfContext) AllocateAmfUeNgapID() (int64, error) {
	return amf.ngapIdGen.Allocate()
}

func (amf *AmfContext) AllocateGutiToUe(ue *AmfUe) {
	servedGuami := amf.cfg.GuamiList[0]
	ue.Tmsi = amf.TmsiAllocate()

	plmnID := servedGuami.PlmnId.Mcc + servedGuami.PlmnId.Mnc
	tmsiStr := fmt.Sprintf("%08x", ue.Tmsi)
	ue.Guti = plmnID + servedGuami.AmfId + tmsiStr
}

func (amf *AmfContext) AllocateRegistrationArea(ue *AmfUe, anType models.AccessType) {
	// clear the previous registration area if need
	if len(ue.RegistrationArea[anType]) > 0 {
		ue.RegistrationArea[anType] = nil
	}

	// allocate a new tai list as a registration area to ue
	// TODO: algorithm to choose TAI list
	for _, supportTai := range amf.cfg.TaiList {
		if reflect.DeepEqual(supportTai, ue.loc.Tai) {
			ue.RegistrationArea[anType] = append(ue.RegistrationArea[anType], supportTai)
			break
		}
	}
}

func (amf *AmfContext) NewAMFStatusSubscription(dat models.SubscriptionData) (subId string) {
	if id, err := amf.statussubIdGen.Allocate(); err != nil {
		log.Errorf("Allocate subscriptionID error: %+v", err)
	} else {
		subId = strconv.Itoa(int(id))
		amf.statussubpool.Store(subId, dat)
	}
	return
}

func (amf *AmfContext) GetStatusSubList() sync.Map {
	return amf.statussubpool
}

func (amf *AmfContext) AddAmfUeToUePool(ue *AmfUe, supi string) {
	if len(supi) == 0 {
		log.Errorf("Supi is nil")
	}
	ue.Supi = supi
	amf.uepool.Store(ue.Supi, ue)
}

func (amf *AmfContext) NewAmfUe(supi string) *AmfUe {
	ue := AmfUe{}
	ue.init(amf)

	if supi != "" {
		amf.AddAmfUeToUePool(&ue, supi)
	}

	amf.AllocateGutiToUe(&ue)

	return &ue
}

/*
	func (amf *AmfContext) RemoveAmfUe(ue *AmfUe) {
		for _, ranue := range ue.RanUe {
			amf.RemoveRanUe(ranue)
		}
	}
*/
func (amf *AmfContext) NewAmfUeByReq(supi string, dat *models.UeContextCreateData) *AmfUe {
	//tungtq: this procedure was written by free5gc, it looks experimential so we have to work on improvements later.
	ue := amf.NewAmfUe(supi)

	// amf.AmfRanSetByRanId(*dat.TargetId.RanNodeId)
	// ue.N1N2Message[ueContextId] = &context.N1N2Message{}
	// ue.N1N2Message[ueContextId].Request.JsonData = &models.N1N2MessageTransferReqData{
	// 	N2InfoContainer: &models.N2InfoContainer{
	// 		SmInfo: &models.N2SmInformation{
	// 			N2InfoContent: dat.SourceToTargetData,
	// 		},
	// 	},
	// }
	ue.HandoverNotifyUri = dat.N2NotifyUri

	amf.AmfRanFindByRanID(*dat.TargetId.RanNodeId)

	/*TungTQ comment - supportedTAI is not used, redundant?
	supportedTAI := NewSupportedTAI()
	supportedTAI.Tai.Tac = dat.TargetId.Tai.Tac
	supportedTAI.Tai.PlmnId = dat.TargetId.Tai.PlmnId
	*/

	// ue.N1N2MessageSubscribeInfo[ueContextID] = &models.UeN1N2InfoSubscriptionCreateData{
	// 	N2NotifyCallbackUri: dat.N2NotifyUri,
	// }
	ue.UnauthenticatedSupi = dat.UeContext.SupiUnauthInd
	// should be smInfo list

	//for _, smInfo := range dat.PduSessionList {
	//if smInfo.N2InfoContent.NgapIeType == "NgapIeType_HANDOVER_REQUIRED" {
	// ue.N1N2Message[amfSelf.Uri].Request.JsonData.N2InfoContainer.SmInfo = &smInfo
	//}
	//}

	ue.ausfcli.info.RoutingIndicator = dat.UeContext.RoutingIndicator

	// optional
	ue.udmcli.info.UdmGroupId = dat.UeContext.UdmGroupId
	ue.ausfcli.info.AusfGroupId = dat.UeContext.AusfGroupId
	// dat.UeContext.HpcfId
	ue.loc.RatType = dat.UeContext.RestrictedRatList[0] // minItem = -1
	// dat.UeContext.ForbiddenAreaList
	// dat.UeContext.ServiceAreaRestriction
	// dat.UeContext.RestrictedCoreNwTypeList

	// it's not in 5.2.2.1.1 step 2a, so don't support
	// ue.Gpsi = dat.UeContext.GpsiList
	// ue.Pei = dat.UeContext.Pei
	// dat.UeContext.GroupList
	// dat.UeContext.DrxParameter
	// dat.UeContext.SubRfsp
	// dat.UeContext.UsedRfsp
	// ue.UEAMBR = dat.UeContext.SubUeAmbr
	// dat.UeContext.SmsSupport
	// dat.UeContext.SmsfId
	// dat.UeContext.SeafData
	// dat.UeContext.Var5gMmCapability
	// dat.UeContext.PcfId
	// dat.UeContext.PcfAmPolicyUri
	// dat.UeContext.AmPolicyReqTriggerList
	// dat.UeContext.EventSubscriptionList
	// dat.UeContext.MmContextList
	// ue.CurPduSession.PduSessionId = dat.UeContext.SessionContextList.
	// ue.TraceData = dat.UeContext.TraceData
	return ue
}

func (amf *AmfContext) AmfUeFindByUeContextID(ueContextID string) (*AmfUe, bool) {
	if strings.HasPrefix(ueContextID, "imsi") {
		return amf.AmfUeFindBySupi(ueContextID)
	}
	if strings.HasPrefix(ueContextID, "imei") {
		return amf.AmfUeFindByPei(ueContextID)
	}
	if strings.HasPrefix(ueContextID, "5g-guti") {
		guti := ueContextID[strings.LastIndex(ueContextID, "-")+1:]
		return amf.AmfUeFindByGuti(guti)
	}
	return nil, false
}

func (amf *AmfContext) AmfUeFindBySupi(supi string) (ue *AmfUe, ok bool) {
	if value, loadOk := amf.uepool.Load(supi); loadOk {
		ue = value.(*AmfUe)
		ok = loadOk
	}
	return
}

func (amf *AmfContext) AmfUeFindByPei(pei string) (ue *AmfUe, ok bool) {
	amf.uepool.Range(func(key, value interface{}) bool {
		candidate := value.(*AmfUe)
		if ok = (candidate.Pei == pei); ok {
			ue = candidate
			return false
		}
		return true
	})
	return
}

func (amf *AmfContext) NewAmfRan(conn net.Conn) *AmfRan {
	ran := newAmfRan(conn, amf)
	amf.ranpool.Store(conn, &ran)
	return ran
}

// ran is not nil
func (amf *AmfContext) RemoveRan(ran *AmfRan) {
	ran.RemoveAllUe()
	amf.ranpool.Delete(ran.conn)
}

// use net.Conn to find RAN context, return *AmfRan and ok bit
func (amf *AmfContext) AmfRanFindByConn(conn net.Conn) (*AmfRan, bool) {
	if value, ok := amf.ranpool.Load(conn); ok {
		return value.(*AmfRan), ok
	}
	return nil, false
}

// use ranNodeID to find RAN context, return *AmfRan and ok bit
func (amf *AmfContext) AmfRanFindByRanID(ranNodeID models.GlobalRanNodeId) (*AmfRan, bool) {
	var ran *AmfRan
	var ok bool
	amf.ranpool.Range(func(key, value interface{}) bool {
		amfRan := value.(*AmfRan)
		switch amfRan.present {
		case RanPresentGNbId:
			if amfRan.id.GNbId.GNBValue == ranNodeID.GNbId.GNBValue {
				ran = amfRan
				ok = true
				return false
			}
		case RanPresentNgeNbId:
			if amfRan.id.NgeNbId == ranNodeID.NgeNbId {
				ran = amfRan
				ok = true
				return false
			}
		case RanPresentN3IwfId:
			if amfRan.id.N3IwfId == ranNodeID.N3IwfId {
				ran = amfRan
				ok = true
				return false
			}
		}
		return true
	})
	return ran, ok
}

func (amf *AmfContext) SupportDnnList() []string {
	return amf.cfg.DnnList
}

func (amf *AmfContext) InSupportDnnList(targetDnn string) bool {
	for _, dnn := range amf.cfg.DnnList {
		if dnn == targetDnn {
			return true
		}
	}
	return false
}

func (amf *AmfContext) InPlmnSupportList(snssai models.Snssai) bool {
	for _, plmnSupportItem := range amf.cfg.PlmnSupportList {
		for _, supportSnssai := range plmnSupportItem.SNssaiList {
			if reflect.DeepEqual(supportSnssai, snssai) {
				return true
			}
		}
	}
	return false
}

func (amf *AmfContext) AmfUeFindByGuti(guti string) (ue *AmfUe, ok bool) {
	amf.uepool.Range(func(key, value interface{}) bool {
		candidate := value.(*AmfUe)
		if ok = (candidate.Guti == guti); ok {
			ue = candidate
			return false
		}
		return true
	})
	return
}

func (amf *AmfContext) AmfUeFindByPolicyAssociationID(polAssoId string) (ue *AmfUe, ok bool) {
	amf.uepool.Range(func(key, value interface{}) bool {
		candidate := value.(*AmfUe)
		if ok = (candidate.pcfcli.Info().PolicyAssociationId == polAssoId); ok {
			ue = candidate
			return false
		}
		return true
	})
	return
}

func (amf *AmfContext) RanUeFindByAmfUeNgapID(amfUeNgapID int64) *RanUe {
	if value, ok := amf.ruepool.Load(amfUeNgapID); ok {
		return value.(*RanUe)
	} else {
		return nil
	}
}

func (amf *AmfContext) GetIPv4Uri() string {
	//return amf.cfg.Configuration.Sbi.GetIPv4Uri()
    return ""
}

// Build AMF profile to register to NRF
func (amf *AmfContext) BuildProfile() (*models.NfProfile, error) {
	profile := models.NfProfile{}
    /*
	var err error
	profile := models.NfProfile{
		NfInstanceId: amf.id,
		NfType:       models.NfType_AMF,
		NfStatus:     models.NfStatus_REGISTERED,
	}

	var plmns []models.PlmnId
	for _, plmnItem := range amf.cfg.PlmnSupportList {
		plmns = append(plmns, *plmnItem.PlmnId)
	}
	if len(plmns) > 0 {
		profile.PlmnList = &plmns
		// TODO: change to Per Plmn Support Snssai List
		profile.SNssais = &amf.cfg.PlmnSupportList[0].SNssaiList
	}
	amfInfo := models.AmfInfo{}
	if len(amf.cfg.GuamiList) == 0 {
		err = fmt.Errorf("Gumai List is Empty in AMF")
		return nil, err
	}
	regionId, setId, _, err1 := SeperateAmfId(amf.ServedGuami().AmfId)
	if err1 != nil {
		err = err1
		return nil, err
	}
	amfInfo.AmfRegionId = regionId
	amfInfo.AmfSetId = setId
	amfInfo.GuamiList = &amf.cfg.GuamiList
	if len(amf.cfg.TaiList) == 0 {
		err = fmt.Errorf("SupportTaiList is Empty in AMF")
		return nil, err
	}
	amfInfo.TaiList = &amf.cfg.TaiList
	profile.AmfInfo = &amfInfo
	if amf.cfg.Configuration.Sbi.RegisterIPv4 == "" {
		err = fmt.Errorf("AMF Address is empty")
		return nil, err
	}
	profile.Ipv4Addresses = append(profile.Ipv4Addresses, amf.cfg.Configuration.Sbi.RegisterIPv4)
	services := []models.NfService{}
	for _, nfService := range amf.services {
		services = append(services, nfService)
	}
	if len(services) > 0 {
		profile.NfServices = &services
	}
	defaultNotificationSubscription := models.DefaultNotificationSubscription{
		CallbackUri:      fmt.Sprintf("%s/namf-callback/v1/n1-message-notify", amf.GetIPv4Uri()),
		NotificationType: models.NotificationType_N1_MESSAGES,
		N1MessageClass:   models.N1MessageClass__5_GMM,
	}
	profile.DefaultNotificationSubscriptions =
		append(profile.DefaultNotificationSubscriptions, defaultNotificationSubscription)
        */
	return &profile, nil
}
