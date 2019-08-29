package controller

import (
	"encoding/json"
	"log"
	"ms-account/accountmain/client"
	"ms-account/accountmain/model"
	"strconv"
	"time"

	"github.com/darahayes/go-boom"

	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	(w).Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	var reqData model.RequestAccount

	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			log.Println("Error parse request body to json. ", err.Error())
			boom.Internal(w, "Error parse request body to json.")
			return
		}
	}

	jsonTest, _ := json.Marshal(&reqData.ReqAccount)
	jsonTest2, _ := json.Marshal(&reqData)
	log.Println("parse body to json All :", string(jsonTest2))
	log.Println("parse body to json :", string(jsonTest))

	// Save to Report table
	log.Println("Save to Report")
	if reqData.UserId != "" {
		t := time.Now()
		ReqSaveReport := model.Report{
			Timestamp:     t.Format("2006-01-02 15:04:05.000-07:00"),
			Userid:        reqData.UserId,
			Activities:    "New Account",
			Applicationid: reqData.UserId + "-" + t.Format("20060102150405"),
			Cif:           reqData.Cif,
			Accountno:     reqData.Accountno,
			Ktpid:         reqData.KtpId,
			Status:        "Send to Approval",
		}
		log.Println("Req to Save report ", ReqSaveReport)
		client.SaveReport(ReqSaveReport)

		// Save Table Report
		log.Println("Save to Approval")
		var status string

		if reqData.Cif != "" {
			status = "Exist"
		} else {
			status = "Not Exist"
		}

		reqSaveApproval := model.CreateAccount{
			RequestId: strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
			Creator:   reqData.UserId,
			Content:   string(reqData.Payload),
			Status:    status,
		}
		log.Println("Req to Save Approval ", reqSaveApproval)
		client.SaveApproval(reqSaveApproval)

	}
	//  else {
	// 	log.Println("Userid null. ")
	// 	boom.Internal(w, "Userid is null.")
	// 	return
	// }

}
