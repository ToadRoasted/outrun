package dbaccess

import (
	"log"

	"github.com/Mtbcooler/outrun/config/eventconf"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/netobj/constnetobjs"
	"github.com/Mtbcooler/outrun/obj"
)

func GetPlayerFromDB(id string) (netobj.Player, error) {
	playerinfo, err := GetPlayerInfo(consts.DBMySQLTableCorePlayerInfo, id)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	playerstate, err := GetPlayerState(consts.DBMySQLTablePlayerStates, id)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	mileagemapstate, err := GetMileageMapState(consts.DBMySQLTableMileageMapStates, id)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	optionuserresult, err := GetOptionUserResult(consts.DBMySQLTableOptionUserResults, id)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	wheeloptions := netobj.DefaultWheelOptions(playerstate.NumRouletteTicket, 0, enums.WheelRankNormal, 5, 0) // TODO: REPLACE ME! FOR TESTING ONLY!
	rouletteinfo, err := GetRouletteInfo(consts.DBMySQLTableRouletteInfos, id)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	allowedCharacters := []string{}
	allowedChao := []string{}
	for _, chao := range playerinfo.ChaoState {
		if chao.Level < 10 { // not max level
			allowedChao = append(allowedChao, chao.ID)
		}
	}
	for _, character := range playerinfo.CharacterState {
		if character.Star < 10 { // not max star
			allowedCharacters = append(allowedCharacters, character.ID)
		}
	}
	chaoroulettedata := netobj.DefaultChaoRouletteGroup(playerstate, allowedCharacters, allowedChao) // TODO: REPLACE ME! FOR TESTING ONLY!
	loginbonusstate, err := GetLoginBonusState(consts.DBMySQLTableLoginBonusStates, id)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	player := netobj.NewPlayer(
		id,
		playerinfo.Username,
		playerinfo.Password,
		playerinfo.MigrationPassword,
		playerinfo.UserPassword,
		playerinfo.Key,
		playerinfo.Language,
		playerstate,
		playerinfo.CharacterState,
		playerinfo.ChaoState,
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

func InitializeTablesIfNecessary() error {
	log.Println("[INFO] Initializing tables... (1/10)")
	_, err := db.Exec(consts.SQLAnalyticsSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (2/10)")
	_, err = db.Exec(consts.SQLCorePlayerInfoSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (3/10)")
	_, err = db.Exec(consts.SQLPlayerStatesSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (4/10)")
	_, err = db.Exec(consts.SQLMileageMapStatesSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (5/10)")
	_, err = db.Exec(consts.SQLOptionUserResultsSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (6/10)")
	_, err = db.Exec(consts.SQLRouletteInfosSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (7/10)")
	_, err = db.Exec(consts.SQLLoginBonusStatesSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (8/10)")
	_, err = db.Exec(consts.SQLOperatorMessagesSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (9/10)")
	_, err = db.Exec(consts.SQLRankingLeagueDataSchema)
	if err != nil {
		return err
	}
	log.Println("[INFO] Initializing tables... (10/10)")
	_, err = db.Exec(consts.SQLSessionIDsSchema)
	return err
}
