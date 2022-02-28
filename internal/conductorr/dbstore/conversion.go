package dbstore

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/lsnow99/conductorr/internal/conductorr/integration"
)

/*
Converts conductorr-annotated SQL to different dbms-compatible sql commands.

See the README file for annotation syntax
*/

/*
CreateMigrationsFor takes a path to a migrations folder. Outputs a
path to a temporary directory containing the same migrations but
with the requested dbms's conversion applied
*/
func CreateMigrationsFor(dbmsKey, inputDir string) (string, error) {
	outputDir, err := integration.MkdirTemp("conductorr_migrations_converted")
	if err != nil {
		return "", err
	}

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
		sqlMap := ConvertSQL(string(sqliteData))
		sql, ok := sqlMap[dbmsKey]
		if !ok {
			return "", fmt.Errorf("conversion not available for requested dbms %s", dbmsKey)
		}
		outData := []byte(sql)
		err = ioutil.WriteFile(outputFile, outData, fs.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return outputDir, nil
}

const sqliteKey = "--sqlite--"
const pgKey = "--postgresql--"
const endKey = "--end--"

var dbmsKeys = []string{sqliteKey, pgKey}

type sqlMapBuilder struct {
	sqlMap map[string]string
	curStr string
	curKey string
}

func (smb *sqlMapBuilder) flush(key string) {
	if smb.curKey == "" {
		for _, k := range dbmsKeys {
			smb.sqlMap[k] += smb.curStr
		}
	} else {
		smb.sqlMap[smb.curKey] += smb.curStr
	}
	smb.curStr = ""
	smb.curKey = key
}

// ConvertSQL takes a commented SQL string and returns a map of DBMS systems to their respective sql
func ConvertSQL(sql string) map[string]string {
	builder := sqlMapBuilder{
		sqlMap: map[string]string{},
	}

	scanner := bufio.NewScanner(strings.NewReader(sql))
	for scanner.Scan() {
		line := scanner.Text()
		var found bool
		for _, dbmsKey := range dbmsKeys {
			if checkLineForKey(line, dbmsKey) {
				builder.flush(dbmsKey)
				found = true
				break
			}
		}
		if !found {
			if checkLineForKey(line, endKey) {
				builder.flush("")
			} else {
				builder.curStr += line + "\n"
			}
		}
	}

	builder.flush("")

	return builder.sqlMap
}

func checkLineForKey(line, key string) bool {
	return strings.HasPrefix(strings.TrimSpace(strings.ToLower(line)), key)
}
