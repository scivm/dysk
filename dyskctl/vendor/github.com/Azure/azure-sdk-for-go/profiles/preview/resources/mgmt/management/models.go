// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 2014fbbf031942474ad27a5a66dffaed5347f3fb

package management

import (
	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-08-31-preview/management"
	uuid "github.com/satori/go.uuid"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type GroupsClient = original.GroupsClient
type ChildType = original.ChildType

const (
	Account      ChildType = original.Account
	Department   ChildType = original.Department
	Enrollment   ChildType = original.Enrollment
	Subscription ChildType = original.Subscription
)

type ChildType1 = original.ChildType1

const (
	ChildType1Account      ChildType1 = original.ChildType1Account
	ChildType1Department   ChildType1 = original.ChildType1Department
	ChildType1Enrollment   ChildType1 = original.ChildType1Enrollment
	ChildType1Subscription ChildType1 = original.ChildType1Subscription
)

type ManagementGroupType = original.ManagementGroupType

const (
	ManagementGroupTypeAccount      ManagementGroupType = original.ManagementGroupTypeAccount
	ManagementGroupTypeDepartment   ManagementGroupType = original.ManagementGroupTypeDepartment
	ManagementGroupTypeEnrollment   ManagementGroupType = original.ManagementGroupTypeEnrollment
	ManagementGroupTypeSubscription ManagementGroupType = original.ManagementGroupTypeSubscription
)

type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Group = original.Group
type GroupChildInfo = original.GroupChildInfo
type GroupDetailsProperties = original.GroupDetailsProperties
type GroupInfo = original.GroupInfo
type GroupInfoProperties = original.GroupInfoProperties
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupProperties = original.GroupProperties
type GroupPropertiesWithChildren = original.GroupPropertiesWithChildren
type GroupPropertiesWithHierarchy = original.GroupPropertiesWithHierarchy
type GroupRecursiveChildInfo = original.GroupRecursiveChildInfo
type GroupWithChildren = original.GroupWithChildren
type GroupWithHierarchy = original.GroupWithHierarchy
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type ParentGroupInfo = original.ParentGroupInfo
type OperationsClient = original.OperationsClient

func New(groupID uuid.UUID) BaseClient {
	return original.New(groupID)
}
func NewWithBaseURI(baseURI string, groupID uuid.UUID) BaseClient {
	return original.NewWithBaseURI(baseURI, groupID)
}
func NewGroupsClient(groupID uuid.UUID) GroupsClient {
	return original.NewGroupsClient(groupID)
}
func NewGroupsClientWithBaseURI(baseURI string, groupID uuid.UUID) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, groupID)
}
func NewOperationsClient(groupID uuid.UUID) OperationsClient {
	return original.NewOperationsClient(groupID)
}
func NewOperationsClientWithBaseURI(baseURI string, groupID uuid.UUID) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, groupID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
