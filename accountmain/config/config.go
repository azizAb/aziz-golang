package config

var Url string

var Urn = struct {
	InquiryID                string
	CreateCustDataAndAccount string
	CreateAccount            string
	InquiryCIFNew            string
	InquirySACADetail        string
}{}

var Http = struct {
	Method      string
	ServingPort string
}{}

func init() {
	// Url = "http://10.35.65.105:8080/WsHybrid/services/SSBServiceImpl.SSBServiceImplHttpSoap12Endpoint/"
	Url = "http://DESKTOP-IU44SKJ:8080/"

	Urn.InquiryID = "inquiryID"
	Urn.CreateCustDataAndAccount = "createCustDataAndAccount"
	Urn.CreateAccount = "createAccount"
	Urn.InquiryCIFNew = "inquiryCIFNew"
	Urn.InquirySACADetail = "inquirySACADetail"

	Http.Method = "POST"
	Http.ServingPort = ":9000"
}
