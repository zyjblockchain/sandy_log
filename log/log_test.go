package log

import (
	"testing"
)

func TestWarn(t *testing.T) {
	Setup("aa", LevelWarn, false, false)
	Debug("should be invisible")
	Info("should be invisible")
	Warn("汉字")
	Error("👽🚀")

	Setup("bb", LevelDebug, false, false)
	Debug("debug level visible")
	Info("info level visible")
	srvLog.GetHandler()
	Setup("cc", LevelDebug, false, true)
	Debug("show code line")
	Info("show code line")

	// Crit("critical error")
}

func TestEventf(t *testing.T) {
	Setup("", LevelWarn, false, false)
	for i := 0; i < 10; i++ {
		Eventf(TxEvent, "test log %d", i)
		Event(TxEvent, "test")
	}
}

func init() {
	Setup("test_module", LevelDebug, true, false)
}
func TestDebugf(t *testing.T) {
	Debugf("test01aaaaa: %d", 122)
	Infof("test02bbb: %s", "sand")
	Warnf("test03ccc: %v", "099")
	Errorf("test04ddd: %f", 0.998)
	sign := "tx sign"
	Warn("invalid length of signture", "sign", string(sign))
	Error("invalid length of signture", "sign", string(sign))
	Info("invalid length of signture", "sign", string(sign))
	Debug("invalid length of signture", "sign", string(sign))
}

func TestCrit(t *testing.T) {
	log := NewLog("log-test", LevelDebug, true, true)
	log.Info("aa", "aa", 22)
	log2 := NewLog("log2-test", LevelDebug, true, true)
	log2.Info("bb", "bb", 22)
}
