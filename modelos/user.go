package modelos

import "time"

type User struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	CreatedAt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, email string, creatat time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.Email = email
	usuario.CreatedAt = creatat
	usuario.Status = status
}
