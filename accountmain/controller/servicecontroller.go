package controller

import (
	"bytes"
	"log"
	"ms-account/accountmain/config"
	"net/http"
)

func HttpPost(payload []byte, soapUrn string) (response *http.Response, err error) {

	method := config.Http.Method
	url := config.Url

	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapUrn)

	// prepare the client request
	client := &http.Client{}

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
	}

	return res, err
}
