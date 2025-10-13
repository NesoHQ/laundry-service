# Roadmap: laundry-service MVP & Beyond

## Overview
- **MVP Goal**: Deliver a robust Go-based backend that enables users to order, track, and pay for laundry services with real-time status updates.
- **Launch Timeline**: Q2 2026.
- **Success Metrics**: 1000+ orders processed, 90% user satisfaction, and <2s average API response time.

## Feature Prioritization (MoSCoW Method)
Use this table to triage features. Focus on Must-Haves for MVP.

| Priority | Feature Name | Description | Status | Est. Effort (Story Points) | Dependencies |
|----------|--------------|-------------|--------|----------------------------|--------------|
| **Must-Have (MVP Core)** | User Authentication | JWT-based authentication with secure password hashing and token validation. | Not Started | 5 | None |
| **Must-Have (MVP Core)** | Product & Order Management | Core ordering and tracking system for laundry services. Handles CRUD and order states. | Not Started | 8 | User Authentication |
| **Must-Have (MVP Core)** | Payment System | Integrates cash and MFS (Bkash/Nagad) payments with transaction validation. | Planned | 5 | Product & Order Management |
| **Should-Have (Post-MVP)** | Shop Management | Allows laundry owners to register and manage their services and pricing. | Planned | 4 | User Authentication |
| **Should-Have (Post-MVP)** | Drop & Delivery Management | Enables home pickup and delivery scheduling using maps API. | Planned | 6 | Product & Order Management |
| **Could-Have (Nice-to-Have)** | Review System | Users with completed orders can leave verified reviews and ratings. | Backlog | 3 | Product & Order Management |
| **Could-Have (Nice-to-Have)** | Notifications | Email or push updates for order status changes. | Backlog | 3 | Product & Order Management |
| **Won't-Have (Out of Scope)** | Multi-language Support | Deferred until after MVP due to localization complexity. | Future | 13 | N/A |

## Phases

### Phase 1: MVP Launch (v0.1.0)
- **Core features**:  
  - User Authentication  
  - Product & Order Management  
  - Payment System  
- **Tech Stack**:  
  - Go 1.22+  
  - PostgreSQL with `sqlx`   
  - REST API using row `net/http`   
  - JWT for authentication  
- **Testing**:  
  - Unit tests with Go’s built-in testing package  
  - Postman/Insomnia API testing  
- **Deployment**:  
  - Docker + Docker Compose  
  - Optional CI/CD with GitHub Actions  
- **Risks**:  
  - Concurrency issues during multiple order updates  
  - Payment gateway API limits  
  - Initial map integration (location accuracy)

### Phase 2: Iteration (v0.2.0)
- **Based on feedback**:  
  - Add Shop Management, Drop & Delivery Management, and Review System.  
  - Improve API response time and error handling.  
- **Timeline**: 4–6 weeks post-MVP launch.  
- **Focus**: Performance tuning and modular refactoring (use cases, repositories, and services).

### Phase 3: Expansion (v0.3.0+)
- **Add-ons**:  
  - Notifications (via SMTP or WebSocket)  
  - Admin dashboard API endpoints  
  - Analytics and reporting for shop owners  
- **Goal**:  
  - Scale horizontally using load balancing  
  - Optimize DB queries with indexes and caching  
  - Add CI/CD with rolling updates  

## How to Contribute
- Open feature proposals via [GitHub Issues](https://github.com/NesoHQ/laundry-service/issues/new?template=feature_request.md).  
- Upvote high-impact ideas with 👍 reactions.  
- Follow `CONTRIBUTING.md` for PR guidelines and code review standards.  

*Last Updated: 2025-10-13. Feedback welcome!*
