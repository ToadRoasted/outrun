package logic

import (
	"github.com/ToadRoastedToadRoasted/outrun/config/campaignconf"
	"github.com/ToadRoasted/outrun/consts"
	"github.com/ToadRoasted/outrun/enums"
	"github.com/ToadRoasted/outrun/logic/conversion"
	"github.com/ToadRoasted/outrun/netobj"
	"github.com/ToadRoasted/outrun/obj"
)

func WheelRefreshLogic(player netobj.Player, wheel netobj.WheelOptions) netobj.WheelOptions {
	freeSpins := consts.RouletteFreeSpins
	campaignList := []obj.Campaign{}
	if campaignconf.CFile.AllowCampaigns {
		for _, confCampaign := range campaignconf.CFile.CurrentCampaigns {
			newCampaign := conversion.ConfiguredCampaignToCampaign(confCampaign)
			campaignList = append(campaignList, newCampaign)
		}
	}
	for index := range campaignList {
		if obj.IsCampaignActive(campaignList[index]) && campaignList[index].Type == enums.CampaignTypeFreeWheelSpinCount {
			freeSpins = campaignList[index].Content
		}
		index++
	}

	// TODO: Find a more standard way of refreshing the wheel status, because this is scary code
	numRouletteTicket := player.PlayerState.NumRouletteTicket  // get roulette tickets
	rouletteCount := player.RouletteInfo.RouletteCountInPeriod // get amount of times we've spun the wheel today
	if player.RouletteInfo.GotJackpotThisPeriod {
		wheel.NumJackpotRing = 1
	}
	wheel.NumRouletteToken = numRouletteTicket
	wheel.NumRemainingRoulette = wheel.NumRouletteToken + freeSpins - rouletteCount // TODO: is this proper?
	if wheel.NumRemainingRoulette < wheel.NumRouletteToken {
		wheel.NumRemainingRoulette = wheel.NumRouletteToken
	}

	return wheel
}
