package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/Turut4/GradeFlow/internal/utils"
	"github.com/go-chi/chi/v5"
)

type CreateExamPayload struct {
	Title       string             `json:"title"        validate:"max=100,min=3,required"`
	Subject     string             `json:"subject"      validate:"max=100"`
	GradeLevel  string             `json:"grade_level"  validate:"required,max=10"`
	Options     int                `json:"options"`
	Questions   int                `json:"questions"`
	AnswerSheet []store.AnswerItem `json:"answer_sheet"                                   gorm:"json"`
}

func (app *application) createExamHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var payload CreateExamPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := parseUserFromCtx(r)

	if err := Validate.Struct(&payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := validateAnswerSheet(&payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	exam := &store.Exam{
		Title:       payload.Title,
		GradeLevel:  payload.GradeLevel,
		Options:     payload.Options,
		AnswerSheet: payload.AnswerSheet,
		UserID:      user.ID,
	}

	totalQuestions := len(exam.AnswerSheet)
	pdfBytes, err := utils.GeneratePDF(
		exam.Title,
		exam.GradeLevel,
		user.Username,
		exam.Options,
		totalQuestions,
	)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	exam.AnswerSheetPDF = pdfBytes

	if err := app.store.Exams.Create(r.Context(), exam); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, exam); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) GetExamHandler(w http.ResponseWriter, r *http.Request) {
	examID, err := strconv.ParseInt(chi.URLParam(r, "examID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	exam, err := app.store.Exams.GetByID(r.Context(), uint(examID))
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, exam); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) GetAnswerSheetHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	examID, err := strconv.ParseInt(chi.URLParam(r, "examID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pdfBytes, err := app.store.Exams.GetAnswerSheet(r.Context(), uint(examID))
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().
		Set("Content-Disposition", fmt.Sprintf("inline; filename=\"prova_%d.pdf\"", examID))

	w.Write(pdfBytes)
}

func validateAnswerSheet(exam *CreateExamPayload) error {
	validAnswers := map[string]bool{
		"A": true,
		"B": true,
		"C": true,
		"D": true,
		"E": true,
	}
	var weight float32

	for i := range exam.Questions {
		if !validAnswers[exam.AnswerSheet[i].A] {
			return fmt.Errorf(
				"opção inválida %s na questão %d",
				exam.AnswerSheet[i].A,
				i+1,
			)
		}

		if exam.AnswerSheet[i].W <= 0 {
			return fmt.Errorf("peso inválido (<= 0) na questão %d", i+1)
		}

		weight += exam.AnswerSheet[i].W
	}

	if weight > 10 {
		return fmt.Errorf(
			"os pesos ultrapassam o limite 10 (atual: %.2f)",
			weight,
		)
	}

	return nil
}
