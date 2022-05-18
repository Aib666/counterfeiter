$ cat tools/tools.go
// +build tools

package tools

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)

// This file imports packages that are used when running go generate, or used
// during the development process but not otherwise depended on by built code.
)

type FakeContext struct {
	DoSomethingStub        func()
	doSomethingMutex       sync.RWMutex
	doSomethingArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContext) DoSomething() {
	fake.doSomethingMutex.Lock()
	fake.doSomethingArgsForCall = append(fake.doSomethingArgsForCall, struct {
	}{})
	stub := fake.DoSomethingStub
	fake.recordInvocation("DoSomething", []interface{}{})
	fake.doSomethingMutex.Unlock()
	if stub != nil {
		fake.DoSomethingStub()
	}
}

func (fake *FakeContext) DoSomethingCallCount() int {
	fake.doSomethingMutex.RLock()
	defer fake.doSomethingMutex.RUnlock()
	return len(fake.doSomethingArgsForCall)
}

func (fake *FakeContext) DoSomethingCalls(stub func()) {
	fake.doSomethingMutex.Lock()
	defer fake.doSomethingMutex.Unlock()
	fake.DoSomethingStub = stub
}

func (fake *FakeContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doSomethingMutex.RLock()
	defer fake.doSomethingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContext) recordInvocation(key string, args []interface{}) {
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

var _ internalpkg.Context = new(FakeContext)
