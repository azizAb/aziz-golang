package soaptemplate

import (
	"encoding/xml"
)

type KodeposInquiryReq struct {
	RequestBy string
	Type      string
	Value     string
}

type KodeposInquiryResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName                xml.Name
		KodePosInquiryResponse struct {
			XMLName xml.Name
			Return  struct {
				ResponseCode string `xml:"responseCode" json:"responseCode"`
				ResponseDesc string `xml:"responseDesc" json:"responseDesc"`
				ResponseItem string `xml:"responseItem" json:"responseItem"`
			} `xml:"return"`
		} `xml:"kodePosInquiryResponse"`
	}
}

func KodeposInquiryString(req *KodeposInquiryReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com" xmlns:xsd="http://service.model.bri.com/xsd">
   <soap:Header/>
   <soap:Body>
      <ser:kodePosInquiry>
         <ser:request>
            <xsd:requestBy>` + req.RequestBy + `</xsd:requestBy>
            <xsd:type>` + req.Type + `</xsd:type>
            <xsd:value>` + req.Value + `</xsd:value>
         </ser:request>
      </ser:kodePosInquiry>
   </soap:Body>
</soap:Envelope>`
	return stringRequest
}
