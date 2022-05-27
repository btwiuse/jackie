package main

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
)

type ApiV1XuperAddrconvX2e struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
}

type ApiV1XuperAddrconvResponse struct {
	Type   string `json:"type" default:""`
	Result string `json:"result" default:""`
}

func (t *ApiV1XuperAddrconvX2e) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperAddrconvResponse
		prog     = "addrconv"
		args     = []string{"x2e", t.Address}
		cmd      = exec.Command(prog, args...)
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

type ApiV1XuperAddrconvE2x struct {
	Address string `uri:"address" binding:"required" json:"address" default:"0"`
}

func (t *ApiV1XuperAddrconvE2x) Handler(c *gin.Context) {
	var (
		resp     ApiV1XuperAddrconvResponse
		prog     = "addrconv"
		args     = []string{"e2x", t.Address}
		cmd      = exec.Command(prog, args...)
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
