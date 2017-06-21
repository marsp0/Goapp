package utils

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		"errors"
		)
const ENDPOINT_SUMMONER_BY_NAME = "https://%s.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s"
const ENDPOINT_MATCHES_BY_ID = "https://%s.api.riotgames.com/lol/match/v3/matchlists/by-account/%d/recent?api_key=%s"
var KEY,ok = ioutil.ReadFile("config.txt")

//struct to help us decode the json
type Matches struct {
	Matches []Match
}

//Match struct to represent each match when we fetch the match history
type Match struct {
	Lane string
	GameId int
	Champion int
	PlatformId string
	Timestamp int
	Queue int
	Role string
	Season int
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
	var Response, ResponseError = http.Get(fmt.Sprintf(ENDPOINT_MATCHES_BY_ID,server, id, string(KEY)))
	var matches Matches
	if ResponseError != nil {
		return errors.New("An error occured with the request to Rito")
	} else if Response.StatusCode != 200 {
		return errors.New("The response code was not 200")
	} else {
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
			var matches_err = profile.GetMatchesByID(profile.AccountId,server)
			if matches_err != nil {
				return &profile, errors.New("something went wrong with the matches")
			}
		}
	}
	defer Response.Body.Close()
	return &profile, nil
}

