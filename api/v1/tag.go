package v1

import (
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询标签是否存在

// 添加标签
func AddTag(c *gin.Context) {
	var data model.Tag
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, code := model.CheckTag(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateTag(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个标签下的文章
func GetArticlesByTag(c *gin.Context) {
	// var articles []model.Article
	// tagName := c.Param("tag")
}

// 查询标签列表
func GetTags(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data := model.GetTags(pageSize, pageNum)
	code := errmsg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑标签
func EditTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var data model.Tag
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, code := model.CheckTag(data.Name)
	if code == errmsg.ERROR_TAGNAME_USED {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
		c.Abort()
	}

	code = model.EditTag(uint(id), &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除标签
func DeleteTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	code := model.DeleteTag(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
