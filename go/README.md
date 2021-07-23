# Tennis in [Go](http://golang.org/)

This application enables to calculate the score for a tennis game. 
There are 3 different implementations added after considering different use cases.

## Installation for newer versions of Go

Newer versions support go modules. In that case the subdirectory 'tenniskata' is a go module and 
you can just use it as is, by checking out the code in whichever directory you want in your local.

## Installation for older versions of Go

Assuming you have a proper [workspace](http://golang.org/doc/code.html#Workspaces) set up, run
```
go get github.com/jeffy-mathew/waylift-refactor/go/tenniskata
```

## Running Tests
From the project root directory on the command line, 
enter ```cd go/tenniskata``` and run
```
go test ./... -cover
```
