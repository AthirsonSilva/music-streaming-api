package repositories

import (
	"context"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/database"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserById(id string) (models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	var user models.User

	if err != nil {
		return user, err
	}

	err = database.UserCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return user, err
	}

	err = database.UserCollection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User

	err := database.UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	err = database.UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := database.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUserByID(user models.User) (models.User, error) {
	user.UpdatedAt = time.Now()

	_, err := database.UserCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return user, err
	}

	return user, nil
}
