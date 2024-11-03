# E-commerce API
This project is a RESTful API for an e-commerce application built with **Golang**. It supports user
management, product catalog management, and order processing, with JWT-based authentication
for secure access. Admin-specific routes are available for managing products and orders. The API is
designed to work with both **MySQL** and **PostgreSQL** databases and is containerized with
Docker and Docker Compose.
## Table of Contents
- [Features](#features)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [API Documentation](#api-documentation)
- [Public Routes](#public-routes)
- [User Routes (Requires Authentication)](#user-routes-requires-authentication)
- [Product Routes (Requires Authentication)](#product-routes-requires-authentication)
- [Order Routes (Requires Authentication)](#order-routes-requires-authentication)

## Features
- **User Management**: Allows users to register, log in, view, and update their profiles.
- **Product Management**: Admin-only functionality to create, update, and delete products.
- **Order Management**: Authenticated users can place, view, and cancel orders. Admins can
update order statuses.
- **Authentication**: JWT-based token authentication, with role-based access control to secure
admin-only routes.
- **Database Flexibility**: Supports both MySQL and PostgreSQL databases, configurable via
environment variables.
- **Docker Integration**: Containerized for easy setup and deployment with Docker and Docker
Compose.
## Getting Started
### Prerequisites
Ensure you have the following installed:
- **Golang** 1.20 or later
- **Docker** and **Docker Compose**
- **Git**
### Installation
1. **Clone the repository**:
```bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
```
2. **Set up environment variables**:
- Copy `sample.env` to `.env` and update it with the required configuration based on your
database and application setup.
```bash
cp .sample.env .env
```
3. **Run with Docker Compose**:
```bash
docker-compose up --build
```
This command builds the Docker image for the Go application, starts up MySQL and PostgreSQL
containers, and sets up the necessary networking.


## API Documentation
### Public Routes
- **POST /api/v1/register**: Register a new user.
- **POST /api/v1/login**: Log in and receive a JWT token.
- **GET /api/v1/healthcheck**: Health check for the API.
### User Routes (Requires Authentication)
- **GET /api/v1/users**: View all users (admin only).
- **GET /api/v1/users/:id**: View user by ID.
- **PUT /api/v1/users/:id**: Update user by ID.
### Product Routes
- **GET /api/v1/products**: View all products (public).
- **GET /api/v1/products/:id**: View product by ID (public).
- **POST /api/v1/products**: Create a new product (admin only).
- **PUT /api/v1/products/:id**: Update product by ID (admin only).
- **DELETE /api/v1/products/:id**: Delete product by ID (admin only).
### Order Routes (Requires Authentication)
- **POST /api/v1/orders**: Place a new order.
- **GET /api/v1/orders**: View all orders for the authenticated user.
- **GET /api/v1/orders/:id**: View order by ID.
- **DELETE /api/v1/orders/:id**: Cancel an order by ID.
- **PUT /api/v1/orders/:id/status**: Update order status (admin only).
