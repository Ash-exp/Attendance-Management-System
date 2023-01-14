package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Ash-exp/Attendance-Management-System/api/models"
	"github.com/Ash-exp/Attendance-Management-System/api/responses"
	"github.com/Ash-exp/Attendance-Management-System/api/utils/formaterror"
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
	attendanceGotten, err := attendance.AttendanceExists(server.DB, uint32(attendance.StudentId))
	
	if (err != nil) {
		if (err.Error() == "Record Not Found"){
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
			return
		}
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if !attendanceGotten.State {
		updatedAttendance, err := attendance.UpdatePunchInAttendance(server.DB, attendanceGotten.ID)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
		responses.JSON(w, http.StatusOK, updatedAttendance)
		return
	} else {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Already Punched In! Try Punching Out"))
		return
	}
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
	attendance := models.Attendance{}
	updatedAttendance, err := attendance.UpdatePunchOutAttendance(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Try Punching In First !! Thank You."))
		return
	}
	responses.JSON(w, http.StatusOK, updatedAttendance)
}