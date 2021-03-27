package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"github.com/i-coder-robot/book_final_code/Chapter16/MyLog"
	"github.com/i-coder-robot/book_final_code/Chapter16/model/my_token"
	"github.com/i-coder-robot/book_final_code/Chapter16/myerr"
	"github.com/i-coder-robot/book_final_code/Chapter16/res"
	"github.com/i-coder-robot/book_final_code/Chapter16/service/wx_service"
	"github.com/i-coder-robot/book_final_code/Chapter16/token"
	"github.com/i-coder-robot/book_final_code/Chapter16/utils"
	"github.com/i-coder-robot/book_final_code/Chapter17/handler/param"
	"github.com/i-coder-robot/book_final_code/Chapter17/model"
	"github.com/i-coder-robot/book_final_code/Chapter17/service"
)

type AccountHandler struct {
	Srv *service.AccountService
}

// 新建一个account.
func (h *AccountHandler) AccountCreate(c *gin.Context) {
	var r param.AccountRequest
	if err := c.Bind(&r); err != nil {
		res.SendResponse(c, myerr.ErrBind, nil)
		return
	}

	if err := utils.CheckParam(r.AccountName, r.Password); err.Err != nil {
		res.SendResponse(c, err.Err, nil)
		return
	}

	accountName := r.AccountName
	MyLog.Log.Infof("用户名: %s", accountName)

	desc := c.Query("desc")
	MyLog.Log.Infof("desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	MyLog.Log.Infof("Header Content-Type: %s", contentType)

	// 把明文密码加密
	md5Pwd, err := utils.Encrypt(r.Password)
	if err != nil {
		res.SendResponse(c, myerr.ErrEncrypt, nil)
		return
	}
	id, err := uuid.GenerateUUID()
	if err != nil {
		res.SendResponse(c, myerr.InternalServerError, nil)
		return
	}
	// 添加用户到数据库
	a := model.Account{
		AccountId:   id,
		AccountName: r.AccountName,
		Password:    md5Pwd,
	}
	if err := h.Srv.CreateAccount(a); err != nil {
		res.SendResponse(c, myerr.ErrDatabase, nil)
		return
	}

	rsp := res.AccountResp{
		AccountName: r.AccountName,
	}

	// 现实新建的 Acccount 信息.
	res.SendResponse(c, nil, rsp)
}

func (h *AccountHandler) ListAccount(c *gin.Context) {
	var r param.AccountListRequest
	if err := c.Bind(&r); err != nil {
		res.SendResponse(c, myerr.ErrBind, nil)
		return
	}
	if r.Offset < 0 {
		r.Offset = 0
	}
	if r.Limit < 1 {
		r.Limit = utils.Limit
	}

	list, count, err := h.Srv.ListAccount(r.Offset, r.Limit)
	if err != nil {
		res.SendResponse(c, err, nil)
		return
	}
	var resp []*res.AccountResp
	for _, item := range list {
		r := res.AccountResp{AccountName: item.AccountName}
		resp = append(resp, &r)
	}

	res.SendResponse(c, nil, res.ListResponse{
		TotalCount:  count,
		AccountList: resp,
	})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	accountName := c.Param("account_name")
	// 从数据库中选择Account.
	account, err := h.Srv.GetAccount(accountName)
	if err != nil {
		res.SendResponse(c, myerr.ErrAccountNotFound, nil)
		return
	}
	r := res.AccountResp{AccountName: account.AccountName}
	res.SendResponse(c, nil, r)
}

func (h *AccountHandler) Update(c *gin.Context) {

	// 绑定 account.
	var m model.Account
	if err := c.Bind(&m); err != nil {
		res.SendResponse(c, myerr.ErrBind, nil)
		return
	}

	// 密码加密处理.
	md5Pwd, err := utils.Encrypt(m.Password)
	if err != nil {
		res.SendResponse(c, myerr.ErrEncrypt, nil)
		return
	}
	m.Password = md5Pwd
	// 保存更新.
	if err := h.Srv.UpdateAccount(m); err != nil {
		res.SendResponse(c, myerr.ErrDatabase, nil)
		return
	}
	res.SendResponse(c, nil, nil)
}

func (h *AccountHandler) Delete(c *gin.Context) {
	accountId := c.Param("id")
	if err := h.Srv.DeleteAccount(accountId); err != nil {
		res.SendResponse(c, myerr.ErrDatabase, nil)
		return
	}
	res.SendResponse(c, nil, nil)
}

func (h *AccountHandler) Login(c *gin.Context) {
	var m model.Account
	if err := c.Bind(&m); err != nil {
		res.SendResponse(c, myerr.ErrBind, nil)
		return
	}
	account, err := h.Srv.Repo.GetAccountByName(m.AccountName)
	if err != nil {
		res.SendResponse(c, myerr.ErrAccountNotFound, nil)
		return
	}

	if err := utils.Compare(account.Password, m.Password); err != nil {
		res.SendResponse(c, myerr.ErrPassword, nil)
		return
	}

	sign, err := token.Sign(c, token.Context{ID: account.AccountId, Username: account.AccountName}, "")
	if err != nil {
		res.SendResponse(c, myerr.ErrToken, nil)
		return
	}

	res.SendResponse(c, nil, my_token.Token{Token: sign})
}

// 微信小程序登录
func (h *AccountHandler) WXLogin(c *gin.Context) {
	code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := wx_service.WXLogin(code)
	if err != nil {
		res.SendResponse(c, nil, nil)
		return
	}
	// 保存登录态
	session := sessions.Default(c)
	session.Set("openid", wxLoginResp.OpenId)
	session.Set("sessionKey", wxLoginResp.SessionKey)

	// 这里可以用openid和sessionkey的串接,或者使用你自己的规则进行拼接,然后进行MD5之后作为该用户的自定义登录态， 要保证mySession唯一,
	mySession := utils.GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库或缓存中, 可以用mySession去索引openid 和sessionkey
	res.SendResponse(c, nil, mySession)
}
