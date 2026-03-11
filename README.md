# Xray limit: Rate Limiting & IP Restriction

[Xray原文档](https://github.com/XTLS/Xray-core/blob/main/README.md)

本文档说明该fork Xray limit 的用户限速和并发 IP 限制功能。

## 功能概述

| 功能 | 配置项 | 说明 |
|------|--------|------|
| 上传限速 | `uplinkSpeedLimit` | 限制用户上传速度 (bytes/s) |
| 下载限速 | `downlinkSpeedLimit` | 限制用户下载速度 (bytes/s) |
| 并发 IP 限制 | `maxConcurrentIPs` | 限制用户同时在线的 IP 数量 |

## 前置条件

1. **用户必须配置 `email` 字段** - 用于唯一标识用户
2. **启用统计功能** - 需要 `statsUserUplink` 和/或 `statsUserDownlink`

## 配置示例

```json
{
  "policy": {
    "levels": {
      "0": {
        "statsUserUplink": true,
        "statsUserDownlink": true,
        "uplinkSpeedLimit": 1250000,
        "downlinkSpeedLimit": 62500000,
        "maxConcurrentIPs": 3
      }
    }
  },
  "inbounds": [
    {
      "protocol": "vless",
      "settings": {
        "clients": [
          {
            "id": "your-uuid",
            "email": "user@example.com",
            "level": 0
          }
        ]
      }
    }
  ]
}
```

### [Collect a Project X NFT to support the development of Project X!](https://opensea.io/item/ethereum/0x5ee362866001613093361eb8569d59c4141b76d1/1)

[<img alt="Project X NFT" width="150px" src="https://raw2.seadn.io/ethereum/0x5ee362866001613093361eb8569d59c4141b76d1/7fa9ce900fb39b44226348db330e32/8b7fa9ce900fb39b44226348db330e32.svg" />](https://opensea.io/item/ethereum/0x5ee362866001613093361eb8569d59c4141b76d1/1)

- **TRX(Tron)/USDT/USDC: `TNrDh5VSfwd4RPrwsohr6poyNTfFefNYan`**
- **TON: `UQApeV-u2gm43aC1uP76xAC1m6vCylstaN1gpfBmre_5IyTH`**
- **BTC: `1JpqcziZZuqv3QQJhZGNGBVdCBrGgkL6cT`**
- **XMR: `4ABHQZ3yJZkBnLoqiKvb3f8eqUnX4iMPb6wdant5ZLGQELctcerceSGEfJnoCk6nnyRZm73wrwSgvZ2WmjYLng6R7sR67nq`**
- **SOL/USDT/USDC: `3x5NuXHzB5APG6vRinPZcsUv5ukWUY1tBGRSJiEJWtZa`**
- **ETH/USDT/USDC: `0xDc3Fe44F0f25D13CACb1C4896CD0D321df3146Ee`**
- **Project X NFT: https://opensea.io/item/ethereum/0x5ee362866001613093361eb8569d59c4141b76d1/1**
- **VLESS NFT: https://opensea.io/collection/vless**
- **REALITY NFT: https://opensea.io/item/ethereum/0x5ee362866001613093361eb8569d59c4141b76d1/2**
- **Related links: [VLESS Post-Quantum Encryption](https://github.com/XTLS/Xray-core/pull/5067), [XHTTP: Beyond REALITY](https://github.com/XTLS/Xray-core/discussions/4113), [Announcement of NFTs by Project X](https://github.com/XTLS/Xray-core/discussions/3633)**

## License

[Mozilla Public License Version 2.0](https://github.com/XTLS/Xray-core/blob/main/LICENSE)

## Documentation

[Project X Official Website](https://xtls.github.io)

## Telegram

[Project X](https://t.me/projectXray)

[Project X Channel](https://t.me/projectXtls)

[Project VLESS](https://t.me/projectVless) (Русский)

[Project XHTTP](https://t.me/projectXhttp) (Persian)

## Installation

- Linux Script
  - [XTLS/Xray-install](https://github.com/XTLS/Xray-install) (**Official**)
  - [tempest](https://github.com/team-cloudchaser/tempest) (supports [`systemd`](https://systemd.io) and [OpenRC](https://github.com/OpenRC/openrc); Linux-only)
- Docker
  - [ghcr.io/xtls/xray-core](https://ghcr.io/xtls/xray-core) (**Official**)
  - [teddysun/xray](https://hub.docker.com/r/teddysun/xray)
  - [wulabing/xray_docker](https://github.com/wulabing/xray_docker)
- Web Panel
  - [Remnawave](https://github.com/remnawave/panel)
  - [3X-UI](https://github.com/MHSanaei/3x-ui)
  - [PasarGuard](https://github.com/PasarGuard/panel)
  - [Xray-UI](https://github.com/qist/xray-ui)
  - [X-Panel](https://github.com/xeefei/X-Panel)
  - [Marzban](https://github.com/Gozargah/Marzban)
  - [Hiddify](https://github.com/hiddify/Hiddify-Manager)
  - [TX-UI](https://github.com/AghayeCoder/tx-ui)
- One Click
  - [Xray-REALITY](https://github.com/zxcvos/Xray-script), [xray-reality](https://github.com/sajjaddg/xray-reality), [reality-ezpz](https://github.com/aleskxyz/reality-ezpz)
  - [Xray_bash_onekey](https://github.com/hello-yunshu/Xray_bash_onekey), [XTool](https://github.com/LordPenguin666/XTool), [VPainLess](https://github.com/vpainless/vpainless)
  - [v2ray-agent](https://github.com/mack-a/v2ray-agent), [Xray_onekey](https://github.com/wulabing/Xray_onekey), [ProxySU](https://github.com/proxysu/ProxySU)
- Magisk
  - [NetProxy-Magisk](https://github.com/Fanju6/NetProxy-Magisk)
  - [Xray4Magisk](https://github.com/Asterisk4Magisk/Xray4Magisk)
  - [Xray_For_Magisk](https://github.com/E7KMbb/Xray_For_Magisk)
- Homebrew
  - `brew install xray`

## Usage

- Example
  - [VLESS-XTLS-uTLS-REALITY](https://github.com/XTLS/REALITY#readme)
  - [VLESS-TCP-XTLS-Vision](https://github.com/XTLS/Xray-examples/tree/main/VLESS-TCP-XTLS-Vision)
  - [All-in-One-fallbacks-Nginx](https://github.com/XTLS/Xray-examples/tree/main/All-in-One-fallbacks-Nginx)
- Xray-examples
  - [XTLS/Xray-examples](https://github.com/XTLS/Xray-examples)
  - [chika0801/Xray-examples](https://github.com/chika0801/Xray-examples)
  - [lxhao61/integrated-examples](https://github.com/lxhao61/integrated-examples)
- Tutorial
  - [XTLS Vision](https://github.com/chika0801/Xray-install)
  - [REALITY (English)](https://cscot.pages.dev/2023/03/02/Xray-REALITY-tutorial/)
  - [XTLS-Iran-Reality (English)](https://github.com/SasukeFreestyle/XTLS-Iran-Reality)
  - [Xray REALITY with 'steal oneself' (English)](https://computerscot.github.io/vless-xtls-utls-reality-steal-oneself.html)
  - [Xray with WireGuard inbound (English)](https://g800.pages.dev/wireguard)

## GUI Clients

- OpenWrt
  - [PassWall](https://github.com/Openwrt-Passwall/openwrt-passwall), [PassWall 2](https://github.com/Openwrt-Passwall/openwrt-passwall2)
  - [ShadowSocksR Plus+](https://github.com/fw876/helloworld)
  - [luci-app-xray](https://github.com/yichya/luci-app-xray) ([openwrt-xray](https://github.com/yichya/openwrt-xray))
- Asuswrt-Merlin
  - [XRAYUI](https://github.com/DanielLavrushin/asuswrt-merlin-xrayui)
  - [fancyss](https://github.com/hq450/fancyss)
- Windows
  - [v2rayN](https://github.com/2dust/v2rayN)
  - [Furious](https://github.com/LorenEteval/Furious)
  - [Invisible Man - Xray](https://github.com/InvisibleManVPN/InvisibleMan-XRayClient)
  - [AnyPortal](https://github.com/AnyPortal/AnyPortal)
  - [GenyConnect](https://github.com/genyleap/GenyConnect)
- Android
  - [v2rayNG](https://github.com/2dust/v2rayNG)
  - [X-flutter](https://github.com/XTLS/X-flutter)
  - [SaeedDev94/Xray](https://github.com/SaeedDev94/Xray)
  - [SimpleXray](https://github.com/lhear/SimpleXray)
  - [XrayFA](https://github.com/Q7DF1/XrayFA)
  - [AnyPortal](https://github.com/AnyPortal/AnyPortal)
  - [NetProxy-Magisk](https://github.com/Fanju6/NetProxy-Magisk)
- iOS & macOS arm64 & tvOS
  - [Happ](https://apps.apple.com/app/happ-proxy-utility/id6504287215) | [Happ RU](https://apps.apple.com/ru/app/happ-proxy-utility-plus/id6746188973) | [Happ tvOS](https://apps.apple.com/us/app/happ-proxy-utility-for-tv/id6748297274)
  - [Streisand](https://apps.apple.com/app/streisand/id6450534064)
  - [OneXray](https://github.com/OneXray/OneXray)
- macOS arm64 & x64
  - [Happ](https://apps.apple.com/app/happ-proxy-utility/id6504287215) | [Happ RU](https://apps.apple.com/ru/app/happ-proxy-utility-plus/id6746188973)
  - [V2rayU](https://github.com/yanue/V2rayU)
  - [V2RayXS](https://github.com/tzmax/V2RayXS)
  - [Furious](https://github.com/LorenEteval/Furious)
  - [OneXray](https://github.com/OneXray/OneXray)
  - [GoXRay](https://github.com/goxray/desktop)
  - [AnyPortal](https://github.com/AnyPortal/AnyPortal)
  - [v2rayN](https://github.com/2dust/v2rayN)
  - [GenyConnect](https://github.com/genyleap/GenyConnect)
- Linux
  - [v2rayA](https://github.com/v2rayA/v2rayA)
  - [Furious](https://github.com/LorenEteval/Furious)
  - [GorzRay](https://github.com/ketetefid/GorzRay)
  - [GoXRay](https://github.com/goxray/desktop)
  - [AnyPortal](https://github.com/AnyPortal/AnyPortal)
  - [v2rayN](https://github.com/2dust/v2rayN)
  - [GenyConnect](https://github.com/genyleap/GenyConnect)

## Others that support VLESS, XTLS, REALITY, XUDP, PLUX...

- iOS & macOS arm64 & tvOS
  - [Shadowrocket](https://apps.apple.com/app/shadowrocket/id932747118)
  - [Loon](https://apps.apple.com/us/app/loon/id1373567447)
  - [Egern](https://apps.apple.com/us/app/egern/id1616105820)
  - [Quantumult X](https://apps.apple.com/us/app/quantumult-x/id1443988620)
- Xray Tools
  - [xray-knife](https://github.com/lilendian0x00/xray-knife)
  - [xray-checker](https://github.com/kutovoys/xray-checker)
- Xray Wrapper
  - [XTLS/libXray](https://github.com/XTLS/libXray)
  - [xtls-sdk](https://github.com/remnawave/xtls-sdk)
  - [xtlsapi](https://github.com/hiddify/xtlsapi)
  - [AndroidLibXrayLite](https://github.com/2dust/AndroidLibXrayLite)
  - [Xray-core-python](https://github.com/LorenEteval/Xray-core-python)
  - [xray-api](https://github.com/XVGuardian/xray-api)
- [XrayR](https://github.com/XrayR-project/XrayR)
  - [XrayR-release](https://github.com/XrayR-project/XrayR-release)
  - [XrayR-V2Board](https://github.com/missuo/XrayR-V2Board)
- Cores
  - [Amnezia VPN](https://github.com/amnezia-vpn)
  - [mihomo](https://github.com/MetaCubeX/mihomo)
  - [sing-box](https://github.com/SagerNet/sing-box)

## Contributing

[Code of Conduct](https://github.com/XTLS/Xray-core/blob/main/CODE_OF_CONDUCT.md)

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/XTLS/Xray-core)

## Credits

- [Xray-core v1.0.0](https://github.com/XTLS/Xray-core/releases/tag/v1.0.0) was forked from [v2fly-core 9a03cc5](https://github.com/v2fly/v2ray-core/commit/9a03cc5c98d04cc28320fcee26dbc236b3291256), and we have made & accumulated a huge number of enhancements over time, check [the release notes for each version](https://github.com/XTLS/Xray-core/releases).
- For third-party projects used in [Xray-core](https://github.com/XTLS/Xray-core), check your local or [the latest go.mod](https://github.com/XTLS/Xray-core/blob/main/go.mod).

## One-line Compilation

### Windows (PowerShell)

```powershell
$env:CGO_ENABLED=0
go build -o xray.exe -trimpath -buildvcs=false -ldflags="-s -w -buildid=" -v ./main

```

## 速度单位换算

| 配置值 | 实际速度 |
|--------|----------|
| 125000 | 1 Mbps |
| 1250000 | 10 Mbps |
| 12500000 | 100 Mbps |
| 62500000 | 500 Mbps |

计算公式：`bytes/s = Mbps × 125000`

## 注意事项

1. **XTLS Vision 兼容**：当限速启用时，Splice 会自动禁用以确保限速生效
2. **并发 IP 计算**：同一 IP 的多个连接只计为 1 个 IP，IP 在 20 秒无活动后自动移除
3. **客户端兼容**：某些客户端（如使用 tproxy 模式的 V2rayA）可能使用内核转发绕过限速

## 常见问题

**Q: 配置了限速但不生效？**
- 检查用户是否配置了 `email`
- 检查用户 `level` 是否匹配 policy 配置
- 确保 `statsUserUplink` / `statsUserDownlink` 已启用

**Q: 如何只限制上传不限制下载？**
- 只设置 `uplinkSpeedLimit`，不设置 `downlinkSpeedLimit`（或设为 0）
