# My gracenote assignment

# Environment

## Dev
* Windows 10 
* Golang 1.8.1
* intellij community edition

## Test
* Windows 10

Although this is tested on windows, there should be no problem running it on other platforms.

# Getting Started
make sure you have [go](https://golang.org/dl/)
```
go get github.com/thinkong/gnassign/runnable
go install github.com/thinkong/gnassign/runnable
```

This should install a `runnable.exe` to your %GOPATH%\bin

copy your blacklist.conf to %GOPATH%\bin

run with `runnable.exe`

open chrome or any http client and go to 
```
http://localhost:8080/download_image?url=gracenote.com/abc.png
```

