package testy

import (
	"time"

	"github.com/gametimesf/testy/orderedmap"
)

type packageTests = orderedmap.OrderedMap[string, testCase]
type packageTestResults = orderedmap.OrderedMap[string, TestResult]
type allTests = orderedmap.OrderedMap[string, packageTests]
type allTestResults = orderedmap.OrderedMap[string, packageTestResults]

type testy struct {
	tests allTests
}

type testCase struct {
	Package string
	Name    string
	tester  Tester
}

type Tester func(t TestingT)

type TestResult struct {
	Package string
	Name    string
	Msgs    []Msg
	Passed  bool
	Dur     time.Duration
}

type Level string

const (
	LevelInfo  Level = "info"
	LevelError Level = "error"
)

type Msg struct {
	Msg   string
	Level Level
}

var instance testy

// TestingT is a subset of testing.T that we have to implement for non-`go test` runs.
//
// TODO flesh this out with more useful stuff from testing.T -- Parallel would be nice but tricky
type TestingT interface {
	Fail()
	FailNow()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Run(string, Tester)
}
