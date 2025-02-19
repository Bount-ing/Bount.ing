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
	ErrResetCodeExpired      = errors.New("password reset code is expired")
	ErrInvalidResetCode      = errors.New("invalid reset code")
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
		//check if user is already verified
		var u models.User
		err = db.DB.Where("email = ?", user.Email).First(&u).Error
		if err == nil && u.Verified {
			return ErrUserEmailAlreadyExist
		} else if err == nil && !u.Verified {
			//delete user and create a new one
			db.DB.Delete(&u)

		}
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

func GetExternalIdentityByUserIDAndHostID(claimerID, hostID uint) (models.Identity, error) {
	var identity models.Identity

	dbc := db.DB.Where("user_id = ? AND host_id = ?", claimerID, hostID).First(&identity)
	if dbc.Error == gorm.ErrRecordNotFound {
		return identity, nil
	}
	return identity, dbc.Error
}

func RequestPasswordReset(email, originIP string) error {
	timestamp := time.Now().Format(time.RFC3339)

	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// If user not found, send an invitation to create an account
		inviteLink := fmt.Sprintf("%s/signup", os.Getenv("APP_BASE_URL"))
		mailContent := fmt.Sprintf(
			`<h3>Password Reset Request</h3>
			<p>A password reset was requested for your account.</p>
			<p><strong>Request Time:</strong> %s</p>
			<p><strong>Origin IP:</strong> %s</p>
			<p>If you didn't make this request, please contact <a href="mailto:support@bount.ing">support@bount.ing</a> immediately.</p>
			<br/>
			<h3>Account Not Found</h3>
			<p>We couldn't find an account with this email address.</p>
			<p>If you want to create an account, click the link below:</p>
			<a href="%s" class="cta-button">Create an Account</a>`,
			timestamp, originIP, inviteLink)
		return tools.SendEmail(email, "Create Your Account", mailContent)
	}

	// Generate reset code
	resetCode, err := tools.RandomString(16)
	if err != nil {
		return err
	}

	// Update user with reset code
	user.VerifCode = resetCode
	user.VerifCodeExpirationTime = time.Now().Add(time.Hour) // Expires in 1 hour
	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}

	// Send reset email
	resetLink := fmt.Sprintf("%s/reset-password/%s", os.Getenv("APP_BASE_URL"), resetCode)
	mailContent := fmt.Sprintf(
		`<h3>Password Reset Request</h3>
		<p>A password reset was requested for your account.</p>
		<p><strong>Request Time:</strong> %s</p>
		<p><strong>Origin IP:</strong> %s</p>
		<p>If you didn't make this request, please contact <a href="mailto:support@bount.ing">support@bount.ing</a> immediately.</p>
		<br/>
		<p>Your reset code is: <strong>%s</strong></p>
		<br/>
		<p>Or click the link below to reset your password:</p>
		<a href="%s" class="cta-button">Reset Password</a>
		<br/>
		<p>This reset code will expire in 1 hour.</p>`,
		timestamp, originIP, resetCode, resetLink,
	)

	return tools.SendEmail(user.Email, "Password Reset Request", mailContent)
}

func ResetPassword(code, newPassword string) error {
	var user models.User
	if err := db.DB.Where("verif_code = ?", code).First(&user).Error; err != nil {
		return ErrInvalidResetCode
	}

	// Check if code is expired (1 hour validity)
	if time.Since(user.VerifCodeExpirationTime) > time.Hour {
		return ErrResetCodeExpired
	}

	// Update password and clear reset code
	user.Password = newPassword
	user.VerifCode = ""
	user.VerifCodeExpirationTime = time.Time{}

	return db.DB.Save(&user).Error
}

func UserLogin(userID uint) error {
	user := models.User{}
	if err := db.DB.First(&user, userID).Error; err != nil {
		return err
	}

	user.LastLogin = time.Now()

	return db.DB.Save(&user).Error
}

func UpdateUserProfileInfo(userID uint, payload models.UserProfileUpdatePayload) error {
	user := models.User{}
	if err := db.DB.First(&user, userID).Error; err != nil {
		return err
	}

	// Check only non null and no empty strings
	if payload.Username != "" {
		user.Username = payload.Username
	}

	if payload.FullName != "" {
		user.FullName = payload.FullName
	}

	return db.DB.Save(&user).Error
}
