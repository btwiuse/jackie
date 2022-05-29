package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
)

type ApiV1XuperCollectionNew struct {
	Address string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey  string `json:"privkey" default:"{...}"`
	Privkey string `json:"pubkey" default:"{...}"`
	Args    string `json:"args" default:"{}"`
	Account string `json:"account" default:"account"`
}

type ApiV1XuperCollectionNewResponse struct {
	Collection  string `json:"collection" default:"deadbeaf"`
	Template    string `json:"template" default:"default"`
	Transaction string `json:"tx" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperCollectionNew) Handler(c *gin.Context) {
	var (
		resp       ApiV1XuperCollectionNewResponse
		keypair, _ = jsonMarshalString(t)
		prog       = "collection.new"
		args       = []string{keypair, t.Account, t.Args}
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

type ApiV1XuperCollectionMint struct {
	Address    string `json:"address" default:"iHXRXwahx4yf6CwYyDQGvJYj4o39Jdgrs"`
	Pubkey     string `json:"privkey" default:"{...}"`
	Privkey    string `json:"pubkey" default:"{...}"`
	Collection string `json:"collection" default:"account"`
	Args       string `json:"args" default:"{}"`
}

type ApiV1XuperCollectionMintResponse struct {
	Collection  string `json:"collection" default:"deadbeaf"`
	Template    string `json:"template" default:"default"`
	Transaction string `json:"tx" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperCollectionMint) Handler(c *gin.Context) {
	var (
		resp       ApiV1XuperCollectionMintResponse
		keypair, _ = jsonMarshalString(t)
		prog       = "collection.mint"
		args       = []string{keypair, t.Collection, t.Args}
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

type ApiV1XuperCollectionInfo struct {
	Collection string `uri:"collection" json:"collection" default:"account"`
}

type ApiV1XuperCollectionInfoResponse struct {
	Name        string `json:"name" default:"deadbeaf"`
	Description string `json:"description" default:"default"`
	Owner       string `json:"owner" default:"cb057a9dce7f8a1d928c46ceb84e8765fab43a5ecf85bf061c59bbbc2e717932"`
}

func (t *ApiV1XuperCollectionInfo) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperCollectionInfoResponse
		prog     = "collection.info"
		args     = []string{t.Collection}
		cmd      = exec.Command(prog, args...)
		out, err = cmd.CombinedOutput()
		output   = strings.TrimSpace(string(out))
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
