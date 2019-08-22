package soaptemplate

import (
	"encoding/xml"
)

type CustDataAndAccountReq struct {
	AccountType            string
	Address                string
	CellPhoneNumber        string
	CustomerName           string
	DailyTransactionAmount string
	DateOfBirth            string
	Education              string
	EmailAddress           string
	FieldWork              string
	IdNo                   string
	IncomeSource           string
	MaritalStatus          string
	MonthlyIncome          string
	MotherName             string
	NpwpNo                 string
	OfficeAddress          string
	OfficeName             string
	OpeningAccountPurpose  string
	PeriodOfWork           string
	PhoneNumber            string
	PlaceOfBirth           string
	ProductType            string
	Religion               string
	RequestBy              string
	RtNo                   string
	RwNo                   string
	Sex                    string
	TypeOfWork             string
	WorkingPosition        string
	ZipCode                string
}

type CustDataAndAccountResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName                          xml.Name
		CreateCustDataAndAccountResponse struct {
			XMLName xml.Name
			Return  struct {
				AccountNumber string `xml:"accountNumber"`
				CifNo         string `xml:"cifNo"`
				ResponseCode  string `xml:"responseCode"`
				ResponseDesc  string `xml:"responseDesc"`
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
