# Project Structure

This document outlines the folder structure of this project.

- **/cmd/${PROJECT_NAME}/**: Contains the entry point of the application. This is where the main application executable is located.

- **/internal/**: Contains the private application and library code. This is code you do not want others importing in their applications.
  - **/config**: Configuration related code.
  - **/handlers**: HTTP handlers.
  - **/models**: Application data models.
  - **/repository**: Database access logic.
  - **/service**: Business logic.

- **/pkg/logger**: Contains the Zap logger setup, which can be imported by other applications.
- **/pkg/testutils**: Contains test utilities that can be imported by other applications.

- **/deployments**: Contains IaC (Infrastructure as Code), container orchestration (Docker, Kubernetes) configurations, and CI/CD setups.

- **/scripts**: Contains scripts for various build, install, analysis, etc., operations.

- **/api**: Intended for OpenAPI/Swagger specs, JSON schema files, protocol definition files.

- **/test**: Contains additional external test apps and test data. Place your integration tests here.