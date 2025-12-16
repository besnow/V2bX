# V2bX

[![](https://img.shields.io/badge/TgChat-UnOfficialV2Board%E4%BA%A4%E6%B5%81%E7%BE%A4-green)](https://t.me/unofficialV2board)
[![](https://img.shields.io/badge/TgChat-YuzukiProjects%E4%BA%A4%E6%B5%81%E7%BE%A4-blue)](https://t.me/YuzukiProjects)

A V2board node server with an xray-first control plane plus sing protocol extensions, modified from XrayR.
一个以 Xray 为主、借助 sing 协议扩展的 V2board 节点服务端，修改自 XrayR，支持 V2ray、Trojan、Shadowsocks 等协议。

**注意： 本项目需要搭配[修改版V2board](https://github.com/wyx2685/v2board)**

## 特点

* 永久开源且免费。
* 支持Vmess/Vless, Trojan， Shadowsocks, Hysteria1/2 多种协议。
* 支持Vless和XTLS等新特性。
* 支持单实例对接多节点，无需重复启动。
* 支持限制在线IP。
* 支持限制Tcp连接数。
* 支持节点端口级别、用户级别限速。
* 配置简单明了。
* 修改配置自动重启实例。
* 双引擎但单管理平面：常规协议优先走 Xray，Hysteria/Hysteria2/TUIC/AnyTLS 自动走 sing。
* 支持条件编译，可仅编译需要的内核。

## 功能介绍

| 功能        | v2ray | trojan | shadowsocks | hysteria1/2 |
|-----------|-------|--------|-------------|----------|
| 自动申请tls证书 | √     | √      | √           | √        |
| 自动续签tls证书 | √     | √      | √           | √        |
| 在线人数统计    | √     | √      | √           | √        |
| 审计规则      | √     | √      | √           | √         |
| 自定义DNS    | √     | √      | √           | √        |
| 在线IP数限制   | √     | √      | √           | √        |
| 连接数限制     | √     | √      | √           | √         |
| 跨节点IP数限制  |√      |√       |√            |√          |
| 按照用户限速    | √     | √      | √           | √         |
| 动态限速(未测试) | √     | √      | √           | √         |

## TODO

- [ ] 重新实现动态限速
- [ ] 完善使用文档

## 引擎与协议路由

* **Xray 核心（优先）**：Vmess / Vless / Trojan / Shadowsocks 等常规协议。
* **sing-box_mod 协议引擎**：Hysteria / Hysteria2 / TUIC / AnyTLS 由 sing 自动承载。
* 配置项 `Core` 默认为 `auto`（等价 `xray_prefer`），无需为 sing/hysteria2 手工指定；保留对 `xray`/`sing` 旧值的兼容。

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/wyx2685/V2bX-script/master/install.sh && bash install.sh
```

### 手动安装

[手动安装教程](https://v2bx.v-50.me/v2bx/v2bx-xia-zai-he-an-zhuang/install/manual)

## 构建
``` bash
git submodule update --init --recursive
go build -v -o build_assets/V2bX -trimpath -ldflags "-X 'github.com/InazumaV/V2bX/cmd.version=$version' -s -w -buildid="

构建时无需再指定 sing/hysteria2 等 core 标签，默认即可同时编译 Xray 与 sing 协议扩展。
```

## 配置文件及详细使用教程

[详细使用教程](https://v2bx.v-50.me/)

## 免责声明

* 此项目用于本人自用，因此本人不能保证向后兼容性。
* 由于本人能力有限，不能保证所有功能的可用性，如果出现问题请在Issues反馈。
* 本人不对任何人使用本项目造成的任何后果承担责任。
* 本人比较多变，因此本项目可能会随想法或思路的变动随性更改项目结构或大规模重构代码，若不能接受请勿使用。

## 协议实现与许可证提示

项目内集成的 sing-box_mod 源自 sing-box（GPL 系列许可证）。请确保分发或再发行时遵守其许可要求，并在使用中关注相应的第三方许可证声明。

## Third-party licenses

* **sing-box / sing-box_mod**：源自 GPL 系列许可。在分发包含本组件的二进制或镜像时，请保留原始版权声明，并根据 GPL 要求提供相应源码、构建脚本和许可证文本。
* **Xray-core (besnow 版本)**：依照上游许可证分发（请参阅上游仓库）。

## 赞助

[赞助链接](https://v-50.me/)

## Thanks

* [Project X](https://github.com/XTLS/)
* [V2Fly](https://github.com/v2fly)
* [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
* [Air-Universe](https://github.com/crossfw/Air-Universe)
* [XrayR](https://github.com/XrayR/XrayR)
* [sing-box](https://github.com/SagerNet/sing-box)

## Stars 增长记录

[![Stargazers over time](https://starchart.cc/wyx2685/V2bX.svg)](https://starchart.cc/wyx2685/V2bX)
