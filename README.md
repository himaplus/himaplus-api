# himaplus-api

ひまぷらのAPIサーバ

## 概要

`3レイヤーアーキテクチャ`を採用したGo-ginでのAPIサーバ
なお、物理的に三層のサーバーを使う`三層アーキテクチャ`（物理三層アーキテクチャ）ではない

- presentaison（handler）: リクエストの受け取り、レスポンスの返却
- application（service）: ユースケースのワークフロー
- data_access: DBへORMでアクセスする

### 環境

Visual Studio Code: 1.88.1  
golang.Go: v0.41.4  
image Golang: go version go1.22.2 linux/amd64

## 環境構築

[docker-himaplus](https://github.com/unSerori/docker-himaplus)を使ってDokcerコンテナーで開発・デプロイする形を想定している  
インストール手順は[docker-himaplusのインストール項目](https://github.com/unSerori/docker-himaplus/blob/main/README.md#インストール)に記載  
cloneしてスクリプト実行で、自動的にコンテナー作成と開発環境（: またはデプロイ）を行う  

### 自前でのローカル環境構築

1. [Goのインストール](https://go.dev/doc/install)
2. このリポジトリをclone

    ```bash
    git clone https://github.com/unSerori/himaplus-api
    ```

3. [.env](#env)ファイルを作成
4. assetsフォルダ内で必要なものをもらう
5. 必要なパッケージの依存関係など

    ```bash
    go mod tidy
    ```

6. プロジェクトを起動

    ```bash
    # 実行(VSCodeならF5keyで実行)
    go run .

    # ワンファイルにビルド
    go build -o output 

    # ビルドで生成されたファイルを実行
    ./output
    ```

#### vscode-ext.txt

Goのデバッグ用VS Code拡張機能や便利な拡張機能のリスト  
以下のコマンドで一括インストールできる

```bash
cat vscode-ext.txt | while read line; do code --install-extension $line; done
```

#### おまけ: Goでプロジェクト作成時のコマンド

```bash
# goモジュールの初期化
go mod init himaplus-api

# ginのインストール
go get -u github.com/gin-gonic/gin

# main.goの作成
echo package main > main.go
```

## ディレクトリ構成

TODO: ディレクトリ構成

## API仕様書

エンドポイント、リクエストレスポンスの形式、その他情報のAPIの仕様書。

### エンドポインツ

<details>
  <summary>カレンダーの情報を取得するエンドポイント</summary>

- **URL:** `/v1/auth/users/calendars`
- **メソッド:** GET
- **説明:** ログインしているアカウントのGoogleカレンダーをAPIから取得する
- **リクエスト:**
  - ヘッダー:
    - `Content-Type`: application/json,
    - `Authorization`: (string) 認証トークン
  -  - ボディ:

- **レスポンス:**
  - ステータスコード:200 OK
    - ボディ:JSON
      ```json
      {
        "srvResMsg":  "OK",
        "srvResData": {
          "calendar": [
            {
            "date": "日付",
            "todos"[
                    "title":";タスクの名前",
                    "time":"かかる時間",
                    "place":"場所",
                    "explanation":"タスクのメモ",
            ]
            }
          ],
        },
      }
      ```

</details>

<details>
  <summary>カレンダーに予定を追加するエンドポイント</summary>

- **URL:** `/v1/auth/users/calendars`
- **メソッド:** POST
- **説明:** ログインしているアカウントのGoogleカレンダーに予定を追加する
- **リクエスト:**
  - ヘッダー:
    - `Content-Type`: application/json
    - `Authorization`: (string) 認証トークン

  - ボディ:
    ```
      {
          "calendar"[
            "date":"日付",
            "todos"[
                    "title":";タスクの名前",
                    "time":"かかる時間",
                    "place":"場所",
                    "explanation":"タスクのメモ",
            ]
          ]
      }
      ```

- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:JSON
      ```json
      {
        "srvResMsg":  "OK",
        "srvResData": {},
      }
      ```

</details>

### API仕様書てんぷれ

<details>
  <summary>＊○○＊するエンドポイント</summary>

- **URL:** `/＊エンドポイントパス＊`
- **メソッド:** ＊HTTPメソッド名＊
- **説明:** ＊○○＊
- **リクエスト:**
  - ヘッダー:
    - `＊HTTPヘッダー名＊`: ＊HTTPヘッダー値＊
  - ボディ:
    ＊さまざまな形式のボディ値＊

- **レスポンス:**
  - ステータスコード: ＊ステータスコード ステータスメッセージ＊
    - ボディ:
      ＊さまざまな形式のレスポンスデータ（基本はJSON）＊

      ```json
      {
        "srvResMsg":  "レスポンスステータスメッセージ",
        "srvResData": {
        
        },
      }
      ```

</details>

## .ENV

.evnファイルの各項目と説明

```env:.env
MYSQL_USER=hima_user: DBに接続する際のログインユーザ名
MYSQL_PASSWORD=hima_pass: パスワード
MYSQL_HOST=mysql-db-srv: ログイン先のDBホスト名（dockerだとサービス名）
MYSQL_PORT=3306: ポート番号（dockerだとコンテナのポート）
MYSQL_DATABASE=hima_db: 使用するdatabase名
JWT_SECRET_KEY=qawsedrftgyhujikolp=: "openssl rand -base64 32"で作ったJWTトークン作成用のキー
JWT_TOKEN_LIFETIME=315360000: JWTトークンの有効期限
MULTIPART_IMAGE_MAX_SIZE=10485760: Multipart/form-dataの画像の制限サイズ（10MBなら10485760）
REQ_BODY_MAX_SIZE=52428800: リクエストボディのマックスサイズ（50MBなら52428800）
```
