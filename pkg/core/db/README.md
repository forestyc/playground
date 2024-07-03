# db

## 1. 使用

```
package main

import (
	"github.com/forestyc/playground/pkg/core/db"
)

func main() {
	mdb := db.NewMysql(db.Config{
		Dsn:              "baal:Baal@123@tcp(140.143.163.171:3306)/baal?charset=utf8mb4&parseTime=true&loc=Local",
		MaxOpen:          10,
		IdleConns:        5,
		MaxIdleTime:      300,
		OperationTimeout: 10,
	})
	defer mdb.Close()
	var user User
	session := mdb.Session()
	result := session.Model(&user).
		//Where("date=? and exchange=? and contract_id=?", "20231212", "DCE", "a2312").
		Take(&user)
	if code, msg := mdb.DBError(result.Error); code > 0 {
		fmt.Println(msg)
	}
}
```

