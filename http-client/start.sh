#!/usr/bin/env bash

set -euo pipefail

######################
count=100    # 計測回数
format="csv" # 出力形式 (csv or json)
filesize[0]="1mb"
filesize[1]="100mb"
# filesize[2]="10mb"
# filesize[3]="1000mb"
######################

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/results"

echo "start initialize..."

tc qdisc del dev enp6s0 root || true

if [ ! -d "${RESULT_DIR}/http2" ]; then
	mkdir -p "${RESULT_DIR}/http2"
fi
if [ ! -d "${RESULT_DIR}/http3" ]; then
	mkdir -p "${RESULT_DIR}/http3"
fi

echo "initialize done"

for ((i = 0; i <= 30; i++)); do
	for ((j = 0; j <= 100; j += 100)); do
		packet_loss=$(printf '%.1f\n' "$(echo "scale=2;${i}/10" | bc)")
		ping_ms=$(printf "%03d\n" "${j}")

		echo "packet_loss: ${packet_loss}%, i: ${i}"
		echo "ping_ms: ${ping_ms}ms, j: ${j}"

		if ((i != 0 && j != 0)); then
			tc qdisc add dev enp6s0 root netem loss "${packet_loss}%" delay "${j}ms"
		elif ((i == 0 && j != 0)); then
			tc qdisc add dev enp6s0 root netem delay "${j}ms"
		elif ((i != 0 && j == 0)); then
			tc qdiscadd dev enp6s0 root netem loss "${packet_loss}%"
		fi

		for fsize in "${filesize[@]}"; do
			go run "${CURDIR}/main.go" --count ${count} --format ${format} --http2 "https://server:18000/${fsize}" \
				>"${RESULT_DIR}/http2/${fsize}/ping_${ping_ms}ms-packet_loss_${packet_loss}%.csv"
			go run "${CURDIR}/main.go" --count ${count} --format ${format} --http3 "https://server:18000/${fsize}" \
				>"${RESULT_DIR}/http3/${fsize}/ping_${ping_ms}ms-packet_loss_${packet_loss}%.csv"
		done

		# パケロスも遅延もない時はエラーが出るため
		if ((i != 0 || j != 0)); then
			tc qdisc del dev enp6s0 root || true
		fi
	done
done
