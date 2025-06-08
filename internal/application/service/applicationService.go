package service

import (
	"database/sql"

	"github.com/Thomas3246/EquipAccounting/internal/infrastructure/database/sqlite"
)

type ApplicationService struct {
	UserService               *UserService
	RequestService            *RequestService
	EquipmentService          *EquipmentService
	DepartmentService         *DepartmentService
	EquipmentDirectoryService *EquipmentDirectoryService
	DocumentService           *DocumentService
}

func NewAppService(db *sql.DB) *ApplicationService {

	userRepo := sqlite.NewUserRepo(db)
	userService := NewUserService(userRepo)

	requestRepo := sqlite.NewRequestRepo(db)
	requestService := NewRequestService(requestRepo)

	equipmentRepo := sqlite.NewEquipmentRepo(db)
	equipmentService := NewEquipmentService(equipmentRepo)

	departmentRepo := sqlite.NewDepartmentRepo(db)
	departmentService := NewDepartmentService(departmentRepo)

	equipmentDirectoryRepo := sqlite.NewEquipmentDirectoryRepo(db)
	equipmentDirectoryService := NewEquipmentDirectoryService(equipmentDirectoryRepo)

	documentRepo := sqlite.NewDocumentRepo(db)
	documentService := NewDocumentService(documentRepo)

	return &ApplicationService{
		UserService:               userService,
		RequestService:            requestService,
		EquipmentService:          equipmentService,
		DepartmentService:         departmentService,
		EquipmentDirectoryService: equipmentDirectoryService,
		DocumentService:           documentService,
	}
}
