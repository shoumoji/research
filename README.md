# 卒業研究

## テーマ

TCP上のprotocolであるhttp2 と、UDP上のprotocolであるQUICのFaaS上での性能比較

## 前提条件

- TLSは両方ともTLS1.3を使用
  - cipher suites は同じにする
- サーバから送られるデータは暗号化のシークレットデータのみ
- QUICは github.com/lucas-clemente/quic-go/http3 を使用
- LinuxではUDPのreceive buffer sizeのデフォルトが小さすぎるため、2.5MBに拡大
  - https://github.com/quic-go/quic-go/wiki/UDP-Receive-Buffer-Size
