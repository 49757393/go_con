编译的时候 如果用到了相对路径的包 那么将在外层编译
go mod init hello_con
go build db/con_db.go
golang的语法：模块中要导出的函数，必须首字母大写
引用函数的方式 package.函数名