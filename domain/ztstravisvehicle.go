package domain

type ZtsTravisVehicleRequest struct {
	VendorID    string `json:"vendor_id" validate:"required"`
	VehicleID   string `json:"vehicle_id" validate:"required"`
	VehicleNo   string `json:"vehicle_no" validate:"required"`
	VehicleType string `json:"vehicle_type" validate:"required"`
	UserID      string `json:"user_id" validate:"required"`
	Udate       string `json:"udate" validate:"required"`
}

type ZtsTravisVehicle struct {
	VendorID    string `json:"vendor_id"`
	VehicleID   string `json:"vehicle_id"`
	VehicleNo   string `json:"vehicle_no"`
	VehicleType string `json:"vehicle_type"`
	UserID      string `json:"user_id"`
	Udate       string `json:"udate"`
}

func (ZtsTravisVehicle) TableName() string {
	return "ZTS_Travis_Vehicle"
}
