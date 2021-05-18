package dbstore

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

/*
Converts sqlite schema to postgresql compatible sql commands.

Is not guaranteed to handle all edge cases.
*/

/*
CreatePGMigrations takes a path to a migrations folder. Outputs a
path to a temporary directory containing the same migrations but
with the postgresql converstion applied.
*/
func CreatePGMigrations(inputDir string) (string, error) {
	outputDir := os.TempDir()

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		inputFile := path.Join(inputDir, entry.Name())
		sqliteData, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return "", err
		}

		outputFile := path.Join(outputDir, entry.Name())
		postgresData := []byte(ConvertSQLite(string(sqliteData)))
		err = ioutil.WriteFile(outputFile, postgresData, fs.ModePerm) 
		if err != nil {
			return "", err
		}
	}

	return outputDir, nil
}

// ConvertSQLite takes a sqlite sql string and returns a postgresql-compatible version
func ConvertSQLite(sqlite string) string {
	sqlite = strings.ToLower(sqlite)
	postgresql := strings.ReplaceAll(sqlite, "integer primary key", "SERIAL PRIMARY KEY")
	postgresql = strings.ReplaceAll(postgresql, "blob", "BYTEA")
	return postgresql
}
