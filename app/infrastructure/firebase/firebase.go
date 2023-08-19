// app/infrastructure/firebase/firebase.go
package firebase

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	App    *firebase.App
	Client *auth.Client
)

func InitializeFirebase(credentialsPath string) error {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	App = app
	client, err := app.Auth(context.Background())
	if err != nil {
		return err
	}

	Client = client
	return nil
}



func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return Client.VerifyIDToken(ctx, idToken)
}

func GetUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error) {
	return Client.GetUserByEmail(ctx, email)
}

func CreateUserWithFirebase(ctx context.Context, user *auth.UserToCreate) (*auth.UserRecord, error) {
	return Client.CreateUser(ctx, user)
}
