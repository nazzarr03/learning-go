package fileupload

import (
	"fmt"
	"log"
	"io"
	"net/http"
	"os"
	"text/template"
)

// Compile templates on start of the application
var templates = template.Must(template.ParseFiles("./upload.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // limit your max input length!

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)
	fmt.Fprintf(w, "%v", handler.Filename)
	fmt.Fprintf(w, "%v", handler.Size)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem	
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	log.Println("Listening on http://localhost:8080/")	
	http.ListenAndServe(":8080", nil)
}