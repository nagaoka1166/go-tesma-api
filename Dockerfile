# ベースとなるDockerイメージ指定
FROM golang:1.20

# コンテナ内に`app`ディレクトリを作成
WORKDIR /app

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY go.mod .
COPY go.sum .

# モジュールをダウンロード
RUN go mod download

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# Goをコンパイルしてバイナリを生成
RUN go build -o main .

# バイナリを実行
CMD ["./main"]
