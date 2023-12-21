package user

import (
	model "github.com/enricoanto/final-project/repository"
	userRepository "github.com/enricoanto/final-project/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	userRepository *userRepository.Repository
}

type Claim struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewService(userRepository *userRepository.Repository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) Register(user model.User) (model.User, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.User{}, err
	}

	user.Password = string(genPassword)

	return s.userRepository.CreateUser(user)
}

func (s *Service) Login(userLogin model.User) (string, error) {

	user, err := s.userRepository.FindBy(model.User{Email: userLogin.Email})
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return "", err
	}

	claims := Claim{
		int(user.ID),
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := newToken.SignedString([]byte("Rahasia"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) UpdateBalance(userID int, balance int) (int, error) {
	user := model.User{
		ID: userID,
	}
	user, err := s.userRepository.FindBy(user)
	if err != nil {
		return 0, err
	}

	user.Balance = user.Balance + balance

	err = s.userRepository.UpdateBalance(user.ID, user.Balance)
	if err != nil {
		return 0, err
	}
	return user.Balance, nil
}

func (s *Service) FindBy(user model.User) (model.User, error) {
	return s.userRepository.FindBy(user)
}

