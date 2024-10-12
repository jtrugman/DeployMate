# DeployMate
An opinionated simple deployment system

DeployMate is a powerful CLI tool designed to streamline the setup and management of deployment configurations for your projects. It helps developers quickly add GitHub Actions workflows and Dockerfiles to their repositories, making the deployment process more efficient and consistent.

## Features

- Add GitHub Actions workflows for sandbox and production environments
- Generate Dockerfiles for various project types (e.g., Next.js)
- Support for monorepo structures
- Easy-to-use command-line interface

## Installation

### Prerequisites

- Go 1.16 or higher

### Building from source

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/deploymate.git
   cd deploymate
   ```

2. Build the project:
   ```
   go build -o deploymate
   ```

3. (Optional) Move the binary to a directory in your PATH:
   ```
   sudo mv deploymate /usr/local/bin/
   ```

## Usage

### Adding deployment components

To add deployment components to your project, use the `add` command:
