package soaptemplate

import (
	"encoding/xml"
)

type CreateAccountReq struct {
	AccountType  string `json:"accountType"`
	CifNo        string `json:"cifNo"`
	CustomerName string `json:"customerName"`
	ProductType  string `json:"productType"`
	RequestBy    string `json:"requestBy"`
}

type CreateAccountResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName               xml.Name
		CreateAccountResponse struct {
			XMLName xml.Name
			Return  struct {
				AccountNumber string `xml:"accountNumber" json:"accountNumber"`
				ResponseCode  string `xml:"responseCode" json:"responseCode"`
				ResponseDesc  string `xml:"responseDesc" json:"responseDesc"`
			} `xml:"return"`
		} `xml:"createAccountResponse"`
	}
}

func CreateAccountString(req CreateAccountReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com" xmlns:xsd="http://service.model.bri.com/xsd">
   <soap:Header/>
   <soap:Body>
      <ser:createAccount>
         <ser:request>
            <xsd:accountType>` + req.AccountType + `</xsd:accountType>
            <xsd:cifNo>` + req.CifNo + `</xsd:cifNo>
            <xsd:customerName>` + req.CustomerName + `</xsd:customerName>
            <xsd:productType>` + req.ProductType + `</xsd:productType>
            <xsd:requestBy>` + req.RequestBy + `</xsd:requestBy>
         </ser:request>
      </ser:createAccount>
   </soap:Body>
</soap:Envelope>`
	return stringRequest
}
