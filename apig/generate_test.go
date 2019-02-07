package apig

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/vharish836/apig/msg"
)

var userModel = &Model{
	Name: "User",
	Fields: []*Field{
		&Field{
			Name:        "ID",
			JSONName:    "id",
			Type:        "uint",
			Tag:         "",
			Association: nil,
		},
		&Field{
			Name:        "Name",
			JSONName:    "name",
			Type:        "string",
			Tag:         "",
			Association: nil,
		},
		&Field{
			Name:        "CreatedAt",
			JSONName:    "created_at",
			Type:        "*time.Time",
			Tag:         "",
			Association: nil,
		},
		&Field{
			Name:        "UpdatedAt",
			JSONName:    "updated_at",
			Type:        "*time.Time",
			Tag:         "",
			Association: nil,
		},
	},
}

var detail = &Detail{
	Module:    "testmodule",
	Project:   "api-server",
	Model:     userModel,
	Models:    []*Model{userModel},
	ImportDir: "github.com/vharish836/api-server",
	Namespace: "",
}

func compareFiles(f1, f2 string) bool {
	c1, _ := ioutil.ReadFile(f1)
	c2, _ := ioutil.ReadFile(f2)

	return bytes.Compare(c1, c2) == 0
}

func setup() {
	msg.Mute = true
}

func teardown() {
	msg.Mute = false
}

// TestMain ...
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// TestGenerateApibIndex ...
func TestGenerateApibIndex(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateApibIndex")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateApibIndex(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "docs", "index.apib")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("API Blueprint index is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "docs", "index.apib")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate API Blueprint index correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateApibModel ...
func TestGenerateApibModel(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateApibModel")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateApibModel(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "docs", "user.apib")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("API Blueprint model is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "docs", "user.apib")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate API Blueprint model correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateController ...
func TestGenerateController(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateController")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateController(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "controllers", "user.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Controller file is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "controllers", "user.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate controller correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateRootController ...
func TestGenerateRootController(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateRootController")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateRootController(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "controllers", "root.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Controller file is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "controllers", "root.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate controller correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateREADME ...
func TestGenerateREADME(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateREADME")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateREADME(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "README.md")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("README is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "README.md")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate README correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateRouter ...
func TestGenerateRouter(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateRouter")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateRouter(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "router", "router.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Router file is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "router", "router.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate router correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateDBSQLite ...
func TestGenerateDBSQLite(t *testing.T) {
	detail.Database = "sqlite"

	outDir, err := ioutil.TempDir("", "generateDBSQLite")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateDB(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "db", "db.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("db.go is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "db", "db_sqlite.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate db.go correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateDBPostgres ...
func TestGenerateDBPostgres(t *testing.T) {
	detail.Database = "postgres"

	outDir, err := ioutil.TempDir("", "generateDBPostgres")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateDB(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "db", "db.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("db.go is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "db", "db_postgres.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate db.go correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

// TestGenerateDBMysql ...
func TestGenerateDBMysql(t *testing.T) {
	detail.Database = "mysql"

	outDir, err := ioutil.TempDir("", "generateDBMysql")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateDB(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "db", "db.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("db.go is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "db", "db_mysql.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate db.go correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}
