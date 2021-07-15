package models

import (
	"testing"
)

func TestInsertClient(t *testing.T) {
	initDBForTest(t)
	defer CloseDB()

	t.Fatalf("insert failed %v", InitClientList())
}

func TestDeleteClientById(t *testing.T) {
	initDBForTest(t)
	defer CloseDB()

	err := DeleteNotaryById(1)
	if err != nil {
		t.Fatalf("delete failed %v", err)
	}
}
