package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

// Post represents a blog post
type Post struct {
	ID       string
	Title    string
	Content  string
	HTML     template.HTML
	FilePath string
}

// MarkdownToHTML converts the Markdown content to HTML
func MarkdownToHTML(content string) template.HTML {
	html := blackfriday.Run([]byte(content))
	return template.HTML(html)
}

// IndexHandler handles the homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts := []Post{}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			filePath := fmt.Sprintf("./posts/%s", file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := Post{
				ID:       file.Name(),
				Title:    strings.TrimSuffix(file.Name(), ".md"),
				Content:  string(content),
				HTML:     MarkdownToHTML(string(content)),
				FilePath: filePath,
			}
			posts = append(posts, post)
		}
	}

	// Check if a post was just deleted
	deletedID := r.URL.Query().Get("deleted")
	if deletedID != "" {
		// Remove the deleted post from the slice
		for i, post := range posts {
			if post.ID == deletedID {
				posts = append(posts[:i], posts[i+1:]...)
				break
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// NewPostHandler handles creating a new blog post
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreatePostHandler handles the creation of a new blog post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	fileName := fmt.Sprintf("%s.md", title)
	filePath := fmt.Sprintf("./posts/%s", fileName)

	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// EditPostHandler handles editing an existing blog post
func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filePath := fmt.Sprintf("./posts/%s", id)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := Post{
		ID:       id,
		Title:    id,
		Content:  string(content),
		HTML:     MarkdownToHTML(string(content)),
		FilePath: filePath,
	}

	tmpl, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdatePostHandler handles updating an existing blog post
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filePath := fmt.Sprintf("./posts/%s", id)

	content := r.FormValue("content")
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// DeletePostHandler handles deleting an existing blog post
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filePath := fmt.Sprintf("./posts/%s", id)

	err := os.Remove(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the homepage and pass the deleted post ID as a query parameter
	http.Redirect(w, r, "/?deleted="+id, http.StatusFound)
}

func main() {
	// Create a router using gorilla/mux
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/posts/new", NewPostHandler).Methods("GET")
	r.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	r.HandleFunc("/posts/{id}/edit", EditPostHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", UpdatePostHandler).Methods("POST")
	r.HandleFunc("/posts/{id}", DeletePostHandler).Methods("POST")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
