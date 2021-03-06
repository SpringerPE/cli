// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7pushaction"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/util/manifestparser"
)

type FakePushActor struct {
	ActualizeStub        func(v7pushaction.PushPlan, v7pushaction.ProgressBar) (<-chan v7pushaction.PushPlan, <-chan v7pushaction.Event, <-chan v7pushaction.Warnings, <-chan error)
	actualizeMutex       sync.RWMutex
	actualizeArgsForCall []struct {
		arg1 v7pushaction.PushPlan
		arg2 v7pushaction.ProgressBar
	}
	actualizeReturns struct {
		result1 <-chan v7pushaction.PushPlan
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}
	actualizeReturnsOnCall map[int]struct {
		result1 <-chan v7pushaction.PushPlan
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}
	ConceptualizeStub        func([]string, string, string, string, v7pushaction.FlagOverrides) ([]v7pushaction.PushPlan, v7pushaction.Warnings, error)
	conceptualizeMutex       sync.RWMutex
	conceptualizeArgsForCall []struct {
		arg1 []string
		arg2 string
		arg3 string
		arg4 string
		arg5 v7pushaction.FlagOverrides
	}
	conceptualizeReturns struct {
		result1 []v7pushaction.PushPlan
		result2 v7pushaction.Warnings
		result3 error
	}
	conceptualizeReturnsOnCall map[int]struct {
		result1 []v7pushaction.PushPlan
		result2 v7pushaction.Warnings
		result3 error
	}
	CreatePushPlansStub        func(string, manifestparser.Parser, v7pushaction.FlagOverrides) ([]v7pushaction.PushPlan, error)
	createPushPlansMutex       sync.RWMutex
	createPushPlansArgsForCall []struct {
		arg1 string
		arg2 manifestparser.Parser
		arg3 v7pushaction.FlagOverrides
	}
	createPushPlansReturns struct {
		result1 []v7pushaction.PushPlan
		result2 error
	}
	createPushPlansReturnsOnCall map[int]struct {
		result1 []v7pushaction.PushPlan
		result2 error
	}
	PrepareSpaceStub        func(string, string, *manifestparser.Parser, v7pushaction.FlagOverrides) (<-chan []string, <-chan v7pushaction.Event, <-chan v7pushaction.Warnings, <-chan error)
	prepareSpaceMutex       sync.RWMutex
	prepareSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *manifestparser.Parser
		arg4 v7pushaction.FlagOverrides
	}
	prepareSpaceReturns struct {
		result1 <-chan []string
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}
	prepareSpaceReturnsOnCall map[int]struct {
		result1 <-chan []string
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePushActor) Actualize(arg1 v7pushaction.PushPlan, arg2 v7pushaction.ProgressBar) (<-chan v7pushaction.PushPlan, <-chan v7pushaction.Event, <-chan v7pushaction.Warnings, <-chan error) {
	fake.actualizeMutex.Lock()
	ret, specificReturn := fake.actualizeReturnsOnCall[len(fake.actualizeArgsForCall)]
	fake.actualizeArgsForCall = append(fake.actualizeArgsForCall, struct {
		arg1 v7pushaction.PushPlan
		arg2 v7pushaction.ProgressBar
	}{arg1, arg2})
	fake.recordInvocation("Actualize", []interface{}{arg1, arg2})
	fake.actualizeMutex.Unlock()
	if fake.ActualizeStub != nil {
		return fake.ActualizeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.actualizeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakePushActor) ActualizeCallCount() int {
	fake.actualizeMutex.RLock()
	defer fake.actualizeMutex.RUnlock()
	return len(fake.actualizeArgsForCall)
}

func (fake *FakePushActor) ActualizeCalls(stub func(v7pushaction.PushPlan, v7pushaction.ProgressBar) (<-chan v7pushaction.PushPlan, <-chan v7pushaction.Event, <-chan v7pushaction.Warnings, <-chan error)) {
	fake.actualizeMutex.Lock()
	defer fake.actualizeMutex.Unlock()
	fake.ActualizeStub = stub
}

func (fake *FakePushActor) ActualizeArgsForCall(i int) (v7pushaction.PushPlan, v7pushaction.ProgressBar) {
	fake.actualizeMutex.RLock()
	defer fake.actualizeMutex.RUnlock()
	argsForCall := fake.actualizeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePushActor) ActualizeReturns(result1 <-chan v7pushaction.PushPlan, result2 <-chan v7pushaction.Event, result3 <-chan v7pushaction.Warnings, result4 <-chan error) {
	fake.actualizeMutex.Lock()
	defer fake.actualizeMutex.Unlock()
	fake.ActualizeStub = nil
	fake.actualizeReturns = struct {
		result1 <-chan v7pushaction.PushPlan
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakePushActor) ActualizeReturnsOnCall(i int, result1 <-chan v7pushaction.PushPlan, result2 <-chan v7pushaction.Event, result3 <-chan v7pushaction.Warnings, result4 <-chan error) {
	fake.actualizeMutex.Lock()
	defer fake.actualizeMutex.Unlock()
	fake.ActualizeStub = nil
	if fake.actualizeReturnsOnCall == nil {
		fake.actualizeReturnsOnCall = make(map[int]struct {
			result1 <-chan v7pushaction.PushPlan
			result2 <-chan v7pushaction.Event
			result3 <-chan v7pushaction.Warnings
			result4 <-chan error
		})
	}
	fake.actualizeReturnsOnCall[i] = struct {
		result1 <-chan v7pushaction.PushPlan
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakePushActor) Conceptualize(arg1 []string, arg2 string, arg3 string, arg4 string, arg5 v7pushaction.FlagOverrides) ([]v7pushaction.PushPlan, v7pushaction.Warnings, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.conceptualizeMutex.Lock()
	ret, specificReturn := fake.conceptualizeReturnsOnCall[len(fake.conceptualizeArgsForCall)]
	fake.conceptualizeArgsForCall = append(fake.conceptualizeArgsForCall, struct {
		arg1 []string
		arg2 string
		arg3 string
		arg4 string
		arg5 v7pushaction.FlagOverrides
	}{arg1Copy, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Conceptualize", []interface{}{arg1Copy, arg2, arg3, arg4, arg5})
	fake.conceptualizeMutex.Unlock()
	if fake.ConceptualizeStub != nil {
		return fake.ConceptualizeStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.conceptualizeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePushActor) ConceptualizeCallCount() int {
	fake.conceptualizeMutex.RLock()
	defer fake.conceptualizeMutex.RUnlock()
	return len(fake.conceptualizeArgsForCall)
}

func (fake *FakePushActor) ConceptualizeCalls(stub func([]string, string, string, string, v7pushaction.FlagOverrides) ([]v7pushaction.PushPlan, v7pushaction.Warnings, error)) {
	fake.conceptualizeMutex.Lock()
	defer fake.conceptualizeMutex.Unlock()
	fake.ConceptualizeStub = stub
}

func (fake *FakePushActor) ConceptualizeArgsForCall(i int) ([]string, string, string, string, v7pushaction.FlagOverrides) {
	fake.conceptualizeMutex.RLock()
	defer fake.conceptualizeMutex.RUnlock()
	argsForCall := fake.conceptualizeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakePushActor) ConceptualizeReturns(result1 []v7pushaction.PushPlan, result2 v7pushaction.Warnings, result3 error) {
	fake.conceptualizeMutex.Lock()
	defer fake.conceptualizeMutex.Unlock()
	fake.ConceptualizeStub = nil
	fake.conceptualizeReturns = struct {
		result1 []v7pushaction.PushPlan
		result2 v7pushaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePushActor) ConceptualizeReturnsOnCall(i int, result1 []v7pushaction.PushPlan, result2 v7pushaction.Warnings, result3 error) {
	fake.conceptualizeMutex.Lock()
	defer fake.conceptualizeMutex.Unlock()
	fake.ConceptualizeStub = nil
	if fake.conceptualizeReturnsOnCall == nil {
		fake.conceptualizeReturnsOnCall = make(map[int]struct {
			result1 []v7pushaction.PushPlan
			result2 v7pushaction.Warnings
			result3 error
		})
	}
	fake.conceptualizeReturnsOnCall[i] = struct {
		result1 []v7pushaction.PushPlan
		result2 v7pushaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePushActor) CreatePushPlans(arg1 string, arg2 manifestparser.Parser, arg3 v7pushaction.FlagOverrides) ([]v7pushaction.PushPlan, error) {
	fake.createPushPlansMutex.Lock()
	ret, specificReturn := fake.createPushPlansReturnsOnCall[len(fake.createPushPlansArgsForCall)]
	fake.createPushPlansArgsForCall = append(fake.createPushPlansArgsForCall, struct {
		arg1 string
		arg2 manifestparser.Parser
		arg3 v7pushaction.FlagOverrides
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreatePushPlans", []interface{}{arg1, arg2, arg3})
	fake.createPushPlansMutex.Unlock()
	if fake.CreatePushPlansStub != nil {
		return fake.CreatePushPlansStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createPushPlansReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePushActor) CreatePushPlansCallCount() int {
	fake.createPushPlansMutex.RLock()
	defer fake.createPushPlansMutex.RUnlock()
	return len(fake.createPushPlansArgsForCall)
}

func (fake *FakePushActor) CreatePushPlansCalls(stub func(string, manifestparser.Parser, v7pushaction.FlagOverrides) ([]v7pushaction.PushPlan, error)) {
	fake.createPushPlansMutex.Lock()
	defer fake.createPushPlansMutex.Unlock()
	fake.CreatePushPlansStub = stub
}

func (fake *FakePushActor) CreatePushPlansArgsForCall(i int) (string, manifestparser.Parser, v7pushaction.FlagOverrides) {
	fake.createPushPlansMutex.RLock()
	defer fake.createPushPlansMutex.RUnlock()
	argsForCall := fake.createPushPlansArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePushActor) CreatePushPlansReturns(result1 []v7pushaction.PushPlan, result2 error) {
	fake.createPushPlansMutex.Lock()
	defer fake.createPushPlansMutex.Unlock()
	fake.CreatePushPlansStub = nil
	fake.createPushPlansReturns = struct {
		result1 []v7pushaction.PushPlan
		result2 error
	}{result1, result2}
}

func (fake *FakePushActor) CreatePushPlansReturnsOnCall(i int, result1 []v7pushaction.PushPlan, result2 error) {
	fake.createPushPlansMutex.Lock()
	defer fake.createPushPlansMutex.Unlock()
	fake.CreatePushPlansStub = nil
	if fake.createPushPlansReturnsOnCall == nil {
		fake.createPushPlansReturnsOnCall = make(map[int]struct {
			result1 []v7pushaction.PushPlan
			result2 error
		})
	}
	fake.createPushPlansReturnsOnCall[i] = struct {
		result1 []v7pushaction.PushPlan
		result2 error
	}{result1, result2}
}

func (fake *FakePushActor) PrepareSpace(arg1 string, arg2 string, arg3 *manifestparser.Parser, arg4 v7pushaction.FlagOverrides) (<-chan []string, <-chan v7pushaction.Event, <-chan v7pushaction.Warnings, <-chan error) {
	fake.prepareSpaceMutex.Lock()
	ret, specificReturn := fake.prepareSpaceReturnsOnCall[len(fake.prepareSpaceArgsForCall)]
	fake.prepareSpaceArgsForCall = append(fake.prepareSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *manifestparser.Parser
		arg4 v7pushaction.FlagOverrides
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("PrepareSpace", []interface{}{arg1, arg2, arg3, arg4})
	fake.prepareSpaceMutex.Unlock()
	if fake.PrepareSpaceStub != nil {
		return fake.PrepareSpaceStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.prepareSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakePushActor) PrepareSpaceCallCount() int {
	fake.prepareSpaceMutex.RLock()
	defer fake.prepareSpaceMutex.RUnlock()
	return len(fake.prepareSpaceArgsForCall)
}

func (fake *FakePushActor) PrepareSpaceCalls(stub func(string, string, *manifestparser.Parser, v7pushaction.FlagOverrides) (<-chan []string, <-chan v7pushaction.Event, <-chan v7pushaction.Warnings, <-chan error)) {
	fake.prepareSpaceMutex.Lock()
	defer fake.prepareSpaceMutex.Unlock()
	fake.PrepareSpaceStub = stub
}

func (fake *FakePushActor) PrepareSpaceArgsForCall(i int) (string, string, *manifestparser.Parser, v7pushaction.FlagOverrides) {
	fake.prepareSpaceMutex.RLock()
	defer fake.prepareSpaceMutex.RUnlock()
	argsForCall := fake.prepareSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePushActor) PrepareSpaceReturns(result1 <-chan []string, result2 <-chan v7pushaction.Event, result3 <-chan v7pushaction.Warnings, result4 <-chan error) {
	fake.prepareSpaceMutex.Lock()
	defer fake.prepareSpaceMutex.Unlock()
	fake.PrepareSpaceStub = nil
	fake.prepareSpaceReturns = struct {
		result1 <-chan []string
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakePushActor) PrepareSpaceReturnsOnCall(i int, result1 <-chan []string, result2 <-chan v7pushaction.Event, result3 <-chan v7pushaction.Warnings, result4 <-chan error) {
	fake.prepareSpaceMutex.Lock()
	defer fake.prepareSpaceMutex.Unlock()
	fake.PrepareSpaceStub = nil
	if fake.prepareSpaceReturnsOnCall == nil {
		fake.prepareSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan []string
			result2 <-chan v7pushaction.Event
			result3 <-chan v7pushaction.Warnings
			result4 <-chan error
		})
	}
	fake.prepareSpaceReturnsOnCall[i] = struct {
		result1 <-chan []string
		result2 <-chan v7pushaction.Event
		result3 <-chan v7pushaction.Warnings
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakePushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.actualizeMutex.RLock()
	defer fake.actualizeMutex.RUnlock()
	fake.conceptualizeMutex.RLock()
	defer fake.conceptualizeMutex.RUnlock()
	fake.createPushPlansMutex.RLock()
	defer fake.createPushPlansMutex.RUnlock()
	fake.prepareSpaceMutex.RLock()
	defer fake.prepareSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePushActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.PushActor = new(FakePushActor)
