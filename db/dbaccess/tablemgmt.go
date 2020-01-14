package dbaccess

import (
	"github.com/Mtbcooler/outrun/config/eventconf"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/netobj/constnetobjs"
	"github.com/Mtbcooler/outrun/obj"
)

func GetPlayerFromDB(id string) (netobj.Player, error) {
	var playerinfo netobj.PlayerInfo
	err := db.QueryRow("SELECT * FROM ? WHERE id = ?", consts.DBMySQLTableCorePlayerInfo, id).Scan(&playerinfo)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	var playerstate netobj.PlayerState
	err = db.QueryRow("SELECT * FROM ? WHERE id = ?", consts.DBMySQLTablePlayerStates, id).Scan(&playerstate)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	characterstate := netobj.DefaultCharacterState() // TODO: REPLACE ME! FOR TESTING ONLY!
	chaostate := constnetobjs.DefaultChaoState()     // TODO: REPLACE ME! FOR TESTING ONLY!
	var mileagemapstate netobj.MileageMapState
	err = db.QueryRow("SELECT * FROM ? WHERE id = ?", consts.DBMySQLTableMileageMapStates, id).Scan(&mileagemapstate)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	var optionuserresult netobj.OptionUserResult
	err = db.QueryRow("SELECT * FROM ? WHERE id = ?", consts.DBMySQLTableOptionUserResults, id).Scan(&optionuserresult)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	wheeloptions := netobj.DefaultWheelOptions(playerstate.NumRouletteTicket, 0, enums.WheelRankNormal, 5) // TODO: REPLACE ME! FOR TESTING ONLY!
	var rouletteinfo netobj.RouletteInfo
	err = db.QueryRow("SELECT * FROM ? WHERE id = ?", consts.DBMySQLTableRouletteInfos, id).Scan(&rouletteinfo)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
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
	chaoroulettedata := netobj.DefaultChaoRouletteGroup(playerstate, allowedCharacters, allowedChao) // TODO: REPLACE ME! FOR TESTING ONLY!
	var loginbonusstate netobj.LoginBonusState
	err = db.QueryRow("SELECT * FROM ? WHERE id = ?", consts.DBMySQLTableLoginBonusStates, id).Scan(&loginbonusstate)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	player := netobj.NewPlayer(
		playerinfo.ID,
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
