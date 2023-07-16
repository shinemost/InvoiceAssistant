package main

import (
	// 引入gin包
	"InvoiceAssistant/dao"
	"InvoiceAssistant/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	conn := gin.New()

	conn.Use(func(c *gin.Context) {
		// 在请求开始时记录请求信息
		logger.Info("收到请求",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		// 处理请求
		c.Next()

		// 在请求结束时记录响应信息
		logger.Info("响应信息",
			zap.Int("status", c.Writer.Status()),
		)
	})

	// 路由分组
	r := conn.Group("/invoiceassistant/")
	{
		r.POST("scanner/:number", func(c *gin.Context) {

			//获取发票二维码解析后的数字
			number := c.Param("number")

			result := dao.InsertInvoceInfo(utils.NumberParser(number))

			//调用发票解析工具类方法
			c.JSON(http.StatusOK, gin.H{

				"message": fmt.Sprintf("发票扫描成功，主键为：%s", strconv.Itoa(int(result))),
			})
		})
	}
	conn.Run(":8080")

}
