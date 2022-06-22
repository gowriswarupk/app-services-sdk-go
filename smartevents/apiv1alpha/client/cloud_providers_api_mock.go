// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package smarteventsclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that CloudProvidersApiMock does implement CloudProvidersApi.
// If this is not the case, regenerate this file with moq.
var _ CloudProvidersApi = &CloudProvidersApiMock{}

// CloudProvidersApiMock is a mock implementation of CloudProvidersApi.
//
// 	func TestSomethingThatUsesCloudProvidersApi(t *testing.T) {
//
// 		// make and configure a mocked CloudProvidersApi
// 		mockedCloudProvidersApi := &CloudProvidersApiMock{
// 			CloudProviderAPIGetCloudProviderFunc: func(ctx _context.Context, id string) ApiCloudProviderAPIGetCloudProviderRequest {
// 				panic("mock out the CloudProviderAPIGetCloudProvider method")
// 			},
// 			CloudProviderAPIGetCloudProviderExecuteFunc: func(r ApiCloudProviderAPIGetCloudProviderRequest) (CloudProviderListResponse, *_nethttp.Response, error) {
// 				panic("mock out the CloudProviderAPIGetCloudProviderExecute method")
// 			},
// 			CloudProviderAPIListCloudProviderRegionsFunc: func(ctx _context.Context, id string) ApiCloudProviderAPIListCloudProviderRegionsRequest {
// 				panic("mock out the CloudProviderAPIListCloudProviderRegions method")
// 			},
// 			CloudProviderAPIListCloudProviderRegionsExecuteFunc: func(r ApiCloudProviderAPIListCloudProviderRegionsRequest) (CloudRegionListResponse, *_nethttp.Response, error) {
// 				panic("mock out the CloudProviderAPIListCloudProviderRegionsExecute method")
// 			},
// 			CloudProviderAPIListCloudProvidersFunc: func(ctx _context.Context) ApiCloudProviderAPIListCloudProvidersRequest {
// 				panic("mock out the CloudProviderAPIListCloudProviders method")
// 			},
// 			CloudProviderAPIListCloudProvidersExecuteFunc: func(r ApiCloudProviderAPIListCloudProvidersRequest) (CloudProviderListResponse, *_nethttp.Response, error) {
// 				panic("mock out the CloudProviderAPIListCloudProvidersExecute method")
// 			},
// 		}
//
// 		// use mockedCloudProvidersApi in code that requires CloudProvidersApi
// 		// and then make assertions.
//
// 	}
type CloudProvidersApiMock struct {
	// CloudProviderAPIGetCloudProviderFunc mocks the CloudProviderAPIGetCloudProvider method.
	CloudProviderAPIGetCloudProviderFunc func(ctx _context.Context, id string) ApiCloudProviderAPIGetCloudProviderRequest

	// CloudProviderAPIGetCloudProviderExecuteFunc mocks the CloudProviderAPIGetCloudProviderExecute method.
	CloudProviderAPIGetCloudProviderExecuteFunc func(r ApiCloudProviderAPIGetCloudProviderRequest) (CloudProviderListResponse, *_nethttp.Response, error)

	// CloudProviderAPIListCloudProviderRegionsFunc mocks the CloudProviderAPIListCloudProviderRegions method.
	CloudProviderAPIListCloudProviderRegionsFunc func(ctx _context.Context, id string) ApiCloudProviderAPIListCloudProviderRegionsRequest

	// CloudProviderAPIListCloudProviderRegionsExecuteFunc mocks the CloudProviderAPIListCloudProviderRegionsExecute method.
	CloudProviderAPIListCloudProviderRegionsExecuteFunc func(r ApiCloudProviderAPIListCloudProviderRegionsRequest) (CloudRegionListResponse, *_nethttp.Response, error)

	// CloudProviderAPIListCloudProvidersFunc mocks the CloudProviderAPIListCloudProviders method.
	CloudProviderAPIListCloudProvidersFunc func(ctx _context.Context) ApiCloudProviderAPIListCloudProvidersRequest

	// CloudProviderAPIListCloudProvidersExecuteFunc mocks the CloudProviderAPIListCloudProvidersExecute method.
	CloudProviderAPIListCloudProvidersExecuteFunc func(r ApiCloudProviderAPIListCloudProvidersRequest) (CloudProviderListResponse, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// CloudProviderAPIGetCloudProvider holds details about calls to the CloudProviderAPIGetCloudProvider method.
		CloudProviderAPIGetCloudProvider []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// CloudProviderAPIGetCloudProviderExecute holds details about calls to the CloudProviderAPIGetCloudProviderExecute method.
		CloudProviderAPIGetCloudProviderExecute []struct {
			// R is the r argument value.
			R ApiCloudProviderAPIGetCloudProviderRequest
		}
		// CloudProviderAPIListCloudProviderRegions holds details about calls to the CloudProviderAPIListCloudProviderRegions method.
		CloudProviderAPIListCloudProviderRegions []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ID is the id argument value.
			ID string
		}
		// CloudProviderAPIListCloudProviderRegionsExecute holds details about calls to the CloudProviderAPIListCloudProviderRegionsExecute method.
		CloudProviderAPIListCloudProviderRegionsExecute []struct {
			// R is the r argument value.
			R ApiCloudProviderAPIListCloudProviderRegionsRequest
		}
		// CloudProviderAPIListCloudProviders holds details about calls to the CloudProviderAPIListCloudProviders method.
		CloudProviderAPIListCloudProviders []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// CloudProviderAPIListCloudProvidersExecute holds details about calls to the CloudProviderAPIListCloudProvidersExecute method.
		CloudProviderAPIListCloudProvidersExecute []struct {
			// R is the r argument value.
			R ApiCloudProviderAPIListCloudProvidersRequest
		}
	}
	lockCloudProviderAPIGetCloudProvider                sync.RWMutex
	lockCloudProviderAPIGetCloudProviderExecute         sync.RWMutex
	lockCloudProviderAPIListCloudProviderRegions        sync.RWMutex
	lockCloudProviderAPIListCloudProviderRegionsExecute sync.RWMutex
	lockCloudProviderAPIListCloudProviders              sync.RWMutex
	lockCloudProviderAPIListCloudProvidersExecute       sync.RWMutex
}

// CloudProviderAPIGetCloudProvider calls CloudProviderAPIGetCloudProviderFunc.
func (mock *CloudProvidersApiMock) CloudProviderAPIGetCloudProvider(ctx _context.Context, id string) ApiCloudProviderAPIGetCloudProviderRequest {
	if mock.CloudProviderAPIGetCloudProviderFunc == nil {
		panic("CloudProvidersApiMock.CloudProviderAPIGetCloudProviderFunc: method is nil but CloudProvidersApi.CloudProviderAPIGetCloudProvider was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockCloudProviderAPIGetCloudProvider.Lock()
	mock.calls.CloudProviderAPIGetCloudProvider = append(mock.calls.CloudProviderAPIGetCloudProvider, callInfo)
	mock.lockCloudProviderAPIGetCloudProvider.Unlock()
	return mock.CloudProviderAPIGetCloudProviderFunc(ctx, id)
}

// CloudProviderAPIGetCloudProviderCalls gets all the calls that were made to CloudProviderAPIGetCloudProvider.
// Check the length with:
//     len(mockedCloudProvidersApi.CloudProviderAPIGetCloudProviderCalls())
func (mock *CloudProvidersApiMock) CloudProviderAPIGetCloudProviderCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockCloudProviderAPIGetCloudProvider.RLock()
	calls = mock.calls.CloudProviderAPIGetCloudProvider
	mock.lockCloudProviderAPIGetCloudProvider.RUnlock()
	return calls
}

// CloudProviderAPIGetCloudProviderExecute calls CloudProviderAPIGetCloudProviderExecuteFunc.
func (mock *CloudProvidersApiMock) CloudProviderAPIGetCloudProviderExecute(r ApiCloudProviderAPIGetCloudProviderRequest) (CloudProviderListResponse, *_nethttp.Response, error) {
	if mock.CloudProviderAPIGetCloudProviderExecuteFunc == nil {
		panic("CloudProvidersApiMock.CloudProviderAPIGetCloudProviderExecuteFunc: method is nil but CloudProvidersApi.CloudProviderAPIGetCloudProviderExecute was just called")
	}
	callInfo := struct {
		R ApiCloudProviderAPIGetCloudProviderRequest
	}{
		R: r,
	}
	mock.lockCloudProviderAPIGetCloudProviderExecute.Lock()
	mock.calls.CloudProviderAPIGetCloudProviderExecute = append(mock.calls.CloudProviderAPIGetCloudProviderExecute, callInfo)
	mock.lockCloudProviderAPIGetCloudProviderExecute.Unlock()
	return mock.CloudProviderAPIGetCloudProviderExecuteFunc(r)
}

// CloudProviderAPIGetCloudProviderExecuteCalls gets all the calls that were made to CloudProviderAPIGetCloudProviderExecute.
// Check the length with:
//     len(mockedCloudProvidersApi.CloudProviderAPIGetCloudProviderExecuteCalls())
func (mock *CloudProvidersApiMock) CloudProviderAPIGetCloudProviderExecuteCalls() []struct {
	R ApiCloudProviderAPIGetCloudProviderRequest
} {
	var calls []struct {
		R ApiCloudProviderAPIGetCloudProviderRequest
	}
	mock.lockCloudProviderAPIGetCloudProviderExecute.RLock()
	calls = mock.calls.CloudProviderAPIGetCloudProviderExecute
	mock.lockCloudProviderAPIGetCloudProviderExecute.RUnlock()
	return calls
}

// CloudProviderAPIListCloudProviderRegions calls CloudProviderAPIListCloudProviderRegionsFunc.
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProviderRegions(ctx _context.Context, id string) ApiCloudProviderAPIListCloudProviderRegionsRequest {
	if mock.CloudProviderAPIListCloudProviderRegionsFunc == nil {
		panic("CloudProvidersApiMock.CloudProviderAPIListCloudProviderRegionsFunc: method is nil but CloudProvidersApi.CloudProviderAPIListCloudProviderRegions was just called")
	}
	callInfo := struct {
		Ctx _context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockCloudProviderAPIListCloudProviderRegions.Lock()
	mock.calls.CloudProviderAPIListCloudProviderRegions = append(mock.calls.CloudProviderAPIListCloudProviderRegions, callInfo)
	mock.lockCloudProviderAPIListCloudProviderRegions.Unlock()
	return mock.CloudProviderAPIListCloudProviderRegionsFunc(ctx, id)
}

// CloudProviderAPIListCloudProviderRegionsCalls gets all the calls that were made to CloudProviderAPIListCloudProviderRegions.
// Check the length with:
//     len(mockedCloudProvidersApi.CloudProviderAPIListCloudProviderRegionsCalls())
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProviderRegionsCalls() []struct {
	Ctx _context.Context
	ID  string
} {
	var calls []struct {
		Ctx _context.Context
		ID  string
	}
	mock.lockCloudProviderAPIListCloudProviderRegions.RLock()
	calls = mock.calls.CloudProviderAPIListCloudProviderRegions
	mock.lockCloudProviderAPIListCloudProviderRegions.RUnlock()
	return calls
}

// CloudProviderAPIListCloudProviderRegionsExecute calls CloudProviderAPIListCloudProviderRegionsExecuteFunc.
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProviderRegionsExecute(r ApiCloudProviderAPIListCloudProviderRegionsRequest) (CloudRegionListResponse, *_nethttp.Response, error) {
	if mock.CloudProviderAPIListCloudProviderRegionsExecuteFunc == nil {
		panic("CloudProvidersApiMock.CloudProviderAPIListCloudProviderRegionsExecuteFunc: method is nil but CloudProvidersApi.CloudProviderAPIListCloudProviderRegionsExecute was just called")
	}
	callInfo := struct {
		R ApiCloudProviderAPIListCloudProviderRegionsRequest
	}{
		R: r,
	}
	mock.lockCloudProviderAPIListCloudProviderRegionsExecute.Lock()
	mock.calls.CloudProviderAPIListCloudProviderRegionsExecute = append(mock.calls.CloudProviderAPIListCloudProviderRegionsExecute, callInfo)
	mock.lockCloudProviderAPIListCloudProviderRegionsExecute.Unlock()
	return mock.CloudProviderAPIListCloudProviderRegionsExecuteFunc(r)
}

// CloudProviderAPIListCloudProviderRegionsExecuteCalls gets all the calls that were made to CloudProviderAPIListCloudProviderRegionsExecute.
// Check the length with:
//     len(mockedCloudProvidersApi.CloudProviderAPIListCloudProviderRegionsExecuteCalls())
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProviderRegionsExecuteCalls() []struct {
	R ApiCloudProviderAPIListCloudProviderRegionsRequest
} {
	var calls []struct {
		R ApiCloudProviderAPIListCloudProviderRegionsRequest
	}
	mock.lockCloudProviderAPIListCloudProviderRegionsExecute.RLock()
	calls = mock.calls.CloudProviderAPIListCloudProviderRegionsExecute
	mock.lockCloudProviderAPIListCloudProviderRegionsExecute.RUnlock()
	return calls
}

// CloudProviderAPIListCloudProviders calls CloudProviderAPIListCloudProvidersFunc.
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProviders(ctx _context.Context) ApiCloudProviderAPIListCloudProvidersRequest {
	if mock.CloudProviderAPIListCloudProvidersFunc == nil {
		panic("CloudProvidersApiMock.CloudProviderAPIListCloudProvidersFunc: method is nil but CloudProvidersApi.CloudProviderAPIListCloudProviders was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCloudProviderAPIListCloudProviders.Lock()
	mock.calls.CloudProviderAPIListCloudProviders = append(mock.calls.CloudProviderAPIListCloudProviders, callInfo)
	mock.lockCloudProviderAPIListCloudProviders.Unlock()
	return mock.CloudProviderAPIListCloudProvidersFunc(ctx)
}

// CloudProviderAPIListCloudProvidersCalls gets all the calls that were made to CloudProviderAPIListCloudProviders.
// Check the length with:
//     len(mockedCloudProvidersApi.CloudProviderAPIListCloudProvidersCalls())
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProvidersCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockCloudProviderAPIListCloudProviders.RLock()
	calls = mock.calls.CloudProviderAPIListCloudProviders
	mock.lockCloudProviderAPIListCloudProviders.RUnlock()
	return calls
}

// CloudProviderAPIListCloudProvidersExecute calls CloudProviderAPIListCloudProvidersExecuteFunc.
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProvidersExecute(r ApiCloudProviderAPIListCloudProvidersRequest) (CloudProviderListResponse, *_nethttp.Response, error) {
	if mock.CloudProviderAPIListCloudProvidersExecuteFunc == nil {
		panic("CloudProvidersApiMock.CloudProviderAPIListCloudProvidersExecuteFunc: method is nil but CloudProvidersApi.CloudProviderAPIListCloudProvidersExecute was just called")
	}
	callInfo := struct {
		R ApiCloudProviderAPIListCloudProvidersRequest
	}{
		R: r,
	}
	mock.lockCloudProviderAPIListCloudProvidersExecute.Lock()
	mock.calls.CloudProviderAPIListCloudProvidersExecute = append(mock.calls.CloudProviderAPIListCloudProvidersExecute, callInfo)
	mock.lockCloudProviderAPIListCloudProvidersExecute.Unlock()
	return mock.CloudProviderAPIListCloudProvidersExecuteFunc(r)
}

// CloudProviderAPIListCloudProvidersExecuteCalls gets all the calls that were made to CloudProviderAPIListCloudProvidersExecute.
// Check the length with:
//     len(mockedCloudProvidersApi.CloudProviderAPIListCloudProvidersExecuteCalls())
func (mock *CloudProvidersApiMock) CloudProviderAPIListCloudProvidersExecuteCalls() []struct {
	R ApiCloudProviderAPIListCloudProvidersRequest
} {
	var calls []struct {
		R ApiCloudProviderAPIListCloudProvidersRequest
	}
	mock.lockCloudProviderAPIListCloudProvidersExecute.RLock()
	calls = mock.calls.CloudProviderAPIListCloudProvidersExecute
	mock.lockCloudProviderAPIListCloudProvidersExecute.RUnlock()
	return calls
}
