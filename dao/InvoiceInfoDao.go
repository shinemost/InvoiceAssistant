/*
 * @Author: shinemost supertain147@163.com
 * @Date: 2023-07-16 21:35:27
 * @LastEditors: shinemost supertain147@163.com
 * @LastEditTime: 2023-07-16 21:55:35
 * @FilePath: /InvoiceAssistant/dao/InvoiceInfoDao.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"InvoiceAssistant/db"
	"InvoiceAssistant/model"
)

func InsertInvoceInfo(invoice *model.InvoiceInfo) uint {

	db.DB.Create(invoice)
	return invoice.ID
}
