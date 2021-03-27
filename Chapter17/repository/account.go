package repository

import (
	"fmt"
	"github.com/i-coder-robot/book_final_code/Chapter16/utils"
	"github.com/i-coder-robot/book_final_code/Chapter17/model"
)

type AccountModelRepo struct {
	DB model.DataBase
}

func (m AccountModelRepo) TableName() string {
	return "account"
}

// 新建一个 account.
func (m *AccountModelRepo) CreateAccount(account model.Account) error {
	return m.DB.MyDB.Create(&account).Error
}

// 通过一个 Id. 删除 Account
func (m *AccountModelRepo) DeleteAccount(id string) error {
	err := m.DB.MyDB.Where("account_id = ?", id).Delete(&model.Account{}).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新 account.
func (m *AccountModelRepo) UpdateAccount(account model.Account) error {
	err := m.DB.MyDB.Model(model.Account{}).Where("account_id=?", account.AccountId).Updates(map[string]interface{}{
		"account_name":     account.AccountName,
		"password": account.Password,
	}).Error
	return err
}

// 获取指定 Account.
func (m *AccountModelRepo) GetAccountByName(name string) (model.Account, error) {
	var account model.Account
	err := m.DB.MyDB.Where("account_name = ?", name).First(&account).Error
	if err != nil {
		return account, err
	}
	return account, nil
}

// Account列表
func (m *AccountModelRepo) ListAccount(offset, limit int) ([]*model.Account, uint64, error) {

	accounts := make([]*model.Account, 0)
	var count uint64

	if err := m.DB.MyDB.Model(&model.Account{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := m.DB.MyDB.Model(&model.Account{}).Limit(limit).Offset(offset).Order("id desc").Find(&accounts).Error
	if err != nil {
		return nil, count, err
	}
	return accounts, count, nil
}

func (m *AccountModelRepo) ListAccountByName(account_name string, offset, limit int) ([]*model.Account, uint64, error) {
	if limit == 0 {
		limit = utils.Limit
	}

	accounts := make([]*model.Account, 0)
	var count uint64

	where := fmt.Sprintf("account_name like '%%%s%%'", account_name)
	if err := model.DB.MyDB.Model(&model.Account{}).Where(where).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := model.DB.MyDB.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&accounts).Error; err != nil {
		return accounts, count, err
	}
	return accounts, count, nil
}

func (m *AccountModelRepo) GetAccountInfo(accountId string) (model.Account, error) {
	var account model.Account
	if err := m.DB.MyDB.Where("account_id=?", accountId).First(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}
