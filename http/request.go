package net

import (
	"github.com/macococo/go-webbase/utils"
	"net/http"
	"strconv"
)

func GetRequestParam(r *http.Request, name string, def string) string {
	param := r.FormValue(name)
	if param == "" {
		return def
	}
	return param
}

func GetRequestParamInt(r *http.Request, name string, def int) int {
	str := r.FormValue(name)
	if str == "" {
		return def
	}

	param, err := strconv.Atoi(str)
	if utils.HandleError(err) != nil {
		return def
	}
	return param
}
