package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"testWB-app/internal/storage"

	"github.com/go-chi/chi/v5"
)

type OrdersHandler struct {
	storage storage.OrdersGetter
}

func New(orders storage.OrdersGetter) OrdersHandler {
	return OrdersHandler{storage: orders}
}

func (m OrdersHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	orders, err := m.storage.Get(id)

	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		http.Error(w, "there is no such order in cache", http.StatusBadRequest)
		return
	}

	main := filepath.Join("..", "..", "internal", "static", "dynamicOrderPage.html")
	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "orders", orders)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
