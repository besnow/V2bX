package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/InazumaV/V2bX/conf"
)

var (
	cores = map[string]func(c *conf.CoreConfig) (Core, error){}
)

func NewCore(c []conf.CoreConfig) (Core, error) {
	if len(c) == 0 {
		return nil, errors.New("no have vail core")
	}
	if len(c) == 1 {
		switch strings.ToLower(c[0].Type) {
		case "auto", "xray_prefer":
			return NewSelector([]conf.CoreConfig{
				{
					Type:       "xray",
					Name:       c[0].Name,
					XrayConfig: c[0].XrayConfig,
				}, {
					Type:       "sing",
					Name:       c[0].Name,
					SingConfig: c[0].SingConfig,
				},
			})
		case "xray", "sing":
			if f, ok := cores[c[0].Type]; ok {
				return f(&c[0])
			}
			return nil, fmt.Errorf("unknown core type: %s", c[0].Type)
		default:
			return nil, fmt.Errorf("unknown core type: %s", c[0].Type)
		}
	}
	return NewSelector(c)
}

func RegisterCore(t string, f func(c *conf.CoreConfig) (Core, error)) {
	cores[t] = f
}

func RegisteredCore() []string {
	cs := make([]string, 0, len(cores))
	for k := range cores {
		cs = append(cs, k)
	}
	return cs
}
