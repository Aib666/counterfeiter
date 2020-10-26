// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"net/http"
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures"
	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures/another_package"
)

type FakeEmbedsInterfaces struct {
	AnotherMethodStub        func([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType)
	anotherMethodMutex       sync.RWMutex
	anotherMethodArgsForCall []struct {
		arg1 []another_package.SomeType
		arg2 map[another_package.SomeType]another_package.SomeType
		arg3 *another_package.SomeType
		arg4 another_package.SomeType
		arg5 chan another_package.SomeType
	}
	DoThingsStub        func()
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
	}
	EmbeddedMethodStub        func() string
	embeddedMethodMutex       sync.RWMutex
	embeddedMethodArgsForCall []struct {
	}
	embeddedMethodReturns struct {
		result1 string
	}
	embeddedMethodReturnsOnCall map[int]struct {
		result1 string
	}
	ServeHTTPStub        func(http.ResponseWriter, *http.Request)
	serveHTTPMutex       sync.RWMutex
	serveHTTPArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmbedsInterfaces) AnotherMethod(arg1 []another_package.SomeType, arg2 map[another_package.SomeType]another_package.SomeType, arg3 *another_package.SomeType, arg4 another_package.SomeType, arg5 chan another_package.SomeType) {
	var arg1Copy []another_package.SomeType
	if arg1 != nil {
		arg1Copy = make([]another_package.SomeType, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.anotherMethodMutex.Lock()
	fake.anotherMethodArgsForCall = append(fake.anotherMethodArgsForCall, struct {
		arg1 []another_package.SomeType
		arg2 map[another_package.SomeType]another_package.SomeType
		arg3 *another_package.SomeType
		arg4 another_package.SomeType
		arg5 chan another_package.SomeType
	}{arg1Copy, arg2, arg3, arg4, arg5})
	fake.recordInvocation("AnotherMethod", []interface{}{arg1Copy, arg2, arg3, arg4, arg5})
	fake.anotherMethodMutex.Unlock()
	if fake.AnotherMethodStub != nil {
		fake.AnotherMethodStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakeEmbedsInterfaces) AnotherMethodCallCount() int {
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	return len(fake.anotherMethodArgsForCall)
}

func (fake *FakeEmbedsInterfaces) AnotherMethodCalls(stub func([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType)) {
	fake.anotherMethodMutex.Lock()
	defer fake.anotherMethodMutex.Unlock()
	fake.AnotherMethodStub = stub
}

func (fake *FakeEmbedsInterfaces) AnotherMethodArgsForCall(i int) ([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType) {
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	argsForCall := fake.anotherMethodArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeEmbedsInterfaces) DoThings() {
	fake.doThingsMutex.Lock()
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
	}{})
	fake.recordInvocation("DoThings", []interface{}{})
	fake.doThingsMutex.Unlock()
	if fake.DoThingsStub != nil {
		fake.DoThingsStub()
	}
}

func (fake *FakeEmbedsInterfaces) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeEmbedsInterfaces) DoThingsCalls(stub func()) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = stub
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethod() string {
	fake.embeddedMethodMutex.Lock()
	ret, specificReturn := fake.embeddedMethodReturnsOnCall[len(fake.embeddedMethodArgsForCall)]
	fake.embeddedMethodArgsForCall = append(fake.embeddedMethodArgsForCall, struct {
	}{})
	fake.recordInvocation("EmbeddedMethod", []interface{}{})
	fake.embeddedMethodMutex.Unlock()
	if fake.EmbeddedMethodStub != nil {
		return fake.EmbeddedMethodStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.embeddedMethodReturns
	return fakeReturns.result1
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethodCallCount() int {
	fake.embeddedMethodMutex.RLock()
	defer fake.embeddedMethodMutex.RUnlock()
	return len(fake.embeddedMethodArgsForCall)
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethodCalls(stub func() string) {
	fake.embeddedMethodMutex.Lock()
	defer fake.embeddedMethodMutex.Unlock()
	fake.EmbeddedMethodStub = stub
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethodReturns(result1 string) {
	fake.embeddedMethodMutex.Lock()
	defer fake.embeddedMethodMutex.Unlock()
	fake.EmbeddedMethodStub = nil
	fake.embeddedMethodReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethodReturnsOnCall(i int, result1 string) {
	fake.embeddedMethodMutex.Lock()
	defer fake.embeddedMethodMutex.Unlock()
	fake.EmbeddedMethodStub = nil
	if fake.embeddedMethodReturnsOnCall == nil {
		fake.embeddedMethodReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.embeddedMethodReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEmbedsInterfaces) ServeHTTP(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.serveHTTPMutex.Lock()
	fake.serveHTTPArgsForCall = append(fake.serveHTTPArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("ServeHTTP", []interface{}{arg1, arg2})
	fake.serveHTTPMutex.Unlock()
	if fake.ServeHTTPStub != nil {
		fake.ServeHTTPStub(arg1, arg2)
	}
}

func (fake *FakeEmbedsInterfaces) ServeHTTPCallCount() int {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	return len(fake.serveHTTPArgsForCall)
}

func (fake *FakeEmbedsInterfaces) ServeHTTPCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.serveHTTPMutex.Lock()
	defer fake.serveHTTPMutex.Unlock()
	fake.ServeHTTPStub = stub
}

func (fake *FakeEmbedsInterfaces) ServeHTTPArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	argsForCall := fake.serveHTTPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEmbedsInterfaces) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	fake.embeddedMethodMutex.RLock()
	defer fake.embeddedMethodMutex.RUnlock()
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmbedsInterfaces) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.EmbedsInterfaces = new(FakeEmbedsInterfaces)
