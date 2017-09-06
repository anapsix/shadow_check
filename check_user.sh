#!/bin/bash

username=${username:-$1}
password=${password:-$2}
shadow_file='/etc/openvpn/vpn_users'

./shadow_check -user "${username}" -password "${password}" -shadow "${shadow_file}"