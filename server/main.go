package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000" , nil)
}
func home(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	log.Println("iniciou a minha request")
	defer log.Println("finalizou a minha request")

	select {
	case <-time.After(time.Second * 5):
		fmt.Fprint(w, "requisicao, processada com sucesso")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("requisicao processada com sucesso"))
	case <- ctx.Done():
			log.Println("request cancelada")
			http.Error(w , " request cancelada" , http.StatusRequestTimeout)
	}

}