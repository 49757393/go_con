package db

//import "hello_con/tools"
import  "database/sql"
import      "fmt"
//import      "time"

import	_ "github.com/lib/pq"
	



const (
    host     = "localhost"
    port     = 5432
    user     = "yangliang"
    password = "123456"
    dbname   = "exampledb"
)

type BrchQryDtl struct {
    TranAate  string
	Timestampl string
	Acc    string
	Amt string
	DrCrFlag string
	RptSum string
	
}

func ConnectDB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}


func InsertUser(db *sql.DB,arr [][]string)  {

	
	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		stmt,err := db.Prepare("insert into brch_qry_dtl(tran_date,timestampl,acc,amt,dr_cr_flag,rpt_sum) values(cast($1 as date),$2, $3, cast($4 as numeric), cast($5 as integer), $6)")
		if err != nil {
			//log.Fatal(err)
		}
		_,err = stmt.Exec(arr[i][0],arr[i][1],arr[i][2],arr[i][3],arr[i][4],arr[i][5])
		if err != nil {
			//log.Fatal(err)
		}else {
			//fmt.Println("insert into brch_qry_dtl success")
		}
	}	
}

func Query(db *sql.DB) []BrchQryDtl{
	
	var arr []BrchQryDtl
	var brch BrchQryDtl
	rows,err:=db.Query(" select tran_date,timestampl,acc from brch_qry_dtl ")

	if err!= nil{
		fmt.Println(err)
	}
	defer rows.Close()
//根据数据库的查询结果将结果集写到对象里面 并放入数组
	for rows.Next(){
		err:= rows.Scan(&brch.TranAate,&brch.Timestampl,&brch.Acc)

		if err!= nil{
			fmt.Println(err)
		}
		arr = append(arr, brch)
	}
	return arr
	
}



	
	
