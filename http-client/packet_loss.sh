#!/usr/bin/env bash

set -euo pipefail

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/results"

tc qdisc del dev enp6s0 root

for ((i = 0; i <= 100; i += 5)); do
	for ((j = 0; j <= 500; j += 100)); do
		packet_loss=$(printf "%03d\n" "${i}")
		if ((i != 0)); then
			tc qdisc add dev enp6s0 root netem loss "${i}%"
		fi

		ping_ms=$(printf "%03d\n" "${j}")
		if ((j != 0)); then
			echo "j: $j"
			tc qdisc add dev enp6s0 root netem delay "${j}ms"
		fi

		go run "${CURDIR}/main.go" --count 100 --format csv --http3 "https://server:18000" \
			>"${RESULT_DIR}/packet_loss_${packet_loss}}%-ping_${ping_ms}ms.csv"

		tc qdisc del dev enp6s0 root
	done
done
