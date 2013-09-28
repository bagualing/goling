/*
$ export GOPATH=/Users/gyuho/src
$ cd $GOPATH/Users/gyuho/src/github.com/
$ go get github.com/golingorg/goling

*/
package main

import "fmt"
import "github.com/golingorg/goling"

func main() {
	fmt.Pritln(goling.Clean_up("Hdslfkjdsf sdf dsf 	d 		")) 
}