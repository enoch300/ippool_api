package log

import (
	"github.com/enoch300/glog"
	"github.com/enoch300/ippool_client/utils"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

var GlobalLog *logrus.Logger

func NewLogger(save uint) {
	logPath := filepath.Dir(utils.GetCurrentAbPath()) + "/logs"
	GlobalLog = glog.NewLogger(logPath, "ip_pool_api", save)
}
