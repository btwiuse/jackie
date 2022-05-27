package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
)

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
	var (
		resp       ApiV1XuperAccountNewResponse
		keypair, _ = jsonMarshalString(t)
		prog       = "account.new"
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

	if err := yaml.Unmarshal(out, &resp); err != nil {
		log.Println("yaml.Unmarshal", err, output)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
