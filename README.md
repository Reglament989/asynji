# Asynji - backend for [Exyji](https://github.com/Reglament989/exyji.git)

* [About]()
  * [Look how it works](#Look-how-it-works)
* [Setup own server](#Im-wonna-setup-own-server)
  * [First setup](#First-setup)
    * [Clone this repo](#Clone-this-repo)
    * [Create private key](#Create-private-key)
  * [Manual](#Manual)
    * [Build](#Build)
    * [Run](#Run)
  * [Docker](#Docker)
    * [Build images](#Build-images)
    * [Run docker](#Run-docker)


### Look how it works
[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=Asynji%20refs&uri=https%3A%2F%2Fraw.githubusercontent.com%2FReglament989%2Fasynji%2Fmain%2Fdocs%2Finsomnia-docs.json)


## Im wonna setup own server

### First setup

#### Clone this repo
```
git clone https://github.com/Reglament989/asynji.git
```

#### Create private key
~~please backup it~~
```
openssl genrsa 4096 | openssl pkcs8 -topk8 -nocrypt > privateKey.pem
```

[Docker-Compose]()

### Manual

You need to be run mongodb, redis and configured it via .env

#### Build
```
make build
```

#### Run
```
make run
```

### Docker

#### Build images
```
make docker-build
```

#### Run docker
```
make drun
```

