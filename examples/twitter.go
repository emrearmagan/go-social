/*
twitter.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/emrearmagan/go-social/config"
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social/twitter"
	"log"
)

func main() {
	var ConfigPath string

	// pass config file
	flag.StringVar(&ConfigPath, "c", "./config/config_example.json", "Specified the config file for account credentials. Default is the \"config_example\" in the config directory.")
	flag.Parse()

	//load config
	accounts, err := config.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	conf := oauth1.NewOAuth(context.TODO(), &accounts.Twitter.Credentials, &accounts.Twitter.Token)
	client := twitter.NewClient(conf)

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
