package models

type User struct {
	ID				  string `bson:"_id,omitempty" json:"id"`
	Email  			  string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"_"`
	IsAdmin			  bool   `bson:"isAdmin" json:"isAdmin"`
	Token			  string `bson:"token" json:"token"`
}