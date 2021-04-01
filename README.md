### First setup

- Create keys

```
openssl genrsa 4096 | openssl pkcs8 -topk8 -nocrypt > privateKey.pem
```
