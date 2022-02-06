package main

import (
	"log"
	"os"

	"github.com/yasir2000/my-hexagonal-go/internal/adaptors/app/api"
	arithamtic "github.com/yasir2000/my-hexagonal-go/internal/adaptors/core/arithmatic"
	gRPC "github.com/yasir2000/my-hexagonal-go/internal/framework/left/grpc"
	"github.com/yasir2000/my-hexagonal-go/internal/framework/right/db"
	"github.com/yasir2000/my-hexagonal-go/internal/ports"
)

// var (
// 	Client   *sql.DB
// 	username = "root"
// 	password = "root"
// 	host     = "127.0.0.1"
// 	schema   = "users_db_01"
// )

func main() {
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	var err error
	//ports
	var dbaseAdaptor ports.DbPort
	var core ports.ArithmaticPort
	var appAdaptor ports.APIPort
	var gRPCAdaptor ports.GRPCPort

	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	dbaseAdaptor, err = db.NewAdaptor(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatal("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdaptor.CloseDbConnection()

	core = arithamtic.NewAdaptor()
	appAdaptor = api.NewApplication(dbaseAdaptor, core)
	gRPCAdaptor = gRPC.NewAdaptor(appAdaptor)
	gRPCAdaptor.Run()

	// creation of type adaptor which has access to all methods
	//arithAdaptor := arithmatic.NewAdaptor()

}
