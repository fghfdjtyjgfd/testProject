package main

import (
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	"github.com/jarbza/backend-api/repository"
	"github.com/jarbza/backend-api/service"
	"github.com/jarbza/logx"
	"github.com/jarbza/middleware"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func init() {

	runtime.GOMAXPROCS(1)
	initViper()
}
func main() {

	var (
		m             = middleware.New(viper.GetString("app.name"))
		g             = initGin(m)
		database      = newMariaConn()
		userReposiory = repository.New(database)
	)
	defer database.Close()

	services := service.NewHandler(service.NewService(userReposiory))

	c := g.Group("/api/jarb-service/v1/login")
	c.POST("/", services.Login)
	c.POST("/status", services.GetStatus)

	history := g.Group("/api/jarb-service/v1/history")
	history.GET("/login", services.GetLoginHistory)

	g.Run(":" + viper.GetString("app.port"))
}

func initGin(m *middleware.Middleware) *gin.Engine {

	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(m.LogRequestInfo())
	g.Use(m.LogResponseInfo())
	g.Use(cors.Default())

	return g
}

func initViper() {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

}

func newMariaConn() *gorm.DB {
	mysql.SetLogger(logx.Logger)
	conf := mysql.Config{
		DBName:               viper.GetString("mysql.database"),
		User:                 viper.GetString("mysql.username"),
		Passwd:               viper.GetString("mysql.password"),
		Net:                  "tcp",
		Addr:                 viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port"),
		AllowNativePasswords: true,
		Timeout:              viper.GetDuration("mysql.timeout"),
		ReadTimeout:          viper.GetDuration("mysql.readtimeout"),
		WriteTimeout:         viper.GetDuration("mysql.writetimeout"),
		ParseTime:            viper.GetBool("mysql.parsetime"),
		Loc:                  time.Local,
	}

	logx.Info("[CONFIG] repositories connection: ", strings.ReplaceAll(conf.FormatDSN(), conf.Passwd, "********"))

	conn, err := gorm.Open("mysql", conf.FormatDSN())
	if err != nil {
		logx.Fatalf("cannot open mysql connection:%s", err)
	}

	conn.SetLogger(logx.Logger)
	conn.DB().SetMaxIdleConns(viper.GetInt("mysql.maxidle"))
	conn.DB().SetMaxOpenConns(viper.GetInt("mysql.maxopen"))
	conn.DB().SetConnMaxLifetime(viper.GetDuration("mysql.maxlifetime"))
	conn.LogMode(viper.GetBool("mysql.debug"))

	if viper.GetBool("mysql.migrate") {
		logx.Info("Performing DB Migration")
	}

	return conn
}
