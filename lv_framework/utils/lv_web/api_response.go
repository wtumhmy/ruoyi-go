package lv_web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"robvi/app/system/model"
)

// 通用api响应
type ApiResp struct {
	r *model.CommonRes
	c *gin.Context
}

// 返回一个成功的消息体
func Sucess(c *gin.Context, data interface{}) {
	msg := model.CommonRes{
		Code: 200,
		Data: data,
		Msg:  "操作成功",
	}
	c.JSON(http.StatusOK, msg)
	c.Abort()
}

// 返回一个成功的消息体
func Fail(c *gin.Context, msg string) {
	ret := model.CommonRes{
		Code: 500,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, ret)
	c.Abort()
}

// 返回一个成功的消息体
func SucessResp(c *gin.Context) *ApiResp {
	msg := model.CommonRes{
		Code:  200,
		Btype: model.Buniss_Other,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// 返回一个错误的消息体
func ErrorResp(c *gin.Context) *ApiResp {
	msg := model.CommonRes{
		Code:  500,
		Btype: model.Buniss_Other,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// 返回一个拒绝访问的消息体
func ForbiddenResp(c *gin.Context) *ApiResp {
	msg := model.CommonRes{
		Code:  403,
		Btype: model.Buniss_Other,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		r: &msg,
		c: c,
	}
	return &a
}

// 设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.r.Msg = msg
	return resp
}

// 设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.r.Code = code
	return resp
}

// 设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.r.Data = data
	return resp
}

// 设置消息体的业务类型
func (resp *ApiResp) SetBtype(btype model.BunissType) *ApiResp {
	resp.r.Btype = btype
	return resp
}

// Deprecated: 不再使用
// punctuation properly. 不再使用的代码.
func (resp *ApiResp) Log(title string, inParam interface{}) *ApiResp {
	fmt.Println("ApiResp.Log 已经废弃，" + title)
	return resp
}

// 输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.c.JSON(http.StatusOK, resp.r)
	resp.c.Abort()
}
