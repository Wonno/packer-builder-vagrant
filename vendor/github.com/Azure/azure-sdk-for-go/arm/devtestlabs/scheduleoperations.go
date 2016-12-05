package devtestlabs

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ScheduleOperationsClient is the azure DevTest Labs REST API.
type ScheduleOperationsClient struct {
	ManagementClient
}

// NewScheduleOperationsClient creates an instance of the
// ScheduleOperationsClient client.
func NewScheduleOperationsClient(subscriptionID string) ScheduleOperationsClient {
	return NewScheduleOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScheduleOperationsClientWithBaseURI creates an instance of the
// ScheduleOperationsClient client.
func NewScheduleOperationsClientWithBaseURI(baseURI string, subscriptionID string) ScheduleOperationsClient {
	return ScheduleOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateResource create or replace an existing schedule.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the schedule.
func (client ScheduleOperationsClient) CreateOrUpdateResource(resourceGroupName string, labName string, name string, schedule Schedule) (result Schedule, err error) {
	req, err := client.CreateOrUpdateResourcePreparer(resourceGroupName, labName, name, schedule)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "CreateOrUpdateResource", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "CreateOrUpdateResource", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "CreateOrUpdateResource", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client ScheduleOperationsClient) CreateOrUpdateResourcePreparer(resourceGroupName string, labName string, name string, schedule Schedule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", pathParameters),
		autorest.WithJSON(schedule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleOperationsClient) CreateOrUpdateResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client ScheduleOperationsClient) CreateOrUpdateResourceResponder(resp *http.Response) (result Schedule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteResource delete schedule.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the schedule.
func (client ScheduleOperationsClient) DeleteResource(resourceGroupName string, labName string, name string) (result autorest.Response, err error) {
	req, err := client.DeleteResourcePreparer(resourceGroupName, labName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "DeleteResource", nil, "Failure preparing request")
	}

	resp, err := client.DeleteResourceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "DeleteResource", resp, "Failure sending request")
	}

	result, err = client.DeleteResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "DeleteResource", resp, "Failure responding to request")
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client ScheduleOperationsClient) DeleteResourcePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleOperationsClient) DeleteResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client ScheduleOperationsClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Execute execute a schedule. This operation can take a while to complete.
// This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling
// and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the schedule.
func (client ScheduleOperationsClient) Execute(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.ExecutePreparer(resourceGroupName, labName, name, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "Execute", nil, "Failure preparing request")
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "Execute", resp, "Failure sending request")
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "Execute", resp, "Failure responding to request")
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client ScheduleOperationsClient) ExecutePreparer(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}/execute", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleOperationsClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client ScheduleOperationsClient) ExecuteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get schedule.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the schedule.
func (client ScheduleOperationsClient) GetResource(resourceGroupName string, labName string, name string) (result Schedule, err error) {
	req, err := client.GetResourcePreparer(resourceGroupName, labName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "GetResource", nil, "Failure preparing request")
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "GetResource", resp, "Failure sending request")
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client ScheduleOperationsClient) GetResourcePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleOperationsClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client ScheduleOperationsClient) GetResourceResponder(resp *http.Response) (result Schedule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list schedules in a given lab.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. filter is the filter to apply on the operation. top is the
// maximum number of resources to return from the operation. orderBy is the
// ordering expression for the results, using OData notation.
func (client ScheduleOperationsClient) List(resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationSchedule, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, filter, top, orderBy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ScheduleOperationsClient) ListPreparer(resourceGroupName string, labName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScheduleOperationsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationSchedule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ScheduleOperationsClient) ListNextResults(lastResults ResponseWithContinuationSchedule) (result ResponseWithContinuationSchedule, err error) {
	req, err := lastResults.ResponseWithContinuationSchedulePreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// PatchResource modify properties of schedules.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the schedule.
func (client ScheduleOperationsClient) PatchResource(resourceGroupName string, labName string, name string, schedule Schedule) (result Schedule, err error) {
	req, err := client.PatchResourcePreparer(resourceGroupName, labName, name, schedule)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "PatchResource", nil, "Failure preparing request")
	}

	resp, err := client.PatchResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "PatchResource", resp, "Failure sending request")
	}

	result, err = client.PatchResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.ScheduleOperationsClient", "PatchResource", resp, "Failure responding to request")
	}

	return
}

// PatchResourcePreparer prepares the PatchResource request.
func (client ScheduleOperationsClient) PatchResourcePreparer(resourceGroupName string, labName string, name string, schedule Schedule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", pathParameters),
		autorest.WithJSON(schedule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchResourceSender sends the PatchResource request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleOperationsClient) PatchResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResourceResponder handles the response to the PatchResource request. The method always
// closes the http.Response Body.
func (client ScheduleOperationsClient) PatchResourceResponder(resp *http.Response) (result Schedule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
