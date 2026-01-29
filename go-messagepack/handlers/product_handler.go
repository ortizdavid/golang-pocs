package handlers

import (
    "net/http"
    "strconv"

    "github.com/ortizdavid/golang-pocs/go-messagepack/helpers"
    "github.com/ortizdavid/golang-pocs/go-messagepack/models"
    "github.com/ortizdavid/golang-pocs/go-messagepack/services"
)

type ProductHandler struct {
    service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

// Helper interno para padronizar respostas de erro binárias

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
    var p models.ProductModel
    if err := helpers.Unmarshal(r.Body, &p); err != nil {
        helpers.SendError(w, "Invalid MessagePack payload", http.StatusBadRequest)
        return
    }

    result, err := h.service.Create(&p)
    if err != nil {
        helpers.SendError(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/x-msgpack")
    w.WriteHeader(http.StatusCreated)
    helpers.Marshal(w, result)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
    // Captura o {id} definido no main.go
    id, _ := strconv.Atoi(r.PathValue("id"))

    var p models.ProductModel
    if err := helpers.Unmarshal(r.Body, &p); err != nil {
        helpers.SendError(w, "Invalid MessagePack payload", http.StatusBadRequest)
        return
    }

    result, err := h.service.Update(id, &p)
    if err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "product not found" {
            status = http.StatusNotFound
        }
        helpers.SendError(w, err.Error(), status)
        return
    }

    w.Header().Set("Content-Type", "application/x-msgpack")
    helpers.Marshal(w, result)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.PathValue("id"))

    result, err := h.service.Delete(id)
    if err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "product not found" {
            status = http.StatusNotFound
        }
        helpers.SendError(w, err.Error(), status)
        return
    }

    w.Header().Set("Content-Type", "application/x-msgpack")
    helpers.Marshal(w, result)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    products, err := h.service.GetAll()
    if err != nil {
        helpers.SendError(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/x-msgpack")
    helpers.Marshal(w, products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.PathValue("id"))

    product, err := h.service.GetByID(id)
    if err != nil {
        helpers.SendError(w, "product not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/x-msgpack")
    helpers.Marshal(w, product)
}