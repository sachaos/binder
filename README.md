# binder

Execute commands from STDOUT, and show logs of those.

## Usage

### Basic

```sh
$ cat commands
app : docker logs -f app
db  : docker logs -f db

# Execute commands & show logs
$ cat commands | binder
app| log line of app
 db| log line of db
```

### See all Docker containers log

```sh
$ docker ps --format '{{.Names}}' | xargs -I {} echo {} ":docker logs --tail=0 -f" {} | binder
```

## Install

### Binary

Go to [release page](https://github.com/sachaos/binder/releases) and download.

or

```shell
$ wget https://github.com/sachaos/binder/releases/download/v0.0.1/binder_darwin_amd64 -O /usr/local/bin/binder
$ chmod +x /usr/local/bin/binder
```

### Manually Build

You need Go 1.13 compiler.

```shell
$ git clone https://github.com/sachaos/binder.git
$ make install
```
