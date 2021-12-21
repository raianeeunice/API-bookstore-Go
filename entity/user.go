package entity

// Users representa a tabela usuarios no banco de dados
type Users struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-; not null" json:"-"` // permitir ler e criar
	Token    string  `gorm:"-" json:"token,omitempty"` // ignorar este campo quando escrever e ler com struct
	Books    *[]Book `json:"books,omitempty"`
}
