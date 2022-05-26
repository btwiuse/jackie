package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		return
	}

	resp = ApiV1XuperBalanceResponse{Balance: strings.TrimSpace(string(out))}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperQueryTx struct {
	Transaction string `uri:"id" binding:"required" json:"id" default:"45abf461aa69fe0135b97a791d5b402da952bcf8f31a63babcb01b7242d130cd"`
}

type ApiV1XuperQueryTxResponse struct{}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (t *ApiV1XuperQueryTx) Handler(c *gin.Context) {
	var (
		prog     = "query.tx"
		args     = []string{t.Transaction}
		cmd      = exec.Command(prog, args...)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
	)

	if err != nil {
		log.Println(prog, args, err, output)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	c.String(http.StatusOK, output)
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
		return
	}

	resp = ApiV1XuperFaucetResponse{Transaction: strings.TrimSpace(string(out))}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperAddrconvX2e struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
}

type ApiV1XuperAddrconvResponse struct {
	Type   string `json:"type" default:""`
	Result string `json:"result" default:""`
}

func (t *ApiV1XuperAddrconvX2e) Handler(c *gin.Context) {
	var resp ApiV1XuperAddrconvResponse

	cmd := exec.Command("addrconv", "x2e", t.Address)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperAddrconvE2x struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
}

func (t *ApiV1XuperAddrconvE2x) Handler(c *gin.Context) {
	var resp ApiV1XuperAddrconvResponse

	cmd := exec.Command("addrconv", "e2x", t.Address)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

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
		return
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

type ApiV1XuperContractQuery struct {
	Template string `json:"template" query:"template" default:"Empty"`
	Contract string `json:"contract" default:"deadbeaf"`
	// ABI      string `json:"abi" default:"{}"`
	Method string `json:"method" default:"method"`
	Args   string `json:"args" default:"{}"`
}

type ApiV1XuperContractQueryResponse struct{}

func (t *ApiV1XuperContractQuery) Handler(c *gin.Context) {
	var resp string

	/*
		fmt.Println("contract.query")
		fmt.Println(t.Template)
		fmt.Println(t.Contract)
		fmt.Println(t.Method)
		fmt.Println(t.Args)
	*/
	cmd := exec.Command("contract.query", t.Template, t.Contract, t.Method, t.Args)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp = strings.TrimSpace(string(out))

	c.String(http.StatusOK, resp)
}

type ApiV1XuperContractInvoke struct {
	Address  string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey   string `json:"privkey" default:"{...}"`
	Privkey  string `json:"pubkey" default:"{...}"`
	Template string `json:"template" query:"template" default:"Empty"`
	Contract string `json:"contract" default:"deadbeaf"`
	// ABI      string `json:"abi" default:"{}"`
	Method string `json:"method" default:"method"`
	Args   string `json:"args" default:"{}"`
}

type ApiV1XuperContractInvokeResponse struct {
	Transaction string `json:"tx" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperContractInvoke) Handler(c *gin.Context) {
	var resp ApiV1XuperContractInvokeResponse

	keypair, _ := jsonMarshalString(t)

	/*
		fmt.Println("contract.invoke")
		fmt.Println(keypair)
		fmt.Println(t.Template)
		fmt.Println(t.Contract)
		fmt.Println(t.Method)
		fmt.Println(t.Args)
	*/
	cmd := exec.Command("contract.invoke", keypair, t.Template, t.Contract, t.Method, t.Args)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperContractDeploy struct {
	Address  string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey   string `json:"privkey" default:"{...}"`
	Privkey  string `json:"pubkey" default:"{...}"`
	Args     string `json:"args" default:"{}"`
	Account  string `json:"account" default:"account"`
	Template string `json:"template" query:"template" default:"Empty"`
}

type ApiV1XuperContractDeployResponse struct {
	Contract string `json:"contract" default:"deadbeaf"`
	// ABI         string `json:"abi" default:"{}"`
	Template    string `json:"template" default:"default"`
	Transaction string `json:"tx" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperContractDeploy) Handler(c *gin.Context) {
	var resp ApiV1XuperContractDeployResponse

	keypair, _ := jsonMarshalString(t)

	/*
		fmt.Println("contract.deploy", keypair, t.Account, t.Template, t.Args)
		fmt.Println(keypair)
		fmt.Println(t.Account)
		fmt.Println(t.Template)
		fmt.Println(t.Args)
	*/
	cmd := exec.Command("contract.deploy", keypair, t.Account, t.Template, t.Args)
	out, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
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

func jsonMarshalString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (t *ApiV1XuperAccountNew) Handler(c *gin.Context) {
	var resp ApiV1XuperAccountNewResponse
	var (
		prog       = "account.new"
		keypair, _ = jsonMarshalString(t)
		args       = []string{keypair}
		cmd        = exec.Command(prog, args...)
		out, err   = cmd.CombinedOutput()
		output     = strings.TrimSpace(string(out))
	)

	if err != nil {
		log.Println(prog, args, err, output)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	if err = yaml.Unmarshal(out, &resp); err != nil {
		log.Println("yaml.Unmarshal", err, out)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
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
		return
	}

	err = yaml.Unmarshal(out, &resp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
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
