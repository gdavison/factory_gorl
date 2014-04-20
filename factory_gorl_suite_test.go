package factory_gorl_test

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/gdavison/factory_gorl"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

var (
	dbMap *gorp.DbMap
)

func TestFactory_gorl(t *testing.T) {
	RegisterFailHandler(Fail)

	dbMap := initDb()
	defer dbMap.Db.Close()

	factory_gorl.DbMap = dbMap

	RunSpecs(t, "Factory_gorl Suite")
}

func initDb() *gorp.DbMap {
	dbFile := "./tmp/foo.db"
	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTable(AutoIncrementIdTest{}).SetKeys(true, "Id")

	if err := dbmap.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}

	return dbmap
}
