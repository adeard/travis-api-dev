package domain

type ZtsTravisAtta struct {
	MANDT           string `json:"mandt" gorm:"column:MANDT"`
	TASKID          string `json:"taskid" gorm:"column:TASKID"`
	FILE_TYPE       string `json:"file_type" gorm:"column:FILE_TYPE"`
	POSNR           string `json:"posnr" gorm:"column:POSNR"`
	URL_TRAVIS      string `json:"url_travis" gorm:"column:URL_TRAVIS"`
	FILE_NAME       string `json:"file_name" gorm:"column:FILE_NAME"`
	FILE_CREATEDATE string `json:"file_createdate" gorm:"column:FILE_CREATEDATE"`
	FILE_CREATETIME string `json:"file_createtime" gorm:"column:FILE_CREATETIME"`
	ERDAT           string `json:"erdat" gorm:"column:ERDAT"`
	ERZET           string `json:"erzet" gorm:"column:ERZET"`
	ERNAM           string `json:"ernam" gorm:"column:ERNAM"`
	AEDAT           string `json:"aedat" gorm:"column:AEDAT"`
	AEZET           string `json:"aezet" gorm:"column:AEZET"`
	AENAM           string `json:"aenam" gorm:"column:AENAM"`
	SEL             string `json:"sel" gorm:"column:SEL"`
}

type ZtsTravisAttaRequest struct {
	MANDT           string `json:"mandt"`
	TASKID          string `json:"taskid"`
	FILE_TYPE       string `json:"file_type"`
	POSNR           string `json:"posnr"`
	URL_TRAVIS      string `json:"url_travis"`
	FILE_NAME       string `json:"file_name"`
	FILE_CREATEDATE string `json:"file_createdate"`
	FILE_CREATETIME string `json:"file_createtime"`
	ERDAT           string `json:"erdat"`
	ERZET           string `json:"erzet"`
	ERNAM           string `json:"ernam"`
	AEDAT           string `json:"aedat"`
	AEZET           string `json:"aezet"`
	AENAM           string `json:"aenam"`
	SEL             string `json:"sel"`
}

type ZtsTravisAttaResponse struct {
	MANDT           string `json:"mandt"`
	TASKID          string `json:"taskid"`
	FILE_TYPE       string `json:"file_type"`
	POSNR           string `json:"posnr"`
	URL_TRAVIS      string `json:"url_travis"`
	FILE_NAME       string `json:"file_name"`
	FILE_CREATEDATE string `json:"file_createdate"`
	FILE_CREATETIME string `json:"file_createtime"`
	ERDAT           string `json:"erdat"`
	ERZET           string `json:"erzet"`
	ERNAM           string `json:"ernam"`
	AEDAT           string `json:"aedat"`
	AEZET           string `json:"aezet"`
	AENAM           string `json:"aenam"`
	SEL             string `json:"sel"`
}
