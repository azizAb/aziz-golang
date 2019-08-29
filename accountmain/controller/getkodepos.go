package controller

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"ms-account/accountmain/config"
	"ms-account/accountmain/soaptemplate"
	"net/http"
	"strings"

	"github.com/darahayes/go-boom"
)

func GetKodePos(w http.ResponseWriter, r *http.Request) {

	reqKopos := new(soaptemplate.KodeposInquiryReq)

	kopos, ok := r.URL.Query()["kodepos"]
	if !ok || len(kopos[0]) < 1 {
		log.Println("Url Param 'kodepos' is missing")
		return
	}

	reqby, ok := r.URL.Query()["requestby"]
	if !ok || len(reqby[0]) < 1 {
		log.Println("Url Param 'requestby' is missing")
		return
	}

	log.Println("Url Param 'kodepos' is: " + string(kopos[0]))
	log.Println("Url Param 'requestby' is: " + string(reqby[0]))

	reqKopos.RequestBy = reqby[0]
	reqKopos.Type = "kodepos"
	reqKopos.Value = kopos[0]

	payload := []byte(strings.TrimSpace(soaptemplate.KodeposInquiryString(reqKopos)))

	// call service controller
	log.Println("Req Payloadnya ", string(payload))
	response, err := HttpPost(payload, config.Urn.KodePos, &w)
	if err != nil {
		log.Println("service controller error. ", err.Error())
		boom.BadGateway(w, "Failed to Connect.")
		return
	}

	// // Parse response Body
	// // log.Println("-> Retrieving and parsing the response")
	// // bodyBytes, err := ioutil.ReadAll(response.Body)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // bodyString := string(bodyBytes)
	// // log.Println(bodyString)

	// read and parse the response body
	result := new(soaptemplate.KodeposInquiryResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Println("Error on unmarshaling xml. ", err.Error())
		boom.Internal(w, "Error on unmarshaling xml.")
		return
	}

	log.Println("-> get result", result)

	// Return Json Response
	js, err := json.Marshal(result.Body.KodePosInquiryResponse.Return)

	if err != nil {
		log.Println("Error on unmarshaling json. ", err.Error())
		boom.Internal(w, "Error on unmarshaling json.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
