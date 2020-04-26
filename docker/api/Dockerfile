# ベースとなるDockerイメージ指定
FROM golang:1.8

# Docker上の作業ディレクトリ設定
WORKDIR /go/src/work

# ローカルのファイルをdocker にコピー
ADD . .

# 必要なパッケージの取得
RUN apt-get update -qq && apt-get install -y mysql-client vim git

RUN go build
