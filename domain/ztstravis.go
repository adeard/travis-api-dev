package domain

type ZtsTravis struct {
	MANDT           string  `json:"mandt" gorm:"column:MANDT"`
	TASKID          string  `json:"taskid" gorm:"column:TASKID"`
	VKORG           string  `json:"vkorg" gorm:"column:VKORG"`
	WERKS           string  `json:"werks" gorm:"column:WERKS"`
	VBELN           string  `json:"vbeln" gorm:"column:VBELN"`
	POSNR           string  `json:"posnr" gorm:"column:POSNR"`
	PICK_LOCATION   string  `json:"pick_location" gorm:"column:PICK_LOCATION"`
	JENIS_KIRIM     string  `json:"jenis_kirim" gorm:"column:JENIS_KIRIM"`
	DO_INDUK        string  `json:"do_induk" gorm:"column:DO_INDUK"`
	BLDAT           string  `json:"bldat" gorm:"column:BLDAT"`
	REQ_DLV_DATE    string  `json:"req_dlv_date" gorm:"column:REQ_DLV_DATE"`
	REQ_DLV_TIME1   string  `json:"req_dlv_time1" gorm:"column:REQ_DLV_TIME1"`
	REQ_DLV_TIME2   string  `json:"req_dlv_time2" gorm:"column:REQ_DLV_TIME2"`
	NOTES_FR_SALES  string  `json:"notes_fr_sales" gorm:"column:NOTES_FR_SALES"`
	NOTES_FR_TRANSP string  `json:"notes_fr_transp" gorm:"column:NOTES_FR_TRANSP"`
	SOLDTO          string  `json:"soldto" gorm:"column:SOLDTO"`
	SHIPTO          string  `json:"shipto" gorm:"column:SHIPTO"`
	REF_PO          string  `json:"ref_po" gorm:"column:REF_PO"`
	MATNR           string  `json:"matnr" gorm:"column:MATNR"`
	MAT_DESC        string  `json:"mat_desc" gorm:"column:MAT_DESC"`
	LFIMG           float32 `json:"lfimg" gorm:"column:LFIMG"`
	VRKME           string  `json:"vrkme" gorm:"column:VRKME"`
	BRGEW           float32 `json:"brgew" gorm:"column:BRGEW"`
	VOLUM           float32 `json:"volum" gorm:"column:VOLUM"`
	LFIMG_RECEIVED  float32 `json:"lfimg_received" gorm:"column:LFIMG_RECEIVED"`
	VRKME_RECEIVED  string  `json:"vrkme_received" gorm:"column:VRKME_RECEIVED"`
	REASON_DIFF     string  `json:"reason_diff" gorm:"column:REASON_DIFF"`
	RECIPIENT_NAME  string  `json:"recipient_name" gorm:"column:RECIPIENT_NAME"`
	RECEIVE_DATE    string  `json:"receive_date" gorm:"column:RECEIVE_DATE"`
	RECEIVE_TIME    string  `json:"receive_time" gorm:"column:RECEIVE_TIME"`
	VENDOR          string  `json:"vendor" gorm:"column:VENDOR"`
	VEHICLE_TYPE    string  `json:"vehicle_type" gorm:"column:VEHICLE_TYPE"`
	DRIVER_NAME     string  `json:"driver_name" gorm:"column:DRIVER_NAME"`
	VEHICLE_NO      string  `json:"vehicle_no" gorm:"column:VEHICLE_NO"`
	LPB_NUMBER      string  `json:"lpb_number" gorm:"column:LPB_NUMBER"`
	TASK_STATUS     string  `json:"task_status" gorm:"column:TASK_STATUS"`
	ERDAT           string  `json:"erdat" gorm:"column:ERDAT"`
	ERZET           string  `json:"erzet" gorm:"column:ERZET"`
	ERNAM           string  `json:"ernam" gorm:"column:ERNAM"`
	AEDAT           string  `json:"aedat" gorm:"column:AEDAT"`
	AEZET           string  `json:"aezet" gorm:"column:AEZET"`
	AENAM           string  `json:"aenam" gorm:"column:AENAM"`
	SEL             string  `json:"sel" gorm:"column:SEL"`
}

func (ZtsTravis) TableName() string {
	return "ZTS_TRAVIS"
}

type ZtsTravisRequest struct {
	MANDT           string  `json:"mandt" validate:"required"`
	TASKID          string  `json:"taskid" validate:"required"`
	VKORG           string  `json:"vkorg" validate:"required"`
	WERKS           string  `json:"werks" validate:"required"`
	VBELN           string  `json:"vbeln" validate:"required"`
	POSNR           string  `json:"posnr" validate:"required"`
	PICK_LOCATION   string  `json:"pick_location" validate:"required"`
	JENIS_KIRIM     string  `json:"jenis_kirim" validate:"required"`
	DO_INDUK        string  `json:"do_induk"`
	BLDAT           string  `json:"bldat" validate:"required"`
	REQ_DLV_DATE    string  `json:"req_dlv_date"`
	REQ_DLV_TIME1   string  `json:"req_dlv_time1"`
	REQ_DLV_TIME2   string  `json:"req_dlv_time2"`
	NOTES_FR_SALES  string  `json:"notes_fr_sales" validate:"required"`
	NOTES_FR_TRANSP string  `json:"notes_fr_transp"`
	SOLDTO          string  `json:"soldto" validate:"required"`
	SHIPTO          string  `json:"shipto" validate:"required"`
	REF_PO          string  `json:"ref_po" validate:"required"`
	MATNR           string  `json:"matnr" validate:"required"`
	MAT_DESC        string  `json:"mat_desc" validate:"required"`
	LFIMG           float32 `json:"lfimg" validate:"required"`
	VRKME           string  `json:"vrkme" validate:"required"`
	BRGEW           float32 `json:"brgew" validate:"required"`
	VOLUM           float32 `json:"volum" validate:"required"`
	LFIMG_RECEIVED  float32 `json:"lfimg_received"`
	VRKME_RECEIVED  string  `json:"vrkme_received"`
	REASON_DIFF     string  `json:"reason_diff"`
	RECIPIENT_NAME  string  `json:"recipient_name"`
	RECEIVE_DATE    string  `json:"receive_date"`
	RECEIVE_TIME    string  `json:"receive_time"`
	VENDOR          string  `json:"vendor" validate:"required"`
	VEHICLE_TYPE    string  `json:"vehicle_type" validate:"required"`
	DRIVER_NAME     string  `json:"driver_name" validate:"required"`
	VEHICLE_NO      string  `json:"vehicle_no" validate:"required"`
	LPB_NUMBER      string  `json:"lpb_number"`
	TASK_STATUS     string  `json:"task_status" validate:"required"`
	ERDAT           string  `json:"erdat" validate:"required"`
	ERZET           string  `json:"erzet" validate:"required"`
	ERNAM           string  `json:"ernam" validate:"required"`
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

type ZtsTravisIdDetail struct {
	MANDT  string `json:"mandt" form:"mandt"`
	TASKID string `json:"taskid" form:"taskid"`
	VKORG  string `json:"vkorg" form:"vkorg"`
	WERKS  string `json:"werks" form:"werks"`
	VBELN  string `json:"vbeln" form:"vbeln"`
	POSNR  string `json:"posnr" form:"posnr"`
}

type ZtsTravisUpdateRequest struct {
	PICK_LOCATION   string  `json:"pick_location"`
	JENIS_KIRIM     string  `json:"jenis_kirim"`
	DO_INDUK        string  `json:"do_induk"`
	BLDAT           string  `json:"bldat"`
	REQ_DLV_DATE    string  `json:"req_dlv_date"`
	REQ_DLV_TIME1   string  `json:"req_dlv_time1"`
	REQ_DLV_TIME2   string  `json:"req_dlv_time2"`
	NOTES_FR_SALES  string  `json:"notes_fr_sales" form:"notes_fr_sales"`
	NOTES_FR_TRANSP string  `json:"notes_fr_transp" form:"notes_fr_transp"`
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
