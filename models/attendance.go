package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Attendance struct {
	ID        	uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	StudentId 	uint32 	  	`gorm:"not null" json:"studentId"`
	Class  		uint32    	`gorm:"not null;" json:"class"`
	State		string		`gorm:"not null;default:'punchIn'" json:"state"`
	Day  		int    		`gorm:"not null;" json:"day"`
	Month  		time.Month  `gorm:"not null;" json:"month"`
	Year  		int    		`gorm:"not null;" json:"year"`
	PunchInAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"punchedIn_at"`
	PunchOutAt 	time.Time 	`json:"punchedOut_at"`
}


func (u *Attendance) Prepare() {
	u.ID = 0
	u.State = html.EscapeString(strings.TrimSpace(u.State))
	u.PunchInAt = time.Now()
	year, month, day := time.Now().Date()
	u.Day = day
	u.Month = month
	u.Year = year

}

func (u *Attendance) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		return nil

	default:
		if u.StudentId == 0 {
			return errors.New("Required StudentId")
		}
		if u.Class == 0 {
			return errors.New("Required Class")
		}
		return nil
	}
}

func (u *Attendance) SaveAttendance(db *gorm.DB) (*Attendance, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Attendance{}, err
	}
	return u, nil
}

func (u *Attendance) FindClassAttendance(db *gorm.DB) (*[]Attendance, error) {
	var err error
	attendance := []Attendance{}
	err = db.Debug().Model(&Attendance{}).Where("class = ? AND year = ? AND month = ?", u.Class, u.Year, u.Month).Limit(100).Find(&attendance).Error
	if err != nil {
		return &[]Attendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Attendance{}, errors.New("Record Not Found")
	}
	return &attendance, err
}

func (u *Attendance) AttendanceExists(db *gorm.DB, studentId uint32) (bool, error) {
	var err error
	err = db.Debug().Model(&Attendance{}).Where("student_id = ? AND year = ? AND month = ? AND day = ?", studentId, u.Year, u.Month, u.Day).Take(&u).Error
	if err != nil {
		return false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return false, errors.New("Record Not Found")
	}
	return true, err
}

func (u *Attendance) FindAttendanceByID(db *gorm.DB, studentId uint32) (*[]Attendance, error) {
	var err error
	attendance := []Attendance{}
	err = db.Debug().Model(&Attendance{}).Where("student_id = ? AND year = ? AND month = ?", studentId, u.Year, u.Month).Limit(100).Find(&attendance).Error
	if err != nil {
		return &[]Attendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Attendance{}, errors.New("Record Not Found")
	}
	return &attendance, err
}

func (u *Attendance) UpdateAttendance(db *gorm.DB, id uint32) (*Attendance, error) {

	db = db.Debug().Model(&Attendance{}).Where("student_id = ?", id).Take(&Attendance{}).UpdateColumns(
		map[string]interface{}{
			"state":     "punchOut",
			"punch_out_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Attendance{}, db.Error
	}
	// This is the display the updated attendance
	err := db.Debug().Model(&Attendance{}).Where("student_id = ?", id).Take(&u).Error
	if err != nil {
		return &Attendance{}, err
	}
	return u, nil
}