package dbaccess

import (
	"log"
	"strconv"

	"github.com/Mtbcooler/outrun/enums"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/obj"
)

// GetHighScores returns the list of high scores, your own entry if applicable, and an error if one is thrown
func GetHighScores(mode, lbtype, offset, limit int64, ownid string) ([]obj.LeaderboardEntry, interface{}, error) {
	CheckIfDBSet()
	leagueColumn := "ranking_league"
	if mode == 1 {
		leagueColumn = "quick_ranking_league"
	}
	columnToSortBy := ""
	switch lbtype {
	case 0: // Friends high score?
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 1: // Friends total score?
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 2: // World high score
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 3: // World total score
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 4: // Runners' League high score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 5: // Runners' League total score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 6: // Historical high score
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	case 7: // Historical total score
		// TODO: Define high_total_score column
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	default:
		log.Printf("[WARN] Unknown leaderboard type %v", lbtype)
		columnToSortBy = "high_score"
	}
	var myEntry interface{}
	leaderboardEntries := []obj.LeaderboardEntry{}
	rows, err := db.Query("SELECT id, " + columnToSortBy + ", " + leagueColumn + ", rank, mainchara_id, subchara_id, mainchao_id, subchao_id FROM `" + consts.DBMySQLTablePlayerStates + "` ORDER BY " + columnToSortBy + " DESC LIMIT " + strconv.Itoa(int(limit)) + " OFFSET " + strconv.Itoa(int(offset)))
	if err != nil {
		return []obj.LeaderboardEntry{}, nil, err
	}
	var uid, username, charasjson, chaojson string
	var highscore, league, rank, mainchara, subchara, mainchao, subchao, lastlogin int64
	var currentEntry obj.LeaderboardEntry
	index := offset
	for rows.Next() {
		err = rows.Scan(&uid, &highscore, &league, &rank, &mainchara, &subchara, &mainchao, &subchao)
		if err != nil {
			rows.Close()
			return []obj.LeaderboardEntry{}, nil, err
		}
		err = db.QueryRow("SELECT username, last_login, characters, chao FROM `"+consts.DBMySQLTableCorePlayerInfo+"` WHERE id = ?", uid).Scan(&username, &lastlogin, &charasjson, &chaojson)
		if err != nil {
			rows.Close()
			return []obj.LeaderboardEntry{}, nil, err
		}
		currentEntry = obj.NewLeaderboardEntry(
			uid,
			username,
			"",
			int64(index+1),
			0,
			highscore,
			0,
			0,
			0,
			rank,
			lastlogin,
			mainchara,
			0,
			subchara,
			0,
			mainchao,
			0,
			subchao,
			0,
			enums.LangEnglish,
			league,
			highscore,
			0,
		)
		if uid == ownid {
			myEntry = currentEntry
		}
		leaderboardEntries = append(leaderboardEntries, currentEntry)
		index++
	}
	err = rows.Err()
	rows.Close()
	return leaderboardEntries, myEntry, err
}

func GetOwnLeaderboardEntry(mode, lbtype int64, ownid string) (interface{}, error) {
	CheckIfDBSet()
	leagueColumn := "ranking_league"
	if mode == 1 {
		leagueColumn = "quick_ranking_league"
	}
	columnToSortBy := ""
	switch lbtype {
	case 0: // Friends high score?
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 1: // Friends total score?
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 2: // World high score
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 3: // World total score
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 4: // Runners' League high score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 5: // Runners' League total score
		log.Println("[WARN] Please use GetLeagueHighScores() for getting Runners' League high scores!")
		// TODO: Define total_score column
		if mode == 1 {
			columnToSortBy = "quick_league_high_score"
		} else {
			columnToSortBy = "league_high_score"
		}
	case 6: // Historical high score
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	case 7: // Historical total score
		// TODO: Define high_total_score column
		if mode == 1 {
			columnToSortBy = "quick_high_score"
		} else {
			columnToSortBy = "high_score"
		}
	default:
		log.Printf("[WARN] Unknown leaderboard type %v", lbtype)
		columnToSortBy = "high_score"
	}
	var myEntry interface{}
	rows, err := db.Query("SELECT id, " + columnToSortBy + ", " + leagueColumn + ", rank, mainchara_id, subchara_id, mainchao_id, subchao_id FROM `" + consts.DBMySQLTablePlayerStates + "` ORDER BY " + columnToSortBy + " DESC")
	if err != nil {
		return nil, err
	}
	var uid, username, charasjson, chaojson string
	var highscore, league, rank, mainchara, subchara, mainchao, subchao, lastlogin int64
	index := 0
	for rows.Next() {
		err = rows.Scan(&uid, &highscore, &league, &rank, &mainchara, &subchara, &mainchao, &subchao)
		if err != nil {
			rows.Close()
			return nil, err
		}
		if uid == ownid {
			err = db.QueryRow("SELECT username, last_login, characters, chao FROM `"+consts.DBMySQLTableCorePlayerInfo+"` WHERE id = ?", uid).Scan(&username, &lastlogin, &charasjson, &chaojson)
			if err != nil {
				rows.Close()
				return nil, err
			}
			myEntry = obj.NewLeaderboardEntry(
				uid,
				username,
				"",
				int64(index+1),
				0,
				highscore,
				0,
				0,
				0,
				rank,
				lastlogin,
				mainchara,
				0,
				subchara,
				0,
				mainchao,
				0,
				subchao,
				0,
				enums.LangEnglish,
				league,
				highscore,
				0,
			)
		}
	}
	err = rows.Err()
	rows.Close()
	return myEntry, err
}

func GetNumOfPlayers() (int64, error) {
	CheckIfDBSet()
	playercount := int64(0)
	err := db.QueryRow("SELECT COUNT(*) FROM `" + consts.DBMySQLTablePlayerStates + "`").Scan(&playercount)
	if err != nil {
		return -1, err
	}
	return playercount, nil
}
