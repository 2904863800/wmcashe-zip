@echo off

chcp 65001

echo 交叉编译 myzip.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/src/bin/myzip ./lib/myzip.go

echo linux 编译完成

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o build/src/bin/myzip.exe ./lib/myzip.go

echo windows 编译完成

echo 再见

