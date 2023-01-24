#!/usr/bin/env bash

CURDIR=$(pwd)
RESULT_DIR="${CURDIR}/results"

for ((i = 5; i <= 100; i += 5)); do
	file_num=$(printf "%02d\n" "${i}")
	tc qdisc add dev enp6s0 root netem loss "1%"
	go run "${CURDIR}/main.go" --count 100 --format csv --http3 "https://server:18000" >"${RESULT_DIR}/packet_loss_${file_num}%.csv"
	tc qdisc del dev enp6s0 root
done
