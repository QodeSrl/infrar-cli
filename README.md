# Infrar CLI

**Command-line interface for Infrar - Initialize, transform, and deploy multi-cloud applications**

## 📌 Status: Basic CLI in Engine Repo

Currently, a basic transformation CLI tool exists in [infrar-engine](https://github.com/QodeSrl/infrar-engine) at `cmd/transform/`.

**Working Now**:
```bash
# Clone infrar-engine
git clone git@github.com:QodeSrl/infrar-engine.git
cd infrar-engine

# Build the transform tool
go build -o bin/transform ./cmd/transform

# Transform code
./bin/transform -provider aws -input app.py -output app_aws.py
```

## 🎯 Planned Full CLI (This Repo)

This repository will house a comprehensive CLI with additional commands:

```bash
# Initialize new project
infrar init my-app --language python

# Transform code
infrar transform --provider aws --input app.py

# Deploy application (future)
infrar deploy --provider aws --region us-east-1

# Compare costs (future)
infrar cost compare

# Manage configuration
infrar config set aws.credentials ~/.aws/credentials
```

## 🗂️ Planned Structure

```
infrar-cli/
├── cmd/
│   ├── init.go         # Project initialization
│   ├── transform.go    # Code transformation
│   ├── deploy.go       # Deployment orchestration
│   ├── cost.go         # Cost comparison
│   └── config.go       # Configuration management
├── pkg/
│   ├── project/        # Project scaffolding
│   ├── provider/       # Cloud provider integrations
│   └── deployer/       # Deployment logic
├── go.mod
└── main.go
```

## 🔧 Why Separate CLI?

**Decision Pending**: We're evaluating whether to:

**Option A**: Keep basic CLI in `infrar-engine` (current approach)
- Simpler architecture
- CLI is thin wrapper around engine
- Easier to maintain

**Option B**: Build comprehensive CLI here (planned approach)
- More features (init, deploy, config)
- Standalone distribution
- Better user experience

## 🚀 What's Available Today

Use the CLI in infrar-engine:

```bash
# Transform stdin
echo "from infrar.storage import upload" | ./bin/transform -provider aws

# Transform file
./bin/transform -provider aws -plugins ../infrar-plugins/packages -input app.py

# Transform to GCP
./bin/transform -provider gcp -input app.py -output app_gcp.py
```

See [infrar-engine README](https://github.com/QodeSrl/infrar-engine#-quick-start) for full CLI documentation.

## 🔗 Related Repositories

- [infrar-engine](https://github.com/QodeSrl/infrar-engine) - ✅ Transformation engine (includes basic CLI)
- [infrar-plugins](https://github.com/QodeSrl/infrar-plugins) - ✅ Transformation rules
- [infrar-sdk-python](https://github.com/QodeSrl/infrar-sdk-python) - 🚧 Python SDK (coming soon)
- [infrar-docs](https://github.com/QodeSrl/infrar-docs) - ✅ Documentation

## 📅 Timeline

Decision on CLI architecture: Within 1 week
Implementation (if separate): Phase 1B (after SDK complete)

## 📄 License

Apache License 2.0

---

**Part of the Infrar project** - Infrastructure Intelligence for the multi-cloud era
