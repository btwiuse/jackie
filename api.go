package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/swagin/security"
	"mime/multipart"
	"net/http"
	"time"
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

type ApiV1XuperHello struct {
	World string `json:"hello" default:"world"`
}

func (t *ApiV1XuperHello) Handler(c *gin.Context) {
	t.World = "world"
	c.JSON(http.StatusOK, t)
}

type ApiV1XuperKeypairNew struct{}

type ApiV1XuperKeypairNewResponse struct {
	Address string `json:"address" default:"addraddraddr"`
	Pubkey  string `json:"privkey" default:"privprivpriv"`
	Privkey string `json:"pubkey" default:"pubpubpub"`
}

func (t *ApiV1XuperKeypairNew) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, &ApiV1XuperKeypairNewResponse{
		Address: "cafebabe",
		Pubkey:  "deadbeef",
		Privkey: time.Now().Local().String(),
	})
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
