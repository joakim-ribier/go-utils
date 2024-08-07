package logsutil

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joakim-ribier/go-utils/pkg/iosutil"
)

var workingDirectory string

func TestMain(m *testing.M) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	workingDirectory = dir

	exitVal := m.Run()

	os.Exit(exitVal)
}

// TestNewLoggerWithInvalidFile calls logger.NewLogger,
// checking for a valid return value.
func TestNewLoggerWithInvalidFile(t *testing.T) {
	r, err := NewLogger("/wrong-folder/go-utils.log", "go-utils")
	if err == nil {
		t.Fatalf(`result: {%v} but expected error`, r)
	}
}

// TestLogInfo calls logger.Info,
// checking for a valid return value.
func TestLogInfo(t *testing.T) {
	logger, err := NewLogger(workingDirectory+"/go-utils.log", "go-utils")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(workingDirectory + "/go-utils.log")

	logger.Info("Hello World", "arg0", "value0", "arg1", "value1")

	data, err := iosutil.Load(workingDirectory + "/go-utils.log")
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(data), `Hello World	{"app": "go-utils", "arg0": "value0", "arg1": "value1"}`) {
		t.Fatalf(`result: {%v} but expected INFO`, string(data))
	}
}

// TestLogInfoWithNamespace calls logger.Namespace,
// checking for a valid return value.
func TestLogInfoWithNamespace(t *testing.T) {
	logger, err := NewLogger(workingDirectory+"/go-utils.log", "go-utils")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(workingDirectory + "/go-utils.log")

	logger.Namespace("my-package").Info("Hello World", "arg0", "value0", "arg1", "value1")

	data, err := iosutil.Load(workingDirectory + "/go-utils.log")
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(data), `Hello World	{"app": "go-utils", "namespace": "my-package", "arg0": "value0", "arg1": "value1"}`) {
		t.Fatalf(`result: {%v} but expected INFO`, string(data))
	}
}

// TestLogError calls logger.Error,
// checking for a valid return value.
func TestLogError(t *testing.T) {
	logger, err := NewLogger(workingDirectory+"/go-utils.log", "go-utils")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(workingDirectory + "/go-utils.log")

	logger.Error(errors.New("io exception"), "file does not exist", "arg0", "value0", "arg1", "value1")

	data, err := iosutil.Load(workingDirectory + "/go-utils.log")
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(data), `file does not exist	{"app": "go-utils", "arg0": "value0", "arg1": "value1", "error": "io exception"}`) {
		t.Fatalf(`result: {%v} but expected ERROR`, string(data))
	}
}
