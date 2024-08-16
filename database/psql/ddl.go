package psql

import (
	"app/config"

	"gorm.io/gorm"
	"fmt"
)

func RunDDLScript(db *gorm.DB){
	script := config.DDLScript()
	err := db.Exec(script).Error
	if err!=nil{
		panic(err)
	}else {
		fmt.Println("ddl script ran successfully")
	}
}

