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
    "net/http"
)

// APIClient manages communication with the NSSF NSSAI Availability API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg    *Configuration
    common service // Reuse a single struct instead of allocating one for each service on the heap.

    // API Services
	NFInstanceIDDocumentApi *NFInstanceIDDocumentApiService
	SubscriptionIDDocumentApi *SubscriptionIDDocumentApiService
	SubscriptionsCollectionApi *SubscriptionsCollectionApiService
}

type service struct {
    client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
    if cfg.httpClient == nil {
        cfg.httpClient = http.DefaultClient
    }

    c := &APIClient{}
    c.cfg = cfg
    c.common.client = c

    // API Services
    c.NFInstanceIDDocumentApi = (*NFInstanceIDDocumentApiService)(&c.common)
    c.SubscriptionIDDocumentApi = (*SubscriptionIDDocumentApiService)(&c.common)
    c.SubscriptionsCollectionApi = (*SubscriptionsCollectionApiService)(&c.common)

    return c
}