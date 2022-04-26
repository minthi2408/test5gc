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

	"etri5gc/nfs/amf/config"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/idgenerator"
)

type IdGenerator interface {
	Allocate() (int64, error)
	FreeID(int64)
}

type AMFContext struct {
	cfg				*config.Config

	uepool			sync.Map
	ranpool			sync.Map
	ruepool			sync.Map
	eventsubpool	sync.Map
	statussubpool		sync.Map

	ladnpool    	map[string]*LADN // dnn as key

	eventsubIdGen   IdGenerator
	statussubIdGen	IdGenerator
	tmsiIdGen			IdGenerator
	ngapIdGen		IdGenerator

	
	relcap          int64 //Relative Capacity
	id              string
	services        map[models.ServiceName]models.NfService // nfservice that amf support

//	UriScheme       				models.UriScheme
	HttpIPv6Address                 string
	TNLWeightFactor                 int64
	SecurityAlgorithm               SecurityAlgorithm
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

func (amf *AMFContext) NfId() string {
	return amf.id
}

func (amf *AMFContext) ServedGuami() *models.Guami  {
	return &amf.cfg.Configuration.ServedGuamiList[0]
}
func (amf *AMFContext) TmsiAllocate() int32 {
	if tmsi, err := amf.tmsiIdGen.Allocate(); err != nil {
		return -1
	} else {
		return int32(tmsi)
	}
}

func (amf *AMFContext) FreeTmsi(tmsi int64) {
	amf.tmsiIdGen.FreeID(tmsi)
}

func (amf *AMFContext) AllocateAmfUeNgapID() (int64, error) {
	return amf.ngapIdGen.Allocate()
}

func (amf *AMFContext) AllocateGutiToUe(ue *AmfUe) {
	servedGuami := amf.cfg.Configuration.ServedGuamiList[0]
	ue.Tmsi = amf.TmsiAllocate()

	plmnID := servedGuami.PlmnId.Mcc + servedGuami.PlmnId.Mnc
	tmsiStr := fmt.Sprintf("%08x", ue.Tmsi)
	ue.Guti = plmnID + servedGuami.AmfId + tmsiStr
}

func (amf *AMFContext) AllocateRegistrationArea(ue *AmfUe, anType models.AccessType) {
	// clear the previous registration area if need
	if len(ue.RegistrationArea[anType]) > 0 {
		ue.RegistrationArea[anType] = nil
	}

	// allocate a new tai list as a registration area to ue
	// TODO: algorithm to choose TAI list
	for _, supportTai := range amf.cfg.Configuration.SupportTAIList {
		if reflect.DeepEqual(supportTai, ue.Tai) {
			ue.RegistrationArea[anType] = append(ue.RegistrationArea[anType], supportTai)
			break
		}
	}
}

func (amf *AMFContext) Ladn(dnn string) (ladn *LADN, ok bool) {
	ladn, ok =  amf.ladnpool[dnn]
	return
}

func (amf *AMFContext) NewAMFStatusSubscription(dat models.SubscriptionData) (subId string) {
	if id, err := amf.statussubIdGen.Allocate(); err != nil {
		//logger.ContextLog.Errorf("Allocate subscriptionID error: %+v", err)
	} else {
		subId = strconv.Itoa(int(id))
		amf.statussubpool.Store(subId, dat)
	}
	return
}


func (amf *AMFContext) AddAmfUeToUePool(ue *AmfUe, supi string) {
	if len(supi) == 0 {
	//	logger.ContextLog.Errorf("Supi is nil")
	}
	ue.Supi = supi
	amf.uepool.Store(ue.Supi, ue)
}

func (amf *AMFContext) NewAmfUe(supi string) *AmfUe {
	ue := AmfUe{}
	ue.init(amf)

	if supi != "" {
		amf.AddAmfUeToUePool(&ue, supi)
	}

	amf.AllocateGutiToUe(&ue)

	return &ue
}

func (amf *AMFContext) AmfUeFindByUeContextID(ueContextID string) (*AmfUe, bool) {
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

func (amf *AMFContext) AmfUeFindBySupi(supi string) (ue *AmfUe, ok bool) {
	if value, loadOk := amf.uepool.Load(supi); loadOk {
		ue = value.(*AmfUe)
		ok = loadOk
	}
	return
}

func (amf *AMFContext) AmfUeFindByPei(pei string) (ue *AmfUe, ok bool) {
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

func (amf *AMFContext) NewAmfRan(conn net.Conn) *AmfRan {
	ran := AmfRan{}
	ran.tailist = make([]SupportedTAI, 0, MaxNumOfTAI*MaxNumOfBroadcastPLMNs)
	ran.conn = conn
	amf.ranpool.Store(conn, &ran)
	return &ran
}

// use net.Conn to find RAN context, return *AmfRan and ok bit
func (amf *AMFContext) AmfRanFindByConn(conn net.Conn) (*AmfRan, bool) {
	if value, ok := amf.ranpool.Load(conn); ok {
		return value.(*AmfRan), ok
	}
	return nil, false
}

// use ranNodeID to find RAN context, return *AmfRan and ok bit
func (amf *AMFContext) AmfRanFindByRanID(ranNodeID models.GlobalRanNodeId) (*AmfRan, bool) {
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

func (amf *AMFContext) DeleteAmfRan(conn net.Conn) {
	amf.ranpool.Delete(conn)
}

func (amf *AMFContext) InSupportDnnList(targetDnn string) bool {
	for _, dnn := range amf.cfg.Configuration.SupportDnnList {
		if dnn == targetDnn {
			return true
		}
	}
	return false
}

func (amf *AMFContext) InPlmnSupportList(snssai models.Snssai) bool {
	for _, plmnSupportItem := range amf.cfg.Configuration.PlmnSupportList {
		for _, supportSnssai := range plmnSupportItem.SNssaiList {
			if reflect.DeepEqual(supportSnssai, snssai) {
				return true
			}
		}
	}
	return false
}

func (amf *AMFContext) AmfUeFindByGuti(guti string) (ue *AmfUe, ok bool) {
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

func (amf *AMFContext) AmfUeFindByPolicyAssociationID(polAssoId string) (ue *AmfUe, ok bool) {
	amf.uepool.Range(func(key, value interface{}) bool {
		candidate := value.(*AmfUe)
		if ok = (candidate.PolicyAssociationId == polAssoId); ok {
			ue = candidate
			return false
		}
		return true
	})
	return
}

func (amf *AMFContext) RanUeFindByAmfUeNgapID(amfUeNgapID int64) *RanUe {
	if value, ok := amf.ranpool.Load(amfUeNgapID); ok {
		return value.(*RanUe)
	} else {
		return nil
	}
}

func (amf *AMFContext) GetIPv4Uri() string {
	return amf.cfg.Configuration.Sbi.GetIPv4Uri()
}

func (context *AMFContext) InitNFService(serivceName []string, version string) {
	/*
	//TODO: tungtq
	tmpVersion := strings.Split(version, ".")
	versionUri := "v" + tmpVersion[0]
	for index, nameString := range serivceName {
		name := models.ServiceName(nameString)
		context.NfService[name] = models.NfService{
			ServiceInstanceId: strconv.Itoa(index),
			ServiceName:       name,
			Versions: &[]models.NfServiceVersion{
				{
					ApiFullVersion:  version,
					ApiVersionInUri: versionUri,
				},
			},
			Scheme:          context.UriScheme,
			NfServiceStatus: models.NfServiceStatus_REGISTERED,
			ApiPrefix:       context.GetIPv4Uri(),
			IpEndPoints: &[]models.IpEndPoint{
				{
					Ipv4Address: context.RegisterIPv4,
					Transport:   models.TransportProtocol_TCP,
					Port:        int32(context.SBIPort),
				},
			},
		}
	}
	*/
}

func (amf *AMFContext) init() {
	//create id generators
	amf.tmsiIdGen = idgenerator.NewGenerator(1, math.MaxInt32)
	amf.eventsubIdGen = idgenerator.NewGenerator(1, math.MaxInt32)
	amf.statussubIdGen = idgenerator.NewGenerator(1, math.MaxInt32)
	amf.ngapIdGen = idgenerator.NewGenerator(1, MaxValueOfAmfUeNgapId)
	//TODO:
}

func (amf *AMFContext) GetConfig() *config.Config {
	return amf.cfg
}

func (amf *AMFContext) BuildNfProfile() (id string, p models.NfProfile, err error) {
	return
}

func CreateAmfContext(cfg *config.Config) *AMFContext {
	//TODO
	ret := &AMFContext{}
	ret.init()
	return ret
}


