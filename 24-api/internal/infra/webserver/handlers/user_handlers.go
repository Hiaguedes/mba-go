package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/hiaguedes/go-api/internal/dto"

	"github.com/hiaguedes/go-api/internal/entity"
	"github.com/hiaguedes/go-api/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

// Create user godoc
// @Summary      Get a JWT User
// @Description  A handler to Get a JWT User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.GetJwtInput  true  "user credentials"
// @Success      200 {object} dto.GetJwtInput
// @Failure      404
// @Failure      500 {object} Error
// @Router       /users/generate-token [post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwt := r.Context().Value(JwtKey).(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value(JwtExpiresInKey).(int)

	mapToEncode := map[string]any{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	}
	_, tokenString, _ := jwt.Encode(mapToEncode)

	accessToken := dto.GetJwtOutput{AcessToken: tokenString}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}

// Create user godoc
// @Summary      Create user
// @Description  A handler for creating a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500  		{object}   Error
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request){
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	err = h.UserDB.Create(u)
		if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
