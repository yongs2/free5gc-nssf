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
    "fmt"
    "sort"
)

type AllowedNssai struct {

	AllowedSnssaiList []AllowedSnssai `json:"allowedSnssaiList"`

	AccessType *AccessType `json:"accessType"`
}

func (a *AllowedNssai) CheckIntegrity() error {
    // Deal with both nil slice and empty slice
    if a.AllowedSnssaiList == nil || len(a.AllowedSnssaiList) == 0 {
        return fmt.Errorf("`allowedSnssaiList` should not be empty")
    } else {
        for i, allowedSnssai := range a.AllowedSnssaiList {
            err := allowedSnssai.CheckIntegrity()
            if err != nil {
                return fmt.Errorf("`allowedSnssaiList`[%d]:%s", i, err.Error())
            }
        }
    }

    if a.AccessType == nil || *a.AccessType == AccessType("") {
        return fmt.Errorf("`accessType` should not be empty")
    } else {
        err := a.AccessType.CheckIntegrity()
        if err != nil {
            return fmt.Errorf("`accessType`:%s", err.Error())
        }
    }

    return nil
}

type ByAccessType []AllowedNssai

func (a ByAccessType) Len() int {
    return len(a)
}

func (a ByAccessType) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByAccessType) Less(i, j int) bool {
    // Only 3GPP Access and Non 3GPP Access could be in Allowed NSSAI list
    // Sort it into 3GPP Access first and Non 3GPP Access last
    if *a[i].AccessType == AccessType__3_GPP_ACCESS {
        return true
    }
    return false
}

func (a *ByAccessType) Sort() {
    for i := range *a {
        sort.Sort(ByAllowedSnssai((*a)[i].AllowedSnssaiList))
    }

    sort.Sort(*a)
}
