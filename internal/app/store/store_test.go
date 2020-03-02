package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func  TestMain( m *testing.M)  {
	databaseURL= os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user='anko' password='Mandarin_137054t' dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}