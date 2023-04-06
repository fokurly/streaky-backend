package postgre

import (
	"fmt"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/lib/pq"
)

func (d *Db) InsertNewTask(task models.Task) (*int64, error) {
	const (
		insertNewUser = `INSERT INTO task_info(id, userid, firstobserver, secondobserver, name, description, punish, frequencyperiod, status, startdate, enddate) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	)

	//var tmpFreq pq.StringArray
	//tmpFreq = task.Frequency
	id, err := d.CountNumberOfColumns()
	if err != nil {
		return nil, fmt.Errorf("[InsertNewTask] - could not get columns in table. error: %v", err)
	}

	task.ID = *id + 1
	_, err = d.db.Exec(insertNewUser, task.ID, task.UserID, task.FirstObserver,
		task.SecondObserver, task.Name, task.Description,
		task.Punish, task.FrequenctPQ, task.State,
		task.StartDate, task.EndDate)

	if err != nil {
		return nil, fmt.Errorf("[InsertNewTask] - could not exec query. error: %v", err)
	}

	{
		err := d.InsertTaskToCreator(task.UserID, task.ID)
		if err != nil {
			return nil, fmt.Errorf("[InsertNewTask] - could not add task to creator. error: %v", err)
		}
		// TODO: Вынести?
		err = d.InsertTaskToObserver(task.FirstObserver, task.ID)
		if err != nil {
			return nil, fmt.Errorf("[InsertNewTask] - could not add observer. error: %v", err)
		}

		err = d.InsertTaskToObserver(task.SecondObserver, task.ID)
		if err != nil {
			return nil, fmt.Errorf("[InsertNewTask] - could not add observer. error: %v", err)
		}
	}
	return &task.ID, nil
}

func (d *Db) CountNumberOfColumns() (*int64, error) {
	const (
		countNumOfСolumns = `SELECT COUNT(*) FROM task_info`
	)

	rows, err := d.db.Query(countNumOfСolumns)
	if err != nil {
		return nil, fmt.Errorf("[CountNumberOfColumns] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var tmp *int64
	if rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("[CountNumberOfColumns] - could not scan rows. error: %v", err)
		}
	}

	return tmp, nil
}

func (d *Db) InsertTaskToCreator(userID, taskID int64) error {
	const (
		insertTaskToCreator = `UPDATE user_tasks SET task_ids=$1 WHERE user_id=$2`
	)

	tasks, err := d.GetUserTasks(userID)
	if err != nil {
		return fmt.Errorf("could not get observer tasks. error: %v", err)
	}

	tasks = append(tasks, taskID)
	var tmp pq.Int64Array
	tmp = tasks

	_, err = d.db.Exec(insertTaskToCreator, tmp, userID)
	if err != nil {
		return fmt.Errorf("could not exec query. error: %v", err)
	}

	return nil
}

func (d *Db) InsertTaskToObserver(userID, taskID int64) error {
	const (
		insertTaskToObserver = `UPDATE user_tasks SET observer_tasks_ids=$1 WHERE user_id=$2`
	)

	tasks, err := d.GetObservedTasks(userID)
	if err != nil {
		return fmt.Errorf("could not get observer tasks. error: %v", err)
	}

	for i := range tasks {
		if tasks[i] == taskID {
			// TODO: брать другого пользователя, если текущий занят
			return fmt.Errorf("user is already observing this task")
		}
	}

	tasks = append(tasks, taskID)
	var tmp pq.Int64Array
	tmp = tasks

	_, err = d.db.Exec(insertTaskToObserver, tmp, userID)
	if err != nil {
		return fmt.Errorf("could not exec query. error: %v", err)
	}

	return nil
}

func (d *Db) GetObservedTasks(userID int64) ([]int64, error) {
	const (
		selectObserverdTasks = `SELECT observer_tasks_ids FROM user_tasks WHERE user_id=$1`
	)

	rows, err := d.db.Query(selectObserverdTasks, userID)
	if err != nil {
		return nil, fmt.Errorf("[GetObservedTasks] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var tasks []int64
	// TODO: заменить tmp везде, где встречается
	var tmp pq.Int64Array
	if rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("[GetObservedTasks] - could not scan rows. error: %v", err)
		}
	}

	if tmp != nil {
		tasks = tmp
	}

	return tasks, nil
}

func (d *Db) GetUserTasks(userID int64) ([]int64, error) {
	const (
		selectTasks = `SELECT task_ids FROM user_tasks WHERE user_id=$1`
	)

	rows, err := d.db.Query(selectTasks, userID)
	if err != nil {
		return nil, fmt.Errorf("[GetUserTasks] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var tasks []int64
	// TODO: заменить tmp везде, где встречается
	var tmp pq.Int64Array
	if rows.Next() {
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("[GetUserTasks] - could not scan rows. error: %v", err)
		}
	}

	if tmp != nil {
		tasks = tmp
	}

	return tasks, nil
}

func (d *Db) UpdateTaskStatus(status string, ID int64) error {
	const (
		updateStatus = `UPDATE task_info SET status=$1 WHERE id=$2`
	)

	_, err := d.db.Exec(updateStatus, status, ID)
	if err != nil {
		return fmt.Errorf("could not exec query. error: %v", err)
	}

	return nil
}

// Метод для получения полной инфы по таскам
func (d *Db) GetTaskInfo(ID int64) (*models.Task, error) {
	const (
		selectTask = `SELECT id, userid, firstobserver, secondobserver, name, description, punish, frequencyperiod, status, startdate, enddate FROM task_info WHERE id=$1`
	)

	rows, err := d.db.Query(selectTask, ID)
	if err != nil {
		return nil, fmt.Errorf("could not get task from database. error: %v", err)
	}

	var info models.Task
	if rows.Next() {
		err := rows.Scan(&info.ID, &info.UserID, &info.FirstObserver, &info.SecondObserver, &info.Name, &info.Description, &info.Punish, &info.FrequenctPQ, &info.State, &info.StartDate, &info.EndDate)
		if err != nil {
			return nil, fmt.Errorf("[GetUserTasks] - could not scan rows. error: %v", err)
		}
	}

	return &info, err
}
