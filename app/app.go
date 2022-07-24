package app

import (
	"StrojProAPI/api"
	//"github.com/mailgun/mailgun-go/v4"
	//"github.com/mailgun/mailgun-go/v4"
	"StrojProAPI/app/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var CFG *config.Config
var DB *sqlx.DB

func init() {
	CFG = config.InitCfg()

	conn := `
           host=` + CFG.Db.Host + `
         dbname=` + CFG.Db.DbName + `
		   user=` + CFG.Db.UserName + `
        sslmode=disable
		   port=` + CFG.Db.Port + `
		password=` + CFG.Db.Password + `
`
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	DB = db
	//	initDb()

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS objects
(
    id       serial primary key,
    user_id integer default 0,
    
    object_name    varchar,
    stage_name    varchar,
    sub_stage_name    varchar,
    period    integer,
    percent_object    numeric(18,2),
    percent_stage numeric(18,2),
    is_completed boolean default false,
    completed_time timestamp default null
);`)
	api.CheckErrInfo(err, "create table objects")
}
