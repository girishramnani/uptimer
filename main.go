package main

import (
	"log"
	"time"

	"net/http"

	"github.com/girishramnani/uptimer/pkg/cache"
	"github.com/girishramnani/uptimer/pkg/database"
	"github.com/girishramnani/uptimer/pkg/req"
	"github.com/girishramnani/uptimer/pkg/types"
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func baseRouter(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	v1.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"health": "OK",
		})
	})
	return v1
}

func initDB(dbm *database.DBManager) {
	if err := dbm.CreateSchema(&types.ServiceResp{}); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	urlList := []string{
		"https://youtube.com",
		"https://facebook.com",
		"https://baidu.com",
		"https://wikipedia.org",
		"https://taobao.com",
		"https://yahoo.com",
		"https://tmall.com",
		"https://amazon.com",
		"https://twitter.com",
		"https://livet.com",
	}
	sc := cache.NewServiceCache()
	dbm := database.NewDBManager(&pg.Options{
		User:     "postgres",
		Password: "postgres",
	})
	initDB(dbm)

	getAndInsert := func() {
		log.Println("getting and inserting urls")
		resps := req.GetAllUrls(urlList)
		for resp := range resps {
			sc.Set(resp.URL, resp.RespCode)
			err := dbm.Insert(&resp)
			if err != nil {
				log.Printf("Error inserting %s: %s\n", resp.URL, err.Error())
			}
		}
	}

	go func() {
		getAndInsert()
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			getAndInsert()
		}

	}()

	// the backend
	e := echo.New()
	v1 := baseRouter(e)
	v1.GET("/services", func(c echo.Context) error {
		return c.JSON(http.StatusOK, sc.ToReadableMap())
	})

	log.Println("Listening on 8080")
	err := e.Start(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
