package db

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Mtbcooler/outrun/config"
	"github.com/Mtbcooler/outrun/config/eventconf"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db/boltdbaccess"
	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/netobj/constnetobjs"
	"github.com/Mtbcooler/outrun/obj"

	bolt "go.etcd.io/bbolt"
)

const (
	SessionIDSchema = "REVIVAL_%s"
)

func NewAccountWithID(uid string) netobj.Player {
	randChar := func(charset string, length int64) string {
		runes := []rune(charset)
		final := make([]rune, 10)
		for i := range final {
			final[i] = runes[rand.Intn(len(runes))]
		}
		return string(final)
	}

	username := ""
	password := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	migrationPassword := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	userPassword := ""
	key := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	playerState := netobj.DefaultPlayerState()
	characterState := netobj.DefaultCharacterState()
	chaoState := constnetobjs.DefaultChaoState()
	mileageMapState := netobj.DefaultMileageMapState()
	mileageFriends := []netobj.MileageFriend{}
	playerVarious := netobj.DefaultPlayerVarious()
	optionUserResult := netobj.DefaultOptionUserResult()
	rouletteInfo := netobj.DefaultRouletteInfo()
	wheelOptions := netobj.DefaultWheelOptions(playerState.NumRouletteTicket, rouletteInfo.RouletteCountInPeriod, enums.WheelRankNormal, consts.RouletteFreeSpins, 0)
	// TODO: get rid of logic here?
	allowedCharacters := []string{}
	allowedChao := []string{}
	for _, chao := range chaoState {
		if chao.Level < 10 { // not max level
			allowedChao = append(allowedChao, chao.ID)
		}
	}
	for _, character := range characterState {
		if character.Star < 10 { // not max star
			allowedCharacters = append(allowedCharacters, character.ID)
		}
	}
	if config.CFile.Debug {
		mileageMapState.Episode = 15
		// testCharacter := netobj.DefaultCharacter(constobjs.CharacterXMasSonic)
		// characterState = append(characterState, testCharacter)
	}
	chaoRouletteGroup := netobj.DefaultChaoRouletteGroup(playerState, allowedCharacters, allowedChao)
	personalEvents := []eventconf.ConfiguredEvent{}
	messages := []obj.Message{}
	operatorMessages := []obj.OperatorMessage{}
	loginBonusState := netobj.DefaultLoginBonusState(0)
	language := int64(enums.LangEnglish)
	return netobj.NewPlayer(
		uid,
		username,
		password,
		migrationPassword,
		userPassword,
		key,
		language,
		playerState,
		characterState,
		chaoState,
		mileageMapState,
		mileageFriends,
		playerVarious,
		optionUserResult,
		wheelOptions,
		rouletteInfo,
		chaoRouletteGroup,
		personalEvents,
		messages,
		operatorMessages,
		loginBonusState,
	)
}

func NewAccount() (netobj.Player, error) {
	// create ID
	attemptsLeft := 500
	newID := ""
	for attemptsLeft > 0 {
		for i := range make([]byte, 10) {
			if i == 0 { // if first character
				newID += strconv.Itoa(rand.Intn(9) + 1)
			} else {
				newID += strconv.Itoa(rand.Intn(10))
			}
		}
		if !DoesPlayerExistInDatabase(newID) {
			return NewAccountWithID(newID), nil
		}
		newID = ""
		attemptsLeft--
	}
	return constnetobjs.BlankPlayer, errors.New("couldn't find an unused ID after 500 attempts")
}

func SavePlayer(player netobj.Player) error {
	playerInfo := netobj.PlayerInfo{
		player.Username,
		player.Password,
		player.MigrationPassword,
		player.UserPassword,
		player.Key,
		player.LastLogin,
		player.Language,
		player.CharacterState,
		player.ChaoState,
	}
	err := dbaccess.SetPlayerInfo(consts.DBMySQLTableCorePlayerInfo, player.ID, playerInfo)
	if err != nil {
		return err
	}
	err = dbaccess.SetPlayerState(consts.DBMySQLTablePlayerStates, player.ID, player.PlayerState)
	if err != nil {
		return err
	}
	err = dbaccess.SetMileageMapState(consts.DBMySQLTableMileageMapStates, player.ID, player.MileageMapState)
	if err != nil {
		return err
	}
	err = dbaccess.SetOptionUserResult(consts.DBMySQLTableOptionUserResults, player.ID, player.OptionUserResult)
	if err != nil {
		return err
	}
	err = dbaccess.SetRouletteInfo(consts.DBMySQLTableRouletteInfos, player.ID, player.RouletteInfo)
	if err != nil {
		return err
	}
	err = dbaccess.SetLoginBonusState(consts.DBMySQLTableLoginBonusStates, player.ID, player.LoginBonusState)
	return err
	// TODO: Add in the rest of the saving!
}

func BoltSavePlayer(player netobj.Player) error {
	j, err := json.Marshal(player)
	if err != nil {
		return err
	}
	err = boltdbaccess.Set(consts.DBBucketPlayers, player.ID, j)
	return err
}

func GetPlayer(uid string) (netobj.Player, error) {
	player, err := dbaccess.GetPlayerFromDB(uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func DoesPlayerExistInDatabase(uid string) bool {
	_, err := dbaccess.GetPlayerInfo(consts.DBMySQLTableCorePlayerInfo, uid)
	if err != nil {
		return false
	} else {
		_, err = dbaccess.GetRouletteInfo(consts.DBMySQLTableRouletteInfos, uid)
		if err != nil {
			return false
		} else {
			_, err = dbaccess.GetLoginBonusState(consts.DBMySQLTableLoginBonusStates, uid)
			if err != nil {
				return false
			} else {
				_, err = dbaccess.GetPlayerState(consts.DBMySQLTablePlayerStates, uid)
				if err != nil {
					return false
				} else {
					_, err = dbaccess.GetOptionUserResult(consts.DBMySQLTableOptionUserResults, uid)
					if err != nil {
						return false
					}
					return true
				}
			}
		}
	}
}

func BoltGetPlayer(uid string) (netobj.Player, error) {
	var player netobj.Player
	playerData, err := boltdbaccess.Get(consts.DBBucketPlayers, uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	err = json.Unmarshal(playerData, &player)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func GetPlayerBySessionID(sid string) (netobj.Player, error) {
	// TODO: Implement this!
	return constnetobjs.BlankPlayer, nil
}

func BoltGetPlayerBySessionID(sid string) (netobj.Player, error) {
	sidResult, err := boltdbaccess.Get(consts.DBBucketSessionIDs, sid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	uid, _ := ParseSIDEntry(sidResult)
	player, err := GetPlayer(uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func GetPlayerIDBySessionID(sid string) (string, error) {
	// TODO: Implement this!
	return "0", nil
}

func BoltGetPlayerIDBySessionID(sid string) (string, error) {
	sidResult, err := boltdbaccess.Get(consts.DBBucketSessionIDs, sid)
	if err != nil {
		return "0", err
	}
	uid, _ := ParseSIDEntry(sidResult)
	return uid, nil
}

func AssignSessionID(uid string) (string, error) {
	// TODO: Implement this!
	datB := []byte(uid + strconv.Itoa(int(time.Now().Unix())))
	hash := sha1.Sum(datB)
	hashStr := fmt.Sprintf("%x", hash)
	sid := fmt.Sprintf(SessionIDSchema, hashStr)
	return sid, nil
}

func BoltAssignSessionID(uid string) (string, error) {
	datB := []byte(uid + strconv.Itoa(int(time.Now().Unix())))
	hash := sha1.Sum(datB)
	hashStr := fmt.Sprintf("%x", hash)
	sid := fmt.Sprintf(SessionIDSchema, hashStr)
	value := fmt.Sprintf("%s/%s", uid, strconv.Itoa(int(time.Now().Unix()))) // register the time that the session ID was assigned
	valueB := []byte(value)
	err := boltdbaccess.Set(consts.DBBucketSessionIDs, sid, valueB)
	return sid, err
}

func ParseSIDEntry(sidResult []byte) (string, int64) {
	split := strings.Split(string(sidResult), "/")
	uid := split[0]
	timeAssigned, _ := strconv.Atoi(split[1])
	return uid, int64(timeAssigned)
}

func IsValidSessionTime(sessionTime int64) bool {
	timeNow := time.Now().Unix()
	if timeNow > sessionTime+consts.DBSessionExpiryTime {
		return false
	}
	return true
}

func IsValidSessionID(sid []byte) (bool, error) {
	// TODO: Implement this!
	return false, nil
}

func BoltIsValidSessionID(sid []byte) (bool, error) {
	sidResult, err := boltdbaccess.Get(consts.DBBucketSessionIDs, string(sid))
	if err != nil {
		return false, err
	}
	_, sessionTime := ParseSIDEntry(sidResult)

	return IsValidSessionTime(sessionTime), err
}

func BoltPurgeSessionID(sid string) error {
	err := boltdbaccess.Delete(consts.DBBucketSessionIDs, sid)
	return err
}

func PurgeAllExpiredSessionIDs() {
	// TODO: Implement this!
}

func BoltPurgeAllExpiredSessionIDs() {
	keysToPurge := [][]byte{}
	each := func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(consts.DBBucketSessionIDs))
		err2 := bucket.ForEach(func(k, v []byte) error { // for each value in the session bucket
			_, sessionTime := ParseSIDEntry(v) // get time the session was created
			if !IsValidSessionTime(sessionTime) {
				keysToPurge = append(keysToPurge, k)
			}
			return nil
		})
		return err2
	}
	boltdbaccess.ForEachLogic(each) // do the logic above
	for _, key := range keysToPurge {
		BoltPurgeSessionID(string(key))
	}
}
