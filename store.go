package main

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

// struct for thing store
type Store struct {
	db     *sql.DB
	dbmap  *gorp.DbMap
	sessMu sync.Mutex
}

func NewStore(dbName string) *Store {
	// create a new store
	store := Store{}
	if dbName == "" {
		dbName = "/tmp/test.db"
	}

	db, err := sql.Open("sqlite3", dbName)
	store.db = db
	if err != nil {
		fmt.Println(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	store.dbmap = dbmap

	// map the objects
	dbmap.AddTable(ThingRef{}).SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		fmt.Print(err)
	}
	return &store
}

// gorp object to store things
type ThingRef struct {
	Id   int64
	URL  string
	Data []byte
}

func (s Store) Close() {
	s.db.Close()
}

// access methods
func (s Store) Query(q string) error {
	rows, err := s.db.Query(q)
	if err != nil {
		return err
	}
	fmt.Println(rows)
	return nil
}

func (s Store) GetThings(page int) ([]ThingRef, int) {
	var thingList []ThingRef
	pages := 10
	_, err := s.dbmap.Select(&thingList, "select * from ThingRef")
	if err != nil {
		panic(err)
	}
	return thingList, pages
}
