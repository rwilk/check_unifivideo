set GOOS=freebsd
go build -o bin/freebsd/check_unifivideo

set GOOS=linux
go build -o bin/linux/check_unifivideo

set GOOS=windows
go build -o bin/windows/check_unifivideo.exe

set GOOS=darwin
go build -o bin/darwin/check_unifivideo

set GOOS=netbsd
go build -o bin/netbsd/check_unifivideo

set GOOS=openbsd
go build -o bin/openbsd/check_unifivideo

set GOOS=solaris
go build -o bin/solaris/check_unifivideo