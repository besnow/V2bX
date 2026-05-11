package core

import (
	"fmt"
	"strings"

	"github.com/InazumaV/V2bX/conf"
)

var (
	cores = map[string]func(c *conf.CoreConfig) (Core, error){}
)

func normalizeCoreType(t string) string {
	return strings.ToLower(strings.TrimSpace(t))
}

func unsupportedCoreError(t string) error {
	return fmt.Errorf("core type %q is not supported by this xray-only build; configure Core/Type as %q", t, "xray")
}

func NewCore(c []conf.CoreConfig) (Core, error) {
	if len(c) == 0 {
		return nil, fmt.Errorf("no valid core configured; this xray-only build requires one xray core")
	}
	// multi core
	if len(c) > 1 {
		return NewSelector(c)
	}
	// one core
	coreType := normalizeCoreType(c[0].Type)
	if f, ok := cores[coreType]; ok {
		return f(&c[0])
	}
	return nil, unsupportedCoreError(c[0].Type)
}

func RegisterCore(t string, f func(c *conf.CoreConfig) (Core, error)) {
	cores[normalizeCoreType(t)] = f
}

func RegisteredCore() []string {
	cs := make([]string, 0, len(cores))
	for k := range cores {
		cs = append(cs, k)
	}
	return cs
}
