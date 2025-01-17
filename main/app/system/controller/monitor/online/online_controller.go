package online

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_web"
	"robvi/app/common/global"
	"robvi/app/system/model/monitor/online"
	"robvi/app/system/service"
	onlineService "robvi/app/system/service/monitor/online"
	"strings"
)

// 列表页
func List(c *gin.Context) {
	sessinIdArr := make([]string, 0)

	global.SessionList.Range(func(k, v interface{}) bool {
		return true
	})
	if len(sessinIdArr) > 0 {
		onlineService.DeleteRecordNotInIds(sessinIdArr)
	}

	lv_web.BuildTpl(c, "monitor/online/list").WriteTpl()
}

// 列表分页数据
func ListAjax(c *gin.Context) {
	var req *online.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	rows := make([]online.UserOnline, 0)
	result, page, err := onlineService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 用户强退
func ForceLogout(c *gin.Context) {
	sessionId := c.PostForm("sessionId")
	if sessionId == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误").Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}
	var userService service.UserService
	err := userService.ForceLogout(sessionId)
	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).Log("用户强退", gin.H{"sessionId": sessionId}).WriteJsonExit()
}

// 批量强退
func BatchForceLogout(c *gin.Context) {
	var userService service.UserService
	ids := c.Query("ids")
	if ids == "" {
		lv_web.ErrorResp(c).SetMsg("参数错误").Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
		return
	}
	ids = strings.ReplaceAll(ids, "[", "")
	ids = strings.ReplaceAll(ids, "]", "")
	ids = strings.ReplaceAll(ids, `"`, "")
	idarr := strings.Split(ids, ",")
	if len(idarr) > 0 {
		for _, sessionId := range idarr {
			if sessionId != "" {
				userService.ForceLogout(sessionId)
			}
		}
	}
	lv_web.SucessResp(c).Log("批量强退", gin.H{"ids": ids}).WriteJsonExit()
}
