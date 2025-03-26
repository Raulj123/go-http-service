package handler

import (
	"net/http"
	"strconv"

	"github.com/Raulj123/go-service/employee"
	"github.com/Raulj123/go-service/utils"
	"github.com/go-chi/chi/v5"
)


type Handler struct {
	http.Handler
	provider employee.EmpProvider
}

func NewHandler(prov employee.EmpProvider) *Handler {
	r := chi.NewRouter()
	h := &Handler{
		Handler: r,
		provider: prov,
	}
	r.Get("/{id}", h.getEmployee)
	r.Post("/", h.postEmployee)
	return h
}

func (h *Handler) getEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	emp, err := h.provider.Employee(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	utils.EncodeJson(w, http.StatusOK, emp)
}

func (h *Handler) postEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method",http.StatusBadRequest)
		return
	}
	d,err := utils.DecodeJson[employee.Employee](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.provider.Store(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}