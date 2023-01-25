#!/usr/bin/env bash

set -euo pipefail

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/results"

echo "start initialize..."
tc qdisc del dev enp6s0 root || true
echo "initialize done"

for ((i = 0; i <= 60; i += 5)); do
	for ((j = 0; j <= 500; j += 100)); do
		packet_loss=$(printf "%03d\n" "${i}")
		if ((i != 0)); then
			echo "packet_loss: ${i}%"
			tc qdisc add dev enp6s0 root netem loss "${i}%"
		fi

		ping_ms=$(printf "%03d\n" "${j}")
		if ((j != 0)); then
			echo "ping_ms: ${j}}ms"
			tc qdisc add dev enp6s0 root netem delay "${j}ms"
		fi

		echo "packet_loss: ${packet_loss}%, ping_ms: ${ping_ms}ms"
		go run "${CURDIR}/main.go" --count 100 --format csv --http3 "https://server:18000" \
			>"${RESULT_DIR}/packet_loss_${packet_loss}%-ping_${ping_ms}ms.csv"

		# パケロスも遅延もない時はエラーが出る為何もしない
		if ((i != 0 && j != 0)); then
			tc qdisc del dev enp6s0 root
		fi
	done
done
