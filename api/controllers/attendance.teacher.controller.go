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

func (server *Server) CreateTeacherAttendance(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	attendance := models.TeacherAttendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher := models.Teacher{}
	_, err = teacher.FindTeacherByID(server.DB, uint32(attendance.TeacherId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	attendance.Prepare()
	attendanceGotten, err := attendance.TeacherAttendanceExists(server.DB, uint32(attendance.TeacherId))
	
	if (err != nil) {
		if (err.Error() == "Record Not Found"){
			err = attendance.Validate("")
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			attendanceCreated, err := attendance.SaveTeacherAttendance(server.DB)
		
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
		updatedAttendance, err := attendance.UpdatePunchInTeacherAttendance(server.DB, attendanceGotten.ID)
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


func (server *Server) GetTeacherAttendance(w http.ResponseWriter, r *http.Request) {

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
	attendance := models.TeacherAttendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	attendanceGotten, err := attendance.FindTeacherAttendanceByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendanceGotten)
}

func (server *Server) UpdateTeacherAttendance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	attendance := models.TeacherAttendance{}
	updatedAttendance, err := attendance.UpdatePunchOutTeacherAttendance(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Try Punching In First !! Thank You."))
		return
	}
	responses.JSON(w, http.StatusOK, updatedAttendance)
}