# Kickstart Go API

Kickstart Go API is a CLI tool that scaffolds Go API projects with a choice of routers: Gin, Mux, or Chi. It allows you to quickly generate boilerplate code for a new Go API project with just a few commands.

## Features
- Choose between popular Go routers: Gin, Mux, or Chi.
- Interactive mode or CLI options to configure your project.
- Generates a basic structure to kickstart your API development.

## Prerequisites

Before using this tool, make sure you have Go installed on your system.

### Install Go

If you don't have Go installed, follow the official installation guide:
- **Go Installation Documentation**: [https://golang.org/doc/install](https://golang.org/doc/install)

#### Quick Install for Go

For most Unix-based systems (Linux/macOS), you can install Go with:

```bash
brew install go
```

For Windows, follow the instructions in the official documentation.

#### Export Go Path

After installing Go, ensure your Go environment variables are correctly set. Add the following lines to your `.bashrc`, `.zshrc`, or equivalent shell configuration:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

You may need to restart your terminal or run `source ~/.bashrc` or `source ~/.zshrc` to apply the changes.

## Installation

Once Go is installed and your environment is set up, you can install the **kickstart-go-api** CLI tool:

```bash
go install github.com/earlysvahn/kickstart-go-api@latest
```

This will install the `kickstart-go-api` binary into your `$GOPATH/bin`, which you should have in your `PATH`.

## Usage

You can use **kickstart-go-api** either in interactive mode or by providing arguments via the CLI.

### 1. Interactive Mode

Simply run the following command:

```bash
kickstart-go-api
```

You will be prompted to enter the project name and select a router from the list of available options (Gin, Mux, Chi).

### 2. CLI Mode

You can also specify the project name and router via command-line arguments:

```bash
kickstart-go-api <project-name> --router=<router-name>
```

Example:

```bash
kickstart-go-api my-awesome-api --router gin
```

In this example, the project `my-awesome-api` will be created using the Gin router.

### Version Check

To check the current version of the tool and see if an update is available, run:

```bash
kickstart-go-api --version
```

This will show you the current installed version and notify you if a new version is available.

## Example Usage

Here's how you can use `kickstart-go-api` to create a new project interactively:

```bash
$ kickstart-go-api
Please enter the project name: my-api
? Choose router: 
  ? gin
    mux
    chi
```

Alternatively, to create a new project using the CLI:

```bash
$ kickstart-go-api my-api --router gin
Creating project with Gin router...
Project my-api has been created with gin router!
```

### Available Routers

- **Gin**: A high-performance HTTP web framework in Go.
- **Mux**: A powerful URL router and dispatcher for matching incoming requests to their handlers.
- **Chi**: A lightweight, idiomatic and composable router for building Go HTTP services.

## Contributing

Feel free to fork this repository and submit pull requests for new features, bug fixes, or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
