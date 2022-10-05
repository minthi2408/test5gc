package context

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"etri5gc/sbi/models"

	"github.com/free5gc/nas/security"
)

func CompareUserLocation(loc1 models.UserLocation, loc2 models.UserLocation) bool {
	/*
		if loc1.EutraLocation != nil && loc2.EutraLocation != nil {
			eutraloc1 := deepcopy.Copy(*loc1.EutraLocation).(models.EutraLocation)
			eutraloc2 := deepcopy.Copy(*loc2.EutraLocation).(models.EutraLocation)
			eutraloc1.UeLocationTimestamp = nil
			eutraloc2.UeLocationTimestamp = nil
			return reflect.DeepEqual(eutraloc1, eutraloc2)
		}
		if loc1.N3gaLocation != nil && loc2.N3gaLocation != nil {
			return reflect.DeepEqual(loc1, loc2)
		}
		if loc1.NrLocation != nil && loc2.NrLocation != nil {
			nrloc1 := deepcopy.Copy(*loc1.NrLocation).(models.NrLocation)
			nrloc2 := deepcopy.Copy(*loc2.NrLocation).(models.NrLocation)
			nrloc1.UeLocationTimestamp = nil
			nrloc2.UeLocationTimestamp = nil
			return reflect.DeepEqual(nrloc1, nrloc2)
		}
	*/
	return false
}

func InTaiList(servedTai models.Tai, taiList []models.Tai) bool {
	for _, tai := range taiList {
		if reflect.DeepEqual(tai, servedTai) {
			return true
		}
	}
	return false
}

func TacInAreas(targetTac string, areas []models.Area) bool {
	for _, area := range areas {
		for _, tac := range area.Tacs {
			if targetTac == tac {
				return true
			}
		}
	}
	return false
}

func AttachSourceUeTargetUe(sourceUe, targetUe *RanUe) {
	if sourceUe == nil || targetUe == nil {
		return
	}

	amfUe := sourceUe.amfUe
	if amfUe == nil {
		return
	}
	targetUe.amfUe = amfUe
	targetUe.handover.SourceUe = sourceUe
	sourceUe.handover.TargetUe = targetUe
}

func DetachSourceUeTargetUe(ranUe *RanUe) {
	if ranUe == nil {
		return
	}
	if ranUe.handover.TargetUe != nil {
		targetUe := ranUe.handover.TargetUe

		ranUe.handover.TargetUe = nil
		targetUe.handover.SourceUe = nil
	} else if ranUe.handover.SourceUe != nil {
		source := ranUe.handover.SourceUe

		ranUe.handover.SourceUe = nil
		source.handover.TargetUe = nil
	}
}

func SnssaiHexToModels(hexString string) (*models.Snssai, error) {
	sst, err := strconv.ParseInt(hexString[:2], 16, 32)
	if err != nil {
		return nil, err
	}
	sNssai := models.Snssai{
		Sst: int32(sst),
		Sd:  hexString[2:],
	}
	return &sNssai, nil
}

func SnssaiModelsToHex(snssai models.Snssai) string {
	sst := fmt.Sprintf("%02x", snssai.Sst)
	return sst + snssai.Sd
}

func SeperateAmfId(amfid string) (regionId, setId, ptrId string, err error) {
	if len(amfid) != 6 {
		err = fmt.Errorf("len of amfId[%s] != 6", amfid)
		return
	}
	// regionId: 16bits, setId: 10bits, ptrId: 6bits
	regionId = amfid[:2]
	byteArray, err1 := hex.DecodeString(amfid[2:])
	if err1 != nil {
		err = err1
		return
	}
	byteSetId := []byte{byteArray[0] >> 6, byteArray[0]<<2 | byteArray[1]>>6}
	setId = hex.EncodeToString(byteSetId)[1:]
	bytePtrId := []byte{byteArray[1] & 0x3f}
	ptrId = hex.EncodeToString(bytePtrId)
	return
}

func PlmnIdStringToModels(plmnId string) (plmnID models.PlmnId) {
	plmnID.Mcc = plmnId[:3]
	plmnID.Mnc = plmnId[3:]
	return
}

func TACConfigToModels(intString string) (hexString string) {
	tmp, err := strconv.ParseUint(intString, 10, 32)
	if err != nil {
	}
	hexString = fmt.Sprintf("%06x", tmp)
	return
}
func getDuration(expiry *time.Time, remainDuration *int32) bool {
	if expiry != nil {
		if time.Now().After(*expiry) {
			return false
		} else {
			duration := time.Until(*expiry)
			*remainDuration = int32(duration.Seconds())
		}
	}
	return true
}

func getIntAlgOrder(integrityOrder []string) (intOrder []uint8) {
	for _, intAlg := range integrityOrder {
		switch intAlg {
		case "NIA0":
			intOrder = append(intOrder, security.AlgIntegrity128NIA0)
		case "NIA1":
			intOrder = append(intOrder, security.AlgIntegrity128NIA1)
		case "NIA2":
			intOrder = append(intOrder, security.AlgIntegrity128NIA2)
		case "NIA3":
			intOrder = append(intOrder, security.AlgIntegrity128NIA3)
		default:
			//logger.UtilLog.Errorf("Unsupported algorithm: %s", intAlg)
		}
	}
	return
}

func getEncAlgOrder(cipheringOrder []string) (encOrder []uint8) {
	for _, encAlg := range cipheringOrder {
		switch encAlg {
		case "NEA0":
			encOrder = append(encOrder, security.AlgCiphering128NEA0)
		case "NEA1":
			encOrder = append(encOrder, security.AlgCiphering128NEA1)
		case "NEA2":
			encOrder = append(encOrder, security.AlgCiphering128NEA2)
		case "NEA3":
			encOrder = append(encOrder, security.AlgCiphering128NEA3)
		default:
			//logger.UtilLog.Errorf("Unsupported algorithm: %s", encAlg)
		}
	}
	return
}
