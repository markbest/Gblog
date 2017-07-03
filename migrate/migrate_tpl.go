package migrate

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const migrate_tpl = "./database/tpl/migrate.sql.tpl"
const migrate_file_path = "./database/migrations/"

func LoadMigrateTpl() ([]byte, error) {
	f, err := os.Open(migrate_tpl)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func CreateMigration(name string) {
	prefix := time.Now().Format("2006_01_02_030405")
	tpl_content, _ := LoadMigrateTpl()

	//创建目录
	os.MkdirAll(migrate_file_path, 0777)

	f, err := os.Create(migrate_file_path + prefix + "_" + name + ".sql")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	io.WriteString(f, string(tpl_content))
	fmt.Println("create migrations " + prefix + "_" + name + ".sql successfully")
}

func LoadMigrationsFile(action string, m []Migrate) (migrate_up []string, migrate_down []string, migrations []string) {
	PthSep := string(os.PathSeparator)
	files := make([]string, 0)
	switch action {
	case "up":
		files = GetMatchMigrations(m)
		break
	case "down":
		if len(m) > 0 {
			for _, v := range m {
				files = append(files, "database"+PthSep+"migrations"+PthSep+v.Migration)
			}
		}
		break
	default:
		break
	}

	var m_up, m_down string
	for _, v := range files {
		m_up = ""
		m_down = ""
		migrations = append(migrations, strings.Replace(v, "database"+PthSep+"migrations"+PthSep, "", -1))
		up, down := ParseMigrationsFile(v)
		for _, m := range up {
			m_up = m_up + m
		}
		for _, n := range down {
			m_down = m_down + n
		}
		migrate_up = append(migrate_up, m_up)
		migrate_down = append(migrate_down, m_down)
	}
	return migrate_up, migrate_down, migrations
}

func GetMatchMigrations(m []Migrate) []string {
	files := make([]string, 0)
	migrations := make([]string, 0)
	matchs := make([]string, 0)
	err := filepath.Walk(migrate_file_path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	if len(m) > 0 {
		for _, v := range m {
			migrations = append(migrations, v.Migration)
		}
	}
	matchs = difference(files, migrations)
	return matchs
}

func ParseMigrationsFile(name string) (up []string, down []string) {
	isUp := false
	isDown := false

	f, err := os.Open(name)
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			if strings.Contains(line, "-- Up") {
				isUp = true
				isDown = false

			}
			if strings.Contains(line, "-- Down") {
				isUp = false
				isDown = true
			}

			if isUp {
				if strings.Contains(line, "--") == false {
					up = append(up, line)
				}
			}
			if isDown {
				if strings.Contains(line, "--") == false {
					down = append(down, line)
				}
			}
		}
	}
	return up, down
}

//获取迁移文件和migrations的差集
func difference(slice1 []string, slice2 []string) []string {
	var diff []string
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if strings.Contains(s1, s2) {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, s1)
		}
	}
	return diff
}
