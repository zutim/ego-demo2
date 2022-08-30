package main

import (
	"ego-demo2/server/route"
	"fmt"
	"github.com/zutim/ego"
	"github.com/zutim/ego/component"
	"gorm.io/driver/mysql"
	gorm2 "gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func main() {
	err := component.Provider().Config().LoadFile("configs/app.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	Run()
}

func InitMysql() error {
	err := component.Provider().Gorm().OpenMySQL(component.Provider().Config().Viper.GetString("mysql.dsn"))
	if err != nil {
		return err
	}

	otherDsn := component.Provider().Config().Viper.GetString("mysql.dsn2")
	resolverErr := component.Provider().Gorm().RegisterResolverConfig(
		dbresolver.Config{
			Sources: []gorm2.Dialector{mysql.Open(otherDsn)},
		}, "tbl_gifts", "tbl_gift_types")
	return resolverErr
}

func Run() {
	aggregator := ego.NewAggregatorServer()

	if err := InitMysql(); err != nil {
		fmt.Println(err)
		return
	}

	// 实例化一个http服务
	httpServer := ego.NewHTTPServer(component.Provider().Config().GetString("http.port")).
		RegisterRouteLoader(
			route.Loader,
		)

	aggregator.WithServer(httpServer)

	aggregator.Run()
}

