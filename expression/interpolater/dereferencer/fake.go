// Code generated by counterfeiter. DO NOT EDIT.
package dereferencer

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	DeReferenceStub        func(ref string, scope map[string]*model.Value, pkgHandle model.PkgHandle) (string, bool, error)
	deReferenceMutex       sync.RWMutex
	deReferenceArgsForCall []struct {
		ref       string
		scope     map[string]*model.Value
		pkgHandle model.PkgHandle
	}
	deReferenceReturns struct {
		result1 string
		result2 bool
		result3 error
	}
	deReferenceReturnsOnCall map[int]struct {
		result1 string
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) DeReference(ref string, scope map[string]*model.Value, pkgHandle model.PkgHandle) (string, bool, error) {
	fake.deReferenceMutex.Lock()
	ret, specificReturn := fake.deReferenceReturnsOnCall[len(fake.deReferenceArgsForCall)]
	fake.deReferenceArgsForCall = append(fake.deReferenceArgsForCall, struct {
		ref       string
		scope     map[string]*model.Value
		pkgHandle model.PkgHandle
	}{ref, scope, pkgHandle})
	fake.recordInvocation("DeReference", []interface{}{ref, scope, pkgHandle})
	fake.deReferenceMutex.Unlock()
	if fake.DeReferenceStub != nil {
		return fake.DeReferenceStub(ref, scope, pkgHandle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.deReferenceReturns.result1, fake.deReferenceReturns.result2, fake.deReferenceReturns.result3
}

func (fake *Fake) DeReferenceCallCount() int {
	fake.deReferenceMutex.RLock()
	defer fake.deReferenceMutex.RUnlock()
	return len(fake.deReferenceArgsForCall)
}

func (fake *Fake) DeReferenceArgsForCall(i int) (string, map[string]*model.Value, model.PkgHandle) {
	fake.deReferenceMutex.RLock()
	defer fake.deReferenceMutex.RUnlock()
	return fake.deReferenceArgsForCall[i].ref, fake.deReferenceArgsForCall[i].scope, fake.deReferenceArgsForCall[i].pkgHandle
}

func (fake *Fake) DeReferenceReturns(result1 string, result2 bool, result3 error) {
	fake.DeReferenceStub = nil
	fake.deReferenceReturns = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *Fake) DeReferenceReturnsOnCall(i int, result1 string, result2 bool, result3 error) {
	fake.DeReferenceStub = nil
	if fake.deReferenceReturnsOnCall == nil {
		fake.deReferenceReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
			result3 error
		})
	}
	fake.deReferenceReturnsOnCall[i] = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deReferenceMutex.RLock()
	defer fake.deReferenceMutex.RUnlock()
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

var _ DeReferencer = new(Fake)
