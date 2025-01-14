package handlers

import (
	"net/http"
	"time"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/serialization"
	"github.com/ortizdavid/go-rest-concepts/entities"
	"github.com/ortizdavid/go-rest-concepts/models"
	"gorm.io/gorm"
)


type UserHandler struct {
	userModel models.UserModel
}


func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		userModel: *models.NewUserModel(db),
	}
}


func (h *UserHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /api/users", h.getAllUsers)
	router.HandleFunc("GET /api/users/{id}", h.getUser)
	router.HandleFunc("POST /api/users", h.createUser)
	router.HandleFunc("PUT /api/users/{id}", h.updateUser)
	router.HandleFunc("DELETE /api/users/{id}", h.deleteUser)
}


func (h *UserHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	currentPage, limit := GetCurrentPageAndLimit(r)
	users, err := h.userModel.FindAllDataLimit(currentPage, limit)
	count := h.userModel.Count()
	
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		httputils.WriteJsonError(w, "unot found", http.StatusNotFound)
		return
	}
	httputils.WriteJsonPaginated(w, r, users, count, currentPage, limit)
}


func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userId := conversion.StringToInt(id)

	user, err := h.userModel.FindById(userId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	httputils.WriteJson(w, http.StatusOK, user)
}


func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if err := serialization.DecodeJson(r.Body, &user); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, _ := h.userModel.ExistsRecord("user_name", user.UserName)
	if exists {
		httputils.WriteJsonError(w, "Username'"+user.UserName+ "' exists", http.StatusConflict)
		return
	}
	user.Active = "Yes"
	user.Password = encryption.HashPassword(user.Password)
	user.UniqueId = encryption.GenerateUUID()
	user.Token = encryption.GenerateRandomToken()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := h.userModel.Create(user)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	httputils.WriteJson(w, http.StatusCreated, user)
}


func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser entities.User
	id := r.PathValue("id")
	userId := conversion.StringToInt(id)

	user, err := h.userModel.FindById(userId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := serialization.DecodeJson(r.Body, updatedUser); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.UserName = updatedUser.UserName
	user.Password = encryption.HashPassword(updatedUser.Password)
	user.Active = updatedUser.Active
	user.UpdatedAt = time.Now()

	_, err = h.userModel.Update(user)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJson(w, http.StatusOK, updatedUser)
}


func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userId := conversion.StringToInt(id)

	user, err := h.userModel.FindById(userId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return 
	}
	_, err = h.userModel.Delete(user)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJson(w, http.StatusNoContent, nil)
}