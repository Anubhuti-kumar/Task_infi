package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"infilon_test/model"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func GetPersonInfo(c *gin.Context) {
	personID := c.Param("person_id")
	id, err := strconv.Atoi(personID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	var person model.Person
	err = db.QueryRow("SELECT p.id, p.name, p.age, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code "+
		"FROM person p "+
		"JOIN phone ph ON p.id = ph.person_id "+
		"JOIN address_join aj ON p.id = aj.person_id "+
		"JOIN address a ON aj.address_id = a.id "+
		"WHERE p.id = ?", id).Scan(&person.ID, &person.Name, &person.Age, &person.PhoneNumber, &person.City, &person.State, &person.Street1, &person.Street2, &person.ZipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person info"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func CreatePerson(c *gin.Context) {
	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := db.Exec("INSERT INTO person (name, age, phone_number, city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		person.Name, person.Age, person.PhoneNumber, person.City, person.State, person.Street1, person.Street2, person.ZipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get last inserted ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person created successfully", "id": id})
}
