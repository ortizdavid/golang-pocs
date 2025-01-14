package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/serialization"
	"github.com/ortizdavid/go-rest-concepts/entities"
	"github.com/ortizdavid/go-rest-concepts/models"
	"gorm.io/gorm"
)


type TaskHandler struct {
	taskModel models.TaskModel
}


func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		taskModel: *models.NewTaskModel(db),
	}
}


func (h TaskHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /api/tasks", h.getAllTasks)
	router.HandleFunc("GET /api/tasks/{id}", h.getTask)
	router.HandleFunc("POST /api/tasks", h.createTask)
	router.HandleFunc("PUT /api/tasks/{id}", h.updateTask)
	router.HandleFunc("DELETE /api/tasks/{id}", h.deleteTask)
	router.HandleFunc("POST /api/tasks/import-csv", h.importTasksCSV)
	router.HandleFunc("GET /api/tasks-xml", h.getAllTasksXml)
	router.HandleFunc("GET /api/tasks-xml/{id}", h.getTaskXml)
}


func (h TaskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	currentPage, limit := GetCurrentPageAndLimit(r)

	tasks, err := h.taskModel.FindAllDataLimit(currentPage, limit)
	count := h.taskModel.Count()
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		httputils.WriteJsonError(w, "Tasks not found", http.StatusNotFound)
		return
	}
	httputils.WriteJsonPaginated(w, r, tasks, count, currentPage, limit)
}


func (h TaskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id := conversion.StringToInt(rawId)
	task, err := h.taskModel.FindById(id)

	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	httputils.WriteJson(w, http.StatusOK, task)
}


func (h TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task entities.Task

	if err := serialization.DecodeJson(r.Body, &task); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Uuid, created, updated
	task.UniqueId = encryption.GenerateUUID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err := h.taskModel.Create(task)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJson(w, http.StatusCreated, task)
}


func (h TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask entities.Task
	id := r.PathValue("id")
	taskId := conversion.StringToInt(id)
	
	task, err := h.taskModel.FindById(taskId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := serialization.DecodeJson(r.Body, &updatedTask); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return 
	}

	task.TaskName = updatedTask.TaskName
	task.UserId = updatedTask.UserId
	task.StartDate = task.EndDate
	task.EndDate = updatedTask.EndDate
	task.Description = updatedTask.Description
	task.UpdatedAt = time.Now()
	
	_, err = h.taskModel.Update(task)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJson(w, http.StatusOK, task)
}


func (h TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	taskId := conversion.StringToInt(id)
	task, err := h.taskModel.FindById(taskId)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = h.taskModel.Delete(task)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJson(w, http.StatusNoContent, nil)
	fmt.Fprintf(w, "Delete a task")
}


func (h TaskHandler) getAllTasksXml(w http.ResponseWriter, r *http.Request) {
	currentPage, limit := GetCurrentPageAndLimit(r)
	tasks, err := h.taskModel.FindAllDataLimit(currentPage, limit)
	count := h.taskModel.Count()
	if err != nil {
		httputils.WriteXmlError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		httputils.WriteXmlError(w, "Tasks not found", http.StatusNotFound)
		return
	}
	httputils.WriteXmlPaginated(w, r, tasks, count, currentPage, limit)
}


func (h TaskHandler) getTaskXml(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id := conversion.StringToInt(rawId)

	task, err := h.taskModel.FindById(id)
	if err != nil {
		httputils.WriteXmlError(w, err.Error(), http.StatusNotFound)
		return
	}
	
	exists, _ := h.taskModel.ExistsRecord("task_id", id)
	if !exists {
		httputils.WriteXmlError(w, "Task wih id: "+rawId+" not exists", http.StatusNotFound)
		return
	}
	httputils.WriteXml(w, http.StatusOK, task)
}



func (h TaskHandler) importTasksCSV(w http.ResponseWriter, r *http.Request) {
	csvFile, _, err := r.FormFile("csv_file")
	if err != nil {
		httputils.WriteJsonError(w, "could not upload csv", http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	if err := models.SkipCsvHeader(reader); err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasksCsv, err := models.ParseTaskFromCSV(reader)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = h.taskModel.CreateBatch(tasksCsv)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	httputils.WriteJson(w, http.StatusCreated, nil)
}

