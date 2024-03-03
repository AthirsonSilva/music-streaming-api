package handlers

import (
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/repositories"
)

// CreateAlbum @Summary Creates an album
// @Tags albums
// @Accept  application/json
// @Produce  application/json
// @Param album body models.AlbumRequest true "Album request"
// @Success 200 {object} api.Response
// @Failure 500 {object} internal.Exception
// @Failure 400 {object} internal.Exception
// @Failure 429 {object} internal.Exception
// @Param Authorization header string true "Authorization"
// @Router /api/v1/albums [post]
func CreateAlbum(res http.ResponseWriter, req *http.Request) {
	var request models.AlbumRequest
	var response api.Response

	if err := api.ReadBody(req, &request); err != nil {
		api.Error(res, req, "Malformed request", err, http.StatusBadRequest)
		return
	}

	album := request.ToModel()
	album, err := repositories.CreateAlbum(album)
	if err != nil {
		api.Error(res, req, "Error while creating album", err, http.StatusInternalServerError)
		return
	}

	response = api.Response{
		Message: "Album created",
		Data:    album,
	}
	api.JSON(res, response, http.StatusCreated)
}
