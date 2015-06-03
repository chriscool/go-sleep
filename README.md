# go-sleep sleeps for some milliseconds

This unix tool is a thin wrapper around `time.Sleep()`.
It aims to provide a portable way to sleep for a number of
milliseconds.

### Install

```sh
go install github.com/chriscool/go-sleep/sleep
```

(The extra /sleep is there because go get is stupidly too proscriptive about
package/repository names and I don't yet know how to change the default binary
output name.)

### Usage:

```
> sleep
Usage: sleep <count>
Sleep for <count> milliseconds.
```

### License

MIT
