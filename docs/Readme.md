***Go Consume Rest Module***

# Softwares Required
1. go https://golang.org/dl/
2. Make (in ubuntu already available)
3. git (git bash)

# Steps Followed in windows 10
1. Create module project \
   `mkdir go-consume-rest-module` \
   `cd go-consume-rest-module` \
   `go mod init github.com/panditb/golearning` <pass ure git hub repo> 
2. For configuration using the go lang community Viper plugin you can get from below command \
`go get github.com/spf13/viper` 
3. Build  \
`go build config\config.go` \
`go build main.go` 
4. Run \
 `go run main.go` \
 `go run github.com/panditb/golearning` 
5. Passing env variable and RUN  \
 `set url=http://dummy.restapiexample.com/api/v1/employees` \
 `go run main.go` 

Note: Code is checking if env is available then use from env variable else fron config file

### For Cross Flatform we can simply run the Makefile
Below steps how to run in Ubuntu
1. Clone the code  or if ure using winodws ubunto go to the mounted path where the code exist \
   `git clone https://github.com/panditb/golearning.git`
2. Get the code for git \
   `go get -u github.com/panditb/golearning`
3. Build  \
  `make all`
4. Run  \
 `make run`

### Build Software in Cross Platform from Windows OS

1. Windows Software Build Open Command Prompt \

   `set GOOS=windows` \
   `set GOARCH=amd64` \
   `go build -o mainwindows.exe github.com/panditb/golearning` \
   `mainwindows.exe`
2. Linux SoftwareBuild Open Command Prompt(cmd) \  
   
   `set GOOS=linux` \
   `set GOARCH=amd64` \
   `go build -o mainlinux github.com/panditb/golearning` \
   `./mainlinux` \

3. MacoS  SoftwareBuild Open Command Prompt(cmd) \  
   `set GOOS=darwin` \
   `set GOARCH=amd64` \
   `go build -o mainmacos github.com/panditb/golearning` \
   `./mainmacos` \   




    
   