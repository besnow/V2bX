package xray

import (
	"github.com/InazumaV/V2bX/api/panel"
	"github.com/InazumaV/V2bX/common/format"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/anytls"
	"github.com/xtls/xray-core/proxy/hysteria2"
	"github.com/xtls/xray-core/proxy/tuic"
)

func buildAnyTLSUsers(tag string, userInfo []panel.UserInfo) (users []*protocol.User) {
	users = make([]*protocol.User, len(userInfo))
	for i, user := range userInfo {
		users[i] = &protocol.User{
			Level: 0,
			Email: format.UserTag(tag, user.Uuid),
			Account: serial.ToTypedMessage(&anytls.Account{
				Password: user.Uuid,
			}),
		}
	}
	return users
}

func buildHysteria2Users(tag string, userInfo []panel.UserInfo) (users []*protocol.User) {
	users = make([]*protocol.User, len(userInfo))
	for i, user := range userInfo {
		users[i] = &protocol.User{
			Level: 0,
			Email: format.UserTag(tag, user.Uuid),
			Account: serial.ToTypedMessage(&hysteria2.Account{
				Password: user.Uuid,
			}),
		}
	}
	return users
}

func buildTuicUsers(tag string, userInfo []panel.UserInfo) (users []*protocol.User) {
	users = make([]*protocol.User, len(userInfo))
	for i, user := range userInfo {
		users[i] = &protocol.User{
			Level: 0,
			Email: format.UserTag(tag, user.Uuid),
			Account: serial.ToTypedMessage(&tuic.Account{
				Uuid:     user.Uuid,
				Password: user.Uuid,
			}),
		}
	}
	return users
}
