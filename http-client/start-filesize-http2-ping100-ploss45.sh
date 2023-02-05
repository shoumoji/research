#!/usr/bin/env bash

set -euo pipefail

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/filesize-http2-results"

filesize[0]="1mb"
filesize[1]="100mb"

echo "start initialize..."

tc qdisc del dev enp6s0 root || true

if [ ! -d "${RESULT_DIR}" ]; then
	mkdir "${RESULT_DIR}"
fi

echo "initialize done"

i=45
j=100

packet_loss=$(printf "%03d\n" "${i}")
ping_ms=$(printf "%03d\n" "${j}")

echo "packet_loss: ${packet_loss}%, i: ${i}"
echo "ping_ms: ${ping_ms}ms, j: ${j}"

if ((i != 0 && j != 0)); then
	tc qdisc add dev enp6s0 root netem loss "${i}%" delay "${j}ms"
elif ((i == 0 && j != 0)); then
	tc qdisc add dev enp6s0 root netem delay "${j}ms"
elif ((i != 0 && j == 0)); then
	tc qdisc add dev enp6s0 root netem loss "${i}%"
fi

for fsize in "${filesize[@]}"; do
	go run "${CURDIR}/main.go" --count 100 --format csv --http2 "https://server:18000/${fsize}" \
		>"${RESULT_DIR}/ping_${ping_ms}ms-packet_loss_${packet_loss}%-filesize_${fsize}.csv"
done

# パケロスも遅延もない時はエラーが出るため
if ((i != 0 || j != 0)); then
	tc qdisc del dev enp6s0 root || true
fi
