package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

type conn struct {
	session *mgo.Session
}

func NewConnection() Connection {
	var c conn
	var err error
	url := getURL()
	c.session, err = mgo.Dial(url)
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) DB() *mgo.Database {
	return c.session.DB(os.Getenv("DATABASE_NAME"))
}

func getURL() string {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Println("error on load db port from env:", err.Error())
		port = 27017
	}
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		port,
		os.Getenv("DATABASE_NAME"))
}
