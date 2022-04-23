<p align="center">
  <a href="https://github.com/abdfnx/gomo#gh-light-mode-only">
    <img src="./.github/assets/logo.png" alt="Gomo">
  </a>
  <a href="https://github.com/abdfnx/gomo#gh-dark-mode-only">
    <img src="./.github/assets/logo-dark.png" alt="Gomo">
  </a>
</p>

> Gomo is a Simple Golang multi modules tool.

## Installation ⬇

### Using script

- Shell

```bash
curl -sL https://bit.ly/gomo-cli | bash
```

- PowerShell

```powershell
iwr -useb https://bit.ly/gomo-win | iex
```

**then restart your powershell**

### Homebrew

```
brew install abdfnx/tap/gomo
```

## Usage

- Creates a new gomo.json file in the current folder

```bash
gomo init
```

- Initialize a new module

```bash
gomo init --mod github.com/x/x2 --path dir
```

- Download go packages through all your modules

```bash
gomo
```

- Get a go package and add it through all modules

```bash
gomo get github.com/gorilla/mux
```

- Delete a go package through all modules

```bash
gomo delete github.com/example/example1
```

- Update all packages

```bash
gomo update
```

- Add any missing packages necessary to build all your modules

```bash
gomo tidy
```

## Gomo config file

```json
{
  "cmds": {
    "download": "go mod download",
    "update": "go get -u"
  },
  "modules": [
    ".",
    "test",
    "test/web"
  ]
}
```

### Technologies Used in Tran

- [**Charm**](https://charm.sh)
- [**Cobra**](https://github.com/spf13/cobra)
- [**Viper**](https://github.com/spf13/viper)
- [**GJson**](https://github.com/tidwall/gjson)

### License

gomo is licensed under the terms of [MIT](https://github.com/abdfnx/gomo/blob/main/LICENSE) license.
