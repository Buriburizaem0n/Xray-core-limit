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
