package controller

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"ms-account/accountmain/client"
	"ms-account/accountmain/config"
	"ms-account/accountmain/model"
	"ms-account/accountmain/soaptemplate"
	"time"

	"github.com/darahayes/go-boom"

	"net/http"
	"strings"
)

func InquiryCIFNew(w http.ResponseWriter, r *http.Request) {

	var reqData model.RequestAccount

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			log.Println("Error parse request body to json. ", err.Error())
			boom.Internal(w, "Error parse request body to json.")
			return
		}
	}

	jsonTest, _ := json.Marshal(&reqData.ReqInquiryCIF)
	log.Println("parse body to json :", string(jsonTest))

	payload := []byte(strings.TrimSpace(soaptemplate.InquiryCIFNewString(reqData.ReqInquiryCIF)))

	// call service controller
	response, err := HttpPost(payload, config.Urn.InquiryCIFNew, &w)
	if err != nil {
		log.Println("service controller error. ", err.Error())
		boom.BadGateway(w, "Failed to Connect.")
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// // Parse response Body
	// log.Println("-> Retrieving and parsing the response")
	// bodyBytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// log.Println(bodyString)

	// read and parse the response body
	result := new(soaptemplate.InquiryCIFNewResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Println("Error on unmarshaling xml. ", err.Error())
		boom.Internal(w, "Error on unmarshaling xml.")
		return
	}

	log.Println("-> get result", result)

	if reqData.UserId != "" {
		log.Println("Save to Report")

		t := time.Now()
		ReqSaveReport := model.Report{
			Timestamp:  t.Format("2006-01-02 15:04:05.000-07:00"),
			Userid:     reqData.UserId,
			Activities: "Inquiry CIF New",
			Cif:        reqData.Cif,
			Accountno:  reqData.Accountno,
			Ktpid:      reqData.KtpId,
			Status:     "Inquiry",
		}

		client.SaveReport(ReqSaveReport)
	}

	// Return Json Response
	js, err := json.Marshal(result.Body.InquiryCIFNewResponse.Return)

	if err != nil {
		log.Println("Error on unmarshaling json. ", err.Error())
		boom.Internal(w, "Error on unmarshaling json.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
