# timeout-exec

Simple CLI for running commands, with a timeout

## Example usage

```bash
$ timeout-exec run --timeout=2 echo hello world, I do not timeout
hello world, I do not timeout

$ timeout-exec run --timeout=5 sleep 6
Process timed out

$ timeout-exec run --timeout=5 -- bash -c 'echo I preserve the exit code too && exit 3'
I preserve the exit code too
$ echo $?
3
```

## License

Licensed under the [MIT License](LICENSE)