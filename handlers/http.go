package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eduardomassami/hime.me/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type handler struct {
	urlService domain.Service
}

func NewHandler(urlService domain.Service) URLHandler {
	return &handler{urlService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alias := chi.URLParam(r, "alias")
	u, err := h.urlService.Get(alias)

	if err != nil {

		if strings.Contains(err.Error(), "URL NOT FOUND") {
			r := domain.ErrorResponse{}
			r.ERR_Code = "002"
			r.Description = "SHORTENED URL NOT FOUND"
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(r)
			return
		}

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&u)
}

func (h *handler) GetMostUsed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u, err := h.urlService.GetMostUsed()

	if err != nil {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(&u)
}

func (h *handler) SaveNoCustomAlias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := r.URL.Query().Get("url")
	uuid := uuid.NewString()
	fmt.Println(uuid)
	us := &domain.URL{}
	us.URL = url
	us.Alias = uuid

	err := h.urlService.SaveNoCustomAlias(us)

	if err != nil {

		if strings.Contains(err.Error(), "Duplicate entry") {
			r := domain.ErrorResponse{}
			r.Alias = uuid
			r.ERR_Code = "001"
			r.Description = "CUSTOM ALIAS ALREADY EXISTS"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(r)
			return
		}

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// json.NewEncoder(w).Encode(&u)
}

func (h *handler) SaveWithCustomAlias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := r.URL.Query().Get("url")
	alias := r.URL.Query().Get("CUSTOM_ALIAS")
	us := domain.URL{}
	us.URL = url
	us.Alias = alias

	err := h.urlService.SaveWithCustomAlias(&us)

	if err != nil {

		if strings.Contains(err.Error(), "Duplicate entry") {
			r := domain.ErrorResponse{}
			r.Alias = alias
			r.ERR_Code = "001"
			r.Description = "CUSTOM ALIAS ALREADY EXISTS"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(r)
			return
		}

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// json.NewEncoder(w).Encode(&u)
}
