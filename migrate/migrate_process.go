package migrate

import (
	_ "blog/initial"
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
func MigrateUp() {
	batch, m := GetLatestMigrationsFile("up")
	upsql, _, files := LoadMigrationsFile("up", m)

	o := orm.NewOrm()
	for _, k := range upsql {
		err := o.Begin()
		_, err = o.Raw(k).Exec()
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	if len(files) > 0 {
		for _, v := range files {
			o.Insert(&Migrate{Migration: v, Batch: batch + 1})
			fmt.Print("migrate " + v + " successfully\n")
		}
	} else {
		fmt.Print("no migrations\n")
	}
}

/* 回滚迁移文件 */
func MigrateDown() {
	batch, m := GetLatestMigrationsFile("down")
	_, downsql, files := LoadMigrationsFile("down", m)

	o := orm.NewOrm()
	for _, k := range downsql {
		err := o.Begin()
		_, err = o.Raw(k).Exec()
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	if len(files) > 0 {
		for _, v := range files {
			qs := o.QueryTable(new(Migrate))
			qs.Filter("batch", batch).Filter("migration", v).Delete()
			fmt.Print("rollback " + v + " successfully\n")
		}
	} else {
		fmt.Print("no rollback\n")
	}
}

func MigrateStatus() {
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
