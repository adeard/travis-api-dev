package domain

type ZtsTravisAtta struct {
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
