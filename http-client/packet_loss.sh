#!/usr/bin/env bash

for i in {0..99}; do
	file=$(printf "%02d\n" "${i}")
	tc qdisc add dev enp6s0 root netem loss "${i}%"
	go run main.go --count 100 --format csv --http3 https://server:18000 >"packet_loss_${file}%.csv"
	tc qdisc del dev enp6s0 root
done
