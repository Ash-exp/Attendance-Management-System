package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)
type TeacherLog struct {
	AttendanceID uint32		`gorm:"not null" json:"attendance_id"`
	PunchInAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP;not null" json:"punchedIn_at"`
	PunchOutAt 	time.Time 	`json:"punchedOut_at"`
}
type TeacherAttendance struct {
	ID        	uint32    			`gorm:"primary_key;auto_increment" json:"id"`
	TeacherId 	uint32 	  			`gorm:"not null" json:"teacherId"`
	State		bool				`gorm:"not null;default: true" json:"state"`
	Day  		int    				`gorm:"not null;" json:"day"`
	Month  		time.Month  		`gorm:"not null;" json:"month"`
	Year  		int    				`gorm:"not null;" json:"year"`
	TeacherLogs		[]TeacherLog 		`gorm:"foreignkey:AttendanceID" json:"teacher_logs"`	
}


func (u *TeacherAttendance) Prepare() {
	u.ID = 0
	u.TeacherLogs = []TeacherLog{{PunchInAt : time.Now(), AttendanceID: u.ID}}
	year, month, day := time.Now().Date()
	u.Day = day
	u.Month = month
	u.Year = year

}

func (u *TeacherAttendance) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		return nil

	default:
		if u.TeacherId == 0 {
			return errors.New("Required TeacherId")
		}
		return nil
	}
}

func (u *TeacherAttendance) SaveTeacherAttendance(db *gorm.DB) (*TeacherAttendance, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	return u, nil
}

func (u *TeacherAttendance) TeacherAttendanceExists(db *gorm.DB, teacherId uint32) (*TeacherAttendance, error) {
	var err error
	err = db.Debug().Model(&TeacherAttendance{}).Where("teacher_id = ? AND year = ? AND month = ? AND day = ?", teacherId, u.Year, u.Month, u.Day).Preload("TeacherLogs").Take(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return &TeacherAttendance{}, errors.New("Record Not Found")
	}
	if err != nil {
		return &TeacherAttendance{}, err
	}
	return u, err
}

func (u *TeacherAttendance) FindTeacherAttendanceByID(db *gorm.DB, teacherId uint32) (*[]TeacherAttendance, error) {
	var err error
	attendance := []TeacherAttendance{}
	err = db.Debug().Model(&TeacherAttendance{}).Where("teacher_id = ? AND year = ? AND month = ?", teacherId, u.Year, u.Month).Preload("TeacherLogs").Limit(100).Find(&attendance).Error
	if err != nil {
		return &[]TeacherAttendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]TeacherAttendance{}, errors.New("Record Not Found")
	}
	return &attendance, err
}

func (u *TeacherAttendance) UpdatePunchInTeacherAttendance(db *gorm.DB, id uint32) (*TeacherAttendance, error) {

	err := db.Debug().Model(&TeacherAttendance{}).Where("id = ?", id).Take(&u).UpdateColumns(
		map[string]interface{}{
			"state":     true,
		},
	).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	err = db.Debug().Create(&TeacherLog{
		AttendanceID: u.ID,
		PunchInAt: time.Now(),
	}).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	
	// This is the display the updated teacher attendance
	err = db.Debug().Model(&TeacherAttendance{}).Preload("TeacherLogs").Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &TeacherAttendance{}, errors.New("Record Not Found")
	}
	return u, nil
}

func (u *TeacherAttendance) UpdatePunchOutTeacherAttendance(db *gorm.DB, id uint32) (*TeacherAttendance, error) {

	err := db.Debug().Model(&TeacherAttendance{}).Where("teacher_id = ?", id).Take(&u).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &TeacherAttendance{}, errors.New("Record Not Found")
	}
	if(!u.State){
		return &TeacherAttendance{}, errors.New("Try Punching In First !! Thank You.")
	}
	err = db.Debug().Model(&TeacherAttendance{}).Where("id = ?", u.ID).UpdateColumns(
		map[string]interface{}{
			"state":     false,
		},
	).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	err = db.Debug().Model(&TeacherLog{}).Where("attendance_id = ?", u.ID).Last(&Log{}).UpdateColumns(
		map[string]interface{}{
			"punch_out_at": time.Now(),
		},
	).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	// This is the display the updated teacher attendance
	err = db.Debug().Model(&TeacherAttendance{}).Where("teacher_id = ?", id).Preload("TeacherLogs").Take(&u).Error
	if err != nil {
		return &TeacherAttendance{}, err
	}
	return u, nil
}