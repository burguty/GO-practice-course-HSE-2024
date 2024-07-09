package handler

import (
	"github.com/labstack/echo/v4"
	"homework_2/models"
	"net/http"
	"sync"
)

type Handler struct {
	data map[string]*models.Account
	mu   *sync.RWMutex
}

func New() *Handler {
	return &Handler{
		data: make(map[string]*models.Account),
		mu:   &sync.RWMutex{},
	}
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request models.CreateAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.data[request.Name]; ok {
		return c.String(http.StatusForbidden, "Account already exists")
	}

	h.data[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: 0,
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	request := models.DeleteAccountRequest{
		Name: c.QueryParam("name"),
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.data[request.Name]; !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	delete(h.data, request.Name)

	return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateAmount(c echo.Context) error {
	var request models.UpdateAmountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.data[request.Name]; !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	h.data[request.Name].Amount += request.Amount

	return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateName(c echo.Context) error {
	var request models.UpdateNameRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}
	if len(request.NewName) == 0 {
		return c.String(http.StatusBadRequest, "Empty new name")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.data[request.Name]; !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	if _, ok := h.data[request.NewName]; ok {
		return c.String(http.StatusConflict, "Account with new name already exists")
	}

	account := h.data[request.Name]
	account.Name = request.NewName
	delete(h.data, request.Name)
	h.data[request.NewName] = account

	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetAccount(c echo.Context) error {
	request := models.GetAccountRequest{
		Name: c.QueryParam("name"),
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	if _, ok := h.data[request.Name]; !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	return c.JSON(http.StatusOK, h.data[request.Name])
}
