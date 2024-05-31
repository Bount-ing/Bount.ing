package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"open-bounties-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedPassword)
	result := s.db.Create(&user)
	return &user, result.Error
}

func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	result := s.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("bcrypt comparison error: %v", err) // Log only on error
		return nil, errors.New("invalid username or password")
	}
	log.Printf("Authentication successful for user: %s", username) // Confirm success only
	return &user, nil
}

func (s *UserService) UpdateUser(id uint, updatedData models.User) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	user.Email = updatedData.Email       // Example update field
	user.Username = updatedData.Username // Update other fields as necessary

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

func (s *UserService) VerifyGitHubToken(token string) (*models.User, error) {
	const githubUserAPIURL = "https://api.github.com/user"
	// Create a new request to the GitHub API to fetch user data
	req, err := http.NewRequest("GET", githubUserAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the OAuth token in the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to GitHub: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned non-OK status: %s", resp.Status)
	}

	// Parse the JSON response into a struct
	var githubUser struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	}
	if err := json.Unmarshal(body, &githubUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Assuming you have a method to find or create your user based on GitHub data
	user, err := s.FindOrCreateUser(githubUser.ID, githubUser.Login)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create user: %w", err)
	}

	return user, nil
}

func (s *UserService) FindOrCreateUser(githubID int, githubLogin string) (*models.User, error) {
	// Check if the user already exists in the database
	var user models.User
	result := s.db.Where("github_id = ?", githubID).First(&user)
	if result.Error == nil {
		return &user, nil
	}

	// If the user does not exist, create a new one
	newUser := models.User{
		GithubID: githubID,
		Username: githubLogin,
	}
	result = s.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
