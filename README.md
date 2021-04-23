# Golang Mongodb Jwt Auth Example Using Echo
> Golang Mongodb Rest Api Example Using Echo
>
<img src="https://github.com/susimsek/golang-echo-mongodb-jwt-auth-example/blob/main/images/golang-echo-mongodb-jwt-auth-example.png" alt="Golang Mongodb Jwt Auth Example Using Echo" width="100%" height="100%"/> 

## Prerequisites

* Golang 1.16.x
* Docker 19.03+
* Docker Compose 1.25+

## Installation

```sh
docker-compose up -d 
```

## Installation Using Vagrant

<img src="https://github.com/susimsek/golang-echo-mongodb-jwt-auth-example/blob/main/images/vagrant-installation.png" alt="Golang Vagrant Installation" width="100%" height="100%"/> 

### Prerequisites

* Vagrant 2.2+
* Virtualbox or Hyperv

```sh
vagrant up
```

```sh
vagrant ssh
```

```sh
cd vagrant/setup
```

```sh
sudo chmod u+x *.sh
```

```sh
./install-prereqs.sh
```

```sh
exit
```

```sh
vagrant ssh
```

```sh
docker-compose up -d
```

You can access the Golang Swagger from the following url.

http://localhost:9000/api

## Used Technologies

* Golang 1.16.3
* Mongodb 4.4
* Echo 
* Content Negotiation Support(Xml,Json Support)
* Pagination Support
* Mongo Driver  
* Mongo Go Pagination
* Echo Swagger
* Swag
* Validator
* Crypto
* Jwt Go  
* Air
* Template
* Sys
* Go Urn  
* Go Colorable
* Universal Translator



## Golang Swagger

> You can access the Golang Swagger from the following url.

http://localhost:9000/api

<img src="https://github.com/susimsek/golang-echo-mongodb-jwt-auth-example/blob/main/images/golang-swagger.png" alt="Golang Swagger" width="100%" height="100%"/> 

## Development

Please note that it requires Go 1.16+ since I use `go mod` to manage dependencies.

```bash

# 1. install dependencies
go mod download

# 2. For less typing, you could add alias air='~/.air' to your .bashrc or .zshrc.
alias air='~/.air'

# 3. run it
air
```


