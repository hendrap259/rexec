# Remote Exec
Simple program to exec command remotely in multiple server using ssh.




### Download

#### using installer
if you use MacOS, you can use brew
```bash
https://brew.sh
```
then after brew installed, please install wget
```bash
brew install wget
```
then you can directly install this repository
```bash
wget https://github.com/hendrap259/rexec/raw/master/rexec
```

#### using git clone
```bash
git clone https://github.com/hendrap259/rexec
go mod init github.com/hendrap259/rexec
go mod tidy
go build
```
then you can move the output binary into your /usr/local/bin
```bash
mv rexec /usr/local/bin/rexec
```

### Usage
```bash
rexec [-e | -h <hosts>|-g <group>] <command>

# example of executing command remotely

# single server
rexec -h root@192.168.100.160 tail -f /var/log/nginx/access.log

# multiple server
rexec -h root@192.168.100.160,root@192.168.100.161 zgrep 500 /var/log/nginx/access.log.1.gz

# using file config
rexec grep 500 /var/log/nginx/access.log

# edit file config
rexec -e
```

### Config
Put host config to be line separated. example:
```
 root@192.168.0.123
 root@192.168.0.124
```
or you can group it. example:
```
 [server1]
 root@192.168.0.1
 root@192.168.0.2

 [other-server]
 root@192.168.0.3
 root@192.168.0.4
```
### Thanks
- https://github.com/alileza/rtail
