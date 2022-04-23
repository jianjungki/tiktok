# tiktok
 Go SDK for TikTok global

## Feature

- [x] Authentication
  - [x] Get your auth code url
  - [x] Use auth code to get access_token
  - [x] Refresh your access token
- [ ] Webhook
- [x] Order API
  - [x] Get order list
  - [x] Get order detail
  - [ ] Ship order
- [ ] Fulfillment API
- [ ] Logistics API
- [ ] Product API
- [ ] Shop API
- [ ] Finance API
- [ ] Reverse Order API

## API Call Flow

https://bytedance.feishu.cn/docs/doccnZ15f4vPK4qOOBTKhgc2xNh

## HOWTO Contribuate

Install [`taskfile`](https://taskfile.dev) before you start.

1. Setup Intergration Environment Variables

edit `.env` file in root directory:

```
APPKEY=<app id for tiktok>
APPSECRET=<app secret for tiktok>
AK=<access token>
RK=<refresh token>
OPENID=<seller's openid>
SELLER=<seller's name>
```

```
task tests  # run tests.
```

If env is not setup, all tests will be skipped.