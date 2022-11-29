package repository

import (
	// "a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {

	// addus := u.db.Select(user).Create(&user)
	// if addus != nil {
	// 	return addus.Error
	// }

	u.db.First(&user)

	leo := u.db.Save(&user)
	if leo != nil {
		return leo.Error
	}

	// result := u.db.Find(&user)
	// if result != nil {
	// 	return result.Error
	// }

	// err := u.db.AutoMigrate(&model.User{})
	// if err != nil {
	// 	return err
	// }

	// res := u.db.Create(&user)
	// if res.Error != nil {
	// 	return res.Error
	// }
	return nil // TODO: replace this
}

func (u *UserRepository) UserAvail(cred model.User) error {
	u.db.First(&cred)
	if len(cred.Username) == 0 {
		return gorm.ErrRecordNotFound
	}

	if len(cred.Password) == 0 {
		return gorm.ErrRecordNotFound
	}

	// result := u.db.First(&cred)

	// // check error ErrRecordNotFound
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	// res := u.db.Create(&cred)
	// if res.Error != nil {
	// 	return res.Error
	// }
	return nil // TODO: replace this
}

func (u *UserRepository) CheckPassLength(pass string) bool {
	if len(pass) <= 5 {
		return true
	}

	return false
}

func (u *UserRepository) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
