package stogo

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

type line struct {
	Name string
	Type string
	Tag  string
}

func GenerateStruct(ssql, dataSourceName string) {
	ssql = strings.TrimSpace(ssql)
	dataSourceName = strings.TrimSpace(dataSourceName)
	if ssql == "" {
		fmt.Printf("%s\n", "error sql 语句 为空")
		return
	}
	if dataSourceName == "" {
		fmt.Printf("%s\n", "error 链接数据库的url 为空")
		return
	}
	//拿到名字和变量
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	rows, err := conn.QueryContext(context.Background(), ssql)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//获取 column type
	cts, err := rows.ColumnTypes()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	ls := make([]*line, 0)
	//储存变量
	for _, ct := range cts {
		l := new(line)
		//下划线
		name := ct.Name()
		//tag
		l.Tag = fmt.Sprintf("`db:\"%s\"`", name)
		//大驼峰
		l.Name = underLineToUC(name)
		//类型
		l.Type = getDataType(ct.DatabaseTypeName())
		ls = append(ls, l)
	}
	//1.拿到最长的大驼峰
	var maxLen int
	for _, l := range ls {
		if len(l.Name) > maxLen {
			maxLen = len(l.Name)
		}
	}
	maxLen++
	//将所有的大驼峰加上空格
	for _, l := range ls {
		s := new(bytes.Buffer)
		for i := 0; i < maxLen-len(l.Name); i++ {
			s.WriteString(" ")
		}
		l.Name = l.Name + s.String()
	}
	//2.最长的类型
	maxLen = 0
	for _, l := range ls {
		if len(l.Type) > maxLen {
			maxLen = len(l.Type)
		}
	}
	maxLen++
	//将所有的大驼峰加上空格
	for _, l := range ls {
		s := new(bytes.Buffer)
		for i := 0; i < maxLen-len(l.Type); i++ {
			s.WriteString(" ")
		}
		l.Type = l.Type + s.String()
	}
	//3.tag 没有
	//输出
	bs := new(bytes.Buffer)
	bs.WriteString(fmt.Sprintf("%s\n", "type AutoDTO struct {"))
	for _, l := range ls {
		bs.WriteString(fmt.Sprintf("\t%s%s%s\n", l.Name, l.Type, l.Tag))
	}
	bs.WriteString(fmt.Sprintf("%s\n", "}"))
	fmt.Printf("%s\n", "")
	fmt.Printf("%s", bs.String())
	fmt.Printf("%s\n", "")
}

//sql 数据类型映射 go 类型
func getDataType(dbType string) string {
	switch strings.ToLower(dbType) {
	case "int":
		return "int64"
	case "varchar":
		return "string"
	case "decimal":
		return "float64"
	default:
		return "string"
	}
}

// 下划线写法转为驼峰写法
func underLineToUC(name string) string {
	//替换
	name = strings.Replace(name, "_", " ", -1)
	//首字母大写
	name = strings.Title(name)
	//缩进
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func uCFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func lCFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
