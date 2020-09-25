package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Mtbcooler/outrun/logic"

	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/responses/responseobjs"
	"github.com/Mtbcooler/outrun/status"

	"github.com/Mtbcooler/outrun/enums"

	"github.com/Mtbcooler/outrun/db/boltdbaccess"
	"github.com/Mtbcooler/outrun/db/dbaccess"

	"github.com/Mtbcooler/outrun/bgtasks"
	"github.com/Mtbcooler/outrun/config"
	"github.com/Mtbcooler/outrun/config/campaignconf"
	"github.com/Mtbcooler/outrun/config/eventconf"
	"github.com/Mtbcooler/outrun/config/gameconf"
	"github.com/Mtbcooler/outrun/config/infoconf"
	"github.com/Mtbcooler/outrun/cryption"
	"github.com/Mtbcooler/outrun/inforeporters"
	"github.com/Mtbcooler/outrun/meta"
	"github.com/Mtbcooler/outrun/muxhandlers"
	"github.com/Mtbcooler/outrun/muxhandlers/muxobj"
	"github.com/Mtbcooler/outrun/orpc"
	"github.com/gorilla/mux"
)

const UNKNOWN_REQUEST_DIRECTORY = "logging/unknown_requests/"

var (
	LogExecutionTime = true
	MaintenanceMode  = false
)

func OutputUnknownRequest(w http.ResponseWriter, r *http.Request) {
	recv := cryption.GetReceivedMessage(r)
	// make a new logging path
	timeStr := strconv.Itoa(int(time.Now().Unix()))
	os.MkdirAll(UNKNOWN_REQUEST_DIRECTORY, 0644)
	normalizedReq := strings.ReplaceAll(r.URL.Path, "/", "-")
	path := UNKNOWN_REQUEST_DIRECTORY + normalizedReq + "_" + timeStr + ".txt"
	err := ioutil.WriteFile(path, recv, 0644)
	if err != nil {
		log.Println("[OUT] UNABLE TO WRITE UNKNOWN REQUEST: " + err.Error())
		HandleUnknownRequest(w, r)
		return
	}
	log.Println("[OUT] !!!!!!!!!!!! Unknown request, output to " + path)
	HandleUnknownRequest(w, r)
}

func HandleUnknownRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

// return Not Found for the favicon; no favicon is intended
func FaviconResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

// Return "OK" for checking if the Outrun instance is alive (intended for uptime monitors)
func GenericRootResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func MaintenanceResponse(w http.ResponseWriter, r *http.Request) {
	baseInfo := responseobjs.NewBaseInfo(emess.OK, status.ServerMaintenance)
	out := responses.NewBaseResponse(baseInfo)
	response := map[string]interface{}{"secure": "0", "param": out}
	toClient, err := json.Marshal(response)
	if err != nil {
		log.Println("[ERR] Error marshalling in MaintenanceResponse")
	}
	w.Write(toClient)
}

func removePrependingSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for len(r.URL.Path) != 0 && string(r.URL.Path[0]) == "/" {
			r.URL.Path = r.URL.Path[1:]
		}
		r.URL.Path = "/" + r.URL.Path
		next.ServeHTTP(w, r)
	})
}

func checkArgs() bool {
	// TODO: _VERY_ dirty command line argument checking. This should be
	// changed into something more robust and less hacky!
	args := os.Args[1:] // drop executable
	amt := len(args)
	if amt >= 1 {
		if args[0] == "--version" {
			fmt.Printf("Outrun for Revival %s\n", meta.Version)
			return true
		}
		if args[0] == "--maintenance" {
			fmt.Println("Launching Outrun for Revival in Maintenance mode...\nOnly RPC commands will be available; no server endpoints will work!")
			MaintenanceMode = true
			return false
		}
		fmt.Println("Unknown given arguments")
		return true
	}
	return false
}

func main() {
	end := checkArgs()
	if end {
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())

	err := config.Parse("config.json")
	if err != nil {
		log.Printf("[INFO] Failure loading config file config.json (%s), using defaults\n", err)
	} else {
		log.Println("[INFO] Config file (config.json) loaded")
	}

	err = eventconf.Parse(config.CFile.EventConfigFilename)
	if err != nil {
		if !config.CFile.SilenceEventConfigErrors {
			log.Printf("[INFO] Failure loading event config file %s (%s), using defaults\n", config.CFile.EventConfigFilename, err)
		}
	} else {
		log.Printf("[INFO] Event config file (%s) loaded\n", config.CFile.EventConfigFilename)
	}

	err = infoconf.Parse(config.CFile.InfoConfigFilename)
	if err != nil {
		if !config.CFile.SilenceInfoConfigErrors {
			log.Printf("[INFO] Failure loading info config file %s (%s), using defaults\n", config.CFile.InfoConfigFilename, err)
		}
	} else {
		log.Printf("[INFO] Info config file (%s) loaded\n", config.CFile.InfoConfigFilename)
	}

	err = gameconf.Parse(config.CFile.GameConfigFilename)
	if err != nil {
		if !config.CFile.SilenceGameConfigErrors {
			log.Printf("[INFO] Failure loading game config file %s (%s), using defaults\n", config.CFile.GameConfigFilename, err)
		}
	} else {
		log.Printf("[INFO] Game config file (%s) loaded\n", config.CFile.GameConfigFilename)
	}

	err = campaignconf.Parse(config.CFile.CampaignConfigFilename)
	if err != nil {
		if !config.CFile.SilenceCampaignConfigErrors {
			log.Printf("[INFO] Failure loading campaign config file %s (%s), using defaults\n", config.CFile.CampaignConfigFilename, err)
		}
	} else {
		log.Printf("[INFO] Campaign config file (%s) loaded\n", config.CFile.CampaignConfigFilename)
	}

	dbaccess.CheckIfDBSet() // make sure we can connect to the mysql database
	err = dbaccess.InitializeTablesIfNecessary()
	if err != nil {
		log.Printf("[WARN] Failed to initialize tables; there may be problems! (%s)\n", err)
	} else {
		leaguestarttime, leagueendtime, err := dbaccess.GetStartAndEndTimesForLeague(enums.RankingLeagueF_M, 0)
		if err != nil {
			log.Println("[INFO] Ranking league data failed to load; resetting...")
			err = dbaccess.ResetAllRankingLeagueData()
			if err != nil {
				log.Printf("[WARN] Failed to reset ranking league data; there may be problems! (%s)\n", err)
			}
		} else {
			if time.Now().UTC().Unix() > leagueendtime {
				log.Printf("[WARN] League reset time has passed! %v - %v (Now: %v) Now resetting...\n", leaguestarttime, leagueendtime, time.Now().UTC().Unix())
				excode, exmsg := logic.CalculateAndResetRunnersLeague()
				log.Printf("[INFO] Result: Error code %v: %s", excode, exmsg)
			}
		}
	}

	if config.CFile.EnableRPC {
		orpc.Start()
	}

	h := muxobj.Handle
	router := mux.NewRouter()
	router.StrictSlash(true)
	SetupShutdownHandler()
	LogExecutionTime = config.CFile.DoTimeLogging
	prefix := config.CFile.EndpointPrefix

	router.HandleFunc("/", GenericRootResponse)
	router.HandleFunc("/favicon.ico", FaviconResponse)

	if MaintenanceMode {
		router.HandleFunc(prefix+"/Login/login/", h(muxhandlers.LoginMaintenance, LogExecutionTime))
		router.PathPrefix("/").HandlerFunc(MaintenanceResponse)
	} else {
		// Login
		router.HandleFunc(prefix+"/Login/login/", h(muxhandlers.Login, LogExecutionTime))
		router.HandleFunc(prefix+"/Login/getCountry/", h(muxhandlers.GetCountry, LogExecutionTime))
		router.HandleFunc(prefix+"/Login/getVariousParameter/", h(muxhandlers.GetVariousParameter, LogExecutionTime))

		// Migration
		router.HandleFunc(prefix+"/Login/getMigrationPassword/", h(muxhandlers.GetMigrationPassword, LogExecutionTime))
		router.HandleFunc(prefix+"/Login/migration/", h(muxhandlers.Migration, LogExecutionTime))

		// Information
		router.HandleFunc(prefix+"/login/getInformation/", h(muxhandlers.GetInformation, LogExecutionTime))
		router.HandleFunc(prefix+"/login/getTicker/", h(muxhandlers.GetTicker, LogExecutionTime))

		// Player operations
		router.HandleFunc(prefix+"/Player/getPlayerState/", h(muxhandlers.GetPlayerState, LogExecutionTime))
		router.HandleFunc(prefix+"/Option/userResult/", h(muxhandlers.GetOptionUserResult, LogExecutionTime))
		router.HandleFunc(prefix+"/Player/setUserName/", h(muxhandlers.SetUsername, LogExecutionTime))

		// Character operations
		router.HandleFunc(prefix+"/Player/getCharacterState/", h(muxhandlers.GetCharacterState, LogExecutionTime))
		router.HandleFunc(prefix+"/Character/changeCharacter/", h(muxhandlers.ChangeCharacter, LogExecutionTime))
		router.HandleFunc(prefix+"/Character/unlockedCharacter/", h(muxhandlers.UnlockedCharacter, LogExecutionTime))
		router.HandleFunc(prefix+"/Character/upgradeCharacter/", h(muxhandlers.UpgradeCharacter, LogExecutionTime))

		// Chao operations
		router.HandleFunc(prefix+"/Player/getChaoState/", h(muxhandlers.GetChaoState, LogExecutionTime))
		router.HandleFunc(prefix+"/Chao/equipChao/", h(muxhandlers.EquipChao, LogExecutionTime))

		// Act operations
		router.HandleFunc(prefix+"/Game/actStart/", h(muxhandlers.ActStart, LogExecutionTime))
		router.HandleFunc(prefix+"/Game/quickActStart/", h(muxhandlers.QuickActStart, LogExecutionTime))
		router.HandleFunc(prefix+"/Game/actRetry/", h(muxhandlers.ActRetry, LogExecutionTime))
		//router.HandleFunc(prefix+"/Game/actRetryFree/", h(muxhandlers.ActRetryFree, LogExecutionTime))

		// Game results operations
		router.HandleFunc(prefix+"/Game/postGameResults/", h(muxhandlers.PostGameResults, LogExecutionTime))
		router.HandleFunc(prefix+"/Game/quickPostGameResults/", h(muxhandlers.QuickPostGameResults, LogExecutionTime))

		// Leaderboard operations
		router.HandleFunc(prefix+"/Leaderboard/getWeeklyLeaderboardOptions/", h(muxhandlers.GetWeeklyLeaderboardOptions, LogExecutionTime))
		router.HandleFunc(prefix+"/Leaderboard/getLeagueData/", h(muxhandlers.GetLeagueData, LogExecutionTime))
		router.HandleFunc(prefix+"/Leaderboard/getLeagueOperatorData/", h(muxhandlers.GetLeagueOperatorData, LogExecutionTime))
		router.HandleFunc(prefix+"/Leaderboard/getWeeklyLeaderboardEntries/", h(muxhandlers.GetWeeklyLeaderboardEntries, LogExecutionTime))

		// Message operations (gift box)
		router.HandleFunc(prefix+"/Message/getMessageList/", h(muxhandlers.GetMessageList, LogExecutionTime))
		router.HandleFunc(prefix+"/Message/getMessage/", h(muxhandlers.GetMessage, LogExecutionTime))

		// Store operations
		router.HandleFunc(prefix+"/Store/getRedstarExchangeList/", h(muxhandlers.GetRedStarExchangeList, LogExecutionTime))
		//router.HandleFunc(prefix+"/Store/getRingExchangeList/", h(muxhandlers.GetRingExchangeList, LogExecutionTime))
		router.HandleFunc(prefix+"/Store/redstarExchange/", h(muxhandlers.RedStarExchange, LogExecutionTime))
		//router.HandleFunc(prefix+"/Store/ringExchange/", h(muxhandlers.RingExchange, LogExecutionTime))
		router.HandleFunc(prefix+"/Store/setBirthday/", h(muxhandlers.SetBirthday, LogExecutionTime))

		// Mileage operations
		router.HandleFunc(prefix+"/Game/getMileageData/", h(muxhandlers.GetMileageData, LogExecutionTime))
		router.HandleFunc(prefix+"/Game/getMileageReward/", h(muxhandlers.GetMileageReward, LogExecutionTime))

		// Event operations
		router.HandleFunc(prefix+"/Event/getEventList/", h(muxhandlers.GetEventList, LogExecutionTime))
		router.HandleFunc(prefix+"/Event/getEventReward/", h(muxhandlers.GetEventReward, LogExecutionTime))
		router.HandleFunc(prefix+"/Event/getEventState/", h(muxhandlers.GetEventState, LogExecutionTime))

		// Campaign operations
		router.HandleFunc(prefix+"/Game/getCampaignList/", h(muxhandlers.GetCampaignList, LogExecutionTime))

		// Daily Challenge operations
		router.HandleFunc(prefix+"/Game/getDailyChalData/", h(muxhandlers.GetDailyChallengeData, LogExecutionTime))

		// Premium Roulette operations
		router.HandleFunc(prefix+"/Chao/getChaoWheelOptions/", h(muxhandlers.GetChaoWheelOptions, LogExecutionTime))
		router.HandleFunc(prefix+"/Chao/getPrizeChaoWheelSpin/", h(muxhandlers.GetPrizeChaoWheelSpin, LogExecutionTime))
		router.HandleFunc(prefix+"/Chao/commitChaoWheelSpin/", h(muxhandlers.CommitChaoWheelSpin, LogExecutionTime))

		// Item Roulette operations
		router.HandleFunc(prefix+"/Spin/getWheelOptions/", h(muxhandlers.GetWheelOptions, LogExecutionTime))
		router.HandleFunc(prefix+"/Spin/commitWheelSpin/", h(muxhandlers.CommitWheelSpin, LogExecutionTime))

		// Raid Boss Roulette operations
		router.HandleFunc(prefix+"/RaidbossSpin/getItemStockNum/", h(muxhandlers.GetItemStockNum, LogExecutionTime))

		// Item operations
		router.HandleFunc(prefix+"/Game/getCostList/", h(muxhandlers.GetCostList, LogExecutionTime))
		router.HandleFunc(prefix+"/Game/getFreeItemList/", h(muxhandlers.GetFreeItemList, LogExecutionTime))

		// Friend operations
		router.HandleFunc(prefix+"/Friend/getFacebookIncentive/", h(muxhandlers.GetFacebookIncentive, LogExecutionTime))

		// Login bonus
		router.HandleFunc(prefix+"/Login/loginBonus/", h(muxhandlers.LoginBonus, LogExecutionTime))
		router.HandleFunc(prefix+"/Login/loginBonusSelect/", h(muxhandlers.LoginBonusSelect, LogExecutionTime))

		// Battle
		router.HandleFunc(prefix+"/Battle/getDailyBattleData/", h(muxhandlers.GetDailyBattleData, LogExecutionTime))
		router.HandleFunc(prefix+"/Battle/updateDailyBattleStatus/", h(muxhandlers.UpdateDailyBattleStatus, LogExecutionTime))
		router.HandleFunc(prefix+"/Battle/resetDailyBattleMatching/", h(muxhandlers.ResetDailyBattleMatching, LogExecutionTime))
		router.HandleFunc(prefix+"/Battle/getDailyBattleDataHistory/", h(muxhandlers.GetDailyBattleHistory, LogExecutionTime))
		router.HandleFunc(prefix+"/Battle/getDailyBattleStatus/", h(muxhandlers.GetDailyBattleStatus, LogExecutionTime))
		router.HandleFunc(prefix+"/Battle/postDailyBattleResult/", h(muxhandlers.PostDailyBattleResult, LogExecutionTime))
		router.HandleFunc(prefix+"/Battle/getPrizeDailyBattle/", h(muxhandlers.GetPrizeDailyBattle, LogExecutionTime))

		// Misc. operations
		router.HandleFunc(prefix+"/Sgn/sendApollo/", h(muxhandlers.SendApollo, LogExecutionTime))

		// Server information
		if config.CFile.EnablePublicStats {
			router.HandleFunc("/outrunInfo/stats", inforeporters.Stats)
		}

		if config.CFile.LogUnknownRequests {
			router.PathPrefix("/").HandlerFunc(OutputUnknownRequest)
		} else {
			router.PathPrefix("/").HandlerFunc(HandleUnknownRequest)
		}
	}

	go bgtasks.TouchAnalyticsDB()

	port := config.CFile.Port
	log.Printf("Starting server on port %s\n", port)
	panic(http.ListenAndServe(":"+port, removePrependingSlashes(router)))
}

func SetupShutdownHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nShutting down...")

		dbaccess.CloseDB()
		boltdbaccess.CloseDB()
		os.Exit(0)
	}()
}
