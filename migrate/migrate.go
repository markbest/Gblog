package migrate

import (
	_ "blog/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Migrate struct {
	Id        int64  `orm:"auto"`
	Migration string `orm:"size(256)"`
	Batch     int64
}

func (m *Migrate) TableName() string {
	return "migrations"
}

func init() {
	orm.RegisterModel(new(Migrate))
	orm.RunSyncdb("default", false, true)
}

/* 获取所有已经执行的迁移文件 */
func GetAllMigrationsFile() (m []string) {
	var migrations []Migrate
	o := orm.NewOrm()
	qs := o.QueryTable(new(Migrate))
	qs.OrderBy("batch").All(&migrations)
	if len(migrations) > 0 {
		for _, v := range migrations {
			m = append(m, v.Migration)
		}
	}
	return m
}

/* 获取最后一批操作的migrations */
func GetLatestMigrationsFile(action string) (batch int64, m []Migrate) {
	var migrate []Migrate
	o := orm.NewOrm()
	qs := o.QueryTable(new(Migrate))
	qs.OrderBy("-batch").Limit(1).All(&migrate)

	if len(migrate) > 0 {
		for _, v := range migrate {
			batch = v.Batch
		}
	} else {
		batch = 0
	}

	if batch > 0 {
		nqs := o.QueryTable(new(Migrate))
		if action == "up" {
			nqs.Filter("batch__lte", batch).All(&m)
		} else if action == "down" {
			nqs.Filter("batch", batch).All(&m)
		}
	}
	return batch, m
}

/* 执行迁移文件 */
func HandleMigrateUp() {
	batch, m := GetLatestMigrationsFile("up")
	upSql, _, files := LoadMigrationsFile("up", m)

	o := orm.NewOrm()
	o.Begin()
	for k, v := range upSql {
		if _, err := o.Raw(v).Exec(); err != nil {
			panic(err)
			o.Rollback()
		} else {
			o.Insert(&Migrate{Migration: files[k], Batch: batch + 1})
			fmt.Println("migrate " + files[k] + " successfully")
		}
	}
	o.Commit()

	if len(files) == 0 {
		fmt.Println("no migration files")
	}
}

/* 回滚迁移文件 */
func HandleMigrateDown() {
	batch, m := GetLatestMigrationsFile("down")
	_, downSql, files := LoadMigrationsFile("down", m)

	o := orm.NewOrm()
	o.Begin()
	for k, v := range downSql {
		if _, err := o.Raw(v).Exec(); err != nil {
			o.Rollback()
			panic(err)
		} else {
			qs := o.QueryTable(new(Migrate))
			qs.Filter("batch", batch).Filter("migration", files[k]).Delete()
			fmt.Println("rollback " + files[k] + " successfully")
		}
	}
	o.Commit()

	if len(files) ==  0 {
		fmt.Println("no rollback files")
	}
}

func HandleMigrateStatus() {
	//已经执行的迁移文件
	files := GetAllMigrationsFile()

	//尚未执行的迁移文件
	_, m := GetLatestMigrationsFile("up")
	_, _, others := LoadMigrationsFile("up", m)

	fmt.Println("+-----+----------------------------------------------------------+")
	fmt.Println("| Ran | Migration                                                |")
	fmt.Println("+-----+----------------------------------------------------------+")

	for _, v := range files {
		fmt.Println("|  Y  | " + v)
	}

	for _, k := range others {
		fmt.Println("|  N  | " + k)
	}

	fmt.Println("+-----+----------------------------------------------------------+")
}
