package services

import (
    "open-bounties-api/models"
        "gorm.io/gorm"
    "errors"
    // Assume a package that handles database operations, for example, GORM
)

// UserService handles business logic related to users.
type UserService struct {
    // db represents the database client which can be injected or set up here.
    db *gorm.DB
}

// NewUserService creates a new instance of UserService.
func NewUserService(db *gorm.DB) *UserService {
    return &UserService{
        db: db,
    }
}

// CreateUser creates a new user in the database.
func (s *UserService) CreateUser(user models.User) (*models.User, error) {
    if result := s.db.Create(&user); result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// GetUserByID fetches a user by their ID.
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if result := s.db.First(&user, id); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

// UpdateUser updates an existing user's details in the database.
func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
    if result := s.db.Save(&user); result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// DeleteUser removes a user from the database.
func (s *UserService) DeleteUser(id uint) error {
    var user models.User
    if result := s.db.Delete(&user, id); result.Error != nil {
        return result.Error
    }
    return nil
}

