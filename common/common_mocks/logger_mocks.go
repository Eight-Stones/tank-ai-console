package commonmocks

import (
	"go-micro-service-template/common"
)

// MockLogger объект логгера, который необходим для моков в тестах.
type MockLogger struct{}

func (m MockLogger) Debug(_ ...interface{}) {}

func (m MockLogger) Info(_ ...interface{}) {}

func (m MockLogger) Warn(_ ...interface{}) {}

func (m MockLogger) Error(_ ...interface{}) {}

func (m MockLogger) Debugf(_ string, _ ...interface{}) {}

func (m MockLogger) Infof(_ string, _ ...interface{}) {}

func (m MockLogger) Warnf(_ string, _ ...interface{}) {}

func (m MockLogger) Errorf(_ string, _ ...interface{}) {}

func (m MockLogger) Print(_ ...interface{}) {}

func (m MockLogger) Printf(_ string, _ ...interface{}) {}

func (m MockLogger) Name(_ string) common.LoggerI { return m }

func (m MockLogger) Fields(_ ...any) common.LoggerI { return m }
