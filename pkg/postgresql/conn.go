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

	var totalMemoryMB string
	err = db.QueryRow(`select total_memory/1024/1024 as total_memory_MB from pg_sys_memory_info();`).Scan(&totalMemoryMB)
	if err != nil {
		log.Fatal(err)
	}
	ret["total_memory_mb"] = totalMemoryMB

	var usedMemoryMB string
	err = db.QueryRow(`select used_memory/1024/1024 as used_memory_MB from pg_sys_memory_info();`).Scan(&usedMemoryMB)
	if err != nil {
		log.Fatal(err)
	}
	ret["used_memory_mb"] = usedMemoryMB

	var freeMemoryMB string
	err = db.QueryRow(`select free_memory/1024/1024 as free_memory_MB from pg_sys_memory_info();`).Scan(&freeMemoryMB)
	if err != nil {
		log.Fatal(err)
	}
	ret["free_memory_mb"] = freeMemoryMB

	var swapMemoryMB string
	err = db.QueryRow(`select swap_used/1024/1024 as swap_used_MB from pg_sys_memory_info();`).Scan(&swapMemoryMB)
	if err != nil {
		log.Fatal(err)
	}
	ret["swap_used_mb"] = swapMemoryMB

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
