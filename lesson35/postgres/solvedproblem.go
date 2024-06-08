package postgres

import (
    "database/sql"
    "github.com/Go11Group/at_lesson/lesson34/model"
)
type SolvedProblemRepo struct {
    db *sql.DB
}

func SolvedProblem_NewBookRepo(db *sql.DB) *SolvedProblemRepo {
    return &SolvedProblemRepo{db: db}
}

func (repo *SolvedProblemRepo) SolvedProblem_Get(id int) (model.SolvedProblem, error) {
    var solvedProblem model.SolvedProblem
    err := repo.db.QueryRow("SELECT id, user_id, problem_id FROM solved_problems WHERE id = $1", id).Scan(
        &solvedProblem.ID, &solvedProblem.UserID, &solvedProblem.ProblemID)
    if err != nil {
        return solvedProblem, err
    }
    return solvedProblem, nil
}

func (repo *SolvedProblemRepo) SolvedProblem_Delete(id string) error {
    _, err := repo.db.Exec("DELETE FROM solved_problems WHERE id = $1", id)
    if err != nil {
        return err
    }
    return nil
}

func (repo *SolvedProblemRepo) SolvedProblem_Create(solvedProblem model.SolvedProblem) error {
    _, err := repo.db.Exec("INSERT INTO solved_problems (user_id, problem_id) VALUES ($1, $2)",
        solvedProblem.UserID, solvedProblem.ProblemID)
    if err != nil {
        return err
    }
    return nil
}

func (repo *SolvedProblemRepo) SolvedProblem_Update(id int, solvedProblem model.SolvedProblem) error {
    query := "UPDATE solved_problems SET user_id=$1, problem_id=$2 WHERE id=$3"
    _, err := repo.db.Exec(query, solvedProblem.UserID, solvedProblem.ProblemID, id)
    if err != nil {
        return err
    }
    return nil
}