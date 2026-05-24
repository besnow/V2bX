# V2bX xray-only fork

这是自用的 V2bX xray-only fork，基于 V2bX / XrayR 节点端维护，用于对接 V2Board API 并运行 Xray core。

**注意：本项目需要搭配[修改版 V2Board](https://github.com/wyx2685/v2board)。**

## 项目定位

- V2bX 是 xray-only fork。
- 这里的 xray-only 仅表示运行核心只使用 Xray core，不代表协议能力收窄。
- 不恢复 V2bX 外部 sing-box / hysteria / hysteria2 core。
- Xray core 使用 [`wyx2685/Xray-core`](https://github.com/wyx2685/Xray-core)，其内部能力（如 anytls、tuic、wireguard、singquic、xhttp、grpc、websocket、reality、routing、DNS、stats、policy 等）按核心实现保留。
- V2bX 外层适配以“尽量实现”为目标，生产可用性以灰度验证结果为准。

## 构建

当前 xray-only 构建不需要额外 core tags：

```bash
mkdir -p build_assets
go build -v -o build_assets/V2bX -trimpath -ldflags "-X 'github.com/InazumaV/V2bX/cmd.version=$version' -s -w -buildid="
```

本分支依赖的 Xray-core 需要 Go 1.26 或更新版本；如果本地 Go 启用了自动工具链，`go build` 会按需下载匹配工具链。

## 安装说明

不要直接使用上游 `wyx2685/V2bX-script` 一键安装脚本安装本 fork；该脚本可能安装到上游原版 V2bX，而不是当前 xray-only fork。

请自行构建本仓库产物，并结合自己的部署方式替换二进制和配置文件。

## TODO

- [ ] 完善本 fork 的部署说明。

## 免责声明

- 此项目为自用 fork，不承诺向后兼容。
- 不添加未经灰度验证的生产可用承诺；实际可用能力以当前代码和所使用的 Xray-core 为准。
- 使用本项目造成的任何后果由使用者自行承担。

## Thanks

- [Project X](https://github.com/XTLS/)
- [V2Fly](https://github.com/v2fly)
- [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
- [Air-Universe](https://github.com/crossfw/Air-Universe)
- [XrayR](https://github.com/XrayR/XrayR)
