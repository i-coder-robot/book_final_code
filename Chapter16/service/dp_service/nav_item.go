package dp_service

import (
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
	"book-code/Chapter13/13-4/res"
)

type ListNavItemService struct {
	Repo *repository.ListNavItemRepo
}

func model2res(nav model.Nav) res.Nav{
	navRes:= res.Nav{
		NavId: nav.NavId,
		Src:   nav.Src,
		Title: nav.Title,
	}
	return navRes
}

func (i *ListNavItemService) ListNavItems(level int) []res.Nav {
	result := i.Repo.ListNavItems(level)
	var newList []res.Nav
	for _,item := range result{
		r:=model2res(item)
		newList= append(newList, r)
	}

	return newList
}
