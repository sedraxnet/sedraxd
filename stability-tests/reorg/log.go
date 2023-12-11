package main

import (
	"github.com/sedraxnet/sedraxd/infrastructure/logger"
	"github.com/sedraxnet/sedraxd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
