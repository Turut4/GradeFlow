# 🌟 **GradeFlow**

## 🚀 **Automatização de Correção de Provas Escolares**

O **GradeFlow** é uma plataforma SaaS inovadora que automatiza a correção de provas escolares, economizando tempo e esforço dos professores. A plataforma integra-se diretamente ao sistema **SIAP** para agilizar o lançamento de notas, utilizando tecnologias modernas e robustas para garantir eficiência e precisão.

---

## 📑 Índice

- [Funcionalidades](#✨ funcionalidades)
- [Tecnologias Utilizadas](#🛠️ tecnologias-utilizadas)
- [Instalação](#📝 instalação)
- [Uso](#🧑‍💻 uso)
- [Contribuição](#💡 contribuição)
- [Licença](#📜 licença)
- [Contato](#🗂️ contato)

---

## ✨ **Funcionalidades**

- **Upload de Provas:**
  Professores podem enviar imagens ou PDFs das provas respondidas pelos alunos.

- **Processamento de Imagens:**
  Correção automática utilizando técnicas de pré-processamento (alinhamento, remoção de ruídos) com **OpenCV**.

- **Reconhecimento de Respostas (OCR):**
  Identificação das respostas através de **Google Cloud Vision API** ou **Tesseract OCR**.

- **Correção Automática:**
  Comparação com o gabarito e geração de notas automaticamente.

- **Integração com SIAP:**
  Lançamento automático de notas no sistema SIAP via **Chromedp** ou **API REST**.

- **Dashboard Intuitivo:**
  Visualização de resultados, estatísticas e correções manuais através de uma interface moderna em **React**.

- **Segurança e Conformidade:**
  Proteção de dados sensíveis em conformidade com a **LGPD**.

---

## 🛠️ **Tecnologias Utilizadas**

### Backend

- **Go (Golang)**: API rápida e escalável com **Gin/Fiber**.
- **Python (FastAPI + UV)**: Microserviço de OCR para reconhecimento de respostas.

### Frontend

- **React + TypeScript**: Interface moderna e responsiva.

### Banco de Dados

- **PostgreSQL**: Armazenamento de dados estruturados.
- **Redis**: Cache e filas para processamento assíncrono.

### Infraestrutura

- **Docker + Kubernetes**: Orquestração e escalabilidade.

---

## 📝 **Instalação**

### Pré-requisitos

- **Go 1.21+**
- **Python 3.13+**
- **Docker e Kubernetes**
- **Node.js 18+**
- **PostgreSQL e Redis**

### Clonando o Projeto

```bash
git clone https://github.com/seu-usuario/gradeflow.git
cd gradeflow
```

### Configuração do Backend

1. Crie o arquivo de variáveis de ambiente:

   ```bash
   cp .env.example .env
   ```

2. Atualize as variáveis no arquivo `.env` conforme necessário.

3. Rode o backend:

   ```bash
   make run-backend
   ```

### Configuração do Frontend

1. Entre no diretório do frontend:

   ```bash
   cd frontend
   ```

2. Instale as dependências:

   ```bash
   npm install
   ```

3. Rode o servidor de desenvolvimento:

   ```bash
   npm start
   ```

### Rodando com Docker

```bash
docker-compose up --build
```

---

## 🧑‍💻 **Uso**

1. Acesse o painel em:

   ```text
   http://localhost:3000
   ```

2. Faça login ou cadastre-se.
3. Realize o upload das provas na seção de correção.
4. Revise os resultados e envie as notas para o SIAP.

---

## 💡 **Contribuição**

Contribuições são bem-vindas! Siga os passos abaixo:

1. Faça um fork do projeto.
2. Crie uma nova branch:

   ```bash
   git checkout -b feature/nova-funcionalidade
   ```

3. Commit suas alterações:

   ```bash
   git commit -m "Adiciona nova funcionalidade"
   ```

4. Faça um push para a branch:

   ```bash
   git push origin feature/nova-funcionalidade
   ```

5. Abra um Pull Request.

---

## 📜 **Licença**

Este projeto está licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para obter mais informações.

---

## 🗂️ **Contato**

- **Site**: [gradeflow.app](https://gradeflow.app)
- **GitHub**: [Link do Repositório](https://github.com/seu-usuario/gradeflow)
- **Documentação**: [Link da Documentação](https://docs.gradeflow.app)
