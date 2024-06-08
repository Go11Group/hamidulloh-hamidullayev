package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type ProblemRepo struct {
	db *sql.DB
}

func Problem_NewBookRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{db: db}
}

func (u *ProblemRepo) Problem_Get(id int) (model.Problem, error) {
	var problem model.Problem
	err := u.db.QueryRow("SELECT id, title, description, difficulty FROM problems WHERE id = $1", id).Scan(
		&problem.ID, &problem.Title, &problem.Description, &problem.Difficulty)
	if err != nil {
		return problem, err
	}
	return problem, nil
}

func (u *ProblemRepo) Problem_delete(id string) error {
	_, err := u.db.Exec(`delete from problems where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProblemRepo) Problem_Create(problem model.Problem) error {
	_, err := u.db.Exec(`INSERT INTO problems (title, description, difficulty) 
                       VALUES ($1, $2, $3)`,
		problem.Title, problem.Description, problem.Difficulty)
	if err != nil {
		return err
	}

	return nil
}
func (repo *ProblemRepo) Problem_Update(id int, problem model.Problem) error {
	query := `UPDATE problems SET title=$1, description=$2, difficulty=$3 WHERE id=$4`
	_, err := repo.db.Exec(query, problem.Title, problem.Description, problem.Difficulty, id)
	if err != nil {
		return err
	}
	return nil
}
