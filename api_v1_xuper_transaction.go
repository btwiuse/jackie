package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type ApiV1XuperQueryTx struct {
	Transaction string `uri:"id" binding:"required" json:"id" default:"45abf461aa69fe0135b97a791d5b402da952bcf8f31a63babcb01b7242d130cd"`
}

type ApiV1XuperQueryTxResponse struct{}

func (t *ApiV1XuperQueryTx) Handler(c *gin.Context) {
	var (
		resp     string
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

	resp = output

	c.String(http.StatusOK, resp)
}
