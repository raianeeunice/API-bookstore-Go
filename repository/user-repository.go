package repository

import (
	"livraria-go/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository é um contrato com o que o userRepository pode fazer com o banco de dados
type UserRepository interface {
	InsertUser(user entity.Users) entity.Users
	UpdateUser(user entity.Users) entity.Users
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.Users
	ProfileUser(userID string) entity.Users
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository cria uma nova instância de UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.Users) entity.Users {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.Users) entity.Users {
	if user.Password != ""{
	user.Password = hashAndSalt([]byte(user.Password))
	}else{
		var tempUser entity.Users
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.Users
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.Users
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.Users {
	var user entity.Users
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.Users {
	var user entity.Users
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}



func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
