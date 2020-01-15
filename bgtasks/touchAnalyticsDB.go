package bgtasks

import (
	"log"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db/boltdbaccess"
)

func TouchAnalyticsDB() {
	err := boltdbaccess.Set(consts.DBMySQLTableAnalytics, "touch", []byte{})
	if err != nil {
		log.Println("[ERR] Unable to touch " + consts.DBBucketAnalytics + ": " + err.Error())
	}
}
