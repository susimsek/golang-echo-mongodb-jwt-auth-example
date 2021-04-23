package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	*UserInput `bson:",inline"`
	ID         primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type UserInput struct {
	FirstName string `json:"firstName" xml:"firstName" bson:"firstName" validate:"required"`
	LastName  string `json:"lastName" xml:"lastName" bson:"lastName" validate:"required"`
	Email     string `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Password  string `json:"password,omitempty" xml:"password,omitempty" bson:"password" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" bson:"password" validate:"required"`
}

type PagedUser struct {
	Data     []User                         `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
