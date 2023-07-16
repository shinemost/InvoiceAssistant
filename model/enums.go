/*
 * @Author: shinemost supertain147@163.com
 * @Date: 2023-07-16 22:17:00
 * @LastEditors: shinemost supertain147@163.com
 * @LastEditTime: 2023-07-16 23:15:57
 * @FilePath: /InvoiceAssistant/model/enums.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

//go:generate stringer -type=InvoiceType
type InvoiceType uint64

const (
	增值税专用发票   InvoiceType = 1
	增值税普通发票   InvoiceType = 4
	增值税专用电子发票 InvoiceType = 8
	增值税普通电子发票 InvoiceType = 10
)
