# ベースとなるDockerイメージ指定
FROM golang:latest

# Docker上の作業ディレクトリ設定
WORKDIR /go/src/github.com/AwataKyosuke/go_api_server

# 必要なパッケージの取得
RUN apt-get update -qq && apt-get install -y mariadb-client vim git lsof
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/ant0ine/go-json-rest/rest
RUN go get github.com/jinzhu/gorm
RUN go get github.com/pkg/errors
RUN go get gopkg.in/ini.v1
RUN go get github.com/derekparker/delve/cmd/dlv
