/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
    "errors"
    "strconv"
)

type AllowedNssai struct {

	AllowedSnssaiList []AllowedSnssai `json:"allowedSnssaiList"`

	AccessType *AccessType `json:"accessType"`
}

func (a *AllowedNssai) CheckIntegrity() error {
    // Deal with both nil slice and empty slice
    if a.AllowedSnssaiList == nil || len(a.AllowedSnssaiList) == 0 {
        return errors.New("`allowedSnssaiList` in query parameter should not be empty")
    } else {
        for i, allowedSnssai := range a.AllowedSnssaiList {
            err := allowedSnssai.CheckIntegrity()
            if err != nil {
                errMsg := "`allowedSnssaiList`[" + strconv.Itoa(i) + "]:" + err.Error()
                return errors.New(errMsg)
            }
        }
    }

    if a.AccessType == nil || *a.AccessType == "" {
        return errors.New("`accessType` in query parameter should not be empty")
    } else {
        err := a.AccessType.CheckIntegrity()
        if err != nil {
            errMsg := "`accessType`:" + err.Error()
            return errors.New(errMsg)
        }
    }

    return nil
}
