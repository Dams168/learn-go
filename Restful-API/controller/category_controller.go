package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(write http.ResponseWriter, read *http.Request, params httprouter.Params)
	Update(write http.ResponseWriter, read *http.Request, params httprouter.Params)
	Delete(write http.ResponseWriter, read *http.Request, params httprouter.Params)
	FindById(write http.ResponseWriter, read *http.Request, params httprouter.Params)
	FindAll(write http.ResponseWriter, read *http.Request, params httprouter.Params)
}
