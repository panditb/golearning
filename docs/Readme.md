***Go Consume Rest Module***

# Softwares Required
1. go https://golang.org/dl/
2. Make (in ubuntu already available)
3. git (git bash)

# Steps Followed in windows 10
1. Create module project \
   `mkdir go-consume-rest-module` \
   `cd go-consume-rest-module` \
   `go mod init github.com/panditb/golearning` <pass ure git hub repo> \
2. For configuration using the go lang community Viper plugin you can get from below command \
`go get github.com/spf13/viper` \
3. Build  \
`go build config\config.go` \
`go build main.go` \
4. Run \
 `go run main.go` \
 `go run github.com/panditb/golearning` \

### For Cross Flatform we can simply run the Makefile
Below steps how to run in Ubuntu
1. Get the code for git \
   `go get -u github.com/panditb/golearning`
2. Build  \ 
  `make all`
3. Run \
 `make run`
