# Choose

Project space management tool utilizing `tmux` and Go.

## Installation

You must have `tmux` installed.

### From Source

Clone the repository and build:

```bash
git clone https://github.com/ericbaukhages/choose.git
cd choose
go build -o choose-tool
touch ~/.tmux.sessions.log
```

Note: The binary must be built with a custom name (e.g., `choose-tool`) due to a directory naming conflict.

### Using go install

```bash
go install github.com/ericbaukhages/choose@latest
touch ~/.tmux.sessions.log
```

## Usage

```
Usage:
  choose [flags]
  choose [command]

Available Commands:
  help        Help about any command
  new         Creates a new tmux session
  open        Open an existing project if possible

Flags:
      --config string   config file (default is $HOME/.choose.yaml)
  -h, --help            help for choose
  -t, --toggle          Help message for toggle

Use "choose [command] --help" for more information about a command.
```

Using `choose` with no options will run the select UI.
