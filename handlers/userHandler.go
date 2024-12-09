package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/AnkitDhawale/TodoListApp/helpers"
	"github.com/AnkitDhawale/TodoListApp/middlewares"
	"github.com/AnkitDhawale/TodoListApp/services"
	"net/http"
)

type UserHandler struct {
	Service     services.UserService
	AuthService services.AuthService
}

// Login authenticates users via email and password.
// Login godoc
// @Summary User login
// @Description Authenticates a user and returns a JWT token if credentials are correct
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.User true "User credentials"
// @Success 200 {object} helpers.Response "JWT token"
// @Failure 400 {object} helpers.Response "Invalid request payload"
// @Failure 401 {object} helpers.Response "Wrong email or password"
// @Router /todoapp/login [post]
func (uh UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u dto.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("invalid request payload"))
		return
	}

	if u.Email == "" || u.Password == "" {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("email & password is required"))
		return
	}

	tokenResponse, err := uh.AuthService.Login(&u)
	if err != nil {
		helpers.WriteResponse(w, http.StatusUnauthorized, nil, errors.New("wrong email or password"))
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, tokenResponse, nil)
		return
	}
}

// GetAllUsers fetches all users from db.
// GetAllUsers godoc
// @Summary Get all users
// @Description Fetches all users from db
// @Tags users
// @Produce json
// @Success 200 {array} helpers.Response "List of all users"
// @Failure 500 {object} helpers.Response "Internal server error"
// @Router /todoapp/users/all [get]
func (uh UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	go HandleTiemout(w, r)

	users, err := uh.Service.GetAllUsers()
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, err)
	} else {
		helpers.WriteResponse(w, http.StatusOK, users, nil)
	}
}

// SignUp godoc
// @Summary User signup
// @Description Creates new user account
// @Tags users
// @Accept json
// @produce json
// @Param userdata body dto.User true "User data for signup"
// @Success 201 {object} helpers.Response "Success"
// @Failure 400 {object} helpers.Response "email & password should not be empty"
// @Failure 500 {object} helpers.Response "Unexpected error"
// @Router /todoapp/signup [post]
func (uh UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("invalid request payload"))
		return
	}

	if user.Email == "" || user.Password == "" {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("email & password should not be empty"))
		return
	}
	userId, err := uh.Service.CreatNewUser(user)
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	} else {
		helpers.WriteResponse(w, http.StatusCreated, fmt.Sprintf("New user created with id:%s", userId), nil)
		return
	}
}

// UpdateUser godoc
// @Summary User update
// @Description Updates an user account
// @Tags users
// @Accept json
// @produce json
// @Param userdata body dto.User true "User data for update"
// @Success 200 {object} helpers.Response "Success update"
// @Failure 400 {object} helpers.Response "email & password should not be empty/at least 1 field required to update profile"
// @Failure 500 {object} helpers.Response "Internal server error"
// @Router /todoapp/user-update [patch]
func (uh UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	usrId := r.Context().Value(middlewares.UserIDKey)
	if usrId == nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, errors.New("userId not found in context"))
		return
	}

	usrEmail := r.Context().Value(middlewares.UserEmailKey)
	if usrEmail == nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, errors.New("userEmail not found in context"))
		return
	}

	userIdFromToken, ok := usrId.(string)
	if !ok {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, errors.New("typecast error: invalid type of userId from context"))
		return
	}

	var user dto.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	}

	if user.Email == "" && user.Password == "" {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("at least 1 field required to update profile"))
		return
	}

	err = uh.Service.UpdateUser(userIdFromToken, user)
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, "user updated successfully...", nil)
		return
	}
}

func HandleTiemout(w http.ResponseWriter, r *http.Request) {
	select {
	case <-r.Context().Done():
		// Handle timeout or cancellation
		helpers.WriteResponse(w, http.StatusRequestTimeout, nil, errors.New("request timed out"))
		return
	default:
	}
}
