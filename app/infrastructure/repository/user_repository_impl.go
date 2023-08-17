// Path: app/infrastructure/repository/user_repository_impl.go
package repository

import (
	"context"
	"log"
	"fmt"
	// "os"
	
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/nagaoka166/go-tesma-api/app/domain/entity"
	"github.com/nagaoka166/go-tesma-api/app/domain/repository"
	"google.golang.org/api/option"
)

type UserRepoImpl struct {
	FirebaseAuth *auth.Client
}

// const credentialsJSON = `
// {
// 	"type": "service_account",
// 	"project_id": "go-tesma",
// 	"private_key_id": "410dd25f804bdfd1a3f4ec23aab8148dd44ccf90",
// 	"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDhhLbkq9Li6jtQ\ngYlWvJQgEzzPaA6TztJaUY5bHGW4SRJd7FzrKXOreXN6jFZQv49B8z2s/6HdVPt2\nc/mo4lkEK4RTiNwJaY0Yq2a+l17kv4z0FdJyryR3MvCyZHJ5VRrX8LLAoS7hyb8v\nWM/8W7MbRigSSq/JhzFTVfjpO8Ro8efk9uGh+aGdgEym1cY9W5/t3ZHZn777zCz4\nOhu88ekxYMMdBGM5TbOo9v6hHZIkR0V8x90CZCAyWy6lyB5UPrSRo0O2hpTQDf14\nHJBGfIq+Y6LC3tbIflEysbXSMP5W7sEm0ZueIg/DCJKnhuPyPFpmHmcjgJTlX44h\nGmHIHhQXAgMBAAECggEAHpZ6nA92HhL5EpVPRHV1fdg+Ij+R4phlzZxMO61hWpqG\npDif7cGEj9h8Qj3woykoTCCq1+EgWj3hkZJvZWm1+yvP2eixN5GacqMk32r2kv4s\nJxsinS2GO6isHvYi1LQnXOynwhramgNcLBbWtvrVEzjzzG9W4nW0ygHaLgBORpFv\nwKphn0yr+JDu38UcLlWgkU43goRGMZC+9G/CEspK9pGuL9lhn8B50K4+hE1if7mj\nFqxNB0iIXyHy5ZEmlecqvPSozhYbvf+HAXZLbtio07wqhWcMqH+grhRYJ4kOJVWC\nhW+HvmGoiLz6vrzQAiBpbGdr4fD7XlEjUw6p9WzhmQKBgQD+0e6NyF4Dq0cg1jgx\ni1u4Nwseik6FGzYCnvBUiAJgnaw9cC1GH00vsjQqK+wCzOOOYkRICwyqVSoIopSI\nN3XWQxLTRrMFP1yYAQNRPAKAZ1eHpLEy32M4e+7mv4IQw+BiH0Dn31rCofURoxDH\nRkUBAPFah9nW2sIvOFrm0PztDQKBgQDikAxDyf1xCmrz+x/QcRc8VFVhAD1c2DyG\npXi09fgYFu009pvNNSDetQKADyFot8cqb+wpeBazMCeOrzEMMH2gET0Dd1/ZSgoo\nH08EqPVerszMBIe5SvzSrn6phY8dFSduWj9d5WMjeTeTdRSpTVjoEu29SIobST1G\n07l9UwikswKBgQDSPfNNF9w5ur9c4GJMvmZWf1O1Rvex6tmNIAW6ON19SZoqtrsj\n9/9/MtDHWqXHSbUWbQ1ZQ4SXpNyhc0KJgPjAZ2bI2rQpyDVlVLCf/ZsyxhD54XAy\nQQ1kNRzPVYLCwloHmzG6HhaML4Q7oQbe46NbSKPZxex90NAWlbCqcBO+oQKBgGdp\nzM9ccXyZfrwi0iC4nhyBef8uttOcN/RVr0UaSMXmIRvdvLvsFjbRdHSzc1JzIEu0\njp1XHTVpM4UjXkXsFEtt0uW0JWHvM7egotbDsEwpiX1OTA09ty0LTPFf7zUdJMRw\nGSBoxdm5SK7b8DRfA8faa/3MEI3n3jpoBfC4Z+eVAoGBAJSjxf50q3e/DpnVPuPB\nusVN1k8/btnoVpknZxp03wND2RngCQ/Dw9c7KpivH+My5cfYWSGiGQ2K8hzgFlqS\nUVksxbRosFy59xpjiU0YqScW09tsnXbQZNeNy0iwjfnqK3z0pHx4nhQVAT/svV6U\nQBS0Kync8IQq2dpphPWQMIkA\n-----END PRIVATE KEY-----\n",
// 	"client_email": "firebase-adminsdk-at6tg@go-tesma.iam.gserviceaccount.com",
// 	"client_id": "104306887339287487570",
// 	"auth_uri": "https://accounts.google.com/o/oauth2/auth",
// 	"token_uri": "https://oauth2.googleapis.com/token",
// 	"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
// 	"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-at6tg%40go-tesma.iam.gserviceaccount.com",
// 	"universe_domain": "googleapis.com"
//   }
// `

// func loadLocalCredentials() string {
// 	data, err := os.ReadFile("Credentials.json")
// 	if err != nil {
// 		log.Fatalf("Failed to load local credentials: %v", err)
// 	}
// 	return string(data)
// }

func NewUserRepo() repository.UserRepository {
	// if os.Getenv("IS_CI") != "" {
	// 	// CI環境
	// 	credentialsJSON := os.Getenv("FIREBASE_CREDENTIALS_JSON")
	// 	if credentialsJSON == "" {
	// 		log.Fatal("FIREBASE_CREDENTIALS_JSON is not set.")
	// 	}
	// 	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON([]byte(credentialsJSON)))
	// 	if err != nil {
	// 		log.Fatalf("error initializing app: %v\n", err)
	// 	}
	// 	auth, err := app.Auth(context.Background())
	// 	if err != nil {
	// 		log.Fatalf("error getting Auth client: %v\n", err)
	// 	}
	// 	return &UserRepoImpl{FirebaseAuth: auth}
	// } else {
		// ローカル環境
		// localCredentials := loadLocalCredentials()
		opt := option.WithCredentialsFile("Credentials.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v", err)
		}
		
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v", err)
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
		log.Printf("Error from Firebase Auth: %v", err)
		if auth.IsUserNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}


func (r *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) error {
    // Check if the user already exists
    exists, err := r.UserExists(ctx, user.Email)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("User already exists")
    }

    // Create the user as the user does not exist
    params := (&auth.UserToCreate{}).Email(user.Email).Password(user.Password)
	_, err = r.FirebaseAuth.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Failed to create user in Firebase: %v", err)
		return err
	}

    log.Println("User successfully created in Firebase")
    return nil
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
