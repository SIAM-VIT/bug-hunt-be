package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/siam-vit/bughunt-be/internal/utils"
)

type Dbinstance struct {
	Db *sqlx.DB
}

var DB Dbinstance

func Connect() {
	p := utils.Config("DB_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Shanghai", utils.Config("DB_HOST"), utils.Config("DB_USER"), utils.Config("DB_PASSWORD"), utils.Config("DB_NAME"), port)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		log.Fatal("Failed to ping the database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")

	runMigrations(db)

	DB = Dbinstance{
		Db: db,
	}
}

func runMigrations(db *sqlx.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS questions (
		id SERIAL PRIMARY KEY,  
		question TEXT NOT NULL,
		language VARCHAR(255) NOT NULL,
		test_case_id INTEGER[],  
		set INT NOT NULL,
		difficulty VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS test_cases (
		id SERIAL PRIMARY KEY,  
		input TEXT NOT NULL,
		output TEXT NOT NULL,
		question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE  
	);

	CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE,
	score INT NOT NULL,
	time_remaining INTEGER DEFAULT 0
);


	
	`)

	if err != nil {
		log.Fatal("Failed to run migrations. \n", err)
		os.Exit(2)
	}

	log.Println("Migrations completed")
}
