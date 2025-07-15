package storage

import (
	"todoapp/internal/models"

	"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
)

func GetAll() ([]models.Todo, error){
	var tasks []models.Todo

	res := DB.Find(&tasks)
	
	if res.Error != nil{
		return nil, res.Error
	}

	logrus.Info("Выведен список задач")
	return tasks, nil
}

func InsertTasks(tasks models.Todo) error {
	res := DB.Create(&models.Todo{Title: tasks.Title, Done: tasks.Done})
	if res.Error != nil{
		return res.Error
	}

	logrus.Info("Добавлена задача")
	return nil
}

func DeleteTaskFromDB(task_id int) error {
	res := DB.Delete(&models.Todo{}, task_id)

	if res.Error != nil{
		return res.Error
	}
	logrus.Infof("Удалена задача %d", task_id)

	return nil
}


func UpdateDone(task_id int) (error) {
	var task models.Todo
	resFind := DB.Find(&task, task_id)

	if resFind.Error != nil{
		return resFind.Error
	}

	if task.Done{
		task.Done = false

		resSave := DB.Save(&task)

		if resSave.Error != nil{
			return resSave.Error
		}
		
		logrus.Info("Изменен статус задачи")
	}else{
		task.Done = true

		resSave := DB.Save(&task)

		if resSave.Error != nil{
			return resSave.Error
		}
		
		logrus.Info("Изменен статус задачи")
	}

	return nil
}
	