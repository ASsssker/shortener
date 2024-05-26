package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"shortener/internal/models"
)

// PostUrl создает короткий адрес
func (a *Application) PostUrl(w http.ResponseWriter, r *http.Request) {
	reqData := &models.RequestDataModels{}

	if err := json.NewDecoder(r.Body).Decode(reqData); err != nil {
		a.serverError(w, err)
		return
	}

	key := generateString(8)
	if err := a.DB.Insert(key, reqData.Url); err != nil {
		a.serverError(w, err)
		return
	}

	respData := &models.UrlResponseModels{
		ResultUrl: a.ServerAddr + a.RootUrl + key,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(respData); err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(buf.Bytes())
}

// PostBatchUrl создает короткий адрес для батча адресов
func (a *Application) PostBatchUrl(w http.ResponseWriter, r *http.Request) {
	reqData := make([]models.BatchRequestModel, 10)

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		a.serverError(w, err)
		return
	}
	respData := make([]models.BatchResponseModel, len(reqData))

	for idx := range reqData {
		key := generateString(8)
		if err := a.DB.Insert(key, reqData[idx].OriginalUrl); err != nil {
			a.serverError(w, err)
			return
		}
		respData[idx].CorrelationId = reqData[idx].CorrelationId
		respData[idx].ShortUrl = a.ServerAddr + a.RootUrl + key
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(respData); err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(buf.Bytes())
}

// GetUrl перенаправляет по адресу
func (a *Application) GetUrl(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	url, err := a.DB.Get(id)
	if err != nil {
		if err.Error() == "url not found" {
			a.notFound(w)
			return
		}
		a.serverError(w, err)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *Application) PingDB(w http.ResponseWriter, r *http.Request) {
	if a.DB == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PONG"))
}
