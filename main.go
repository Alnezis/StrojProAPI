package main

import (
	"StrojProAPI/api"
	"StrojProAPI/app"
	"StrojProAPI/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	//object.ParseFILE()
	_cors := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	router := mux.NewRouter()

	router.HandleFunc("/objects/getStagesByObjName", controllers.GetStagesByObjName).Methods("GET")
	router.HandleFunc("/objects/getObjList", controllers.GetObjList).Methods("GET")
	router.HandleFunc("/objects/getSubStageListByObjNameStage", controllers.GetSubStageListByObjNameStage).Methods("GET")

	router.HandleFunc("/uploadImage", controllers.UploadFile).Methods("POST")
	router.PathPrefix("/images").Handler(http.StripPrefix("/images", http.FileServer(http.Dir("images/")))).Methods("GET")

	//router.HandleFunc("/user/transactions", controllers.GetUser).Methods("GET")

	//	go firebase.Demon()

	cert := "/etc/letsencrypt/live/alnezis.riznex.ru/fullchain.pem"
	key := "/etc/letsencrypt/live/alnezis.riznex.ru/privkey.pem"
	if _, err := os.Stat(cert); err != nil {
		if os.IsNotExist(err) {
			log.Println("no ssl")
			handler := _cors.Handler(router)
			err := http.ListenAndServe(fmt.Sprintf(":%d", app.CFG.Port), handler)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
	log.Println("yes ssl")
	handler := _cors.Handler(router)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", app.CFG.Port), cert, key, handler)
	if err != nil {
		api.CheckErrInfo(err, "ListenAndServeTLS")
		//	return
	}
}
