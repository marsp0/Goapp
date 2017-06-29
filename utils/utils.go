package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const ENDPOINT_SUMMONER_BY_NAME = "https://%s.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s"
const ENDPOINT_RECENT_MATCHES_BY_ID = "https://%s.api.riotgames.com/lol/match/v3/matchlists/by-account/%d/recent?api_key=%s"
const ENDPOINT_RANKED_BY_ID = "https://%s.api.riotgames.com/lol/match/v3/matchlists/by-account/%d?api_key=%s"
const ENDPOINT_CHAMPIONS_BY_ID = "https://%s.api.riotgames.com/lol/static-data/v3/champions/%d"
const ENDPOINT_FEATURED_GAMES = "https://%s.api.riotgames.com/lol/spectator/v3/featured-games"
const ENDPOINT_MATCH_BY_GAME_ID = "https://%s.api.riotgames.com/lol/match/v3/matches/%s?api_key=%s"

var KEY, ok = ioutil.ReadFile("config.txt")

var GameModes = map[int]string{
	0:   "Custom",
	8:   "Normal 3v3",
	2:   "Normal 5v5 Blind Pick",
	14:  "Normal 5v5 Draft Pick",
	4:   "Ranked Solo 5v5",
	9:   "Ranked Flex Twisted Treeline",
	42:  "Ranked Team 5v5",
	31:  "Summoner's Rift Coop vs AI Intro",
	32:  "Summoner's Rift Coop vs AI Beginner",
	33:  "Summoner's Rift Coop vs AI Intermediate",
	52:  "Twisted Treeline Coop vs AI games",
	61:  "Team Builder",
	65:  "ARAM",
	70:  "One for All",
	72:  "Snowdown Showdown 1v1",
	73:  "Snowdown Showdown 2v2",
	75:  "Summoner's Rift 6x6 Hexakill",
	76:  "Ultra Rapid Fire",
	78:  "One for All (Mirror mode)",
	83:  "Ultra Rapid Fire AI",
	91:  "Doom Bots Rank 1",
	92:  "Doom Bots Rank 2",
	93:  "Doom Bots Rank 5",
	96:  "Ascension",
	98:  "Treeline 6x6 Hexakill",
	100: "Butcher's Bridge",
	300: "King Poro",
	310: "Nemesis",
	313: "Black Market Brawlers",
	315: "Nexus Siege",
	317: "Definitely Not Dominion",
	318: "All Random URF",
	325: "All Random Summoner's Rift",
	400: "Normal 5v5 Draft Pick",
	420: "Ranked Solo",
	430: "Normal 5v5 Blind Pick",
	440: "Ranked Flex",
	600: "Blood Hunt Assassin",
	610: "Dark Star"}

var Champions = map[int]string{24: "Jax",
	37:  "Sona",
	18:  "Tristana",
	110: "Varus",
	114: "Fiora",
	27:  "Singed",
	223: "TahmKench",
	7:   "Leblanc",
	412: "Thresh",
	43:  "Karma",
	202: "Jhin",
	68:  "Rumble",
	77:  "Udyr",
	64:  "LeeSin",
	83:  "Yorick",
	38:  "Kassadin",
	15:  "Sivir",
	21:  "MissFortune",
	119: "Draven",
	157: "Yasuo",
	10:  "Kayle",
	35:  "Shaco",
	58:  "Renekton",
	120: "Hecarim",
	105: "Fizz",
	96:  "KogMaw",
	57:  "Maokai",
	127: "Lissandra",
	222: "Jinx",
	6:   "Urgot",
	9:   "Fiddlesticks",
	3:   "Galio",
	80:  "Pantheon",
	91:  "Talon",
	41:  "Gangplank",
	81:  "Ezreal",
	150: "Gnar",
	17:  "Teemo",
	1:   "Annie",
	82:  "Mordekaiser",
	268: "Azir",
	85:  "Kennen",
	92:  "Riven",
	31:  "Cho'Gath",
	266: "Aatrox",
	78:  "Poppy",
	163: "Taliyah",
	420: "Illaoi",
	74:  "Heimerdinger",
	12:  "Alistar",
	5:   "XinZhao",
	236: "Lucian",
	106: "Volibear",
	113: "Sejuani",
	76:  "Nidalee",
	86:  "Garen",
	89:  "Leona",
	238: "Zed",
	53:  "Blitzcrank",
	33:  "Rammus",
	161: "Vel'Koz",
	51:  "Caitlyn",
	48:  "Trundle",
	203: "Kindred",
	133: "Quinn",
	245: "Ekko",
	267: "Nami",
	50:  "Swain",
	44:  "Taric",
	134: "Syndra",
	72:  "Skarner",
	201: "Braum",
	45:  "Veigar",
	101: "Xerath",
	42:  "Corki",
	111: "Nautilus",
	103: "Ahri",
	126: "Jayce",
	122: "Darius",
	23:  "Tryndamere",
	40:  "Janna",
	60:  "Elise",
	67:  "Vayne",
	63:  "Brand",
	104: "Graves",
	16:  "Soraka",
	30:  "Karthus",
	8:   "Vladimir",
	26:  "Zilean",
	55:  "Katarina",
	102: "Shyvana",
	19:  "Warwick",
	115: "Ziggs",
	240: "Kled",
	121: "Khazix",
	2:   "Olaf",
	4:   "TwistedFate",
	20:  "Nunu",
	107: "Rengar",
	432: "Bard",
	39:  "Irelia",
	427: "Ivern",
	62:  "Wukong",
	22:  "Ashe",
	429: "Kalista",
	84:  "Akali",
	254: "Vi",
	32:  "Amumu",
	117: "Lulu",
	25:  "Morgana",
	56:  "Nocturne",
	131: "Diana",
	136: "AurelionSol",
	143: "Zyra",
	112: "Viktor",
	69:  "Cassiopeia",
	75:  "Nasus",
	29:  "Twitch",
	36:  "DrMundo",
	61:  "Orianna",
	28:  "Evelynn",
	421: "Rek'Sai",
	99:  "Lux",
	14:  "Sion",
	164: "Camille",
	11:  "MasterYi",
	13:  "Ryze",
	54:  "Malphite",
	34:  "Anivia",
	98:  "Shen",
	59:  "Jarvan IV",
	90:  "Malzahar",
	154: "Zac",
	79:  "Gragas",
	498: "Xayah",
	497: "Rakan"}

//Match struct to represent each match when we fetch the match history
type Match struct {
	Lane         string `lane`
	GameId       int    `gameId`
	Champion     int    `champion`
	ChampionName string
	PlatformId   string `platformId`
	Timestamp    int    `timestamp`
	Date         string
	Queue        int `queue`
	Mode         string
	Role         string `role`
	Season       int    `season`
}

type DetailedMatch struct {
	SeasonId              int
	QueueId               int
	GameId                int
	ParticipantIdentities []ParticipantIty
	GameVersion           string
	PlatformId            string
	GameMode              string
	MapId                 int
	GameType              string
	Teams                 []Team
	Participants          []Participant
	GameDuration          int
	GameCreation          int
	Separator int
}

type ParticipantIty struct {
	ParticipantId int
	Player        Player
}

type Player struct {
	CurrentPlatformId string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryUri   string `json:"matchHistoryUri"`
	PlatformId        string `json:"platformId"`
	CurrentAccountId  int    `json:"currentAccountId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerId        int    `json:"summonerId"`
	AccountId         int    `json:"accountId"`
}

type Team struct {
	FirstDragon          bool
	FirstInhibitor       bool
	Bans                 []TeamBans
	BaronKills           int
	FirstRiftHerald      bool
	FirstBaron           bool
	FiftHeraldKills      int
	FirstBlood           bool
	TeamId               int
	FirstTower           bool
	VilemawKills         int
	InhibitorKills       int
	TowerKills           int
	DominionVictoryScore int
	Win                  string
	DragonKills          int
}

type TeamBans struct {
	PickTurn   int
	ChampionId int
}

type Participant struct {
	Stats                     Stats
	ParticipantId             int
	Runes                     []Rune
	Timeline                  ParticipantTimeline
	TeamId                    int
	Spell2Id                  int
	Masteries                 []Mastery
	HighestAchievedSeasonTier string
	Spell1Id                  int
	ChampionId                int
}

type Stats struct {
	PhysicalDamageDealt             int
	NeutralMinionsKilledTeamJungle  int
	MagicDamageDealt                int
	TotalPlayerScore                int
	Deaths                          int
	Win                             bool
	NeutralMinionsKilledEnemyJungle int
	AltarsCaptured                  int
	LargestCriticalStrike           int
	TotalDamageDealt                int
	MagicDamageDealtToChampions     int
	VisionWardsBoughtInGame         int
	DamageDealtToObjectives         int
	LargestKillingSpree             int
	Item1                           int
	QuadraKills                     int
	TeamObjective                   int
	TotalTimeCrowdControlDealt      int
	LongestTimeSpentLiving          int
	WardsKilled                     int
	FirstTowerAssist                bool
	FirstTowerKill                  bool
	Item2                           int
	Item3                           int
	Item0                           int
	FirstBloodAssist                bool
	VisionScore                     int
	WardsPlaced                     int
	Item4                           int
	Item5                           int
	Item6                           int
	TurretKills                     int
	TripleKills                     int
	DamageSelfMitigated             int
	ChampLevel                      int
	NodeNeutralizeAssist            int
	FirstInhibitorKill              bool
	GoldEarned                      int
	MagicalDamageTaken              int
	Kills                           int
	DoubleKills                     int
	NodeCaptureAssist               int
	TrueDamageTaken                 int
	NodeNeutralize                  int
	FirstInhibitorAssist            bool
	Assists                         int
	UnrealKills                     int
	NeutralMinionsKilled            int
	ObjectivePlayerScore            int
	CombatPlayerScore               int
	DamageDealtToTurrets            int
	AltarsNeutralized               int
	PhysicalDamageDealtToChampions  int
	GoldSpent                       int
	TrueDamageDealt                 int
	TrueDamageDealtToChampions      int
	ParticipantId                   int
	PentaKills                      int
	TotalHeal                       int
	TotalMinionsKilled              int
	FirstBloodKill                  bool
	LargestMultiKill                int
	SightWardsBoughtInGame          int
	TotalDamageDealtToChampions     int
	TotalUnitsHealed                int
	InhibitorKills                  int
	TotalScoreRank                  int
	TotalDamageTaken                int
	KillingSprees                   int
	TimeCCingOthers                 int
	PhysicalDamageTaken             int
}

type Rune struct {
	RuneId int
	Rank   int
}

type ParticipantTimeline struct {
	Lane                        string
	ParticipantId               int
	CsDiffPerMinDeltas          map[string]float64
	GoldPerMinDeltas            map[string]float64
	XpDiffPerMinDeltas          map[string]float64
	CreepsPerMinDeltas          map[string]float64
	XpPerMinDeltas              map[string]float64
	Role                        string
	DamageTakenDiffPerMinDeltas map[string]float64
	DamageTakenPerMinDeltas     map[string]float64
}

type Mastery struct {
	MasteryId int
	Rank      int
}

//Struct for the Profile of each summoner
type SummonerProfile struct {
	ProfileIconId int
	Name          string
	SummonerLevel int
	RevisionDate  int //when was the profile last modified. It is given as epoch milliseconds (w/e that means, need to check it out)
	LastSeen      string
	Id            int //Summoner ID - NOT ACCOUNT ID
	AccountId     int
	Matches       []Match
	Ranked        []Match
}

func (summoner *SummonerProfile) GetMatchesByAccountID(id int, server string, endpoint string, ranked bool) (*[]Match, error) {
	// Call the end point to get the matches
	if ranked {
		var UnixTime = (time.Now().Unix() - 1296000) * 1000
		endpoint += fmt.Sprintf("&beginTime=%d", UnixTime)
	}
	var Response, ResponseError = http.Get(fmt.Sprintf(endpoint, server, id, string(KEY)))
	//use anon struct for the unmarshal function later on
	var matches = struct{ Matches []Match }{}
	//we check if the call was ok
	//Need to figure out a way to handle the errors better.
	if ResponseError != nil {
		return &[]Match{}, ResponseError
	} else {
		//we have received 200 response and now we need to read the body.
		var ByteResponse, ByteError = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return &[]Match{}, ByteError
		} else {

			var err = json.Unmarshal(ByteResponse, &matches)
			if err != nil {
				return &[]Match{}, err
			}
		}
	}
	defer Response.Body.Close()
	return &matches.Matches, nil
}

func GetSummonerByName(name string, server string) (*SummonerProfile, error) {
	//The function should return profile address and an error. We need it in case where we cannot get the profile for some reason
	var Response, ResponseError = http.Get(fmt.Sprintf(ENDPOINT_SUMMONER_BY_NAME, server, name, string(KEY)))
	var profile = SummonerProfile{}
	// a bunch of returns, but am not currently able to 'predefine' an error variable that should hold the eventual errors and then just use 1 return at the end.
	if ResponseError != nil {
		return &profile, ResponseError
	} else if Response.StatusCode != http.StatusOK {
		return &profile, errors.New("The response code was not 200")
	} else {
		//
		var ByteResponse, ByteError = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			return &profile, ByteError
		} else {
			var err = json.Unmarshal(ByteResponse, &profile)
			if err != nil {
				return &profile, err

			}
			//We need to find a way to cache this
			// One way could be to just call for summoner and if revision time is different then call
			var RecentMatches, _ = profile.GetMatchesByAccountID(profile.AccountId, server, ENDPOINT_RECENT_MATCHES_BY_ID, false)
			profile.Matches = *RecentMatches
			for i := 0; i < len(profile.Matches); i++ {
				var year, month, day = time.Unix(int64(profile.Matches[i].Timestamp)/1000, 0).Date()
				profile.Matches[i].Date = fmt.Sprintf("%02d-%02d-%d", day, month, year)
				profile.Matches[i].Mode = GameModes[profile.Matches[i].Queue]
				profile.Matches[i].ChampionName = Champions[profile.Matches[i].Champion]
			}
			var RankedMatches, _ = profile.GetMatchesByAccountID(profile.AccountId, server, ENDPOINT_RANKED_BY_ID, true)
			profile.Ranked = *RankedMatches
			for i := 0; i < len(profile.Ranked); i++ {
				var year, month, day = time.Unix(int64(profile.Ranked[i].Timestamp)/1000, 0).Date()
				profile.Ranked[i].Date = fmt.Sprintf("%02d-%02d-%d", day, month, year)
				profile.Ranked[i].Mode = GameModes[profile.Ranked[i].Queue]
				profile.Ranked[i].ChampionName = Champions[profile.Ranked[i].Champion]
			}
		}
	}
	defer Response.Body.Close()
	var year, month, day = time.Unix(int64(profile.RevisionDate)/1000, 0).Date()
	profile.LastSeen = fmt.Sprintf("%d-%02d-%02d", year, month, day)
	return &profile, nil
}

func GetMatchById(matchId string, server string) (*DetailedMatch, error) {
	var Response, err = http.Get(fmt.Sprintf(ENDPOINT_MATCH_BY_GAME_ID, server, matchId, string(KEY)))
	var Details = DetailedMatch{}

	if err != nil {
		fmt.Println(123)
		return &Details, err
	} else {
		var ByteResponse, ByteError = ioutil.ReadAll(Response.Body)
		if ByteError != nil {
			fmt.Println(124)
			return &Details, ByteError
		} else {
			var UnmarshalError = json.Unmarshal(ByteResponse, &Details)
			if UnmarshalError != nil {
				fmt.Println(UnmarshalError)
				return &Details, UnmarshalError
			} else {
				Details.Separator = 6
				return &Details, nil
			}
		}
	}
}
