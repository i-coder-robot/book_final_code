package res

type AccountResp struct {
	AccountName string `json:"accountName"`
}

type ListResponse struct {
	TotalCount  uint64             `json:"totalCount"`
	AccountList []*AccountResp `json:"accountList"`
}
