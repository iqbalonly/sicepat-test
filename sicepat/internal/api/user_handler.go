package api

import (
	"net/http"
	"sicepat/internal/constant"
	"sicepat/internal/dto"
	"sicepat/internal/ierr"
	"sicepat/internal/storage"
	"strconv"

	"github.com/gorilla/mux"
)

func NewUserHandler(storage storage.StorageDAO) *userHandler {
	return &userHandler{
		storage: storage,
	}
}

type userHandler struct {
	storage storage.StorageDAO
}

// AddRoutes adds the routers for this API to the provided router (or subrouter)
func (h *userHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.getHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", h.upsertHandler).Methods(http.MethodPost)
	router.HandleFunc("/users", h.deleteHandler).Methods(http.MethodDelete)
}

func (h *userHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.storage.GetUsers(r.Context())
	if err != nil {
		renderErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	renderResponse(w, users, http.StatusOK, constant.Success)
}

func (h *userHandler) upsertHandler(w http.ResponseWriter, r *http.Request) {
	req := new(dto.UserRequest)
	err := parseAPIRequest(r, req)
	if err != nil {
		renderErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.storage.CreateOrUpdateUser(r.Context(), req)
	if err != nil {
		renderErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	renderResponse(w, nil, http.StatusOK, constant.Success)
}

func (h *userHandler) deleteHandler(w http.ResponseWriter, r *http.Request) {
	userIDparam := r.URL.Query().Get("userId")

	userID, err := strconv.ParseInt(userIDparam, 10, 64)
	if err != nil {
		renderErrorResponse(w, ierr.FailedParameterType, http.StatusBadRequest)
		return
	}

	err = h.storage.DeleteUser(r.Context(), int(userID))
	if err != nil {
		renderErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	renderResponse(w, nil, http.StatusOK, constant.Success)
}
