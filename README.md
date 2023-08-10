# go-tesma-api

### データベース設計
<img src="https://github.com/nagaoka1166/go-tesma-api/assets/69971830/df7226f5-1ce1-4677-84e9-01d706932dcd" width="500">

## アーキテクチャ
このプロジェクトは、クリーンアーキテクチャの原則に基づいて設計されています。クリーンアーキテクチャの目的は、フレームワーク、データベース、UIなどの外部の詳細からビジネスロジックを独立させここ変更やテストを行いやすいように機能の独立性を担保しております。

### 構造

**エンティティ層 (app/domain/entity**)

アプリケーションの中心的ビジネスルールを表現する層
**
ユースケース層 (app/domain/usecase)**

アプリケーションの具体的なビジネスロジックを担当。エンティティを操作を実装する。

**インターフェース層 (app/interfaces)**

アプリケーションの入出力のためのインターフェースを提供。

**インフラストラクチャ (app/infrastructure)**

アプリケーションのインフラストラクチャ関連のロジックを管理している。
説外部のフレームワークやツール、データベースなどとの接続。具体的な技術やフレームワークの実装が含まれている。
