/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 */

package factory

import (
    "errors"
    "strconv"

    . "../model"
)

type Configuration struct {

    SupportedNssaiInPlmn []Snssai `yaml:"supportedNssaiInPlmn"`

    NsiList []NsiConfig `yaml:"nsiList,omitempty"`

    AmfSetList []AmfSetConfig `yaml:"amfSetList"`

    TaList []TaConfig `yaml:"taList"`

    MappingListFromPlmn []MappingFromPlmnConfig `yaml:"mappingListFromPlmn"`
}

func (c *Configuration) checkIntegrity() error {
    if c.SupportedNssaiInPlmn == nil || len(c.SupportedNssaiInPlmn) == 0 {
        return errors.New("`supportedNssaiInPlmn` in configuration should not be empty")
    } else {
        for i, supportedSnssaiInPlmn := range c.SupportedNssaiInPlmn {
            err := supportedSnssaiInPlmn.CheckIntegrity()
            if err != nil {
                errMsg := "`supportedNssaiInPlmn`[" + strconv.Itoa(i) + "]:" + err.Error()
                return errors.New(errMsg)
            }
        }
    }

    if c.NsiList != nil && len(c.NsiList) != 0 {
        for i, nsiConfig := range c.NsiList {
            err := nsiConfig.checkIntegrity()
            if err != nil {
                errMsg := "`nsiList`[" + strconv.Itoa(i) + "]:" + err.Error()
                return errors.New(errMsg)
            }
        }
    }

    if c.AmfSetList == nil || len(c.AmfSetList) == 0 {
        return errors.New("`amfSetList` in configuration should not be empty")
    } else {
        for i, amfSetConfig := range c.AmfSetList {
            err := amfSetConfig.checkIntegrity()
            if err != nil {
                errMsg := "`amfSetList`[" + strconv.Itoa(i) + "]:" + err.Error()
                return errors.New(errMsg)
            }
        }
    }

    if c.TaList == nil || len(c.TaList) == 0 {
        return errors.New("`taList` in configuration should not be empty")
    } else {
        for i, taConfig := range c.TaList {
            err := taConfig.checkIntegrity()
            if err != nil {
                errMsg := "`taList`[" + strconv.Itoa(i) + "]:" + err.Error()
                return errors.New(errMsg)
            }
        }
    }

    if c.MappingListFromPlmn == nil || len(c.MappingListFromPlmn) == 0 {
        return errors.New("`mappingListFromPlmn` in configuration should not be empty")
    } else {
        for i, mappingFromPlmnConfig := range c.MappingListFromPlmn {
            err := mappingFromPlmnConfig.checkIntegrity()
            if err != nil {
                errMsg := "`mappingListFromPlmn`[" + strconv.Itoa(i) + "]:" + err.Error()
                return errors.New(errMsg)
            }
        }
    }

    return nil
}
