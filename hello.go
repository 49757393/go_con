package main


//import "hello_con/tools"
import "hello_con/db"
import      "fmt"


func main() {

	//arr := tools.ReadCsv("/home/yangliang/data.csv")//读csv 并拼接成二维数组
	databe:=db.ConnectDB()
	//db.InsertUser(databe,arr)//插入数据
	brchs := db.Query(databe)
	
	for i := 0; i < len(brchs); i++ {
		
		fmt.Println(brchs[i].Acc + "       " + brchs[i].TranAate)
	}
	}