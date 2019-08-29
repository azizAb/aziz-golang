package config

var Url = struct {
	Soap            string
	ApprovalAddress string
	ReportAddress   string
}{}

var Urn = struct {
	InquiryID                string
	CreateCustDataAndAccount string
	CreateAccount            string
	InquiryCIFNew            string
	InquirySACADetail        string
	GetCIFParam              string
	KodePos                  string
}{}

var Http = struct {
	Method      string
	ServingPort string
}{}

var KeyCifParam = [5]string{
	"jenispekerjaan",
	"listpenghasilan",
	"listtransaksiharian",
	"listsumberutama",
	"listtujuanpembukaanrekening",
}

func init() {
	Url.Soap = "http://10.35.65.105:8080/WsHybrid/services/SSBServiceImpl.SSBServiceImplHttpSoap12Endpoint/"
	// Url.Soap = "http://DESKTOP-IU44SKJ:8080/"
	Url.ApprovalAddress = "10.111.0.114:7001"
	Url.ReportAddress = "10.111.0.138:7777"

	Urn.InquiryID = "inquiryID"
	Urn.CreateCustDataAndAccount = "createCustDataAndAccount"
	Urn.CreateAccount = "createAccount"
	Urn.InquiryCIFNew = "inquiryCIFNew"
	Urn.InquirySACADetail = "inquirySACADetail"
	Urn.GetCIFParam = "getCIFParameter"
	Urn.KodePos = "kodePosInquiry"

	Http.Method = "POST"
	Http.ServingPort = ":9000"

}
