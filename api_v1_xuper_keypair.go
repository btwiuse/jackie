package main

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
)

type ApiV1XuperKeypairNew struct{}

type ApiV1XuperKeypairNewResponse struct {
	Address string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey  string `json:"privkey" default:"{...}"`
	Privkey string `json:"pubkey" default:"{...}"`
}

func (t *ApiV1XuperKeypairNew) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperKeypairNewResponse
		prog     = "keypair.new"
		cmd      = exec.Command(prog)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: output})
		return
	}

	if err := yaml.Unmarshal(out, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
