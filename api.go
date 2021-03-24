/*
	20 march 2021
	go run api.go
*/

package main

import (
	c "Project/config"
	e "Project/endpoint"
	f "Project/functions"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)


func main(){

	conf := loadEnv()
	handleReq(conf)
}

func handleReq(conf c.Configuration){
	//endpoints
	router := mux.NewRouter()
	router.HandleFunc("/", e.Index{}.SendResponse).Methods("GET")
	router.HandleFunc("/whoami", e.Whoami{}.SendResponse).Methods("GET")
	router.HandleFunc("/alert", func (w http.ResponseWriter, r *http.Request) {
		e.Alert{}.SendResponse(w,r,conf)
	})
	
	//about port
	if os.Getenv("PORT")==""{
		fmt.Println("Listening...",conf.Server.Host+conf.Server.Port)
		log.Fatal(http.ListenAndServe(conf.Server.Host+conf.Server.Port, router))	
	}else{
		fmt.Println("Listening...",conf.Server.Host+os.Getenv("PORT"))
		log.Fatal(http.ListenAndServe(conf.Server.Host+os.Getenv("PORT"), router))	
	}	
}

//for loading env. variables
func loadEnv() (conf c.Configuration){
	v := viper.New()
	v.SetConfigFile("config.yml")

	err := v.ReadInConfig()
	f.ErrCheck(err, nil, "Problem with reading config process")

	err = v.Unmarshal(&conf)
	f.ErrCheck(err,nil,"Problem on decoding into struct")

	fmt.Println(conf.Environment.DUMMY_WEBHOOK_URL)

	return
}








