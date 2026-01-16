package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID				  string `bson:"_id,omitempty" json:"id"`
	Email  			  string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"_"`
	IsAdmin			  bool   `bson:"isAdmin" json:"isAdmin"`
	Token			  string `bson:"token" json:"token"`
}

//creatring new user
func NewUser(email, password string) (*User, error){
	epw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Email: 			   email,
		EncryptedPassword: string(epw),
	}, nil
}

