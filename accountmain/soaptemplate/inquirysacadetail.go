package soaptemplate

import (
	"encoding/xml"
)

type InquirySACADetailReq struct {
	AccountNumber string
	Branch        string
	RequestBy     string
	UserId        string
}

type InquirySACADetailResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName                   xml.Name
		InquirySACADetailResponse struct {
			XMLName xml.Name
			Return  struct {
				AvaliableBalance string `xml:"avaliableBalance"`
				CifNo            string `xml:"cifNo"`
				IdNo             string `xml:"idNo"`
				IdType           string `xml:"idType"`
				LedgerBalance    string `xml:"ledgerBalance"`
				MinimumBalance   string `xml:"minimumBalance"`
				Name             string `xml:"name"`
				ResponseCode     string `xml:"responseCode"`
				ResponseDesc     string `xml:"responseDesc"`
				Sccode           string `xml:"sccode"`
				StatusRekening   string `xml:"statusRekening"`
			} `xml:"return"`
		} `xml:"inquirySACADetailResponse"`
	}
}

func InquirySACADetailString(req InquirySACADetailReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com" xmlns:xsd="http://service.model.bri.com/xsd">
   <soap:Header/>
   <soap:Body>
      <ser:inquirySACADetail>
         <ser:request>
            <xsd:accountNumber>` + req.AccountNumber + `</xsd:accountNumber>
            <xsd:branch>` + req.Branch + `</xsd:branch>
            <xsd:requestBy>` + req.RequestBy + `</xsd:requestBy>
            <xsd:userId>` + req.UserId + `</xsd:userId>
         </ser:request>
      </ser:inquirySACADetail>
   </soap:Body>
</soap:Envelope>`
	return stringRequest
}
