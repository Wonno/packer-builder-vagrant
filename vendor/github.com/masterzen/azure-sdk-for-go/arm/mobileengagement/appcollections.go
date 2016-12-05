package mobileengagement

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

// AppCollectionsClient is the microsoft Azure Mobile Engagement REST APIs.
type AppCollectionsClient struct {
	ManagementClient
}

// NewAppCollectionsClient creates an instance of the AppCollectionsClient
// client.
func NewAppCollectionsClient(subscriptionID string, resourceGroupName string, appCollection string, appName string) AppCollectionsClient {
	return NewAppCollectionsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, appCollection, appName)
}

// NewAppCollectionsClientWithBaseURI creates an instance of the
// AppCollectionsClient client.
func NewAppCollectionsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, appCollection string, appName string) AppCollectionsClient {
	return AppCollectionsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, appCollection, appName)}
}

// CheckNameAvailability checks availability of an app collection name in the
// Engagement domain.
//
func (client AppCollectionsClient) CheckNameAvailability(parameters AppCollectionNameAvailability) (result AppCollectionNameAvailability, err error) {
	req, err := client.CheckNameAvailabilityPreparer(parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "CheckNameAvailability", nil, "Failure preparing request")
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "CheckNameAvailability", resp, "Failure sending request")
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client AppCollectionsClient) CheckNameAvailabilityPreparer(parameters AppCollectionNameAvailability) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MobileEngagement/checkAppCollectionNameAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client AppCollectionsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client AppCollectionsClient) CheckNameAvailabilityResponder(resp *http.Response) (result AppCollectionNameAvailability, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists app collections in a subscription.
func (client AppCollectionsClient) List() (result AppCollectionListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AppCollectionsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MobileEngagement/appCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AppCollectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AppCollectionsClient) ListResponder(resp *http.Response) (result AppCollectionListResult, err error) {
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
func (client AppCollectionsClient) ListNextResults(lastResults AppCollectionListResult) (result AppCollectionListResult, err error) {
	req, err := lastResults.AppCollectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.AppCollectionsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
