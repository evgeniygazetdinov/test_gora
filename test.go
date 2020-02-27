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

type Image struct{
  id int
  path string
}






func show_all(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		database, _ := sql.Open("sqlite3", "./images.db")
		statement, _  := database.Prepare("CREATE TABLE IF NOT EXISTS images(id INTEGER PRIMARY KEY AUTOINCREMENT, path TEXT)")
		statement.Exec()
		rows, _ := database.Query("SELECT id, path from images")
		var id int
		var all_images []Image
		var path string
		for rows.Next(){
	 	rows.Scan(&id, &path)
		append(all_images, Image{id, path})
		}
			w.Write([]byte(`{"message":1)`))
	// } else {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	 w.Header().Set("Content-Type", "application/json")
	// 	 w.Write([]byte(`{"message": "method not allowed"}`))
	// }
}






func main() {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", show_all)
		log.Fatal(http.ListenAndServe(":8080", router))

}
