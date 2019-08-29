package main

import (
	"log"
	"net/http"

	"ms-account/accountmain/config"
	"ms-account/accountmain/controller"
)

func main() {

	http.HandleFunc("/getid", controller.GetInquiryID)
	http.HandleFunc("/newcifaccount", controller.CustDataandAccount)
	http.HandleFunc("/newaccount", controller.CreateAccount)
	http.HandleFunc("/inquirycif", controller.InquiryCIFNew)
	http.HandleFunc("/inquirysaca", controller.InquirySACADetail)
	http.HandleFunc("/getcifparam", controller.GetCIFParameter)
	http.HandleFunc("/getkodepos", controller.GetKodePos)

	var address = config.Http.ServingPort
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("http error ", err.Error())
	}

}
