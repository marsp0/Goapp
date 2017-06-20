package utils

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		)
const ENDPOINT = "https://%s.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s"
var KEY,ok = ioutil.ReadFile("config.txt")

type SummonerProfile struct {
	ProfileIconId int
	Name string
	SummonerLevel int
	RevisionDate int //when was the profile last modified. It is given as epoch milliseconds (w/e that means, need to check it out)
	id int //Summoner ID - NOT ACCOUNT ID
	AccountId int
}

func GetSummonerByName(name string, server string) {
	var Response,ResponseError  = http.Get(fmt.Sprintf(ENDPOINT, server,name,string(KEY)))
	fmt.Printf(ENDPOINT, server,name,string(KEY))
	if ResponseError != nil {
		fmt.Println("There was an error with the summoner requested")
	} else {
		//
		var ByteResponse, ByteError  = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			fmt.Println("There was an error with the ReadAll method")
		} else {
			var profile = SummonerProfile{}
			var err = json.Unmarshal(ByteResponse, &profile)
			if err != nil { 
				fmt.Println("we fucked")
			} else {
				fmt.Println("do the work with the vlaue")
			}
		}
	}
	defer Response.Body.Close()
}