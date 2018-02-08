package users

import (
	"big_projects/pagination"
	"database/sql"
	"log"

	"github.com/tokopedia/sqlt"
)

// User model
type User struct {
	UserID     int            `db:"user_id"`
	FullName   string         `db:"full_name"`
	MSISDN     string         `db:"msisdn"`
	UserEmail  string         `db:"user_email"`
	BirthDate  sql.NullString `db:"birth_date"`
	CreateTime string         `db:"create_time"`
	UpdateTime sql.NullString `db:"update_time"`
}

// GetUser get users
func GetUser(db *sqlt.DB, limit int, lastID int) ([]User, pagination.SimplePagination) {
	var users []User
	var total int
	query := `
		SELECT count(user_id) from ws_user
	`
	if err := db.QueryRow(query).Scan(&total); err != nil {
		log.Println(err)
	}

	query = `
		SELECT 
			user_id, full_name, msisdn, user_email, birth_date, create_time, update_time 
		FROM ws_user
		WHERE user_id > $1
		ORDER BY user_id 
		LIMIT $2
	`
	err := db.Select(&users, query, lastID, limit)
	if err != nil {
		log.Println(err)
	}
	userTot := len(users)
	if userTot == 0 {
		return users, pagination.SimplePagination{}
	}

	paginate := pagination.SimplePagination{
		Prev:     lastID,
		Next:     users[limit-1].UserID,
		First:    0,
		Last:     getLastUserID(db),
		Total:    userTot,
		TotalAll: total,
	}

	return users, paginate
}

func Total(db *sqlt.DB) int {
	var total int
	if err := db.QueryRow("Select count(user_id) FROM ws_user").Scan(&total); err != nil {
		log.Println(err)
	}

	return total
}

//GetUserByName get user filter by full_name
func GetUserByName(db *sqlt.DB, name string, limit int, lastID int) ([]User, pagination.SimplePagination) {
	var users []User
	var total int
	query := `
		SELECT 
			count(user_id)
		FROM ws_user 
		WHERE 
			full_name ILIKE $1
		AND  
			user_id > $2
		ORDER BY user_id
	`
	if err := db.QueryRow(query).Scan(&total); err != nil {
		log.Println(err)
	}

	if err := db.QueryRow("Select count(user_id) FROM ws_user").Scan(&total); err != nil {
		log.Println(err)
	}
	searchName := "%" + name + "%"

	query = `
		SELECT 
			user_id, full_name, msisdn, user_email, birth_date, create_time, update_time 
		FROM ws_user 
		WHERE 
			full_name ILIKE $1
		AND  
			user_id > $2
		ORDER BY user_id
		LIMIT $3
	`
	err := db.Select(&users, query, searchName, lastID, limit)
	if err != nil {
		log.Println(err)
	}

	userTot := len(users)
	if userTot == 0 {
		return users, pagination.SimplePagination{}
	}

	paginate := pagination.SimplePagination{
		Prev:     lastID,
		Next:     users[userTot-1].UserID,
		Last:     getLastUserID(db),
		Total:    userTot,
		TotalAll: total,
		Limit:    limit,
		Params:   "&name=" + name,
	}

	return users, paginate
}

func getLastUserID(db *sqlt.DB) int {
	var id int

	query := `
		SELECT 
			user_id
		FROM ws_user
		ORDER BY user_id DESC
		LIMIT 1
	`
	err := db.QueryRow(query).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id
}
