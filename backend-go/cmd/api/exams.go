package main

import (
	"bytes"
	"fmt"
	"log"
	"math"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
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
	pdfBytes, err := GeneratePDF(
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

func GeneratePDF(title, grade, username string, optionsNum, totalQuestions int) ([]byte, error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetMargins(5, 5, 5)

	pdf.AddUTF8Font("DejaVu", "", "DejaVuSans.ttf")

	if optionsNum < 4 || optionsNum > 5 {
		optionsNum = 4
	}

	if totalQuestions == 0 {
		return nil, fmt.Errorf("nenhuma questão fornecida")
	}
	if totalQuestions > 100 {
		totalQuestions = 100
	}

	pdf.AddPage()
	gabaritoWidth := 138.5
	pageBottom := 200.0
	headerY, headerHeight := 12.0, 23.0
	questionStartY := headerY + headerHeight + 10
	questionHeight := 5.5
	colMargin := 4.5

	availableHeight := pageBottom - questionStartY
	questionsPerColumn := int((availableHeight - 10) / questionHeight)
	log.Printf("question: %d", questionsPerColumn)

	requiredColumns := int(math.Ceil(float64(totalQuestions) / float64(questionsPerColumn)))

	colWidth := (gabaritoWidth - colMargin*float64(requiredColumns-1)) / float64(requiredColumns)

	for g := range 2 {
		gabaritoX := 15.0 + float64(g)*gabaritoWidth
		if g == 0 {
			gabaritoX = 5 + float64(g)*gabaritoWidth
		}

		// Cabeçalho do gabarito
		pdf.SetFillColor(230, 230, 230)
		pdf.Rect(gabaritoX, headerY, gabaritoWidth, headerHeight, "FD")

		pdf.SetFont("DejaVu", "", 10)
		pdf.SetXY(gabaritoX+5, headerY+2)
		pdf.Cell(0, 5, title)

		pdf.SetFont("DejaVu", "", 10)
		pdf.SetXY(gabaritoX+5, headerY+10)
		pdf.Cell(0, 3, fmt.Sprintf("Série: %s   Turma: ___   Data: ___/___/___   Prof: %s", grade, username))

		pdf.SetXY(gabaritoX+5, headerY+18)
		pdf.Cell(0, 3, "Aluno:_________________________________")

		// Imprime as questões distribuídas pelas colunas
		pdf.SetFont("DejaVu", "", 7)
		for i := range totalQuestions {
			col := i / questionsPerColumn
			row := i % questionsPerColumn
			currentX := gabaritoX + float64(col)*(colWidth+colMargin) - 2
			currentY := questionStartY + float64(row)*questionHeight

			// Imprime número da questão e alternativas
			pdf.SetXY(currentX+2, currentY)
			pdf.Cell(4, 3, fmt.Sprintf("%02d)", i+1))

			optionX := currentX + 8
			options := []string{"A", "B", "C", "D"}
			if optionsNum == 5 {
				options = append(options, "E")
			}
			for _, option := range options {
				pdf.Circle(optionX+1.5, currentY+1.5, 1.7, "D")
				pdf.SetXY(optionX, currentY+0.5)
				pdf.CellFormat(3, 2, option, "", 0, "C", false, 0, "")
				optionX += 5
			}
		}
	}

	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(10+gabaritoWidth, headerY, 10+gabaritoWidth, pageBottom)

	// Marcadores de alinhamento
	pdf.Rect(5, 5, 5, 5, "D")
	pdf.Rect(287, 5, 5, 5, "D")
	pdf.Rect(138.5, 5, 5, 5, "D")
	pdf.Rect(153.5, 5, 5, 5, "D")
	pdf.Rect(138.5, 200, 5, 5, "D")
	pdf.Rect(153.5, 200, 5, 5, "D")
	pdf.Rect(5, 200, 5, 5, "D")
	pdf.Rect(287, 200, 5, 5, "D")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("falha ao gerar PDF: %v", err)
	}

	//test output
	// err = os.WriteFile("saida.pdf", buf.Bytes(), 0644)
	// if err != nil {
	// 	return nil, fmt.Errorf("falha ao salvar PDF no disco: %v", err)
	// }

	return buf.Bytes(), nil
}
