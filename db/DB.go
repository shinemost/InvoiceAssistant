/*
 * @Author: shinemost supertain147@163.com
 * @Date: 2023-07-16 20:32:34
 * @LastEditors: shinemost supertain147@163.com
 * @LastEditTime: 2023-07-16 21:52:04
 * @FilePath: /InvoiceAssistant/db/DB.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:test@tcp(localhost:3306)/invoiceassistant?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
