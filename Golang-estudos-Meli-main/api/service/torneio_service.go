package service

import (
	"errors"
	"math/rand"
	"sync"

	"api-campeonato/model"
)

type TorneioService struct {
	mu       sync.Mutex
	torneios map[int]model.Torneio
}

func NewTorneioService() *TorneioService {
	return &TorneioService{
		torneios: make(map[int]model.Torneio),
	}
}

func (s *TorneioService) Create(nome string, ano int) model.Torneio {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := rand.Intn(1000000)

	t := model.Torneio{
		ID:   id,
		Nome: nome,
		Ano:  ano,
	}

	s.torneios[id] = t
	return t
}

func (s *TorneioService) List() []model.Torneio {
	s.mu.Lock()
	defer s.mu.Unlock()

	var lista []model.Torneio
	for _, t := range s.torneios {
		lista = append(lista, t)
	}
	return lista
}

func (s *TorneioService) GetByID(id int) (model.Torneio, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.torneios[id]
	if !ok {
		return model.Torneio{}, errors.New("torneio não encontrado")
	}
	return t, nil
}

func (s *TorneioService) Update(id int, nome string, ano int) (model.Torneio, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.torneios[id]
	if !ok {
		return model.Torneio{}, errors.New("torneio não encontrado")
	}

	t := model.Torneio{
		ID:   id,
		Nome: nome,
		Ano:  ano,
	}

	s.torneios[id] = t
	return t, nil
}

func (s *TorneioService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.torneios[id]; !ok {
		return errors.New("torneio não encontrado")
	}

	delete(s.torneios, id)
	return nil
}
