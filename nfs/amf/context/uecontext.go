package context

import (

	//	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"sync"

	"etri5gc/openapi/models"

	//	"github.com/mohae/deepcopy"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/util/fsm"
	"github.com/free5gc/util/idgenerator"
)

type OnGoingProcedure string

const (
	OnGoingProcedureNothing      OnGoingProcedure = "Nothing"
	OnGoingProcedurePaging       OnGoingProcedure = "Paging"
	OnGoingProcedureN2Handover   OnGoingProcedure = "N2Handover"
	OnGoingProcedureRegistration OnGoingProcedure = "Registration"
)

const (
	NgRanCgiPresentNRCGI    int32 = 0
	NgRanCgiPresentEUTRACGI int32 = 1
)

const (
	RecommendRanNodePresentRanNode int32 = 0
	RecommendRanNodePresentTAI     int32 = 1
)

// GMM state for UE
const (
	Deregistered            fsm.StateType = "Deregistered"
	DeregistrationInitiated fsm.StateType = "DeregistrationInitiated"
	Authentication          fsm.StateType = "Authentication"
	SecurityMode            fsm.StateType = "SecurityMode"
	ContextSetup            fsm.StateType = "ContextSetup"
	Registered              fsm.StateType = "Registered"
)

type NssfInfo struct {
	NssfId                            string
	NssfUri                           string
	NetworkSliceInfo                  *models.AuthorizedNetworkSliceInfo
	AllowedNssai                      map[models.AccessType][]models.AllowedSnssai
	ConfiguredNssai                   []models.ConfiguredSnssai
	NetworkSlicingSubscriptionChanged bool
	SdmSubscriptionId                 string
	UeCmRegistered                    bool
}

//func (info *NssfInfo) copy(ueContext *models.UeContext) {
//}

type LocInfo struct {
	RatType                  models.RatType
	Location                 models.UserLocation
	Tai                      models.Tai
	LocationChanged          bool
	LastVisitedRegisteredTai models.Tai
	TimeZone                 string
}

/*
func (ue *AmfUe) updateLoc(ranUe *RanUe) {
	if ue.loc.Tai != ranUe.Tai {
		ue.loc.LocationChanged = true
	}
	ue.loc.Location = deepcopy.Copy(ranUe.Location).(models.UserLocation)
	ue.loc.Tai = deepcopy.Copy(*ue.loc.Location.NrLocation.Tai).(models.Tai)
}
*/
type AmfUe struct {
	/* the AMF which serving this AmfUe now */
	amf *AMFContext // never nil

	/* Network Slicing related context and Nssf */
	nssf NssfInfo

	/* udm client */
	udmcli udmClient

	/* pcf client */
	pcfcli pcfClient

	/* ausf client */
	ausfcli ausfClient

	/* amf client */
	amfcli amfClient

	/* callback client */
	callback ueCallback

	/* Gmm State */
	State map[models.AccessType]*fsm.State
	/* Registration procedure related context */
	RegistrationType5GS                uint8
	IdentityTypeUsedForRegistration    uint8
	RegistrationRequest                *nasMessage.RegistrationRequest
	ServingAmfChanged                  bool
	DeregistrationTargetAccessType     uint8 // only used when deregistration procedure is initialized by the network
	RegistrationAcceptForNon3GPPAccess []byte
	RetransmissionOfInitialNASMsg      bool
	RequestIdentityType                uint8
	/* Used for AMF relocation */
	TargetAmfProfile *models.NfProfile
	TargetAmfUri     string
	/* Ue Identity*/
	PlmnId              models.PlmnId
	Suci                string
	Supi                string
	UnauthenticatedSupi bool
	Gpsi                string
	Pei                 string
	Tmsi                int32 // 5G-Tmsi
	Guti                string
	GroupID             string
	EBI                 int32
	/* Ue Identity*/
	EventSubscriptionsInfo map[string]*AmfUeEventSubscription
	/* User Location*/
	loc LocInfo

	/* UeContextForHandover*/
	HandoverNotifyUri string
	/* N1N2Message */
	N1N2MessageIDGenerator          IdGenerator
	N1N2Message                     *N1N2Message
	N1N2MessageSubscribeIDGenerator IdGenerator
	// map[int64]models.UeN1N2InfoSubscriptionCreateData; use n1n2MessageSubscriptionID as key
	N1N2MessageSubscription sync.Map
	/* Pdu Sesseion context */
	SmContextList sync.Map // map[int32]*SmContext, pdu session id as key
	/* Related Context*/
	RanUe map[models.AccessType]*RanUe
	/* other */
	onGoing                       map[models.AccessType]*OnGoing
	UeRadioCapability             string // OCTET string
	Capability5GMM                nasType.Capability5GMM
	ConfigurationUpdateIndication nasType.ConfigurationUpdateIndication
	/* context related to Paging */
	UeRadioCapabilityForPaging                 *UERadioCapabilityForPaging
	InfoOnRecommendedCellsAndRanNodesForPaging *InfoOnRecommendedCellsAndRanNodesForPaging
	UESpecificDRX                              uint8

	/* Registration Area */
	RegistrationArea map[models.AccessType][]models.Tai
	LadnInfo         []LADN

	/* T3513(Paging) */
	T3513 *Timer // for paging
	/* T3565(Notification) */
	T3565 *Timer // for NAS Notification
	/* T3560 (for authentication request/security mode command retransmission) */
	T3560 *Timer
	/* T3550 (for registration accept retransmission) */
	T3550 *Timer
	/* T3522 (for deregistration request) */
	T3522 *Timer
	/* T3570 (for identity request) */
	T3570 *Timer
	/* Ue Context Release Cause */
	ReleaseCause map[models.AccessType]*CauseAll
	/* T3502 (Assigned by AMF, and used by UE to initialize registration procedure) */
	T3502Value                      int // Second
	T3512Value                      int // default 54 min
	Non3gppDeregistrationTimerValue int // default 54 min
	Lock                            sync.Mutex
}

type N1N2Message struct {
	Request     models.N1N2MessageTransferRequest
	Status      models.N1N2MessageTransferCause
	ResourceUri string
}

type OnGoing struct {
	Procedure OnGoingProcedure
	Ppi       int32 // Paging priority
}

type UERadioCapabilityForPaging struct {
	NR    string // OCTET string
	EUTRA string // OCTET string
}

// TS 38.413 9.3.1.100
type InfoOnRecommendedCellsAndRanNodesForPaging struct {
	RecommendedCells    []RecommendedCell  // RecommendedCellsForPaging
	RecommendedRanNodes []RecommendRanNode // RecommendedRanNodesForPaging
}

// TS 38.413 9.3.1.71
type RecommendedCell struct {
	NgRanCGI         NGRANCGI
	TimeStayedInCell *int64
}

// TS 38.413 9.3.1.101
type RecommendRanNode struct {
	Present         int32
	GlobalRanNodeId *models.GlobalRanNodeId
	Tai             *models.Tai
}

type NGRANCGI struct {
	Present  int32
	NRCGI    *models.Ncgi
	EUTRACGI *models.Ecgi
}

func (ue *AmfUe) init(amf *AMFContext) {
	ue.amf = amf
	ue.State = make(map[models.AccessType]*fsm.State)
	ue.State[models.AccessType__3_GPP_ACCESS] = fsm.NewState(Deregistered)
	ue.State[models.AccessType_NON_3_GPP_ACCESS] = fsm.NewState(Deregistered)
	ue.UnauthenticatedSupi = true
	ue.EventSubscriptionsInfo = make(map[string]*AmfUeEventSubscription)
	ue.RanUe = make(map[models.AccessType]*RanUe)
	ue.RegistrationArea = make(map[models.AccessType][]models.Tai)
	ue.nssf.AllowedNssai = make(map[models.AccessType][]models.AllowedSnssai)
	ue.N1N2MessageIDGenerator = idgenerator.NewGenerator(1, 2147483647)
	ue.N1N2MessageSubscribeIDGenerator = idgenerator.NewGenerator(1, 2147483647)
	ue.onGoing = make(map[models.AccessType]*OnGoing)
	ue.onGoing[models.AccessType_NON_3_GPP_ACCESS] = new(OnGoing)
	ue.onGoing[models.AccessType_NON_3_GPP_ACCESS].Procedure = OnGoingProcedureNothing
	ue.onGoing[models.AccessType__3_GPP_ACCESS] = new(OnGoing)
	ue.onGoing[models.AccessType__3_GPP_ACCESS].Procedure = OnGoingProcedureNothing
	ue.ReleaseCause = make(map[models.AccessType]*CauseAll)
}

func (ue *AmfUe) ServingAMF() *AMFContext {
	return ue.amf
}

func (ue *AmfUe) GetNssfInfo() *NssfInfo {
	return &ue.nssf
}

func (ue *AmfUe) GetLocInfo() *LocInfo {
	return &ue.loc
}

func (ue *AmfUe) UdmClient() *udmClient {
	return &ue.udmcli
}

func (ue *AmfUe) PcfClient() *pcfClient {
	return &ue.pcfcli
}

func (ue *AmfUe) AusfClient() *ausfClient {
	return &ue.ausfcli
}

func (ue *AmfUe) AmfClient() *amfClient {
	return &ue.amfcli
}

func (ue *AmfUe) CallbackClient() *ueCallback {
	return &ue.callback
}

func (ue *AmfUe) CmConnect(anType models.AccessType) bool {
	if _, ok := ue.RanUe[anType]; !ok {
		return false
	}
	return true
}

func (ue *AmfUe) CmIdle(anType models.AccessType) bool {
	return !ue.CmConnect(anType)
}

func (ue *AmfUe) Remove() {
	for _, ranUe := range ue.RanUe {
		if err := ranUe.Remove(); err != nil {
			//		logger.ContextLog.Errorf("Remove RanUe error: %v", err)
		}
	}
	ue.amf.tmsiIdGen.FreeID(int64(ue.Tmsi))
	if len(ue.Supi) > 0 {
		//AMF_Self().UePool.Delete(ue.Supi)
	}
}

func (ue *AmfUe) DetachRanUe(anType models.AccessType) {
	delete(ue.RanUe, anType)
}

func (ue *AmfUe) AttachRanUe(ranUe *RanUe) {
	if oldRanUe := ue.RanUe[ranUe.ran.atype]; oldRanUe != nil {
		//oldRanUe.Log.Infof("Implicit Deregistration - RanUeNgapID[%d]", oldRanUe.RanUeNgapId)
		oldRanUe.DetachAmfUe()
		// TODO: stop Implicit Deregistration timer
	}
	ue.RanUe[ranUe.ran.atype] = ranUe
	ranUe.amfUe = ue
}

func (ue *AmfUe) GetAnType() models.AccessType {
	if ue.CmConnect(models.AccessType__3_GPP_ACCESS) {
		return models.AccessType__3_GPP_ACCESS
	} else if ue.CmConnect(models.AccessType_NON_3_GPP_ACCESS) {
		return models.AccessType_NON_3_GPP_ACCESS
	}
	return ""
}

func (ue *AmfUe) GetCmInfo() (cmInfos []models.CmInfo) {
	var cmInfo models.CmInfo
	cmInfo.AccessType = models.AccessType__3_GPP_ACCESS
	if ue.CmConnect(cmInfo.AccessType) {
		cmInfo.CmState = models.CmState_CONNECTED
	} else {
		cmInfo.CmState = models.CmState_IDLE
	}
	cmInfos = append(cmInfos, cmInfo)
	cmInfo.AccessType = models.AccessType_NON_3_GPP_ACCESS
	if ue.CmConnect(cmInfo.AccessType) {
		cmInfo.CmState = models.CmState_CONNECTED
	} else {
		cmInfo.CmState = models.CmState_IDLE
	}
	cmInfos = append(cmInfos, cmInfo)
	return
}

func (ue *AmfUe) InAllowedNssai(targetSNssai models.Snssai, anType models.AccessType) bool {
	for _, allowedSnssai := range ue.nssf.AllowedNssai[anType] {
		if reflect.DeepEqual(*allowedSnssai.AllowedSnssai, targetSNssai) {
			return true
		}
	}
	return false
}

//TODO: tqtung - need cleaning the long chained object reference
func (ue *AmfUe) InSubscribedNssai(targetSNssai models.Snssai) bool {
	for _, sNssai := range ue.udmcli.info.SubscribedNssai {
		if reflect.DeepEqual(*sNssai.SubscribedSnssai, targetSNssai) {
			return true
		}
	}
	return false
}

func (ue *AmfUe) GetNsiInformationFromSnssai(anType models.AccessType, snssai models.Snssai) *models.NsiInformation {
	for _, allowedSnssai := range ue.nssf.AllowedNssai[anType] {
		if reflect.DeepEqual(*allowedSnssai.AllowedSnssai, snssai) {
			// TODO: select NsiInformation based on operator policy
			if len(allowedSnssai.NsiInformationList) != 0 {
				return &allowedSnssai.NsiInformationList[0]
			}
		}
	}
	return nil
}

func (ue *AmfUe) TaiListInRegistrationArea(taiList []models.Tai, accessType models.AccessType) bool {
	for _, tai := range taiList {
		if !InTaiList(tai, ue.RegistrationArea[accessType]) {
			return false
		}
	}
	return true
}

//TODO: tqtung need cleaning the long chained object reference
func (ue *AmfUe) HasWildCardSubscribedDNN() bool {
	for _, snssaiInfo := range ue.udmcli.info.SmfSelectionData.SubscribedSnssaiInfos {
		for _, dnnInfo := range snssaiInfo.DnnInfos {
			if dnnInfo.Dnn == "*" {
				return true
			}
		}
	}
	return false
}

func (ue *AmfUe) ClearRegistrationRequestData(accessType models.AccessType) {
	ue.RegistrationRequest = nil
	ue.RegistrationType5GS = 0
	ue.IdentityTypeUsedForRegistration = 0
	ue.ausfcli.clear()
	ue.ServingAmfChanged = false
	ue.RegistrationAcceptForNon3GPPAccess = nil
	if ranUe := ue.RanUe[accessType]; ranUe != nil {
		ranUe.ueContextRequest = false
	}
	ue.RetransmissionOfInitialNASMsg = false
	if onGoing := ue.onGoing[accessType]; onGoing != nil {
		onGoing.Procedure = OnGoingProcedureNothing
	}
}

func (ue *AmfUe) SetOnGoing(anType models.AccessType, onGoing *OnGoing) {
	//prevOnGoing := ue.onGoing[anType]
	ue.onGoing[anType] = onGoing
	//ue.GmmLog.Debugf("OnGoing[%s]->[%s] PPI[%d]->[%d]", prevOnGoing.Procedure, onGoing.Procedure,
	//	prevOnGoing.Ppi, onGoing.Ppi)
}

func (ue *AmfUe) OnGoing(anType models.AccessType) OnGoing {
	return *ue.onGoing[anType]
}

func (ue *AmfUe) RemoveAmPolicyAssociation() {
	ue.pcfcli.info.AmPolicyAssociation = nil
	ue.pcfcli.info.PolicyAssociationId = ""
}

func (ue *AmfUe) CopyDataFromUeContextModel(ueContext models.UeContext) {
	/*
		if ueContext.Supi != "" {
			ue.Supi = ueContext.Supi
			ue.UnauthenticatedSupi = ueContext.SupiUnauthInd
		}

		if ueContext.Pei != "" {
			ue.Pei = ueContext.Pei
		}

		if ueContext.AusfGroupId != "" {
			ue.ausfcli.info.AusfGroupId = ueContext.AusfGroupId
		}

		if ueContext.RoutingIndicator != "" {
			ue.ausfcli.info.RoutingIndicator = ueContext.RoutingIndicator
		}
		ue.udmcli.info.copy(&ueContext)

		if ueContext.SeafData != nil {
			seafData := ueContext.SeafData

			ue.seccon.NgKsi = *seafData.NgKsi
			if seafData.KeyAmf != nil {
				if seafData.KeyAmf.KeyType == models.KeyAmfType_KAMF {
					ue.ausfcli.info.Kamf = seafData.KeyAmf.KeyVal
				}
			}
			if nh, err := hex.DecodeString(seafData.Nh); err != nil {
				//	logger.ContextLog.Error(err)
				return
			} else {
				ue.seccon.NH = nh
			}
			ue.seccon.NCC = uint8(seafData.Ncc)
		}

		ue.pcfcli.Info().copy(&ueContext)
		if len(ueContext.SessionContextList) > 0 {
			for _, pduSessionContext := range ueContext.SessionContextList {
				smContext := SmContext{
					pduSessionID: pduSessionContext.PduSessionId,
					smContextRef: pduSessionContext.SmContextRef,
					snssai:       *pduSessionContext.SNssai,
					dnn:          pduSessionContext.Dnn,
					accessType:   pduSessionContext.AccessType,
					hSmfID:       pduSessionContext.HsmfId,
					vSmfID:       pduSessionContext.VsmfId,
					nsInstance:   pduSessionContext.NsInstance,
				}
				ue.StoreSmContext(pduSessionContext.PduSessionId, &smContext)
			}
		}

		if len(ueContext.MmContextList) > 0 {
			for _, mmContext := range ueContext.MmContextList {
				if mmContext.AccessType == models.AccessType__3_GPP_ACCESS {
					if nasSecurityMode := mmContext.NasSecurityMode; nasSecurityMode != nil {
						switch nasSecurityMode.IntegrityAlgorithm {
						case models.IntegrityAlgorithm_NIA0:
							ue.seccon.IntegrityAlg = security.AlgIntegrity128NIA0
						case models.IntegrityAlgorithm_NIA1:
							ue.seccon.IntegrityAlg = security.AlgIntegrity128NIA1
						case models.IntegrityAlgorithm_NIA2:
							ue.seccon.IntegrityAlg = security.AlgIntegrity128NIA2
						case models.IntegrityAlgorithm_NIA3:
							ue.seccon.IntegrityAlg = security.AlgIntegrity128NIA3
						}

						switch nasSecurityMode.CipheringAlgorithm {
						case models.CipheringAlgorithm_NEA0:
							ue.seccon.CipheringAlg = security.AlgCiphering128NEA0
						case models.CipheringAlgorithm_NEA1:
							ue.seccon.CipheringAlg = security.AlgCiphering128NEA1
						case models.CipheringAlgorithm_NEA2:
							ue.seccon.CipheringAlg = security.AlgCiphering128NEA2
						case models.CipheringAlgorithm_NEA3:
							ue.seccon.CipheringAlg = security.AlgCiphering128NEA3
						}

						if mmContext.NasDownlinkCount != 0 {
							overflow := uint16((uint32(mmContext.NasDownlinkCount) & 0x00ffff00) >> 8)
							sqn := uint8(uint32(mmContext.NasDownlinkCount & 0x000000ff))
							ue.seccon.DLCount.Set(overflow, sqn)
						}

						if mmContext.NasUplinkCount != 0 {
							overflow := uint16((uint32(mmContext.NasUplinkCount) & 0x00ffff00) >> 8)
							sqn := uint8(uint32(mmContext.NasUplinkCount & 0x000000ff))
							ue.seccon.ULCount.Set(overflow, sqn)
						}

						// TS 29.518 Table 6.1.6.3.2.1
						if mmContext.UeSecurityCapability != "" {
							// ue.SecurityCapabilities
							buf, err := base64.StdEncoding.DecodeString(mmContext.UeSecurityCapability)
							if err != nil {
								//logger.ContextLog.Error(err)
								return
							}
							ue.seccon.UESecurityCapability.Buffer = buf
							ue.seccon.UESecurityCapability.SetLen(uint8(len(buf)))
						}
					}
				}

				if mmContext.AllowedNssai != nil {
					for _, snssai := range mmContext.AllowedNssai {
						allowedSnssai := models.AllowedSnssai{
							AllowedSnssai: &snssai,
						}
						ue.nssf.AllowedNssai[mmContext.AccessType] = append(ue.nssf.AllowedNssai[mmContext.AccessType], allowedSnssai)
					}
				}
			}
		}
	*/
}

// SM Context related function

func (ue *AmfUe) StoreSmContext(pduId int32, smContext *SmContext) {
	ue.SmContextList.Store(pduId, smContext)
}

func (ue *AmfUe) SmContextFindByPDUSessionID(pduId int32) (*SmContext, bool) {
	if value, ok := ue.SmContextList.Load(pduId); ok {
		return value.(*SmContext), true
	} else {
		return nil, false
	}
}

func (ue *AmfUe) DeleteSmContext(pduId int32) {
	ue.SmContextList.Delete(pduId)
}

// /sbi/producer/location

func (ue *AmfUe) BuildLocInfo(req *models.RequestLocInfo) *models.ProvideLocInfo {
	anType := ue.GetAnType()

	if anType == "" {
		return nil
	}

	locinfo := new(models.ProvideLocInfo)

	ranUe := ue.RanUe[anType]
	if req.Req5gsLoc || req.ReqCurrentLoc {
		locinfo.CurrentLoc = true
		locinfo.Location = &ue.loc.Location
	}

	if req.ReqRatType {
		locinfo.RatType = ue.loc.RatType
	}

	if req.ReqTimeZone {
		locinfo.Timezone = ue.loc.TimeZone
	}

	if req.SupportedFeatures != "" {
		locinfo.SupportedFeatures = ranUe.supportedFeatures
	}
	return locinfo

}

// /sbi/producer/mt

func (ue *AmfUe) GetContextInfo(class string, feature string) *models.UeContextInfo {
	info := &models.UeContextInfo{}

	// TODO: Error Status 307, 403 in TS29.518 Table 6.3.3.3.3.1-3
	anType := ue.GetAnType()
	if anType != "" && class != "" {
		ranUe := ue.RanUe[anType]
		info.AccessType = anType
		info.RatType = ue.loc.RatType
		info.LastActTime = ranUe.lastActTime
		info.SupportedFeatures = ranUe.supportedFeatures
		info.SupportVoPS = ranUe.supportVoPS
		info.SupportVoPSn3gpp = ranUe.supportVoPSn3gpp
	}

	return info
}

// /sbi/producer/

//

func (ue *AmfUe) BuildCreateSmContextData(sm *SmContext, reqtype *models.RequestType) (dat models.SmContextCreateData) {
	dat.Supi = ue.Supi
	dat.UnauthenticatedSupi = ue.UnauthenticatedSupi
	dat.Pei = ue.Pei
	dat.Gpsi = ue.Gpsi
	dat.PduSessionId = sm.PduSessionID()
	snssai := sm.Snssai()
	dat.SNssai = &snssai
	dat.Dnn = sm.Dnn()
	dat.ServingNfId = ue.amf.id
	dat.Guami = &ue.amf.cfg.Configuration.ServedGuamiList[0]
	dat.ServingNetwork = ue.amf.cfg.Configuration.ServedGuamiList[0].PlmnId
	if reqtype != nil {
		dat.RequestType = *reqtype
	}
	dat.N1SmMsg = new(models.RefToBinaryData)
	dat.N1SmMsg.ContentId = "n1SmMsg"
	dat.AnType = sm.AccessType()
	if ue.loc.RatType != "" {
		dat.RatType = ue.loc.RatType
	}
	// TODO: location is used in roaming scenerio
	// if ue.Location != nil {
	// 	dat.UeLocation = ue.Location
	// }
	dat.UeTimeZone = ue.loc.TimeZone
	dat.SmContextStatusUri = ue.amf.GetIPv4Uri() + "/namf-callback/v1/smContextStatus/" +
		ue.Guti + "/" + strconv.Itoa(int(sm.PduSessionID()))

	return dat
}

func (ue *AmfUe) BuildReleaseSmContextData(cause *CauseAll, n2SmInfoType models.N2SmInfoType, n2Info []byte) (
	releaseData models.SmContextReleaseData) {
	if cause != nil {
		if cause.Cause != nil {
			releaseData.Cause = *cause.Cause
		}
		if cause.NgapCause != nil {
			releaseData.NgApCause = cause.NgapCause
		}
		if cause.Var5GmmCause != nil {
			releaseData.Var5gMmCauseValue = *cause.Var5GmmCause
		}
	}
	if ue.loc.TimeZone != "" {
		releaseData.UeTimeZone = ue.loc.TimeZone
	}
	if n2Info != nil {
		releaseData.N2SmInfoType = n2SmInfoType
		releaseData.N2SmInfo = &models.RefToBinaryData{
			ContentId: "n2SmInfo",
		}
	}
	// TODO: other param(ueLocation...)
	return
}
func (ue *AmfUe) BuildContextTransferResponse(reqdat *models.UeContextTransferReqData, res *models.UeContextTransferResponse) *models.ProblemDetails {
	resdat := res.JsonData
	//if ue.GetAnType() != UeContextTransferReqData.AccessType {
	//for _, tai := range ue.RegistrationArea[ue.GetAnType()] {
	//if UeContextTransferReqData.PlmnId == tai.PlmnId {
	// TODO : generate N2 signalling
	//}
	//}
	//}
	fn := func() {
		resdat.UeContext = ue.buildContextModel()
		clist := &resdat.UeContext.SessionContextList
		ue.SmContextList.Range(func(key, value interface{}) bool {
			smContext := value.(*SmContext)
			snssai := smContext.Snssai()
			pduSessionContext := models.PduSessionContext{
				PduSessionId: smContext.PduSessionID(),
				SmContextRef: smContext.SmContextRef(),
				SNssai:       &snssai,
				Dnn:          smContext.Dnn(),
				AccessType:   smContext.AccessType(),
				HsmfId:       smContext.HSmfID(),
				VsmfId:       smContext.VSmfID(),
				NsInstance:   smContext.NsInstance(),
			}
			*clist = append(*clist, pduSessionContext)
			return true
		})
		resdat.UeRadioCapability = &models.N2InfoContent{
			NgapMessageType: 0,
			NgapIeType:      models.NgapIeType_UE_RADIO_CAPABILITY,
			NgapData: &models.RefToBinaryData{
				ContentId: "n2Info",
			},
		}
		b := []byte(ue.UeRadioCapability)
		copy(res.BinaryDataN2Information, b)

	}

	switch reqdat.Reason {
	case models.TransferReason_INIT_REG:
		// TODO: check integrity of the registration request included in ueContextTransferRequest
		// TODO: handle condition of TS 29.518 5.2.2.2.1.1 step 2a case b
		resdat.UeContext = ue.buildContextModel()
	case models.TransferReason_MOBI_REG:
		// TODO: check integrity of the registration request included in ueContextTransferRequest
		fn()
	case models.TransferReason_MOBI_REG_UE_VALIDATED:
		fn()
	default:
		//	logger.ProducerLog.Warnf("Invalid Transfer Reason: %+v", UeContextTransferReqData.Reason)
		return &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "MANDATORY_IE_INCORRECT",
			InvalidParams: []models.InvalidParam{
				{
					Param: "reason",
				},
			},
		}
	}
	return nil
}

func (ue *AmfUe) buildContextModel() *models.UeContext {
	ueContext := &models.UeContext{
		Supi:          ue.Supi,
		SupiUnauthInd: ue.UnauthenticatedSupi,
	}
	/*
		if ue.Gpsi != "" {
			ueContext.GpsiList = append(ueContext.GpsiList, ue.Gpsi)
		}

		if ue.Pei != "" {
			ueContext.Pei = ue.Pei
		}

		if ue.udm.UdmGroupId != "" {
			ueContext.UdmGroupId = ue.UdmGroupId
		}

		if ue.AusfGroupId != "" {
			ueContext.AusfGroupId = ue.AusfGroupId
		}

		if ue.RoutingIndicator != "" {
			ueContext.RoutingIndicator = ue.RoutingIndicator
		}

		if ue.AccessAndMobilitySubscriptionData != nil {
			if ue.AccessAndMobilitySubscriptionData.SubscribedUeAmbr != nil {
				ueContext.SubUeAmbr = &models.Ambr{
					Uplink:   ue.AccessAndMobilitySubscriptionData.SubscribedUeAmbr.Uplink,
					Downlink: ue.AccessAndMobilitySubscriptionData.SubscribedUeAmbr.Downlink,
				}
			}
			if ue.AccessAndMobilitySubscriptionData.RfspIndex != 0 {
				ueContext.SubRfsp = ue.AccessAndMobilitySubscriptionData.RfspIndex
			}
		}

		if ue.PcfId != "" {
			ueContext.PcfId = ue.PcfId
		}

		if ue.AmPolicyUri != "" {
			ueContext.PcfAmPolicyUri = ue.AmPolicyUri
		}

		if ue.AmPolicyAssociation != nil {
			if len(ue.AmPolicyAssociation.Triggers) > 0 {
				ueContext.AmPolicyReqTriggerList = buildAmPolicyReqTriggers(ue.AmPolicyAssociation.Triggers)
			}
		}

		for _, eventSub := range ue.EventSubscriptionsInfo {
			if eventSub.EventSubscription != nil {
				ueContext.EventSubscriptionList = append(ueContext.EventSubscriptionList, *eventSub.EventSubscription)
			}
		}

		if ue.TraceData != nil {
			ueContext.TraceData = ue.TraceData
		}
	*/
	return ueContext
}

func buildAmPolicyReqTriggers(triggers []models.RequestTrigger) (amPolicyReqTriggers []models.AmPolicyReqTrigger) {
	for _, trigger := range triggers {
		switch trigger {
		case models.RequestTrigger_LOC_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_LOCATION_CHANGE)
		case models.RequestTrigger_PRA_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_PRA_CHANGE)
		case models.RequestTrigger_SERV_AREA_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_SARI_CHANGE)
		case models.RequestTrigger_RFSP_CH:
			amPolicyReqTriggers = append(amPolicyReqTriggers, models.AmPolicyReqTrigger_RFSP_INDEX_CHANGE)
		}
	}
	return
}

func (ue *AmfUe) N1N2InfoSubscribe(dat *models.UeN1N2InfoSubscriptionCreateData) (string, *models.ProblemDetails) {
	if id, err := ue.N1N2MessageSubscribeIDGenerator.Allocate(); err != nil {
		return "", &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
		}
	} else {
		ue.N1N2MessageSubscription.Store(id, dat)
		return strconv.Itoa(int(id)), nil
	}
}

func (ue *AmfUe) N1N2InfoUnsubscribe(subId string) {
	ue.N1N2MessageSubscription.Delete(subId)
}

func (ue *AmfUe) N1N2MessageUnsubscribe(subId string) {
	ue.N1N2MessageSubscription.Delete(subId)
}

func (ue *AmfUe) N1N2MessageStatus(uri string) (models.N1N2MessageTransferCause, bool) {
	if ue.N1N2Message == nil || ue.N1N2Message.ResourceUri != uri {
		return "", false
	}
	return ue.N1N2Message.Status, true
}
