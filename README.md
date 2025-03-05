![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/joshua-temple/goconf)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/joshua-temple/goconf/test.yml?label=tests)
![GitHub issues](https://img.shields.io/github/issues/joshua-temple/goconf)
![GitHub stars](https://img.shields.io/github/stars/joshua-temple/goconf?style=social)

# goconf

goconf is a Go tool that provides a CLI and library for generating configuration constants from YAML files and updating Go code with new constant values. Use it as a standalone CLI or import its functions directly into your tooling.

## Features

- **Configuration Generation:** Generate Go constants from YAML configuration files.
- **Constant Updates:** Replace old constant usages with new ones by comparing constant values.
- **Backup & Dry-run:** Optionally backup files before modifying and run in dry-run mode.
- **Concurrent File Processing:** Speeds up updates on large codebases.
- **Cobra-based CLI:** Easy-to-use command line interface.

## Installation

1. Clone the repository:

```bash
go install github.com/joshua-temple/goconf@latest
```

2. Build the CLI:

```bash
go build -o goconf ./cmd/goconf
```

## Usage

### As a CLI

#### Generate Configuration Constants

Generate constants from YAML configuration files.  
You can specify the target path for your YAML files, output file path, package name, and whether to create backups.

```bash
./goconf generate -t ./configs -o ./generated -p config
```

#### Update Constant Usages

Replace old constant names with new ones based on matching values from two Go files.  
Options include dry-run and backup.  
Provide the new and old constant files, and the directories containing files to update.

```bash
./goconf update -n new.go -o old.go --dirs ./pkg,./internal -d -b
```

- `-d`/`--dry-run`: Display changes without writing.
- `-b`/`--backup`: Create a backup file (with `.bak` extension) before modifying.

### As a Library

Import the core functions into your project without using the CLI initialization:

```go
package main

import "github.com/joshua-temple/goconf/pkg/goconf"

var (
    targetPath  = "./configs"
    outPath     = "./generated"
    packageName = "config"
    backup      = true

    oldFile     = "old.go"
    newFile     = "new.go"
    dryRun      = false
    directories = []string{"./pkg", "./internal"}
)

func main() {
	// ...
	
	// Generate configuration constants:
	err := goconf.GenerateConstants(targetPath, outPath, packageName, backup)
	if err != nil {
		// handle error
	}

	// Update constant usages:
	err = goconf.UpdateConstants(oldFile, newFile, dryRun, backup, directories)
	if err != nil {
		// handle error
	}

}
```

## Testing

Run unit tests with:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please submit issues and pull requests on GitHub. Make sure your changes are covered by tests and follow the existing style.

## License

[MIT License](LICENSE)

## Donations

![Static Badge](https://img.shields.io/badge/XRP_(2174028412)-rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg-blue?style=for-the-badge&logo=XRP&logoColor=white)

![Static Badge](https://img.shields.io/badge/Hedera_(3927122871)-0.0.1133968-blue?style=for-the-badge&logo=Hedera&logoColor=white)

![Static Badge](https://img.shields.io/badge/Bitcoin-bc1q7f49uuzrq2hwyclct78whaaekwpcd6r4p69mj9-blue?style=for-the-badge&logo=Bitcoin&logoColor=yellow)

![Static Badge](https://img.shields.io/badge/Polygon-0x242F47E66d725fDd5116c87a586066c5D0Cd1d3C-blue?style=for-the-badge&logo=Polygon&logoColor=pink)

![Static Badge](https://img.shields.io/badge/SUI-0x4e637f0dca15d71f3ed967ffd188e20ce9fe49f372396f251c85a1e0541d7de4-blue?style=for-the-badge&logo=Sui&logoColor=blue)

![Static Badge](https://img.shields.io/badge/Base-0x242F47E66d725fDd5116c87a586066c5D0Cd1d3C-blue?style=for-the-badge&logo=Ethereum&logoColor=white)

![Static Badge](https://img.shields.io/badge/Ethereum-0x242F47E66d725fDd5116c87a586066c5D0Cd1d3C-blue?style=for-the-badge&logo=Ethereum&logoColor=white)

![Static Badge](https://img.shields.io/badge/Solana-FuTked8fnRAiFsYNjQPsnuPQBcdQd2QZ2K822LvjYxvu-blue?style=for-the-badge&logo=Solana&logoColor=teal)




