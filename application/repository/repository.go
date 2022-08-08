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

func (d URLRepositoryDb) Get(alias string) (*[]domain.URL, error) {
	urlSql := "SELECT ALIAS, URL, RETRIEVAL_COUNT from URLS WHERE ALIAS = ?"
	result, err := newGetQuery(d, urlSql, alias)
	if err == nil {
		for i := 0; i < len(*result); i++ {
			(*result)[i].RetrievalCount = (*result)[i].RetrievalCount + 1
			updateSql := "UPDATE URLS SET `RETRIEVAL_COUNT` = ? WHERE ALIAS = ?"
			_, e := d.client.Exec(updateSql, (*result)[i].RetrievalCount, alias)
			if e != nil {
				fmt.Printf(e.Error())
			}
		}
	}

	return result, err
}

func (d URLRepositoryDb) GetMostUsed() (*[]domain.URL, error) {
	urlSql := "SELECT ALIAS, URL, RETRIEVAL_COUNT from URLS ORDER BY RETRIEVAL_COUNT DESC limit 10"
	return newGetQuery(d, urlSql, "")
}

func (d URLRepositoryDb) Save(url *domain.URL) error {
	urlSql := `INSERT INTO URLS(ALIAS, URL) VALUES(?, ?)`
	_, err := d.client.Exec(urlSql,
		url.Alias,
		url.URL,
	)

	return err
}

func newGetQuery(d URLRepositoryDb, urlSql, alias string) (*[]domain.URL, error) {
	var url domain.URL

	var res *sql.Rows
	var err error
	if alias != "" {
		res, err = d.client.Query(urlSql, alias)
	} else {
		res, err = d.client.Query(urlSql)
	}

	if err != nil {
		return nil, errors.New("fail to get URL")
	}

	var urls []domain.URL

	for res.Next() {

		err := res.Scan(
			&url.Alias,
			&url.URL,
			&url.RetrievalCount,
		)

		if err != nil {
			return nil, errors.New("fail to get URL")
		}
		urls = append(urls, url)
	}

	if len(urls) == 0 {
		return nil, errors.New("URL NOT FOUND")
	}

	return &urls, nil
}
