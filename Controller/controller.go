// Package controller Student API.
//
// the purpose of this application is to provide an application
// that is using go code to define an  Rest API
//
//     Schemes: http, https
//     Host:student.com:31000
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package Controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Configuration struct
type Configuration struct {
	Port             string // port no
	ConnectionString string // connection string
	Database         string // database name
	Collection       string // collection
}

/*ReadConfig Reading the configs from  db.properties
 */
func ReadConfig() Configuration {
	var configfile = "config.properties"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Configuration
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}

func HandleRequests() {

	config := ReadConfig()
	var port = ":" + config.Port

	Router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "api_key", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT"})

	Router.HandleFunc("/", Homepage).Methods(http.MethodGet)
	Router.HandleFunc("/students", Poststudent).Methods(http.MethodPost)
	Router.HandleFunc("/students", GetStudent).Methods(http.MethodGet)
	Router.HandleFunc("/students", GetStudentbyName).Methods(http.MethodGet)
	Router.HandleFunc("/students/{id}", GetStudentbyid).Methods(http.MethodGet)
	Router.HandleFunc("/students/{id}", Deletestudent).Methods(http.MethodDelete)
	Router.HandleFunc("/students/{id}", Updatestudent).Methods(http.MethodPut)
	Router.HandleFunc("/students/{id}", Updatestudent_patch).Methods(http.MethodPatch)
	// log.Fatal(http.ListenAndServe(":8091", Router))
	fmt.Printf("application listening port%s\n", port)
	http.ListenAndServe(port, handlers.CORS(corsObj, methodsOk, headersOk)(Router))

}
