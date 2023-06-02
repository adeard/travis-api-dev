package domain

type ZtsTravis struct {
	MANDT           string  `json:"mandt"`
	TASKID          string  `json:"taskid"`
	VKORG           string  `json:"vkorg"`
	WERKS           string  `json:"werks"`
	VBELN           string  `json:"vbeln"`
	POSNR           string  `json:"posnr"`
	PICK_LOCATION   string  `json:"pick_location"`
	JENIS_KIRIM     string  `json:"jenis_kirim"`
	DO_INDUK        string  `json:"do_induk"`
	BLDAT           string  `json:"bldat"`
	REQ_DLV_DATE    string  `json:"req_dlv_date"`
	REQ_DLV_TIME1   string  `json:"req_dlv_time1"`
	REQ_DLV_TIME2   string  `json:"req_dlv_time2"`
	NOTES_FR_SALES  string  `json:"notes_fr_sales"`
	NOTES_FR_TRANSP string  `json:"notes_fr_transp"`
	SOLDTO          string  `json:"soldto"`
	SHIPTO          string  `json:"shipto"`
	REF_PO          string  `json:"ref_po"`
	MATNR           string  `json:"matnr"`
	MAT_DESC        string  `json:"mat_desc"`
	LFIMG           float32 `json:"lfimg"`
	VRKME           string  `json:"vrkme"`
	BRGEW           float32 `json:"brgew"`
	VOLUM           float32 `json:"volum"`
	LFIMG_RECEIVED  float32 `json:"lfimg_received"`
	VRKME_RECEIVED  string  `json:"vrkme_received"`
	REASON_DIFF     string  `json:"reason_diff"`
	RECIPIENT_NAME  string  `json:"recipient_name"`
	RECEIVE_DATE    string  `json:"receive_date"`
	RECEIVE_TIME    string  `json:"receive_time"`
	VENDOR          string  `json:"vendor"`
	VEHICLE_TYPE    string  `json:"vehicle_type"`
	DRIVER_NAME     string  `json:"driver_name"`
	VEHICLE_NO      string  `json:"vehicle_no"`
	LPB_NUMBER      string  `json:"lpb_number"`
	TASK_STATUS     string  `json:"task_status"`
	ERDAT           string  `json:"erdat"`
	ERZET           string  `json:"erzet"`
	ERNAM           string  `json:"ernam"`
	AEDAT           string  `json:"aedat"`
	AEZET           string  `json:"aezet"`
	AENAM           string  `json:"aenam"`
	SEL             string  `json:"sel"`
}

type ZtsTravisRequest struct {
	MANDT           string  `json:"mandt"`
	TASKID          string  `json:"taskid"`
	VKORG           string  `json:"vkorg"`
	WERKS           string  `json:"werks"`
	VBELN           string  `json:"vbeln"`
	POSNR           string  `json:"posnr"`
	PICK_LOCATION   string  `json:"pick_location"`
	JENIS_KIRIM     string  `json:"jenis_kirim"`
	DO_INDUK        string  `json:"do_induk"`
	BLDAT           string  `json:"bldat"`
	REQ_DLV_DATE    string  `json:"req_dlv_date"`
	REQ_DLV_TIME1   string  `json:"req_dlv_time1"`
	REQ_DLV_TIME2   string  `json:"req_dlv_time2"`
	NOTES_FR_SALES  string  `json:"notes_fr_sales"`
	NOTES_FR_TRANSP string  `json:"notes_fr_transp"`
	SOLDTO          string  `json:"soldto"`
	SHIPTO          string  `json:"shipto"`
	REF_PO          string  `json:"ref_po"`
	MATNR           string  `json:"matnr"`
	MAT_DESC        string  `json:"mat_desc"`
	LFIMG           float32 `json:"lfimg"`
	VRKME           string  `json:"vrkme"`
	BRGEW           float32 `json:"brgew"`
	VOLUM           float32 `json:"volum"`
	LFIMG_RECEIVED  float32 `json:"lfimg_received"`
	VRKME_RECEIVED  string  `json:"vrkme_received"`
	REASON_DIFF     string  `json:"reason_diff"`
	RECIPIENT_NAME  string  `json:"recipient_name"`
	RECEIVE_DATE    string  `json:"receive_date"`
	RECEIVE_TIME    string  `json:"receive_time"`
	VENDOR          string  `json:"vendor"`
	VEHICLE_TYPE    string  `json:"vehicle_type"`
	DRIVER_NAME     string  `json:"driver_name"`
	VEHICLE_NO      string  `json:"vehicle_no"`
	LPB_NUMBER      string  `json:"lpb_number"`
	TASK_STATUS     string  `json:"task_status"`
	ERDAT           string  `json:"erdat"`
	ERZET           string  `json:"erzet"`
	ERNAM           string  `json:"ernam"`
	AEDAT           string  `json:"aedat"`
	AEZET           string  `json:"aezet"`
	AENAM           string  `json:"aenam"`
	SEL             string  `json:"sel"`
}

type ZtsTravisResponse struct {
	MANDT           string  `json:"mandt"`
	TASKID          string  `json:"taskid"`
	VKORG           string  `json:"vkorg"`
	WERKS           string  `json:"werks"`
	VBELN           string  `json:"vbeln"`
	POSNR           string  `json:"posnr"`
	PICK_LOCATION   string  `json:"pick_location"`
	JENIS_KIRIM     string  `json:"jenis_kirim"`
	DO_INDUK        string  `json:"do_induk"`
	BLDAT           string  `json:"bldat"`
	REQ_DLV_DATE    string  `json:"req_dlv_date"`
	REQ_DLV_TIME1   string  `json:"req_dlv_time1"`
	REQ_DLV_TIME2   string  `json:"req_dlv_time2"`
	NOTES_FR_SALES  string  `json:"notes_fr_sales"`
	NOTES_FR_TRANSP string  `json:"notes_fr_transp"`
	SOLDTO          string  `json:"soldto"`
	SHIPTO          string  `json:"shipto"`
	REF_PO          string  `json:"ref_po"`
	MATNR           string  `json:"matnr"`
	MAT_DESC        string  `json:"mat_desc"`
	LFIMG           float32 `json:"lfimg"`
	VRKME           string  `json:"vrkme"`
	BRGEW           float32 `json:"brgew"`
	VOLUM           float32 `json:"volum"`
	LFIMG_RECEIVED  float32 `json:"lfimg_received"`
	VRKME_RECEIVED  string  `json:"vrkme_received"`
	REASON_DIFF     string  `json:"reason_diff"`
	RECIPIENT_NAME  string  `json:"recipient_name"`
	RECEIVE_DATE    string  `json:"receive_date"`
	RECEIVE_TIME    string  `json:"receive_time"`
	VENDOR          string  `json:"vendor"`
	VEHICLE_TYPE    string  `json:"vehicle_type"`
	DRIVER_NAME     string  `json:"driver_name"`
	VEHICLE_NO      string  `json:"vehicle_no"`
	LPB_NUMBER      string  `json:"lpb_number"`
	TASK_STATUS     string  `json:"task_status"`
	ERDAT           string  `json:"erdat"`
	ERZET           string  `json:"erzet"`
	ERNAM           string  `json:"ernam"`
	AEDAT           string  `json:"aedat"`
	AEZET           string  `json:"aezet"`
	AENAM           string  `json:"aenam"`
	SEL             string  `json:"sel"`
}

type ZtsTravisFilter struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}
