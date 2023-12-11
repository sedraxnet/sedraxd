package rpchandlers

import (
	"github.com/sedraxnet/sedraxd/infrastructure/logger"
	"github.com/sedraxnet/sedraxd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
