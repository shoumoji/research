#!/usr/bin/env bash

set -euo pipefail

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/results"

echo "start initialize..."
tc qdisc del dev enp6s0 root || true
echo "initialize done"

for ((i = 0; i <= 60; i += 5)); do
	for ((j = 0; j <= 300; j += 100)); do
		packet_loss=$(printf "%03d\n" "${i}")
		echo "packet_loss: ${packet_loss}%, i: ${i}"
		if ((i != 0)); then
			tc qdisc add dev enp6s0 root netem loss "${i}%"
		fi

		ping_ms=$(printf "%03d\n" "${j}")
		echo "ping_ms: ${ping_ms}ms, j: ${j}"
		if ((j != 0)); then
			tc qdisc add dev enp6s0 root netem delay "${j}ms"
		fi

		go run "${CURDIR}/main.go" --count 100 --format csv --http3 "https://server:18000" \
			>"${RESULT_DIR}/packet_loss_${packet_loss}%-ping_${ping_ms}ms.csv"

		# パケロスも遅延もない時はエラーが出る為強制的に成功させる
		tc qdisc del dev enp6s0 root || true
	done
done
