package main

import (
	"net/http"
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
