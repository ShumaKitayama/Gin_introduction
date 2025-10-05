# Gin Introduction - Go学習プロジェクト

## 概要
このプロジェクトは、Go言語のWebフレームワーク「Gin」を学習するための実践的なRESTful API開発プロジェクトです。商品管理システムとユーザー認証機能を実装しており、Ginの基本的な使い方からデータベース連携、認証機能まで学ぶことができます。

## 技術スタック
- **言語**: Go 1.23.3
- **Webフレームワーク**: Gin
- **ORM**: GORM
- **データベース**: PostgreSQL
- **管理ツール**: pgAdmin4
- **コンテナ**: Docker & Docker Compose
- **認証**: JWT (JSON Web Token)

## プロジェクト構成

### アーキテクチャ
このプロジェクトは、レイヤードアーキテクチャ（層化アーキテクチャ）を採用しており、責任の分離が明確になっています。

```
├── main.go                  # エントリーポイント
├── controllers/             # コントローラー層（HTTPリクエスト処理）
│   ├── auth_controller.go   # 認証API
│   └── item_controller.go   # 商品管理API
├── services/                # サービス層（ビジネスロジック）
│   ├── auth_service.go      # 認証ビジネスロジック
│   └── item_service.go      # 商品管理ビジネスロジック
├── repositories/            # リポジトリ層（データアクセス）
│   ├── auth_repository.go   # ユーザーデータアクセス
│   └── item_repository.go   # 商品データアクセス
├── models/                  # データモデル
│   ├── user.go             # ユーザーモデル
│   └── item.go             # 商品モデル
├── dto/                     # データ転送オブジェクト
│   ├── auth_dto.go         # 認証用DTO
│   └── item_dto.go         # 商品用DTO
├── infra/                   # インフラ層
│   ├── db.go               # データベース接続
│   └── intializer.go       # 初期化処理
├── migrations/              # データベースマイグレーション
└── docker/                  # Docker関連設定
```

## 機能

### 商品管理API
- `GET /items` - 商品一覧取得
- `GET /items/:id` - 商品詳細取得
- `POST /items` - 商品作成
- `PUT /items/:id` - 商品更新
- `DELETE /items/:id` - 商品削除

### 認証API
- `POST /auth/signup` - ユーザー登録
- `POST /auth/login` - ログイン

## セットアップ

### 前提条件
- Go 1.23.3以上
- Docker & Docker Compose

### 環境構築

1. **リポジトリのクローン**
```bash
git clone <repository-url>
cd Gin_introduction
```

2. **データベース起動**
```bash
docker-compose up -d
```

3. **依存関係のインストール**
```bash
go mod download
```

4. **アプリケーション起動**
```bash
go run main.go
```

### データベース設定
- **PostgreSQL**: `localhost:5432`
- **pgAdmin**: `http://localhost:81`
  - Email: `gin@example.com`
  - Password: `ginpassword`

## 学習ポイント

### 1. Ginフレームワークの基本
- ルーティング設定
- HTTPハンドラーの実装
- ミドルウェアの使用
- JSONレスポンスの処理

### 2. レイヤードアーキテクチャ
- Controller - Service - Repository パターン
- 責任の分離と依存性注入
- ビジネスロジックとデータアクセスの分離

### 3. データベース連携
- GORMによるORM操作
- PostgreSQLとの接続
- マイグレーション

### 4. 認証・認可
- JWT トークンによる認証
- パスワードハッシュ化
- セキュアなAPI設計

### 5. Docker環境
- コンテナ化されたデータベース
- 開発環境の統一

## API使用例

### 商品作成
```bash
curl -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{
    "name": "サンプル商品",
    "price": 1500,
    "description": "これはサンプル商品です",
    "sold_out": false
  }'
```

### ユーザー登録
```bash
curl -X POST http://localhost:8080/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

## 今後の拡張案
- [ ] JWTミドルウェアによる認証保護
- [ ] 商品画像アップロード機能
- [ ] ページネーション実装
- [ ] バリデーション強化
- [ ] テストコード追加
- [ ] OpenAPI/Swagger文書化

## 参考資料
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM Guide](https://gorm.io/docs/)
- [Go Documentation](https://golang.org/doc/)

---
このプロジェクトは学習目的で作成されており、実際のプロダクション環境で使用する場合は、セキュリティとパフォーマンスの追加考慮が必要です。