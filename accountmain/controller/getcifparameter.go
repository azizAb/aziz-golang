package controller

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"ms-account/accountmain/config"
	"ms-account/accountmain/model"
	"ms-account/accountmain/soaptemplate"

	"github.com/darahayes/go-boom"

	"net/http"
	"strings"
)

func GetCIFParameter(w http.ResponseWriter, r *http.Request) {

	reqData := new(soaptemplate.CifParameterReq)

	reqby, ok := r.URL.Query()["requestby"]
	if !ok || len(reqby[0]) < 1 {
		log.Println("Url Param 'requestby' is missing")
		return
	}

	log.Println("Url Param 'requestby' is: " + string(reqby[0]))

	reqData.RequestBy = reqby[0]

	// var reqData soaptemplate.CifParameterReq

	// if r.Method == "POST" {

	// 	err := json.NewDecoder(r.Body).Decode(&reqData)
	// 	if err != nil {
	// 		log.Println("Error parse request body to json. ", err.Error())
	// 		boom.Internal(w, "Error parse request body to json.")
	// 		return
	// 	}
	// }

	var js []byte
	var err error
	resultCIFParam := new(model.RespCIFParameter)

	for i := 0; i < 5; i++ {
		reqData.Key = config.KeyCifParam[i]

		jsonTest, _ := json.Marshal(&reqData)
		log.Println("parse body to json :", string(jsonTest))

		payload := []byte(strings.TrimSpace(soaptemplate.CifParameterString(reqData)))

		// call service controller
		response, err := HttpPost(payload, config.Urn.GetCIFParam, &w)
		if err != nil {
			log.Println("service controller error. ", err.Error())
			boom.BadGateway(w, "Failed to Connect.")
			return
		}

		log.Println("-> Retrieving and parsing the response")

		// read and parse the response body
		result := new(soaptemplate.CifParameterResp)
		err = xml.NewDecoder(response.Body).Decode(result)
		if err != nil {
			log.Println("Error on unmarshaling xml. ", err.Error())
			boom.Internal(w, "Error on unmarshaling xml.")
			return
		}

		log.Println("-> get result", result)

		switch config.KeyCifParam[i] {
		case "jenispekerjaan":
			resultCIFParam.JenisPekerjaan.ResponseCode = result.Body.GetCIFParameterResponse.Return.ResponseCode
			resultCIFParam.JenisPekerjaan.ResponseDesc = result.Body.GetCIFParameterResponse.Return.ResponseDesc
			resultCIFParam.JenisPekerjaan.ResponseItem = result.Body.GetCIFParameterResponse.Return.ResponseItem
		case "listpenghasilan":
			resultCIFParam.PenghasilanPerbulan.ResponseCode = result.Body.GetCIFParameterResponse.Return.ResponseCode
			resultCIFParam.PenghasilanPerbulan.ResponseDesc = result.Body.GetCIFParameterResponse.Return.ResponseDesc
			resultCIFParam.PenghasilanPerbulan.ResponseItem = result.Body.GetCIFParameterResponse.Return.ResponseItem
		case "listtransaksiharian":
			resultCIFParam.TransaksiHarian.ResponseCode = result.Body.GetCIFParameterResponse.Return.ResponseCode
			resultCIFParam.TransaksiHarian.ResponseDesc = result.Body.GetCIFParameterResponse.Return.ResponseDesc
			resultCIFParam.TransaksiHarian.ResponseItem = result.Body.GetCIFParameterResponse.Return.ResponseItem
		case "listsumberutama":
			resultCIFParam.SumberUtama.ResponseCode = result.Body.GetCIFParameterResponse.Return.ResponseCode
			resultCIFParam.SumberUtama.ResponseDesc = result.Body.GetCIFParameterResponse.Return.ResponseDesc
			resultCIFParam.SumberUtama.ResponseItem = result.Body.GetCIFParameterResponse.Return.ResponseItem
		case "listtujuanpembukaanrekening":
			resultCIFParam.TujuanBukaRekening.ResponseCode = result.Body.GetCIFParameterResponse.Return.ResponseCode
			resultCIFParam.TujuanBukaRekening.ResponseDesc = result.Body.GetCIFParameterResponse.Return.ResponseDesc
			resultCIFParam.TujuanBukaRekening.ResponseItem = result.Body.GetCIFParameterResponse.Return.ResponseItem
		}

	}

	// Return Json Response
	js, err = json.Marshal(resultCIFParam)

	if err != nil {
		log.Println("Error on unmarshaling json. ", err.Error())
		boom.Internal(w, "Error on unmarshaling json.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
