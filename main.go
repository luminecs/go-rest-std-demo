package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &RecipesHandler{})
	mux.Handle("/recipes/", &RecipesHandler{})
	//mux.Handle("/picture.jpg", http.FileServer(http.Dir("/home/lumine/图片/ckxx.jpg")))

	// 定义处理 /picture.jpg 请求的处理器
	mux.HandleFunc("/picture.jpg", func(w http.ResponseWriter, r *http.Request) {
		imgPath := "/home/lumine/图片/ckxx.png"
		f, err := os.Open(imgPath)
		if err != nil {
			http.Error(w, "Image not found", http.StatusNotFound)
			log.Printf("Error opening image: %v", err)
			return
		}
		defer f.Close()

		fileInfo, err := f.Stat()
		if err != nil {
			http.Error(w, "Error getting file info", http.StatusInternalServerError)
			log.Printf("Error getting file info: %v", err)
			return
		}

		buffer := make([]byte, 512)
		f.Read(buffer)
		contentType := http.DetectContentType(buffer)
		w.Header().Set("Content-Type", contentType)

		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), f)
	})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct {
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

type RecipesHandler struct {
}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my recipe page"))
}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request)  {}
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request)    {}
func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)
