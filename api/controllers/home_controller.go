package controllers

import (
	"net/http"

	"github.com/krittawatcode/go-rest-testing/api/responses"
)

// Home ...
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
