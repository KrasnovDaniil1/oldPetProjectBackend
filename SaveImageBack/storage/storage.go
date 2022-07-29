package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"saveimage/variable"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

type TableCard struct {
	Id       string `json:"id"`
	Tags     string `json:"tags"`
	Filename string `json:"src"`
}

type TableTags struct {
	Id   string `json:"id"`
	Tags string `json:"tags"`
}

var STORAGE PostgresStorage

func OpenPostgres() PostgresStorage {
	db, err := sql.Open("postgres", variable.OPEN_POSTGRES)
	if err != nil {
		log.Fatal(errors.New("OpenPostgres - не открылась база данных"))
	}
	storage := PostgresStorage{db: db}
	storage.CreateTable()

	return storage
}

func (p PostgresStorage) CreateTable() {
	_, err := p.db.Exec(variable.CREATE_TABLE)
	if err != nil {
		log.Fatal(errors.New("CreateTable - не получилось создать базу данных"))
		return
	}
	fmt.Println("Таблица готова к работе")
}

func (p PostgresStorage) UpdateTable(tags []string, filename string) error {
	_, err := p.db.Exec(variable.UPDATE_TABLE, pq.Array(tags), filename)
	return err
}

func (p PostgresStorage) GetAllCard(limit, offset string) ([]TableCard, error) {
	var allTableCard []TableCard
	rows, err := p.db.Query(variable.GET_ALL_CARD, limit, offset)
	for rows.Next() {
		var p TableCard
		rows.Scan(&p.Id, &p.Tags, &p.Filename)
		p.Filename = "http://localhost:8080/images/" + p.Filename
		allTableCard = append(allTableCard, p)
	}
	return allTableCard, err
}

func (p PostgresStorage) GetSearchCard(tag, limit, offset string) ([]TableCard, error) {
	var allSearchCard []TableCard
	rows, err := p.db.Query(variable.GET_SEARCH_CARD, tag, limit, offset)
	for rows.Next() {
		var p TableCard
		rows.Scan(&p.Id, &p.Tags, &p.Filename)
		p.Filename = "http:/localhost/images/" + p.Filename
		allSearchCard = append(allSearchCard, p)
	}
	return allSearchCard, err
}

func (p PostgresStorage) DeleteCard(id string) error {
	_, err := p.db.Exec(variable.DELETE_CARD_BY_ID, id)
	return err
}

func (p PostgresStorage) GetAllTags() ([]TableTags, error) {
	var allTableTags []TableTags
	rows, err := p.db.Query(variable.GET_ALL_TAGS)
	for rows.Next() {
		var p TableTags
		rows.Scan(&p.Id, &p.Tags)
		allTableTags = append(allTableTags, p)
	}
	return allTableTags, err
}

func (p PostgresStorage) GetSearchTag(search string) (string, error) {
	rows, err := p.db.Query(variable.GET_SEARCH_TAG, search)
	if rows != nil && err == nil {
		return search, nil
	}
	return "null", err
}

func (p PostgresStorage) GetLastId() (string, error) {
	var id string

	row := p.db.QueryRow(variable.GET_LAST_ID)
	err := row.Scan(&id)
	if err != nil {
		return "0", nil
	}
	return id, err
}
