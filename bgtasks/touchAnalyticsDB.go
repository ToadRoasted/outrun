package bgtasks

import (
	"log"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db/dbaccess"
)

func TouchAnalyticsDB() {
	err := dbaccess.SetAnalyticsEntry(consts.DBMySQLTableAnalytics, "touch", []byte{})
	if err != nil {
		log.Println("[ERR] Unable to touch " + consts.DBMySQLTableAnalytics + ": " + err.Error())
	}
}
