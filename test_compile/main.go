package main

import (
	amfcomm "etrib5gc/sbi/amf/comm"
	amfee "etrib5gc/sbi/amf/ee"
	amfloc "etrib5gc/sbi/amf/loc"
	amfmt "etrib5gc/sbi/amf/mt"
	ausfsor "etrib5gc/sbi/ausf/sor"
	ausfuea "etrib5gc/sbi/ausf/uea"
	ausfupu "etrib5gc/sbi/ausf/upu"
	"etrib5gc/sbi/models"
	sbimodels "etrib5gc/sbi/models"
	pcfampc "etrib5gc/sbi/pcf/ampc"
	pcfbtdpc "etrib5gc/sbi/pcf/btdpc"
	pcfee "etrib5gc/sbi/pcf/ee"
	pcfpa "etrib5gc/sbi/pcf/pa"
	pcfsmpc "etrib5gc/sbi/pcf/smpc"
	pcfuepc "etrib5gc/sbi/pcf/uepc"
	smfee "etrib5gc/sbi/smf/ee"
	smfnidd "etrib5gc/sbi/smf/nidd"
	smfpdu "etrib5gc/sbi/smf/pdu"
	udmee "etrib5gc/sbi/udm/ee"
	udmmt "etrib5gc/sbi/udm/mt"
	udmniddau "etrib5gc/sbi/udm/niddau"
	udmpp "etrib5gc/sbi/udm/pp"
	udmsdm "etrib5gc/sbi/udm/sdm"
	udmueau "etrib5gc/sbi/udm/ueau"
	udmuecm "etrib5gc/sbi/udm/uecm"
	udrdr "etrib5gc/sbi/udr/dr"
	udrgroup "etrib5gc/sbi/udr/group"
)

func main() {
	check_sbi_compile()
}

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
