package repository

import (
	"context"
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-echo-mongodb-jwt-auth-example/exception"
	"golang-echo-mongodb-jwt-auth-example/model"
)

var cntx context.Context

type UserRepository interface {
	GetAllUser(page int64, limit int64) (*model.PagedUser, error)
	SaveUser(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	GetUser(id string) (*model.User, error)
	UpdateUser(id string, user *model.User) (*model.User, error)
	DeleteUser(id string) error
}

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func (userRepository *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var existingUser model.User
	filter := bson.M{"email": email}
	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func NewUserRepository(Connection *mongo.Database) UserRepository {
	return &userRepositoryImpl{Connection: Connection}
}

func (userRepository *userRepositoryImpl) GetAllUser(page int64, limit int64) (*model.PagedUser, error) {
	var users []model.User

	filter := bson.M{}

	collection := userRepository.Connection.Collection("users")

	projection := bson.D{
		{"id", 1},
		{"firstName", 1},
		{"lastName", 1},
		{"email", 1},
	}

	paginatedData, err := paginate.New(collection).Context(cntx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedUser{
		Data:     users,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (userRepository *userRepositoryImpl) SaveUser(user *model.User) (*model.User, error) {
	user.ID = primitive.NewObjectID()

	_, err := userRepository.Connection.Collection("users").InsertOne(cntx, user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (userRepository *userRepositoryImpl) GetUser(id string) (*model.User, error) {
	var existingUser model.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	existingUser.Password = ""
	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	result, err := userRepository.Connection.Collection("users").ReplaceOne(cntx, filter, user)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	user.Password = ""
	return user, nil
}

func (userRepository *userRepositoryImpl) DeleteUser(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	result, err := userRepository.Connection.Collection("users").DeleteOne(cntx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return exception.ResourceNotFoundException("User", "id", id)
	}

	return nil
}
