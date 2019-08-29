package model

import (
	. "ms-account/accountmain/soaptemplate"
)

type RequestAccount struct {
	UserId                string
	Activities            string
	Cif                   string
	Accountno             string
	KtpId                 string
	Payload               string
	ReqInquiry            InquiryIDReq          `json:"reqInquiry"`
	ReqCustDataAndAccount CustDataAndAccountReq `json:"reqCustDataAndAccount"`
	ReqAccount            CreateAccountReq      `json:"reqAccount"`
	ReqInquiryCIF         InquiryCIFNewReq      `json:"reqInquiryCIF"`
	ReqInquirySACA        InquirySACADetailReq  `json:"reqInquirySACA"`
}
