package dbaccess

import (
	"reflect"

	"github.com/Mtbcooler/outrun/config/eventconf"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/netobj/constnetobjs"
	"github.com/Mtbcooler/outrun/obj"
)

func GetPlayerFromDB(id string) (netobj.Player, error) {
	var playerinfo netobj.PlayerInfo
	var playerstate netobj.PlayerState
	var characterstate []netobj.Character
	var chaostate []netobj.Chao
	var mileagemapstate netobj.MileageMapState
	var optionuserresult netobj.OptionUserResult
	var wheeloptions netobj.WheelOptions
	var rouletteinfo netobj.RouletteInfo
	var chaoroulettedata netobj.ChaoRouletteGroup
	var loginbonusstate netobj.LoginBonusState
	playerinfoI, err := GetNamed(consts.DBMySQLTableCorePlayerInfo, id, reflect.TypeOf(playerinfo))
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	playerinfo = playerinfoI.(netobj.PlayerInfo)
	playerstateI, err := GetNamed(consts.DBMySQLTablePlayerStates, id, reflect.TypeOf(playerstate))
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	playerstate = playerstateI.(netobj.PlayerState)
	characterstate = netobj.DefaultCharacterState() // TODO: REPLACE ME! FOR TESTING ONLY!
	chaostate = constnetobjs.DefaultChaoState()     // TODO: REPLACE ME! FOR TESTING ONLY!
	mileagemapstateI, err := GetNamed(consts.DBMySQLTableMileageMapStates, id, reflect.TypeOf(mileagemapstate))
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	mileagemapstate = mileagemapstateI.(netobj.MileageMapState)
	optionuserresultI, err := GetNamed(consts.DBMySQLTableOptionUserResults, id, reflect.TypeOf(optionuserresult))
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	optionuserresult = optionuserresultI.(netobj.OptionUserResult)
	wheeloptions = netobj.DefaultWheelOptions(playerstate.NumRouletteTicket, 0, enums.WheelRankNormal, 5) // TODO: REPLACE ME! FOR TESTING ONLY!
	rouletteinfoI, err := GetNamed(consts.DBMySQLTableRouletteInfos, id, reflect.TypeOf(rouletteinfo))
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	rouletteinfo = rouletteinfoI.(netobj.RouletteInfo)
	allowedCharacters := []string{}
	allowedChao := []string{}
	for _, chao := range chaostate {
		if chao.Level < 10 { // not max level
			allowedChao = append(allowedChao, chao.ID)
		}
	}
	for _, character := range characterstate {
		if character.Star < 10 { // not max star
			allowedCharacters = append(allowedCharacters, character.ID)
		}
	}
	chaoroulettedata = netobj.DefaultChaoRouletteGroup(playerstate, allowedCharacters, allowedChao) // TODO: REPLACE ME! FOR TESTING ONLY!
	loginbonusstateI, err := GetNamed(consts.DBMySQLTableLoginBonusStates, id, reflect.TypeOf(loginbonusstate))
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	loginbonusstate = loginbonusstateI.(netobj.LoginBonusState)
	player := netobj.NewPlayer(
		id,
		playerinfo.Username,
		playerinfo.Password,
		playerinfo.MigrationPassword,
		playerinfo.UserPassword,
		playerinfo.Key,
		playerstate,
		characterstate,
		chaostate,
		mileagemapstate,
		[]netobj.MileageFriend{},
		netobj.DefaultPlayerVarious(),
		optionuserresult,
		wheeloptions,
		rouletteinfo,
		chaoroulettedata,
		[]eventconf.ConfiguredEvent{},
		[]obj.Message{},
		[]obj.OperatorMessage{},
		loginbonusstate,
	)
	return player, nil
}
