package ReadExcel

import (
	"common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadExcel(c *gin.Context) {
	//接收前端发来的文件
	//file, _, err := c.Request.FormFile("file")
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    common.ERROR,
	//		"message": "文件读取失败，请按格式填写excel",
	//	})
	//	return
	//}
	////读取文件中的内容
	//xlsx, err := excelize.OpenReader(file)
	//fmt.Println(xlsx)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    common.ERROR,
	//		"message": "文件读取失败，请按格式填写excel",
	//	})
	//	return
	//}

	//index := 2
	//for true {
	//	Variable := "A" + strconv.Itoa(index)
	//	order_id,_ := xlsx.GetCellValue("Sheet1", Variable) //12345678909876
	//	if order_id == "" {
	//		index--
	//		break
	//	}
	//	index++
	//	continue
	//}

	//Subscript := 0
	//result := make([]interface{}, index-1) //定义一个空切片
	//value := make(map[string]interface{})  //定义一个map
	//for {
	//	index_A := "A" + strconv.Itoa(Subscript+2)
	//	index_B := "B" + strconv.Itoa(Subscript+2)
	//	index_C := "C" + strconv.Itoa(Subscript+2)
	//	orderId,_ := xlsx.GetCellValue("Sheet1", index_A) //12345678909876
	//	if orderId == "" {
	//
	//		break
	//	}
	//postWay,_ := xlsx.GetCellValue("Sheet1", index_B)        //申通快递
	//postId,_ := xlsx.GetCellValue("Sheet1", index_C)         //3860480177227
	////err, res := comments.DeliveryShowOrder(orderId) //查询show_order表，转model层
	////错误捕捉
	//if err != nil {
	//	//logging.Error(err)
	//	value = map[string]interface{}{
	//		"code":    common.ERROR,
	//		"message": "该订单数字id错误",
	//		"orderId": orderId,
	//	}
	//	result[Subscript] = value
	//	Subscript++
	//	continue
	//}
	//订单状态不为3或4判断
	//if 3 > res.Status || res.Status > 4 {
	//	value = map[string]interface{}{
	//		"code":    common.ERROR,
	//		"message": "该订单状态不是待发货或发货状态,",
	//		"orderId": orderId,
	//		"status":  res.Status,
	//	}
	//	result[Subscript] = value
	//	Subscript++
	//	continue
	//}
	//修改状态信息，改为已经发货
	//err = comments.UpdateShowOrder1(orderId, postId, postWay)
	//if err != nil {
	//	value = map[string]interface{}{
	//		"code":    common.ERROR,
	//		"message": "订单状态修改为已发货状态失败,",
	//		"orderId": orderId,
	//		"status":  res.Status,
	//	}
	//	result[Subscript] = value
	//	Subscript++
	//	continue
	//}
	//for循环内部返回数据
	//value = map[string]interface{}{
	//	"code":    common.SUCCESS,
	//	"status":  4,
	//	"orderId": orderId,
	//	"postWay": postWay,
	//	"postId":  postId,
	//}
	//result[Subscript] = value
	//Subscript++

	//}
	//*********返回给前端的数据
	c.JSON(http.StatusOK, gin.H{
		"code":    common.SUCCESS,
		"message": "excel文件导入成功",
	})
}
