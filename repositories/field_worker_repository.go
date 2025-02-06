package repositories

import (
	"AgringBackend/config"
)

type VirtualFieldWorker struct {
	EmployeeName    string
	EmployeeSurname string
	FieldID         int
	WorkTypeName    string
}

func GetVirtualFieldWorkers() ([]VirtualFieldWorker, error) {
	var virtualWorkers []VirtualFieldWorker
	query := `
        SELECT u.name AS employee_name, u.surname AS employee_surname, fw.field_id, wt.name AS work_type_name
        FROM field_workers fw
        JOIN users u ON fw.employee_id = u.id
        JOIN work_types wt ON fw.work_type = wt.id
    `
	result := config.DB.Raw(query).Scan(&virtualWorkers)
	return virtualWorkers, result.Error
}
