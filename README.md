# API Structure

## Run

```bash
go run main.go
```

```bash
go run main.go --listenaddr :5000
```

## Model
**_NOTE_**: Package 'types' is the same as 'models', the guy doesn't like MVC

### Validation
You could create another package for validation, but is not needed, could be done on types package

## Data Layer
Package 'storage'
Always make an interface for your storage: storage/storage.go
Then the implementation will be on that package: storage/mongodb.go
For testing is important to create something in memory: storage/memory.go

# Routing
Here's is allowed to use mux router
