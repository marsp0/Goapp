package utils

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		"errors"
		"time"
		)
const ENDPOINT_SUMMONER_BY_NAME = "https://%s.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s"
const ENDPOINT_MATCHES_BY_ID = "https://%s.api.riotgames.com/lol/match/v3/matchlists/by-account/%d/recent?api_key=%s"
const ENDPOINT_CHAMPIONS_BY_ID = "https://%s.api.riotgames.com/lol/static-data/v3/champions/%d"
var KEY,ok = ioutil.ReadFile("config.txt")

var GameModes = map[int]string{
	0 : "Custom",
	8 : "Normal 3v3",
	2 : "Normal 5v5 Blind Pick",
	14 : "Normal 5v5 Draft Pick",
	4 : "Ranked Solo 5v5",
	9 : "Ranked Flex Twisted Treeline",
	42 : "Ranked Team 5v5",
	31 : "Summoner's Rift Coop vs AI Intro",
	32 : "Summoner's Rift Coop vs AI Beginner",
	33 : "Summoner's Rift Coop vs AI Intermediate",
	52 : "Twisted Treeline Coop vs AI games",
	61 : "Team Builder",
	65 : "ARAM",
	70 : "One for All",
	72 : "Snowdown Showdown 1v1",
	73 : "Snowdown Showdown 2v2",
	75 : "Summoner's Rift 6x6 Hexakill",
	76 : "Ultra Rapid Fire",
	78 : "One for All (Mirror mode)",
	83 : "Ultra Rapid Fire AI",
	91 : "Doom Bots Rank 1",
	92 : "Doom Bots Rank 2",
	93 : "Doom Bots Rank 5",
	96 : "Ascension",
	98 : "Treeline 6x6 Hexakill",
	100 : "Butcher's Bridge",
	300 : "King Poro",
	310 : "Nemesis",
	313 : "Black Market Brawlers",
	315 : "Nexus Siege",
	317 : "Definitely Not Dominion",
	318 : "All Random URF",
	325 : "All Random Summoner's Rift",
	400 : "Normal 5v5 Draft Pick",
	420 : "Ranked Solo",
	430 : "Normal 5v5 Blind Pick",
	440 : "Ranked Flex",
	600 : "Blood Hunt Assassin",
	610 : "Dark Star"}

//Champion struct
// might add additional data given that the static endpoint provides a lot of data
type Champion struct {
	Title string `title`
	Name string `name`
	id int `id`
}


//Match struct to represent each match when we fetch the match history
type Match struct {
	Lane string `lane`
	GameId int `gameId`
	Champion int `champion`
	ChampionName string
	PlatformId string `platformId`
	Timestamp int `timestamp`
	Date time.Time
	Queue int `queue`
	Mode string
	Role string `role`
	Season int `season`
	
}

//Struct for the Profile of each summoner
type SummonerProfile struct {
	ProfileIconId int
	Name string
	SummonerLevel int
	RevisionDate int //when was the profile last modified. It is given as epoch milliseconds (w/e that means, need to check it out)
	Id int //Summoner ID - NOT ACCOUNT ID
	AccountId int
	Matches []Match
}

func (summoner *SummonerProfile) GetMatchesByID(id int, server string) error{
	// Call the end point to get the matches
	var Response, ResponseError = http.Get(fmt.Sprintf(ENDPOINT_MATCHES_BY_ID,server, id, string(KEY)))
	//use anon struct for the unmarshal function later on
	var matches = struct {Matches []Match}{}
	//we check if the call was ok
	//Need to figure out a way to handle the errors better.
	if ResponseError != nil {
		return errors.New("An error occured with the request to Rito")
	} else if Response.StatusCode != 200 {
		return errors.New("The response code was not 200")
	} else {
		//we have received 200 response and now we need to read the body.
		var ByteResponse, ByteError = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return errors.New("The Byte Reader did not finish properly")
		} else {
			
			var err = json.Unmarshal(ByteResponse,&matches)
			if err != nil {
				return errors.New("The decoding went wrong")
			}
		}
	}
	defer Response.Body.Close()
	summoner.Matches = matches.Matches
	for i := 0; i < len(summoner.Matches); i++ {
		summoner.Matches[i].Date = time.Unix(int64(summoner.Matches[i].Timestamp)/1000,0)
		summoner.Matches[i].Mode = GameModes[summoner.Matches[i].Queue]
	}
	return nil
}

func GetSummonerByName(name string, server string) (*SummonerProfile, error ) {
	//The function should return profile address and an error. We need it in case where we cannot get the profile for some reason
	var Response,ResponseError  = http.Get(fmt.Sprintf(ENDPOINT_SUMMONER_BY_NAME, server,name, string(KEY)))
	var profile = SummonerProfile{}
	// a bunch of returns, but am not currently able to 'predefine' an error variable that should hold the eventual errors and then just use 1 return at the end.
	if ResponseError != nil {
		return &profile, errors.New("Error with the GET request from RITO's API")
	} else if Response.StatusCode != 200 {
		return &profile, errors.New("The response code was not 200")
	} else {
		//
		var ByteResponse, ByteError  = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return &profile, errors.New("The ByteReader did not work properly")
		} else {
			var err = json.Unmarshal(ByteResponse, &profile)
			if err != nil {
				return &profile, errors.New("Something went wrong with the decoding")
					
			}
			//We need to find a way to cache this
			// One way could be to just call for summoner and if revision time is different then call
			var matches_err = profile.GetMatchesByID(profile.AccountId,server)
			if matches_err != nil {
				return &profile, errors.New("something went wrong with the matches")
			}
		}
	}
	for i := 0; i < len(profile.Matches); i++ {
		
	}
	defer Response.Body.Close()
	return &profile, nil
}

