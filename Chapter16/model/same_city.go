package model

type SameCity struct {
	Id 	string `json:"id"`
	Title string `json:"title"`
	PhotoAddr string `json:"photoAddr"`
	Distance string `json:"distance"`
	AvatarUrl string `json:"avatarUrl"`
	Location string `json:"location"`
	Random int32 `json:"random"`
	Desc string `json:"desc"`
}

type SameCityItems struct {
	Items1 []SameCity `json:"items1"`
	Items2 []SameCity `json:"items2"`
}
