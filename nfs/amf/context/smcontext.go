package context

import (
	"sync"

	"etrib5gc/sbi/models"

	"github.com/free5gc/nas/nasMessage"
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

// create a session management context and populate its attributes with the information from an input NAS message
// the session is not attached to the AmfUe yet
func (ue *AmfUe) CreateSmContext(msg *nasMessage.ULNASTransport, pduSessionId int32, anType models.AccessType) (smc *SmContext, err error) {
	/*
		//1. identity the contexts of the target SMF
		var (
			snssai models.Snssai
			dnn    string
		)
		// A) AMF shall select an SMF

		// If the S-NSSAI IE is not included and the user's subscription context obtained from UDM. AMF shall
		// select a default snssai
		if msg.SNSSAI != nil {
			snssai = nasConvert.SnssaiToModels(msg.SNSSAI)
		} else {
			if allowedNssai, ok := ue.GetNssfInfo().AllowedNssai[anType]; ok {
				snssai = allowedNssai[0].AllowedSnssai
			} else {
				return nil, errors.New("Ue doesn't have allowedNssai")
			}
		}

		if msg.DNN != nil {
			dnn = msg.DNN.GetDNN()
		} else {
			// if user's subscription context obtained from UDM does not contain the default DNN for the,
			// S-NSSAI, the AMF shall use a locally configured DNN as the DNN
			udminfo := ue.udmcli.Info()
			dnn = ue.amf.SupportDnnList()[0]
			if udminfo.SmfSelectionData != nil {
				snssaiStr := SnssaiModelsToHex(snssai)
				if snssaiInfo, ok := udminfo.SmfSelectionData.SubscribedSnssaiInfos[snssaiStr]; ok {
					for _, dnnInfo := range snssaiInfo.DnnInfos {
						if dnnInfo.DefaultDnnIndicator {
							dnn = dnnInfo.Dnn
						}
					}
				}
			}
		}

		//2. create the session management context (and populate SMF discovery
		//query)

		//TODO tqtung - session creattion is incompleted
		//also, we need to populate an SMF discoery query
		smc = NewSmContext(pduSessionId)
		smc.SetSnssai(snssai)
		smc.SetDnn(dnn)
		smc.SetAccessType(anType)
	*/
	return
}
