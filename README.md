# go-monitor

## What?

Monitoring collector tools golang. 

## Running

This project uses a `makefile` to automate some of the build steps needed.

To build the binaries you need to run

```shell
make build
```

To clean compilation artifacts run

```shell
make clean
```

If you invoke `make` the default associated target steps are the equivalent of

```shell
make format build
```

