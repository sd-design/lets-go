package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
)

type Row struct {
	ID       int
	Filename string
}
type List struct {
	FileList []Row
}

func countGoFiles(folder string) (int, []*Row) {
	Error := "Error"
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Printf(Error)
	}
	//Struct1 := List{}
	//Struct2 := Row{}
	count := 0
	FileList := []*Row{}
	for _, f := range files {
		if path.Ext(f.Name()) == ".sql" {
			//fmt.Println(f.Name())
			Struct2 := new(Row)
			Struct2.ID = count
			Struct2.Filename = f.Name()
			FileList = append(FileList, Struct2)
			//Struct1.FileList = append(Struct1.FileList, Struct2)
			count++
		}
	}
	//fmt.Printf("%+v", Struct1)
	return count, FileList
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	folder := "/Volumes/ARCHIVE/PRIVATE/WS/github/lets-go/db_backups"
	counter, Struct := countGoFiles(folder)
	log.Println(counter)
	//fmt.Printf("%+v", Struct)
	data, err := json.Marshal(Struct)
	if err != nil {
		return
	}
	w.Write([]byte(data))
}

func main() {

	//countGoFiles(folder)
	//fmt.Printf("Count of files: %d\n", count)
	http.HandleFunc("/", home)
	log.Println("Starting server on :4200")
	err := http.ListenAndServe(":4200", nil)
	log.Fatal(err)
}
