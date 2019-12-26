# binder

## Usage

```sh
$ cat commands
app : docker logs -f app
db  : docker logs -f db

# Execute commands & show logs
$ cat commands | binder
app| log line of app
 db| log line of db
```