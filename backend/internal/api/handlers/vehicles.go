package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"workshop-platform/backend/internal/auth"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type VehiclesHandler struct {
	repo repositories.VehicleRepository
}

func NewVehiclesHandler(repo repositories.VehicleRepository) *VehiclesHandler {
	return &VehiclesHandler{repo: repo}
}

func (h *VehiclesHandler) List(w http.ResponseWriter, r *http.Request) {
	claims := auth.FromContext(r.Context())
	if claims == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	// TODO: parse limit/offset from query
	list, _, err := h.repo.ListByWorkshop(r.Context(), claims.WorkshopID, 50, 0)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (h *VehiclesHandler) Create(w http.ResponseWriter, r *http.Request) {
	claims := auth.FromContext(r.Context())
	if claims == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	var v domain.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
		return
	}
	v.WorkshopID = claims.WorkshopID
	if err := h.repo.Create(r.Context(), &v); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}

func (h *VehiclesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, `{"error":"id required"}`, http.StatusBadRequest)
		return
	}
	v, err := h.repo.GetByID(r.Context(), id)
	if err != nil || v == nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
