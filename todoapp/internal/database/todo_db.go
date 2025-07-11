package storage

import (
	"todoapp/internal/models"

	"github.com/sirupsen/logrus"
)

func GetAll() ([]models.Todo, error) {
	rows, err :=  DB.Query("SELECT id, title, done FROM tasks")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next(){
		var t models.Todo

		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil{
			return nil,  err
		}
		todos = append(todos, t)
	}

	logrus.Info("Выведен список задач")
	return todos, nil
}

func InsertTasks(tasks models.Todo) (error) {
	stmt, err := DB.Prepare("INSERT INTO tasks (title, done) VALUES (?,?)")
	if err != nil{
		return err
	}

	_, err = stmt.Exec(tasks.Title, tasks.Done)
	if err != nil{
		return err
	}
	logrus.Info("Добавлена задача")

	return nil
}

func DeleteTaskFromDB(task_id int) (error) {
	stmt, err := DB.Prepare("DELETE FROM tasks WHERE id = ?")

	if err != nil{
		return err
	}

	if _, err := stmt.Exec(task_id); err != nil{
		return err
	}
	logrus.Info("Произошло удаление задачи")

	return nil
}

func UpdateDone(task_id int) (error) {

	rows, err := DB.Query("SELECT done FROM tasks WHERE id = ?", task_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var task models.Todo
	for rows.Next(){

		if err := rows.Scan(&task.Done); err != nil{
			return err
		}
	}

	if task.Done{
		stmt, err := DB.Prepare("UPDATE tasks SET done = 0 WHERE id = ?")
		if err != nil{
			return err
		}

		if _, err := stmt.Exec(task_id); err != nil{
			return err
		}
		logrus.Info("Изменен статус задачи")
	}else{
		stmt, err := DB.Prepare("UPDATE tasks SET done = 1 WHERE id = ?")
		if err != nil{
			return err
		}

		if _, err := stmt.Exec(task_id); err != nil{
			return err
		}
		logrus.Info("Изменен статус задачи")
	}

	return nil
}