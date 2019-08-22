package soaptemplate

import (
	"encoding/xml"
)

type CreateAccountReq struct {
	AccountType  string
	CifNo        string
	CustomerName string
	productType  string
	requestBy    string
}

type CreateAccountResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName               xml.Name
		CreateAccountResponse struct {
			XMLName xml.Name
			Return  struct {
				AccountNumber string `xml:"accountNumber"`
				ResponseCode  string `xml:"responseCode"`
				ResponseDesc  string `xml:"responseDesc"`
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
            <xsd:productType>` + req.productType + `</xsd:productType>
            <xsd:requestBy>` + req.requestBy + `</xsd:requestBy>
         </ser:request>
      </ser:createAccount>
   </soap:Body>
</soap:Envelope>`
	return stringRequest
}
