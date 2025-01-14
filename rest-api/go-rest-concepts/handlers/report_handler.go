package handlers

import (
	"fmt"
	"net/http"
	"github.com/ortizdavid/go-nopain/docgen"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-rest-concepts/models"
	"gorm.io/gorm"
)


type ReportHandler struct {
	tableReport models.TableReport
}


func NewReportHandler(db *gorm.DB) *ReportHandler {
	return &ReportHandler{
		tableReport: models.NewTableReport(db),
	}
}


func (h ReportHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /api/reports", h.reports)
}


func (h ReportHandler) reports(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	pdfGen :=  docgen.NewHtmlPdfGenerator()
	var report models.TableReport

	switch param {
	case "users":
		report = h.tableReport.GetAllUsers()
	case "tasks":
		report = h.tableReport.GetAllTasks()
	default:
		httputils.WriteJsonError(w, "Invalid report type", http.StatusBadRequest)
		return
	}
	//-----------------------
	templateFile :=  param +".html"
	title := "Report: " +report.Title
	fileName := title +".pdf"
	data := map[string]any{
		"Title": title,
		"AppName": "Task Management App",
		"Rows": report.Rows,
		"Count": report.Count,
	}
	//----------- Render PDF
	pdfBytes, err := pdfGen.GeneratePDF(fmt.Sprintf("./templates/reports/%s", templateFile), data)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pdfGen.SetOutput(w, pdfBytes, fileName)
	if err != nil {
		httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
}