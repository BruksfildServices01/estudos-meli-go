package service

import (
	"errors"
	"sync"

	"api-campeonato/model"
)

type TorneioTimeService struct {
	mu       sync.Mutex
	relacoes []model.TorneioTime
}

func NewTorneioTimeService() *TorneioTimeService {
	return &TorneioTimeService{
		relacoes: []model.TorneioTime{},
	}
}

func (s *TorneioTimeService) AddTimeToTorneio(torneioID, timeID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, rel := range s.relacoes {
		if rel.TorneioID == torneioID && rel.TimeID == timeID {
			return errors.New("time já está nesse torneio")
		}
	}

	s.relacoes = append(s.relacoes, model.TorneioTime{
		TorneioID: torneioID,
		TimeID:    timeID,
	})
	return nil
}

func (s *TorneioTimeService) RemoveTimeFromTorneio(torneioID, timeID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := -1
	for i, rel := range s.relacoes {
		if rel.TorneioID == torneioID && rel.TimeID == timeID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("relação torneio/time não encontrada")
	}

	s.relacoes = append(s.relacoes[:idx], s.relacoes[idx+1:]...)
	return nil
}

func (s *TorneioTimeService) ListTimesByTorneio(torneioID int) []int {
	s.mu.Lock()
	defer s.mu.Unlock()

	var ids []int
	for _, rel := range s.relacoes {
		if rel.TorneioID == torneioID {
			ids = append(ids, rel.TimeID)
		}
	}
	return ids
}
