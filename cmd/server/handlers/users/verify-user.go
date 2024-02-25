package handlers

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/utils/api"
)

// @Summary Verifies an user
// @Tags users
// @Produce  json
// @Success 200 {object} api.Response
// @Failure 500 {object} api.Response
// @Failure 400 {object} api.Response
// @Failure 404 {object} api.Response
// @Param token query string false "token"
// @Router /api/v1/users/verify [get]
func VerifyUser(res http.ResponseWriter, req *http.Request) {
	token := api.Param(req, "token")
	if token == "" {
		api.JSON(res, api.Response{
			Message: "No token provided",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	user, err := repositories.FindUserById(token)
	if err != nil {
		api.JSON(res, api.Response{
			Message: "Invalid token",
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	user.Enabled = true
	user, err = repositories.UpdateUserByID(user)
	if err != nil {
		api.JSON(res, api.Response{
			Message: err.Error(),
			Data:    nil,
		}, http.StatusBadRequest)
		return
	}

	api.JSON(res, api.Response{
		Message: "User verified successfully, you can now login",
		Data:    user,
	}, http.StatusOK)
}