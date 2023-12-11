package grpcclient

import (
	"github.com/sedraxnet/sedraxd/infrastructure/logger"
	"github.com/sedraxnet/sedraxd/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
