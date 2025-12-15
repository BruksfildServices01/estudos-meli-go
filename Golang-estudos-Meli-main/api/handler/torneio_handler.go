package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"api-campeonato/service"
)

type TorneioHandler struct {
	Service         *service.TorneioService
	TorneioTimeServ *service.TorneioTimeService
}

func NewTorneioHandler(s *service.TorneioService, tt *service.TorneioTimeService) *TorneioHandler {
	return &TorneioHandler{
		Service:         s,
		TorneioTimeServ: tt,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func parseIDFromPath(path string) (int, error) {

	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, strconv.ErrSyntax
	}
	return strconv.Atoi(parts[1])
}

func (h *TorneioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/torneios" {
		switch r.Method {
		case http.MethodPost:
			h.create(w, r)
		case http.MethodGet:
			h.list(w, r)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/torneios/") && strings.Contains(r.URL.Path, "/times") {
		h.handleTorneioTimes(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/torneios/") {
		id, err := parseIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "id inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.getByID(w, r, id)
		case http.MethodPut:
			h.update(w, r, id)
		case http.MethodDelete:
			h.delete(w, r, id)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

func (h *TorneioHandler) handleTorneioTimes(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) < 3 || parts[0] != "torneios" || parts[2] != "times" {
		http.NotFound(w, r)
		return
	}

	torneioID, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "torneio_id inválido", http.StatusBadRequest)
		return
	}

	if len(parts) == 3 {
		switch r.Method {
		case http.MethodPost:
			h.addTime(w, r, torneioID)
		case http.MethodGet:
			h.listTimes(w, r, torneioID)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	if len(parts) == 4 {
		timeID, err := strconv.Atoi(parts[3])
		if err != nil {
			http.Error(w, "time_id inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodDelete:
			h.removeTime(w, r, torneioID, timeID)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

func (h *TorneioHandler) addTime(w http.ResponseWriter, r *http.Request, torneioID int) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		TimeID int `json:"time_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	if err := h.TorneioTimeServ.AddTimeToTorneio(torneioID, input.TimeID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TorneioHandler) listTimes(w http.ResponseWriter, r *http.Request, torneioID int) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	ids := h.TorneioTimeServ.ListTimesByTorneio(torneioID)
	writeJSON(w, http.StatusOK, ids)
}

func (h *TorneioHandler) removeTime(w http.ResponseWriter, r *http.Request, torneioID, timeID int) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := h.TorneioTimeServ.RemoveTimeFromTorneio(torneioID, timeID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ===== Métodos de CRUD de torneio =====

// @Summary      Cria um novo torneio
// @Description  Cria um torneio com nome e ano
// @Tags         torneios
// @Accept       json
// @Produce      json
// @Param        body  body      model.Torneio  true  "Dados do torneio"
// @Success      201   {object}  model.Torneio
// @Failure      400   {string}  string  "json inválido"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /torneios [post]
func (h *TorneioHandler) create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Ler JSON
	var input struct {
		Nome string `json:"nome"`
		Ano  int    `json:"ano"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	// Chamar o service
	t := h.Service.Create(input.Nome, input.Ano)

	// Responder
	writeJSON(w, http.StatusCreated, t)
}

// @Summary      Lista todos os torneios
// @Description  Retorna todos os torneios em memória
// @Tags         torneios
// @Produce      json
// @Success      200  {array}   model.Torneio
// @Failure      405  {string}  string  "método não permitido"
// @Router       /torneios [get]
func (h *TorneioHandler) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	lista := h.Service.List()
	writeJSON(w, http.StatusOK, lista)
}

// @Summary      Busca torneio por ID
// @Description  Retorna um torneio específico
// @Tags         torneios
// @Produce      json
// @Param        id   path      int  true  "ID do torneio"
// @Success      200  {object}  model.Torneio
// @Failure      404  {string}  string  "torneio não encontrado"
// @Router       /torneios/{id} [get]
func (h *TorneioHandler) getByID(w http.ResponseWriter, r *http.Request, id int) {
	t, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "torneio não encontrado", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// @Summary      Atualiza um torneio
// @Description  Atualiza nome e ano de um torneio existente
// @Tags         torneios
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "ID do torneio"
// @Param        body  body      model.Torneio true  "Dados do torneio"
// @Success      200   {object}  model.Torneio
// @Failure      400   {string}  string  "json inválido"
// @Failure      404   {string}  string  "torneio não encontrado"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /torneios/{id} [put]
func (h *TorneioHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodPut {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nome string `json:"nome"`
		Ano  int    `json:"ano"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	t, err := h.Service.Update(id, input.Nome, input.Ano)
	if err != nil {
		http.Error(w, "torneio não encontrado", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// @Summary      Deleta um torneio
// @Description  Remove um torneio pelo ID
// @Tags         torneios
// @Param        id  path  int  true  "ID do torneio"
// @Success      204  {string}  string  "No Content"
// @Failure      404  {string}  string  "torneio não encontrado"
// @Failure      405  {string}  string  "método não permitido"
// @Router       /torneios/{id} [delete]
func (h *TorneioHandler) delete(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "torneio não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
