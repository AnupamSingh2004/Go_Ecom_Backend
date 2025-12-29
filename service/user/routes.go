package user

import (
	"net/http"

	"github.com/AnupamSingh2004/Go_Ecom_Backend/types"
	"github.com/AnupamSingh2004/Go_Ecom_Backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	if err := utils.ParsePayload(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}
