package main

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/swagin/security"
)

type TestQuery struct {
	Name     string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	Token    string `header:"token" binding:"required" json:"token" default:"test"`
	Optional string `query:"optional" json:"optional"`
}

func (t *TestQuery) Handler(c *gin.Context) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, t)
}

type TestQueryList struct {
	Name  string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	Token string `header:"token" binding:"required" json:"token" default:"test"`
}

func (t *TestQueryList) Handler(c *gin.Context) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, []TestQueryList{*t})
}

type TestQueryPath struct {
	Name  string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	ID    int    `uri:"id" binding:"required" json:"id" description:"id of model" default:"1"`
	Token string `header:"token" binding:"required" json:"token" default:"test"`
}

func (t *TestQueryPath) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t)
}

type TestForm struct {
	ID   int    `query:"id" binding:"required" json:"id" description:"id of model" default:"1"`
	Name string `form:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	List []int  `form:"list" binding:"required" json:"list" description:"list of model"`
}

func (t *TestForm) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t)
}

type ApiV1XuperAdminBalance struct{}

func (t *ApiV1XuperAdminBalance) Handler(c *gin.Context) {
	var resp ApiV1XuperBalanceResponse

	cmd := exec.Command("balance.of")
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	resp = ApiV1XuperBalanceResponse{Balance: strings.TrimSpace(string(out))}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperQueryTx struct {
	Transaction string `uri:"id" binding:"required" json:"id" default:"45abf461aa69fe0135b97a791d5b402da952bcf8f31a63babcb01b7242d130cd"`
}

type ApiV1XuperQueryTxResponse struct{}

func (t *ApiV1XuperQueryTx) Handler(c *gin.Context) {
	var resp string

	cmd := exec.Command("query.tx", t.Transaction)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	resp = strings.TrimSpace(string(out))

	c.String(http.StatusOK, resp)
}

type ApiV1XuperFaucet struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
	Amount  string `query:"amount" json:"amount" default:"1000"`
}

type ApiV1XuperFaucetResponse struct {
	Transaction string `json:"tx" default:"45abf461aa69fe0135b97a791d5b402da952bcf8f31a63babcb01b7242d130cd"`
}

func (t *ApiV1XuperFaucet) Handler(c *gin.Context) {
	var resp ApiV1XuperFaucetResponse

	cmd := exec.Command("transfer.to", t.Address, t.Amount)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	resp = ApiV1XuperFaucetResponse{Transaction: strings.TrimSpace(string(out))}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperBalance struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
}

type ApiV1XuperBalanceResponse struct {
	Balance string `json:"balance" default:"0"`
}

func (t *ApiV1XuperBalance) Handler(c *gin.Context) {
	var resp ApiV1XuperBalanceResponse

	cmd := exec.Command("balance.of", t.Address)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	resp = ApiV1XuperBalanceResponse{Balance: strings.TrimSpace(string(out))}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperHello struct {
	World string `json:"hello" default:"world"`
}

func (t *ApiV1XuperHello) Handler(c *gin.Context) {
	t.World = "world"
	c.JSON(http.StatusOK, t)
}

type ApiV1XuperAccountNew struct {
	Address string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey  string `json:"privkey" default:"{...}"`
	Privkey string `json:"pubkey" default:"{...}"`
}

type ApiV1XuperAccountNewResponse struct {
	Account     string `json:"account" default:"XC1111111111111111@xuper"`
	Transaction string `json:"tx" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperAccountNew) Handler(c *gin.Context) {
	var resp ApiV1XuperAccountNewResponse

	jb, err := json.Marshal(t)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	keypair := string(jb)

	cmd := exec.Command("account.new", keypair)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperKeypairNew struct{}

type ApiV1XuperKeypairNewResponse struct {
	Address string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey  string `json:"privkey" default:"{...}"`
	Privkey string `json:"pubkey" default:"{...}"`
}

func (t *ApiV1XuperKeypairNew) Handler(c *gin.Context) {
	var resp ApiV1XuperKeypairNewResponse

	cmd := exec.Command("keypair.new")
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, resp)
}

type Health struct {
	OK bool `json:"ok" default:"true"`
}

func (t *Health) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, Health{OK: true})
}

type TestNoModel struct {
}

func (t *TestNoModel) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

type TestFile struct {
	File *multipart.FileHeader `form:"file" binding:"required" description:"file upload"`
}

func (t *TestFile) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"file": t.File.Filename})
}
