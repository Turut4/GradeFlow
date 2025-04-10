package store

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"

	"gorm.io/gorm"
)

type AnswerItem struct {
	S string  `json:"s"`
	A string  `json:"a"`
	W float32 `json:"w"`
}

type Exam struct {
	gorm.Model
	Title          string       `gorm:"size:48;not null"  json:"title"`
	UserID         uint         `json:"user_id"`
	User           User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user"`
	GradeLevel     string       `gorm:"size:10"            json:"grade_level"`
	Options        int          `json:"options"`
	AnswerSheetPDF []byte       `gorm:"bytea" json:"-"`
	AnswerSheet    []AnswerItem `gorm:"json" json:"answer_sheet"`
}

type ExamStore struct {
	db *gorm.DB
}

func (s *ExamStore) Create(ctx context.Context, exam *Exam) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	compressedPDF, err := compressPDF(exam.AnswerSheetPDF)
	if err != nil {
		return err
	}

	exam.AnswerSheetPDF = compressedPDF

	return s.db.WithContext(ctx).Create(exam).Error
}

func (s *ExamStore) GetByID(ctx context.Context, examID uint) (*Exam, error) {
	exam := new(Exam)
	if err := s.db.WithContext(ctx).Omit("AnswerSheetPDF").First(exam, examID).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return exam, nil
}

func (s *ExamStore) GetExamPDF(ctx context.Context, examID uint) ([]byte, error) {
	var exam Exam
	if err := s.db.WithContext(ctx).
		Select("answer_sheet_pdf").
		Where("id = ?", examID).
		First(&exam).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("exam not found: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get exam PDF: %w", err)
	}

	if len(exam.AnswerSheetPDF) == 0 {
		return nil, errors.New("exam PDF is empty")
	}

	pdfBytes, err := decompressPDF(exam.AnswerSheetPDF)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress PDF: %w", err)
	}

	return pdfBytes, nil
}

func compressPDF(pdfBytes []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(pdfBytes)
	if err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decompressPDF(compressedPDF []byte) ([]byte, error) {
	buf := bytes.NewReader(compressedPDF)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	return io.ReadAll(gz)
}
