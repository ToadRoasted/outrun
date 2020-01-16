package netobj

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/jinzhu/now"

	"github.com/Mtbcooler/outrun/config/gameconf"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/obj/constobjs"
)

/*
Notes:
  - rankingLeague gets converted to enums.RankingLeague*
    - The internal icon name for the league icons is "ui_ranking_league_icon_s_" + value.lower(), where value is a string of the rank (S, A, et cetera.)
*/

type PlayerState struct {
	Items                  []obj.Item `json:"items" db:"items,"`                 // items owned
	EquippedItemIDs        []string   `json:"equipItemList" db:"equipped_items"` // default is list of 3 "-1"s. look to be item ids
	MainCharaID            string     `json:"mainCharaID" db:"mainchara_id"`
	SubCharaID             string     `json:"subCharaID" db:"subchara_id"`
	MainChaoID             string     `json:"mainChaoID" db:"mainchao_id"`
	SubChaoID              string     `json:"subChaoID" db:"subchao_id"`
	NumRings               int64      `json:"numRings,string" db:"num_rings"`                      // number of rings
	NumBuyRings            int64      `json:"numBuyRings,string" db:"num_buy_rings"`               // number of rings purchased
	NumRedRings            int64      `json:"numRedRings,string" db:"num_red_rings"`               // number of red rings
	NumBuyRedRings         int64      `json:"numBuyRedRings,string" db:"num_buy_red_rings"`        // number of red rings purchased
	Energy                 int64      `json:"energy,string" db:"energy"`                           // energy/'lives'
	EnergyBuy              int64      `json:"energyBuy,string" db:"energy_buy"`                    // ?
	EnergyRenewsAt         int64      `json:"energyRenewsAt" db:"energy_renews_at"`                // does 0 mean it is instant?
	MumMessages            int64      `json:"mumMessages" db:"num_messages"`                       // number of unread messages
	RankingLeague          int64      `json:"rankingLeague,string" db:"ranking_league"`            // 'league index'
	QuickRankingLeague     int64      `json:"quickRankingLeague,string" db:"quick_ranking_league"` // same as above, but for timed mode
	NumRouletteTicket      int64      `json:"numRouletteTicket,string" db:"num_roulette_ticket"`
	NumChaoRouletteTicket  int64      `json:"numChaoRouletteTicket" db:"num_chao_roulette_ticket"` // This isn't a requirement from the game for PlayerState, but is useful to have here
	ChaoEggs               int64      `json:"chaoEggs" db:"chao_eggs"`                             // Same as above
	HighScore              int64      `json:"totalHighScore,string" db:"high_score"`
	TimedHighScore         int64      `json:"quickTotalHighScore,string" db:"quick_high_score"`
	TotalDistance          int64      `json:"totalDistance,string" db:"total_distance"`
	HighDistance           int64      `json:"maximumDistance,string" db:"best_distance"` // high distance in one go?
	DailyMissionID         int64      `json:"dailyMissionId,string" db:"daily_mission_id"`
	DailyMissionEndTime    int64      `json:"dailyMissionEndTime" db:"daily_mission_end_time"` // 11:59 pm of current day
	DailyChallengeValue    int64      `json:"dailyChallengeValue" db:"daily_challenge_value"`  // internally listed as ProgressStatus... Current day of the challenge?
	DailyChallengeComplete int64      `json:"dailyChallengeComplete" db:"daily_challenge_complete"`
	NumDailyChallenge      int64      `json:"numDailyChalCont" db:"num_daily_chal_cont"`
	NumPlaying             int64      `json:"numPlaying,string" db:"num_plays"` // ?
	Animals                int64      `json:"numAnimals,string" db:"num_animals"`
	Rank                   int64      `json:"numRank,string" db:"rank"`
	DailyChalCatNum        int64      `json:"ORN_dailyChalCatNum,string" db:"dm_cat"`
	DailyChalSetNum        int64      `json:"ORN_dailyChalSetNum,string" db:"dm_set"`
	DailyChalPosNum        int64      `json:"ORN_dailyChalPosNum,string" db:"dm_pos"`
	NextNumDailyChallenge  int64      `json:"ORN_nextNumDailyChalCont" db:"dm_nextcont"`
	LeagueHighScore        int64      `json:"leagueHighScore" db:"league_high_score"`
	QuickLeagueHighScore   int64      `json:"quickLeagueHighScore" db:"quick_league_high_score"`
	LeagueResetTime        int64      `json:"leagueResetTime" db:"league_reset_time"`
}

type SqlCompatiblePlayerState struct {
	ID                     int64  `db:"id"`
	Items                  []byte `json:"items" db:"items"`                  // items owned
	EquippedItemIDs        []byte `json:"equipItemList" db:"equipped_items"` // default is list of 3 "-1"s. look to be item ids
	MainCharaID            string `json:"mainCharaID" db:"mainchara_id"`
	SubCharaID             string `json:"subCharaID" db:"subchara_id"`
	MainChaoID             string `json:"mainChaoID" db:"mainchao_id"`
	SubChaoID              string `json:"subChaoID" db:"subchao_id"`
	NumRings               int64  `json:"numRings,string" db:"num_rings"`                      // number of rings
	NumBuyRings            int64  `json:"numBuyRings,string" db:"num_buy_rings"`               // number of rings purchased
	NumRedRings            int64  `json:"numRedRings,string" db:"num_red_rings"`               // number of red rings
	NumBuyRedRings         int64  `json:"numBuyRedRings,string" db:"num_buy_red_rings"`        // number of red rings purchased
	Energy                 int64  `json:"energy,string" db:"energy"`                           // energy/'lives'
	EnergyBuy              int64  `json:"energyBuy,string" db:"energy_buy"`                    // ?
	EnergyRenewsAt         int64  `json:"energyRenewsAt" db:"energy_renews_at"`                // does 0 mean it is instant?
	MumMessages            int64  `json:"mumMessages" db:"num_messages"`                       // number of unread messages
	RankingLeague          int64  `json:"rankingLeague,string" db:"ranking_league"`            // 'league index'
	QuickRankingLeague     int64  `json:"quickRankingLeague,string" db:"quick_ranking_league"` // same as above, but for timed mode
	NumRouletteTicket      int64  `json:"numRouletteTicket,string" db:"num_roulette_ticket"`
	NumChaoRouletteTicket  int64  `json:"numChaoRouletteTicket" db:"num_chao_roulette_ticket"` // This isn't a requirement from the game for PlayerState, but is useful to have here
	ChaoEggs               int64  `json:"chaoEggs" db:"chao_eggs"`                             // Same as above
	HighScore              int64  `json:"totalHighScore,string" db:"high_score"`
	TimedHighScore         int64  `json:"quickTotalHighScore,string" db:"quick_high_score"`
	TotalDistance          int64  `json:"totalDistance,string" db:"total_distance"`
	HighDistance           int64  `json:"maximumDistance,string" db:"best_distance"` // high distance in one go?
	DailyMissionID         int64  `json:"dailyMissionId,string" db:"daily_mission_id"`
	DailyMissionEndTime    int64  `json:"dailyMissionEndTime" db:"daily_mission_end_time"` // 11:59 pm of current day
	DailyChallengeValue    int64  `json:"dailyChallengeValue" db:"daily_challenge_value"`  // internally listed as ProgressStatus... Current day of the challenge?
	DailyChallengeComplete int64  `json:"dailyChallengeComplete" db:"daily_challenge_complete"`
	NumDailyChallenge      int64  `json:"numDailyChalCont" db:"num_daily_chal_cont"`
	NumPlaying             int64  `json:"numPlaying,string" db:"num_plays"` // ?
	Animals                int64  `json:"numAnimals,string" db:"num_animals"`
	Rank                   int64  `json:"numRank,string" db:"rank"`
	DailyChalCatNum        int64  `json:"ORN_dailyChalCatNum,string" db:"dm_cat"`
	DailyChalSetNum        int64  `json:"ORN_dailyChalSetNum,string" db:"dm_set"`
	DailyChalPosNum        int64  `json:"ORN_dailyChalPosNum,string" db:"dm_pos"`
	NextNumDailyChallenge  int64  `json:"ORN_nextNumDailyChalCont" db:"dm_nextcont"`
	LeagueHighScore        int64  `json:"leagueHighScore" db:"league_high_score"`
	QuickLeagueHighScore   int64  `json:"quickLeagueHighScore" db:"quick_league_high_score"`
	LeagueResetTime        int64  `json:"leagueResetTime" db:"league_reset_time"`
}

func PlayerStateToSQLCompatiblePlayerState(ps PlayerState) SqlCompatiblePlayerState {
	items, _ := json.Marshal(ps.Items)
	equippeditems, _ := json.Marshal(ps.EquippedItemIDs)
	return SqlCompatiblePlayerState{
		0,
		items,
		equippeditems,
		ps.MainCharaID,
		ps.SubCharaID,
		ps.MainChaoID,
		ps.SubChaoID,
		ps.NumRings,
		ps.NumBuyRings,
		ps.NumRedRings,
		ps.NumBuyRedRings,
		ps.Energy,
		ps.EnergyBuy,
		ps.EnergyRenewsAt,
		ps.MumMessages,
		ps.RankingLeague,
		ps.QuickRankingLeague,
		ps.NumRouletteTicket,
		ps.NumChaoRouletteTicket,
		ps.ChaoEggs,
		ps.HighScore,
		ps.TimedHighScore,
		ps.TotalDistance,
		ps.HighDistance,
		ps.DailyMissionID,
		ps.DailyMissionEndTime,
		ps.DailyChallengeValue,
		ps.DailyChallengeComplete,
		ps.NumDailyChallenge,
		ps.NumPlaying,
		ps.Animals,
		ps.Rank,
		ps.DailyChalCatNum,
		ps.DailyChalSetNum,
		ps.DailyChalPosNum,
		ps.NextNumDailyChallenge,
		ps.LeagueHighScore,
		ps.QuickLeagueHighScore,
		ps.LeagueResetTime,
	}
}

func SQLCompatiblePlayerStateToPlayerState(ps SqlCompatiblePlayerState) PlayerState {
	var items []obj.Item
	json.Unmarshal(ps.Items, &items)
	var equippeditems []string
	json.Unmarshal(ps.EquippedItemIDs, &equippeditems)
	return PlayerState{
		items,
		equippeditems,
		ps.MainCharaID,
		ps.SubCharaID,
		ps.MainChaoID,
		ps.SubChaoID,
		ps.NumRings,
		ps.NumBuyRings,
		ps.NumRedRings,
		ps.NumBuyRedRings,
		ps.Energy,
		ps.EnergyBuy,
		ps.EnergyRenewsAt,
		ps.MumMessages,
		ps.RankingLeague,
		ps.QuickRankingLeague,
		ps.NumRouletteTicket,
		ps.NumChaoRouletteTicket,
		ps.ChaoEggs,
		ps.HighScore,
		ps.TimedHighScore,
		ps.TotalDistance,
		ps.HighDistance,
		ps.DailyMissionID,
		ps.DailyMissionEndTime,
		ps.DailyChallengeValue,
		ps.DailyChallengeComplete,
		ps.NumDailyChallenge,
		ps.NumPlaying,
		ps.Animals,
		ps.Rank,
		ps.DailyChalCatNum,
		ps.DailyChalSetNum,
		ps.DailyChalPosNum,
		ps.NextNumDailyChallenge,
		ps.LeagueHighScore,
		ps.QuickLeagueHighScore,
		ps.LeagueResetTime,
	}
}

var ChaoIDs = []string{"400000", "400001", "400002", "400003", "400004", "400005", "400006", "400007", "400008", "400009", "400010", "400011", "400012", "400013", "400014", "400015", "400016", "400017", "400018", "400019", "400020", "400021", "400022", "400023", "400024", "400025", "401000", "401001", "401002", "401003", "401004", "401005", "401006", "401007", "401008", "401009", "401010", "401011", "401012", "401013", "401014", "401015", "401016", "401017", "401018", "401019", "401020", "401021", "401022", "401023", "401024", "401025", "401026", "401027", "401028", "401029", "401030", "401031", "401032", "401033", "401034", "401035", "401036", "401037", "401038", "401039", "401040", "401041", "401042", "401043", "401044", "401045", "401046", "401047", "402000", "402001", "402002", "402003", "402004", "402005", "402006", "402007", "402008", "402009", "402010", "402011", "402012", "402013", "402014", "402015", "402016", "402017", "402018", "402019", "402020", "402021", "402022", "402023", "402024", "402025", "402026", "402027", "402028", "402029", "402030", "402031", "402032", "402033", "402034"}

func DefaultPlayerState() PlayerState {
	// TODO: establish as constants
	items := constobjs.DefaultPlayerStateItems
	equippedItemIDs := []string{"-1", "-1", "-1"}
	//mainCharaID := enums.CTStrSonic
	mainCharaID := gameconf.CFile.DefaultMainCharacter
	//subCharaID := enums.CTStrTails
	subCharaID := gameconf.CFile.DefaultSubCharacter
	//mainChaoID := ChaoIDs[0]
	mainChaoID := gameconf.CFile.DefaultMainChao
	//subChaoID := ChaoIDs[5] // changed from [1]...
	subChaoID := gameconf.CFile.DefaultSubChao
	numRings := int64(gameconf.CFile.StartingRings)
	//numBuyRings := int64(1)
	numBuyRings := int64(0)
	numRedRings := int64(gameconf.CFile.StartingRedRings)
	//numBuyRedRings := int64(7)
	numBuyRedRings := int64(0)
	energy := int64(gameconf.CFile.StartingEnergy)
	energyBuy := int64(0)
	energyRenewsAt := time.Now().Unix() + 600 // in ten minutes
	mumMessages := int64(0)
	rankingLeague := int64(enums.RankingLeagueF_M)
	quickRankingLeague := int64(enums.RankingLeagueF_M)
	numRouletteTicket := int64(3)
	numChaoRouletteTicket := int64(7)
	//chaoEggs := int64(11)
	chaoEggs := int64(0)
	highScore := int64(0)
	timedHighScore := int64(0)
	totalDistance := int64(0)
	highDistance := int64(0)
	dcCatNum := int64(rand.Intn(5))
	dcSetNum := int64(0)
	dcPosNum := int64(1 + rand.Intn(2))
	dailyMissionID := int64((dcCatNum * 33) + (dcSetNum * 3) + dcPosNum)
	dailyMissionEndTime := now.EndOfDay().UTC().Unix() // TODO: should this be in UTC, or local time?
	dailyChallengeValue := int64(0)
	dailyChallengeComplete := int64(0)
	numDailyChallenge := int64(0)
	numPlayer := int64(0)
	animals := int64(0)
	rank := int64(0)
	nextNumDailyChallenge := int64(1)
	leagueHighScore := int64(0)
	quickLeagueHighScore := int64(0)
	leagueResetTime := now.EndOfWeek().UTC().Unix()
	return PlayerState{
		items,
		equippedItemIDs,
		mainCharaID,
		subCharaID,
		mainChaoID,
		subChaoID,
		numRings,
		numBuyRings,
		numRedRings,
		numBuyRedRings,
		energy,
		energyBuy,
		energyRenewsAt,
		mumMessages,
		rankingLeague,
		quickRankingLeague,
		numRouletteTicket,
		numChaoRouletteTicket,
		chaoEggs,
		highScore,
		timedHighScore,
		totalDistance,
		highDistance,
		dailyMissionID,
		dailyMissionEndTime,
		dailyChallengeValue,
		dailyChallengeComplete,
		numDailyChallenge,
		numPlayer,
		animals,
		rank,
		dcCatNum,
		dcSetNum,
		dcPosNum,
		nextNumDailyChallenge,
		leagueHighScore,
		quickLeagueHighScore,
		leagueResetTime,
	}
}
