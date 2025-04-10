package main

import (
	"fmt"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/Turut4/GradeFlow/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateExamPayload struct {
	Title       string             `json:"title" validate:"max=100,min=3,required"`
	Subject     string             `json:"subject" validate:"max=100"`
	GradeLevel  string             `json:"grade_level" validate:"required,max=10"`
	Options     int                `json:"options"`
	Questions   int                `json:"questions"`
	AnswerSheet []store.AnswerItem `gorm:"json" json:"answer_sheet"`
}

func (api *application) createExamHandler(c *fiber.Ctx) error {
	var payload CreateExamPayload
	if err := c.BodyParser(&payload); err != nil {
		return api.badRequestResponse(c, err)
	}

	user := parseUserFromCtx(c)

	if err := Validate.Struct(&payload); err != nil {
		return api.badRequestResponse(c, err)
	}

	if err := validateAnswerSheet(&payload); err != nil {
		return api.badRequestResponse(c, err)
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
		return api.internalError(c, err)
	}

	exam.AnswerSheetPDF = pdfBytes

	if err := api.store.Exams.Create(c.UserContext(), exam); err != nil {
		return api.internalError(c, err)
	}

	return api.jsonResponse(c, fiber.StatusCreated, exam)
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
			return fmt.Errorf("opção inválida %s na questão %d", exam.AnswerSheet[i].A, i+1)
		}

		if exam.AnswerSheet[i].W <= 0 {
			return fmt.Errorf("peso inválido (<= 0) na questão %d", i+1)
		}

		weight += exam.AnswerSheet[i].W
	}

	if weight > 10 {
		return fmt.Errorf("os pesos ultrapassam o limite 10 (atual: %.2f)", weight)
	}

	return nil
}

func (api *application) GetExamPDFHandler(c *fiber.Ctx) error {
	examID, err := c.ParamsInt("examID")
	if err != nil {
		return api.badRequestResponse(c, err)
	}

	pdfBytes, err := api.store.Exams.GetExamPDF(c.Context(), uint(examID))
	if err != nil {
		switch err {
		case store.ErrNotFound:
			return api.notFoundResponse(c, err)
		default:
			return api.internalError(c, err)
		}
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", fmt.Sprintf("inline; filename=\"prova_%d.pdf\"", examID))

	return c.Send(pdfBytes)
}
