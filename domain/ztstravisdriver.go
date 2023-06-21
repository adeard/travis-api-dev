package domain

type ZtsTravisDriverRequest struct {
	VendorID      string `json:"vendor_id" validate:"required"`
	DriverID      string `json:"driver_id" validate:"required"`
	Driver_Name   string `json:"driver_name" validate:"required"`
	KTP           string `json:"ktp" validate:"required"`
	No_HP         string `json:"no_hp" validate:"required"`
	VehicleID     string `json:"vehicle_id" validate:"required"`
	Driver_Status string `json:"driver_status" validate:"required"`
	UserID        string `json:"user_id" validate:"required"`
	Udate         string `json:"udate" validate:"required"`
}

type ZtsTravisDriver struct {
	VendorID      string `json:"vendor_id" gorm:"column:VendorID"`
	DriverID      string `json:"driver_id" gorm:"column:DriverID"`
	Driver_Name   string `json:"driver_name" gorm:"column:Driver_Name"`
	KTP           string `json:"ktp" gorm:"column:KTP"`
	No_HP         string `json:"no_hp" gorm:"column:No_HP"`
	VehicleID     string `json:"vehicle_id" gorm:"column:VehicleID"`
	Driver_Status string `json:"driver_status" gorm:"column:Driver_Status"`
	UserID        string `json:"user_id" gorm:"column:UserID"`
	Udate         string `json:"udate" gorm:"column:Udate"`
}

func (ZtsTravisDriver) TableName() string {
	return "ZTS_Travis_Driver"
}
