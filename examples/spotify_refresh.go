/*
spotify_refresh.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/emrearmagan/go-social/config"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social/spotify"
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

	conf := oauth2.NewOAuth(context.TODO(), &accounts.Spotify.Credentials, &accounts.Spotify.Token)
	client := spotify.NewClient(conf)

	newToken, _ := client.Account.RefreshToken()
	fmt.Printf("Refreshed Token: %v \n\n", newToken)

	u, _ := client.User.UserCredentials()
	fmt.Printf("User credentials: %v \n\n", u)

	p, _ := client.Playlist.UserPlaylists(&spotify.UserPlaylistParams{
		Limit: 10,
	})
	fmt.Printf("User Playlist: %v \n\n", p)
}
