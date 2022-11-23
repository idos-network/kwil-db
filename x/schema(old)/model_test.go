package schema

import (
	"fmt"
	"os"
	"testing"
)

func Test_ReadYaml(t *testing.T) {
	// read in test.yaml

	bts, err := os.ReadFile("test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	db, err := readYaml(bts)
	if err != nil {
		t.Fatal(err)
	}

	if db.Owner != "kwil" {
		t.Fatal("owner should be kwil")
	}

	if db.Name != "mydb" {
		t.Fatal("name should be mydb")
	}

	fmt.Println(db.Tables)

	pgDB, err := convertToPG(db)
	if err != nil {
		t.Fatal(err)
	}

	ddlStr := BuildDDL(pgDB)
	fmt.Println(ddlStr)

	err = Verify(db)
	if err != nil {
		t.Fatal(err)
	}

}
