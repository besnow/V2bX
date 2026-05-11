# V2bX

[![](https://img.shields.io/badge/TgChat-UnOfficialV2Board%E4%BA%A4%E6%B5%81%E7%BE%A4-green)](https://t.me/unofficialV2board)
[![](https://img.shields.io/badge/TgChat-YuzukiProjects%E4%BA%A4%E6%B5%81%E7%BE%A4-blue)](https://t.me/YuzukiProjects)

A V2board node server based on Xray core, modified from XrayR.<br>
一个基于 Xray core 的 V2board 节点服务端，修改自 XrayR，支持 Vmess/Vless、Trojan、Shadowsocks 等 Xray 能力。

**注意： 本项目需要搭配[修改版V2board](https://github.com/wyx2685/v2board)**

## 特点

* 永久开源且免费。
* xray-only 版本支持 Xray core 提供的 Vmess/Vless、Trojan、Shadowsocks 等协议能力。
* 支持Vless和XTLS等新特性。
* 支持单实例对接多节点，无需重复启动。
* 支持限制在线IP。
* 支持限制Tcp连接数。
* 支持节点端口级别、用户级别限速。
* 配置简单明了。
* 修改配置自动重启实例。
* xray-only：仅内置 Xray core，不再编译 sing-box / hysteria2 独立 core。
* 保留 V2Board API 对接、用户同步、流量上报、在线 IP/在线用户上报等节点逻辑。

## 功能介绍

| 功能 | vmess/vless | trojan | shadowsocks |
|------|-------------|--------|-------------|
| 自动申请 tls 证书 | √ | √ | √ |
| 自动续签 tls 证书 | √ | √ | √ |
| 在线人数统计 | √ | √ | √ |
| 审计规则 | √ | √ | √ |
| 自定义 DNS | √ | √ | √ |
| 在线 IP 数限制 | √ | √ | √ |
| 连接数限制 | √ | √ | √ |
| 跨节点 IP 数限制 | √ | √ | √ |
| 按照用户限速 | √ | √ | √ |
| 动态限速(未测试) | √ | √ | √ |

## TODO

- [ ] 重新实现动态限速
- [ ] 完善使用文档

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/wyx2685/V2bX-script/master/install.sh && bash install.sh
```

### 手动安装

[手动安装教程](https://v2bx.v-50.me/v2bx/v2bx-xia-zai-he-an-zhuang/install/manual)

## 构建
``` bash
# xray-only 版本默认只包含 xray core；tags 仅用于启用 Xray/ACME 相关能力
go build -v -o build_assets/V2bX -tags "with_quic with_grpc with_utls with_wireguard with_acme with_gvisor" -trimpath -ldflags "-X 'github.com/InazumaV/V2bX/cmd.version=$version' -s -w -buildid="
```

## 配置文件及详细使用教程

[详细使用教程](https://v2bx.v-50.me/)

## 免责声明

* 此项目用于本人自用，因此本人不能保证向后兼容性。
* 由于本人能力有限，不能保证所有功能的可用性，如果出现问题请在Issues反馈。
* 本人不对任何人使用本项目造成的任何后果承担责任。
* 本人比较多变，因此本项目可能会随想法或思路的变动随性更改项目结构或大规模重构代码，若不能接受请勿使用。

## 赞助

[赞助链接](https://v-50.me/)

## Thanks

* [Project X](https://github.com/XTLS/)
* [V2Fly](https://github.com/v2fly)
* [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
* [Air-Universe](https://github.com/crossfw/Air-Universe)
* [XrayR](https://github.com/XrayR/XrayR)

## Stars 增长记录

[![Stargazers over time](https://starchart.cc/wyx2685/V2bX.svg)](https://starchart.cc/wyx2685/V2bX)
