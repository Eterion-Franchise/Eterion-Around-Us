package database

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var FirebaseDBClient *db.Client

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

	FirebaseDBClient = client
}
