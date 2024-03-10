package rest

type BasePageWrapper struct {
	Items PageItemsWrapper `json:"items"`
}

type PageResponseWrapper struct {
	Data BasePageWrapper `json:"data"`
	Meta MetaResponse    `json:"meta"`
}
type Pageable struct {
	PageNumber int  `json:"pageNumber"`
	PageSize   int  `json:"pageSize"`
	Sort       Sort `json:"sort"`
	Offset     int  `json:"offset"`
	Unpaged    bool `json:"unpaged"`
	Paged      bool `json:"paged"`
}

type Sort struct {
	Empty    bool `json:"empty"`
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
}

type PageItemsWrapper struct {
	Content          interface{} `json:"content"`
	Pageable         Pageable    `json:"pageable"`
	Last             bool        `json:"last"`
	TotalElements    int         `json:"totalElements"`
	TotalPages       int         `json:"totalPages"`
	Size             int         `json:"size"`
	Number           int         `json:"number"`
	Sort             Sort        `json:"sort"`
	First            bool        `json:"first"`
	NumberOfElements int         `json:"numberOfElements"`
	Empty            bool        `json:"empty"`
}
