package mysql_test

import (
	"testing"

	"github.com/khaiql/dbcleaner"
	"github.com/khaiql/dbcleaner/helper/mysql"
)

func TestGetTableQuery(t *testing.T) {
	h := mysql.Helper{}
	query := h.GetTablesQuery()
	if query != "SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE()" {
		t.Error("Wrong query")
	}
}

func TestTruncateTableCommand(t *testing.T) {
	h := mysql.Helper{}
	cmd := h.TruncateTableCommand("users")
	expectedCmd := "TRUNCATE TABLE users;"

	if cmd != expectedCmd {
		t.Errorf("Expected %s. Got %s", expectedCmd, cmd)
	}
}

func TestInit(t *testing.T) {
	_, err := dbcleaner.FindHelper("mysql")
	if err != nil {
		t.Errorf("Shouldn't get error but got %s", err.Error())
	}
}
