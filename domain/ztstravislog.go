package domain

type ZtsTravisLog struct {
	MANDT        string `json:"mandt"`
	TASKID       string `json:"taskid"`
	OBJECT_CLASS string `json:"object_class"`
	OBJECT_VALUE string `json:"object_value"`
	CHANGENR     string `json:"changenr"`
	OLD_VALUE    string `json:"old_value"`
	NEW_VALUE    string `json:"new_value"`
	PUBLISH      string `json:"publish"`
	USERNAME     string `json:"username"`
	UDATE        string `json:"udate"`
	UTIME        string `json:"utime"`
	ERDAT        string `json:"erdat"`
	ERZET        string `json:"erzet"`
	ERNAM        string `json:"ernam"`
	AEDAT        string `json:"aedat"`
	AEZET        string `json:"aezet"`
	AENAM        string `json:"aenam"`
	SEL          string `json:"sel"`
}

type ZtsTravisLogRequest struct {
	MANDT        string `json:"mandt"`
	TASKID       string `json:"taskid"`
	OBJECT_CLASS string `json:"object_class"`
	OBJECT_VALUE string `json:"object_value"`
	CHANGENR     string `json:"changenr"`
	OLD_VALUE    string `json:"old_value"`
	NEW_VALUE    string `json:"new_value"`
	PUBLISH      string `json:"publish"`
	USERNAME     string `json:"username"`
	UDATE        string `json:"udate"`
	UTIME        string `json:"utime"`
	ERDAT        string `json:"erdat"`
	ERZET        string `json:"erzet"`
	ERNAM        string `json:"ernam"`
	AEDAT        string `json:"aedat"`
	AEZET        string `json:"aezet"`
	AENAM        string `json:"aenam"`
	SEL          string `json:"sel"`
}

type ZtsTravisLogResponse struct {
	MANDT        string `json:"mandt"`
	TASKID       string `json:"taskid"`
	OBJECT_CLASS string `json:"object_class"`
	OBJECT_VALUE string `json:"object_value"`
	CHANGENR     string `json:"changenr"`
	OLD_VALUE    string `json:"old_value"`
	NEW_VALUE    string `json:"new_value"`
	PUBLISH      string `json:"publish"`
	USERNAME     string `json:"username"`
	UDATE        string `json:"udate"`
	UTIME        string `json:"utime"`
	ERDAT        string `json:"erdat"`
	ERZET        string `json:"erzet"`
	ERNAM        string `json:"ernam"`
	AEDAT        string `json:"aedat"`
	AEZET        string `json:"aezet"`
	AENAM        string `json:"aenam"`
	SEL          string `json:"sel"`
}
