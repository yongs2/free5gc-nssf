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
    "reflect"
)

type PatchItem struct {

	Op PatchOperation `json:"op"`

	Path string `json:"path"`

	From string `json:"from,omitempty"`

	// Value *map[string]interface{} `json:"value,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

func (p *PatchItem) CheckIntegrity() error {
    if p.Op == PatchOperation("") {
        return fmt.Errorf("`op` should not be empty")
    } else {
        err := p.Op.CheckIntegrity()
        if err != nil {
            return fmt.Errorf("`op`:%s", err.Error())
        }
    }

    if p.Path == "" {
        return fmt.Errorf("`path` should not be empty")
    }

    if (p.Op == PatchOperation_MOVE || p.Op == PatchOperation_COPY) && p.From == "" {
        return fmt.Errorf("`from` should not be empty with `op`:'%s' operation", string(p.Op))
    }

    if (p.Op == PatchOperation_ADD || p.Op == PatchOperation_REPLACE || p.Op == PatchOperation_TEST) &&
        (p.Value == nil || reflect.DeepEqual(p.Value, reflect.Zero(reflect.TypeOf(p.Value)).Interface())) {
        return fmt.Errorf("`value` should not be empty with `op`:'%s' operation", string(p.Op))
    }

    return nil
}
