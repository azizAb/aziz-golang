package controller

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"ms-account/accountmain/client"
	"ms-account/accountmain/config"
	"ms-account/accountmain/model"
	"ms-account/accountmain/soaptemplate"

	"github.com/darahayes/go-boom"

	"time"

	"net/http"
	"strings"
)

func GetInquiryID(w http.ResponseWriter, r *http.Request) {

	var reqData model.RequestAccount

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			log.Println("Error parse request body to json. ", err.Error())
			boom.Internal(w, "Error parse request body to json.")
			return
		}
	}

	jsonTest, _ := json.Marshal(&reqData.ReqInquiry)
	log.Println("parse body to json :", string(jsonTest))
	jsonTest2, _ := json.Marshal(&reqData)
	log.Println("parse body to json all :", string(jsonTest2))

	payload := []byte(strings.TrimSpace(soaptemplate.InquiriIDString(reqData.ReqInquiry)))

	// call service controller
	log.Println("Req Payloadnya ", string(payload))
	response, err := HttpPost(payload, config.Urn.InquiryID, &w)
	if err != nil {
		log.Println("service controller error. ", err.Error())
		boom.BadGateway(w, "Failed to Connect.")
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse response Body
	// log.Println("-> Retrieving and parsing the response")
	// bodyBytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// log.Println(bodyString)

	// read and parse the response body
	result := new(soaptemplate.InquiryIDResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Println("Error on unmarshaling xml. ", err.Error())
		boom.Internal(w, "Error on unmarshaling xml.")
		return
	}

	log.Println("-> get result", result)

	// Save to Report Table
	log.Println("Save to Report")

	if reqData.UserId != "" {
		t := time.Now()
		ReqSaveReport := model.Report{
			Timestamp:  t.Format("2006-01-02 15:04:05.000-07:00"),
			Userid:     reqData.UserId,
			Activities: "Inquiry ID",
			Cif:        reqData.Cif,
			Accountno:  reqData.Accountno,
			Ktpid:      reqData.KtpId,
			Status:     "Inquery",
		}

		client.SaveReport(ReqSaveReport)

	}

	// Return Json Response
	js, err := json.Marshal(result.Body.InquiryIDResponse.Return)

	if err != nil {
		log.Println("Error on unmarshaling json. ", err.Error())
		boom.Internal(w, "Error on unmarshaling json.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
