package order

import (
	"github.com/gin-gonic/gin"
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
