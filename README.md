# tools
Set of several tools used by all or most projects at Friends fo Go

1. [portscan](#portscan)

## portscan

A fast and cool port scanner.

```shell script
╰─ portscan

A fast and cool port scanner built with love by Friends of Go.
Complete documentation is available at https://github.com/friendsofgo/tools.

Usage:
  portscan [command]

Available Commands:
  help        Help about any command
  range       Does a range port scan (i.e. 0-1024)
  single      Does a single port scan (i.e. 8080)

Flags:
  -f, --file string   results file (stdin by default)
      --fmt string    results format, allowed: plain, json, yaml (default "plain")
  -h, --help          help for portscan

Use "portscan [command] --help" for more information about a command.
```

### Usage examples

**Scan a single localhost TCP port:**

```shell script
╰─ portscan single 127.0.0.1 tcp 80
Protocol: tcp    Port: 80    State: closed
```

**Scan a range of localhost UDP ports:**

```shell script
╰─ portscan single 127.0.0.1 tcp 80 90
Protocol: udp    Port: 81    State: closed
Protocol: udp    Port: 90    State: closed
Protocol: udp    Port: 89    State: closed
Protocol: udp    Port: 83    State: closed
Protocol: udp    Port: 84    State: closed
Protocol: udp    Port: 80    State: closed
Protocol: udp    Port: 85    State: closed
Protocol: udp    Port: 86    State: closed
Protocol: udp    Port: 82    State: closed
Protocol: udp    Port: 87    State: closed
Protocol: udp    Port: 88    State: closed
```

**Dump range scan results into a JSON file:**

```shell script
╰─ portscan range 127.0.0.1 tcp 80 83 -f out.json --fmt json
```

Results file:

```json
[{
	"protocol": "tcp",
	"port": 80,
	"state": "closed"
}, {
	"protocol": "tcp",
	"port": 83,
	"state": "closed"
}, {
	"protocol": "tcp",
	"port": 81,
	"state": "closed"
}, {
	"protocol": "tcp",
	"port": 82,
	"state": "closed"
}]
```

## TODOs / Ideas

- [ ] Set up GoReleaser for binaries generation.
- [ ] Add unitary & integration coverage.
- [ ] Use only the root command for the `portscan` tool.