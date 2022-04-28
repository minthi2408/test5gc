package context

import (
	"fmt"
	"net"

	"github.com/free5gc/ngap/ngapConvert"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

const (
	RanPresentGNbId   = 1
	RanPresentNgeNbId = 2
	RanPresentN3IwfId = 3
)

type AmfRan struct {
	amf				*AMFContext
	present			int
	id      		*models.GlobalRanNodeId
	name  		    string
	atype   		models.AccessType
	/* socket Connect*/
	conn 			net.Conn
	/* Supported TA List */
	tailist 		[]SupportedTAI

	/* RAN UE List */
	uelist 			[]*RanUe // RanUeNgapId as key

}

func (r *AmfRan) Id() *models.GlobalRanNodeId {
	return r.id
}

func (r *AmfRan) Conn() net.Conn {
	return r.conn
}

func (r *AmfRan) Name() string {
	return r.name
}

func (r *AmfRan) SetName(name string) {
	r.name = name
}

func (r *AmfRan) AnType() models.AccessType {
	return r.atype
}
func (r *AmfRan) Present() int {
	return r.present
}

func (r *AmfRan) AccessType() models.AccessType {
	return r.atype
}

func (r *AmfRan) SupportedTAList() []SupportedTAI {
	return r.tailist
}

func (r *AmfRan) SetSupportedTAList(l []SupportedTAI) {
	r.tailist = l
}

type SupportedTAI struct {
	Tai        models.Tai
	SNssaiList []models.Snssai
}

func NewSupportedTAI() (tai SupportedTAI) {
	tai.SNssaiList = make([]models.Snssai, 0, MaxNumOfSlice)
	return
}

func (ran *AmfRan) Remove() {
	//ran.Log.Infof("Remove RAN Context[ID: %+v]", ran.RanID())
	ran.RemoveAllUeInRan()
	ran.amf.DeleteAmfRan(ran.conn)
}


func (ran *AmfRan) RemoveAllUeInRan() {
	saveRanUeList := make([]*RanUe, len(ran.uelist))
	copy(saveRanUeList, ran.uelist)
	for _, ranUe := range saveRanUeList {
		if err := ranUe.Remove(); err != nil {
		//	logger.ContextLog.Errorf("Remove RanUe error: %v", err)
		}
	}
}

func (ran *AmfRan) RanUeFindByRanUeNgapID(ranUeNgapID int64) *RanUe {
	for _, ranUe := range ran.uelist {
		if ranUe.RanUeNgapId == ranUeNgapID {
			return ranUe
		}
	}
	return nil
}

func (ran *AmfRan) SetRanId(ranNodeId *ngapType.GlobalRANNodeID) {
	ranId := ngapConvert.RanIdToModels(*ranNodeId)
	ran.present = ranNodeId.Present
	ran.id = &ranId
	if ranNodeId.Present == ngapType.GlobalRANNodeIDPresentGlobalN3IWFID {
		ran.atype = models.AccessType_NON_3_GPP_ACCESS
	} else {
		ran.atype = models.AccessType__3_GPP_ACCESS
	}
}

func (ran *AmfRan) RanID() string {
	switch ran.present {
	case RanPresentGNbId:
		return fmt.Sprintf("<PlmnID: %+v, GNbID: %s>", *ran.id.PlmnId, ran.id.GNbId.GNBValue)
	case RanPresentN3IwfId:
		return fmt.Sprintf("<PlmnID: %+v, N3IwfID: %s>", *ran.id.PlmnId, ran.id.N3IwfId)
	case RanPresentNgeNbId:
		return fmt.Sprintf("<PlmnID: %+v, NgeNbID: %s>", *ran.id.PlmnId, ran.id.NgeNbId)
	default:
		return ""
	}
}

func (ran *AmfRan) NewRanUe(ranUeNgapID int64) (*RanUe, error) {
	if amfUeNgapID, err := ran.amf.AllocateAmfUeNgapID(); err != nil {
		return nil, fmt.Errorf("Allocate AMF UE NGAP ID error: %+v", err)
	} else {
		ranUe := &RanUe{
			AmfUeNgapId: amfUeNgapID,
			RanUeNgapId: ranUeNgapID,
			Ran: ran,
		}

		ran.uelist = append(ran.uelist, ranUe)
		ran.amf.ruepool.Store(ranUe.AmfUeNgapId, ranUe)
		return ranUe, nil
	}
}


