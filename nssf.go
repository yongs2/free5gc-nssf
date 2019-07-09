/*
 * NSSF
 *
 * Network Slice Selection Function
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
    "./nssf_service"
)

func main() {
    var nssfService nssf_service.Nssf
    nssfService.Cfg = "./conf/nssf_config.yaml"

    nssfService.Initialize()
    nssfService.Start()
}
