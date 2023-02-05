# 卒業研究

## テーマ

Function as a Service における HTTP3/QUIC のコスト削減効果について

## 前提条件

- TLSは両方ともTLS1.3を使用
  - cipher suites は同じにする
    - 参照: <https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml>
    - http2
      - TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
    - http3
      - TLS_AES_128_GCM_SHA256
- QUICは <https://github.com/quic-go/quic-go> を使用
- LinuxではUDPのreceive buffer sizeのデフォルトが小さすぎるため、2.5MBに拡大する
  - <https://github.com/quic-go/quic-go/wiki/UDP-Receive-Buffer-Size>

## 使い方

```bash
# start HTTP/2 and HTTP/3 server
make start -j 2
```

```bash
cd http-client
./start-http2.sh
./start-http3.sh
```
