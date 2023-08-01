// Path: app/infrastructure/repository/user_repository_impl.go
package repository

import (
	"context"
	"log"

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
	opt := option.WithCredentialsFile("credentials.json") // Firebase の秘密鍵へのパス
	app, err := firebase.NewApp(context.Background(), nil, opt)
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


func (r *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) error {
	params := (&auth.UserToCreate{}).Email(user.Email).Password(user.Password)
	_, err := r.FirebaseAuth.CreateUser(ctx, params)
	return err
}


func (r *UserRepoImpl) UserExists(ctx context.Context, email string) (bool, error) {
	_, err := r.FirebaseAuth.GetUserByEmail(ctx, email)
	if err != nil {
		// Firebaseからのエラーメッセージは err.Error() で取得できます
		if err.Error() == "firebase: user does not exist" {
			return false, nil
		}
		return false, err
	}
	return true, nil
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
		// Password は取得できませんが、他の必要なフィールドをここで設定できます。
	}

	return user, nil
}





// その他の UserRepository インターフェースのメソッドを実装します
// GetUserByEmail, UpdateUser, RefreshToken, Login...
