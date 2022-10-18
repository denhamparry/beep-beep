# beepbeep

Sample Go application to log out `beep beep` until the application recieves a `SIGINT` or `SIGTERM` signal where it dumps out its current state.

> Application to be used as a way to capture a flag at the KubeCon NA 2022 Hack Back CTF talk.

## Build

> Uses [ko](https://github.com/ko-build/ko) to build the application

```sh
ko build ./cmd/app
```

