package service

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/Thomas3246/EquipAccounting/internal/domain"
	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database"
)

type EquipmentService struct {
	repo database.EquipmentRepo
}

func NewEquipmentService(repo database.EquipmentRepo) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (s *EquipmentService) GetAvailableEquipment(cookieValue string) ([]domain.EquipmentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	parts := strings.Split(cookieValue, "|")
	if len(parts) != 2 {
		return nil, ErrInvalidCookieParameter
	}

	isAdmin, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	if isAdmin == 1 {
		equipment, err := s.repo.GetActiveEquipment(ctx)
		if err != nil {
			return nil, err
		}
		return equipment, nil
	}

	if isAdmin == 0 {
		equipment, err := s.repo.GetActiveEquipmentForUserLogin(ctx, parts[0])
		if err != nil {
			return nil, err
		}
		return equipment, nil
	}

	return nil, ErrInvalidIsAdminValue
}

func (s *EquipmentService) GetEquipmentStates() ([]domain.EquipmentState, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	states, err := s.repo.GetEquipmentStates(ctx)
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (s *EquipmentService) GetEquipmentViewByFilter(department, state int) ([]domain.EquipmentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	equipment, err := s.repo.GetEquipmentViewByFilter(ctx, department, state)
	if err != nil {
		return nil, err
	}

	return equipment, nil
}

func (s *EquipmentService) GetEquipmentById(id int) (domain.Equipment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	equipment, err := s.repo.GetEquipmentById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Equipment{}, ErrNotFound
		}
		return domain.Equipment{}, err
	}

	return equipment, nil
}

func (s *EquipmentService) CheckInvNumForFreeToChange(id int, newNum string) (isFree bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	busyEquipment, err := s.repo.GetEquipmentByInvNum(ctx, newNum)
	if err != nil {
		if err == sql.ErrNoRows {
			// запись с таким инв.номером не найдена --> инв. номер свободен
			return true, nil
		}
		return false, err
	}

	// если найденая запись и есть редактируемая запись
	if id == busyEquipment.Id {
		return true, nil
	}

	return false, nil
}

func (s *EquipmentService) UpdateEquipment(equipment domain.Equipment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.UpdateEquipment(ctx, equipment)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentService) UpdatePC(equipment domain.Equipment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.UpdatePC(ctx, equipment)
	return err
}

func (s *EquipmentService) DeleteEquipment(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.DeleteEquipment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentService) CheckInvNumForFree(invNum string) (isFree bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = s.repo.GetEquipmentByInvNum(ctx, invNum)
	if err != nil {
		if err == sql.ErrNoRows {
			// запись с таким инв.номером не найдена --> инв. номер свободен
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (s *EquipmentService) NewEquipment(equipment domain.Equipment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	equipment.RegDate = time.Now().Format("2006.01.02")
	equipment.StatusId = 1

	err := s.repo.AddEquipment(ctx, equipment)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentService) ChangeEquipStatusByResult(equipId int, resultId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var statusToSet int
	if resultId == 1 {
		statusToSet = 1
	}

	err := s.repo.ChangeEquipStatus(ctx, equipId, statusToSet)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentService) ChangeEquipmentStatus(equipmentId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.repo.ChangeEquipStatus(ctx, equipmentId, 3)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentService) DecomEquipment(equipId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	decomDate := time.Now().Format("2006.01.02")
	err := s.repo.DecomEquipment(ctx, equipId, decomDate)
	if err != nil {
		return err
	}
	return nil
}

func (s *EquipmentService) GetAllEquipment() ([]domain.EquipmentView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	equipment, err := s.repo.GetAllEquipment(ctx)
	if err != nil {
		return nil, err
	}
	return equipment, nil
}
