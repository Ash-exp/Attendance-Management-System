package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Ash-exp/Attendance-Management-System/models"
	"github.com/Ash-exp/Attendance-Management-System/responses"
	"github.com/Ash-exp/Attendance-Management-System/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateAttendance(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	attendance := models.Attendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(attendance.StudentId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	attendance.Class = userGotten.Class
	attendance.Prepare()
	attendanceGotten, _ := attendance.AttendanceExists(server.DB, uint32(attendance.StudentId))
	if attendanceGotten {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Already Punched In! Try Punching Out"))
		return
	}
	err = attendance.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendanceCreated, err := attendance.SaveAttendance(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, attendanceCreated.ID))
	responses.JSON(w, http.StatusCreated, attendanceCreated)
}

func (server *Server) GetAttendances(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendance := models.Attendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendances, err := attendance.FindClassAttendance(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendances)
}

func (server *Server) GetAttendance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendance := models.Attendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendanceGotten, err := attendance.FindAttendanceByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendanceGotten)
}

func (server *Server) UpdateAttendance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendance := models.Attendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendance.Prepare()
	err = attendance.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedAttendance, err := attendance.UpdateAttendance(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedAttendance)
}