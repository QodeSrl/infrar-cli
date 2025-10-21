# Infrar CLI

**Command-line interface for Infrar - Transform and manage multi-cloud applications**

## Status: ✅ Active Development

The Infrar CLI is now the primary command-line tool for working with Infrar transformations.

## Installation

### Build from Source

```bash
git clone https://github.com/QodeSrl/infrar-cli.git
cd infrar-cli
go build -o bin/infrar .
```

### Add to PATH (Optional)

```bash
# Copy to /usr/local/bin
sudo cp bin/infrar /usr/local/bin/

# Or add to your PATH
export PATH="$PATH:/path/to/infrar-cli/bin"
```

## Quick Start

### Transform Code

```bash
# Transform from file to AWS
infrar transform --provider aws --input app.py --output app_aws.py

# Transform to GCP
infrar transform --provider gcp --input app.py --output app_gcp.py

# Transform from stdin
cat app.py | infrar transform --provider aws

# Use custom plugin directory
infrar transform --provider aws --plugins ./my-plugins --input app.py
```

### Get Help

```bash
# General help
infrar --help

# Transform command help
infrar transform --help

# Version
infrar --version
```

## Commands

### `infrar transform`

Transform Infrar SDK code to provider-specific code.

**Flags**:
- `--provider, -p` - Target cloud provider (aws, gcp, azure) [default: aws]
- `--plugins` - Path to plugins directory [default: ../infrar-plugins/packages]
- `--capability, -c` - Capability to transform (storage, database, etc.) [default: storage]
- `--input, -i` - Input file to transform (or use stdin)
- `--output, -o` - Output file (or use stdout)

**Examples**:

```bash
# Basic transformation
infrar transform --provider aws --input app.py --output app_aws.py

# Pipe from stdin
echo "from infrar.storage import upload
upload(bucket='test', source='file.txt', destination='file.txt')" | infrar transform --provider aws

# Specify plugins location
infrar transform --provider gcp --plugins /path/to/plugins --input app.py
```

## Architecture

The Infrar CLI uses [infrar-engine](https://github.com/QodeSrl/infrar-engine) as a library for the transformation logic.

```
infrar-cli/
├── cmd/
│   ├── root.go          # Root command setup (Cobra)
│   └── transform.go     # Transform command implementation
├── main.go              # CLI entry point
├── go.mod               # Dependencies (including infrar-engine)
└── bin/
    └── infrar           # Built binary
```

## Development

### Dependencies

- Go 1.21 or later
- infrar-engine (included as Go module dependency)
- Cobra CLI framework

### Build

```bash
# Build CLI
go build -o bin/infrar .

# Run tests
go test ./...

# Install dependencies
go mod tidy
```

### Local Development with Engine

The CLI uses a local replace directive to use the local infrar-engine:

```go
// go.mod
replace github.com/QodeSrl/infrar-engine => ../infrar-engine
```

This allows development of both CLI and engine simultaneously without publishing.

## Related Repositories

- [infrar-engine](https://github.com/QodeSrl/infrar-engine) - Transformation engine library
- [infrar-sdk-python](https://github.com/QodeSrl/infrar-sdk-python) - Python SDK (v0.1.0)
- [infrar-plugins](https://github.com/QodeSrl/infrar-plugins) - Transformation rules
- [infrar-docs](https://github.com/QodeSrl/infrar-docs) - Documentation

## Future Commands

Planned commands for future releases:

- `infrar init` - Initialize new project
- `infrar deploy` - Deploy application
- `infrar cost` - Cost comparison
- `infrar config` - Configuration management

## License

Apache License 2.0

---

**Part of the Infrar project** - Infrastructure Intelligence for the multi-cloud era
