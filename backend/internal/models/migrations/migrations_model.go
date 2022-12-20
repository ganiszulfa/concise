package migrations

type Migration struct {
	Key string `json:"key" gorm:"primaryKey"`
}
