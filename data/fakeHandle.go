// Code generated by counterfeiter. DO NOT EDIT.
package data

import (
	"context"
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type FakeHandle struct {
	ListDescendantsStub        func(ctx context.Context) ([]*model.DirEntry, error)
	listDescendantsMutex       sync.RWMutex
	listDescendantsArgsForCall []struct {
		ctx context.Context
	}
	listDescendantsReturns struct {
		result1 []*model.DirEntry
		result2 error
	}
	listDescendantsReturnsOnCall map[int]struct {
		result1 []*model.DirEntry
		result2 error
	}
	GetContentStub        func(ctx context.Context, contentPath string) (model.ReadSeekCloser, error)
	getContentMutex       sync.RWMutex
	getContentArgsForCall []struct {
		ctx         context.Context
		contentPath string
	}
	getContentReturns struct {
		result1 model.ReadSeekCloser
		result2 error
	}
	getContentReturnsOnCall map[int]struct {
		result1 model.ReadSeekCloser
		result2 error
	}
	PathStub        func() *string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 *string
	}
	pathReturnsOnCall map[int]struct {
		result1 *string
	}
	RefStub        func() string
	refMutex       sync.RWMutex
	refArgsForCall []struct{}
	refReturns     struct {
		result1 string
	}
	refReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandle) ListDescendants(ctx context.Context) ([]*model.DirEntry, error) {
	fake.listDescendantsMutex.Lock()
	ret, specificReturn := fake.listDescendantsReturnsOnCall[len(fake.listDescendantsArgsForCall)]
	fake.listDescendantsArgsForCall = append(fake.listDescendantsArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("ListDescendants", []interface{}{ctx})
	fake.listDescendantsMutex.Unlock()
	if fake.ListDescendantsStub != nil {
		return fake.ListDescendantsStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDescendantsReturns.result1, fake.listDescendantsReturns.result2
}

func (fake *FakeHandle) ListDescendantsCallCount() int {
	fake.listDescendantsMutex.RLock()
	defer fake.listDescendantsMutex.RUnlock()
	return len(fake.listDescendantsArgsForCall)
}

func (fake *FakeHandle) ListDescendantsArgsForCall(i int) context.Context {
	fake.listDescendantsMutex.RLock()
	defer fake.listDescendantsMutex.RUnlock()
	return fake.listDescendantsArgsForCall[i].ctx
}

func (fake *FakeHandle) ListDescendantsReturns(result1 []*model.DirEntry, result2 error) {
	fake.ListDescendantsStub = nil
	fake.listDescendantsReturns = struct {
		result1 []*model.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeHandle) ListDescendantsReturnsOnCall(i int, result1 []*model.DirEntry, result2 error) {
	fake.ListDescendantsStub = nil
	if fake.listDescendantsReturnsOnCall == nil {
		fake.listDescendantsReturnsOnCall = make(map[int]struct {
			result1 []*model.DirEntry
			result2 error
		})
	}
	fake.listDescendantsReturnsOnCall[i] = struct {
		result1 []*model.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeHandle) GetContent(ctx context.Context, contentPath string) (model.ReadSeekCloser, error) {
	fake.getContentMutex.Lock()
	ret, specificReturn := fake.getContentReturnsOnCall[len(fake.getContentArgsForCall)]
	fake.getContentArgsForCall = append(fake.getContentArgsForCall, struct {
		ctx         context.Context
		contentPath string
	}{ctx, contentPath})
	fake.recordInvocation("GetContent", []interface{}{ctx, contentPath})
	fake.getContentMutex.Unlock()
	if fake.GetContentStub != nil {
		return fake.GetContentStub(ctx, contentPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getContentReturns.result1, fake.getContentReturns.result2
}

func (fake *FakeHandle) GetContentCallCount() int {
	fake.getContentMutex.RLock()
	defer fake.getContentMutex.RUnlock()
	return len(fake.getContentArgsForCall)
}

func (fake *FakeHandle) GetContentArgsForCall(i int) (context.Context, string) {
	fake.getContentMutex.RLock()
	defer fake.getContentMutex.RUnlock()
	return fake.getContentArgsForCall[i].ctx, fake.getContentArgsForCall[i].contentPath
}

func (fake *FakeHandle) GetContentReturns(result1 model.ReadSeekCloser, result2 error) {
	fake.GetContentStub = nil
	fake.getContentReturns = struct {
		result1 model.ReadSeekCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeHandle) GetContentReturnsOnCall(i int, result1 model.ReadSeekCloser, result2 error) {
	fake.GetContentStub = nil
	if fake.getContentReturnsOnCall == nil {
		fake.getContentReturnsOnCall = make(map[int]struct {
			result1 model.ReadSeekCloser
			result2 error
		})
	}
	fake.getContentReturnsOnCall[i] = struct {
		result1 model.ReadSeekCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeHandle) Path() *string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pathReturns.result1
}

func (fake *FakeHandle) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeHandle) PathReturns(result1 *string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeHandle) PathReturnsOnCall(i int, result1 *string) {
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeHandle) Ref() string {
	fake.refMutex.Lock()
	ret, specificReturn := fake.refReturnsOnCall[len(fake.refArgsForCall)]
	fake.refArgsForCall = append(fake.refArgsForCall, struct{}{})
	fake.recordInvocation("Ref", []interface{}{})
	fake.refMutex.Unlock()
	if fake.RefStub != nil {
		return fake.RefStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.refReturns.result1
}

func (fake *FakeHandle) RefCallCount() int {
	fake.refMutex.RLock()
	defer fake.refMutex.RUnlock()
	return len(fake.refArgsForCall)
}

func (fake *FakeHandle) RefReturns(result1 string) {
	fake.RefStub = nil
	fake.refReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHandle) RefReturnsOnCall(i int, result1 string) {
	fake.RefStub = nil
	if fake.refReturnsOnCall == nil {
		fake.refReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.refReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHandle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listDescendantsMutex.RLock()
	defer fake.listDescendantsMutex.RUnlock()
	fake.getContentMutex.RLock()
	defer fake.getContentMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.refMutex.RLock()
	defer fake.refMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandle) recordInvocation(key string, args []interface{}) {
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

var _ model.DataHandle = new(FakeHandle)
