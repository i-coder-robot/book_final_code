package param

type AccountRequest struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

type AccountListRequest struct {
	AccountName string `json:"accountName"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}
