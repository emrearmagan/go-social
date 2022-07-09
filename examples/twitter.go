/*
twitter.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package main

import (
	"context"
	"fmt"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social/twitter"
	"log"
)

func main() {
	auther := oauth1.NewOAuth(context.TODO(), oauth.NewCredentials("xxxx", "xxxxx"), oauth1.NewToken("xxx", "xxx"))
	client := twitter.NewClient(auther)

	u, err := client.User.UserCredentials(nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("User credentials: %v \n\n", u)

	f, _ := client.Follower.FollowerIDs(nil)
	fmt.Printf("Follower IDs: %v \n\n", f)

	f2, _ := client.Follower.FollowingIDs(nil)
	fmt.Printf("Following IDs: %v \n\n", f2)
}
