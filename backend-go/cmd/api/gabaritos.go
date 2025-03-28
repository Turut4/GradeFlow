package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

func (api *application) GenerateGabarito(c *fiber.Ctx) error {
	var payload GabaritoPayload

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	resultChan := make(chan []byte, 1)
	errorChan := make(chan error, 1)

	if err := c.BodyParser(&payload); err != nil {
		return api.badResquestResponse(c, err)
	}

	go func() {
		pdfBytes, err := GeneratePDF(payload)
		if err != nil {
			errorChan <- err
			return
		}

		resultChan <- pdfBytes
	}()

	select {
	case pdfBytes := <-resultChan:
		c.Set("Content-Type", "application/pdf")
		return c.Send(pdfBytes)
	case err := <-errorChan:
		return api.internalError(c, err)
	case <-ctx.Done():
		return api.timeoutResponse(c)
	}
}

type GabaritoPayload struct {
	Test         string `json:"test"`
	Grade        string `json:"grade"`
	Alternatives int    `json:"alternatives"`
	Questions    int    `json:"questions"`
	TeacherName  string `json:"teacher_name"`
}

func GeneratePDF(payload GabaritoPayload) ([]byte, error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetMargins(5, 5, 5)

	pdf.AddUTF8Font("DejaVu", "", "DejaVuSans.ttf")

	if payload.Alternatives < 4 || payload.Alternatives > 5 {
		payload.Alternatives = 4
	}

	totalQuestions := payload.Questions
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
		pdf.Cell(0, 5, payload.Test)

		pdf.SetFont("DejaVu", "", 10)
		pdf.SetXY(gabaritoX+5, headerY+10)
		pdf.Cell(0, 3, fmt.Sprintf("Série: %s   Turma: ___   Data: ___/___/___   Prof: %s", payload.Grade, payload.TeacherName))

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
			if payload.Alternatives == 5 {
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
	return buf.Bytes(), nil
}
