package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	"go-logstash/config"

	_ "github.com/lib/pq"
)

func ConnectPG() (*sql.DB, error) {
	//Open db
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Conf.Host, config.Conf.Port, config.Conf.User, config.Conf.Password,
		config.Conf.DBName, config.Conf.SSLMode)
	db, err := sql.Open("postgres", psqlconn)
	return db, err
}

func GetSysInfo() (ret map[string]string) {
	db, err := ConnectPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ret = make(map[string]string)
	ret["instance"] = "pg-sys"

	var totalMemory string
	err = db.QueryRow(`select total from memory.f_memory();`).Scan(&totalMemory)
	if err != nil {
		log.Fatal(err)
	}
	ret["total_memory"] = totalMemory

	var usedMemory string
	err = db.QueryRow(`select used from memory.f_memory();`).Scan(&usedMemory)
	if err != nil {
		log.Fatal(err)
	}
	ret["used_memory"] = usedMemory

	var freeMemory string
	err = db.QueryRow(`select free from memory.f_memory();`).Scan(&freeMemory)
	if err != nil {
		log.Fatal(err)
	}
	ret["free_memory"] = freeMemory

	var loadAvgOneMinute string
	err = db.QueryRow(`select load_avg_one_minute from pg_sys_load_avg_info();`).Scan(&loadAvgOneMinute)
	if err != nil {
		log.Fatal(err)
	}
	ret["load_avg_one_minute"] = loadAvgOneMinute

	var loadAvgFiveMinutes string
	err = db.QueryRow(`select load_avg_five_minutes from pg_sys_load_avg_info();`).Scan(&loadAvgFiveMinutes)
	if err != nil {
		log.Fatal(err)
	}
	ret["load_avg_five_minutes"] = loadAvgFiveMinutes

	var loadAvgTenMinutes string
	err = db.QueryRow(`select load_avg_ten_minutes from pg_sys_load_avg_info();`).Scan(&loadAvgTenMinutes)
	if err != nil {
		log.Fatal(err)
	}
	ret["load_avg_ten_minutes"] = loadAvgTenMinutes

	return ret
}
