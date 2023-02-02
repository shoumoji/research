#!/usr/bin/env bash

set -euo pipefail

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/filesize-http2-results"

filesize[0]=""
filesize[1]="1mb"
filesize[2]="10mb"
filesize[3]="100mb"
filesize[4]="1000mb"

echo "start initialize..."

tc qdisc del dev enp6s0 root || true

if [ ! -d "${RESULT_DIR}" ]; then
	mkdir "${RESULT_DIR}"
fi

echo "initialize done"

for ((i = 0; i <= 50; i += 5)); do
	for ((j = 0; j <= 300; j += 100)); do
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
			go run "${CURDIR}/main.go" --count 10 --format csv --http2 "https://server:18000/${fsize}" \
				>"${RESULT_DIR}/ping_${ping_ms}ms-packet_loss_${packet_loss}%-filesize_${fsize}.csv"
		done

		# パケロスも遅延もない時はエラーが出るため
		if ((i != 0 || j != 0)); then
			tc qdisc del dev enp6s0 root || true
		fi
	done
done
