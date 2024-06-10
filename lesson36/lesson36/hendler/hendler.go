package handler

import (
	"dodi/postgres"
)

type Handler struct {
	User          *postgres.UserRepo
	Problem       *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}
