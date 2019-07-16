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
	"fmt"

    "free5gc-nssf/util/client"
)

// Linger please
var (
	_ context.Context
)

type SubscriptionIDDocumentApiService service

/*
SubscriptionIDDocumentApiService Deletes an already existing NSSAI availability notification subscription
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param subscriptionId Identifier of the subscription for notification
*/

func (a *SubscriptionIDDocumentApiService) NSSAIAvailabilityUnsubscribe(ctx context.Context, subscriptionId string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/nssai-availability/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", fmt.Sprintf("%v", subscriptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}



    localVarHTTPContentTypes := []string{"application/json"}

    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := client.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	r, err := client.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := client.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioclient.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

    apiError := client.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 204:
        return localVarHTTPResponse, nil
        case 400:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarHTTPResponse, apiError
        case 401:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarHTTPResponse, apiError
        case 404:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarHTTPResponse, apiError
        case 429:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarHTTPResponse, apiError
        case 500:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarHTTPResponse, apiError
        case 503:
            var v ProblemDetails
            err = client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarHTTPResponse, apiError
        default:
        return localVarHTTPResponse, nil
    }
}
