package entity

// Book representa a tabela book no banco de dados
type Book struct {
	ID          uint64 `gorm:"primary_Key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        Users  `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
