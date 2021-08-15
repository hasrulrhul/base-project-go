# menginisalisasi projek baru
go mod init base-project-go
buat file main.go

# jalankan project
go run main.go
go build
atau

# buat file makefile kemudian
run makefile dengan perintah make dev

# install echo
go get -u github.com/labstack/echo/v4

# publish vendor
go mod vendor

# package godotenv
go get package github.com/joho/godotenv

# buat file .env

# migration

# gorm 
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# mysql
go get package github.com/go-sql-driver/mysql
