# Golang Backend with Docker & CLI: A Developer’s Approach 🚀

In this project, I’ve built a complete Golang server using the go-fiber framework and integrated all services with a Command-Line Interface (CLI) using the cobra library. This setup provides a seamless development experience, combining the power of a backend server with the flexibility of CLI commands.

## Steps to run CLI

1. Start the server using Docker by running make compose-with-debug. Additional commands are available in the Makefile for your convenience.

2. Navigate to the cli folder and build the CLI binary using the command:

   ```bash
       go build -o task
   ```

3. Run the CLI using the generated binary:
   ```bash
   ./task
   ```

This will allow you to interact with the backend services directly through the CLI.

<b>For more understanding how service is created, the database connection and how the CLI is created please refer to this medium blog</b>: https://medium.com/@sharmavivek1709/golang-backend-with-docker-cli-a-devs-approach-7700665f2daf
