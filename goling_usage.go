/*
$ export GOPATH=/Users/gyuho/src
$ cd $GOPATH/Users/gyuho/src/
$ go get github.com/golingorg/goling

// good idea to create a different GOPATH for each
$ mkdir $HOME/gopath
$ export GOPATH="$HOME/gopath"
$ export PATH="$GOPATH/bin:$PATH"

// to create a new project
$ mkdir -p $GOPATH/src/github.com/golingorg/goling
$ cd $GOPATH/src/github.com/golingorg/goling
$ git init

*/
package main

import "fmt"
import "goling.go"

func main() {
	fmt.Pritln(goling.Clean_up("Hdslfkjdsf sdf dsf 	d 		")) 
}