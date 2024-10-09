# Market Data Analysis Project

This project is a hybrid Go-Python system designed for the quantitative analysis of market data, 
utilizing gRPC for communication between the services. 
The system ingests market data in real-time, preprocesses it in Go, 
and performs data analysis and machine learning tasks in Python.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)

## Overview

This project aims to provide a high-performance, low-latency system for analyzing market data using a hybrid approach:
- **Go Service**: Handles data ingestion, preprocessing, and calls the Python service for analysis.
- **Python Service**: Performs data analysis and machine learning tasks and returns results to the Go service.
- **gRPC Communication**: Facilitates efficient and secure data exchange between the two services.

## Architecture

1. **Go Service**: Ingests market data, preprocesses it, and sends requests to the Python service.
2. **Python Service**: Analyzes data using statistical methods and machine learning models, returning results to the Go service.
3. **gRPC**: Used for high-performance communication between Go and Python services.

## Technologies Used

- **Go**: For data ingestion, preprocessing, and high-throughput processing.
- **Python**: For data analysis, machine learning, and advanced statistical computations.
- **gRPC**: For inter-service communication between Go and Python.
- **AWS**: For deployment, including EC2 for hosting, S3 for storage, and TimescaleDB for time-series data.
- **Docker**: For containerizing services and ensuring consistent deployments.
