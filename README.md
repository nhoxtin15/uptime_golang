# Uptime Monitor Service

This is a simple uptime monitor service that checks the uptime of a given URL and stores the data in a database.
Coding guidelines are provided under for beginners to try for the project with me.
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

- Check the uptime of a given URL
- Store the data in a database
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

## 2. Coding Guidelines (for developers)
### 2.1. Technologies
- Backend: Go
- Database: PostgreSQL
- Deployment: Docker and Docker Compose
### Estimated Time to Complete the Project
- 10 hours ~ 20 hours (depends on the experience of the developer)

### 2.2. Software Requirements Specification (SRS)

#### 2.2.1. Introduction
##### 2.2.1.1. User Stories

- As a user, I want to be able to register a domain so that I can check the uptime of the domain.
    - I should be able to provide the domain name and the API key.
    - I should be able to provide the body, headers, and the method of the request.
    - I should be able to provide the frequency of the uptime check.
    - I should be able to provide the timeout of the uptime check.
    - I should be able to provide the max retries of the uptime check.
    - I should be able to provide the retry interval of the uptime check.
    - I should be able to provide the retry backoff of the uptime check.
    - I should be able to provide the retry backoff factor of the uptime check.
    - I should be able to provide the retry backoff max of the uptime check.
    - I should be able to provide the retry backoff max jitter of the uptime check.


##### 2.2.1.2. Acceptance Criteria

- The domain is registered successfully if the domain is not already registered.
- The domain is updated successfully if the domain is already registered.
- The uptime history is displayed successfully if the domain is registered.

##### 2.2.1.3. Scope
The system provides:
- Domain registration  
- Domain configuration updates  
- Automated periodic uptime checks  
- Retry, timeout, and backoff logic  
- History storage and retrieval  
This SRS covers backend requirements only. UI, billing, and analytics are out of scope.

##### 2.2.1.4. Definitions, Acronyms, and Abbreviations
- **API**: Application Programming Interface  
- **REST**: Representational State Transfer  
- **Uptime Check**: Automated HTTP request performed to determine availability  
- **Backoff**: Retrying strategy with delays  
- **Domain Monitor**: A registered domain with monitoring configurations  

##### 2.2.1.5. References
- PostgreSQL Documentation  
- Go Standard Library HTTP Package  
- Docker & Docker Compose Documentation  

##### 2.2.1.6. Overview
This document describes system features, constraints, API requirements, non-functional requirements, models, risks, and appendices.

#### 2.2.2. Overall Description
The system is a standalone backend service exposing a REST API.  
It executes periodic tasks using internal schedulers and stores data in PostgreSQL.

##### 2.2.2.1. Product Functions
- Register a domain  
- Update monitoring settings  
- Perform automated uptime checks  
- Store uptime results  
- Provide uptime history  

##### 2.2.2.2. User Classes and Characteristics
- **User**: Registers domains and views history  
- **System Scheduler**: Performs periodic automated checks  
- **Admin (Optional)**: Database and system maintenance  

##### 2.2.2.3. Operating Environment
- Backend: Go  
- Database: PostgreSQL  
- Containerized: Docker + Docker Compose  
- OS: Linux Ubuntu Server  

##### 2.2.2.4. Design and Implementation Constraints
- REST API over HTTPS  
- JSON request/response format  
- Maximum check frequency limited to avoid overload  
- Scheduled tasks must run concurrently but safely  

##### 2.2.2.5. User Documentation
- API documentation (Swagger/OpenAPI)  
- Developer README  
- Deployment guide  

##### 2.2.2.6. Assumptions and Dependencies
- Internet required for uptime checks  
- Valid domains must be provided  
- PostgreSQL database must be available  

#### 2.2.3. System Features

##### 2.2.3.1. Register Domain

##### 2.2.3.1.1. Description and Priority
Allows user to register a domain with monitoring configuration.  
Priority: High.

##### 2.2.3.1.2. Stimulus/Response Sequence
- User calls `POST /domains`  
- System validates input  
- System stores domain in database  
- System schedules first uptime check  

##### 2.2.3.1.3. Functional Requirements
- **FR-1.1**: The system shall allow users to provide the domain URL.  
- **FR-1.2**: The system shall accept method, headers, and body.  
- **FR-1.3**: The system shall allow setting frequency, timeout, retries, retry interval, and backoff parameters.  
- **FR-1.4**: The system shall return error if domain is already registered.  
- **FR-1.5**: The system shall schedule recurring tasks after registration.

##### 2.2.3.2. Update Domain Configuration

##### 2.2.3.2.1. Description and Priority
Allows users to modify monitoring configuration of an existing domain.  
Priority: High.

##### 2.2.3.2.2. Stimulus/Response Sequence
- User calls `PUT /domains/{id}`  
- System validates and updates configuration  
- System regenerates monitoring schedule  

##### 2.2.3.2.3. Functional Requirements
- **FR-2.1**: The system shall update frequency, timeout, retries, and backoff.  
- **FR-2.2**: The system shall prevent duplicates.  
- **FR-2.3**: The system shall validate field types and limits.  

##### 2.2.3.3. View Uptime History

##### 2.2.3.3.1. Description and Priority
Allows users to retrieve recorded uptime history.  
Priority: Medium.

##### 2.2.3.3.2. Stimulus/Response Sequence
- User calls `GET /domains/{id}/history`  
- System fetches history from database  
- System returns JSON list of check results  

##### 2.2.3.3.3. Functional Requirements
- **FR-3.1**: The system shall store every check result.  
- **FR-3.2**: The system shall allow retrieval with pagination.  
- **FR-3.3**: The system shall include timestamps, status code, response time, and error messages.  

---

#### 2.2.4. External Interface Requirements

##### 2.2.4.1. User Interfaces
No UI included; API only.

##### 2.2.4.2. API Interfaces
- `POST /domains`  
- `PUT /domains/{id}`  
- `GET /domains/{id}`  
- `GET /domains/{id}/history`  

All in JSON format.

##### 2.2.4.3. Hardware Interfaces
N/A.

#### 2.2.4.4. Software Interfaces
- PostgreSQL  
- DNS/HTTP network services  

#### 2.2.4.5. Communications Interfaces
- HTTPS  
- JSON  
- UTF-8 encoding  

---

#### 2.2.5. Non-Functional Requirements

##### 2.2.5.1. Performance Requirements
- Uptime check execution must be within timeout.  
- API must respond within 200ms under normal load.  

##### 2.2.5.2. Security Requirements
- API key authentication required.  
- All communications must use HTTPS.  

##### 2.2.5.3. Reliability Requirements
- Retry mechanism must follow configured rules.  
- Application must gracefully handle network failures.  

##### 2.2.5.4. Scalability Requirements
- System must support up to 10,000 monitored domains.  

##### 2.2.5.5. Maintainability Requirements
- Modular Go packages.  
- Code must follow Go linting standards.  

##### 2.2.5.6. Portability Requirements
- Must run under Docker on Linux.  

---

#### 2.2.6. Data Requirements

##### 2.2.6.1. Database Design
Tables:
- **domains**: domain info & configuration  
- **uptime_checks**: results history  

##### 2.2.6.2. Data Storage Requirements
- Data must persist in PostgreSQL volumes.  
- History retained for minimum 90 days.

##### 2.2.6.3. Data Validation Rules
- Domain must be valid URL.  
- Frequency must be positive integer.  
- Retries must be >= 0.  
- Method must be valid HTTP method.  

