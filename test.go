package main
//each and every file in go must have a package name
import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
	"fmt"
)
func main() {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/{id}", show_all)
		log.Fatal(http.ListenAndServe(":8080", router))

}





func show_all(w http.ResponseWriter, r *http.Request) {

		if r.Method =="GET" {
			database, _ := sql.Open("sqllite3", "images.db")
			rows,_ := database.Query("SELECT * images")
			fmt.Println(rows)


			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": 4`))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		 w.Header().Set("Content-Type", "application/json")
		 w.Write([]byte(`{"message": "method not allowed"}`))
	}


}
