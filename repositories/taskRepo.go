package repositories

import (
	"database/sql"
	"errors"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"log"
)

type TaskRepoDb struct {
	DbClient *sql.DB
}

func NewTaskRepoDb(dbClient *sql.DB) *TaskRepoDb {
	return &TaskRepoDb{DbClient: dbClient}
}

func (t TaskRepoDb) FindAllTaskByUserId(userIdFromToken string) ([]domains.Task, error) {
	query := `SELECT * FROM tasks WHERE user_id = $1`
	rows, err := t.DbClient.Query(query, userIdFromToken)
	if err != nil {
		log.Println("error while getting rows", err)

		return nil, err
	}
	defer rows.Close()

	var tasks []domains.Task
	var task domains.Task
	for rows.Next() {
		if err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.Priority,
			&task.Category,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			log.Println("error while scanning task data", err)

			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}

func (t TaskRepoDb) CreateNewTask(task *domains.Task) (string, error) {
	query := `INSERT INTO tasks VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	res, err := t.DbClient.Exec(query,
		task.Id,
		task.UserId,
		task.Title,
		task.Description,
		task.DueDate,
		task.Priority,
		task.Category,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	)
	if err != nil {
		return "", err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Println("failed to create a task, err: ", err)
		return "", err
	}

	log.Println(affectedRows)

	return task.Id, nil
}

func (t TaskRepoDb) UpdateTask(task *domains.Task) error {
	updateQuery := `UPDATE tasks SET 
                 	title = $1,
                 	description = $2,
                 	due_date = $3,
                 	priority = $4,
                 	category = $5,
                 	status = $6,
                 	updated_at = $7
                 	WHERE id = $8`

	res, err := t.DbClient.Exec(updateQuery,
		task.Title,
		task.Description,
		task.DueDate,
		task.Priority,
		task.Category,
		task.Status,
		task.UpdatedAt,
		task.Id,
	)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Println("failed to create a task, err: ", err)
		return err
	}

	log.Println(affectedRows)

	return nil
}

func (t TaskRepoDb) DeleteTask(id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	res, err := t.DbClient.Exec(query, id)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Println("failed to create a task, err: ", err)
		return err
	}

	log.Println(affectedRows)

	return nil
}

func (t TaskRepoDb) FindTaskById(taskId string) (*domains.Task, error) {
	var task domains.Task
	query := `SELECT * FROM tasks WHERE id = $1`
	err := t.DbClient.QueryRow(query, taskId).Scan(
		&task.Id,
		&task.UserId,
		&task.Title,
		&task.Description,
		&task.DueDate,
		&task.Priority,
		&task.Category,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("task does not exists")
	}
	if err != nil {
		log.Println("failed to scan task")
		return nil, err
	}

	return &task, nil
}

func (t TaskRepoDb) FindTaskByFilter(query string, values ...any) ([]domains.Task, error) {
	rows, err := t.DbClient.Query(query, values...)
	if err != nil {
		log.Println("rows err- ", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []domains.Task
	var task domains.Task
	for rows.Next() {
		if err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.Priority,
			&task.Category,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			log.Println("error while scanning task data", err)

			return nil, err
		}
		tasks = append(tasks, task)
	}
	log.Println("total tasks found: ", len(tasks))
	return tasks, err

}
