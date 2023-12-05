package addons

import (
	"database/sql"
	"fmt"
	"path"

	"github.com/cgrates/cgrates/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	cfgCgr, _ = config.NewCGRConfigFromPath(path.Join("/etc", "cgrates"))
	host      = cfgCgr.StorDbCfg().Host
	port      = cfgCgr.StorDbCfg().Port
	user      = cfgCgr.StorDbCfg().User
	password  = cfgCgr.StorDbCfg().Password
	dbName    = "addons" //cfgCgr.StorDbCfg().Name
	location  = "Local"
)

func StordbConnect() {
	// cfgCgr, _ := config.NewCGRConfigFromPath(path.Join("/etc", "cgrates"))
	host := cfgCgr.StorDbCfg().User
	pwd := cfgCgr.StorDbCfg().Password
	fmt.Println(host, pwd)
}

func mkDBConn() (cDB *sql.DB, err error) {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true&sql_mode='ALLOW_INVALID_DATES'", user, password, host, port, dbName, location)
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{AllowGlobalUpdate: true})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// defer db.Close()
	cDB, err = db.DB()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if cDB.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cDB, nil
}

func runQuery(qryStr string) (qOut []string, err error) {

	cDB, err := mkDBConn()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := cDB.Query(qryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	ids := make([]string, 0)
	i := 0
	for rows.Next() {
		i++
		var id string
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ids = append(ids, id)
	}
	qOut = ids
	return qOut, nil
}

func MySqlConnect() {

	var rows *sql.Rows
	var err error
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true&sql_mode='ALLOW_INVALID_DATES'", user, password, host, port, dbName, location)
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{AllowGlobalUpdate: true})
	if err != nil {
		fmt.Println(err)
	}
	// defer db.Close()
	cDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	if cDB.Ping(); err != nil {
		fmt.Println(err)
	}
	qryStr := "show tables"
	rows, err = cDB.Query(qryStr)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	ids := make([]string, 0)
	i := 0
	for rows.Next() {
		i++ //Keep here a reference so we know we got at least one
		var id string
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
		ids = append(ids, id)
	}
	fmt.Println(ids, err)
}
