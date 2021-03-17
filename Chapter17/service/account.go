package service

import (
	"errors"
	"github.com/i-coder-robot/book_final_code/Chapter17/model"
	"github.com/i-coder-robot/book_final_code/Chapter17/repository"
)

type AccountService struct {
	Repo *repository.AccountModelRepo
}

func (ac *AccountService) ListAccountByName(username string, offset, limit int) ([]*model.AccountInfo, uint64, error) {
	infos := make([]*model.AccountInfo, 0)
	accounts, count, err := ac.Repo.ListAccount(offset, limit)
	if err != nil {
		return nil, count, err
	}

	for _, item := range accounts {
		info := &model.AccountInfo{
			AccountId:   item.AccountId,
			AccountName: item.AccountName,
			Password:    item.Password,
		}
		infos = append(infos, info)
	}

	return infos, count, nil
}

func (ac *AccountService) ListAccount(offset, limit int) ([]*model.AccountInfo, uint64, error) {
	infos := make([]*model.AccountInfo, 0)
	accounts, count, err := ac.Repo.ListAccount(offset, limit)
	if err != nil {
		return nil, count, err
	}

	for _, item := range accounts {
		info := &model.AccountInfo{
			AccountId:   item.AccountId,
			AccountName: item.AccountName,
			Password:    item.Password,
		}
		infos = append(infos, info)
	}

	return infos, count, nil
}

func (ac *AccountService) CreateAccount(account model.Account) error {
	return ac.Repo.CreateAccount(account)
}

func (ac *AccountService) DeleteAccount(id string) error {
	return ac.Repo.DeleteAccount(id)
}

func (ac *AccountService) UpdateAccount(account model.Account) error {
	accountInfo, err := ac.Repo.GetAccountInfo(account.AccountId)
	if err != nil {
		return err
	}
	if accountInfo.AccountId == "" {
		return errors.New("用户不存在")
	}
	return ac.Repo.UpdateAccount(account)
}

func (ac *AccountService) GetAccount(accountName string) (model.Account, error) {
	return ac.Repo.GetAccountByName(accountName)
}
