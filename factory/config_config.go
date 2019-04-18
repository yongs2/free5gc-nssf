/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 */

package factory

import (
    "fmt"
)

type Config struct {

    Info *Info `yaml:"info"`

    Configuration *Configuration `yaml:"configuration"`
}

func (c *Config) CheckIntegrity() error {
    if c.Info == nil {
        return fmt.Errorf("`info` in configuration should not be empty")
    } else {
        err := c.Info.checkIntegrity()
        if err != nil {
            return fmt.Errorf("`info`:%s", err.Error())
        }
    }

    if c.Configuration == nil {
        return fmt.Errorf("`configuration` in configuration should not be empty")
    } else {
        err := c.Configuration.checkIntegrity()
        if err != nil {
            return fmt.Errorf("`configuration`:%s", err.Error())
        }
    }

    return nil
}