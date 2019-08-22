package controller

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"ms-account/accountmain/config"
	"ms-account/accountmain/soaptemplate"

	"net/http"
	"strings"
)

func GetInquiryID(w http.ResponseWriter, r *http.Request) {

	var reqData soaptemplate.InquiryIDReq

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("parse body to json :", reqData)

	payload := []byte(strings.TrimSpace(soaptemplate.InquiriIDString(reqData)))

	// call service controller
	response, err := HttpPost(payload, config.Urn.InquiryID)
	if err != nil {
		log.Fatal("service controller error. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(soaptemplate.InquiryIDResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> get result", result)

	// Return Json Response
	js, err := json.Marshal(result.Body.InquiryIDResponse.Return)

	if err != nil {
		log.Fatal("Error on marshaling json. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func CustDataandAccount(w http.ResponseWriter, r *http.Request) {

	var reqData soaptemplate.CustDataAndAccountReq

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("parse body to json :", reqData)

	payload := []byte(strings.TrimSpace(soaptemplate.CustDataAndAccountString(reqData)))

	// call service controller
	response, err := HttpPost(payload, config.Urn.CreateCustDataAndAccount)
	if err != nil {
		log.Fatal("service controller error. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(soaptemplate.CustDataAndAccountResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> get result", result)

	// Return Json Response
	js, err := json.Marshal(result.Body.CreateCustDataAndAccountResponse.Return)

	if err != nil {
		log.Fatal("Error on marshaling json. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var reqData soaptemplate.CreateAccountReq

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("parse body to json :", reqData)

	payload := []byte(strings.TrimSpace(soaptemplate.CreateAccountString(reqData)))

	// call service controller
	response, err := HttpPost(payload, config.Urn.CreateAccount)
	if err != nil {
		log.Fatal("service controller error. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(soaptemplate.CreateAccountResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> get result", result)

	// Return Json Response
	js, err := json.Marshal(result.Body.CreateAccountResponse.Return)

	if err != nil {
		log.Fatal("Error on marshaling json. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func InquiryCIFNew(w http.ResponseWriter, r *http.Request) {

	var reqData soaptemplate.InquiryCIFNewReq

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("parse body to json :", reqData)

	payload := []byte(strings.TrimSpace(soaptemplate.InquiryCIFNewString(reqData)))

	// call service controller
	response, err := HttpPost(payload, config.Urn.InquiryCIFNew)
	if err != nil {
		log.Fatal("service controller error. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(soaptemplate.InquiryCIFNewResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> get result", result)

	// Return Json Response
	js, err := json.Marshal(result.Body.InquiryCIFNewResponse.Return)

	if err != nil {
		log.Fatal("Error on marshaling json. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func InquirySACADetail(w http.ResponseWriter, r *http.Request) {

	var reqData soaptemplate.InquirySACADetailReq

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("parse body to json :", reqData)

	payload := []byte(strings.TrimSpace(soaptemplate.InquirySACADetailString(reqData)))

	// call service controller
	response, err := HttpPost(payload, config.Urn.InquirySACADetail)
	if err != nil {
		log.Fatal("service controller error. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(soaptemplate.InquirySACADetailResp)
	err = xml.NewDecoder(response.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("-> get result", result)

	// Return Json Response
	js, err := json.Marshal(result.Body.InquirySACADetailResponse.Return)

	if err != nil {
		log.Fatal("Error on marshaling json. ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
