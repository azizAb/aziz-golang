package soaptemplate

import (
	"encoding/xml"
)

type InquiryIDReq struct {
	IdNo      string
	RequestBy string
}

type InquiryIDResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName           xml.Name
		InquiryIDResponse struct {
			XMLName xml.Name
			Return  struct {
				Address       string `xml:"address"`
				CifNo         string `xml:"cifNo"`
				CustomerName  string `xml:"customerName"`
				DateOfBirth   string `xml:"dateOfBirth"`
				Education     string `xml:"education"`
				IdNo          string `xml:"idNo"`
				MaritalStatus string `xml:"maritalStatus"`
				MotherName    string `xml:"motherName"`
				PlaceOfBirth  string `xml:"placeOfBirth"`
				Religion      string `xml:"religion"`
				ResponseCode  string `xml:"responseCode"`
				ResponseDesc  string `xml:"responseDesc"`
				RtNo          string `xml:"rtNo"`
				RwNo          string `xml:"rwNo"`
				Sex           string `xml:"sex"`
				TypeOfWork    string `xml:"typeOfWork"`
				ZipCode       string `xml:"zipCode"`
			} `xml:"return"`
		} `xml:"inquiryIDResponse"`
	}
}

func InquiriIDString(req InquiryIDReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com" xmlns:xsd="http://service.model.bri.com/xsd">
	<soap:Header/>
	<soap:Body>
	<ser:inquiryID>
		<ser:request>
			<xsd:idNo>` + req.IdNo + `</xsd:idNo>
			<xsd:requestBy>` + req.RequestBy + `</xsd:requestBy>
		</ser:request>
	</ser:inquiryID>
	</soap:Body>
	</soap:Envelope>`
	return stringRequest
}
