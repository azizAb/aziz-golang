package client

import (
	"context"
	"log"
	"ms-account/accountmain/config"
	"ms-account/accountmain/model"
	"time"

	"google.golang.org/grpc"
)

func SaveReport(saveReq model.Report) {
	log.Println("Connecting to ", config.Url.ReportAddress)
	conn, err := grpc.Dial(config.Url.ReportAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	log.Println("Connected")

	defer conn.Close()
	client := model.NewReportsClient(conn)

	log.Printf("Send Request")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.AddReport(ctx, &saveReq)

	log.Printf("Finish")

}

func SaveApproval(saveReq model.CreateAccount) {
	conn, err := grpc.Dial(config.Url.ApprovalAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := model.NewCreateAccountsClient(conn)

	// user1 := model.CreateAccount{
	// 	RequestId: "00002",
	// 	Creator:   "Cireng",
	// 	Content:   "{\"accountType\":\"S\",\"cifNo\":\"ACFB633\",\"customerName\":\"AWANG HERMAWAN\",\"productType\":\"SU\",\"requestBy\":\"02060\"}",
	// 	// Content: "{\"AccountType\":\"S\",\"Address\":\"Cinere\",\"CellPhoneNumber\":\"089630823030\",\"CustomerName\":\"Abdul Asis\",\"DateOfBirth\":\"01/01/1968\","+
	// 	// "\"Education\":\"Otodidak\",\"EmailAddress\":\"abasis@com\",\"FieldWork\":\"FieldWork\",\"IdNo\":\"008\",\"IncomeSource\":\"IncomeSource\"}",
	// 	Status: "Exist",
	// }

	log.Printf("jalanin client")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Register(ctx, &saveReq)

	log.Printf("kelar client")

}

// func main() {
// 	log.Printf("Connecting to ", config.Url.ApprovalAddress)
// 	conn, err := grpc.Dial(config.Url.ApprovalAddress, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	log.Printf("Connected")

// 	defer conn.Close()
// 	client := model.NewCreateAccountsClient(conn)

// 	user1 := model.CreateAccount{
// 		RequestId: "00002",
// 		Creator:   "Cireng",
// 		Content:   "{\"accountType\":\"S\",\"cifNo\":\"ACFB633\",\"customerName\":\"AWANG HERMAWAN\",\"productType\":\"SU\",\"requestBy\":\"02060\"}",
// 		// Content: "{\"AccountType\":\"S\",\"Address\":\"Cinere\",\"CellPhoneNumber\":\"089630823030\",\"CustomerName\":\"Abdul Asis\",\"DateOfBirth\":\"01/01/1968\","+
// 		// "\"Education\":\"Otodidak\",\"EmailAddress\":\"abasis@com\",\"FieldWork\":\"FieldWork\",\"IdNo\":\"008\",\"IncomeSource\":\"IncomeSource\"}",
// 		Status: "Exist",
// 	}

// 	log.Printf("jalanin client")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	client.Register(ctx, &user1)

// 	log.Printf("kelar client")

// }

// func main() {
// 	log.Println("Connecting to ", config.Url.ReportAddress)
// 	conn, err := grpc.Dial(config.Url.ReportAddress, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	log.Println("Success Connect")

// 	defer conn.Close()
// 	client := model.NewReportsClient(conn)

// 	log.Printf("Sending Request")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

// 	t := time.Now()
// 	ReqSaveReport := model.Report{
// 		Timestamp:     t.Format("2006-01-02 15:04:05.000-07:00"),
// 		Userid:        "4212",
// 		Activities:    "Inquiry ID",
// 		Applicationid: "aziz",
// 		Cif:           "cif",
// 		Accountno:     "reqData.Accountno",
// 		Ktpid:         " reqData.Ktpid",
// 		Status:        "Query",
// 	}

// 	defer cancel()
// 	client.AddReport(ctx, &ReqSaveReport)

// 	log.Printf("Finish Process")

// }
