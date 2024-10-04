package database

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var firebaseDBClient *db.Client

func Init() {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID:   os.Getenv("FIREBASE_PROJECT_ID"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	opt := option.WithCredentialsFile("./config/credentials.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		panic("Unable to init firebase app:" + err.Error())
	}

	client, err := app.Database(ctx)
	if err != nil {
		panic("Unable to init firebase db:" + err.Error())
	}

	firebaseDBClient = client
}

func GetUserData(username string) User {
	ref := firebaseDBClient.NewRef(fmt.Sprintf("users/%s", username))

	var user User
	if err := ref.Get(context.TODO(), &user); err != nil {
		log.Fatalln("error reading value:", err)
	}

	return user
}

func AddUserData(user User) {
	ref := firebaseDBClient.NewRef("users")

	var emptyInterface interface{}
	newUserRef, err := ref.Push(context.TODO(), emptyInterface)
	if err != nil {
		log.Fatalln("error creating new user:", err)
	}

	if err := newUserRef.Set(context.TODO(), user); err != nil {
		log.Fatalln("error setting user data:", err)
	}
}

func UpdateUserData(user User, username string) {
	ref := firebaseDBClient.NewRef("users/" + username)

	if err := ref.Set(context.TODO(), user); err != nil {
		log.Fatalln("error setting user data:", err)
	}
}

func IsUserExists(username string) bool {
	ref := firebaseDBClient.NewRef("users/" + username)

	var result map[string]interface{}

	err := ref.Get(context.TODO(), &result)

	if err != nil {
		log.Printf("Error checking existence of user: %v\n", err)
		return false
	}

	return result != nil && len(result) > 0
}

func GetCampaignData(campaignName string) Campaign {
	ref := firebaseDBClient.NewRef(fmt.Sprintf("campaigns/%s", campaignName))

	var campaign Campaign
	if err := ref.Get(context.TODO(), &campaign); err != nil {
		log.Fatalln("error reading value:", err)
	}

	return campaign
}

func AddCampaignData(campaign Campaign) {
	ref := firebaseDBClient.NewRef("campaigns")

	var emptyInterface interface{}
	newCampaignRef, err := ref.Push(context.TODO(), emptyInterface)
	if err != nil {
		log.Fatalln("error creating new campaign:", err)
	}

	if err := newCampaignRef.Set(context.TODO(), campaign); err != nil {
		log.Fatalln("error setting campaign data:", err)
	}
}

func GetSecretData() Secret {
	ref := firebaseDBClient.NewRef("secrets")

	var secret Secret
	if err := ref.Get(context.TODO(), &secret); err != nil {
		log.Fatalln("error retrieving secret data:", err)
	}

	return secret
}

// TODO

//func GetRandomKey(path string) (string, error) {
//	ref := firebaseDBClient.NewRef(path)
//	keys := []string{}
//
//	iter := ref.OrderByKey().LimitToFirst(100)
//	for {
//		snapshot, err := iter.Next()
//	}
//
//	randomIndex := rand.Intn(len(keys))
//
//	return keys[randomIndex], nil
//}
