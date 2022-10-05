package service

import (
	amfcomm "etri5gc/sbi/amf/comm"
	amfee "etri5gc/sbi/amf/ee"
	amfloc "etri5gc/sbi/amf/loc"
	amfmt "etri5gc/sbi/amf/mt"
	ausfsor "etri5gc/sbi/ausf/sor"
	ausfuea "etri5gc/sbi/ausf/uea"
	ausfupu "etri5gc/sbi/ausf/upu"
	"etri5gc/sbi/models"
	sbimodels "etri5gc/sbi/models"
	pcfampc "etri5gc/sbi/pcf/ampc"
	pcfbtdpc "etri5gc/sbi/pcf/btdpc"
	pcfee "etri5gc/sbi/pcf/ee"
	pcfpa "etri5gc/sbi/pcf/pa"
	pcfsmpc "etri5gc/sbi/pcf/smpc"
	pcfuepc "etri5gc/sbi/pcf/uepc"
	smfee "etri5gc/sbi/smf/ee"
	smfnidd "etri5gc/sbi/smf/nidd"
	smfpdu "etri5gc/sbi/smf/pdu"
	udmee "etri5gc/sbi/udm/ee"
	udmmt "etri5gc/sbi/udm/mt"
	udmniddau "etri5gc/sbi/udm/niddau"
	udmpp "etri5gc/sbi/udm/pp"
	udmsdm "etri5gc/sbi/udm/sdm"
	udmueau "etri5gc/sbi/udm/ueau"
	udmuecm "etri5gc/sbi/udm/uecm"
	udrdr "etri5gc/sbi/udr/dr"
	udrgroup "etri5gc/sbi/udr/group"
)

func check_sbi_compile() {
	body := sbimodels.SubscriptionData{}
	amfcomm.AMFStatusChangeSubscribeModfy(nil, "", body)
	amfcomm.AMFStatusChangeUnSubscribe(nil, "")
	body1 := sbimodels.UeN1N2InfoSubscriptionCreateData{}
	amfcomm.N1N2MessageSubscribe(nil, "", body1)
	amfee.DeleteSubscription(nil, "")
	body3 := sbimodels.CancelPosInfo{}
	amfloc.CancelLocation(nil, "", body3)
	amfmt.ProvideDomainSelectionInfo(nil, "", nil, "", nil)
	body4 := sbimodels.SorInfo{}
	ausfsor.SupiUeSorPost(nil, "", body4)
	body5 := sbimodels.UpuInfo{}
	ausfupu.SupiUeUpuPost(nil, "", body5)
	ausfuea.Delete5gAkaAuthenticationResult(nil, "")
	udrdr.QuerySubsToNotify(nil, "", "")
	udrgroup.GetNfGroupIDs(nil, nil, "")
	smfpdu.ReleasePduSession(nil, "", nil)
	smfee.DeleteIndividualSubcription(nil, "")
	body6 := models.DeliverRequest{}
	smfnidd.Deliver(nil, "", body6)

	pcfee.DeletePcEventExposureSubsc(nil, "")
	pcfampc.DeleteIndividualAMPolicyAssociation(nil, "")
	body7 := models.SmPolicyDeleteData{}
	pcfsmpc.DeleteSMPolicy(nil, "", body7)
	pcfuepc.DeleteIndividualUEPolicyAssociation(nil, "")
	body8 := models.BdtReqData{}
	pcfbtdpc.CreateBDTPolicy(nil, body8)
	body9 := models.AppSessionContext{}
	pcfpa.PostAppSessions(nil, body9)

	udmee.DeleteEeSubscription(nil, "", "")
	udmmt.QueryUeInfo(nil, "", nil, "")
	udmpp.Delete5GVNGroup(nil, "", "", "")
	udmsdm.GetAmData(nil, "", "", nil, "", "")
	body10 := models.AuthEvent{}
	udmueau.ConfirmAuth(nil, "", body10)
	udmuecm.Get3GppRegistration(nil, "", "")
	body11 := models.AuthorizationInfo{}
	udmniddau.AuthorizeNiddData(nil, "", body11)
}
