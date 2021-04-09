package zap

type BaseMap map[string]interface{}

func (log *log) Debug(msg string, data ...BaseMap) {
	if len(data) == 0 {
		log.logger.Debugw(msg)
	} else {
		log.logger.Debugw(msg, "data", data[0])
	}
}

func (log *log) Info(msg string, data ...BaseMap) {
	if len(data) == 0 {
		log.logger.Infow(msg)
	} else {
		log.logger.Infow(msg, "data", data[0])
	}
}

func (log *log) Warn(msg string, data ...BaseMap) {
	if len(data) == 0 {
		log.logger.Warnw(msg)
	} else {
		log.logger.Warnw(msg, "data", data[0])
	}
}

func (log *log) Error(msg string, data ...BaseMap) {
	if len(data) == 0 {
		log.logger.Errorw(msg)
	} else {
		log.logger.Errorw(msg, "data", data[0])
	}
}

func (log *log) Panic(msg string, data ...BaseMap) {
	if len(data) == 0 {
		log.logger.Panicw(msg)
	} else {
		log.logger.Panicw(msg, "data", data[0])
	}
}

func (log *log) Fatal(msg string, data ...BaseMap) {
	if len(data) == 0 {
		log.logger.Fatalw(msg)
	} else {
		log.logger.Fatalw(msg, "data", data[0])
	}
}
