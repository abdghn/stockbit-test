/*
 * Created on 11/10/21 16.08
 *
 * Copyright (c) 2021 Abdul Ghani Abbasi
 */

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Persistent interface {
}

type persistent struct {
	conn *sql.DB
}

func NewPersistent(dataSourceName string) (Persistent, error) {
	c, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[resource.db] open sql connection failed: %v", err)
	}

	err = c.Ping()
	if err != nil {
		return nil, fmt.Errorf("[resource.db] ping db failed: %v", err)
	}

	return &persistent{conn: c}, nil
}

