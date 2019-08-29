package controller

import (
	"bytes"
	"log"
	"ms-account/accountmain/config"
	"net/http"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

func HttpPost(payload []byte, soapUrn string, w *http.ResponseWriter) (response *http.Response, err error) {

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	method := config.Http.Method
	url := config.Url.Soap

	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		log.Println("Error on creating request object. ", err.Error())
	}

	req.Header.Set("Content-type", "application/soap+xml")
	req.Header.Set("SOAPAction", soapUrn)

	// prepare the client request
	client := &http.Client{}

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		// log.Fatal("Error on dispatching request. ", err.Error())
	}

	return res, err
}
