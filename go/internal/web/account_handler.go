package web

import (
	"encoding/json"
	"net/http"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
)

type WebAccountHandler struct {
	create_accountUseCase create_account.CreateAccountUseCase
}

func NewWebAccountHandler(c create_account.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		create_accountUseCase: c,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto create_account.CreateAccountInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.create_accountUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
