package rest

type PageResponseWrapper struct {
	Data interface{}  `json:"data"`
	Meta MetaResponse `json:"meta"`
}
