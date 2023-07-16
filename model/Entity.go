/*
 * @Author: shinemost supertain147@163.com
 * @Date: 2023-07-16 20:28:14
 * @LastEditors: shinemost supertain147@163.com
 * @LastEditTime: 2023-07-16 22:03:34
 * @FilePath: /InvoiceAssistant/model/Entity.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"gorm.io/gorm"
)

type InvoiceInfo struct {
	Code           string  `gorm:"size:255" json:"code"`
	Number         string  `gorm:"size:255" json:"number"`
	Date           string  `gorm:"size:10" json:"date"`
	ValidateNumber string  `gorm:"size:255" json:"validate_number"`
	Type           string  `gorm:"size:255" json:"type"`
	Amount         float64 `json:"amount"`
	InputPerson    string  `gorm:"size:255" json:"input_person"`
	Remark         string  `gorm:"size:1000" json:"remark"`
	Status         string  `gorm:"size:255" json:"status"`
	gorm.Model
}

func (InvoiceInfo) TableName() string {
	return "invoice_info"
}
