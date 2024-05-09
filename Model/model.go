package Model

import "database/sql"

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

func (p *Person) Create(db *sql.DB) error {
	result, err := db.Exec("INSERT INTO person (name, phone_number, city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.Name, p.PhoneNumber, p.City, p.State, p.Street1, p.Street2, p.ZipCode)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	p.ID = int(id) // Convert int64 to int
	return nil
}
