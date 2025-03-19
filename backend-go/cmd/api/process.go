package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TestResultPayload struct {
	Answers []string `json:"answers"`
	Result  []bool   `json:"results"`
}

func (api *application) processGabarito(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return api.badResquestResponse(c, fmt.Errorf("arquivo não enviado"))
	}

	src, err := file.Open()
	if err != nil {
		return api.internalError(c, fmt.Errorf("imagem não pode ser aberta"))
	}

	defer src.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return api.internalError(c, fmt.Errorf("não foi possível criar o formulário: %v", err))
	}
	_, err = io.Copy(part, src)
	if err != nil {
		return api.internalError(c, fmt.Errorf("não foi possível copiar o arquivo: %v", err))
	}

	if err := writer.Close(); err != nil {
		return api.internalError(c, fmt.Errorf("erro ao fechar o writer: %v", err))
	}

	resp, err := http.Post(api.cfg.ocr.addr, writer.FormDataContentType(), body)
	if err != nil {
		return api.internalError(c, fmt.Errorf("erro ao enviar a imagem para o microserviço: %v", err))
	}
	if resp.StatusCode != http.StatusOK {
		return api.processError(c, resp.StatusCode, fmt.Errorf("erro ao processar a imagem"))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.internalError(c, fmt.Errorf("erro ao ler a resposta: %v", err))
	}

	var result TestResultPayload
	if err := json.Unmarshal(respBody, &result); err != nil {
		return api.internalError(c, fmt.Errorf("erro ao decodificar a resposta: %v", err))
	}

	return api.jsonResponse(c, http.StatusCreated, result)
}
