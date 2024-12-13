# Pandora Server

## Introduction
Secret management is a critical aspect of modern DevOps environments, where securely handling authentication tokens and other sensitive credentials is essential.  
To address these needs, we introduce **Pandora**, a secret management system built using the Golang Gin framework. Pandora is designed to run in a containerized environment and provides an API for seamless user interaction. Any HTTPS client operating from a machine with access to the server can act as a Pandora client.  
Pandora aims to offer an efficient and streamlined solution for securely managing secrets in DevOps workflows.

## Prerequisites
1. **Golang Installed**: Install Golang to compile the code.  
2. **Docker Installed**: Install Docker for containerized deployment.  
3. **Digital Certificate and Private Key**: Obtain a valid digital certificate and private key to enable secure HTTPS communication.

## Installation Guide
Follow these steps to set up the Pandora HTTPS server:

1. **Clone the Repository**: Clone the Pandora repository from GitHub:
```shell
   git clone https://github.com/MustafaAbdulazizHamza/Pandora-Server.git
```
2. **Compile the Project**: Compile the project so that it can be run:
```shell
go build -o Pandora
```
3.	**Add Digital Certificate and Private Key**: Place your digital certificate and private key files (named server.*) in the project directory.
4.	**Build Pandora Docker Image**: Create a custom Docker image for Pandora using the following command:
```shell
docker build -t pandora:v1 .
```
4.	**Create a Database Storage Volume**: Set up a volume for the Pandora database using:
```shell
docker volume create pandora-data
```
6. **Run Pandora Container**: Launch the Pandora container and attach the previously created volume:
```shell
docker run -d -p 8080:8080 --name pandora -v pandora-data:/root/ pandora:v1  
```
## Notes:
1. The default root password is root.
2. You can use Pandora-CLI for server configuration or service acquisition.
3. Several clients have been developed for this server, such as Pandora-CLI and Go-Pandora.
## Related Tools
1. **Pandora-Cli**: An HTTPS client CLI tool for simple user interaction with the Pandora Secrets Management System.
[Pandora-CLI](https://github.com/MustafaAbdulazizHamza/Pandora-CLI)
2. **Go-Pandora**: A Golang library that provides a client interface for using and managing the Pandora Secrets Management System was publish at this repository:
[Go-Pandora](https://github.com/MustafaAbdulazizHamza/go-pandora)
