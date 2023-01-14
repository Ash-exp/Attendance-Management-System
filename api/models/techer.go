package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Teacher struct {
	ID        	uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Firstname  	string    `gorm:"size:255;not null;unique" json:"firstName"`
	Lastname  	string    `gorm:"size:255;not null;unique" json:"lastName"`
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


func (u *Teacher) Prepare() {
	u.ID = 0
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Teacher) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		return nil

	default:
		if u.Firstname == "" {
			return errors.New("Required Firstname")
		}
		if u.Lastname == "" {
			return errors.New("Required Lastname")
		}
		return nil
	}
}

func (u *Teacher) SaveTeacher(db *gorm.DB) (*Teacher, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Teacher{}, err
	}
	return u, nil
}

func (u *Teacher) FindAllTeachers(db *gorm.DB) (*[]Teacher, error) {
	var err error
	teachers := []Teacher{}
	err = db.Debug().Model(&Teacher{}).Limit(100).Find(&teachers).Error
	if err != nil {
		return &[]Teacher{}, err
	}
	return &teachers, err
}

func (u *Teacher) FindTeacherByID(db *gorm.DB, uid uint32) (*Teacher, error) {
	var err error
	err = db.Debug().Model(Teacher{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Teacher{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Teacher{}, errors.New("Teacher Not Found")
	}
	return u, err
}

func (u *Teacher) UpdateATeacher(db *gorm.DB, uid uint32) (*Teacher, error) {

	db = db.Debug().Model(&Teacher{}).Where("id = ?", uid).Take(&Teacher{}).UpdateColumns(
		map[string]interface{}{
			"firstname":  u.Firstname,
			"lastname":  u.Lastname,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Teacher{}, db.Error
	}
	// This is the display the updated teacher
	err := db.Debug().Model(&Teacher{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Teacher{}, err
	}
	return u, nil
}

func (u *Teacher) DeleteATeacher(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Teacher{}).Where("id = ?", uid).Take(&Teacher{}).Delete(&Teacher{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}