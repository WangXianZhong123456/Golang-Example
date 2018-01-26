package main

import (
	"database/sql"
	"fmt"
	"strings"
)

import (
	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
	port   int
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	conf = append(conf, "Initial Catalog="+m.database)

	// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器
	// (需要在安装sqlserver时候设置)，
	// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
	if m.windows {
		conf = append(conf, "integrated security=SSPI")
	} else {
		conf = append(conf, "user id = "+m.sa.user)
		conf = append(conf, "password = "+m.sa.passwd)
		conf = append(conf, "port = "+fmt.Sprint(m.sa.port))
	}

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	//fmt.Println(strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db := Mssql{
		// 如果数据库是默认实例（MSSQLSERVER）则直接使用IP，命名实例需要指明。
		dataSource: "172.16.2.33",
		database:   "FCMS_BSDB",
		// windows: true 为windows身份验证，false 必须设置sa账号和密码
		windows: false,
		sa: SA{
			user:   "sa",
			passwd: "yx_3.0",
			port:   1433,
		},
	}
	// 连接数据库
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()

	// 执行SQL语句

	/*1、PreparedStatement 可以实现自定义参数的查询
	2、PreparedStatement 通常来说, 比手动拼接字符串 SQL 语句高效.
	3、PreparedStatement 可以防止SQL注入攻击*/
	/*stmt, _ := db.Prepare(`UPDATE UCML_Organize SET OrgName = ? WHERE flag = 3 AND PersonGroupID = ?`)

	rows, err := stmt.Query("研发部", 48)
	if err != nil {
		fmt.Println("Prepare: ", err)
		return
	}
	var result int
	rows.Scan(&result)
	fmt.Printf("insert result %v\n", result)
	rows.Close()*/
	rows, err := db.Query("select top 10 orgName from ucml_organize where flag = 3")

	for rows.Next() {
		// 查询结果字段和声明变量数量相等，否则数据为空。
		var orgName string
		rows.Scan(&orgName)
		fmt.Printf("orgName: %s\n", orgName)
		// var number int
		// rows.Scan(&name, &number)
		// fmt.Printf("Name: %s \t Number: %d\n", name, number)
	}

	/*
		返回不定字段
		columns, err := rows.Columns()

		values := make([]sql.RawBytes, len(columns))
		scans := make([]interface{}, len(columns))

		for i := range values {
			scans[i] = &values[i]
		}

		var result []map[string]string
		for rows.Next() {
			_ = rows.Scan(scans...)
			each := make(map[string]string)

			for i, col := range values {
				each[columns[i]] = string(col)
			}

			result = append(result, each)

		}

		fmt.Println(result[2]["OrgName"])*/

}
