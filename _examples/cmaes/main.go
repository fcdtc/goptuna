package main

import (
	"github.com/c-bata/goptuna"
	"github.com/c-bata/goptuna/cmaes"
	"github.com/c-bata/goptuna/dashboard"
	"github.com/c-bata/goptuna/rdb.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math"
	"net/http"
	"os"
)

func objective(trial goptuna.Trial) (float64, error) {
	x1, _ := trial.SuggestFloat("x1", -10, 10)
	x2, _ := trial.SuggestFloat("x2", -10, 10)
	return math.Pow(x1-2, 2) + math.Pow(x2+5, 2), nil
}

func main() {

	var db *gorm.DB
	var err error

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

	if err != nil {
		log.Fatal("failed to open db:", err)
	}
	err = rdb.RunAutoMigrate(db)
	if err != nil {
		log.Fatal("failed to run auto migrate:", err)
	}

	relativeSampler := cmaes.NewSampler(
		cmaes.SamplerOptionNStartupTrials(5))
	study, err := goptuna.CreateStudy(
		"goptuna-example",
		goptuna.StudyOptionStorage(rdb.NewStorage(db)),
		goptuna.StudyOptionRelativeSampler(relativeSampler),
	)
	if err != nil {
		log.Fatal("failed to create study:", err)
	}

	go runDashboard()
	if err = study.Optimize(objective, 20000); err != nil {
		log.Fatal("failed to optimize:", err)
	}

	v, _ := study.GetBestValue()
	params, _ := study.GetBestParams()
	log.Printf("Best evaluation=%f (x1=%f, x2=%f)",
		v, params["x1"].(float64), params["x2"].(float64))


}

func runDashboard()  {
	// 可以用goptuna封装的函数
	//storageURL := "sqlite:///test.db"
	//db2, err := sqlalchemy.GetGormDBFromURL(storageURL, nil)
	//if err != nil {
	//	os.Exit(1)
	//}

	//if db2.Dialector.Name() == "sqlite" {
	//	err = db2.Exec("PRAGMA foreign_keys = ON").Error
	//	if err != nil {
	//		os.Exit(1)
	//	}
	//}

	// 也可以不用他提供的url，自己写
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	storage := rdb.NewStorage(db)
	server, err := dashboard.NewServer(storage)
	if err != nil {
		os.Exit(1)
	}
	if err := http.ListenAndServe("127.0.0.1:8000", server); err != nil {
		os.Exit(1)
	}
}