// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures"
)

type FakeInlineStructParams struct {
	DoSomethingStub func(context.Context, struct {
		SomeString        string
		SomeStringPointer *string
		SomeTime          time.Time
		SomeTimePointer   *time.Time
		HTTPRequest       http.Request
	}) error
	doSomethingMutex       sync.RWMutex
	doSomethingArgsForCall []struct {
		arg1 context.Context
		arg2 struct {
			SomeString        string
			SomeStringPointer *string
			SomeTime          time.Time
			SomeTimePointer   *time.Time
			HTTPRequest       http.Request
		}
	}
	doSomethingReturns struct {
		result1 error
	}
	doSomethingReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInlineStructParams) DoSomething(arg1 context.Context, arg2 struct {
	SomeString        string
	SomeStringPointer *string
	SomeTime          time.Time
	SomeTimePointer   *time.Time
	HTTPRequest       http.Request
}) error {
	fake.doSomethingMutex.Lock()
	ret, specificReturn := fake.doSomethingReturnsOnCall[len(fake.doSomethingArgsForCall)]
	fake.doSomethingArgsForCall = append(fake.doSomethingArgsForCall, struct {
		arg1 context.Context
		arg2 struct {
			SomeString        string
			SomeStringPointer *string
			SomeTime          time.Time
			SomeTimePointer   *time.Time
			HTTPRequest       http.Request
		}
	}{arg1, arg2})
	fake.recordInvocation("DoSomething", []interface{}{arg1, arg2})
	fake.doSomethingMutex.Unlock()
	if fake.DoSomethingStub != nil {
		return fake.DoSomethingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doSomethingReturns
	return fakeReturns.result1
}

func (fake *FakeInlineStructParams) DoSomethingCallCount() int {
	fake.doSomethingMutex.RLock()
	defer fake.doSomethingMutex.RUnlock()
	return len(fake.doSomethingArgsForCall)
}

func (fake *FakeInlineStructParams) DoSomethingCalls(stub func(context.Context, struct {
	SomeString        string
	SomeStringPointer *string
	SomeTime          time.Time
	SomeTimePointer   *time.Time
	HTTPRequest       http.Request
}) error) {
	fake.doSomethingMutex.Lock()
	defer fake.doSomethingMutex.Unlock()
	fake.DoSomethingStub = stub
}

func (fake *FakeInlineStructParams) DoSomethingArgsForCall(i int) (context.Context, struct {
	SomeString        string
	SomeStringPointer *string
	SomeTime          time.Time
	SomeTimePointer   *time.Time
	HTTPRequest       http.Request
}) {
	fake.doSomethingMutex.RLock()
	defer fake.doSomethingMutex.RUnlock()
	argsForCall := fake.doSomethingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInlineStructParams) DoSomethingReturns(result1 error) {
	fake.doSomethingMutex.Lock()
	defer fake.doSomethingMutex.Unlock()
	fake.DoSomethingStub = nil
	fake.doSomethingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInlineStructParams) DoSomethingReturnsOnCall(i int, result1 error) {
	fake.doSomethingMutex.Lock()
	defer fake.doSomethingMutex.Unlock()
	fake.DoSomethingStub = nil
	if fake.doSomethingReturnsOnCall == nil {
		fake.doSomethingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doSomethingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInlineStructParams) Invocations() map[string][][]interface{} {
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

func (fake *FakeInlineStructParams) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.InlineStructParams = new(FakeInlineStructParams)
