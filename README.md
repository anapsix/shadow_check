## shadow_check

Simple utility to verify user's password, comparing it with salted SHA512 (like in `/etc/shadow`)

### Install
```
go get github.com/anapsix/shadow_check
```

### Build

For example:
```
env GOOS=linux GOARCH=amd64 go build -o pkgs/shadow_check-linux_amd64
env GOOS=darwin GOARCH=amd64 go build -o pkgs/shadow_check-darwin_amd64
```

Reduce binary size with some UPX packing:
```
# install UPX with `brew install upx`
upx pkgs/*
```

### Usage

Use it directly, like so:
```
echo 'testuser:$6$Su1Gj94CfnCd78lA$BCxPneWhGybzbKM2f6E3O0TDOdMBqN6pCRJUAUWIB6BvzUSoAC7Ccxt9er5P6VKO1H1JpbG2U9eRgJW5QCvgG0' > /tmp/testshadow
shadow_check -user testuser -password mygoodpassword -shadow /tmp/testshadow
shadow_check -user testuser -password mybadpassword -shadow /tmp/testshadow
```

Could be used for OpenVPN verification of user's password.  
Add the following into OpenVPN's `server.conf`:
```
auth-user-pass-verify check_user.sh via-env
script-security 3  # required to pass user's password to script
```

The content of the `check_user.sh` is as follows:
```bash
#!/bin/bash

username=${username:-$1}
password=${password:-$2}
shadow_file='/etc/openvpn/vpn_users'

./shadow_check -user "${username}" -password "${password}" -shadow "${shadow_file}"
```