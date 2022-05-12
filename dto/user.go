package dto

// User ...
type User struct {
	Name string `json:"name" example:"John"`
	MCI  int    `json:"mci" example:"12345678"`
}
