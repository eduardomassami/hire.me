package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/eduardomassami/hime.me/domain"
	_ "github.com/go-sql-driver/mysql"
)

type URLRepositoryDb struct {
	client *sql.DB
}

func newSqlClient(user, password, address, port, name string, timeout int) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, address, port, name)
	client, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	client.SetMaxIdleConns(10)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(time.Duration(timeout) * time.Second)
	client.SetConnMaxLifetime(time.Duration(timeout) * time.Second)
	return client, nil
}

func NewURLRepositoryDb(user, password, address, port, name string, timeout int) (domain.Repository, error) {
	sqlClient, err := newSqlClient(user, password, address, port, name, timeout)
	if err != nil {
		return nil, err
	}
	return &URLRepositoryDb{sqlClient}, nil
}

func (d URLRepositoryDb) Get(alias string) ([]*domain.URL, error) {
	urlSql := "SELECT Alias, URL from URLS WHERE ALIAS = ?"
	return newGetQuery(d, urlSql, alias)
}

func (d URLRepositoryDb) GetMostUsed() ([]*domain.URL, error) {
	urlSql := "SELECT TOP (10) Alias, URL from URLS ORDER BY ACCESS"
	return newGetQuery(d, urlSql, "")
}

func (d URLRepositoryDb) Save(url *domain.URL) error {
	urlSql := `INSERT INTO URLS(Alias, URL) VALUES(?, ?)`
	_, err := d.client.Exec(urlSql,
		url.Alias,
		url.URL,
	)

	return err
}

func newGetQuery(d URLRepositoryDb, urlSql, alias string) ([]*domain.URL, error) {
	var url domain.URL

	res, err := d.client.Query(urlSql, alias)

	if err != nil {
		return nil, errors.New("fail to get URL")
	}

	var urls []*domain.URL

	for res.Next() {

		err := res.Scan(
			&url.Alias,
			&url.URL,
		)

		if err != nil {
			return nil, errors.New("fail to get URL")
		}
		urls = append(urls, &url)
	}

	if len(urls) == 0 {
		return nil, errors.New("URL NOT FOUND")
	}

	return urls, nil
}
