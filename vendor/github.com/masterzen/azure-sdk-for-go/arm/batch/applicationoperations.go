package batch

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

// ApplicationOperationsClient is the client for the ApplicationOperations
// methods of the Batch service.
type ApplicationOperationsClient struct {
	ManagementClient
}

// NewApplicationOperationsClient creates an instance of the
// ApplicationOperationsClient client.
func NewApplicationOperationsClient(subscriptionID string) ApplicationOperationsClient {
	return NewApplicationOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationOperationsClientWithBaseURI creates an instance of the
// ApplicationOperationsClient client.
func NewApplicationOperationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationOperationsClient {
	return ApplicationOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ActivateApplicationPackage activates the specified application package.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. id is the id of the
// application. version is the version of the application to activate.
// parameters is the parameters for the request.
func (client ApplicationOperationsClient) ActivateApplicationPackage(resourceGroupName string, accountName string, id string, version string, parameters ActivateApplicationPackageParameters) (result autorest.Response, err error) {
	req, err := client.ActivateApplicationPackagePreparer(resourceGroupName, accountName, id, version, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "ActivateApplicationPackage", nil, "Failure preparing request")
	}

	resp, err := client.ActivateApplicationPackageSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "ActivateApplicationPackage", resp, "Failure sending request")
	}

	result, err = client.ActivateApplicationPackageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "ActivateApplicationPackage", resp, "Failure responding to request")
	}

	return
}

// ActivateApplicationPackagePreparer prepares the ActivateApplicationPackage request.
func (client ApplicationOperationsClient) ActivateApplicationPackagePreparer(resourceGroupName string, accountName string, id string, version string, parameters ActivateApplicationPackageParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"id":                autorest.Encode("path", id),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{id}/versions/{version}/activate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ActivateApplicationPackageSender sends the ActivateApplicationPackage request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) ActivateApplicationPackageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ActivateApplicationPackageResponder handles the response to the ActivateApplicationPackage request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) ActivateApplicationPackageResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// AddApplication adds an application to the specified Batch account.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application. parameters is the parameters for the request.
func (client ApplicationOperationsClient) AddApplication(resourceGroupName string, accountName string, applicationID string, parameters *AddApplicationParameters) (result Application, err error) {
	req, err := client.AddApplicationPreparer(resourceGroupName, accountName, applicationID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "AddApplication", nil, "Failure preparing request")
	}

	resp, err := client.AddApplicationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "AddApplication", resp, "Failure sending request")
	}

	result, err = client.AddApplicationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "AddApplication", resp, "Failure responding to request")
	}

	return
}

// AddApplicationPreparer prepares the AddApplication request.
func (client ApplicationOperationsClient) AddApplicationPreparer(resourceGroupName string, accountName string, applicationID string, parameters *AddApplicationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare(&http.Request{})
}

// AddApplicationSender sends the AddApplication request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) AddApplicationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddApplicationResponder handles the response to the AddApplication request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) AddApplicationResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddApplicationPackage creates an application package record.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application. version is the version of the application.
func (client ApplicationOperationsClient) AddApplicationPackage(resourceGroupName string, accountName string, applicationID string, version string) (result AddApplicationPackageResult, err error) {
	req, err := client.AddApplicationPackagePreparer(resourceGroupName, accountName, applicationID, version)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "AddApplicationPackage", nil, "Failure preparing request")
	}

	resp, err := client.AddApplicationPackageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "AddApplicationPackage", resp, "Failure sending request")
	}

	result, err = client.AddApplicationPackageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "AddApplicationPackage", resp, "Failure responding to request")
	}

	return
}

// AddApplicationPackagePreparer prepares the AddApplicationPackage request.
func (client ApplicationOperationsClient) AddApplicationPackagePreparer(resourceGroupName string, accountName string, applicationID string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// AddApplicationPackageSender sends the AddApplicationPackage request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) AddApplicationPackageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddApplicationPackageResponder handles the response to the AddApplicationPackage request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) AddApplicationPackageResponder(resp *http.Response) (result AddApplicationPackageResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteApplication deletes an application.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application.
func (client ApplicationOperationsClient) DeleteApplication(resourceGroupName string, accountName string, applicationID string) (result autorest.Response, err error) {
	req, err := client.DeleteApplicationPreparer(resourceGroupName, accountName, applicationID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "DeleteApplication", nil, "Failure preparing request")
	}

	resp, err := client.DeleteApplicationSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "DeleteApplication", resp, "Failure sending request")
	}

	result, err = client.DeleteApplicationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "DeleteApplication", resp, "Failure responding to request")
	}

	return
}

// DeleteApplicationPreparer prepares the DeleteApplication request.
func (client ApplicationOperationsClient) DeleteApplicationPreparer(resourceGroupName string, accountName string, applicationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteApplicationSender sends the DeleteApplication request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) DeleteApplicationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteApplicationResponder handles the response to the DeleteApplication request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) DeleteApplicationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteApplicationPackage deletes an application package record and its
// associated binary file.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application. version is the version of the application to
// delete.
func (client ApplicationOperationsClient) DeleteApplicationPackage(resourceGroupName string, accountName string, applicationID string, version string) (result autorest.Response, err error) {
	req, err := client.DeleteApplicationPackagePreparer(resourceGroupName, accountName, applicationID, version)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "DeleteApplicationPackage", nil, "Failure preparing request")
	}

	resp, err := client.DeleteApplicationPackageSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "DeleteApplicationPackage", resp, "Failure sending request")
	}

	result, err = client.DeleteApplicationPackageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "DeleteApplicationPackage", resp, "Failure responding to request")
	}

	return
}

// DeleteApplicationPackagePreparer prepares the DeleteApplicationPackage request.
func (client ApplicationOperationsClient) DeleteApplicationPackagePreparer(resourceGroupName string, accountName string, applicationID string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteApplicationPackageSender sends the DeleteApplicationPackage request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) DeleteApplicationPackageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteApplicationPackageResponder handles the response to the DeleteApplicationPackage request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) DeleteApplicationPackageResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetApplication gets information about the specified application.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application.
func (client ApplicationOperationsClient) GetApplication(resourceGroupName string, accountName string, applicationID string) (result Application, err error) {
	req, err := client.GetApplicationPreparer(resourceGroupName, accountName, applicationID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "GetApplication", nil, "Failure preparing request")
	}

	resp, err := client.GetApplicationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "GetApplication", resp, "Failure sending request")
	}

	result, err = client.GetApplicationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "GetApplication", resp, "Failure responding to request")
	}

	return
}

// GetApplicationPreparer prepares the GetApplication request.
func (client ApplicationOperationsClient) GetApplicationPreparer(resourceGroupName string, accountName string, applicationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetApplicationSender sends the GetApplication request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) GetApplicationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetApplicationResponder handles the response to the GetApplication request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) GetApplicationResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetApplicationPackage gets information about the specified application
// package.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application. version is the version of the application.
func (client ApplicationOperationsClient) GetApplicationPackage(resourceGroupName string, accountName string, applicationID string, version string) (result GetApplicationPackageResult, err error) {
	req, err := client.GetApplicationPackagePreparer(resourceGroupName, accountName, applicationID, version)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "GetApplicationPackage", nil, "Failure preparing request")
	}

	resp, err := client.GetApplicationPackageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "GetApplicationPackage", resp, "Failure sending request")
	}

	result, err = client.GetApplicationPackageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "GetApplicationPackage", resp, "Failure responding to request")
	}

	return
}

// GetApplicationPackagePreparer prepares the GetApplicationPackage request.
func (client ApplicationOperationsClient) GetApplicationPackagePreparer(resourceGroupName string, accountName string, applicationID string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetApplicationPackageSender sends the GetApplicationPackage request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) GetApplicationPackageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetApplicationPackageResponder handles the response to the GetApplicationPackage request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) GetApplicationPackageResponder(resp *http.Response) (result GetApplicationPackageResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all of the applications in the specified account.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. maxresults is the
// maximum number of items to return in the response.
func (client ApplicationOperationsClient) List(resourceGroupName string, accountName string, maxresults *int32) (result ListApplicationsResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, accountName, maxresults)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationOperationsClient) ListPreparer(resourceGroupName string, accountName string, maxresults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) ListResponder(resp *http.Response) (result ListApplicationsResult, err error) {
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
func (client ApplicationOperationsClient) ListNextResults(lastResults ListApplicationsResult) (result ListApplicationsResult, err error) {
	req, err := lastResults.ListApplicationsResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// UpdateApplication updates settings for the specified application.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is
// the id of the application. parameters is the parameters for the request.
func (client ApplicationOperationsClient) UpdateApplication(resourceGroupName string, accountName string, applicationID string, parameters UpdateApplicationParameters) (result autorest.Response, err error) {
	req, err := client.UpdateApplicationPreparer(resourceGroupName, accountName, applicationID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "UpdateApplication", nil, "Failure preparing request")
	}

	resp, err := client.UpdateApplicationSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "UpdateApplication", resp, "Failure sending request")
	}

	result, err = client.UpdateApplicationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationOperationsClient", "UpdateApplication", resp, "Failure responding to request")
	}

	return
}

// UpdateApplicationPreparer prepares the UpdateApplication request.
func (client ApplicationOperationsClient) UpdateApplicationPreparer(resourceGroupName string, accountName string, applicationID string, parameters UpdateApplicationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateApplicationSender sends the UpdateApplication request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationOperationsClient) UpdateApplicationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateApplicationResponder handles the response to the UpdateApplication request. The method always
// closes the http.Response Body.
func (client ApplicationOperationsClient) UpdateApplicationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
