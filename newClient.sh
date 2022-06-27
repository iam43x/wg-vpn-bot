#!/bin/sh

#Description
# $1 - client_name
# $2 - ip_address
# $3 - wg_server.conf 

mkdir clients/$1
wg genkey > clients/$1/privatekey
cat clients/$1/privatekey | wg pubkey > clients/$1/publickey

echo "\n[Peer]\nPublicKey = $(cat clients/$1/publickey)\nAllowedIPs = $2\n" >> $3
echo "\n[Interface]\nPrivateKey = $(cat clients/$1/privatekey)\nAddress = $2\nDNS = 8.8.8.8\n\n[Peer]\nPublicKey = $(cat server/publickey)\nEndpoint = vpn.iam43x.online:51831\nAllowedIPs = 0.0.0.0/0\nPersistentKeepalive = 20\n" > clients/$1/wg_client.conf

systemctl restart wg-quick@wgvpn.service

echo "
########################################################################
\033[0;32mUser=$1 with ip_address=$2 successful added!\033[0m
########################################################################
"

exit 0