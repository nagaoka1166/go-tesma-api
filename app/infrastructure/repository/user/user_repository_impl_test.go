package repository

func TestUserRepositoryImpl_Login(t *testing.T) {
    db, _ := gorm.Open(/* テスト用データベースの設定 */)
    repo := repository.NewUserRepoImpl(db)

    // テストデータの挿入
    testUser := &entity.User{
        Email:    "test@email.com",
        Password: "hashed_password",  // 実際はbcryptなどでハッシュ化する
    }
    db.Create(testUser)

    // テスト対象メソッドの実行
    user, err := repo.Login(context.Background(), "test@email.com", "password")

    // アサート
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if user.Email != testUser.Email {
        t.Errorf("Expected email %s; got %s", testUser.Email, user.Email)
    }
}
