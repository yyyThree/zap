package zap

type BaseMap map[string]interface{}

func (logger *Logger) Debug(msg string, data ...BaseMap) {
	if len(data) == 0 {
		logger.logger.Debugw(msg)
	} else {
		logger.logger.Debugw(msg, "data", data[0])
	}
}

func (logger *Logger) Info(msg string, data ...BaseMap) {
	if len(data) == 0 {
		logger.logger.Infow(msg)
	} else {
		logger.logger.Infow(msg, "data", data[0])
	}
}

func (logger *Logger) Warn(msg string, data ...BaseMap) {
	if len(data) == 0 {
		logger.logger.Warnw(msg)
	} else {
		logger.logger.Warnw(msg, "data", data[0])
	}
}

func (logger *Logger) Error(msg string, data ...BaseMap) {
	if len(data) == 0 {
		logger.logger.Errorw(msg)
	} else {
		logger.logger.Errorw(msg, "data", data[0])
	}
}

func (logger *Logger) Panic(msg string, data ...BaseMap) {
	if len(data) == 0 {
		logger.logger.Panicw(msg)
	} else {
		logger.logger.Panicw(msg, "data", data[0])
	}
}

func (logger *Logger) Fatal(msg string, data ...BaseMap) {
	if len(data) == 0 {
		logger.logger.Fatalw(msg)
	} else {
		logger.logger.Fatalw(msg, "data", data[0])
	}
}
