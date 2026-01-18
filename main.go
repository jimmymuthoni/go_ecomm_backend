package main

import (
	"fmt"
	"net/http"

	"github.com/anthdm/weavebox"
)


func handleAPIError(ctx *weavebox.Context, err error){
	fmt.Println("API Error !!!:", err)
	ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}