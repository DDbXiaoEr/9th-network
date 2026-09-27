---
title: ssl证书自签名快捷制作
date: 2026-06-27 00:00:00
categories:
- 运维经验记录
- linux
tags:
---

测试环境使用自签名证书就够了，主要记录两种场景：单域名和泛域名（多域名）。

环境说明：OpenSSL 3.0+

## 单域名自签名

`--nodes` 表示私钥不加密，一步生成 `key` + `crt`：

```bash
#!/bin/bash
set -e

execute_date=$(date "+%Y-%m-%d_%H-%M-%S")

if [[ $# -eq 0 ]]; then
    echo -e "CN\nBJ\nBJ\nLocal\nLocal\nwww.example.com\ntest@localhost\n" \
        | openssl req -new --nodes -x509 -keyout "${execute_date}.key" -out "${execute_date}.crt" > /dev/null 2>&1
else
    echo -e "CN\nBJ\nBJ\nLocal\nLocal\n$1\ntest@localhost\n" \
        | openssl req -new --nodes -x509 -keyout "${execute_date}.key" -out "${execute_date}.crt" > /dev/null 2>&1
fi

echo -e "${execute_date}.key\n${execute_date}.crt"
```

用法：

```bash
# 默认域名 www.example.com
./generate.sh

# 自定义域名
./generate.sh www.test.com
```

也可直接命令行：

```bash
echo -e "CN\nBJ\nBJ\nLocal\nLocal\nwww.example.com\ntest@localhost\n" \
    | openssl req -new --nodes -x509 -keyout server.key -out server.crt > /dev/null 2>&1
```

## 泛域名自签名

需要配置文件 `openssl.cnf`：

```ini
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req

[req_distinguished_name]
countryName = CN
countryName_default = CN
stateOrProvinceName = BJ
stateOrProvinceName_default = BJ
localityName = WZQ
localityName_default = WZQ
organizationalUnitName = WXQ
organizationalUnitName = WXQ
commonName = *.DOMAIN.COM
commonName_max = 64

[v3_req]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation,digitalSignature,keyEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = DOMAIN.COM
DNS.2 = *.DOMAIN.COM
```

生成证书：

```bash
openssl genrsa -out server.key 2048
openssl req -new -out server.csr -key server.key -config openssl.cnf
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt -extensions v3_req -extfile openssl.cnf
```

批量脚本 `generate.sh`：

```bash
#!/bin/bash
set -e

declare domain
if [[ $# -eq 0 ]]; then
    domain="example.com"
else
    domain=$1
fi

sed -e "s/DOMAIN.COM/$domain/g" openssl_template.cnf > openssl.cnf
openssl genrsa -out server.key 2048
openssl req -new -out server.csr -key server.key -config openssl.cnf
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt -extensions v3_req -extfile openssl.cnf
rm -rf openssl.cnf
```

用法：

```bash
# 默认 example.com
./generate.sh

# 自定义泛域名
./generate.sh test.com
```
