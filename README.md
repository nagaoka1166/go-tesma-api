# 被験者募集アプリケーション
## サービスの目的
[サービスの目的やリプレイス前のアプリについて別リポジトリにて記載してあります](https://github.com/nagaoka1166/rubybook-service-examinee
)

## API設計
[こちらのtesma-openapiのリポジトリにて記載しております
](https://github.com/nagaoka1166/tesma-openapi/blob/main/reference/api.yaml)
## データベース設計
<img src="https://github.com/nagaoka1166/go-tesma-api/assets/69971830/005f04cc-1029-43fc-9fd0-021890569815" width="500">


## アーキテクチャ
このプロジェクトは、クリーンアーキテクチャの原則に基づいて設計されています。クリーンアーキテクチャの目的は、フレームワーク、データベース、UIなどの外部の詳細からビジネスロジックを独立させここ変更やテストを行いやすいように機能の独立性を担保しております。

| 層の名前                                  | 責務 |
| ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| エンティティ層 (app/domain/entity) | アプリケーションの中心的ビジネスルールを表現する層                                      |
| ユースケース層 (app/domain/usecase)                          | プリケーションの具体的なビジネスロジックを担当。エンティティを操作を実装する。 |
| インターフェース層 (app/interfaces)                                  |アプリケーションの入出力のためのインターフェース                         |
| インフラストラクチャ (app/infrastructure)                      | アプリケーションのインフラストラクチャ関連のロジックを管理している。外部のフレームワークやツール、データベースなどとの接続。具体的な技術やフレームワークの実装が含まれている。


## Firebase Authentication の技術選定の背景

モバイルアプリを想定している。認証、ユーザーのUX向上や利用状況を把握できる以下の機能が欲しいモチベーションからFirebase Authenticationを選定しました。

### 1. **Firebase Console**
- ユーザーの管理、確認、削除などの操作をGUI上で簡単に行いたい

### 2. **Firebase Cloud Messaging (FCM)**
- メッセージの返信速度や確認を促したいため、プッシュ通知の送信が必要

### 3. **Firebase Analytics**
- ユーザーの行動をトラッキングして分析できる

### 4. **SDK**
- ログイン後にサーバーに送信されたidトークンのuidからcurrent_userを簡単に取得できる

### 5. **JWT認証**
- 後ほど解説

加えて、GoはFirebase Authenticationに公式にサポートされており、認証関連の実装が効率的に行えます。このようなメリットから、本アプリケーションでFirebase Authenticationを採用することとしました。

## JWT認証

本アプリケーションでは、ユーザー認証にJWT (JSON Web Token) 認証を採用しています。


### JWT認証の背景

JWT認証は、ステートレスでセキュアな方法でユーザー認証情報をトークンとして交換する方式です。Firebase Authenticationとの相性が良く、以下のような理由から本アプリケーションでの採用を決定しました。

1. **セキュリティ**: JWTは暗号化されており、中間者攻撃を防ぐことができる
2. **ステートレス性**: JWTはサーバーサイドでセッションを保持する必要がないため、スケーラビリティが高い。
3. **Firebaseとの統合**: Firebase AuthenticationはJWTをサポートしており、認証処理の実装が容易だったため。


## JWTの認証のデメリットと対策

- **リスク**: セッションを即時無効化できない
- **対策**:
  - 即時無効化はできないが、[Firebaseでは1時間で切れる。](https://firebase.google.com/docs/auth/admin/manage-sessions?hl=ja)
  - 有効期間を短く設定することでリスクを低減することが可能。
  - 更新トークンを適切に利用する。

### 更新トークンを適切に利用するための方針

更新トークンは、期限がないため更新トークンを第三者が取得した際にユーザーのアカウントにアクセスし続ける可能性があるため
Firebase Authenticationでは、以下の条件で更新トークンが無効となります。([Firebaseの公式ドキュメント参照](https://firebase.google.com/docs/auth/admin/manage-sessions?hl=ja))

- ユーザーが削除される場合
- ユーザーが無効になる場合
- ユーザーのアカウントで大きな変更が検出される場合（例: パスワードやメールアドレスの変更）

上記以外のユースケースでら更新トークンを削除したい場合は、クライアント側で`RevokeRefreshTokens`メソッドを使用します。

本アプリでは、被験者への申請時にセキュリティ上の理由から更新トークンを削除し再認証を促す予定です。

## **JWT認証の流れ**

### 1. 認証後の処理

- **IDトークンの返却**: ログイン時やサインアップ時に入力される情報がサーバーサイドとFirebase authで承認された際、IDトークンが返される。
  
- **トークンの保存**: クライアント側では返却されたIDトークンを保存する。（保存先は未定）

### 2. 再認証

リクエスト時にIDトークンをヘッダーに含め、サーバーサイドで認証を行う。

### 3. トークンの更新

IDトークンの有効期間は短いため、期限切れ時に更新トークンを使用して新しいIDトークンを取得します。これにより、ユーザーは再ログインの必要がなくなります。

## Dockerコマンド

db_login:
   `docker exec -it go-tesma-api-db-1 bin/bash`
   `mysql -u user -p`
   ` USE dbname;`
