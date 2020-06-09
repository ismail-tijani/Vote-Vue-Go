package controllers

import (
	"net/http"

	"github.com/ismail-tijani/Vote-Vue-Go/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
