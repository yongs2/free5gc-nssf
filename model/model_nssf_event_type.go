/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
    "fmt"
)

type NssfEventType string

// List of NssfEventType
const (
	NssfEventType_SNSSAI_STATUS_CHANGE_REPORT NssfEventType = "SNSSAI_STATUS_CHANGE_REPORT"
)

func (n *NssfEventType) CheckIntegrity() error {
    if *n != NssfEventType_SNSSAI_STATUS_CHANGE_REPORT {
        return fmt.Errorf("'%s' is unrecognized", string(*n))
    }

    return nil
}
