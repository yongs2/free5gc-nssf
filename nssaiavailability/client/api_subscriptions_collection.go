/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

    . "free5gc-nssf/model"
    "free5gc-nssf/util/client"
)

// Linger please
var (
	_ context.Context
)

type SubscriptionsCollectionApiService service

/*
SubscriptionsCollectionApiService Creates subscriptions for notification about updates to NSSAI availability information
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param nssfEventSubscriptionCreateData Subscription for notification about updates to NSSAI availability information
@return NssfEventSubscriptionCreatedData
*/

func (a *SubscriptionsCollectionApiService) NSSAIAvailabilityPost(ctx context.Context, nssfEventSubscriptionCreateData NssfEventSubscriptionCreateData) (NssfEventSubscriptionCreatedData, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NssfEventSubscriptionCreatedData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nssai-availability/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


    localVarHTTPContentTypes := []string{ "application/json" }


    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "application/json", "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := client.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = &nssfEventSubscriptionCreateData

	r, err := client.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

    apiError := client.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 201:
            err = client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
            }
            return localVarReturnValue, localVarHTTPResponse, nil
        case 400:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 401:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 403:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 404:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 411:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 413:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 415:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 429:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 500:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 503:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        default:
        return localVarReturnValue, localVarHTTPResponse, nil
    }
}