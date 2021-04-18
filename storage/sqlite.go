package storage

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	// path is the absolute path tp the password
	// database.
	path = "/var/lib/vault/vault.db"
)

type Store struct {
	db *gorm.DB
}

var Storage = &Store{}

type Password struct {
	gorm.Model
	Name string
	Body string
	Meta []Meta
}

type Meta struct {
	gorm.Model
	PasswordID uint
	Name       string
	Value      string
}

func init() {
	var migrate bool
	if _, err := os.Stat(path); os.IsNotExist(err) {
		migrate = true
	}

	var err error
	Storage.db, err = gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	if migrate {
		Storage.db.AutoMigrate(&Password{}, &Meta{})
	}
}

func (s *Store) Put(name string, password string, meta map[string]string) {
	p := Password{Name: name}
	Storage.db.Where("name = ?", name).First(&p)
	p.Body = password
	Storage.db.Save(&p)

	for name, value := range meta {
		model := Meta{
			PasswordID: p.ID,
			Name:       name,
			Value:      value,
		}
		Storage.db.Create(&model)
	}
}

func (s *Store) Show(name string) (*Password, []*Meta) {
	p := Password{}
	Storage.db.Where("name = ?", name).Find(&p)
	var m []*Meta
	Storage.db.Model(&p).Related(&m)
	return &p, m
}

func (s *Store) Drop(name string) {
	p := Password{Name: name}
	Storage.db.Where("name = ?", name).Find(&p)
	Storage.db.Delete(&p)
}

func (s *Store) List() []Password {
	var passwords []Password
	Storage.db.Find(&passwords)
	return passwords
}
