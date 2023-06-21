package domain

type VehicleTypeRequest struct {
	VehicleType          string `json:"vehicle_type" validate:"required"`
	VehicleType_Name     string `json:"vehicle_type_name" validate:"required"`
	VehicleType_Category string `json:"vehicle_type_category" validate:"required"`
}

type VehicleType struct {
	VehicleType          string `json:"vehicle_type" gorm:"column:VehicleType"`
	VehicleType_Name     string `json:"vehicle_type_name" gorm:"column:VehicleType_Name"`
	VehicleType_Category string `json:"vehicle_type_category" gorm:"column:VehicleType_Category"`
}

func (VehicleType) TableName() string {
	return "VehicleType"
}
