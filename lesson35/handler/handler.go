package handler

import (
	// "encoding/json"
	"net/http"

	// "github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/Go11Group/at_lesson/lesson34/postgres"
)

type user struct{
	ID       int 
    Username string 
    Email    string 
}
type Problem struct{
	ID          int
	Title       string
	Description string
	Difficulty  string
}

type SolvedProblem struct{
	ID        int 
    UserID    int 
    ProblemID int 
}

type Handler struct {
	user          *postgres.UserRepo
	Problem       *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	// User handlers
	mux.HandleFunc("/user/", handler.User)
	mux.HandleFunc("/user/create", handler.UserCreate)
	mux.HandleFunc("/user/update/", handler.UserUpdate)
	mux.HandleFunc("/user/delete/", handler.Userdelete)

	// Problem handlers
	mux.HandleFunc("/problem/", handler.Problem1)
	mux.HandleFunc("/problem/create", handler.ProblemCreate)
	mux.HandleFunc("/problem/update/", handler.ProblemUpdate)
	mux.HandleFunc("/problem/delete/", handler.Problemdelete)

	// SolvedProblem handlers
	mux.HandleFunc("/solvedproblem/", handler.SolvedProblemGet)
	mux.HandleFunc("/solvedproblem/create", handler.SolvedProblemCreate)
	mux.HandleFunc("/solvedproblem/update/", handler.SolvedProblemUpdate)
	mux.HandleFunc("/solvedproblem/delete/", handler.SolvedProblemDelete)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}