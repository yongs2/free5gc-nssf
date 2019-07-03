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
)

type SliceInfoForPduSession struct {

	SNssai *Snssai `json:"sNssai"`

	RoamingIndication *RoamingIndication `json:"roamingIndication"`

	HomeSnssai *Snssai `json:"homeSnssai,omitempty"`
}

func (s *SliceInfoForPduSession) CheckIntegrity() error {
    if s.SNssai == nil {
        return fmt.Errorf("`sNssai` should not be empty")
    } else {
        err := s.SNssai.CheckIntegrity()
        if err != nil {
            return fmt.Errorf("`sNssai`:%s", err.Error())
        }
    }

    if s.RoamingIndication == nil || *s.RoamingIndication == "" {
        return fmt.Errorf("`roamingIndication` should not be empty")
    } else {
        err := s.RoamingIndication.CheckIntegrity()
        if err != nil {
            return fmt.Errorf("`roamingIndication`:%s", err.Error())
        }
    }

    if *s.RoamingIndication == RoamingIndication_HOME_ROUTED_ROAMING && s.HomeSnssai == nil {
        return fmt.Errorf("`homeSnssai` should be included in home routed roaming scenario")
    } else if s.HomeSnssai != nil {
        err := s.HomeSnssai.CheckIntegrity()
        if err != nil {
            return fmt.Errorf("`homeSnssai`:%s", err.Error())
        }
    }

    return nil
}
