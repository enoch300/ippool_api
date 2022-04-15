package api

import (
	"github.com/gin-gonic/gin"
	"ippool_api/controller/peers"
	. "ippool_api/utils/log"
	"net/http"
	"strings"
)

type query struct {
	Network   string `form:"network"`
	Appid     string `form:"appid"`
	Machineid string `form:"machineid"`
}

type response struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func GetPeers(c *gin.Context) {
	var q query
	if err := c.ShouldBindQuery(&q); err != nil {
		GlobalLog.Errorf("request bind query %s", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	results, err := peers.Get(q.Network, q.Appid, q.Machineid)
	if err != nil {
		GlobalLog.Errorf("get peers %s", err)
		c.JSON(http.StatusOK, gin.H{
			"code":      1,
			"network":   q.Network,
			"appid":     q.Appid,
			"machineid": q.Machineid,
			"msg":       "获取peers出错:" + err.Error(),
			"data":      results,
		})
		return
	}

	var r []response
	for _, info := range results {
		strArray := strings.Split(info, "_")
		r = append(r, response{
			Ip:   strArray[0],
			Port: strArray[1],
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"network":   q.Network,
		"appid":     q.Appid,
		"machineid": q.Machineid,
		"msg":       "success",
		"data":      r,
	})
}
