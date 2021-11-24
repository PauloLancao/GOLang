# Go
The exercise contains three approaches for analysis,
1º normal pattern MVC like controller -> Services -> Repositories
2º channels GO style using controller -> Repository

## Mod init
go mod init github.com/{your_username}/{repo_name}

## Mod Tidy
go mod tidy

## Mod Vendor
go mod vendor

## Go packages
go install .\controllers\ .\bootstrap\ .\utils\ 
.\webserver\ .\logging\  .\models\ .\repository\ .\services\ .\di\ .\channels\ .\gerrors\

## Go run
go run -race .\main.go

## Run test from src directory
go test ./... 