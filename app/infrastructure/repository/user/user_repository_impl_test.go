package repository

import (
	"testing"
    infrastructureDb "github.com/nagaoka1166/go-tesma-api/app/infrastructure"
	"github.com/nagaoka1166/go-tesma-api/app/domain/entity"
	"context"
    "golang.org/x/crypto/bcrypt"
)

func TestUserRepositoryImpl_Login(t *testing.T) {
    db := infrastructureDb.InitDB()
    repo := NewUserRepo(db)

    // 既存のユーザーを削除
    db.Where("email = ?", "test75@ed.ritsumei.ac.jp").Delete(&entity.User{})

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password1234"), bcrypt.DefaultCost)
    if err != nil {
        t.Fatalf("Failed to hash password: %v", err)
    }

    testUser := &entity.User{
        Email:    "test75@ed.ritsumei.ac.jp",
        Password: string(hashedPassword),
    }

    result := db.Create(testUser)
    if result.Error != nil {
        t.Fatalf("Failed to create test user: %v", result.Error)
    }

    defer db.Delete(testUser)

    user, idToken, err := repo.Login(context.Background(), "test75@ed.ritsumei.ac.jp", "password1234")

    // アサート
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if user == nil {
        t.Fatalf("Expected user to be returned, got nil")
    }
    if user.Email != testUser.Email {
        t.Errorf("Expected email %s; got %s", testUser.Email, user.Email)
    }
    if idToken == "" {
        t.Errorf("Expected idToken to be returned, got empty string")
    }
}
