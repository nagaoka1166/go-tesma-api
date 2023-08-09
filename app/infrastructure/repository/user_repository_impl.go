// Path: app/infrastructure/repository/user_repository_impl.go
package repository

import (
	"context"
	"log"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/nagaoka166/go-tesma-api/app/domain/entity"
	"github.com/nagaoka166/go-tesma-api/app/domain/repository"
	"google.golang.org/api/option"
)

type UserRepoImpl struct {
	FirebaseAuth *auth.Client
}

func NewUserRepo() repository.UserRepository {
    app, err := firebase.NewApp(context.Background(), &firebase.Config{
        CredentialsJSON: []byte(CredentialsJSON),
    })
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }
    auth, err := app.Auth(context.Background())
    if err != nil {
        log.Fatalf("error getting Auth client: %v\n", err)
    }
    return &UserRepoImpl{FirebaseAuth: auth}
}



func (r *UserRepoImpl) Login(ctx context.Context, email, password string) (*entity.User, error) {
	// TODO: 実際のログインロジックを書く
	return nil, nil
}

func (r *UserRepoImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	// TODO: 実際のユーザー更新ロジックを書く
	return nil
}

func (r *UserRepoImpl) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	// TODO: 実際のトークン更新ロジックを書く
	return "", nil
}





func (r *UserRepoImpl) UserExists(ctx context.Context, email string) (bool, error) {
	_, err := r.FirebaseAuth.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("Error from Firebase Auth: %v", err)  // This is new
		if auth.IsUserNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}



func (r *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) error {
	exists, err := r.UserExists(ctx, user.Email)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User already exists")
	}

	params := (&auth.UserToCreate{}).Email(user.Email).Password(user.Password)
	_, err = r.FirebaseAuth.CreateUser(ctx, params)
	return err
}




func (r *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	userRecord, err := r.FirebaseAuth.GetUserByEmail(ctx, email)
	if err != nil {
		if err.Error() == "firebase: user does not exist" {
			// ユーザーが存在しない場合はエラーログを出力せずにnilを返す
			return nil, nil
		}
		// それ以外のエラーが発生した場合は、そのままエラーを返す
		return nil, err
	}
	
	// ユーザーレコードを entity.User に変換
	user := &entity.User{
		Email: userRecord.Email,
	}

	return user, nil
}
