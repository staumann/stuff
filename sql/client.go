package sql

import (
	"database/sql"
	"fmt"
	"github.com/staumann/caluclation/database"
	"io/ioutil"
	"log"
	"os"
	"strings"
	// Import driver only works with blank import
	_ "github.com/go-sql-driver/mysql"
)

const (
	host      = "127.0.0.1"
	port      = 3306
	user      = "root"
	password  = "root"
	dbname    = "billing"
	scriptDir = "sql/scripts"
	initDir   = "init"
)

var adapter *Adapter

type Adapter struct {
	scriptCache map[string]string
	db          *sql.DB
}

func GetBillRepository() database.BillRepository {
	return &billRepo{adapter: getAdapter()}
}

func GetUserRepository() database.UserRepository {
	return &userRepo{adapter: getAdapter()}
}

func GetPositionRepository() database.PositionRepository {
	return &positionRepo{adapter: getAdapter()}
}

func getAdapter() *Adapter {
	if adapter == nil {
		adapter = &Adapter{}
		adapter.init()
	}
	return adapter
}

func (a *Adapter) init() {
	sqlDnsString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname)
	log.Print(sqlDnsString)
	if db, err := sql.Open("mysql", sqlDnsString); err != nil {
		log.Fatalf("could not open connection to database: %s", err.Error())
	} else {
		a.scriptCache = make(map[string]string)
		a.db = db
		files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", scriptDir, initDir))

		for _, f := range files {
			_, err = a.db.Exec(a.getScript(fmt.Sprintf("%s/%s", initDir, f.Name())))
			if err != nil {
				log.Fatalf("unable to execute %s script: %s", f.Name(), err)
			}
		}
	}
}

func (a *Adapter) getScript(scriptName string) string {
	sanitizedScriptName := strings.Replace(scriptName, ".sql", "", -1)
	if v, ok := a.scriptCache[sanitizedScriptName]; ok {
		return v
	} else {
		file, err := os.Open(fmt.Sprintf("%s/%s.sql", scriptDir, sanitizedScriptName))
		if err != nil {
			log.Printf("error getting script %s. Cause: %s", sanitizedScriptName, err.Error())
			return ""
		}
		bts, _ := ioutil.ReadAll(file)
		script := string(bts)
		a.scriptCache[sanitizedScriptName] = script
		return script
	}
}
