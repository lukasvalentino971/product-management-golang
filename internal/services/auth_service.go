package services

import (
    "errors"
    "jwt-auth-crud/internal/dto"
    "jwt-auth-crud/internal/models"
    "jwt-auth-crud/internal/repositories"
    "jwt-auth-crud/internal/utils"

    "gorm.io/gorm"
)

type AuthService interface {
    Register(req *dto.RegisterRequest) (*dto.AuthResponse, error)
    Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
}

type authService struct {
    userRepo  repositories.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) AuthService {
    return &authService{
        userRepo:  userRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *authService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
    if err := utils.ValidateStruct(req); err != nil {
        return nil, err
    }

    // Check if user already exists
    _, err := s.userRepo.GetByEmail(req.Email)
    if err == nil {
        return nil, errors.New("user already exists")
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    // Set default role if not provided
    role := req.Role
    if role == "" {
        role = "user"
    }

    // Create user
    user := &models.User{
        Email:    req.Email,
        Password: hashedPassword,
        Name:     req.Name,
        Role:     role,
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    // Generate token
    token, err := utils.GenerateJWT(user.ID, user.Email, user.Role, s.jwtSecret)
    if err != nil {
        return nil, err
    }

    return &dto.AuthResponse{
        Token: token,
        User:  user,
    }, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
    if err := utils.ValidateStruct(req); err != nil {
        return nil, err
    }

    // Get user by email
    user, err := s.userRepo.GetByEmail(req.Email)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("invalid credentials")
        }
        return nil, err
    }

    // Check password
    if !utils.CheckPassword(req.Password, user.Password) {
        return nil, errors.New("invalid credentials")
    }

    // Generate token
    token, err := utils.GenerateJWT(user.ID, user.Email, user.Role, s.jwtSecret)
    if err != nil {
        return nil, err
    }

    return &dto.AuthResponse{
        Token: token,
        User:  user,
    }, nil
}
