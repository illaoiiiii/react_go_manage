package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"os"
	"react_go_manage/models"
	"react_go_manage/utils"
	"strings"
)

type TypeController struct{}

type Page struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

func (con TypeController) List(c *gin.Context) {
	state, _ := utils.Int(c.Query("state"))
	orderId := strings.Trim(c.Query("orderId"), "")
	driverName := strings.Trim(c.Query("driverName"), "")
	userName := strings.Trim(c.Query("userName"), "")
	cityName := strings.Trim(c.Query("cityName"), "")

	pageNum, _ := utils.Int(c.Query("pageNum"))
	pageSize, _ := utils.Int(c.Query("pageSize"))

	var total int64
	var orderList []models.Order

	query := models.DB
	if orderId != "" {
		query = query.Where("order_id=?", orderId)
	}
	if driverName != "" {
		query = query.Where("driver_name=?", driverName)
	}
	if userName != "" {
		query = query.Where("user_name=?", userName)
	}
	if cityName != "" {
		query = query.Where("city_name=?", cityName)
	}
	if state != 0 {
		query = query.Where("state=?", state)
	}
	countQuery := query
	countQuery.Model(&models.Order{}).Count(&total)

	query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&orderList)

	//事实证明建立索引远远比开多协程快
	/*a := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex
	GoruntineNumber := 1
	orderListsCh := make(chan []models.Order)
	perSize := pageSize / GoruntineNumber
	for i := 0; i < GoruntineNumber; i++ {
		wg.Add(1)
		go func(ch chan<- []models.Order, i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			var orderLists []models.Order
			err := query.Limit(perSize).Offset((pageNum-1)*pageSize + i*perSize).Find(&orderLists).Error
			if err != nil {
				fmt.Println("Error querying users:", err)
				return
			}
			ch <- orderLists // 将查询结果发送到通道
		}(orderListsCh, i)
	}
	go func() {
		wg.Wait()
		close(orderListsCh) // 关闭通道，表示所有goroutine都已完成
	}()

	var allOrderLists []models.Order
	for orderList := range orderListsCh {
		allOrderLists = append(allOrderLists, orderList...)
	}
	fmt.Println("查询时间：", time.Since(a))*/
	//1个协程查询时间：  67.5623ms 1w条数据    618.7465ms 10w条数据
	//5个协程查询时间：  85.3694ms 1w条数据   886.6754ms 10w条数据
	//10个协程查询时间： 118.8443ms 1w条数据  1.4042216s 10w条数据

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"list": orderList,
			"page": gin.H{
				"pageNum":  pageNum,
				"pageSize": pageSize,
				"total":    total,
			},
		},
		"msg": "获取成功",
	})
}

func (con TypeController) VehicleList(c *gin.Context) {
	var vehicleList []models.Vehicle
	models.DB.Find(&vehicleList)
	c.JSON(200, gin.H{
		"code": 0,
		"data": vehicleList,
	})
}

func (con TypeController) CityList(c *gin.Context) {
	var cityList []models.City
	models.DB.Find(&cityList)
	c.JSON(200, gin.H{
		"code": 0,
		"data": cityList,
	})
}

func (con TypeController) OrderExport(c *gin.Context) {
	orderId := c.PostForm("orderId")
	userName := c.PostForm("userName")
	state, _ := utils.Int(c.PostForm("state"))
	fmt.Println(orderId, userName, state)
	//TODO 数据库查询然后把数据写入到excel中
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 设置头.
	f.SetCellValue("Sheet1", "B1", "订单编号")
	f.SetCellValue("Sheet1", "C1", "城市")
	f.SetCellValue("Sheet1", "D1", "下单地址（开始）")
	f.SetCellValue("Sheet1", "E1", "下单地址（结束）")
	f.SetCellValue("Sheet1", "F1", "下单时间")
	f.SetCellValue("Sheet1", "G1", "订单价格")
	f.SetCellValue("Sheet1", "H1", "订单状态")
	f.SetCellValue("Sheet1", "I1", "用户名称")
	f.SetCellValue("Sheet1", "J1", "司机名称")
	// 设置单元格宽度
	f.SetColWidth("Sheet1", "B", "J", 20)
	// 查询数据库
	var total int64
	var orderList []models.Order

	query := models.DB
	if orderId != "" {
		query = query.Where("order_id=?", orderId)
	}
	if userName != "" {
		query = query.Where("user_name=?", userName)
	}
	if state != 0 {
		query = query.Where("state=?", state)
	}
	countQuery := query
	countQuery.Model(&models.Order{}).Count(&total)
	query.Find(&orderList)
	// 设置具体的内容
	for i, order := range orderList {
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), order.OrderId)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), order.CityName)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), order.StartAddress)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), order.EndAddress)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), order.CreateTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), order.OrderAmount)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), formatState(order.State))
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), order.UserName)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), order.DriverName)
	}
	//文件处理格式
	filePath := "./static/upload/export.xlsx"

	// 检查文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		// 存在同名文件，删除它
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Failed to delete existing file:", err)
			return
		}
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	}

	// 直接写死export是导出的xlsx，而import是导入的xlsx
	// 检查文件是否存在
	_, err1 := os.Stat(filePath)
	if os.IsNotExist(err1) {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	// 发送文件给前端
	c.File(filePath)
}

func (con TypeController) OrderImport(c *gin.Context) {
	/*用老js的上传可以
	file, err := c.FormFile("face")

	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"dst":     dst,
	})*/
	authorization := c.GetHeader("Authorization")
	contentType := c.GetHeader("Content-Type")
	fmt.Println(authorization, contentType)
	// 检查请求方法是否为 POST
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	// 解析表单数据
	err := c.Request.ParseMultipartForm(50 << 20) // 设置最大内存为10MB
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 获取文件句柄
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}
	defer file.Close()
	// 打印文件信息
	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Type: %+v\n", handler.Header.Get("Content-Type"))

	//下载文件到本地指定目录
	// 由前端指定文件名称
	//c.SaveUploadedFile(handler, "./static/upload/"+handler.Filename)
	// 由后端指定文件名称 直接写死export是导出的xlsx，而import是导入的xlsx

	//文件处理格式
	filePath := "./static/upload/import.xlsx"
	// 检查文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		// 存在同名文件，删除它
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Failed to delete existing file:", err)
			return
		}
	}
	// 保存文件到指定路径好做后续处理
	c.SaveUploadedFile(handler, filePath)

	// 读取excel文件
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		// 处理一些数据
		rowSix, err := utils.Int(row[6])
		if err != nil {
			fmt.Println(err)
		}
		//不要第一行的介绍
		if i == 0 {
			fmt.Println(row)
			continue // 跳过第一行
		}
		order := models.Order{
			//OrderId:  这个之后在做什么分布式ID
			CityName:     row[2],
			StartAddress: row[3],
			EndAddress:   row[4],
			//TODO 之后这里要改成读取时间
			CreateTime:  utils.StringToTime(row[5]),
			OrderAmount: rowSix,
			State:       formatState1(row[7]),
			UserName:    row[8],
			DriverName:  row[9],
		}

		err1 := models.DB.Create(&order).Error
		if err1 != nil {
			c.JSON(400, gin.H{
				"msg": "导入数据失败",
			})
			return
		}
		fmt.Println()
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// 处理一些文字内容
func formatState(state int) string {
	switch state {
	case 1:
		return "未支付"
	case 2:
		return "已支付"
	case 3:
		return "已完成"
	case 4:
		return "已取消"
	default:
		return ""
	}
}

func formatState1(state string) int {
	switch state {
	case "未支付":
		return 1
	case "已支付":
		return 2
	case "已完成":
		return 3
	case "已取消":
		return 4
	default:
		return -1
	}
}
