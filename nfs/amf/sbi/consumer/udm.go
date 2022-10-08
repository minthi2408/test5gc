package consumer

import (
	"etrib5gc/sbi"
	udmsdm "etrib5gc/sbi/udm/sdm"
	udmuecm "etrib5gc/sbi/udm/uecm"
	"etrib5gc/sbi/models"
)

type AmfUe interface {
	RatType() models.RatType
	NfId() string
	Guami() models.Guami
	Supi() string
	PlmnId() *models.PlmnId
	SdmSubscriptionId() string
	UpdateAccessAndMobilitySubscriptionData(*models.AccessAndMobilitySubscriptionData) error
	UpdateSmfSelectionSubscriptionData(*models.SmfSelectionSubscriptionData) error
	UpdateUeContextInAmfData(*models.UeContextInAmfData) error
	UpdateSdmSubscription(*models.SdmSubscription) error
	UpdateNssai(*models.Nssai) error
	UpdateAmf3GppAccessRegistration(*models.Amf3GppAccessRegistration) error
	UpdateAmfNon3GppAccessRegistration(*models.AmfNon3GppAccessRegistration) error
	Update3GppRegistration(*models.PatchResult) error
	UpdateNon3GppRegistration(*models.PatchResult) error
}

type UdmConsumer struct {
	ue	AmfUe
	cli sbi.ConsumerClient //connection to a remote UDM
	//other attributes to maintain remote context
}

func NewUdmConsumer(cli sbi.ConsumerClient) *UdmConsumer {
	return &UdmConsumer{
		cli: cli,
	}
}

// subcriber data management
func (c *UdmConsumer) SdmPutUpuAck(upuMacIue string) (err error) {
	ackInfo := &models.AcknowledgeInfo{
		UpuMacIue: upuMacIue,
	}
	err = udmsdm.UpuAck(c.cli, c.ue.Supi(), ackInfo)
	return		
}

func (c *UdmConsumer) SdmGetAmData() (err error) {
	var dat models.AccessAndMobilitySubscriptionData
	if dat, err = udmsdm.GetAmData(c.cli, c.ue.Supi(), "", c.ue.PlmnId(), "", ""); err == nil {
		err = c.ue.UpdateAccessAndMobilitySubscriptionData(&dat)
			//ue.GetUdmInfo().AccessAndMobilitySubscriptionData = &data
			//ue.Gpsi = data.Gpsis[0] // TODO: select GPSI
	}	
	return
}

func (c *UdmConsumer) SdmGetSmfSelectData() (err error) {
	var dat models.SmfSelectionSubscriptionData
	if dat, err = udmsdm.GetSmfSelData(c.cli, c.ue.Supi(), "", c.ue.PlmnId(), "", ""); err == nil {
		err = c.ue.UpdateSmfSelectionSubscriptionData(&dat)
	}
	return
}

func (c *UdmConsumer) SdmGetUeContextInSmfData() (err error) {
	var dat models.UeContextInAmfData
	if dat, err = udmsdm.GetUeCtxInAmfData(c.cli, c.ue.Supi(), ""); err == nil {
		err = c.ue.UpdateUeContextInAmfData(&dat)
		//ue.GetUdmInfo().UeContextInSmfData = &data
	}
	return
}

func (c *UdmConsumer) SdmSubscribe() (err error) {
	var dat models.SdmSubscription
	req := models.SdmSubscription{
		NfInstanceId: c.ue.NfId(),
		PlmnId:	*c.ue.PlmnId(),
	}

	if dat, err = udmsdm.Subscribe(c.cli, c.ue.Supi(), req); err == nil {
		err = c.ue.UpdateSdmSubscription(&dat)
		//ue.GetNssfInfo().SdmSubscriptionId = resSubscription.SubscriptionId
	}
	return
}

func (c *UdmConsumer) SdmGetSliceSelectionSubscriptionData() (err error) {
	var dat models.Nssai
	if dat, err = udmsdm.GetNSSAI(c.cli, c.ue.Supi(), "", c.ue.PlmnId(), "", ""); err == nil {
		err = c.ue.UpdateNssai(&dat)
		/*
				for _, defaultSnssai := range nssai.DefaultSingleNssais {
				subscribedSnssai := models.SubscribedSnssai{
					SubscribedSnssai: &models.Snssai{
						Sst: defaultSnssai.Sst,
						Sd:  defaultSnssai.Sd,
					},
					DefaultIndication: true,
				}

				udminfo.SubscribedNssai = append(udminfo.SubscribedNssai, subscribedSnssai)
			}
			for _, snssai := range nssai.SingleNssais {
				subscribedSnssai := models.SubscribedSnssai{
					SubscribedSnssai: &models.Snssai{
						Sst: snssai.Sst,
						Sd:  snssai.Sd,
					},
					DefaultIndication: false,
				}
				udminfo.SubscribedNssai = append(udminfo.SubscribedNssai, subscribedSnssai)
			}
		*/
		} 	
	return
}

func (c *UdmConsumer) SDMUnsubscribe() (err error) {
	err = udmsdm.Unsubscribe(c.cli, c.ue.Supi(), c.ue.SdmSubscriptionId())
	return
}

//uecontext management

func (c *UdmConsumer) UeCmRegistration(accessType models.AccessType, initialRegistrationInd bool) (err error) {

	switch accessType {
	case models.ACCESSTYPE__3_GPP_ACCESS:
		var dat models.Amf3GppAccessRegistration
		req := models.Amf3GppAccessRegistration{
				AmfInstanceId:          c.ue.NfId(),
				InitialRegistrationInd: initialRegistrationInd,
				Guami:                  c.ue.Guami(),
				RatType:                c.ue.RatType(),
				// TODO: not support Homogenous Support of IMS Voice over PS Sessions this stage
				ImsVoPs: models.IMSVOPS_HOMOGENEOUS_NON_SUPPORT,
			}
		if dat, err = udmuecm.Call3GppRegistration(c.cli, c.ue.Supi(), req); err == nil {
			err = c.ue.UpdateAmf3GppAccessRegistration(&dat)
			//ue.GetNssfInfo().UeCmRegistered = true
		}
	case models.ACCESSTYPE_NON_3_GPP_ACCESS:
		var dat models.AmfNon3GppAccessRegistration
		req := models.AmfNon3GppAccessRegistration{
			AmfInstanceId: c.ue.NfId(),
			Guami:         c.ue.Guami(),
			RatType:       c.ue.RatType(),
		}
		if 	dat, err = udmuecm.Non3GppRegistration(c.cli, c.ue.Supi(), req); err == nil {
			err = c.ue.UpdateAmfNon3GppAccessRegistration(&dat)
				//ue.GetNssfInfo().UeCmRegistered = true
		}
	}
	return
}

func (c *UdmConsumer) UeCmDeregistration(accessType models.AccessType) (err error) {
	switch accessType {
	case models.ACCESSTYPE__3_GPP_ACCESS:
		var dat models.PatchResult
		req :=  models.Amf3GppAccessRegistrationModification{
			Guami:     c.ue.Guami(),
			PurgeFlag: true,
		}
		if dat, err = udmuecm.Update3GppRegistration(c.cli, "", c.ue.Supi(), req); err == nil {
			err = c.ue.Update3GppRegistration(&dat)
		}

	case models.ACCESSTYPE_NON_3_GPP_ACCESS:
		var dat models.PatchResult
		req :=  models.AmfNon3GppAccessRegistrationModification{
			Guami:     c.ue.Guami(),
			PurgeFlag: true,
		}
		if dat, err = udmuecm.UpdateNon3GppRegistration(c.cli, c.ue.Supi(), "", req); err == nil {
			err = c.ue.UpdateNon3GppRegistration(&dat)
		}
	}
	return 
}
