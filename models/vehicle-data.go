package models

// VehicleAccidentStatus xxx
type VehicleAccidentStatus uint

const (
	// YesAccident xxx
	YesAccident VehicleAccidentStatus = 0
	// NoAccident xxx
	NoAccident VehicleAccidentStatus = 1
)

// VehicleData xxx
type VehicleData struct {
	Name               string                `json:"name"`
	Make               string                `json:"make"`
	Model              string                `json:"model"`
	SubModel           string                `json:"submodel"`
	Trim               string                `json:"trim"`
	Price              uint                  `json:"price"`
	YearlyDepreciation uint                  `json:"yearlyDepreciation"`
	KMDepreciation     uint                  `json:"kmDepreciation"`
	KM                 uint                  `json:"km"`
	Year               uint                  `json:"year"`
	Category           uint                  `json:"category"`
	AccidentStatus     VehicleAccidentStatus `json:"accidentStatus"`
}
