package database

// 要想链接到数据库，首先需要加载目标数据库的驱动，驱动里面包含着与该数据库交互的逻辑。
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/denisenkon/go-mssqldb" // sqlserver 数据库驱动，引入时进行了自我注册，sql.Register("sqlserver",&DRV{})
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

// const (
// 	server = ""
// 	port = 1433
// 	user = "xxxx"
// 	password = "xxxx"
// 	database = "go-db"
// )

func mssqldb(server string,port string,user string,password string,database string) {
	// connStr := fmt.Sprintf("server=%s; user id=%s;password=%s;port=%d;database=%s;",server,user,password,port,database)

	// "user:password@/dbname"
	connStr := fmt.Sprintf("%s:%s@/%s",user,password,database)

	db, err := sql.Open("mysql",connStr) 
	// Open函数需要两个参数，数据库驱动的名称和数据源（链接字符串）的名称，得到一个指向sql.DB这个struct的指针。open函数并不会连接数据库，甚至不会验证其参数。它只是把后续连接到数据库所必须的structs给设置好了，而真正的连接是在被需要的时候才进行懒设置的。
	// sql.DB 是用来操作数据库的，它代表了0个或者多个底层连接的池，这些连接由sql包来维护，sql包会自动的创建和释放这些连接。它对于多个goroutine并发的使用是安全的。sql.DB不需要进行关闭（当然你想关闭也是可以的），它是用来处理数据库的，而不是实际的连接。这个抽象包含了数据库连接的池，而且会对此进行维护。在使用sql.DB的时候，可以定义它的全局变量进行使用，也可以将它传递函数/方法里。
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Connected!")
}

// 如何获得驱动？
// 正常的做法是使用sql.Register()函数、数据库驱动的名称和一个实现了driver.Driver接口的struct，来注册数据库的驱动。例如：
// sql.Register("sqlserver",&DRV{})
// 当go-mssqldb包被引入的时候，它的init函数将会运行并进行自我注册（在go语言里，每个包的init函数都会在自动的调用）
// 在引入go-mssqldb包的时候，把该包的名设置为下划线_ ，这是因为我们不直接使用数据库驱动（只需要它起的“副作用”），我们只使用database/sql
// Go语言没有提供官方的数据库驱动，所有的数据库驱动都是第三方驱动，但是他们都遵循sql.driver包里面定义的接口
// go get github.com/demisenkom/go-mssqldb