# Uptime Monitor Service

This is a simple uptime monitor service that checks the uptime of a given URL and stores the data in a database.
Note: This is not a production-ready service and is for learning purposes only.
Also Note: This is also also a project for me to learn Go. If you have any suggestions, please feel free to contribute.


## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Technologies](#technologies)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Coding Guidelines](#coding-guidelines)
7. [Estimated Time to Complete the Project](#estimated-time-to-complete-the-project)
8. [Requirements](#requirements)
9. [User Stories](#user-stories)
10. [Acceptance Criteria](#acceptance-criteria)

## 1. Overview

This is a simple uptime monitor service that checks the uptime of a given URL and stores the data in a database.

### 1.1. Features

- Check the uptime of a given URL and payload of the URL 
- Store the data in a database 

### 1.2. Technologies

- Go
- PostgreSQL
- Docker
- Docker Compose

### 1.2. Installation
#### 1.2.1. Prerequisites
- Docker (Docker desktop)
    - [Download](https://www.docker.com/products/docker-desktop/)
    - For windows users, you'll need to install the WSL2 and then install the Docker desktop.
#### 1.2.2. Installation
1. Clone the repository
2. Run `docker-compose up --build`
3. The server will be available at `http://localhost:8080`

### 1.3. Usage

1. Register a domain and content for the domain.
2. Update a domain and content for the domain.
3. View the uptime history of the domain.

