package log

import (
	"github.com/enoch300/glog"
	"github.com/sirupsen/logrus"
	"ippool_api/utils"
	"path/filepath"
)

var GlobalLog *logrus.Logger

func NewLogger(save uint) {
	logPath := filepath.Dir(utils.GetCurrentAbPath()) + "/logs"
	GlobalLog = glog.NewLogger(logPath, "ip_pool_api", save)
}
