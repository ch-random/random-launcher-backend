[![Actions Status](https://github.com/ch-random/random-launcher-backend/actions/workflows/ci.yml/badge.svg)](https://github.com/ch-random/random-launcher-backend/actions/workflows/ci.yml)
[![go-version](https://img.shields.io/github/go-mod/go-version/ch-random/random-launcher-backend)](https://github.com/ch-random/random-launcher-backend/blob/master/go.mod)
[![license](https://img.shields.io/badge/license-CC0--1.0-blue)](https://github.com/ch-random/random-launcher-backend/blob/master/LICENSE)

# random-launcher-backend

## Install `fly`

```sh
curl -L https://fly.io/install.sh | sh
```

## Login

```sh
fly auth login
fly logs -a random-launcher
```

## Execute

```sh
go run cmd/app/main.go
```

## Initialize all tables

```sh
go run cmd/app/main.go migrate --drop
```

## Deploy

```sh
# fly wireguard reset
fly deploy
```

## License

これらのコードや文章は [CC0](https://creativecommons.org/publicdomain/zero/1.0/deed.ja) で許諾されています。すなわち、引用元に記載せずに、これらのコードや文章の一部または全部を使用できます。
