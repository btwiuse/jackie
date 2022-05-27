package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
)

type ApiV1XuperContractQuery struct {
	Template string `json:"template" query:"template" default:"Empty"`
	Contract string `json:"contract" default:"deadbeaf"`
	// ABI      string `json:"abi" default:"{}"`
	Method string `json:"method" default:"method"`
	Args   string `json:"args" default:"{}"`
}

type ApiV1XuperContractQueryResponse struct{}

func (t *ApiV1XuperContractQuery) Handler(c *gin.Context) {
	var (
		resp     string
		prog     = "contract.query"
		args     = []string{t.Template, t.Contract, t.Method, t.Args}
		cmd      = exec.Command(prog, args...)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	resp = output

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
	var (
		resp       ApiV1XuperContractInvokeResponse
		keypair, _ = jsonMarshalString(t)
		prog       = "contract.invoke"
		args       = []string{keypair, t.Template, t.Contract, t.Method, t.Args}
		cmd        = exec.Command(prog, args...)
		out, err   = cmd.CombinedOutput()
		output     = strings.TrimSpace(string(out))
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	if err = yaml.Unmarshal(out, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
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
	Contract    string `json:"contract" default:"deadbeaf"`
	Template    string `json:"template" default:"default"`
	Transaction string `json:"tx" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperContractDeploy) Handler(c *gin.Context) {
	var (
		resp       ApiV1XuperContractDeployResponse
		keypair, _ = jsonMarshalString(t)
		prog       = "contract.deploy"
		args       = []string{keypair, t.Account, t.Template, t.Args}
		cmd        = exec.Command(prog, args...)
		out, err   = cmd.CombinedOutput()
		output     = strings.TrimSpace(string(out))
	)

	if err != nil {
		log.Println(prog, args, err, output)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	if err := yaml.Unmarshal(out, &resp); err != nil {
		log.Println("yaml.Unmarshal", err, output)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
