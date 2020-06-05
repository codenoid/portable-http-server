package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	var port string
	flag.StringVar(&port, "port", "3000", "-port <wanted-port>")
	flag.Parse()

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20) // limit your max input length!

		// in your case file would be fileupload
		file, header, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		fmt.Printf("[UPLOAD] From %v File name %s\n", r.RemoteAddr, name[0])
		// Copy the file data to disk
		f, err := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		io.Copy(f, file)
		// do something else
		// etc write header
		w.Write([]byte("done"))
		return
	})

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening on :" + port + "...")
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
