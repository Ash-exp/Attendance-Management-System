package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/Ash-exp/Attendance-Management-System/models"
)

func Load(db *gorm.DB) {

	err := db.Debug().AutoMigrate(&models.User{}, &models.Teacher{}, &models.Attendance{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Attendance{}).AddForeignKey("student_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}