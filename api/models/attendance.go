package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)
type Log struct {
	AttendanceID uint32		`gorm:"not null" json:"attendance_id"`
	PunchInAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP;not null" json:"punchedIn_at"`
	PunchOutAt 	time.Time 	`json:"punchedOut_at"`
}
type Attendance struct {
	ID        	uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	StudentId 	uint32 	  	`gorm:"not null" json:"studentId"`
	Class  		uint32    	`gorm:"not null;" json:"class"`
	State		bool		`gorm:"not null;default: true" json:"state"`
	Day  		int    		`gorm:"not null;" json:"day"`
	Month  		time.Month  `gorm:"not null;" json:"month"`
	Year  		int    		`gorm:"not null;" json:"year"`
	Logs		[]Log 		`gorm:"foreignkey:AttendanceID" json:"logs"`	
}


func (u *Attendance) Prepare() {
	u.ID = 0
	u.Logs = []Log{{PunchInAt : time.Now(), AttendanceID: u.ID}}
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
	err = db.Debug().Model(&Attendance{}).Preload("Logs").Where("class = ? AND year = ? AND month = ?", u.Class, u.Year, u.Month).Limit(100).Find(&attendance).Error
	if err != nil {
		return &[]Attendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Attendance{}, errors.New("Record Not Found")
	}
	return &attendance, err
}

func (u *Attendance) AttendanceExists(db *gorm.DB, studentId uint32) (*Attendance, error) {
	var err error
	err = db.Debug().Model(&Attendance{}).Where("student_id = ? AND year = ? AND month = ? AND day = ?", studentId, u.Year, u.Month, u.Day).Preload("Logs").Take(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Attendance{}, errors.New("Record Not Found")
	}
	if err != nil {
		return &Attendance{}, err
	}
	return u, err
}

func (u *Attendance) FindAttendanceByID(db *gorm.DB, studentId uint32) (*[]Attendance, error) {
	var err error
	attendance := []Attendance{}
	err = db.Debug().Model(&Attendance{}).Where("student_id = ? AND year = ? AND month = ?", studentId, u.Year, u.Month).Preload("Logs").Limit(100).Find(&attendance).Error
	if err != nil {
		return &[]Attendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Attendance{}, errors.New("Record Not Found")
	}
	return &attendance, err
}

func (u *Attendance) UpdatePunchInAttendance(db *gorm.DB, id uint32) (*Attendance, error) {

	err := db.Debug().Model(&Attendance{}).Where("id = ?", id).Take(&u).UpdateColumns(
		map[string]interface{}{
			"state":     true,
		},
	).Error
	if err != nil {
		return &Attendance{}, err
	}
	err = db.Debug().Create(&Log{
		AttendanceID: u.ID,
		PunchInAt: time.Now(),
	}).Error
	if err != nil {
		return &Attendance{}, err
	}
	
	// This is the display the updated attendance
	err = db.Debug().Model(&Attendance{}).Preload("Logs").Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &Attendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Attendance{}, errors.New("Record Not Found")
	}
	return u, nil
}

func (u *Attendance) UpdatePunchOutAttendance(db *gorm.DB, id uint32) (*Attendance, error) {

	err := db.Debug().Model(&Attendance{}).Where("student_id = ?", id).Take(&u).Error
	if err != nil {
		return &Attendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Attendance{}, errors.New("Record Not Found")
	}
	if(!u.State){
		return &Attendance{}, errors.New("Try Punching In First !! Thank You.")
	}
	err = db.Debug().Model(&Attendance{}).Where("id = ?", u.ID).UpdateColumns(
		map[string]interface{}{
			"state":     false,
		},
	).Error
	if err != nil {
		return &Attendance{}, err
	}
	err = db.Debug().Model(&Log{}).Where("attendance_id = ?", u.ID).Last(&Log{}).UpdateColumns(
		map[string]interface{}{
			"punch_out_at": time.Now(),
		},
	).Error
	if err != nil {
		return &Attendance{}, err
	}
	// This is the display the updated attendance
	err = db.Debug().Model(&Attendance{}).Where("student_id = ?", id).Preload("Logs").Take(&u).Error
	if err != nil {
		return &Attendance{}, err
	}
	return u, nil
}