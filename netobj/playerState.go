package netobj

import (
	"math/rand"
	"time"

	"github.com/jinzhu/now"

	"github.com/ToadRoasted/outrun/config/gameconf"
	"github.com/ToadRoasted/outrun/enums"
	"github.com/ToadRoasted/outrun/obj"
	"github.com/ToadRoasted/outrun/obj/constobjs"
)

/*
Notes:
  - rankingLeague gets converted to enums.RankingLeague*
    - The internal icon name for the league icons is "ui_ranking_league_icon_s_" + value.lower(), where value is a string of the rank (S, A, et cetera.)
*/

type PlayerState struct {
	Items                  []obj.Item `json:"items"`         // items owned
	EquippedItemIDs        []string   `json:"equipItemList"` // default is list of 3 "-1"s. look to be item ids
	MainCharaID            string     `json:"mainCharaID"`
	SubCharaID             string     `json:"subCharaID"`
	MainChaoID             string     `json:"mainChaoID"`
	SubChaoID              string     `json:"subChaoID"`
	NumRings               int64      `json:"numRings,string"`           // number of rings
	NumBuyRings            int64      `json:"numBuyRings,string"`        // number of rings purchased
	NumRedRings            int64      `json:"numRedRings,string"`        // number of red rings
	NumBuyRedRings         int64      `json:"numBuyRedRings,string"`     // number of red rings purchased
	Energy                 int64      `json:"energy,string"`             // energy/'lives'
	EnergyBuy              int64      `json:"energyBuy,string"`          // ?
	EnergyRenewsAt         int64      `json:"energyRenewsAt"`            // does 0 mean it is instant?
	MumMessages            int64      `json:"mumMessages"`               // number of unread messages
	RankingLeague          int64      `json:"rankingLeague,string"`      // 'league index'
	QuickRankingLeague     int64      `json:"quickRankingLeague,string"` // same as above, but for timed mode
	NumRouletteTicket      int64      `json:"numRouletteTicket,string"`
	NumChaoRouletteTicket  int64      `json:"numChaoRouletteTicket"` // This isn't a requirement from the game for PlayerState, but is useful to have here
	ChaoEggs               int64      `json:"chaoEggs"`              // Same as above
	HighScore              int64      `json:"totalHighScore,string"`
	TimedHighScore         int64      `json:"quickTotalHighScore,string"`
	TotalDistance          int64      `json:"totalDistance,string"`
	HighDistance           int64      `json:"maximumDistance,string"` // high distance in one go?
	DailyMissionID         int64      `json:"dailyMissionId,string"`
	DailyMissionEndTime    int64      `json:"dailyMissionEndTime"` // 11:59 pm of current day
	DailyChallengeValue    int64      `json:"dailyChallengeValue"` // internally listed as ProgressStatus... Current day of the challenge?
	DailyChallengeComplete int64      `json:"dailyChallengeComplete"`
	NumDailyChallenge      int64      `json:"numDailyChalCont"`
	NumPlaying             int64      `json:"numPlaying,string"` // ?
	Animals                int64      `json:"numAnimals,string"`
	Rank                   int64      `json:"numRank,string"`
	DailyChalCatNum        int64      `json:"ORN_dailyChalCatNum,string"`
	DailyChalSetNum        int64      `json:"ORN_dailyChalSetNum,string"`
	DailyChalPosNum        int64      `json:"ORN_dailyChalPosNum,string"`
	NextNumDailyChallenge  int64      `json:"ORN_nextNumDailyChalCont"`
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
	rankingLeague := int64(enums.RankingLeagueNone)
	quickRankingLeague := int64(enums.RankingLeagueNone)
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
	}
}
