package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	// fmt.Println()
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	auth := authenticationInfo{
		username: "erickkramer",
		password: "1234",
	}
	fmt.Println(auth.getBasicAuth())
}
