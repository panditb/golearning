Softwares Required
1. go
2. Make 
3. git

create module project 
mkdir go-consume-rest-module
cd go-consume-rest-module

go mod init github.com/panditb/golearning <pass ure git hub repo>

For configuration using the go lang community Viper plugin you can get from below command

go get github.com/spf13/viper

Build 

go build config\config.go
go build main.go

Run 
go run main.go
go run github.com/panditb/golearning

For Cross Flatform we can simply run the Makefile

make all

