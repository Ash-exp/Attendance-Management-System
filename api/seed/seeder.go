package seed

import (
	"log"

	"github.com/Ash-exp/Attendance-Management-System/api/models"
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().AutoMigrate(&models.User{}, &models.Teacher{}, &models.Attendance{}, &models.Log{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Attendance{}).AddForeignKey("student_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Log{}).AddForeignKey("attendance_id", "attendances(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.TeacherAttendance{}).AddForeignKey("teacher_id", "teachers(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.TeacherLog{}).AddForeignKey("attendance_id", "teacher_attendances(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}