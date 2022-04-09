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
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social/twitter"
	"log"
)

var ConfigPath string

func main() {

	// pass config file

	flag.StringVar(&ConfigPath, "c", "./config/config_example.json", "Specified the config file for running server. Default is the \"config_example\" in the config directory.")
	flag.Parse()

	//load config
	accounts, err := config.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	credentials := oauth.NewCredentials(accounts.Twitter.ConsumerKey, accounts.Twitter.ConsumerSecret)
	token := oauth1.NewToken(accounts.Twitter.Token, accounts.Twitter.TokenSecret)
	conf := oauth1.NewOAuth(context.TODO(), credentials, token)
	client := twitter.NewClient(conf)

	u, _ := client.User.UserCredentials(nil)
	fmt.Printf("Usercredentials: %v \n\n", u)

	f, _ := client.Follower.FollowerIDs(nil)
	fmt.Printf("Follower IDs: %v \n\n", f)

	f2, _ := client.Follower.FollowingIDs(nil)
	fmt.Printf("Following IDs: %v \n\n", f2)
}
