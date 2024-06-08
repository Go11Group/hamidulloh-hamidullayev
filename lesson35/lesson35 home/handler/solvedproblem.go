package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

func (h *Handler) SolvedProblemGet(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/solvedproblem/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid solved problem ID: %s", err.Error()), http.StatusBadRequest)
		return
	}

	solvedProblem, err := h.SolvedProblem.SolvedProblem_Get(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while getting solved problem: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(solvedProblem); err != nil {
		http.Error(w, fmt.Sprintf("Error while encoding response: %s", err.Error()), http.StatusInternalServerError)
	}
}

func (h *Handler) SolvedProblemDelete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/solvedproblemdelete/")

	err := h.SolvedProblem.SolvedProblem_Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while deleting solved problem, err: %s", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Solved problem deleted successfully"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while encoding response, err: %s", err.Error())))
	}
}

func (h *Handler) SolvedProblemCreate(w http.ResponseWriter, r *http.Request) {
	var solvedProblem model.SolvedProblem
	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding, err: %s", err.Error())))
		return
	}

	err = h.SolvedProblem.SolvedProblem_Create(solvedProblem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while creating solved problem, err: %s", err.Error())))
		return
	}

	w.Write([]byte("Solved problem created successfully"))
}

func (h *Handler) SolvedProblemUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/solvedproblemupdate/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid solved problem ID, err: %s", err.Error())))
		return
	}

	var updatedSolvedProblem model.SolvedProblem
	err = json.NewDecoder(r.Body).Decode(&updatedSolvedProblem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding, err: %s", err.Error())))
		return
	}

	err = h.SolvedProblem.SolvedProblem_Update(id, updatedSolvedProblem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while updating solved problem, err: %s", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Solved problem updated successfully"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while encoding response, err: %s", err.Error())))
	}
}
