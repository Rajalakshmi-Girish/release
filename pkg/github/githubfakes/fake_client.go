/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package githubfakes

import (
	"context"
	"sync"

	githuba "github.com/google/go-github/v29/github"
	"k8s.io/release/pkg/github"
)

type FakeClient struct {
	GetCommitStub        func(context.Context, string, string, string) (*githuba.Commit, *githuba.Response, error)
	getCommitMutex       sync.RWMutex
	getCommitArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getCommitReturns struct {
		result1 *githuba.Commit
		result2 *githuba.Response
		result3 error
	}
	getCommitReturnsOnCall map[int]struct {
		result1 *githuba.Commit
		result2 *githuba.Response
		result3 error
	}
	GetPullRequestStub        func(context.Context, string, string, int) (*githuba.PullRequest, *githuba.Response, error)
	getPullRequestMutex       sync.RWMutex
	getPullRequestArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 int
	}
	getPullRequestReturns struct {
		result1 *githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}
	getPullRequestReturnsOnCall map[int]struct {
		result1 *githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}
	GetRepoCommitStub        func(context.Context, string, string, string) (*githuba.RepositoryCommit, *githuba.Response, error)
	getRepoCommitMutex       sync.RWMutex
	getRepoCommitArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getRepoCommitReturns struct {
		result1 *githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}
	getRepoCommitReturnsOnCall map[int]struct {
		result1 *githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}
	ListCommitsStub        func(context.Context, string, string, *githuba.CommitsListOptions) ([]*githuba.RepositoryCommit, *githuba.Response, error)
	listCommitsMutex       sync.RWMutex
	listCommitsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *githuba.CommitsListOptions
	}
	listCommitsReturns struct {
		result1 []*githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}
	listCommitsReturnsOnCall map[int]struct {
		result1 []*githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}
	ListPullRequestsWithCommitStub        func(context.Context, string, string, string, *githuba.PullRequestListOptions) ([]*githuba.PullRequest, *githuba.Response, error)
	listPullRequestsWithCommitMutex       sync.RWMutex
	listPullRequestsWithCommitArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 *githuba.PullRequestListOptions
	}
	listPullRequestsWithCommitReturns struct {
		result1 []*githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}
	listPullRequestsWithCommitReturnsOnCall map[int]struct {
		result1 []*githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}
	ListReleasesStub        func(context.Context, string, string, *githuba.ListOptions) ([]*githuba.RepositoryRelease, *githuba.Response, error)
	listReleasesMutex       sync.RWMutex
	listReleasesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *githuba.ListOptions
	}
	listReleasesReturns struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}
	listReleasesReturnsOnCall map[int]struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}
	ListTagsStub        func(context.Context, string, string, *githuba.ListOptions) ([]*githuba.RepositoryTag, *githuba.Response, error)
	listTagsMutex       sync.RWMutex
	listTagsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *githuba.ListOptions
	}
	listTagsReturns struct {
		result1 []*githuba.RepositoryTag
		result2 *githuba.Response
		result3 error
	}
	listTagsReturnsOnCall map[int]struct {
		result1 []*githuba.RepositoryTag
		result2 *githuba.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) GetCommit(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*githuba.Commit, *githuba.Response, error) {
	fake.getCommitMutex.Lock()
	ret, specificReturn := fake.getCommitReturnsOnCall[len(fake.getCommitArgsForCall)]
	fake.getCommitArgsForCall = append(fake.getCommitArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetCommit", []interface{}{arg1, arg2, arg3, arg4})
	fake.getCommitMutex.Unlock()
	if fake.GetCommitStub != nil {
		return fake.GetCommitStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getCommitReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) GetCommitCallCount() int {
	fake.getCommitMutex.RLock()
	defer fake.getCommitMutex.RUnlock()
	return len(fake.getCommitArgsForCall)
}

func (fake *FakeClient) GetCommitCalls(stub func(context.Context, string, string, string) (*githuba.Commit, *githuba.Response, error)) {
	fake.getCommitMutex.Lock()
	defer fake.getCommitMutex.Unlock()
	fake.GetCommitStub = stub
}

func (fake *FakeClient) GetCommitArgsForCall(i int) (context.Context, string, string, string) {
	fake.getCommitMutex.RLock()
	defer fake.getCommitMutex.RUnlock()
	argsForCall := fake.getCommitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) GetCommitReturns(result1 *githuba.Commit, result2 *githuba.Response, result3 error) {
	fake.getCommitMutex.Lock()
	defer fake.getCommitMutex.Unlock()
	fake.GetCommitStub = nil
	fake.getCommitReturns = struct {
		result1 *githuba.Commit
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) GetCommitReturnsOnCall(i int, result1 *githuba.Commit, result2 *githuba.Response, result3 error) {
	fake.getCommitMutex.Lock()
	defer fake.getCommitMutex.Unlock()
	fake.GetCommitStub = nil
	if fake.getCommitReturnsOnCall == nil {
		fake.getCommitReturnsOnCall = make(map[int]struct {
			result1 *githuba.Commit
			result2 *githuba.Response
			result3 error
		})
	}
	fake.getCommitReturnsOnCall[i] = struct {
		result1 *githuba.Commit
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) GetPullRequest(arg1 context.Context, arg2 string, arg3 string, arg4 int) (*githuba.PullRequest, *githuba.Response, error) {
	fake.getPullRequestMutex.Lock()
	ret, specificReturn := fake.getPullRequestReturnsOnCall[len(fake.getPullRequestArgsForCall)]
	fake.getPullRequestArgsForCall = append(fake.getPullRequestArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 int
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetPullRequest", []interface{}{arg1, arg2, arg3, arg4})
	fake.getPullRequestMutex.Unlock()
	if fake.GetPullRequestStub != nil {
		return fake.GetPullRequestStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getPullRequestReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) GetPullRequestCallCount() int {
	fake.getPullRequestMutex.RLock()
	defer fake.getPullRequestMutex.RUnlock()
	return len(fake.getPullRequestArgsForCall)
}

func (fake *FakeClient) GetPullRequestCalls(stub func(context.Context, string, string, int) (*githuba.PullRequest, *githuba.Response, error)) {
	fake.getPullRequestMutex.Lock()
	defer fake.getPullRequestMutex.Unlock()
	fake.GetPullRequestStub = stub
}

func (fake *FakeClient) GetPullRequestArgsForCall(i int) (context.Context, string, string, int) {
	fake.getPullRequestMutex.RLock()
	defer fake.getPullRequestMutex.RUnlock()
	argsForCall := fake.getPullRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) GetPullRequestReturns(result1 *githuba.PullRequest, result2 *githuba.Response, result3 error) {
	fake.getPullRequestMutex.Lock()
	defer fake.getPullRequestMutex.Unlock()
	fake.GetPullRequestStub = nil
	fake.getPullRequestReturns = struct {
		result1 *githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) GetPullRequestReturnsOnCall(i int, result1 *githuba.PullRequest, result2 *githuba.Response, result3 error) {
	fake.getPullRequestMutex.Lock()
	defer fake.getPullRequestMutex.Unlock()
	fake.GetPullRequestStub = nil
	if fake.getPullRequestReturnsOnCall == nil {
		fake.getPullRequestReturnsOnCall = make(map[int]struct {
			result1 *githuba.PullRequest
			result2 *githuba.Response
			result3 error
		})
	}
	fake.getPullRequestReturnsOnCall[i] = struct {
		result1 *githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) GetRepoCommit(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*githuba.RepositoryCommit, *githuba.Response, error) {
	fake.getRepoCommitMutex.Lock()
	ret, specificReturn := fake.getRepoCommitReturnsOnCall[len(fake.getRepoCommitArgsForCall)]
	fake.getRepoCommitArgsForCall = append(fake.getRepoCommitArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetRepoCommit", []interface{}{arg1, arg2, arg3, arg4})
	fake.getRepoCommitMutex.Unlock()
	if fake.GetRepoCommitStub != nil {
		return fake.GetRepoCommitStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getRepoCommitReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) GetRepoCommitCallCount() int {
	fake.getRepoCommitMutex.RLock()
	defer fake.getRepoCommitMutex.RUnlock()
	return len(fake.getRepoCommitArgsForCall)
}

func (fake *FakeClient) GetRepoCommitCalls(stub func(context.Context, string, string, string) (*githuba.RepositoryCommit, *githuba.Response, error)) {
	fake.getRepoCommitMutex.Lock()
	defer fake.getRepoCommitMutex.Unlock()
	fake.GetRepoCommitStub = stub
}

func (fake *FakeClient) GetRepoCommitArgsForCall(i int) (context.Context, string, string, string) {
	fake.getRepoCommitMutex.RLock()
	defer fake.getRepoCommitMutex.RUnlock()
	argsForCall := fake.getRepoCommitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) GetRepoCommitReturns(result1 *githuba.RepositoryCommit, result2 *githuba.Response, result3 error) {
	fake.getRepoCommitMutex.Lock()
	defer fake.getRepoCommitMutex.Unlock()
	fake.GetRepoCommitStub = nil
	fake.getRepoCommitReturns = struct {
		result1 *githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) GetRepoCommitReturnsOnCall(i int, result1 *githuba.RepositoryCommit, result2 *githuba.Response, result3 error) {
	fake.getRepoCommitMutex.Lock()
	defer fake.getRepoCommitMutex.Unlock()
	fake.GetRepoCommitStub = nil
	if fake.getRepoCommitReturnsOnCall == nil {
		fake.getRepoCommitReturnsOnCall = make(map[int]struct {
			result1 *githuba.RepositoryCommit
			result2 *githuba.Response
			result3 error
		})
	}
	fake.getRepoCommitReturnsOnCall[i] = struct {
		result1 *githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListCommits(arg1 context.Context, arg2 string, arg3 string, arg4 *githuba.CommitsListOptions) ([]*githuba.RepositoryCommit, *githuba.Response, error) {
	fake.listCommitsMutex.Lock()
	ret, specificReturn := fake.listCommitsReturnsOnCall[len(fake.listCommitsArgsForCall)]
	fake.listCommitsArgsForCall = append(fake.listCommitsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *githuba.CommitsListOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ListCommits", []interface{}{arg1, arg2, arg3, arg4})
	fake.listCommitsMutex.Unlock()
	if fake.ListCommitsStub != nil {
		return fake.ListCommitsStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listCommitsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListCommitsCallCount() int {
	fake.listCommitsMutex.RLock()
	defer fake.listCommitsMutex.RUnlock()
	return len(fake.listCommitsArgsForCall)
}

func (fake *FakeClient) ListCommitsCalls(stub func(context.Context, string, string, *githuba.CommitsListOptions) ([]*githuba.RepositoryCommit, *githuba.Response, error)) {
	fake.listCommitsMutex.Lock()
	defer fake.listCommitsMutex.Unlock()
	fake.ListCommitsStub = stub
}

func (fake *FakeClient) ListCommitsArgsForCall(i int) (context.Context, string, string, *githuba.CommitsListOptions) {
	fake.listCommitsMutex.RLock()
	defer fake.listCommitsMutex.RUnlock()
	argsForCall := fake.listCommitsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) ListCommitsReturns(result1 []*githuba.RepositoryCommit, result2 *githuba.Response, result3 error) {
	fake.listCommitsMutex.Lock()
	defer fake.listCommitsMutex.Unlock()
	fake.ListCommitsStub = nil
	fake.listCommitsReturns = struct {
		result1 []*githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListCommitsReturnsOnCall(i int, result1 []*githuba.RepositoryCommit, result2 *githuba.Response, result3 error) {
	fake.listCommitsMutex.Lock()
	defer fake.listCommitsMutex.Unlock()
	fake.ListCommitsStub = nil
	if fake.listCommitsReturnsOnCall == nil {
		fake.listCommitsReturnsOnCall = make(map[int]struct {
			result1 []*githuba.RepositoryCommit
			result2 *githuba.Response
			result3 error
		})
	}
	fake.listCommitsReturnsOnCall[i] = struct {
		result1 []*githuba.RepositoryCommit
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListPullRequestsWithCommit(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 *githuba.PullRequestListOptions) ([]*githuba.PullRequest, *githuba.Response, error) {
	fake.listPullRequestsWithCommitMutex.Lock()
	ret, specificReturn := fake.listPullRequestsWithCommitReturnsOnCall[len(fake.listPullRequestsWithCommitArgsForCall)]
	fake.listPullRequestsWithCommitArgsForCall = append(fake.listPullRequestsWithCommitArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 *githuba.PullRequestListOptions
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("ListPullRequestsWithCommit", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.listPullRequestsWithCommitMutex.Unlock()
	if fake.ListPullRequestsWithCommitStub != nil {
		return fake.ListPullRequestsWithCommitStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listPullRequestsWithCommitReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListPullRequestsWithCommitCallCount() int {
	fake.listPullRequestsWithCommitMutex.RLock()
	defer fake.listPullRequestsWithCommitMutex.RUnlock()
	return len(fake.listPullRequestsWithCommitArgsForCall)
}

func (fake *FakeClient) ListPullRequestsWithCommitCalls(stub func(context.Context, string, string, string, *githuba.PullRequestListOptions) ([]*githuba.PullRequest, *githuba.Response, error)) {
	fake.listPullRequestsWithCommitMutex.Lock()
	defer fake.listPullRequestsWithCommitMutex.Unlock()
	fake.ListPullRequestsWithCommitStub = stub
}

func (fake *FakeClient) ListPullRequestsWithCommitArgsForCall(i int) (context.Context, string, string, string, *githuba.PullRequestListOptions) {
	fake.listPullRequestsWithCommitMutex.RLock()
	defer fake.listPullRequestsWithCommitMutex.RUnlock()
	argsForCall := fake.listPullRequestsWithCommitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeClient) ListPullRequestsWithCommitReturns(result1 []*githuba.PullRequest, result2 *githuba.Response, result3 error) {
	fake.listPullRequestsWithCommitMutex.Lock()
	defer fake.listPullRequestsWithCommitMutex.Unlock()
	fake.ListPullRequestsWithCommitStub = nil
	fake.listPullRequestsWithCommitReturns = struct {
		result1 []*githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListPullRequestsWithCommitReturnsOnCall(i int, result1 []*githuba.PullRequest, result2 *githuba.Response, result3 error) {
	fake.listPullRequestsWithCommitMutex.Lock()
	defer fake.listPullRequestsWithCommitMutex.Unlock()
	fake.ListPullRequestsWithCommitStub = nil
	if fake.listPullRequestsWithCommitReturnsOnCall == nil {
		fake.listPullRequestsWithCommitReturnsOnCall = make(map[int]struct {
			result1 []*githuba.PullRequest
			result2 *githuba.Response
			result3 error
		})
	}
	fake.listPullRequestsWithCommitReturnsOnCall[i] = struct {
		result1 []*githuba.PullRequest
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListReleases(arg1 context.Context, arg2 string, arg3 string, arg4 *githuba.ListOptions) ([]*githuba.RepositoryRelease, *githuba.Response, error) {
	fake.listReleasesMutex.Lock()
	ret, specificReturn := fake.listReleasesReturnsOnCall[len(fake.listReleasesArgsForCall)]
	fake.listReleasesArgsForCall = append(fake.listReleasesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *githuba.ListOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ListReleases", []interface{}{arg1, arg2, arg3, arg4})
	fake.listReleasesMutex.Unlock()
	if fake.ListReleasesStub != nil {
		return fake.ListReleasesStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listReleasesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListReleasesCallCount() int {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	return len(fake.listReleasesArgsForCall)
}

func (fake *FakeClient) ListReleasesCalls(stub func(context.Context, string, string, *githuba.ListOptions) ([]*githuba.RepositoryRelease, *githuba.Response, error)) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = stub
}

func (fake *FakeClient) ListReleasesArgsForCall(i int) (context.Context, string, string, *githuba.ListOptions) {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	argsForCall := fake.listReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) ListReleasesReturns(result1 []*githuba.RepositoryRelease, result2 *githuba.Response, result3 error) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = nil
	fake.listReleasesReturns = struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListReleasesReturnsOnCall(i int, result1 []*githuba.RepositoryRelease, result2 *githuba.Response, result3 error) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = nil
	if fake.listReleasesReturnsOnCall == nil {
		fake.listReleasesReturnsOnCall = make(map[int]struct {
			result1 []*githuba.RepositoryRelease
			result2 *githuba.Response
			result3 error
		})
	}
	fake.listReleasesReturnsOnCall[i] = struct {
		result1 []*githuba.RepositoryRelease
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListTags(arg1 context.Context, arg2 string, arg3 string, arg4 *githuba.ListOptions) ([]*githuba.RepositoryTag, *githuba.Response, error) {
	fake.listTagsMutex.Lock()
	ret, specificReturn := fake.listTagsReturnsOnCall[len(fake.listTagsArgsForCall)]
	fake.listTagsArgsForCall = append(fake.listTagsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 *githuba.ListOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ListTags", []interface{}{arg1, arg2, arg3, arg4})
	fake.listTagsMutex.Unlock()
	if fake.ListTagsStub != nil {
		return fake.ListTagsStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listTagsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListTagsCallCount() int {
	fake.listTagsMutex.RLock()
	defer fake.listTagsMutex.RUnlock()
	return len(fake.listTagsArgsForCall)
}

func (fake *FakeClient) ListTagsCalls(stub func(context.Context, string, string, *githuba.ListOptions) ([]*githuba.RepositoryTag, *githuba.Response, error)) {
	fake.listTagsMutex.Lock()
	defer fake.listTagsMutex.Unlock()
	fake.ListTagsStub = stub
}

func (fake *FakeClient) ListTagsArgsForCall(i int) (context.Context, string, string, *githuba.ListOptions) {
	fake.listTagsMutex.RLock()
	defer fake.listTagsMutex.RUnlock()
	argsForCall := fake.listTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) ListTagsReturns(result1 []*githuba.RepositoryTag, result2 *githuba.Response, result3 error) {
	fake.listTagsMutex.Lock()
	defer fake.listTagsMutex.Unlock()
	fake.ListTagsStub = nil
	fake.listTagsReturns = struct {
		result1 []*githuba.RepositoryTag
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListTagsReturnsOnCall(i int, result1 []*githuba.RepositoryTag, result2 *githuba.Response, result3 error) {
	fake.listTagsMutex.Lock()
	defer fake.listTagsMutex.Unlock()
	fake.ListTagsStub = nil
	if fake.listTagsReturnsOnCall == nil {
		fake.listTagsReturnsOnCall = make(map[int]struct {
			result1 []*githuba.RepositoryTag
			result2 *githuba.Response
			result3 error
		})
	}
	fake.listTagsReturnsOnCall[i] = struct {
		result1 []*githuba.RepositoryTag
		result2 *githuba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCommitMutex.RLock()
	defer fake.getCommitMutex.RUnlock()
	fake.getPullRequestMutex.RLock()
	defer fake.getPullRequestMutex.RUnlock()
	fake.getRepoCommitMutex.RLock()
	defer fake.getRepoCommitMutex.RUnlock()
	fake.listCommitsMutex.RLock()
	defer fake.listCommitsMutex.RUnlock()
	fake.listPullRequestsWithCommitMutex.RLock()
	defer fake.listPullRequestsWithCommitMutex.RUnlock()
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	fake.listTagsMutex.RLock()
	defer fake.listTagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ github.Client = new(FakeClient)
