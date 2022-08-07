package handlers

import "net/http"

type URLHandler interface {
	Get(http.ResponseWriter, *http.Request)
	GetMostUsed(http.ResponseWriter, *http.Request)
	SaveNoCustomAlias(http.ResponseWriter, *http.Request)
	SaveWithCustomAlias(http.ResponseWriter, *http.Request)
}
