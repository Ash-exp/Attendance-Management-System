package controllers

import "github.com/Ash-exp/Attendance-Management-System/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

	//Teachers routes
	s.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(s.CreateTeacher)).Methods("POST")
	s.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(s.GetTeachers)).Methods("GET")
	s.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(s.GetTeacher)).Methods("GET")
	s.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareJSON(s.UpdateTeacher))).Methods("PUT")
	s.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(s.DeleteTeacher)).Methods("DELETE")
	
	//Student Attendance routes
	s.Router.HandleFunc("/attendance", middlewares.SetMiddlewareJSON(s.CreateAttendance)).Methods("POST")
	s.Router.HandleFunc("/attendance", middlewares.SetMiddlewareJSON(s.GetAttendances)).Methods("GET")
	s.Router.HandleFunc("/attendance/{id}", middlewares.SetMiddlewareJSON(s.GetAttendance)).Methods("GET")
	s.Router.HandleFunc("/attendance/{id}", middlewares.SetMiddlewareJSON(s.UpdateAttendance)).Methods("PUT")
	
	//Teacher Attendance routes
	s.Router.HandleFunc("/teacher-attendance", middlewares.SetMiddlewareJSON(s.CreateTeacherAttendance)).Methods("POST")
	s.Router.HandleFunc("/teacher-attendance/{id}", middlewares.SetMiddlewareJSON(s.GetTeacherAttendance)).Methods("GET")
	s.Router.HandleFunc("/teacher-attendance/{id}", middlewares.SetMiddlewareJSON(s.UpdateTeacherAttendance)).Methods("PUT")

}