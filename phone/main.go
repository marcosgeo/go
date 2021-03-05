package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
	phoneDB "github.com/marcosgeo/go/phone/db"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "zapgis"
	password = ".senha."
	dbname   = "gophercises"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(phoneDB.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phoneDB.Migrate("postgres", psqlInfo))

	db, err := phoneDB.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)
	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	row := db.QueryRow("SELECT value FROM phone_numbers WHERE id=$1", id)
	err := row.Scan(&number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

func normalizere(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}
