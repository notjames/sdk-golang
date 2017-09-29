// Code generated by counterfeiter. DO NOT EDIT.
package files

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	InterpretStub        func(pkgHandle model.PkgHandle, scope map[string]*model.Value, scgContainerCallFiles map[string]interface{}, scratchDirPath string) (map[string]string, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		pkgHandle             model.PkgHandle
		scope                 map[string]*model.Value
		scgContainerCallFiles map[string]interface{}
		scratchDirPath        string
	}
	interpretReturns struct {
		result1 map[string]string
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Interpret(pkgHandle model.PkgHandle, scope map[string]*model.Value, scgContainerCallFiles map[string]interface{}, scratchDirPath string) (map[string]string, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		pkgHandle             model.PkgHandle
		scope                 map[string]*model.Value
		scgContainerCallFiles map[string]interface{}
		scratchDirPath        string
	}{pkgHandle, scope, scgContainerCallFiles, scratchDirPath})
	fake.recordInvocation("Interpret", []interface{}{pkgHandle, scope, scgContainerCallFiles, scratchDirPath})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(pkgHandle, scope, scgContainerCallFiles, scratchDirPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *Fake) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *Fake) InterpretArgsForCall(i int) (model.PkgHandle, map[string]*model.Value, map[string]interface{}, string) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].pkgHandle, fake.interpretArgsForCall[i].scope, fake.interpretArgsForCall[i].scgContainerCallFiles, fake.interpretArgsForCall[i].scratchDirPath
}

func (fake *Fake) InterpretReturns(result1 map[string]string, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *Fake) InterpretReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ Files = new(Fake)
