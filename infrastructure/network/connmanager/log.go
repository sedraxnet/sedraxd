package connmanager

import (
	"github.com/sedraxnet/sedraxd/infrastructure/logger"
	"github.com/sedraxnet/sedraxd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
