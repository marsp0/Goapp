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

type SummonerProfile struct {
	ProfileIconId int
	Name string
	SummonerLevel int
	RevisionDate int //when was the profile last modified. It is given as epoch milliseconds (w/e that means, need to check it out)
	Id int //Summoner ID - NOT ACCOUNT ID
	AccountId int
}

func (summoner *SummonerProfile) GetMatchesByID(id int, server string) (*[]Match, error){
	var Response, ResponseError = http.Get(fmt.Sprintf(ENDPOINT_MATCHES_BY_ID,server, id, string(KEY)))
	var matches []Match
	if ResponseError != nil {
		return &[]Match{}, errors.New("An error occured with the request to Rito")
	} else if Response.StatusCode != 200 {
		return &[]Match{}, errors.New("The response code was not 200")
	} else {
		var ByteResponse, ByteError = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return &[]Match{}, errors.New("The Byte Reader did not finish properly")
		} else {
			var err = json.Unmarshal(ByteResponse,&matches)
			if err != nil {
				return &matches, errors.New("The decoding went wrong")
			}
		}
	}
	defer Response.Body.Close()
	fmt.Println(matches)
	return &matches, nil
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
				return &profile, errors.New("The decoding of the JSON went wrong")
			}
		}
	}
	defer Response.Body.Close()
	return &profile, nil
}

