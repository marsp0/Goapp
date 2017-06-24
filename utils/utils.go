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

var Champions = map[int]string{ 24 : "Jax" ,
									37 :"Sona",
									18 :"Tristana",
									110: "Varus",
									114:"Fiora",
									27 : "Singed",
									223: "TahmKench",
									7  : "Leblanc",
									412: "Thresh",
									43 : "Karma",
									202: "Jhin",
									68 : "Rumble",
									77 : "Udyr",
									64 : "LeeSin",
									83 : "Yorick",
									38 : "Kassadin",
									15 : "Sivir",
									21 : "MissFortune",
									119: "Draven",
									157: "Yasuo",
									10 : "Kayle",
									35 : "Shaco",
									58 :"Renekton",
									120:"Hecarim",
									105:"Fizz",
									96 :"KogMaw",
									57 :"Maokai",
									127:"Lissandra",
									222:"Jinx",
									6  : "Urgot",
									9  : "Fiddlesticks",
									3  : "Galio",
									80 :"Pantheon",
									91 : "Talon",
									41 : "Gangplank",
									81 :"Ezreal",
									150:"Gnar",
									17 :"Teemo",
									1  :"Annie",
									82 : "Mordekaiser",
									268: "Azir",
									85 : "Kennen",
									92 : "Riven",
									31 : "Cho'Gath",
									266: "Aatrox",
									78 : "Poppy",
									163: "Taliyah",
									420: "Illaoi",
									74 : "Heimerdinger",
									12 : "Alistar",
									5  : "XinZhao",
									236: "Lucian",
									106: "Volibear",
									113: "Sejuani",
									76 : "Nidalee",
									86 : "Garen",
									89 : "Leona",
									238: "Zed",
									53 : "Blitzcrank",
									33 : "Rammus",
									161: "Vel'Koz",
									51 : "Caitlyn",
									48 : "Trundle",
									203: "Kindred",
									133: "Quinn",
									245: "Ekko",
									267: "Nami",
									50 : "Swain",
									44 :"Taric",
									134:"Syndra",
									72 : "Skarner",
									201:"Braum",
									45 : "Veigar",
									101: "Xerath",
									42 : "Corki",
									111:"Nautilus",
									103: "Ahri",
									126: "Jayce",
									122: "Darius",
									23 : "Tryndamere",
									40 :"Janna",
									60 :"Elise",
									67 : "Vayne",
									63 : "Brand",
									104:"Graves",
									16 :"Soraka",
									30 : "Karthus",
									8  : "Vladimir",
									26 : "Zilean",
									55 : "Katarina",
									102: "Shyvana",
									19 : "Warwick",
									115:"Ziggs",
									240:"Kled",
									121:"Khazix",
									2  : "Olaf",
									4  :"TwistedFate",
									20 :"Nunu",
									107:"Rengar",
									432:"Bard",
									39 :"Irelia",
									427:"Ivern",
									62 : "Wukong",
									22 :"Ashe",
									429:"Kalista",
									84 : "Akali",
									254:"Vi",
									32 : "Amumu",
									117:"Lulu",
									25 :"Morgana",
									56 : "Nocturne",
									131: "Diana",
									136: "AurelionSol",
									143:"Zyra",
									112: "Viktor",
									69 : "Cassiopeia",
									75 : "Nasus",
									29 : "Twitch",
									36 : "DrMundo",
									61 : "Orianna",
									28 : "Evelynn",
									421: "Rek'Sai",
									99 : "Lux",
									14 : "Sion",
									164: "Camille",
									11 : "MasterYi",
									13 : "Ryze",
									54 : "Malphite",
									34 : "Anivia",
									98 : "Shen",
									59 : "Jarvan IV",
									90 : "Malzahar",
									154: "Zac",
									79 : "Gragas"}

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
	Date string
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
	LastSeen string
	Id int //Summoner ID - NOT ACCOUNT ID
	AccountId int
	Matches []Match
}

func (summoner *SummonerProfile) GetMatchesByID(id int, server string) error {
	// Call the end point to get the matches
	var Response, ResponseError = http.Get(fmt.Sprintf(ENDPOINT_MATCHES_BY_ID,server, id, string(KEY)))
	//use anon struct for the unmarshal function later on
	var matches = struct {Matches []Match}{}
	//we check if the call was ok
	//Need to figure out a way to handle the errors better.
	if ResponseError != nil {
		return errors.New(ResponseError.Error())
	} else {
		//we have received 200 response and now we need to read the body.
		var ByteResponse, ByteError = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return errors.New(ByteError.Error())
		} else {
			
			var err = json.Unmarshal(ByteResponse,&matches)
			if err != nil {
				return errors.New(err.Error())
			}
		}
	}
	defer Response.Body.Close()
	summoner.Matches = matches.Matches
	for i := 0; i < len(summoner.Matches); i++ {
		var year,month,day = time.Unix(int64(summoner.Matches[i].Timestamp)/1000,0).Date()
		summoner.Matches[i].Date = fmt.Sprintf("%d-%02d-%02d", year,month,day)
		summoner.Matches[i].Mode = GameModes[summoner.Matches[i].Queue]
		summoner.Matches[i].ChampionName = Champions[summoner.Matches[i].Champion]
	}
	return nil
}

func GetSummonerByName(name string, server string) (*SummonerProfile, error ) {
	//The function should return profile address and an error. We need it in case where we cannot get the profile for some reason
	var Response,ResponseError  = http.Get(fmt.Sprintf(ENDPOINT_SUMMONER_BY_NAME, server,name, string(KEY)))
	var profile = SummonerProfile{}
	// a bunch of returns, but am not currently able to 'predefine' an error variable that should hold the eventual errors and then just use 1 return at the end.
	if ResponseError != nil {
		return &profile, errors.New(ResponseError.Error())
	} else if Response.StatusCode != http.StatusOK {
		return &profile, errors.New("The response code was not 200")
	} else {
		//
		var ByteResponse, ByteError  = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return &profile, errors.New(ByteError.Error())
		} else {
			var err = json.Unmarshal(ByteResponse, &profile)
			if err != nil {
				return &profile, errors.New(err.Error())
					
			}
			//We need to find a way to cache this
			// One way could be to just call for summoner and if revision time is different then call
			var matches_err = profile.GetMatchesByID(profile.AccountId,server)
			if matches_err != nil {
				fmt.Println("dsadsad")
				return &profile, nil
			}
		}
	}
	defer Response.Body.Close()
	var year, month, day = time.Unix(int64(profile.RevisionDate)/1000,0).Date()
	profile.LastSeen = fmt.Sprintf("%d-%02d-%02d",year,month,day)
	return &profile, nil
}

