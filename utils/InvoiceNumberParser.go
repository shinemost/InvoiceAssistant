/*
 * @Author: shinemost supertain147@163.com
 * @Date: 2023-07-16 20:42:44
 * @LastEditors: shinemost supertain147@163.com
 * @LastEditTime: 2023-07-16 23:19:52
 * @FilePath: /InvoiceAssistant/utils/InvoiceNumberParser.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import (
	"InvoiceAssistant/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NumberParser(invoiceNumber string) *model.InvoiceInfo {

	// 01,10,034002200111,96767252,15.92,20230712,14591557031922141366,C841,
	result := strings.Split(invoiceNumber, ",")

	if len(result) == 1 {
		err := errors.New("解析发票报错，请检查发票二维码！")
		fmt.Println("错误信息：", err)
	}
	// 获取结果的第二个元素
	num, err := strconv.ParseUint(result[1], 10, 64)
	if err != nil {
		fmt.Println("无法将字符串转换为无符号整数")
	}

	invoiceType := model.InvoiceType(num)

	code := result[2]
	number := result[3]
	amount := result[4]
	input_date := result[5]
	validateNumber := result[6]

	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Printf("解析发票金额报错：%s\n", err)

	}

	t, err := time.Parse("20060102", input_date)
	if err != nil {
		fmt.Printf("解析开票日期报错：%s\n", err)
	}

	formattedStr := t.Format("2006-01-02")

	return &model.InvoiceInfo{
		Code:           code,
		Number:         number,
		Amount:         f,
		Type:           invoiceType.String(),
		Date:           formattedStr,
		ValidateNumber: validateNumber,
		InputPerson:    "admin",
	}
}
