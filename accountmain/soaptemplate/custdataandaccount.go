package soaptemplate

import (
	"encoding/xml"
)

type CustDataAndAccountReq struct {
	AccountType            string `json:"accountType"`
	Address                string `json:"address"`
	CellPhoneNumber        string `json:"cellPhoneNumber"`
	CustomerName           string `json:"customerName"`
	DailyTransactionAmount string `json:"dailyTransactionAmount"`
	DateOfBirth            string `json:"dateOfBirth"`
	Education              string `json:"education"`
	EmailAddress           string `json:"emailAddress"`
	FieldWork              string `json:"fieldWork"`
	IdNo                   string `json:"idNo"`
	IncomeSource           string `json:"incomeSource"`
	MaritalStatus          string `json:"maritalStatus"`
	MonthlyIncome          string `json:"monthlyIncome"`
	MotherName             string `json:"motherName"`
	NpwpNo                 string `json:"npwpNo"`
	OfficeAddress          string `json:"officeAddress"`
	OfficeName             string `json:"officeName"`
	OpeningAccountPurpose  string `json:"openingAccountPurpose"`
	PeriodOfWork           string `json:"periodOfWork"`
	PhoneNumber            string `json:"phoneNumber"`
	PlaceOfBirth           string `json:"placeOfBirth"`
	ProductType            string `json:"productType"`
	Religion               string `json:"religion"`
	RequestBy              string `json:"requestBy"`
	RtNo                   string `json:"rtNo"`
	RwNo                   string `json:"rwNo"`
	Sex                    string `json:"sex"`
	TypeOfWork             string `json:"typeOfWork"`
	WorkingPosition        string `json:"workingPosition"`
	ZipCode                string `json:"zipCode"`
}

type CustDataAndAccountResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName                          xml.Name
		CreateCustDataAndAccountResponse struct {
			XMLName xml.Name
			Return  struct {
				AccountNumber string `xml:"accountNumber" json:"accountNumber"`
				CifNo         string `xml:"cifNo" json:"cifNo"`
				ResponseCode  string `xml:"responseCode" json:"responseCode"`
				ResponseDesc  string `xml:"responseDesc" json:"responseDesc"`
			} `xml:"return"`
		} `xml:"createCustDataAndAccountResponse"`
	}
}

func CustDataAndAccountString(req CustDataAndAccountReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com" xmlns:xsd="http://service.model.bri.com/xsd">
   <soap:Header/>
   <soap:Body>
      <ser:createCustDataAndAccount>
         <ser:request>
            <xsd:accountType>` + req.AccountType + `</xsd:accountType>
            <xsd:address>` + req.Address + `</xsd:address>
            <xsd:cellPhoneNumber>` + req.CellPhoneNumber + `</xsd:cellPhoneNumber>
            <xsd:customerName>` + req.CustomerName + `</xsd:customerName>
            <xsd:dailyTransactionAmount>` + req.DailyTransactionAmount + `</xsd:dailyTransactionAmount>
            <xsd:dateOfBirth>` + req.DateOfBirth + `</xsd:dateOfBirth>
            <xsd:education>` + req.Education + `</xsd:education>
            <xsd:emailAddress>` + req.EmailAddress + `</xsd:emailAddress>
            <xsd:fieldWork>` + req.FieldWork + `</xsd:fieldWork>
            <xsd:idNo>` + req.IdNo + `</xsd:idNo>
            <xsd:incomeSource>` + req.IncomeSource + `</xsd:incomeSource>
            <xsd:maritalStatus>` + req.MaritalStatus + `</xsd:maritalStatus>
            <xsd:monthlyIncome>` + req.MonthlyIncome + `</xsd:monthlyIncome>
            <xsd:motherName>` + req.MotherName + `</xsd:motherName>
            <xsd:npwpNo>` + req.NpwpNo + `</xsd:npwpNo>
            <xsd:officeAddress>` + req.OfficeAddress + `</xsd:officeAddress>
            <xsd:officeName>` + req.OfficeName + `</xsd:officeName>
            <xsd:openingAccountPurpose>` + req.OpeningAccountPurpose + `</xsd:openingAccountPurpose>
            <xsd:periodOfWork>` + req.PeriodOfWork + `</xsd:periodOfWork>
            <xsd:phoneNumber>` + req.PhoneNumber + `</xsd:phoneNumber>
            <xsd:placeOfBirth>` + req.PlaceOfBirth + `</xsd:placeOfBirth>
            <xsd:productType>` + req.ProductType + `</xsd:productType>
            <xsd:religion>` + req.Religion + `</xsd:religion>
            <xsd:requestBy>` + req.RequestBy + `</xsd:requestBy>
            <xsd:rtNo>` + req.RtNo + `</xsd:rtNo>
            <xsd:rwNo>` + req.RwNo + `</xsd:rwNo>
            <xsd:sex>` + req.Sex + `</xsd:sex>
            <xsd:typeOfWork>` + req.TypeOfWork + `</xsd:typeOfWork>
            <xsd:workingPosition>` + req.WorkingPosition + `</xsd:workingPosition>
            <xsd:zipCode>` + req.ZipCode + `</xsd:zipCode>
         </ser:request>
      </ser:createCustDataAndAccount>
   </soap:Body>
</soap:Envelope>`
	return stringRequest
}
