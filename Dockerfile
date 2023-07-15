# ベースとなる Docker イメージ指定
FROM golang:1.20

# コンテナ内に `app` ディレクトリを作成
WORKDIR /app

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# Go をコンパイルしてバイナリを生成
RUN go build -o main .

# バイナリを実行
CMD ["/app/main"]