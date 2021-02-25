package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ToadRoasted/outrun/analytics"
	"github.com/ToadRoasted/outrun/analytics/factors"
	"github.com/ToadRoasted/outrun/config"
	"github.com/ToadRoasted/outrun/config/campaignconf"
	"github.com/ToadRoasted/outrun/config/gameconf"
	"github.com/ToadRoasted/outrun/consts"
	"github.com/ToadRoasted/outrun/db"
	"github.com/ToadRoasted/outrun/emess"
	"github.com/ToadRoasted/outrun/enums"
	"github.com/ToadRoasted/outrun/helper"
	"github.com/ToadRoasted/outrun/logic/campaign"
	"github.com/ToadRoasted/outrun/logic/conversion"
	"github.com/ToadRoasted/outrun/logic/gameplay"
	"github.com/ToadRoasted/outrun/netobj"
	"github.com/ToadRoasted/outrun/obj"
	"github.com/ToadRoasted/outrun/obj/constobjs"
	"github.com/ToadRoasted/outrun/requests"
	"github.com/ToadRoasted/outrun/responses"
	"github.com/ToadRoasted/outrun/status"
	"github.com/jinzhu/now"
)

func GetDailyChallengeData(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DailyChallengeData(baseInfo, player.PlayerState.NumDailyChallenge, player.PlayerState.NextNumDailyChallenge)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetCostList(helper *helper.Helper) {
	// no player, agonstic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultCostList(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetMileageData(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultMileageData(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetCampaignList(helper *helper.Helper) {
	campaignList := []obj.Campaign{}
	if campaignconf.CFile.AllowCampaigns {
		for _, confCampaign := range campaignconf.CFile.CurrentCampaigns {
			newCampaign := conversion.ConfiguredCampaignToCampaign(confCampaign)
			campaignList = append(campaignList, newCampaign)
		}
	}
	helper.DebugOut("Campaign list: %v", campaignList)
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.CampaignList(baseInfo, campaignList)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func QuickActStart(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.QuickActStartRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	responseStatus := status.OK
	campaignList := []obj.Campaign{}
	if campaignconf.CFile.AllowCampaigns {
		for _, confCampaign := range campaignconf.CFile.CurrentCampaigns {
			newCampaign := conversion.ConfiguredCampaignToCampaign(confCampaign)
			campaignList = append(campaignList, newCampaign)
		}
	}
	helper.DebugOut("Campaign list: %v", campaignList)
	// consume items
	modToStringSlice := func(ns []int64) []string {
		result := []string{}
		for _, n := range ns {
			result = append(result, fmt.Sprintf("%v", n))
		}
		return result
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}
	if player.PlayerState.Energy > 0 {
		if gameconf.CFile.EnableEnergyConsumption {
			if player.PlayerState.Energy >= player.PlayerVarious.EnergyRecoveryMax {
				player.PlayerState.EnergyRenewsAt = time.Now().UTC().Unix() + player.PlayerVarious.EnergyRecoveryTime
			}
			player.PlayerState.Energy--
		}
		player.PlayerState.NumPlaying++
		if !gameconf.CFile.AllItemsFree {
			consumedItems := modToStringSlice(request.Modifier)
			consumedRings := gameplay.GetRequiredItemPayment(consumedItems, player)
			for _, citemID := range consumedItems {
				if citemID[:2] == "11" { // boosts, not items
					continue
				}
				index := player.IndexOfItem(citemID)
				if index == -1 {
					helper.Uncatchable(fmt.Sprintf("Player sent bad item ID '%s', cannot continue", citemID))
					helper.InvalidRequest()
					return
				}
				if player.PlayerState.Items[index].Amount >= 1 { // can use item
					player.PlayerState.Items[index].Amount--
				}
			}
			if player.PlayerState.NumRings < consumedRings { // not enough rings
				responseStatus = status.NotEnoughRings
			}
			player.PlayerState.NumRings -= consumedRings
		}
		helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))
	} else {
		responseStatus = status.NotEnoughEnergy
	}
	baseInfo := helper.BaseInfo(emess.OK, responseStatus)
	response := responses.DefaultQuickActStart(baseInfo, player, campaignList)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeTimedStarts)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeTimedStarts)", err)
	}
}

func ActStart(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.ActStartRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))
	responseStatus := status.OK
	campaignList := []obj.Campaign{}
	if campaignconf.CFile.AllowCampaigns {
		for _, confCampaign := range campaignconf.CFile.CurrentCampaigns {
			newCampaign := conversion.ConfiguredCampaignToCampaign(confCampaign)
			campaignList = append(campaignList, newCampaign)
		}
	}
	helper.DebugOut("Campaign list: %v", campaignList)
	// consume items
	modToStringSlice := func(ns []int64) []string {
		result := []string{}
		for _, n := range ns {
			result = append(result, fmt.Sprintf("%v", n))
		}
		return result
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}
	if player.PlayerState.Energy > 0 {
		if gameconf.CFile.EnableEnergyConsumption {
			if player.PlayerState.Energy >= player.PlayerVarious.EnergyRecoveryMax {
				player.PlayerState.EnergyRenewsAt = time.Now().UTC().Unix() + player.PlayerVarious.EnergyRecoveryTime
			}
			player.PlayerState.Energy--
		}
		player.PlayerState.NumPlaying++
		if !gameconf.CFile.AllItemsFree {
			consumedItems := modToStringSlice(request.Modifier)
			consumedRings := gameplay.GetRequiredItemPayment(consumedItems, player)
			for _, citemID := range consumedItems {
				if citemID[:2] == "11" { // boosts, not items
					continue
				}
				index := player.IndexOfItem(citemID)
				if index == -1 {
					helper.Uncatchable(fmt.Sprintf("Player sent bad item ID '%s', cannot continue", citemID))
					helper.InvalidRequest()
					return
				}
				if player.PlayerState.Items[index].Amount >= 1 { // can use item
					player.PlayerState.Items[index].Amount--
				}
			}
			if player.PlayerState.NumRings < consumedRings { // not enough rings
				responseStatus = status.NotEnoughRings
			}
			player.PlayerState.NumRings -= consumedRings
		}
		helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))
	} else {
		responseStatus = status.NotEnoughEnergy
	}
	baseInfo := helper.BaseInfo(emess.OK, responseStatus)
	response := responses.DefaultActStart(baseInfo, player, campaignList)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeStoryStarts)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeStoryStarts)", err)
	}
}

func ActRetry(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	redRingContinuePrice := 5
	// TODO: Add campaign support
	responseStatus := status.OK
	if player.PlayerState.NumRedRings >= int64(redRingContinuePrice) { //does the player actually have enough red rings?
		//if so, subtract 5 red rings and respond with an OK status
		player.PlayerState.NumRedRings -= int64(redRingContinuePrice)
		err = db.SavePlayer(player)
		if err != nil {
			helper.InternalErr("Error saving player", err)
			return
		}
	} else {
		//if not, respond with NotEnoughRedRings
		responseStatus = status.NotEnoughRedRings
	}
	baseInfo := helper.BaseInfo(emess.OK, responseStatus)
	response := responses.NewBaseResponse(baseInfo)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeRevives)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeRevives)", err)
	}
}

func QuickPostGameResults(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.QuickPostGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}

	hasSubCharacter := player.PlayerState.SubCharaID != "-1"
	var subC netobj.Character
	mainC, err := player.GetMainChara()
	if err != nil {
		helper.InternalErr("Error getting main character", err)
		return
	}
	playCharacters := []netobj.Character{ // assume only main character active right now
		mainC,
	}
	if hasSubCharacter {
		subC, err = player.GetSubChara()
		if err != nil {
			helper.InternalErr("Error getting sub character", err)
			return
		}
		playCharacters = []netobj.Character{ // add sub character to playCharacters
			mainC,
			subC,
		}
	}
	if request.Closed == 0 { // If the game wasn't exited out of
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
		player.PlayerState.NumRouletteTicket += request.RedRings // TODO: URGENT! Remove as soon as possible!
		player.PlayerState.Animals += request.Animals
		player.OptionUserResult.NumTakeAllRings += request.Rings
		player.OptionUserResult.NumTakeAllRedRings += request.RedRings
		playerTimedHighScore := player.PlayerState.TimedHighScore
		if request.Score > playerTimedHighScore {
			player.PlayerState.TimedHighScore = request.Score
		}
		helper.DebugOut("request.DailyChallengeValue: %v", request.DailyChallengeValue)
		helper.DebugOut("request.DailyChallengeComplete: %v", request.DailyChallengeComplete)
		if player.PlayerState.DailyChallengeComplete == 0 && request.DailyChallengeComplete == 1 {
			if player.PlayerState.NextNumDailyChallenge <= 0 {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1)
			}
			player.AddOperatorMessage(
				"A Daily Challenge Reward.",
				obj.NewMessageItem(
					consts.DailyMissionRewards[player.PlayerState.NextNumDailyChallenge-1],
					consts.DailyMissionRewardCounts[player.PlayerState.NextNumDailyChallenge-1],
					0,
					0,
				),
				2592000,
			)
			player.PlayerState.NumDailyChallenge = player.PlayerState.NextNumDailyChallenge
		}
		if player.PlayerState.DailyChallengeComplete == 0 {
			player.PlayerState.DailyChallengeComplete = request.DailyChallengeComplete
		}
		player.PlayerState.DailyChallengeValue = request.DailyChallengeValue
		if time.Now().UTC().Unix() >= player.PlayerState.DailyMissionEndTime {
			if player.PlayerState.DailyChallengeComplete == 1 && player.PlayerState.DailyChalSetNum < 10 {
				helper.DebugOut("Advancing to next daily mission...")
				player.PlayerState.DailyChalSetNum++
			} else {
				player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
				player.PlayerState.DailyChalSetNum = int64(0)
			}
			if player.PlayerState.DailyChallengeComplete == 0 {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1)
			} else {
				player.PlayerState.NextNumDailyChallenge++
				if int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
					player.PlayerState.NumDailyChallenge = int64(0)
					player.PlayerState.NextNumDailyChallenge = int64(1) //restart from beginning
					player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
					player.PlayerState.DailyChalSetNum = int64(0)
				}
			}
			player.PlayerState.DailyChalPosNum = int64(1 + rand.Intn(2))
			player.PlayerState.DailyMissionID = int64((player.PlayerState.DailyChalCatNum * 33) + (player.PlayerState.DailyChalSetNum * 3) + player.PlayerState.DailyChalPosNum)
			player.PlayerState.DailyChallengeValue = int64(0)
			player.PlayerState.DailyChallengeComplete = int64(0)
			player.PlayerState.DailyMissionEndTime = now.EndOfDay().UTC().Unix() + 1
			helper.DebugOut("New daily mission ID: %v", player.PlayerState.DailyMissionID)
		}
		//player.PlayerState.TotalDistance += request.Distance  // We don't do this in timed mode!

		sum := func(in []int64) int64 {
			v := int64(0)
			for _, val := range in {
				v += val
			}
			return v
		}

		// increase character(s)'s experience
		expIncrease := request.Rings + request.FailureRings // all rings collected
		mainAbilityIndex := 1
		mainAbilitySum := sum(mainC.AbilityLevel)
		if mainAbilitySum < 100 {
			for mainAbilityIndex == 1 || mainC.AbilityLevel[mainAbilityIndex] >= 10 { // unused ability is at index 1
				mainAbilityIndex = rand.Intn(len(mainC.AbilityLevel))
			}
		} else {
			helper.DebugOut("Main character seems to be maxed out on abilities!")
		}
		subAbilityIndex := 1
		subAbilitySum := mainAbilitySum
		if hasSubCharacter {
			subAbilitySum = sum(subC.AbilityLevel)
			if subAbilitySum < 100 {
				for subAbilityIndex == 1 || subC.AbilityLevel[subAbilityIndex] >= 10 { // unused ability is at index 1
					subAbilityIndex = rand.Intn(len(subC.AbilityLevel))
				}
			} else {
				helper.DebugOut("Sub character seems to be maxed out on abilities!")
			}
		}
		// check that increases exist
		_, ok := consts.UpgradeIncreases[mainC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", mainC.ID))
			return
		}
		if hasSubCharacter {
			_, ok = consts.UpgradeIncreases[subC.ID]
			if !ok {
				helper.InternalErr("Error getting upgrade increase for sub character", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", subC.ID))
				return
			}
		}
		if playCharacters[0].Level < 100 {
			playCharacters[0].Exp += expIncrease
			for playCharacters[0].Exp >= playCharacters[0].Cost {
				// more exp than cost = level up
				if playCharacters[0].Level < 100 {
					playCharacters[0].Level++                                               // increase level
					playCharacters[0].AbilityLevel[mainAbilityIndex]++                      // increase ability level
					playCharacters[0].Exp -= playCharacters[0].Cost                         // remove cost from exp
					playCharacters[0].Cost += consts.UpgradeIncreases[playCharacters[0].ID] // increase cost
					mainAbilitySum = sum(playCharacters[0].AbilityLevel)
					if mainAbilitySum < 100 {
						for mainAbilityIndex == 1 || playCharacters[0].AbilityLevel[mainAbilityIndex] >= 10 { // reroll ability index
							mainAbilityIndex = rand.Intn(len(playCharacters[0].AbilityLevel))
						}
					}
				} else {
					helper.DebugOut("Main character is level 100; cannot level up anymore!")
					playCharacters[0].Exp -= playCharacters[0].Cost
				}
			}
		}
		if hasSubCharacter {
			if playCharacters[1].Level < 100 {
				playCharacters[1].Exp += expIncrease
				for playCharacters[1].Exp >= playCharacters[1].Cost {
					// more exp than cost = level up
					if playCharacters[1].Level < 100 {
						playCharacters[1].Level++                                               // increase level
						playCharacters[1].AbilityLevel[subAbilityIndex]++                       // increase ability level
						playCharacters[1].Exp -= playCharacters[1].Cost                         // remove cost from exp
						playCharacters[1].Cost += consts.UpgradeIncreases[playCharacters[1].ID] // increase cost
						subAbilitySum = sum(playCharacters[1].AbilityLevel)
						if subAbilitySum < 100 {
							for subAbilityIndex == 1 || playCharacters[1].AbilityLevel[subAbilityIndex] >= 10 { // reroll ability index
								subAbilityIndex = rand.Intn(len(playCharacters[1].AbilityLevel))
							}
						}
					} else {
						helper.DebugOut("Sub character is level 100; cannot level up anymore!")
						playCharacters[1].Exp -= playCharacters[1].Cost
					}
				}
			}
		}

		helper.DebugOut("Old mainC Exp: %v / %v", mainC.Exp, mainC.Cost)
		helper.DebugOut("Old mainC Level: %v", mainC.Level)
		if hasSubCharacter {
			helper.DebugOut("Old subC Exp: %v / %v", subC.Exp, subC.Cost)
			helper.DebugOut("Old subC Level: %v", subC.Level)
		}
		helper.DebugOut("New mainC Exp: %v / %v", playCharacters[0].Exp, playCharacters[0].Cost)
		helper.DebugOut("New mainC Level: %v", playCharacters[0].Level)
		if hasSubCharacter {
			helper.DebugOut("New subC Exp: %v / %v", playCharacters[1].Exp, playCharacters[1].Cost)
			helper.DebugOut("New subC Level: %v", playCharacters[1].Level)
		}
	}

	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := -1
	if hasSubCharacter {
		subCIndex = player.IndexOfChara(subC.ID) // TODO: check if -1
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultQuickPostGameResults(baseInfo, player, playCharacters)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	// apply the save after the response so that we don't break the leveling
	mainC = playCharacters[0]
	if hasSubCharacter {
		subC = playCharacters[1]
	}
	player.CharacterState[mainCIndex] = mainC
	if hasSubCharacter {
		player.CharacterState[subCIndex] = subC
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))

	_, err = analytics.Store(player.ID, factors.AnalyticTypeTimedEnds)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeTimedEnds)", err)
	}
}

func PostGameResults(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.PostGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}

	hasSubCharacter := player.PlayerState.SubCharaID != "-1"
	var subC netobj.Character
	mainC, err := player.GetMainChara()
	if err != nil {
		helper.InternalErr("Error getting main character", err)
		return
	}
	playCharacters := []netobj.Character{ // assume only main character active right now
		mainC,
	}
	if hasSubCharacter {
		subC, err = player.GetSubChara()
		if err != nil {
			helper.InternalErr("Error getting sub character", err)
			return
		}
		playCharacters = []netobj.Character{ // add sub character to playCharacters
			mainC,
			subC,
		}
	}
	helper.DebugOut("Pre-function")
	helper.DebugOut("Chapter: %v", player.MileageMapState.Chapter)
	helper.DebugOut("Episode: %v", player.MileageMapState.Episode)
	helper.DebugOut("StageTotalScore: %v", player.MileageMapState.StageTotalScore)
	helper.DebugOut("Point: %v", player.MileageMapState.Point)
	helper.DebugOut("request.Score: %v", request.Score)

	incentives := constobjs.GetMileageIncentives(player.MileageMapState.Episode, player.MileageMapState.Chapter) // Game wants incentives in _current_ episode-chapter
	var oldRewardEpisode, newRewardEpisode int64
	var oldRewardChapter, newRewardChapter int64
	var oldRewardPoint, newRewardPoint int64

	if request.Closed == 0 { // If the game wasn't exited out of
		oldRewardEpisode = player.MileageMapState.Episode
		oldRewardChapter = player.MileageMapState.Chapter
		oldRewardPoint = player.MileageMapState.Point
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
		player.PlayerState.NumRouletteTicket += request.RedRings // TODO: URGENT! Remove as soon as possible!
		player.PlayerState.Animals += request.Animals
		player.OptionUserResult.NumTakeAllRings += request.Rings
		player.OptionUserResult.NumTakeAllRedRings += request.RedRings
		playerHighScore := player.PlayerState.HighScore
		if request.Score > playerHighScore {
			player.PlayerState.HighScore = request.Score
		}
		helper.DebugOut("request.DailyChallengeValue: %v", request.DailyChallengeValue)
		helper.DebugOut("request.DailyChallengeComplete: %v", request.DailyChallengeComplete)
		if player.PlayerState.DailyChallengeComplete == 0 && request.DailyChallengeComplete == 1 {
			if player.PlayerState.NextNumDailyChallenge <= 0 {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1)
			}
			player.AddOperatorMessage(
				"A Daily Challenge Reward.",
				obj.NewMessageItem(
					consts.DailyMissionRewards[player.PlayerState.NextNumDailyChallenge-1],
					consts.DailyMissionRewardCounts[player.PlayerState.NextNumDailyChallenge-1],
					0,
					0,
				),
				2592000,
			)
			player.PlayerState.NumDailyChallenge = player.PlayerState.NextNumDailyChallenge
		}
		if player.PlayerState.DailyChallengeComplete == 0 {
			player.PlayerState.DailyChallengeComplete = request.DailyChallengeComplete
		}
		player.PlayerState.DailyChallengeValue = request.DailyChallengeValue
		if time.Now().UTC().Unix() >= player.PlayerState.DailyMissionEndTime {
			if player.PlayerState.DailyChallengeComplete == 1 && player.PlayerState.DailyChalSetNum < 10 {
				helper.DebugOut("Advancing to next daily mission...")
				player.PlayerState.DailyChalSetNum++
			} else {
				player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
				player.PlayerState.DailyChalSetNum = int64(0)
			}
			if player.PlayerState.DailyChallengeComplete == 0 {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1)
			} else {
				player.PlayerState.NextNumDailyChallenge++
				if int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
					player.PlayerState.NumDailyChallenge = int64(0)
					player.PlayerState.NextNumDailyChallenge = int64(1) //restart from beginning
					player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
					player.PlayerState.DailyChalSetNum = int64(0)
				}
			}
			player.PlayerState.DailyChalPosNum = int64(1 + rand.Intn(2))
			player.PlayerState.DailyMissionID = int64((player.PlayerState.DailyChalCatNum * 33) + (player.PlayerState.DailyChalSetNum * 3) + player.PlayerState.DailyChalPosNum)
			player.PlayerState.DailyChallengeValue = int64(0)
			player.PlayerState.DailyChallengeComplete = int64(0)
			player.PlayerState.DailyMissionEndTime = now.EndOfDay().UTC().Unix() + 1
			helper.DebugOut("New daily mission ID: %v", player.PlayerState.DailyMissionID)
		}
		player.PlayerState.TotalDistance += request.Distance
		playerHighDistance := player.PlayerState.HighDistance
		if request.Distance > playerHighDistance {
			player.PlayerState.HighDistance = request.Distance
		}

		sum := func(in []int64) int64 {
			v := int64(0)
			for _, val := range in {
				v += val
			}
			return v
		}

		// increase character(s)'s experience
		expIncrease := request.Rings + request.FailureRings // all rings collected
		mainAbilityIndex := 1
		mainAbilitySum := sum(mainC.AbilityLevel)
		if mainAbilitySum < 100 {
			for mainAbilityIndex == 1 || mainC.AbilityLevel[mainAbilityIndex] >= 10 { // unused ability is at index 1
				mainAbilityIndex = rand.Intn(len(mainC.AbilityLevel))
			}
		} else {
			helper.DebugOut("Main character seems to be maxed out on abilities!")
		}
		subAbilityIndex := 1
		subAbilitySum := mainAbilitySum
		if hasSubCharacter {
			subAbilitySum = sum(subC.AbilityLevel)
			if subAbilitySum < 100 {
				for subAbilityIndex == 1 || subC.AbilityLevel[subAbilityIndex] >= 10 { // unused ability is at index 1
					subAbilityIndex = rand.Intn(len(subC.AbilityLevel))
				}
			} else {
				helper.DebugOut("Sub character seems to be maxed out on abilities!")
			}
		}
		// check that increases exist
		_, ok := consts.UpgradeIncreases[mainC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase for main character", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", mainC.ID))
			return
		}
		if hasSubCharacter {
			_, ok = consts.UpgradeIncreases[subC.ID]
			if !ok {
				helper.InternalErr("Error getting upgrade increase for sub character", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", subC.ID))
				return
			}
		}
		if playCharacters[0].Level < 100 {
			playCharacters[0].Exp += expIncrease
			for playCharacters[0].Exp >= playCharacters[0].Cost {
				// more exp than cost = level up
				if playCharacters[0].Level < 100 {
					playCharacters[0].Level++                                               // increase level
					playCharacters[0].AbilityLevel[mainAbilityIndex]++                      // increase ability level
					playCharacters[0].Exp -= playCharacters[0].Cost                         // remove cost from exp
					playCharacters[0].Cost += consts.UpgradeIncreases[playCharacters[0].ID] // increase cost
					mainAbilitySum = sum(playCharacters[0].AbilityLevel)
					if mainAbilitySum < 100 {
						for mainAbilityIndex == 1 || playCharacters[0].AbilityLevel[mainAbilityIndex] >= 10 { // reroll ability index
							mainAbilityIndex = rand.Intn(len(playCharacters[0].AbilityLevel))
						}
					}
				} else {
					helper.DebugOut("Main character is level 100; cannot level up anymore!")
					playCharacters[0].Exp -= playCharacters[0].Cost
				}
			}
		}
		if hasSubCharacter {
			if playCharacters[1].Level < 100 {
				playCharacters[1].Exp += expIncrease
				for playCharacters[1].Exp >= playCharacters[1].Cost {
					// more exp than cost = level up
					if playCharacters[1].Level < 100 {
						playCharacters[1].Level++                                               // increase level
						playCharacters[1].AbilityLevel[subAbilityIndex]++                       // increase ability level
						playCharacters[1].Exp -= playCharacters[1].Cost                         // remove cost from exp
						playCharacters[1].Cost += consts.UpgradeIncreases[playCharacters[1].ID] // increase cost
						subAbilitySum = sum(playCharacters[1].AbilityLevel)
						if subAbilitySum < 100 {
							for subAbilityIndex == 1 || playCharacters[1].AbilityLevel[subAbilityIndex] >= 10 { // reroll ability index
								subAbilityIndex = rand.Intn(len(playCharacters[1].AbilityLevel))
							}
						}
					} else {
						helper.DebugOut("Sub character is level 100; cannot level up anymore!")
						playCharacters[1].Exp -= playCharacters[1].Cost
					}
				}
			}
		}

		helper.DebugOut("Old mainC Exp: %v / %v", mainC.Exp, mainC.Cost)
		helper.DebugOut("Old mainC Level: %v", mainC.Level)
		if hasSubCharacter {
			helper.DebugOut("Old subC Exp: %v / %v", subC.Exp, subC.Cost)
			helper.DebugOut("Old subC Level: %v", subC.Level)
		}
		helper.DebugOut("New mainC Exp: %v / %v", playCharacters[0].Exp, playCharacters[0].Cost)
		helper.DebugOut("New mainC Level: %v", playCharacters[0].Level)
		if hasSubCharacter {
			helper.DebugOut("New subC Exp: %v / %v", playCharacters[1].Exp, playCharacters[1].Cost)
			helper.DebugOut("New subC Level: %v", playCharacters[1].Level)
		}

		player.MileageMapState.StageTotalScore += request.Score

		goToNextChapter := request.ChapterClear == 1
		chaoEggs := request.GetChaoEgg
		player.PlayerState.ChaoEggs += chaoEggs
		if chaoEggs > 0 || player.PlayerState.ChaoEggs >= 10 {
			player.ChaoRouletteGroup.ChaoWheelOptions = netobj.DefaultChaoWheelOptions(player.PlayerState) // create a new wheel
			if player.PlayerState.ChaoEggs >= 10 {
				player.PlayerState.ChaoEggs = 10
			}
		}

		// TODO: Add chao eggs to player
		newPoint := request.ReachPoint

		goToNextEpisode := true
		if goToNextChapter {
			// Assumed this just means next episode...
			player.PlayerState.Rank++
			if player.PlayerState.Rank > 998 { // rank going above 999
				player.PlayerState.Rank = 998
			}
			if player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
				player.PlayerState.Energy = player.PlayerVarious.EnergyRecoveryMax //restore energy
			}
			player.MileageMapState.Point = 0
			player.MileageMapState.StageTotalScore = 0
			maxChapters, episodeHasMultipleChapters := consts.EpisodeWithChapters[player.MileageMapState.Episode]
			if episodeHasMultipleChapters {
				goToNextEpisode = false
				player.MileageMapState.Chapter++
				if player.MileageMapState.Chapter > maxChapters {
					// there's no more chapters for this episode!
					goToNextEpisode = true
				}
			}
			if goToNextEpisode {
				player.MileageMapState.Episode++
				player.MileageMapState.Chapter = 1
				helper.DebugOut("goToNextEpisode -> Episode: %v", player.MileageMapState.Episode)
				if config.CFile.Debug {
					player.MileageMapState.Episode = 11
				}
			}
		} else {
			player.MileageMapState.Point = newPoint
		}
		/*if config.CFile.Debug {
			if player.MileageMapState.Episode < 11 {
				//player.MileageMapState.Episode = 11
			}
		}*/
		newRewardEpisode = player.MileageMapState.Episode
		newRewardChapter = player.MileageMapState.Chapter
		newRewardPoint = player.MileageMapState.Point
		if goToNextChapter && goToNextEpisode && player.MileageMapState.Episode > 50 { // if beat game, reset to 50-1
			player.MileageMapState.Episode = 50
			player.MileageMapState.Chapter = 1
			helper.DebugOut("goToNextEpisode: Player (%s) beat the game!", player.ID)
		}
		// add rewards to PlayerState
		wonRewards := campaign.GetWonRewards(oldRewardEpisode, oldRewardChapter, oldRewardPoint, newRewardEpisode, newRewardChapter, newRewardPoint)
		helper.DebugOut("wonRewards length: %v", wonRewards)
		helper.DebugOut("Previous red rings: %v", player.PlayerState.NumRings)
		helper.DebugOut("Previous rings: %v", player.PlayerState.NumRings)
		newItems := player.PlayerState.Items
		for _, reward := range wonRewards { // TODO: This is O(n^2). Maybe alleviate this?
			helper.DebugOut("Reward: %s", reward.ItemID)
			helper.DebugOut("Reward amount: %v", reward.NumItem)
			if reward.ItemID[:2] == "12" { // ID is an item
				// check if the item is already in the player's inventory
				itemIndex := player.IndexOfItem(reward.ItemID)
				if itemIndex != -1 {
					player.PlayerState.Items[itemIndex].Amount += reward.NumItem
				} else {
					helper.Warn("Unknown item reward '%s'", reward.ItemID)
				}
			} else if reward.ItemID == strconv.Itoa(enums.ItemIDRing) { // Rings
				player.PlayerState.NumRings += reward.NumItem
			} else if reward.ItemID == strconv.Itoa(enums.ItemIDRedRing) { // Red rings
				player.PlayerState.NumRedRings += reward.NumItem
			} else if reward.ItemID == enums.CTStrTails { // Tails node
				tailsIndex := player.IndexOfChara(enums.CTStrTails)
				player.CharacterState[tailsIndex].Status = enums.CharacterStatusUnlocked
			} else if reward.ItemID == enums.CTStrKnuckles { // Knuckles node
				knucklesIndex := player.IndexOfChara(enums.CTStrKnuckles)
				player.CharacterState[knucklesIndex].Status = enums.CharacterStatusUnlocked
			} else {
				helper.Warn("Unknown reward '%s', ignoring", reward.ItemID)
			}
			// TODO: allow for any character joining the cast, for custom story episodes
		}
		helper.DebugOut("Current rings: %v", player.PlayerState.NumRings)
		player.PlayerState.Items = newItems
	}

	helper.DebugOut("Chapter: %v", player.MileageMapState.Chapter)
	helper.DebugOut("Episode: %v", player.MileageMapState.Episode)
	helper.DebugOut("StageTotalScore: %v", player.MileageMapState.StageTotalScore)
	helper.DebugOut("Point: %v", player.MileageMapState.Point)
	helper.DebugOut("request.Score: %v", request.Score)

	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := -1
	if hasSubCharacter {
		subCIndex = player.IndexOfChara(subC.ID) // TODO: check if -1
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultPostGameResults(baseInfo, player, playCharacters, incentives)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	// apply the save after the response so that we don't break the leveling
	mainC = playCharacters[0]
	if hasSubCharacter {
		subC = playCharacters[1]
	}
	player.CharacterState[mainCIndex] = mainC
	if hasSubCharacter {
		player.CharacterState[subCIndex] = subC
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	helper.DebugOut(fmt.Sprintf("%v", player.PlayerState.Items))

	_, err = analytics.Store(player.ID, factors.AnalyticTypeStoryEnds)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeStoryEnds)", err)
	}
}

func GetFreeItemList(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response responses.FreeItemListResponse
	if gameconf.CFile.AllItemsFree {
		response = responses.DefaultFreeItemList(baseInfo)
	} else {
		response = responses.FreeItemList(baseInfo, []obj.Item{}) // No free items
	}
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetMileageReward(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.MileageRewardRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	/*
		player, err := helper.GetCallingPlayer()
		if err != nil {
			helper.InternalErr("Error getting calling player", err)
			return
		}
	*/
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultMileageReward(baseInfo, request.Chapter, request.Episode)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
