package utils

import (
		//"net/http"
		"fmt"
		"io/ioutil"
		)
const ENDPOINT = "https://%s.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s"
var KEY,ok = ioutil.ReadFile("config.txt")

type SummonerProfile struct {
	Icon int
	Name string
	Level int
	RevisionDate int //when was the profile last modified. It is given as epoch milliseconds (w/e that means, need to check it out)
	ID int //Summoner ID - NOT ACCOUNT ID
	AccountID int
}

func GetSummonerByName(name string, server string) {
	fmt.Println(string(KEY))
}