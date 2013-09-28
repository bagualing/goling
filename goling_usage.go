/*
http://stackoverflow.com/questions/10130341/go-go-get-go-install-local-packages-and-version-control

$ export GOPATH=/Users/gyuho/src
$ cd $GOPATH/Users/gyuho/src/
$ go get github.com/golingorg/goling
*/
package main

import "fmt"
import "goling.go"

func main() {
	fmt.Pritln(goling.Clean_up("Hdslfkjdsf sdf dsf 	d 		")) 
}