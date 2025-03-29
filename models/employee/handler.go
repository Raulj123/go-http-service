package employee

import (
	"net/http"
	"strconv"

	"github.com/Raulj123/go-service/httpjson"
	"github.com/go-chi/chi/v5"
)


type Handler struct {
	http.Handler
	provider Provider
}

func NewHandler(prov Provider) *Handler {
	r := chi.NewRouter()
	h := &Handler{
		Handler: r,
		provider: prov,
	}
	r.Get("/", h.GetEmployees)
	r.Post("/", h.postEmployee)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.getEmployee)
	}) 
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
	if err := httpjson.Encode(w, http.StatusOK,emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) postEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method",http.StatusBadRequest)
		return
	}
	d,err := httpjson.Decode[Employee](r)
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

func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request){
	emps,err := h.provider.Employees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := httpjson.Encode(w, http.StatusOK,emps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}