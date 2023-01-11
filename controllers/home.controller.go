package controllers

import (
	"net/http"

	"github.com/Ash-exp/Attendance-Management-System/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Attendance Management APIs")

}