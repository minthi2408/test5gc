package service

import (
	"etri5gc/fabric"
	fabric_common "etri5gc/fabric/common"
	fabric_config "etri5gc/fabric/config"

	amf_config "etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"

	//	"etri5gc/nfs/amf/nas"
	"etri5gc/nfs/amf/ngap"

	"etri5gc/nfs/amf/sbi/producer"
	"fmt"

	//"time"

	//	"github.com/free5gc/openapi/models"
	//	"github.com/gin-gonic/gin"
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
	udrdr "etri5gc/sbi/udr/dr"
	udrgroup "etri5gc/sbi/udr/group"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "service"})
}

type AMF struct {
	producer *producer.Producer    //handling Sbi requests received at the server
	ngap     *ngap.Server          //ngap server handling Ran connections and ngap messages
	context  *context.AmfContext   // AMF context
	config   *amf_config.AmfConfig // loaded AMF config
	agent    fabric.ServiceAgent
	sender   fabric.Forwarder
}

func New(config *amf_config.AmfConfig) (nf *AMF, err error) {
	nf = &AMF{
		config: config,
	}

	//create NGAP server (including a NAS handler)
	nf.ngap = ngap.NewServer(nf)

	//create sbi producer
	nf.producer = producer.New(nf, nf.ngap.Sender(), nf.ngap.Nas())

	//createfabric agent

	nf.agent, err = fabric.NewAgent(nf)
	nf.sender = nf.agent.Forwarder()

	// initialize AMF context
	nf.context = context.NewAmfContext(config, nf.sender)

	return
}

func (nf *AMF) AgentConfig() *fabric_config.AgentConfig {
	//TODO: get it from the nf config
	return nil
}

func (nf *AMF) Services() []fabric_common.Service {
	return nf.producer.Services()
}

func (nf *AMF) Context() *context.AmfContext {
	return nf.context
}

func (nf *AMF) Config() *amf_config.AmfConfig {
	return nf.config
}
func (nf *AMF) Profile() fabric_common.NfProfile {
	return nil
}
func (nf *AMF) Start() (err error) {
	check_amf_compile()
	// start ngap server
	log.Info("Starting NGAP server")
	nf.ngap.Run(nf.config.NgapIpList, 38412)

	if err != nil {
		nf.ngap.Stop()
		return
	}

	if err = nf.agent.Start(); err != nil {
		nf.ngap.Stop()
	}
	return

}

func (nf *AMF) Terminate() {
	//TODO: notify Ran of connection termination?
	//stop ngap server
	nf.ngap.Stop()
	//stop sbi server
	nf.agent.Terminate()

	fmt.Println("Kill it")
}

func check_amf_compile() {
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
}
