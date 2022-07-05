package repository

import (
	"context"
	"errors"
	m "go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/utils/errors/check"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type authRepository struct {
	sqldb *gorm.DB
	mongo *mongo.Database
}

func NewAuthRepository(sqldb *gorm.DB, mongodb *mongo.Database) *authRepository {
	return &authRepository{sqldb: sqldb, mongo: mongodb}
}

func (repo authRepository) FindByCode(code string) (users response.User, err error) {
	db := repo.sqldb.Model(&m.User{}).
		Select(`users.*, Role.name as role, MedicalFacility.name as facility`).
		Joins("Role").Joins("MedicalFacility").
		Where("users.code = ?", code).Find(&users)
	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *authRepository) Register(reg m.User) (err error) {
	db := repo.sqldb.Create(&reg)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (repo *authRepository) FindByEmail(email string) (res response.User, err error) {
	db := repo.sqldb.Model(&m.User{}).
		Select(`users.*, Role.name as role, MedicalFacility.name as facility`).
		Joins("Role").Joins("MedicalFacility").
		Where("users.email = ?", email).Find(&res)
	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *authRepository) ChangePassword(u m.User) (err error) {
	db := repo.sqldb.Where("code = ?", u.Code).Updates(&u)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo authRepository) SaveToken(token m.Token) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	db := repo.mongo.Collection("token")
	_, err = db.InsertOne(context.TODO(), token)

	return
}

func (repo authRepository) RevokeToken(token string) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	filter := bson.D{
		{Key: "access_token", Value: token},
	}

	db := repo.mongo.Collection("token")
	res := db.FindOneAndDelete(context.TODO(), filter)

	if res.Err() != nil {
		err = errors.New("invalid or expired token")
	}

	return
}

func (repo authRepository) UpdateToken(old_token m.Token, new_token m.Token) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	filter := bson.D{
		{
			Key:   "$set",
			Value: bson.D{{Key: "refresh_token", Value: new_token.RefreshToken}},
		},
		{
			Key:   "$set",
			Value: bson.D{{Key: "access_token", Value: new_token.AccessToken}},
		},
	}

	db := repo.mongo.Collection("token")
	res := db.FindOneAndUpdate(context.TODO(), old_token, filter)

	if res.Err() != nil {
		err = errors.New("invalid or expired token")
	}

	return
}
