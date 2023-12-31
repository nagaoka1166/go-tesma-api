// Pth: app/infrastructure/repository/user_repository_impl.go
package repository

import (
	"context"
	"log"
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
	
	
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/nagaoka1166/go-tesma-api/app/domain/entity"
	"github.com/nagaoka1166/go-tesma-api/app/domain/repository"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
	FirebaseAuth *auth.Client
}


func NewUserRepo(db *gorm.DB) repository.UserRepository {
	if os.Getenv("IS_CI") != "" {
		// CI環境
		credentialsJSON := os.Getenv("FIREBASE_CREDENTIALS_JSON")
		if credentialsJSON == "" {
			log.Fatal("FIREBASE_CREDENTIALS_JSON is not set.")
		}
		app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON([]byte(credentialsJSON)))
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		if err := db.AutoMigrate(&entity.User{}); err != nil {
			log.Fatalf("Failed to auto migrate User: %v", err)
		}

		return &UserRepoImpl{
			DB:  db,
            FirebaseAuth: auth,
        }
	} else {
		// opt := option.WithCredentialsFile("/Users/nagaokaryuunotasuku/go-tesma-api/Credentials.json") #テストの時絶対パスで読み込む
		opt := option.WithCredentialsFile("./Credentials.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v", err)
		}
		
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v", err)
		}

		if err := db.AutoMigrate(&entity.User{}); err != nil {
			log.Fatalf("Failed to auto migrate User: %v", err)
		}
		
		return &UserRepoImpl{
		DB:    db,
		FirebaseAuth: auth,
	}
	
	}
}


func (r *UserRepoImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	// TODO: 実際のユーザー更新ロジックを書く
	return nil
}


func (r *UserRepoImpl) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	client := r.FirebaseAuth
	
	_, err := client.VerifyIDTokenAndCheckRevoked(ctx, refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to verify and refresh ID token: %v", err)
	}

	return refreshToken, nil
}

func (r *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
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


func (r *UserRepoImpl) CreateUser(ctx context.Context, user *entity.User) (string, error) {
    // Check if the user already exists
    exists, err := r.UserExists(ctx, user.Email)
    if err != nil {
        return "", err
    }
    if exists {
        return "", fmt.Errorf("User already exists")
    }

    originalPassword := user.Password

    // パスワードをハッシュ化
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Failed to hash password: %v", err)
        return "", fmt.Errorf("User already exists")
    }
    user.Password = string(hashedPassword)

    if err := r.DB.Create(&entity.User{
        Email:    user.Email,
        Password: user.Password,
    }).Error; err != nil {
        log.Printf("Failed to create user in MySQL: %v", err)
        return "", err
    }
    log.Println("User successfully created in MySQL")

    // Firebase registration
    params := (&auth.UserToCreate{}).Email(user.Email).Password(originalPassword) // オリジナルのパスワードを使用
    firebaseUser, err := r.FirebaseAuth.CreateUser(ctx, params)
    if err != nil {
        log.Printf("Failed to create user in Firebase: %v", err)

        // Firebase失敗時にMySQLからロールバック
        if delErr := r.DB.Delete(&entity.User{
            Email: user.Email,
        }).Error; delErr != nil {
            log.Printf("Failed to delete user from MySQL after Firebase registration failure: %v", delErr)
        }

        return "",err
    }

	idToken, err := r.FirebaseAuth.CustomToken(ctx, firebaseUser.UID)
	    if err != nil {
        log.Printf("Failed to create custom token for user: %v", err)
        return "", err
    }

	

    log.Println("User successfully created in Firebase")
    return idToken, nil
}



func (r *UserRepoImpl) Login(ctx context.Context, email, password string) (*entity.User, string, error) {
    // Firebase Authenticationのユーザー情報取得
    fbUser, err := r.FirebaseAuth.GetUserByEmail(ctx, email)
    if err != nil {
        return nil, "", err
    }

    localUser, err := r.GetUserByEmail(ctx, email)
    if err != nil {
        return nil, "", err
    }

	if err := bcrypt.CompareHashAndPassword([]byte(localUser.Password), []byte(password)); err != nil {
        return nil, "", fmt.Errorf("invalid password")
    }


    idToken, tokenErr := r.FirebaseAuth.CustomToken(ctx, fbUser.UID)
    if tokenErr != nil {
        return nil, "", tokenErr
    }

    return localUser, idToken, nil
}
