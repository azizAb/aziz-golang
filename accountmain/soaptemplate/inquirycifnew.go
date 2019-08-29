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
				Address                 string `xml:"address" json:"address"`
				Address2                string `xml:"address2" json:"address2"`
				Address3                string `xml:"address3" json:"address3"`
				Address4                string `xml:"address4" json:"address4"`
				AddressRT               string `xml:"addressRT" json:"addressRT"`
				AddressRW               string `xml:"addressRW" json:"addressRW"`
				Agama                   string `xml:"agama" json:"agama"`
				AlamatSuratMenyurat     string `xml:"alamatSuratMenyurat" json:"alamatSuratMenyurat"`
				BirthDate               string `xml:"birthDate" json:"birthDate"`
				CcCode                  string `xml:"ccCode" json:"ccCode"`
				CifNo                   string `xml:"cifNo" json:"cifNo"`
				City                    string `xml:"city" json:"city"`
				DailyTrx                string `xml:"dailyTrx" json:"dailyTrx"`
				Education               string `xml:"education" json:"education"`
				EducationDesc           string `xml:"educationDesc" json:"educationDesc"`
				Email                   string `xml:"email" json:"email"`
				Hobby                   string `xml:"hobby" json:"hobby"`
				HobbyDesc               string `xml:"hobbyDesc" json:"hobbyDesc"`
				IdNo                    string `xml:"idNo" json:"idNo"`
				IdType                  string `xml:"idType" json:"idType"`
				Jabatan                 string `xml:"jabatan" json:"jabatan"`
				JabatanDesc             string `xml:"jabatanDesc" json:"jabatanDesc"`
				Kecamatan               string `xml:"kecamatan" json:"kecamatan"`
				Kelurahan               string `xml:"kelurahan" json:"kelurahan"`
				MaritalStatus           string `xml:"maritalStatus" json:"maritalStatus"`
				MotherName              string `xml:"motherName" json:"motherName"`
				Nama                    string `xml:"nama" json:"nama"`
				NoHP                    string `xml:"noHP" json:"noHP"`
				NoTelp                  string `xml:"noTelp" json:"noTelp"`
				OfficeAddress           string `xml:"officeAddress" json:"officeAddress"`
				OfficeCity              string `xml:"officeCity" json:"officeCity"`
				OfficeKecamatan         string `xml:"officeKecamatan" json:"officeKecamatan"`
				OfficeKelurahan         string `xml:"officeKelurahan" json:"officeKelurahan"`
				OfficeName              string `xml:"officeName" json:"officeName"`
				OfficeNoTelp            string `xml:"officeNoTelp" json:"officeNoTelp"`
				OfficeZipCode           string `xml:"officeZipCode" json:"officeZipCode"`
				Pekerjaan               string `xml:"pekerjaan" json:"pekerjaan"`
				PekerjaanDesc           string `xml:"pekerjaanDesc" json:"pekerjaanDesc"`
				PenghasilanPerbulan     string `xml:"penghasilanPerbulan" json:"penghasilanPerbulan"`
				PepFlag                 string `xml:"pepFlag" json:"pepFlag"`
				PepFlagKeluarga         string `xml:"pepFlagKeluarga" json:"pepFlagKeluarga"`
				PepStatusKeluarga       string `xml:"pepStatusKeluarga" json:"pepStatusKeluarga"`
				PlaceOfBirth            string `xml:"placeOfBirth" json:"placeOfBirth"`
				Province                string `xml:"province" json:"province"`
				ResponseCode            string `xml:"responseCode" json:"responseCode"`
				ResponseDesc            string `xml:"responseDesc" json:"responseDesc"`
				Sex                     string `xml:"sex" json:"sex"`
				SumberUtama             string `xml:"sumberUtama" json:"sumberUtama"`
				SumberUtamaDesc         string `xml:"sumberUtamaDesc" json:"sumberUtamaDesc"`
				TujuanPembukaanRekening string `xml:"tujuanPembukaanRekening" json:"tujuanPembukaanRekening"`
				TypeOfWork              string `xml:"typeOfWork" json:"typeOfWork"`
				ZipCode                 string `xml:"zipCode" json:"zipCode"`
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
