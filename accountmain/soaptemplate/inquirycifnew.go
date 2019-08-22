package soaptemplate

import (
	"encoding/xml"
)

type InquiryCIFNewReq struct {
	Branch    string
	CifNo     string
	RequestBy string
	UserId    string
}

type InquiryCIFNewResp struct {
	XMLName xml.Name
	Body    struct {
		XMLName               xml.Name
		InquiryCIFNewResponse struct {
			XMLName xml.Name
			Return  struct {
				Address                 string `xml:"address"`
				Address2                string `xml:"address2"`
				Address3                string `xml:"address3"`
				Address4                string `xml:"address4"`
				AddressRT               string `xml:"addressRT"`
				AddressRW               string `xml:"addressRW"`
				Agama                   string `xml:"agama"`
				AlamatSuratMenyurat     string `xml:"alamatSuratMenyurat"`
				BirthDate               string `xml:"birthDate"`
				CcCode                  string `xml:"ccCode"`
				CifNo                   string `xml:"cifNo"`
				City                    string `xml:"city"`
				DailyTrx                string `xml:"dailyTrx"`
				Education               string `xml:"education"`
				EducationDesc           string `xml:"educationDesc"`
				Email                   string `xml:"email"`
				Hobby                   string `xml:"hobby"`
				HobbyDesc               string `xml:"hobbyDesc"`
				IdNo                    string `xml:"idNo"`
				IdType                  string `xml:"idType"`
				Jabatan                 string `xml:"jabatan"`
				JabatanDesc             string `xml:"jabatanDesc"`
				Kecamatan               string `xml:"kecamatan"`
				Kelurahan               string `xml:"kelurahan"`
				MaritalStatus           string `xml:"maritalStatus"`
				MotherName              string `xml:"motherName"`
				Nama                    string `xml:"nama"`
				NoHP                    string `xml:"noHP"`
				NoTelp                  string `xml:"noTelp"`
				OfficeAddress           string `xml:"officeAddress"`
				OfficeCity              string `xml:"officeCity"`
				OfficeKecamatan         string `xml:"officeKecamatan"`
				OfficeKelurahan         string `xml:"officeKelurahan"`
				OfficeName              string `xml:"officeName"`
				OfficeNoTelp            string `xml:"officeNoTelp"`
				OfficeZipCode           string `xml:"officeZipCode"`
				Pekerjaan               string `xml:"pekerjaan"`
				PekerjaanDesc           string `xml:"pekerjaanDesc"`
				PenghasilanPerbulan     string `xml:"penghasilanPerbulan"`
				PepFlag                 string `xml:"pepFlag"`
				PepFlagKeluarga         string `xml:"pepFlagKeluarga"`
				PepStatusKeluarga       string `xml:"pepStatusKeluarga"`
				PlaceOfBirth            string `xml:"placeOfBirth"`
				Province                string `xml:"province"`
				ResponseCode            string `xml:"responseCode"`
				ResponseDesc            string `xml:"responseDesc"`
				Sex                     string `xml:"sex"`
				SumberUtama             string `xml:"sumberUtama"`
				SumberUtamaDesc         string `xml:"sumberUtamaDesc"`
				TujuanPembukaanRekening string `xml:"tujuanPembukaanRekening"`
				TypeOfWork              string `xml:"typeOfWork"`
				ZipCode                 string `xml:"zipCode"`
			} `xml:"return"`
		} `xml:"inquiryCIFNewResponse"`
	}
}

func InquiryCIFNewString(req InquiryCIFNewReq) (resp string) {
	var stringRequest = `
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ser="http://service.bri.com" xmlns:xsd="http://service.model.bri.com/xsd">
   <soap:Header/>
   <soap:Body>
      <ser:inquiryCIFNew>
         <ser:request>
            <xsd:branch>` + req.Branch + `</xsd:branch>
            <xsd:cifNo>` + req.CifNo + `</xsd:cifNo>
            <xsd:requestBy>` + req.RequestBy + `</xsd:requestBy>
            <xsd:userId>` + req.UserId + `</xsd:userId>
         </ser:request>
      </ser:inquiryCIFNew>
   </soap:Body>
	</soap:Envelope>`
	return stringRequest
}
