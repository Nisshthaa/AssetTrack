package middlewares

import (
	"log"
	"net/http"
)

func CommonMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		log.Println(r.Method, r.URL.Path)

		defer func(){
			if err:=recover();err!=nil{
				http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
