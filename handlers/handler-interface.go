package handlers

import "net/http"

type URLHandler interface {
	Get(http.ResponseWriter, *http.Request)
	GetMostUsed(http.ResponseWriter, *http.Request)
	Save(http.ResponseWriter, *http.Request)
}
