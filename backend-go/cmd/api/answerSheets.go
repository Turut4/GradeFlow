package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type TestResultPayload struct {
	Answers []string `json:"answers"`
	Result  []bool   `json:"results"`
}

func (app *application) processAnswersSheet(
	w http.ResponseWriter,
	r *http.Request,
) {
	body := &bytes.Buffer{}
	writer, err := parseFile(r, body)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	resp, err := http.Post(
		app.config.ocr.addr,
		writer.FormDataContentType(),
		body,
	)
	if err != nil {
		app.internalServerError(
			w,
			r,
			fmt.Errorf("erro ao enviar a imagem para o microserviço: %v", err),
		)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		app.internalServerError(
			w,
			r,
			fmt.Errorf("erro ao processar a imagem: status %d, resposta: %s",
				resp.StatusCode, string(body)),
		)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		app.internalServerError(
			w,
			r,
			fmt.Errorf("erro ao ler a resposta: %v", err),
		)
		return
	}

	var result TestResultPayload
	if err := json.Unmarshal(respBody, &result); err != nil {
		app.internalServerError(
			w,
			r,
			fmt.Errorf("erro ao decodificar a resposta: %v", err),
		)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, result); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func parseFile(r *http.Request, body *bytes.Buffer) (*multipart.Writer, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do formulário: %v", err)
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("arquivo não enviado")
	}

	defer file.Close()

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", handler.Filename)
	if err != nil {
		return nil, fmt.Errorf("não foi possível criar o formulário: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("não foi possível copiar o arquivo: %v", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("erro ao fechar o writer: %v", err)
	}

	return writer, err
}
