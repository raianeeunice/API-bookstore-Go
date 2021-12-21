package service

import (
	"livraria-go/dto"
	"livraria-go/entity"
	"livraria-go/repository"
	"log"

	"github.com/mashingan/smapping"
)

//UserService é um contrato do que userService pode fazer
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.Users
	Profile(userID string) entity.Users
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService cria uma nova instância de UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.Users {
	userToUpdate := entity.Users{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.Users {
	return service.userRepository.ProfileUser(userID)
}