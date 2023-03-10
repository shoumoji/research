# 卒業研究

## テーマ

Function as a Service における HTTP3/QUIC のコスト削減効果について

## 前提条件

- プロトコル
  - HTTP/2 は Go 標準ライブラリの net/http を使用
  - QUICプロトコルはサードパーティライブラリの [quic-go](<https://github.com/quic-go/quic-go>) を使用
- TLSは両方ともTLS1.3を使用
  - cipher suites は TLS_AES_128_GCM_SHA256 を使用
    - 参照: <https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml>
  - 楕円曲線は x25519 を使用
- Linuxではデフォルトの UDP receive buffer size が小さすぎるため、2.5MBに拡大する
  - 参考 <https://github.com/quic-go/quic-go/wiki/UDP-Receive-Buffer-Size>
  - サーバ側は送信側も同じ大きさに拡大しておく
    - `sudo sysctl -w net.core.rmem_max=2500000 && sudo sysctl -w net.core.wmem_max=2500000`

## 使い方

```bash
make start -j 2 # start HTTP/2 and HTTP/3 server
```

```bash
cd http-client
./start-http2.sh # start HTTP/2 client
./start-http3.sh # start HTTP/3 client
```
