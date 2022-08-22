package context

import (
	"sync"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/openapi/models"
)

type SmContext struct {
	mu sync.RWMutex // protect the following fields

	// pdu session information
	pduSessionID int32
	smContextRef string
	snssai       models.Snssai
	dnn          string
	accessType   models.AccessType
	nsInstance   string
	userLocation models.UserLocation
	plmnID       models.PlmnId

	// SMF information
	smfID  string
	smfUri string
	hSmfID string
	vSmfID string

	/* Smf client that sends request to a SMF producer */
	smfcli smfClient

	// for duplicate pdu session id handling
	ulNASTransport *nasMessage.ULNASTransport
	duplicated     bool
}

func (c *SmContext) SmfClient() *smfClient {
	return &c.smfcli
}

func NewSmContext(pduSessionID int32) *SmContext {
	c := &SmContext{pduSessionID: pduSessionID}
	return c
}

func (c *SmContext) PduSessionID() int32 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.pduSessionID
}

func (c *SmContext) SetPduSessionID(id int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pduSessionID = id
}

func (c *SmContext) SmContextRef() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.smContextRef
}

func (c *SmContext) SetSmContextRef(ref string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.smContextRef = ref
}

func (c *SmContext) AccessType() models.AccessType {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.accessType
}

func (c *SmContext) SetAccessType(accessType models.AccessType) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.accessType = accessType
}

func (c *SmContext) Snssai() models.Snssai {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.snssai
}

func (c *SmContext) SetSnssai(snssai models.Snssai) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.snssai = snssai
}

func (c *SmContext) Dnn() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.dnn
}

func (c *SmContext) SetDnn(dnn string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dnn = dnn
}

func (c *SmContext) NsInstance() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.nsInstance
}

func (c *SmContext) SetNsInstance(nsInstanceID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nsInstance = nsInstanceID
}

func (c *SmContext) UserLocation() models.UserLocation {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.userLocation
}

func (c *SmContext) SetUserLocation(userLocation models.UserLocation) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.userLocation = userLocation
}

func (c *SmContext) PlmnID() models.PlmnId {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.plmnID
}

func (c *SmContext) SetPlmnID(plmnID models.PlmnId) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.plmnID = plmnID
}

func (c *SmContext) SmfID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.smfID
}

func (c *SmContext) SetSmfID(smfID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.smfID = smfID
}

func (c *SmContext) SmfUri() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.smfUri
}

func (c *SmContext) SetSmfUri(smfUri string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.smfUri = smfUri
}

func (c *SmContext) HSmfID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.hSmfID
}

func (c *SmContext) SetHSmfID(hsmfID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hSmfID = hsmfID
}

func (c *SmContext) VSmfID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.vSmfID
}

func (c *SmContext) SetVSmfID(vsmfID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vSmfID = vsmfID
}

func (c *SmContext) PduSessionIDDuplicated() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.duplicated
}

func (c *SmContext) SetDuplicatedPduSessionID(duplicated bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.duplicated = duplicated
}

func (c *SmContext) ULNASTransport() *nasMessage.ULNASTransport {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.ulNASTransport
}

func (c *SmContext) StoreULNASTransport(msg *nasMessage.ULNASTransport) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ulNASTransport = msg
}

func (c *SmContext) DeleteULNASTransport() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ulNASTransport = nil
}

func (ue *AmfUe) CreateSmContext(msg *nasMessage.ULNASTransport) (smc *SmContext, err error) {
	/* TODO: tqtung - comment out as a reference for network selection implementation
	var (
		snssai models.Snssai
		dnn    string
	)
	// A) AMF shall select an SMF

	// If the S-NSSAI IE is not included and the user's subscription context obtained from UDM. AMF shall
	// select a default snssai
	if ulNasTransport.SNSSAI != nil {
		snssai = nasConvert.SnssaiToModels(ulNasTransport.SNSSAI)
	} else {
		if allowedNssai, ok := ue.GetNssfInfo().AllowedNssai[anType]; ok {
			snssai = *allowedNssai[0].AllowedSnssai
		} else {
			return errors.New("Ue doesn't have allowedNssai")
		}
	}

	if ulNasTransport.DNN != nil {
		dnn = ulNasTransport.DNN.GetDNN()
	} else {
		// if user's subscription context obtained from UDM does not contain the default DNN for the,
		// S-NSSAI, the AMF shall use a locally configured DNN as the DNN
		dnn = gmm.nas.amf().SupportDnnList()[0]
		if udminfo.SmfSelectionData != nil {
			snssaiStr := context.SnssaiModelsToHex(snssai)
			if snssaiInfo, ok := udminfo.SmfSelectionData.SubscribedSnssaiInfos[snssaiStr]; ok {
				for _, dnnInfo := range snssaiInfo.DnnInfos {
					if dnnInfo.DefaultDnnIndicator {
						dnn = dnnInfo.Dnn
					}
				}
			}
		}
	}
	//TODO: tungtq - implement this in the context package
	/*
		if newSmContext, cause, err := gmm.nas.backend.NfSelector().SelectSmf(ue, anType, pduSessionID, snssai, dnn); err != nil {
			log.Errorf("Select SMF failed: %+v", err)
			gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
				smMessage, pduSessionID, cause, nil, 0)
		} else {
			ue.Lock.Lock()
			defer ue.Lock.Unlock()
			_, smContextRef, errResponse, problemDetail, err :=
				gmm.consumer.Smf().SendCreateSmContextRequest(ue, newSmContext, nil, smMessage)
			if err != nil {
				log.Errorf("CreateSmContextRequest Error: %+v", err)
				return nil
			} else if problemDetail != nil {
				// TODO: error handling
				return fmt.Errorf("Failed to Create smContext[pduSessionID: %d], Error[%v]", pduSessionID, problemDetail)
			} else if errResponse != nil {
				log.Warnf("PDU Session Establishment Request is rejected by SMF[pduSessionId:%d]",
					pduSessionID)
				gmm.nas.SendDLNASTransport(ue.RanUe[anType], nasMessage.PayloadContainerTypeN1SMInfo,
					errResponse.BinaryDataN1SmMessage, pduSessionID, 0, nil, 0)
			} else {
				newSmContext.SetSmContextRef(smContextRef)
				newSmContext.SetUserLocation(deepcopy.Copy(ue.GetLocInfo().Location).(models.UserLocation))
				ue.StoreSmContext(pduSessionID, newSmContext)
				log.Infof("create smContext[pduSessionID: %d] Success", pduSessionID)
				// TODO: handle response(response N2SmInfo to RAN if exists)
			}
		}
	*/

	return
}
