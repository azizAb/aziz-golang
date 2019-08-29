package soaptemplate

import (
	"encoding/xml"
)

type CifParameterReq struct {
	Key       string
	RequestBy string
}

type CifParameterResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName                 xml.Name
		GetCIFParameterResponse struct {
			XMLName xml.Name
			Return  struct {
				ResponseCode string `xml:"responseCode" json:"responseCode"`
				ResponseDesc string `xml:"responseDesc" json:"responseDesc"`
				ResponseItem string `xml:"responseItem" json:"responseItem"`
			} `xml:"return"`
		} `xml:"getCIFParameterResponse"`
	}
}

func CifParameterString(req *CifParameterReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com">
   <soap:Header/>
   <soap:Body>
      <ser:getCIFParameter>
         <ser:key>` + req.Key + `</ser:key>
         <ser:requestBy>` + req.RequestBy + `</ser:requestBy>
      </ser:getCIFParameter>
   </soap:Body>
</soap:Envelope>`
	return stringRequest
}
