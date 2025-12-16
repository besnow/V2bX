package conf

import (
	"encoding/json"
)

type CoreConfig struct {
	Type       string      `json:"Type"`
	Name       string      `json:"Name"`
	XrayConfig *XrayConfig `json:"-"`
	SingConfig *SingConfig `json:"-"`
}

type _CoreConfig CoreConfig

func (c *CoreConfig) UnmarshalJSON(b []byte) error {
	if len(c.Type) == 0 {
		c.Type = "auto"
	}
	err := json.Unmarshal(b, (*_CoreConfig)(c))
	if err != nil {
		return err
	}
	switch c.Type {
	case "auto", "xray_prefer":
		c.XrayConfig = NewXrayConfig()
		c.SingConfig = NewSingConfig()
		if err := json.Unmarshal(b, c.XrayConfig); err != nil {
			return err
		}
		return json.Unmarshal(b, c.SingConfig)
	case "xray":
		c.XrayConfig = NewXrayConfig()
		return json.Unmarshal(b, c.XrayConfig)
	case "sing":
		c.SingConfig = NewSingConfig()
		return json.Unmarshal(b, c.SingConfig)
	}
	return nil
}
