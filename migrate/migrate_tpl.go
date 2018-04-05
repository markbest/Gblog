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

const (
	migrateTpl      = "./migrate/tpl/migrate.sql.tpl"
	migrateFilePath = "./migrate/migrations/"
)

/* 加载迁移模板的内容 */
func LoadMigrateTpl() ([]byte, error) {
	f, err := os.Open(migrateTpl)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

/* 创建迁移文件 */
func CreateMigration(name string) {
	prefix := time.Now().Format("2006_01_02_030405")
	tplContent, err := LoadMigrateTpl()
	if err != nil {
		panic(err)
	}

	//创建目录
	os.MkdirAll(migrateFilePath, 0777)

	//创建数据库迁移文件
	migrateFile := prefix + "_" + name + ".sql"
	f, err := os.Create(migrateFilePath + migrateFile)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	io.WriteString(f, string(tplContent))
	fmt.Println("create db " + migrateFile + " successfully")
}

/* 获取迁移文件的内容以及文件列表 */
func LoadMigrationsFile(action string, m []Migrate) (migrateUp []string, migrateDown []string, migrations []string) {
	PthSep := string(os.PathSeparator)
	files := make([]string, 0)
	switch action {
	case "up":
		files = GetMatchMigrations(m)
		break
	case "down":
		if len(m) > 0 {
			for _, v := range m {
				files = append(files, "migrate"+PthSep+"migrations"+PthSep+v.Migration)
			}
		}
		break
	default:
		break
	}

	var mUp, mDown string
	for _, v := range files {
		mUp = ""
		mDown = ""
		migrations = append(migrations, strings.Replace(v, "migrate"+PthSep+"migrations"+PthSep, "", -1))
		up, down := ParseMigrationsFile(v)
		for _, m := range up {
			mUp = mUp + m
		}
		for _, n := range down {
			mDown = mDown + n
		}
		migrateUp = append(migrateUp, mUp)
		migrateDown = append(migrateDown, mDown)
	}
	return migrateUp, migrateDown, migrations
}

/* 获取up和down匹配的迁移文件 */
func GetMatchMigrations(m []Migrate) []string {
	files := make([]string, 0)
	migrations := make([]string, 0)
	matchFiles := make([]string, 0)
	if _, err := os.Stat(migrateFilePath); os.IsNotExist(err) {
		return matchFiles
	}
	err := filepath.Walk(migrateFilePath, func(path string, f os.FileInfo, err error) error {
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
	matchFiles = difference(files, migrations)
	return matchFiles
}

/* 解析迁移文件的SQL语句 */
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

/* 获取所有迁移文件和已经执行过迁移的文件的差集 */
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