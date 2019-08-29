package model

type KeyValuePair struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
	ResponseItem string `json:"responseItem"`
}

type RespCIFParameter struct {
	JenisPekerjaan      KeyValuePair `json:"jenisPekerjaan"`
	PenghasilanPerbulan KeyValuePair `json:"penghasilanPerbulan"`
	TransaksiHarian     KeyValuePair `json:"transaksiHarian"`
	SumberUtama         KeyValuePair `json:"sumberUtama"`
	TujuanBukaRekening  KeyValuePair `json:"tujuanBukaRekening"`
}
