package main

import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
	"fmt"
	"encoding/json"
)

func show_all(w http.ResponseWriter, r *http.Request){
			if r.Method == "GET"{
				var id int
				var path string
	 			all_images := []map[int]string{}
				result := make(map[string][]map[int]string)

				database, _ := sql.Open("sqlite3", "./images.db")
				statement, _  := database.Prepare("CREATE TABLE IF NOT EXISTS images(id INTEGER PRIMARY KEY AUTOINCREMENT, path TEXT)")
				statement.Exec()
				rows, _ := database.Query("SELECT id, path from images")

			for rows.Next(){
				rows.Scan(&id, &path)
				all_images = append(all_images, map[int]string{id: path})
			}
			result["messages"]=all_images
			 w.WriteHeader(http.StatusOK)
	 		 w.Header().Set("Content-Type", "application/json")
			 json.NewEncoder(w).Encode(result)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		 	w.Header().Set("Content-Type", "appli
				cation/json")
		 	w.Write([]byte(`{"message": "method not allowed"}`))
		 }
	}
func add_image(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST"{
			image_path, ok := r.URL.Query()["image_path"]
	    if !ok || len(image_path[0]) < 1 {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"message": "path_missed"}`))
	    }
			//TODO add check by path
			fmt.Println(image_path[0])
			database, _ := sql.Open("sqlite3", "./images.db")
			statement, _  := database.Prepare("INSERT INTO images (path) VALUES(?)")
			statement.Exec(image_path[0])
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "path_added}`))
		}else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		 	w.Header().Set("Content-Type", "application/json")
		 	w.Write([]byte(`{"message": "method not allowed"}`))
		 }

}
func delete_image_by_id(w http.ResponseWriter, r *http.Request){
		if r.Method == "POST"{
			image_id, ok := r.URL.Query()["id"]
	    if !ok || len(image_id[0]) < 1 {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"message": "id_missed"}`))
	    }

			fmt.Println(image_id[0])
			database, _ := sql.Open("sqlite3", "./images.db")
			statement, _  := database.Prepare("DELETE FROM images WHERE id = (?)")
			statement.Exec(image_id[0])
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "path_deleted}`))
		} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			 	w.Header().Set("Content-Type", "application/json")
			 	w.Write([]byte(`{"message": "method not allowed"}`))
			 }

}

func main() {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", show_all)
		router.HandleFunc("/add_image/",add_image)
		router.HandleFunc("/delete_image/",delete_image_by_id)
		log.Fatal(http.ListenAndServe(":8080", router))
}
