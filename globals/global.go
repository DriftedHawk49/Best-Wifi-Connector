package globals

import "go.uber.org/zap"

// Needs to be initialized as soon as logger is created for use in every module
var Logger *zap.Logger
