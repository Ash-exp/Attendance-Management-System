package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Ash-exp/Attendance-Management-System/api/models"
	"github.com/Ash-exp/Attendance-Management-System/api/responses"
	"github.com/Ash-exp/Attendance-Management-System/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateTeacher(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	teacher := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher.Prepare()
	err = teacher.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacherCreated, err := teacher.SaveTeacher(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, teacherCreated.ID))
	responses.JSON(w, http.StatusCreated, teacherCreated)
}

func (server *Server) GetTeachers(w http.ResponseWriter, r *http.Request) {

	teacher := models.Teacher{}

	teachers, err := teacher.FindAllTeachers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, teachers)
}

func (server *Server) GetTeacher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	teacher := models.Teacher{}
	teacherGotten, err := teacher.FindTeacherByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, teacherGotten)
}

func (server *Server) UpdateTeacher(w http.ResponseWriter, r *http.Request) {

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
	teacher := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher.Prepare()
	err = teacher.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedTeacher, err := teacher.UpdateATeacher(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedTeacher)
}

func (server *Server) DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	teacher := models.Teacher{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = teacher.DeleteATeacher(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}