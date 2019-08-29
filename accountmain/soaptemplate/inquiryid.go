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
				Address       string `xml:"address" json:"address"`
				CifNo         string `xml:"cifNo" json:"cifNo"`
				CustomerName  string `xml:"customerName" json:"customerName"`
				DateOfBirth   string `xml:"dateOfBirth" json:"dateOfBirth"`
				Education     string `xml:"education" json:"education"`
				IdNo          string `xml:"idNo" json:"idNo"`
				MaritalStatus string `xml:"maritalStatus" json:"maritalStatus"`
				MotherName    string `xml:"motherName" json:"motherName"`
				PlaceOfBirth  string `xml:"placeOfBirth" json:"placeOfBirth"`
				Religion      string `xml:"religion" json:"religion"`
				ResponseCode  string `xml:"responseCode" json:"responseCode"`
				ResponseDesc  string `xml:"responseDesc" json:"responseDesc"`
				RtNo          string `xml:"rtNo" json:"rtNo"`
				RwNo          string `xml:"rwNo" json:"rwNo"`
				Sex           string `xml:"sex" json:"sex"`
				TypeOfWork    string `xml:"typeOfWork" json:"typeOfWork"`
				ZipCode       string `xml:"zipCode" json:"zipCode"`
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
