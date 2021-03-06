// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeDisableOrgIsolationActor struct {
	DeleteIsolationSegmentOrganizationByNameStub        func(string, string) (v3action.Warnings, error)
	deleteIsolationSegmentOrganizationByNameMutex       sync.RWMutex
	deleteIsolationSegmentOrganizationByNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteIsolationSegmentOrganizationByNameReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	deleteIsolationSegmentOrganizationByNameReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDisableOrgIsolationActor) DeleteIsolationSegmentOrganizationByName(arg1 string, arg2 string) (v3action.Warnings, error) {
	fake.deleteIsolationSegmentOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.deleteIsolationSegmentOrganizationByNameReturnsOnCall[len(fake.deleteIsolationSegmentOrganizationByNameArgsForCall)]
	fake.deleteIsolationSegmentOrganizationByNameArgsForCall = append(fake.deleteIsolationSegmentOrganizationByNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteIsolationSegmentOrganizationByName", []interface{}{arg1, arg2})
	fake.deleteIsolationSegmentOrganizationByNameMutex.Unlock()
	if fake.DeleteIsolationSegmentOrganizationByNameStub != nil {
		return fake.DeleteIsolationSegmentOrganizationByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteIsolationSegmentOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDisableOrgIsolationActor) DeleteIsolationSegmentOrganizationByNameCallCount() int {
	fake.deleteIsolationSegmentOrganizationByNameMutex.RLock()
	defer fake.deleteIsolationSegmentOrganizationByNameMutex.RUnlock()
	return len(fake.deleteIsolationSegmentOrganizationByNameArgsForCall)
}

func (fake *FakeDisableOrgIsolationActor) DeleteIsolationSegmentOrganizationByNameCalls(stub func(string, string) (v3action.Warnings, error)) {
	fake.deleteIsolationSegmentOrganizationByNameMutex.Lock()
	defer fake.deleteIsolationSegmentOrganizationByNameMutex.Unlock()
	fake.DeleteIsolationSegmentOrganizationByNameStub = stub
}

func (fake *FakeDisableOrgIsolationActor) DeleteIsolationSegmentOrganizationByNameArgsForCall(i int) (string, string) {
	fake.deleteIsolationSegmentOrganizationByNameMutex.RLock()
	defer fake.deleteIsolationSegmentOrganizationByNameMutex.RUnlock()
	argsForCall := fake.deleteIsolationSegmentOrganizationByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDisableOrgIsolationActor) DeleteIsolationSegmentOrganizationByNameReturns(result1 v3action.Warnings, result2 error) {
	fake.deleteIsolationSegmentOrganizationByNameMutex.Lock()
	defer fake.deleteIsolationSegmentOrganizationByNameMutex.Unlock()
	fake.DeleteIsolationSegmentOrganizationByNameStub = nil
	fake.deleteIsolationSegmentOrganizationByNameReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDisableOrgIsolationActor) DeleteIsolationSegmentOrganizationByNameReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.deleteIsolationSegmentOrganizationByNameMutex.Lock()
	defer fake.deleteIsolationSegmentOrganizationByNameMutex.Unlock()
	fake.DeleteIsolationSegmentOrganizationByNameStub = nil
	if fake.deleteIsolationSegmentOrganizationByNameReturnsOnCall == nil {
		fake.deleteIsolationSegmentOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.deleteIsolationSegmentOrganizationByNameReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDisableOrgIsolationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteIsolationSegmentOrganizationByNameMutex.RLock()
	defer fake.deleteIsolationSegmentOrganizationByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDisableOrgIsolationActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.DisableOrgIsolationActor = new(FakeDisableOrgIsolationActor)
