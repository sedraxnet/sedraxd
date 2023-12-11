package blockrelay

import (
	"github.com/sedraxnet/sedraxd/infrastructure/logger"
	"github.com/sedraxnet/sedraxd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
