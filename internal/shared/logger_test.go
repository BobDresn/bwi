package shared

import (
	"bytes",
	"strings",
	"testing",
	"regexp",
)

func TestLogLevelAllowsOutput(t *testing.T) {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	SetLogLevel(LevelInfo)

	Debug("Debug should NOT appear")
	Info("info should appear")

	output := buf.String()
	if strings.Contains(output, "Debug should NOT appear") {
		t.Error("Debug message appeared despite LevelInfo")
	}
	if !strings.Contains(output, "info should appear") {
		t.Error("Info message did not appear")
	}
}

func TestLogIncludesFileLine(t *testing.T) {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	SetLogLevel(LevelDebug)

	Debug("testing file and line")

	output := buf.String()
	if !strings.Contains(output, "logger_test.go") {
		t.Error("Log output missing file info")
	}
	if !strings.Contains(output, ":") {
		t.Error("Log output missing line number")
	}
}

func TestTimeFormat(t *testing.T) {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	SetLogLevel(levelDebug)

	Debug("check timestamp")

	output := buf.String()
	re := regexp.MustCompile(`\[\d{2}:\d{2}:\d{2}.\d{3}\]`)
	if !re.MatchString(output) {
		t.Error("Timestamp format invalid or missing")
	}
}
