package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/long2ice/swagin/router"
	"github.com/long2ice/swagin/security"
)

var (
	query = router.New(
		&TestQuery{},
		router.Summary("Test query"),
		router.Description("Test query model"),
		router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model:       TestQuery{},
				Description: "response model description",
			},
		}),
	)
	queryList = router.New(
		&TestQueryList{},
		router.Summary("Test query list"),
		router.Description("Test query list model"),
		router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: []TestQueryList{},
			},
		}),
	)
	health = router.New(
		&Health{},
		router.Summary("Health check"),
		router.Description("Health check. Should return `200 OK`"),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: Health{OK: true},
			},
		}),
	)
	noModel = router.New(
		&TestNoModel{},
		router.Summary("Test no model"),
		router.Description("Test no model"),
	)
	queryPath = router.New(
		&TestQueryPath{},
		router.Summary("Test query path"),
		router.Description("Test query path model"),
	)
	formEncode = router.New(
		&TestForm{},
		router.Summary("Test form"),
		router.ContentType(binding.MIMEPOSTForm),
	)
	body = router.New(
		&TestForm{},
		router.Summary("Test json body"),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: TestForm{},
			},
		}),
	)
	file = router.New(
		&TestFile{},
		router.Summary("Test file upload"),
		router.ContentType(binding.MIMEMultipartPOSTForm),
	)
	apiV1XuperHello = router.New(
		&ApiV1XuperHello{},
		router.Summary("Xuper Hello"),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: ApiV1XuperHello{},
			},
		}),
	)
	apiV1XuperKeypairNew = router.New(
		&ApiV1XuperKeypairNew{},
		router.Summary("New Keypair"),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: ApiV1XuperKeypairNewResponse{},
			},
		}),
	)
)