// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeNetworkPoliciesActor struct {
	NetworkPoliciesBySpaceAndAppNameStub        func(spaceGUID string, srcAppName string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error)
	networkPoliciesBySpaceAndAppNameMutex       sync.RWMutex
	networkPoliciesBySpaceAndAppNameArgsForCall []struct {
		spaceGUID  string
		srcAppName string
	}
	networkPoliciesBySpaceAndAppNameReturns struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	networkPoliciesBySpaceAndAppNameReturnsOnCall map[int]struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	NetworkPoliciesBySpaceStub        func(spaceGUID string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error)
	networkPoliciesBySpaceMutex       sync.RWMutex
	networkPoliciesBySpaceArgsForCall []struct {
		spaceGUID string
	}
	networkPoliciesBySpaceReturns struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	networkPoliciesBySpaceReturnsOnCall map[int]struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppName(spaceGUID string, srcAppName string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error) {
	fake.networkPoliciesBySpaceAndAppNameMutex.Lock()
	ret, specificReturn := fake.networkPoliciesBySpaceAndAppNameReturnsOnCall[len(fake.networkPoliciesBySpaceAndAppNameArgsForCall)]
	fake.networkPoliciesBySpaceAndAppNameArgsForCall = append(fake.networkPoliciesBySpaceAndAppNameArgsForCall, struct {
		spaceGUID  string
		srcAppName string
	}{spaceGUID, srcAppName})
	fake.recordInvocation("NetworkPoliciesBySpaceAndAppName", []interface{}{spaceGUID, srcAppName})
	fake.networkPoliciesBySpaceAndAppNameMutex.Unlock()
	if fake.NetworkPoliciesBySpaceAndAppNameStub != nil {
		return fake.NetworkPoliciesBySpaceAndAppNameStub(spaceGUID, srcAppName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.networkPoliciesBySpaceAndAppNameReturns.result1, fake.networkPoliciesBySpaceAndAppNameReturns.result2, fake.networkPoliciesBySpaceAndAppNameReturns.result3
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameCallCount() int {
	fake.networkPoliciesBySpaceAndAppNameMutex.RLock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.RUnlock()
	return len(fake.networkPoliciesBySpaceAndAppNameArgsForCall)
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameArgsForCall(i int) (string, string) {
	fake.networkPoliciesBySpaceAndAppNameMutex.RLock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.RUnlock()
	return fake.networkPoliciesBySpaceAndAppNameArgsForCall[i].spaceGUID, fake.networkPoliciesBySpaceAndAppNameArgsForCall[i].srcAppName
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameReturns(result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.NetworkPoliciesBySpaceAndAppNameStub = nil
	fake.networkPoliciesBySpaceAndAppNameReturns = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameReturnsOnCall(i int, result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.NetworkPoliciesBySpaceAndAppNameStub = nil
	if fake.networkPoliciesBySpaceAndAppNameReturnsOnCall == nil {
		fake.networkPoliciesBySpaceAndAppNameReturnsOnCall = make(map[int]struct {
			result1 []cfnetworkingaction.Policy
			result2 cfnetworkingaction.Warnings
			result3 error
		})
	}
	fake.networkPoliciesBySpaceAndAppNameReturnsOnCall[i] = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpace(spaceGUID string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error) {
	fake.networkPoliciesBySpaceMutex.Lock()
	ret, specificReturn := fake.networkPoliciesBySpaceReturnsOnCall[len(fake.networkPoliciesBySpaceArgsForCall)]
	fake.networkPoliciesBySpaceArgsForCall = append(fake.networkPoliciesBySpaceArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("NetworkPoliciesBySpace", []interface{}{spaceGUID})
	fake.networkPoliciesBySpaceMutex.Unlock()
	if fake.NetworkPoliciesBySpaceStub != nil {
		return fake.NetworkPoliciesBySpaceStub(spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.networkPoliciesBySpaceReturns.result1, fake.networkPoliciesBySpaceReturns.result2, fake.networkPoliciesBySpaceReturns.result3
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceCallCount() int {
	fake.networkPoliciesBySpaceMutex.RLock()
	defer fake.networkPoliciesBySpaceMutex.RUnlock()
	return len(fake.networkPoliciesBySpaceArgsForCall)
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceArgsForCall(i int) string {
	fake.networkPoliciesBySpaceMutex.RLock()
	defer fake.networkPoliciesBySpaceMutex.RUnlock()
	return fake.networkPoliciesBySpaceArgsForCall[i].spaceGUID
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceReturns(result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.NetworkPoliciesBySpaceStub = nil
	fake.networkPoliciesBySpaceReturns = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceReturnsOnCall(i int, result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.NetworkPoliciesBySpaceStub = nil
	if fake.networkPoliciesBySpaceReturnsOnCall == nil {
		fake.networkPoliciesBySpaceReturnsOnCall = make(map[int]struct {
			result1 []cfnetworkingaction.Policy
			result2 cfnetworkingaction.Warnings
			result3 error
		})
	}
	fake.networkPoliciesBySpaceReturnsOnCall[i] = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.networkPoliciesBySpaceAndAppNameMutex.RLock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.RUnlock()
	fake.networkPoliciesBySpaceMutex.RLock()
	defer fake.networkPoliciesBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetworkPoliciesActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.NetworkPoliciesActor = new(FakeNetworkPoliciesActor)