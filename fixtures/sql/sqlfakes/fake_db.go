// Code generated by counterfeiter. DO NOT EDIT.
package sqlfakes

import (
	sqla "database/sql"
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures/sql"
)

type FakeDB struct {
	ExecStub        func(string, ...interface{}) (sqla.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	execReturns struct {
		result1 sqla.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sqla.Result
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDB) Exec(arg1 string, arg2 ...interface{}) (sqla.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDB) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeDB) ExecCalls(stub func(string, ...interface{}) (sqla.Result, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeDB) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDB) ExecReturns(result1 sqla.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sqla.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeDB) ExecReturnsOnCall(i int, result1 sqla.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sqla.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sqla.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ sql.DB = new(FakeDB)
