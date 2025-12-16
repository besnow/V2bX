package core

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/InazumaV/V2bX/api/panel"
	"github.com/InazumaV/V2bX/conf"
)

type Selector struct {
	cores map[string]Core
	types map[string]Core
	nodes sync.Map
}

func NewSelector(c []conf.CoreConfig) (Core, error) {
	cs := make(map[string]Core, len(c))
	types := make(map[string]Core, len(c))
	for _, t := range c {
		f, ok := cores[strings.ToLower(t.Type)]
		if !ok {
			return nil, errors.New("unknown core type: " + t.Type)
		}
		core1, err := f(&t)
		if err != nil {
			return nil, err
		}
		if t.Name == "" {
			cs[t.Type] = core1
		} else {
			cs[t.Name] = core1
		}
		types[core1.Type()] = core1
	}
	return &Selector{
		cores: cs,
		types: types,
	}, nil
}

func (s *Selector) Start() error {
	for i := range s.cores {
		err := s.cores[i].Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Selector) Close() error {
	var errs []error
	for i := range s.cores {
		if err := s.cores[i].Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func isSupported(protocol string, protocols []string) bool {
	protocol = strings.ToLower(protocol)
	for i := range protocols {
		if protocol == strings.ToLower(protocols[i]) {
			return true
		}
	}
	return false
}

func (s *Selector) AddNode(tag string, info *panel.NodeInfo, option *conf.Options) error {
	core, err := s.selectCore(info, option)
	if err != nil {
		return err
	}
	if len(option.Core) == 0 {
		option.Core = core.Type()
		err := option.UnmarshalJSON(option.RawOptions)
		if err != nil {
			return fmt.Errorf("unmarshal option error: %s", err)
		}
		option.RawOptions = nil
	}
	err = core.AddNode(tag, info, option)
	if err != nil {
		return err
	}
	s.nodes.Store(tag, core)
	return nil
}

func (s *Selector) selectCore(info *panel.NodeInfo, option *conf.Options) (Core, error) {
	if len(option.CoreName) > 0 {
		if c, ok := s.cores[option.CoreName]; ok {
			return c, nil
		}
		return nil, fmt.Errorf("unknown core name: %s", option.CoreName)
	}

	requested := strings.ToLower(option.Core)
	switch requested {
	case "", "auto", "xray_prefer":
		requested = ""
	default:
		if c, ok := s.types[requested]; ok {
			if !isSupported(info.Type, c.Protocols()) {
				return nil, fmt.Errorf("core %s does not support protocol %s", requested, info.Type)
			}
			return c, nil
		}
		if c, ok := s.cores[option.Core]; ok {
			if !isSupported(info.Type, c.Protocols()) {
				return nil, fmt.Errorf("core %s does not support protocol %s", option.Core, info.Type)
			}
			return c, nil
		}
		return nil, fmt.Errorf("unknown core type: %s", option.Core)
	}

	preferred := "xray"
	switch strings.ToLower(info.Type) {
	case "hysteria", "hysteria2", "tuic", "anytls":
		preferred = "sing"
	}

	if c, ok := s.types[preferred]; ok {
		if isSupported(info.Type, c.Protocols()) {
			return c, nil
		}
	}

	for _, c := range s.types {
		if isSupported(info.Type, c.Protocols()) {
			return c, nil
		}
	}
	return nil, errors.New("the node type is not support")
}

func (s *Selector) DelNode(tag string) error {
	if t, e := s.nodes.Load(tag); e {
		err := t.(Core).DelNode(tag)
		if err != nil {
			return err
		}
		s.nodes.Delete(tag)
		return nil
	}
	return errors.New("the node is not have")
}

func (s *Selector) AddUsers(p *AddUsersParams) (added int, err error) {
	t, e := s.nodes.Load(p.Tag)
	if !e {
		return 0, errors.New("the node is not have")
	}
	return t.(Core).AddUsers(p)
}

func (s *Selector) GetUserTrafficSlice(tag string, reset bool) ([]panel.UserTraffic, error) {
	t, e := s.nodes.Load(tag)
	if !e {
		return nil, errors.New("the node is not have")
	}
	return t.(Core).GetUserTrafficSlice(tag, reset)
}

func (s *Selector) DelUsers(users []panel.UserInfo, tag string, info *panel.NodeInfo) error {
	t, e := s.nodes.Load(tag)
	if !e {
		return errors.New("the node is not have")
	}
	return t.(Core).DelUsers(users, tag, info)
}

func (s *Selector) Protocols() []string {
	protocols := make([]string, 0)
	for i := range s.types {
		protocols = append(protocols, s.types[i].Protocols()...)
	}
	return protocols
}

func (s *Selector) Type() string {
	t := "Selector("
	var flag bool
	for n, c := range s.cores {
		if flag {
			t += " "
		} else {
			flag = true
		}
		if len(n) == 0 {
			t += c.Type()
		} else {
			t += n
		}
	}
	t += ")"
	return t
}
