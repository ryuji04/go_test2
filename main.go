package main

import (
	"database/sql"
	"fmt"
	"hello"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

// 出力メソッド
func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\"" + m.Mail + "," +
		strconv.Itoa(m.Age) + ">"
}

// レコード情報の取り出し
func cursor(md *Mydata, rs *sql.Rows) *Mydata {
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return md
}

// 全件検索
func findAll(con *sql.DB) {
	rs, er := con.Query(all)
	if er != nil {
		panic(er)
	}
	for rs.Next() {
		var md Mydata
		md = *cursor(&md, rs)
		fmt.Println(md.Str())
	}

	fmt.Println("end")
}

// ID検索
func findUserById(con *sql.DB) {
	for true {
		s := hello.Input("id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}

		rs, er := con.Query(qry, n)
		if er != nil {
			panic(er)
		}
		for rs.Next() {
			var md Mydata
			md = *cursor(&md, rs)
			fmt.Println(md.Str())
		}

	}
	fmt.Println("end")
}

// 名前orメール検索
func findUserByNameOrMail(con *sql.DB) {
	for true {
		s := hello.Input("findUserByNameOrMail")
		if s == "" {
			break
		}
		rs, er := con.Query(qry2, "%"+s+"%", "%"+s+"%")
		if er != nil {
			panic(er)
		}
		for rs.Next() {
			var md Mydata
			md = *cursor(&md, rs)
			fmt.Println(md.Str())
		}

		fmt.Println("end")
	}
}

// insertメソッド
func insert(con *sql.DB) {
	nm := hello.Input("name")
	ml := hello.Input("mail")
	age := hello.Input("age")
	ag, _ := strconv.Atoi(age)

	con.Exec(qry3, nm, ml, ag)
	fmt.Println("入力完了")
}

// 構造体の生成
func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

// updateメソッド
func update(con *sql.DB) {
	ids := hello.Input("update ID")             //アップデートするIDを入力
	id, _ := strconv.Atoi(ids)                  //stringからintに変換
	rw := con.QueryRow(qry, id)                 //クエリの結果を格納
	tgt := mydatafmRw(rw)                       //クエリの結果から構造体を生成
	ae := strconv.Itoa(tgt.Age)                 //intからstringに変換
	nm := hello.Input("name(" + tgt.Name + ")") //変更後のユーザー名を入力
	ml := hello.Input("mail(" + tgt.Mail + ")") //変更後のメールアドレスを入力
	ge := hello.Input("age(" + ae + ")")        //変更後の年齢を入力
	ag, _ := strconv.Atoi(ge)

	//入力がなければ変更を加えない
	if nm == "" {
		nm = tgt.Name
	}
	if ml == "" {
		ml = tgt.Mail
	}
	if ge == "" {
		ag = tgt.Age
	}

	con.Exec(qry4, nm, ml, ag, id)
	fmt.Println("更新完了")

}

// deleteメソッド
func delete(con *sql.DB) {
	idd := hello.Input("select id for deleting")
	id, er := strconv.Atoi(idd)
	if er != nil {
		panic(er)
	}

	con.Exec(qry5, id)
	fmt.Println("更新完了")
}

// クエリ文
var all string = "select * from mydata"
var qry string = "select * from mydata where id=?"
var qry2 string = "select * from mydata where name like ? or mail like ?"
var qry3 string = "insert into  mydata (name,mail,age) values(?,?,?)"
var qry4 string = "update mydata set name=?,mail=?,age=? where id=?"
var qry5 string = "delete from mydata where id=?"

func main2() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	flg := hello.Input("all or id or other or insert or update or delete ")

	if flg == "all" {
		findAll(con)
	} else if flg == "id" {
		findUserById(con)
	} else if flg == "other" {
		findUserByNameOrMail(con)
	} else if flg == "insert" {
		insert(con)
	} else if flg == "update" {
		update(con)
	} else if flg == "delete" {
		delete(con)
	} else {
		fmt.Println("error")
		return
	}

}
