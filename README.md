# microservice-template

Example use practices for robovoice.

### Getting Started

```
make run
```

##### Prerequisites

+ docker
+ docker-compose
+ protoc 3.7.1+

### Architecture

```
.
├── /src/                       # The public source code of the application
├── /internal/                  # The private source code of the application
│   ├── /application/           # Write business logic
│   ├── /domian/                # Entity struct that represent mapping to data model
│   └── /infrastructure/        # Solves backend technical topics
├── /ops/                       # All infrastructure configuration for IoC
├── .gitignore                  # A gitignore file specifies untracked files
└── README.md                   # README
```
