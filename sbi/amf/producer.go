package amf
import (
	"etr5gc/sbi/amf/comm"
	"etr5gc/sbi/amf/ee"
	"etr5gc/sbi/amf/loc"
	"etr5gc/sbi/amf/mt"
)

type Producer interface {
	comm.IndividualSubscriptionDocumentApiHandler
	comm.IndividualUeContextDocumentApiHandler
	comm.N1N2IndividualSubscriptionDocumentApiHandler
	comm.N1N2MessageCollectionDocumentApiHandler
	comm.N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApiHandler
	comm.NonUEN2MessageNotificationIndividualSubscriptionDocumentApiHandler
	comm.NonUEN2MessagesCollectionDocumentApiHandler
	comm.NonUEN2MessagesSubscriptionsCollectionDocumentApiHandler
	comm.SubscriptionsCollectionDocumentApiHandler
	ee.IndividualSubscriptionDocumentApiHandler
	ee.SubscriptionsCollectionDocumentApiHandler
	loc.IndividualUEContextDocumentApiHandler
	mt.UeContextDocumentApiHandler
	mt.UeReachIndDocumentApiHandler
}
