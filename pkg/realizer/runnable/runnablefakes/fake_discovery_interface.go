// Code generated by counterfeiter. DO NOT EDIT.
package runnablefakes

import (
	"sync"

	openapi_v2 "github.com/google/gnostic/openapiv2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/openapi"
	"k8s.io/client-go/rest"
)

type FakeDiscoveryInterface struct {
	OpenAPISchemaStub        func() (*openapi_v2.Document, error)
	openAPISchemaMutex       sync.RWMutex
	openAPISchemaArgsForCall []struct {
	}
	openAPISchemaReturns struct {
		result1 *openapi_v2.Document
		result2 error
	}
	openAPISchemaReturnsOnCall map[int]struct {
		result1 *openapi_v2.Document
		result2 error
	}
	OpenAPIV3Stub        func() openapi.Client
	openAPIV3Mutex       sync.RWMutex
	openAPIV3ArgsForCall []struct {
	}
	openAPIV3Returns struct {
		result1 openapi.Client
	}
	openAPIV3ReturnsOnCall map[int]struct {
		result1 openapi.Client
	}
	RESTClientStub        func() rest.Interface
	rESTClientMutex       sync.RWMutex
	rESTClientArgsForCall []struct {
	}
	rESTClientReturns struct {
		result1 rest.Interface
	}
	rESTClientReturnsOnCall map[int]struct {
		result1 rest.Interface
	}
	ServerGroupsStub        func() (*v1.APIGroupList, error)
	serverGroupsMutex       sync.RWMutex
	serverGroupsArgsForCall []struct {
	}
	serverGroupsReturns struct {
		result1 *v1.APIGroupList
		result2 error
	}
	serverGroupsReturnsOnCall map[int]struct {
		result1 *v1.APIGroupList
		result2 error
	}
	ServerGroupsAndResourcesStub        func() ([]*v1.APIGroup, []*v1.APIResourceList, error)
	serverGroupsAndResourcesMutex       sync.RWMutex
	serverGroupsAndResourcesArgsForCall []struct {
	}
	serverGroupsAndResourcesReturns struct {
		result1 []*v1.APIGroup
		result2 []*v1.APIResourceList
		result3 error
	}
	serverGroupsAndResourcesReturnsOnCall map[int]struct {
		result1 []*v1.APIGroup
		result2 []*v1.APIResourceList
		result3 error
	}
	ServerPreferredNamespacedResourcesStub        func() ([]*v1.APIResourceList, error)
	serverPreferredNamespacedResourcesMutex       sync.RWMutex
	serverPreferredNamespacedResourcesArgsForCall []struct {
	}
	serverPreferredNamespacedResourcesReturns struct {
		result1 []*v1.APIResourceList
		result2 error
	}
	serverPreferredNamespacedResourcesReturnsOnCall map[int]struct {
		result1 []*v1.APIResourceList
		result2 error
	}
	ServerPreferredResourcesStub        func() ([]*v1.APIResourceList, error)
	serverPreferredResourcesMutex       sync.RWMutex
	serverPreferredResourcesArgsForCall []struct {
	}
	serverPreferredResourcesReturns struct {
		result1 []*v1.APIResourceList
		result2 error
	}
	serverPreferredResourcesReturnsOnCall map[int]struct {
		result1 []*v1.APIResourceList
		result2 error
	}
	ServerResourcesForGroupVersionStub        func(string) (*v1.APIResourceList, error)
	serverResourcesForGroupVersionMutex       sync.RWMutex
	serverResourcesForGroupVersionArgsForCall []struct {
		arg1 string
	}
	serverResourcesForGroupVersionReturns struct {
		result1 *v1.APIResourceList
		result2 error
	}
	serverResourcesForGroupVersionReturnsOnCall map[int]struct {
		result1 *v1.APIResourceList
		result2 error
	}
	ServerVersionStub        func() (*version.Info, error)
	serverVersionMutex       sync.RWMutex
	serverVersionArgsForCall []struct {
	}
	serverVersionReturns struct {
		result1 *version.Info
		result2 error
	}
	serverVersionReturnsOnCall map[int]struct {
		result1 *version.Info
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDiscoveryInterface) OpenAPISchema() (*openapi_v2.Document, error) {
	fake.openAPISchemaMutex.Lock()
	ret, specificReturn := fake.openAPISchemaReturnsOnCall[len(fake.openAPISchemaArgsForCall)]
	fake.openAPISchemaArgsForCall = append(fake.openAPISchemaArgsForCall, struct {
	}{})
	stub := fake.OpenAPISchemaStub
	fakeReturns := fake.openAPISchemaReturns
	fake.recordInvocation("OpenAPISchema", []interface{}{})
	fake.openAPISchemaMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiscoveryInterface) OpenAPISchemaCallCount() int {
	fake.openAPISchemaMutex.RLock()
	defer fake.openAPISchemaMutex.RUnlock()
	return len(fake.openAPISchemaArgsForCall)
}

func (fake *FakeDiscoveryInterface) OpenAPISchemaCalls(stub func() (*openapi_v2.Document, error)) {
	fake.openAPISchemaMutex.Lock()
	defer fake.openAPISchemaMutex.Unlock()
	fake.OpenAPISchemaStub = stub
}

func (fake *FakeDiscoveryInterface) OpenAPISchemaReturns(result1 *openapi_v2.Document, result2 error) {
	fake.openAPISchemaMutex.Lock()
	defer fake.openAPISchemaMutex.Unlock()
	fake.OpenAPISchemaStub = nil
	fake.openAPISchemaReturns = struct {
		result1 *openapi_v2.Document
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) OpenAPISchemaReturnsOnCall(i int, result1 *openapi_v2.Document, result2 error) {
	fake.openAPISchemaMutex.Lock()
	defer fake.openAPISchemaMutex.Unlock()
	fake.OpenAPISchemaStub = nil
	if fake.openAPISchemaReturnsOnCall == nil {
		fake.openAPISchemaReturnsOnCall = make(map[int]struct {
			result1 *openapi_v2.Document
			result2 error
		})
	}
	fake.openAPISchemaReturnsOnCall[i] = struct {
		result1 *openapi_v2.Document
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) OpenAPIV3() openapi.Client {
	fake.openAPIV3Mutex.Lock()
	ret, specificReturn := fake.openAPIV3ReturnsOnCall[len(fake.openAPIV3ArgsForCall)]
	fake.openAPIV3ArgsForCall = append(fake.openAPIV3ArgsForCall, struct {
	}{})
	stub := fake.OpenAPIV3Stub
	fakeReturns := fake.openAPIV3Returns
	fake.recordInvocation("OpenAPIV3", []interface{}{})
	fake.openAPIV3Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDiscoveryInterface) OpenAPIV3CallCount() int {
	fake.openAPIV3Mutex.RLock()
	defer fake.openAPIV3Mutex.RUnlock()
	return len(fake.openAPIV3ArgsForCall)
}

func (fake *FakeDiscoveryInterface) OpenAPIV3Calls(stub func() openapi.Client) {
	fake.openAPIV3Mutex.Lock()
	defer fake.openAPIV3Mutex.Unlock()
	fake.OpenAPIV3Stub = stub
}

func (fake *FakeDiscoveryInterface) OpenAPIV3Returns(result1 openapi.Client) {
	fake.openAPIV3Mutex.Lock()
	defer fake.openAPIV3Mutex.Unlock()
	fake.OpenAPIV3Stub = nil
	fake.openAPIV3Returns = struct {
		result1 openapi.Client
	}{result1}
}

func (fake *FakeDiscoveryInterface) OpenAPIV3ReturnsOnCall(i int, result1 openapi.Client) {
	fake.openAPIV3Mutex.Lock()
	defer fake.openAPIV3Mutex.Unlock()
	fake.OpenAPIV3Stub = nil
	if fake.openAPIV3ReturnsOnCall == nil {
		fake.openAPIV3ReturnsOnCall = make(map[int]struct {
			result1 openapi.Client
		})
	}
	fake.openAPIV3ReturnsOnCall[i] = struct {
		result1 openapi.Client
	}{result1}
}

func (fake *FakeDiscoveryInterface) RESTClient() rest.Interface {
	fake.rESTClientMutex.Lock()
	ret, specificReturn := fake.rESTClientReturnsOnCall[len(fake.rESTClientArgsForCall)]
	fake.rESTClientArgsForCall = append(fake.rESTClientArgsForCall, struct {
	}{})
	stub := fake.RESTClientStub
	fakeReturns := fake.rESTClientReturns
	fake.recordInvocation("RESTClient", []interface{}{})
	fake.rESTClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDiscoveryInterface) RESTClientCallCount() int {
	fake.rESTClientMutex.RLock()
	defer fake.rESTClientMutex.RUnlock()
	return len(fake.rESTClientArgsForCall)
}

func (fake *FakeDiscoveryInterface) RESTClientCalls(stub func() rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = stub
}

func (fake *FakeDiscoveryInterface) RESTClientReturns(result1 rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = nil
	fake.rESTClientReturns = struct {
		result1 rest.Interface
	}{result1}
}

func (fake *FakeDiscoveryInterface) RESTClientReturnsOnCall(i int, result1 rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = nil
	if fake.rESTClientReturnsOnCall == nil {
		fake.rESTClientReturnsOnCall = make(map[int]struct {
			result1 rest.Interface
		})
	}
	fake.rESTClientReturnsOnCall[i] = struct {
		result1 rest.Interface
	}{result1}
}

func (fake *FakeDiscoveryInterface) ServerGroups() (*v1.APIGroupList, error) {
	fake.serverGroupsMutex.Lock()
	ret, specificReturn := fake.serverGroupsReturnsOnCall[len(fake.serverGroupsArgsForCall)]
	fake.serverGroupsArgsForCall = append(fake.serverGroupsArgsForCall, struct {
	}{})
	stub := fake.ServerGroupsStub
	fakeReturns := fake.serverGroupsReturns
	fake.recordInvocation("ServerGroups", []interface{}{})
	fake.serverGroupsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiscoveryInterface) ServerGroupsCallCount() int {
	fake.serverGroupsMutex.RLock()
	defer fake.serverGroupsMutex.RUnlock()
	return len(fake.serverGroupsArgsForCall)
}

func (fake *FakeDiscoveryInterface) ServerGroupsCalls(stub func() (*v1.APIGroupList, error)) {
	fake.serverGroupsMutex.Lock()
	defer fake.serverGroupsMutex.Unlock()
	fake.ServerGroupsStub = stub
}

func (fake *FakeDiscoveryInterface) ServerGroupsReturns(result1 *v1.APIGroupList, result2 error) {
	fake.serverGroupsMutex.Lock()
	defer fake.serverGroupsMutex.Unlock()
	fake.ServerGroupsStub = nil
	fake.serverGroupsReturns = struct {
		result1 *v1.APIGroupList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerGroupsReturnsOnCall(i int, result1 *v1.APIGroupList, result2 error) {
	fake.serverGroupsMutex.Lock()
	defer fake.serverGroupsMutex.Unlock()
	fake.ServerGroupsStub = nil
	if fake.serverGroupsReturnsOnCall == nil {
		fake.serverGroupsReturnsOnCall = make(map[int]struct {
			result1 *v1.APIGroupList
			result2 error
		})
	}
	fake.serverGroupsReturnsOnCall[i] = struct {
		result1 *v1.APIGroupList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerGroupsAndResources() ([]*v1.APIGroup, []*v1.APIResourceList, error) {
	fake.serverGroupsAndResourcesMutex.Lock()
	ret, specificReturn := fake.serverGroupsAndResourcesReturnsOnCall[len(fake.serverGroupsAndResourcesArgsForCall)]
	fake.serverGroupsAndResourcesArgsForCall = append(fake.serverGroupsAndResourcesArgsForCall, struct {
	}{})
	stub := fake.ServerGroupsAndResourcesStub
	fakeReturns := fake.serverGroupsAndResourcesReturns
	fake.recordInvocation("ServerGroupsAndResources", []interface{}{})
	fake.serverGroupsAndResourcesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeDiscoveryInterface) ServerGroupsAndResourcesCallCount() int {
	fake.serverGroupsAndResourcesMutex.RLock()
	defer fake.serverGroupsAndResourcesMutex.RUnlock()
	return len(fake.serverGroupsAndResourcesArgsForCall)
}

func (fake *FakeDiscoveryInterface) ServerGroupsAndResourcesCalls(stub func() ([]*v1.APIGroup, []*v1.APIResourceList, error)) {
	fake.serverGroupsAndResourcesMutex.Lock()
	defer fake.serverGroupsAndResourcesMutex.Unlock()
	fake.ServerGroupsAndResourcesStub = stub
}

func (fake *FakeDiscoveryInterface) ServerGroupsAndResourcesReturns(result1 []*v1.APIGroup, result2 []*v1.APIResourceList, result3 error) {
	fake.serverGroupsAndResourcesMutex.Lock()
	defer fake.serverGroupsAndResourcesMutex.Unlock()
	fake.ServerGroupsAndResourcesStub = nil
	fake.serverGroupsAndResourcesReturns = struct {
		result1 []*v1.APIGroup
		result2 []*v1.APIResourceList
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDiscoveryInterface) ServerGroupsAndResourcesReturnsOnCall(i int, result1 []*v1.APIGroup, result2 []*v1.APIResourceList, result3 error) {
	fake.serverGroupsAndResourcesMutex.Lock()
	defer fake.serverGroupsAndResourcesMutex.Unlock()
	fake.ServerGroupsAndResourcesStub = nil
	if fake.serverGroupsAndResourcesReturnsOnCall == nil {
		fake.serverGroupsAndResourcesReturnsOnCall = make(map[int]struct {
			result1 []*v1.APIGroup
			result2 []*v1.APIResourceList
			result3 error
		})
	}
	fake.serverGroupsAndResourcesReturnsOnCall[i] = struct {
		result1 []*v1.APIGroup
		result2 []*v1.APIResourceList
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDiscoveryInterface) ServerPreferredNamespacedResources() ([]*v1.APIResourceList, error) {
	fake.serverPreferredNamespacedResourcesMutex.Lock()
	ret, specificReturn := fake.serverPreferredNamespacedResourcesReturnsOnCall[len(fake.serverPreferredNamespacedResourcesArgsForCall)]
	fake.serverPreferredNamespacedResourcesArgsForCall = append(fake.serverPreferredNamespacedResourcesArgsForCall, struct {
	}{})
	stub := fake.ServerPreferredNamespacedResourcesStub
	fakeReturns := fake.serverPreferredNamespacedResourcesReturns
	fake.recordInvocation("ServerPreferredNamespacedResources", []interface{}{})
	fake.serverPreferredNamespacedResourcesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiscoveryInterface) ServerPreferredNamespacedResourcesCallCount() int {
	fake.serverPreferredNamespacedResourcesMutex.RLock()
	defer fake.serverPreferredNamespacedResourcesMutex.RUnlock()
	return len(fake.serverPreferredNamespacedResourcesArgsForCall)
}

func (fake *FakeDiscoveryInterface) ServerPreferredNamespacedResourcesCalls(stub func() ([]*v1.APIResourceList, error)) {
	fake.serverPreferredNamespacedResourcesMutex.Lock()
	defer fake.serverPreferredNamespacedResourcesMutex.Unlock()
	fake.ServerPreferredNamespacedResourcesStub = stub
}

func (fake *FakeDiscoveryInterface) ServerPreferredNamespacedResourcesReturns(result1 []*v1.APIResourceList, result2 error) {
	fake.serverPreferredNamespacedResourcesMutex.Lock()
	defer fake.serverPreferredNamespacedResourcesMutex.Unlock()
	fake.ServerPreferredNamespacedResourcesStub = nil
	fake.serverPreferredNamespacedResourcesReturns = struct {
		result1 []*v1.APIResourceList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerPreferredNamespacedResourcesReturnsOnCall(i int, result1 []*v1.APIResourceList, result2 error) {
	fake.serverPreferredNamespacedResourcesMutex.Lock()
	defer fake.serverPreferredNamespacedResourcesMutex.Unlock()
	fake.ServerPreferredNamespacedResourcesStub = nil
	if fake.serverPreferredNamespacedResourcesReturnsOnCall == nil {
		fake.serverPreferredNamespacedResourcesReturnsOnCall = make(map[int]struct {
			result1 []*v1.APIResourceList
			result2 error
		})
	}
	fake.serverPreferredNamespacedResourcesReturnsOnCall[i] = struct {
		result1 []*v1.APIResourceList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	fake.serverPreferredResourcesMutex.Lock()
	ret, specificReturn := fake.serverPreferredResourcesReturnsOnCall[len(fake.serverPreferredResourcesArgsForCall)]
	fake.serverPreferredResourcesArgsForCall = append(fake.serverPreferredResourcesArgsForCall, struct {
	}{})
	stub := fake.ServerPreferredResourcesStub
	fakeReturns := fake.serverPreferredResourcesReturns
	fake.recordInvocation("ServerPreferredResources", []interface{}{})
	fake.serverPreferredResourcesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiscoveryInterface) ServerPreferredResourcesCallCount() int {
	fake.serverPreferredResourcesMutex.RLock()
	defer fake.serverPreferredResourcesMutex.RUnlock()
	return len(fake.serverPreferredResourcesArgsForCall)
}

func (fake *FakeDiscoveryInterface) ServerPreferredResourcesCalls(stub func() ([]*v1.APIResourceList, error)) {
	fake.serverPreferredResourcesMutex.Lock()
	defer fake.serverPreferredResourcesMutex.Unlock()
	fake.ServerPreferredResourcesStub = stub
}

func (fake *FakeDiscoveryInterface) ServerPreferredResourcesReturns(result1 []*v1.APIResourceList, result2 error) {
	fake.serverPreferredResourcesMutex.Lock()
	defer fake.serverPreferredResourcesMutex.Unlock()
	fake.ServerPreferredResourcesStub = nil
	fake.serverPreferredResourcesReturns = struct {
		result1 []*v1.APIResourceList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerPreferredResourcesReturnsOnCall(i int, result1 []*v1.APIResourceList, result2 error) {
	fake.serverPreferredResourcesMutex.Lock()
	defer fake.serverPreferredResourcesMutex.Unlock()
	fake.ServerPreferredResourcesStub = nil
	if fake.serverPreferredResourcesReturnsOnCall == nil {
		fake.serverPreferredResourcesReturnsOnCall = make(map[int]struct {
			result1 []*v1.APIResourceList
			result2 error
		})
	}
	fake.serverPreferredResourcesReturnsOnCall[i] = struct {
		result1 []*v1.APIResourceList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerResourcesForGroupVersion(arg1 string) (*v1.APIResourceList, error) {
	fake.serverResourcesForGroupVersionMutex.Lock()
	ret, specificReturn := fake.serverResourcesForGroupVersionReturnsOnCall[len(fake.serverResourcesForGroupVersionArgsForCall)]
	fake.serverResourcesForGroupVersionArgsForCall = append(fake.serverResourcesForGroupVersionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ServerResourcesForGroupVersionStub
	fakeReturns := fake.serverResourcesForGroupVersionReturns
	fake.recordInvocation("ServerResourcesForGroupVersion", []interface{}{arg1})
	fake.serverResourcesForGroupVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiscoveryInterface) ServerResourcesForGroupVersionCallCount() int {
	fake.serverResourcesForGroupVersionMutex.RLock()
	defer fake.serverResourcesForGroupVersionMutex.RUnlock()
	return len(fake.serverResourcesForGroupVersionArgsForCall)
}

func (fake *FakeDiscoveryInterface) ServerResourcesForGroupVersionCalls(stub func(string) (*v1.APIResourceList, error)) {
	fake.serverResourcesForGroupVersionMutex.Lock()
	defer fake.serverResourcesForGroupVersionMutex.Unlock()
	fake.ServerResourcesForGroupVersionStub = stub
}

func (fake *FakeDiscoveryInterface) ServerResourcesForGroupVersionArgsForCall(i int) string {
	fake.serverResourcesForGroupVersionMutex.RLock()
	defer fake.serverResourcesForGroupVersionMutex.RUnlock()
	argsForCall := fake.serverResourcesForGroupVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDiscoveryInterface) ServerResourcesForGroupVersionReturns(result1 *v1.APIResourceList, result2 error) {
	fake.serverResourcesForGroupVersionMutex.Lock()
	defer fake.serverResourcesForGroupVersionMutex.Unlock()
	fake.ServerResourcesForGroupVersionStub = nil
	fake.serverResourcesForGroupVersionReturns = struct {
		result1 *v1.APIResourceList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerResourcesForGroupVersionReturnsOnCall(i int, result1 *v1.APIResourceList, result2 error) {
	fake.serverResourcesForGroupVersionMutex.Lock()
	defer fake.serverResourcesForGroupVersionMutex.Unlock()
	fake.ServerResourcesForGroupVersionStub = nil
	if fake.serverResourcesForGroupVersionReturnsOnCall == nil {
		fake.serverResourcesForGroupVersionReturnsOnCall = make(map[int]struct {
			result1 *v1.APIResourceList
			result2 error
		})
	}
	fake.serverResourcesForGroupVersionReturnsOnCall[i] = struct {
		result1 *v1.APIResourceList
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerVersion() (*version.Info, error) {
	fake.serverVersionMutex.Lock()
	ret, specificReturn := fake.serverVersionReturnsOnCall[len(fake.serverVersionArgsForCall)]
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct {
	}{})
	stub := fake.ServerVersionStub
	fakeReturns := fake.serverVersionReturns
	fake.recordInvocation("ServerVersion", []interface{}{})
	fake.serverVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDiscoveryInterface) ServerVersionCallCount() int {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}

func (fake *FakeDiscoveryInterface) ServerVersionCalls(stub func() (*version.Info, error)) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = stub
}

func (fake *FakeDiscoveryInterface) ServerVersionReturns(result1 *version.Info, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1 *version.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) ServerVersionReturnsOnCall(i int, result1 *version.Info, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	if fake.serverVersionReturnsOnCall == nil {
		fake.serverVersionReturnsOnCall = make(map[int]struct {
			result1 *version.Info
			result2 error
		})
	}
	fake.serverVersionReturnsOnCall[i] = struct {
		result1 *version.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeDiscoveryInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openAPISchemaMutex.RLock()
	defer fake.openAPISchemaMutex.RUnlock()
	fake.openAPIV3Mutex.RLock()
	defer fake.openAPIV3Mutex.RUnlock()
	fake.rESTClientMutex.RLock()
	defer fake.rESTClientMutex.RUnlock()
	fake.serverGroupsMutex.RLock()
	defer fake.serverGroupsMutex.RUnlock()
	fake.serverGroupsAndResourcesMutex.RLock()
	defer fake.serverGroupsAndResourcesMutex.RUnlock()
	fake.serverPreferredNamespacedResourcesMutex.RLock()
	defer fake.serverPreferredNamespacedResourcesMutex.RUnlock()
	fake.serverPreferredResourcesMutex.RLock()
	defer fake.serverPreferredResourcesMutex.RUnlock()
	fake.serverResourcesForGroupVersionMutex.RLock()
	defer fake.serverResourcesForGroupVersionMutex.RUnlock()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDiscoveryInterface) recordInvocation(key string, args []interface{}) {
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

var _ discovery.DiscoveryInterface = new(FakeDiscoveryInterface)