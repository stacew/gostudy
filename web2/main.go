package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadfile, header, err := r.FormFile("upload_file")
	defer uploadfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	io.Copy(file, uploadfile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)

}

func main() {
	path := "C:/a.png"
	if _, err := os.Stat(path); os.IsExist(err) {
	}

	_, err := os.Stat(path)
	if err != nil {

	}
	_, err2 := os.Stat(path)
	if err2 != nil {

	}

	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
