package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
	"github.com/bount-ing/bount.ing/api/tools"
	"gorm.io/gorm"
)

var (
	ErrUserEmailAlreadyExist = errors.New("a user with this mail already exists")
	ErrUserAlreadyVerified   = errors.New("this user has already been activated")
	ErrUserHasNoPasswd       = errors.New("user has no password")
	ErrUserVerifCodeExpired  = errors.New("verification code is expired")
)

func CreateUser(user models.User) error {

	user.Email = strings.ToLower(user.Email)

	var exists bool
	err := db.DB.Model(user).
		Select("count(*) > 0").
		Where("email = ?", user.Email).
		Find(&exists).
		Error
	if exists {
		return ErrUserEmailAlreadyExist
	} else if err != nil {
		return err
	}

	nGens := 0
	for {
		var n int64
		user.VerifCode, err = tools.RandomString(16)
		if err != nil {
			nGens += 1
			if nGens == 3 {
				return err
			}
			continue
		}

		err = db.DB.Model(&models.User{}).Where("verif_code = ?", user.VerifCode).Count(&n).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		} else if n > 0 {
			continue
		}
		break
	}

	user.VerifCodeExpirationTime = time.Now()

	if dbc := db.DB.Create(&user); dbc.Error != nil {
		return dbc.Error
	}

	mailb64 := base64.StdEncoding.EncodeToString([]byte(user.Email))
	log.Println(mailb64)

	userValidationLink := fmt.Sprintf("%s/signup/verify/%s/%s",
		os.Getenv("APP_BASE_URL"),
		mailb64,
		user.VerifCode,
	)
	mailContent := fmt.Sprintf(
		`<h3>Welcome to Bount.ing !</h3>
		<br/>
		<br/>
		<a href="%s" class="cta-button"> Click here to activate your account! </a>
		<br/>
		<br/>
		<br/>
		`,
		userValidationLink,
	)
	err = tools.SendEmail(user.Email, "user creation", mailContent)
	if err != nil {
		db.DB.Delete(&user)
		return err
	}

	return nil
}

func CreateUserPassword(pwd, code string) error {
	if len(pwd) == 0 {
		return ErrUserHasNoPasswd
	}

	err := CheckUserVerifCode(code)
	if err != nil {
		return err
	}

	dbc := db.DB.Model(&models.User{}).Where("verif_code = ?", code).Updates(
		map[string]interface{}{
			"password": pwd,
			"verified": true,
		},
	)
	return dbc.Error
}

func CheckUserVerifCode(verifCode string) error {
	var user models.User
	dbc := db.DB.Where("verif_code = ?", verifCode).Take(&user)
	if dbc.Error != nil {
		return dbc.Error
	}

	if user.Verified {
		return ErrUserAlreadyVerified
	}
	now := time.Now()
	limit := user.VerifCodeExpirationTime.Add(72 * time.Hour)
	if now.After(limit) {
		return ErrUserVerifCodeExpired
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	dbc := db.DB.Find(&users)
	if dbc.Error == gorm.ErrRecordNotFound {
		return users, nil
	}

	return users, dbc.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User

	dbc := db.DB.First(&user, id)
	if dbc.Error == gorm.ErrRecordNotFound {
		return user, nil
	}
	return user, dbc.Error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	dbc := db.DB.Where("email = ?", email).First(&user)
	if dbc.Error == gorm.ErrRecordNotFound {
		return user, nil
	}
	return user, dbc.Error
}

func UpdateUserStripeID(id uint, stripeUserId string) error {
	var user models.User

	result := db.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.StipeAccountID = stripeUserId
	saveResult := db.DB.Save(&user)
	if saveResult.Error != nil {
		return saveResult.Error
	}
	return nil
}
