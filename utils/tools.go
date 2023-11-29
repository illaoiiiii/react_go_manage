package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// DateToUnix 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetUnix 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetUnixNano 获取纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// GetDate 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// GetDateWithZone 获取当前日期加上时区
func GetDateWithZone() time.Time {
	//模版
	template := "2006-01-02 15:04:05"
	//实际要输出的时间
	t, err := time.Parse(template, time.Now().Format(template))
	if err != nil {
		// 处理解析错误
		fmt.Println("err 时间解析错误", err)
	}
	return t
}

// GetDay 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// Md5 md5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Str2Html 把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// Int 表示把string转换成int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// Float 表示把string转换成Float64
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

// String 表示把int转换成string
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}

// StringToArray 把字符串转换成数组
func StringToArray(str string) []string {
	// 使用正则表达式提取数字部分
	str = strings.Trim(str, "[]")
	matches := strings.Split(str, ",")
	// 将字符串转换为整数并放入切片中
	result := make([]string, len(matches))
	for i, v := range matches {
		result[i] = v
	}
	return result
}

// Substr Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])

}

//Oss上传
/*func OssUplod(file *multipart.FileHeader, dst string) (string, error) {

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "GJoqWHXB2c9S9gwP", "Lgf3weXuWITUUb17vDJfveg1jmKEe9")
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket("beego")
	if err != nil {
		return "", err
	}

	// 上传文件流。
	err = bucket.PutObject(dst, f)
	if err != nil {
		return "", err
	}
	return dst, nil
}*/

//通过列获取值
/*func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}*/

// 获取Oss的状态
/*func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := Int(config.Section("oss").Key("status").String())
	return ossStatus
}
*/
//格式化输出图片
/*func FormatImg(str string) string {
	ossStatus := GetOssStatus()
	fmt.Println(ossStatus)
	if ossStatus == 1 {
		return GetSettingFromColumn("OssDomain") + str
	} else {
		return "/" + str
	}
}*/

func Sub(a int, b int) int {
	return a - b
}

func Mul(price float64, num int) float64 {
	return price * float64(num)
}

//上传图片
/*func UploadImg(c *gin.Context, picName string) (string, error) {
	ossStatus := GetOssStatus()
	if ossStatus == 1 {
		return OssUploadImg(c, picName)
	} else {
		return LocalUploadImg(c, picName)
	}
}*/

//  上传图片到Oss
/*func OssUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、定义图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "static/upload/" + day

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)

	OssUplod(file, dst)
	return dst, nil
}
*/
// LocalUploadImg 上传图片到本地
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	// fmt.Println(file)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil

}

//生成商品缩略图
/*func ResizeGoodsImage(filename string) {
	extname := path.Ext(filename)
	ThumbnailSize := strings.ReplaceAll(GetSettingFromColumn("ThumbnailSize"), "，", ",")
	thumbnailSizeSlice := strings.Split(ThumbnailSize, ",")
	//static/upload/tao_400.png
	//static/upload/tao_400.png_100x100.png
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		savepath := filename + "_" + thumbnailSizeSlice[i] + "x" + thumbnailSizeSlice[i] + extname
		w, _ := Int(thumbnailSizeSlice[i])
		err := ThumbnailF2F(filename, savepath, w, w)
		if err != nil {
			fmt.Println(err) //写个日志模块  处理日志
		}
	}
}*/

// 生成随机数
func GetRandomNum() string {
	var str string

	for i := 0; i < 4; i++ {
		current := rand.Intn(10)

		str += String(current)
	}
	return str
}

// GetOrderId
func GetOrderId() string {
	// 2022020312233
	template := "20060102150405"
	return time.Now().Format(template) + GetRandomNum()
}
