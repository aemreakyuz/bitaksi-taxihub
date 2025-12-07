# 🚕 Bitaksi TaxiHub - Taxi Management System

A microservices-based taxi management system built with Go, featuring geolocation services, API Gateway, and MongoDB persistence.

## 🏗️ Architecture
```
User → API Gateway (Port 3000) → Driver Service (Port 8081) → MongoDB
```

### Services

- **API Gateway**: Routes and manages incoming requests
- **Driver Service**: Manages driver data with geolocation features
- **MongoDB**: Persistent data storage

## ✨ Features

### Core Features
- ✅ Driver CRUD operations (Create, Read, Update)
- ✅ Pagination support for driver listings
- ✅ Geolocation-based nearby driver search (6km radius using Haversine formula)
- ✅ RESTful API design
- ✅ Microservices architecture

### Bonus Features
- ⭐ Interactive Swagger/OpenAPI documentation
- ⭐ Docker Compose for one-command deployment
- ⭐ Request logging middleware

