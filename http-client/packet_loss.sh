#!/usr/bin/env bash

for i in {0..99}; do
	tc qdisc add dev enp6s0 root netem loss "${i}%"
	go run main.go --count 100 --format csv >"packet_loss_${i}%.csv"
	tc qdisc del dev enp6s0 root
done
