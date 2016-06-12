// This file was generated by counterfeiter
package opspec

import (
	"sync"
)

type fakeCompositionRoot struct {
	CreateOpUseCaseStub        func() createOpUseCase
	createOpUseCaseMutex       sync.RWMutex
	createOpUseCaseArgsForCall []struct{}
	createOpUseCaseReturns     struct {
		result1 createOpUseCase
	}
	SetCollectionDescriptionUseCaseStub        func() setCollectionDescriptionUseCase
	setCollectionDescriptionUseCaseMutex       sync.RWMutex
	setCollectionDescriptionUseCaseArgsForCall []struct{}
	setCollectionDescriptionUseCaseReturns     struct {
		result1 setCollectionDescriptionUseCase
	}
	SetOpDescriptionUseCaseStub        func() setOpDescriptionUseCase
	setOpDescriptionUseCaseMutex       sync.RWMutex
	setOpDescriptionUseCaseArgsForCall []struct{}
	setOpDescriptionUseCaseReturns     struct {
		result1 setOpDescriptionUseCase
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCompositionRoot) CreateOpUseCase() createOpUseCase {
	fake.createOpUseCaseMutex.Lock()
	fake.createOpUseCaseArgsForCall = append(fake.createOpUseCaseArgsForCall, struct{}{})
	fake.recordInvocation("CreateOpUseCase", []interface{}{})
	fake.createOpUseCaseMutex.Unlock()
	if fake.CreateOpUseCaseStub != nil {
		return fake.CreateOpUseCaseStub()
	} else {
		return fake.createOpUseCaseReturns.result1
	}
}

func (fake *fakeCompositionRoot) CreateOpUseCaseCallCount() int {
	fake.createOpUseCaseMutex.RLock()
	defer fake.createOpUseCaseMutex.RUnlock()
	return len(fake.createOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) CreateOpUseCaseReturns(result1 createOpUseCase) {
	fake.CreateOpUseCaseStub = nil
	fake.createOpUseCaseReturns = struct {
		result1 createOpUseCase
	}{result1}
}

func (fake *fakeCompositionRoot) SetCollectionDescriptionUseCase() setCollectionDescriptionUseCase {
	fake.setCollectionDescriptionUseCaseMutex.Lock()
	fake.setCollectionDescriptionUseCaseArgsForCall = append(fake.setCollectionDescriptionUseCaseArgsForCall, struct{}{})
	fake.recordInvocation("SetCollectionDescriptionUseCase", []interface{}{})
	fake.setCollectionDescriptionUseCaseMutex.Unlock()
	if fake.SetCollectionDescriptionUseCaseStub != nil {
		return fake.SetCollectionDescriptionUseCaseStub()
	} else {
		return fake.setCollectionDescriptionUseCaseReturns.result1
	}
}

func (fake *fakeCompositionRoot) SetCollectionDescriptionUseCaseCallCount() int {
	fake.setCollectionDescriptionUseCaseMutex.RLock()
	defer fake.setCollectionDescriptionUseCaseMutex.RUnlock()
	return len(fake.setCollectionDescriptionUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) SetCollectionDescriptionUseCaseReturns(result1 setCollectionDescriptionUseCase) {
	fake.SetCollectionDescriptionUseCaseStub = nil
	fake.setCollectionDescriptionUseCaseReturns = struct {
		result1 setCollectionDescriptionUseCase
	}{result1}
}

func (fake *fakeCompositionRoot) SetOpDescriptionUseCase() setOpDescriptionUseCase {
	fake.setOpDescriptionUseCaseMutex.Lock()
	fake.setOpDescriptionUseCaseArgsForCall = append(fake.setOpDescriptionUseCaseArgsForCall, struct{}{})
	fake.recordInvocation("SetOpDescriptionUseCase", []interface{}{})
	fake.setOpDescriptionUseCaseMutex.Unlock()
	if fake.SetOpDescriptionUseCaseStub != nil {
		return fake.SetOpDescriptionUseCaseStub()
	} else {
		return fake.setOpDescriptionUseCaseReturns.result1
	}
}

func (fake *fakeCompositionRoot) SetOpDescriptionUseCaseCallCount() int {
	fake.setOpDescriptionUseCaseMutex.RLock()
	defer fake.setOpDescriptionUseCaseMutex.RUnlock()
	return len(fake.setOpDescriptionUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) SetOpDescriptionUseCaseReturns(result1 setOpDescriptionUseCase) {
	fake.SetOpDescriptionUseCaseStub = nil
	fake.setOpDescriptionUseCaseReturns = struct {
		result1 setOpDescriptionUseCase
	}{result1}
}

func (fake *fakeCompositionRoot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOpUseCaseMutex.RLock()
	defer fake.createOpUseCaseMutex.RUnlock()
	fake.setCollectionDescriptionUseCaseMutex.RLock()
	defer fake.setCollectionDescriptionUseCaseMutex.RUnlock()
	fake.setOpDescriptionUseCaseMutex.RLock()
	defer fake.setOpDescriptionUseCaseMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeCompositionRoot) recordInvocation(key string, args []interface{}) {
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