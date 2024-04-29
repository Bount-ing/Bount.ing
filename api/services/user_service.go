package services

import (
    "open-bounties-api/models"
    "gorm.io/gorm"
    "errors"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

// FetchAllUsers returns all users from the database
func (s *UserService) FetchAllUsers() ([]models.User, error) {
    var users []models.User
    result := s.db.Find(&users)
    return users, result.Error
}

// FetchUserByID retrieves an user by its ID from the database
func (s *UserService) FindUserById(id uint) (*models.User, error) {
    var user models.User
    result := s.db.First(&user, id)
    return &user, result.Error
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user models.User) (*models.User, error) {
    result := s.db.Create(&user)
    return &user, result.Error
}


func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
    // Example logic for user authentication
    // This should actually check a user repository for a user that matches the credentials
    if username == "admin" && password == "pass" {
        return &models.User{Username: username}, nil
    }
    return nil, errors.New("invalid credentials")
}


func (s *UserService) UpdateUser(id uint, updatedData models.User) (*models.User, error) {
    var user models.User
    result := s.db.First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }

    user.Email = updatedData.Email // Example update field
    user.Username = updatedData.Username   // Update other fields as necessary

    saveResult := s.db.Save(&user)
    if saveResult.Error != nil {
        return nil, saveResult.Error
    }
    return &user, nil
}

func (s *UserService) DeleteUser(id uint) error {
    var user models.User
    result := s.db.First(&user, id)
    if result.Error != nil {
        return result.Error
    }
    deleteResult := s.db.Delete(&user)
    return deleteResult.Error
}
