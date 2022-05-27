package main

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type ApiV1XuperAdminBalance struct{}

func (t *ApiV1XuperAdminBalance) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperBalanceResponse
		prog     = "balance.of"
		cmd      = exec.Command(prog)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	resp = ApiV1XuperBalanceResponse{Balance: output}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperFaucet struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
	Amount  string `query:"amount" json:"amount" default:"1000"`
}

type ApiV1XuperFaucetResponse struct {
	Transaction string `json:"tx" default:"45abf461aa69fe0135b97a791d5b402da952bcf8f31a63babcb01b7242d130cd"`
}

func (t *ApiV1XuperFaucet) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperFaucetResponse
		prog     = "transfer.to"
		args     = []string{t.Address, t.Amount}
		cmd      = exec.Command(prog, args...)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	resp = ApiV1XuperFaucetResponse{Transaction: output}

	c.JSON(http.StatusOK, resp)
}

type ApiV1XuperBalance struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
}

type ApiV1XuperBalanceResponse struct {
	Balance string `json:"balance" default:"0"`
}

func (t *ApiV1XuperBalance) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperBalanceResponse
		prog     = "balance.of"
		args     = []string{t.Address}
		cmd      = exec.Command(prog, args...)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	resp = ApiV1XuperBalanceResponse{Balance: output}

	c.JSON(http.StatusOK, resp)
}
