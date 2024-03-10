package rest

type BaseDataWrapper struct {
	Items interface{} `json:"items"`
}

type DataResponse struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}
