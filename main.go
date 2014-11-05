package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/jaijiv/go-social-api/bootstrapper"
	"github.com/jaijiv/go-social-api/database"
	"github.com/jaijiv/go-social-api/environment"
	"github.com/jaijiv/go-social-api/middleware"

	"github.com/adampresley/sigint"
	"github.com/justinas/alice"
)

var serverPort = flag.Int("serverport", 8081, "Port for web server")
var serverAddress = flag.String("serveraddress", "localhost", "Address for web server")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
	 * Handle SIGINT (CTRL+C)
	 */
	sigint.ListenForSIGINT(func() {
		log.Println("Shutting down...")
		os.Exit(0)
	})

	/*
	 * Setup database
	 */

	connectionInfo, err := environment.GetDBConnectionInformation()
	if err != nil {
		log.Fatal(err)
	}

	err = database.ConnectMongo(connectionInfo)
	if err != nil {
		log.Fatal(err)
	}

	/*
	 * Setup routing and middleware
	 */
	router := bootstrapper.SetupWebRouter()
	middleware := alice.New(middleware.Logger, middleware.Auth).Then(router)

	/*
	 * Start web server
	 */
	log.Printf("go-social-api services started on %s:%d\n\n", *serverAddress, *serverPort)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *serverAddress, *serverPort), middleware)
}
