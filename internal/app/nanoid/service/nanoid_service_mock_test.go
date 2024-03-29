// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"sync"
)

// Ensure, that NanoidRepositoryMock does implement NanoidRepository.
// If this is not the case, regenerate this file with moq.
var _ NanoidRepository = &NanoidRepositoryMock{}

// NanoidRepositoryMock is a mock implementation of NanoidRepository.
//
// 	func TestSomethingThatUsesNanoidRepository(t *testing.T) {
//
// 		// make and configure a mocked NanoidRepository
// 		mockedNanoidRepository := &NanoidRepositoryMock{
// 			GenerateUniqueIDFunc: func() string {
// 				panic("mock out the GenerateUniqueID method")
// 			},
// 		}
//
// 		// use mockedNanoidRepository in code that requires NanoidRepository
// 		// and then make assertions.
//
// 	}
type NanoidRepositoryMock struct {
	// GenerateUniqueIDFunc mocks the GenerateUniqueID method.
	GenerateUniqueIDFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// GenerateUniqueID holds details about calls to the GenerateUniqueID method.
		GenerateUniqueID []struct {
		}
	}
	lockGenerateUniqueID sync.RWMutex
}

// GenerateUniqueID calls GenerateUniqueIDFunc.
func (mock *NanoidRepositoryMock) GenerateUniqueID() string {
	if mock.GenerateUniqueIDFunc == nil {
		panic("NanoidRepositoryMock.GenerateUniqueIDFunc: method is nil but NanoidRepository.GenerateUniqueID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGenerateUniqueID.Lock()
	mock.calls.GenerateUniqueID = append(mock.calls.GenerateUniqueID, callInfo)
	mock.lockGenerateUniqueID.Unlock()
	return mock.GenerateUniqueIDFunc()
}

// GenerateUniqueIDCalls gets all the calls that were made to GenerateUniqueID.
// Check the length with:
//     len(mockedNanoidRepository.GenerateUniqueIDCalls())
func (mock *NanoidRepositoryMock) GenerateUniqueIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGenerateUniqueID.RLock()
	calls = mock.calls.GenerateUniqueID
	mock.lockGenerateUniqueID.RUnlock()
	return calls
}
