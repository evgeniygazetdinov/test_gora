package main
//each and every file in go must have a package name
import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
	"fmt"
	"encoding/json"
)

type Image struct{
  id int
  path string
}


func show_all(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET"{
			var id int
 			all_images := []map[int]string{}
			var path string

			database, _ := sql.Open("sqlite3", "./images.db")
			statement, _  := database.Prepare("CREATE TABLE IF NOT EXISTS images(id INTEGER PRIMARY KEY AUTOINCREMENT, path TEXT)")
			statement.Exec()
			rows, _ := database.Query("SELECT id, path from images")

		for rows.Next(){
			rows.Scan(&id, &path)
			fmt.Println(id,path)
			all_images = append(all_images, map[int]string{id: path})
		}

		  //js,err := json.Marshall(all_images)
			// json.NewEncoder(w).Encode(all_images)
		 w.WriteHeader(http.StatusOK)
 		 w.Header().Set("Content-Type", "application/json")
		 output := make(map[string][]map)
	 	 output["message"] = all_images
		 json.NewEncoder(w).Encode(output)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	 	w.Header().Set("Content-Type", "application/json")
	 	w.Write([]byte(`{"message": "method not allowed"}`))
	 }
}








func main() {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", show_all)
		// router.HandleFunc("/add_image",add_image)
		log.Fatal(http.ListenAndServe(":8080", router))

}
