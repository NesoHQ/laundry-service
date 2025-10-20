# Feature: User Authentication

## Problem Solved
Unauthenticated users cannot place or track laundry orders securely.

## Solution Overview
Provides secure login and registration for users and admins using email and password authentication.

## User Flow
1. User registers with email or social login.
2. User logs in with credentials.
3. System verifies token/session for every protected request.

## Technical Details
- **Implementation**: Go net/http handlers with JWT or cookie-based authentication.
- **Dependencies**: None (no third-party library).
- **Non-Functional**: Passwords hashed using bcrypt or SHA-256 + salt; uses HTTPS.

## Acceptance Criteria
- [x] Registration and login endpoints working.
- [x] Only authenticated users can access protected routes.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add token refresh feature.

## Contribution Notes
Help wanted for token middleware improvements.

---

# Feature: User Management

## Problem Solved
Admin needs to manage users and shops centrally.

## Solution Overview
Allows the admin to create, update, delete, and view users or shops.

## User Flow
1. Admin logs in.
2. Admin views user/shop list.
3. Admin performs CRUD actions.

## Technical Details
- **Implementation**: Admin-only endpoints using Go net/http.
- **Dependencies**: PostgreSQL for data storage.
- **Non-Functional**: Role-based access control.

## Acceptance Criteria
- [x] Admin can CRUD users and shops.
- [x] Unauthorized users cannot access admin routes.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add activity logging.

## Contribution Notes
Seeking improvements for RBAC middleware.

---

# Feature: Shop Management

## Problem Solved
Each laundry shop must manage its own services and customers.

## Solution Overview
Enables laundry shops to define their services, pricing, and customer management.

## User Flow
1. Shop logs in.
2. Shop adds or edits available laundry services.
3. Shop views and manages user orders.

## Technical Details
- **Implementation**: HTTP routes for CRUD service management.
- **Dependencies**: PostgreSQL (tables: shops, services).
- **Non-Functional**: Optimized queries with indexing.

## Acceptance Criteria
- [ ] Shops can create and edit services.
- [ ] Shop data linked correctly with orders.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add analytics dashboard.

## Contribution Notes
Need optimization suggestions for Postgres schema.

---

# Feature: Order Management

## Problem Solved
Users need to order multiple laundry services and track progress.

## Solution Overview
Provides order creation, tracking, updating, and cancellation functionality.

## User Flow
1. User selects services and uploads images.
2. User places order.
3. System updates order status automatically (collected → in process → completed).
4. User can cancel before “in process”.

## Technical Details
- **Implementation**: Order and status APIs using Go handlers.
- **Dependencies**: PostgreSQL (orders, order_items tables).
- **Non-Functional**: Atomic DB transactions for order consistency.

## Acceptance Criteria
- [ ] Users can place and cancel orders.
- [ ] Order status transitions follow defined logic.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add automated notifications.

## Contribution Notes
Help wanted for status transition logic testing.

---

# Feature: Image Upload

## Problem Solved
Users need to attach photos of their clothes for identification.

## Solution Overview
Handles secure image upload and validation.

## User Flow
1. User selects image file.
2. System validates format and size.
3. System stores image securely (local directory or cloud).

## Technical Details
- **Implementation**: Go file upload endpoint using multipart/form-data.
- **Dependencies**: Standard library (no third party).
- **Non-Functional**: Max size limit, file type whitelist (JPEG, PNG).

## Acceptance Criteria
- [ ] Valid images are stored successfully.
- [ ] Invalid formats are rejected with error message.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add image compression.

## Contribution Notes
Testing volunteers needed for upload validation.

---

# Feature: Payment System

## Problem Solved
Users need a way to pay for laundry services.

## Solution Overview
Allows users to choose between cash or MFS (mobile financial services) payment methods.

## User Flow
1. User places an order.
2. User selects payment method.
3. System marks payment status upon confirmation.

## Technical Details
- **Implementation**: Go payment handler with transaction records.
- **Dependencies**: PostgreSQL for payment table.
- **Non-Functional**: Transaction integrity, secure record storage.

## Acceptance Criteria
- [ ] Users can select payment method.
- [ ] Payment status updates correctly.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add payment gateway integration.

## Contribution Notes
Open for MFS integration contributors.

---

# Feature: Review System

## Problem Solved
Customers need to share feedback after completing an order.

## Solution Overview
Allows verified customers to post reviews for completed orders.

## User Flow
1. User completes an order.
2. User posts a review.
3. System stores review linked to order and laundry shop.

## Technical Details
- **Implementation**: Go handler for reviews (CRUD).
- **Dependencies**: PostgreSQL (reviews table).
- **Non-Functional**: Review posting limited to completed orders.

## Acceptance Criteria
- [ ] Only completed orders allow reviews.
- [ ] Reviews linked to correct shop/user.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add rating aggregation.

## Contribution Notes
Need input on average rating calculation.

---

# Feature: Drop & Delivery Management

## Problem Solved
Users need flexibility in dropping off or receiving clothes.

## Solution Overview
Supports both home pickup and self-drop services with tracking.

## User Flow
1. User selects delivery method during order.
2. Laundry updates collection and delivery status.
3. User confirms receipt.

## Technical Details
- **Implementation**: Delivery scheduling logic using Go handlers.
- **Dependencies**: PostgreSQL (delivery table).
- **Non-Functional**: Delivery time estimation field in DB.

## Acceptance Criteria
- [ ] Users can choose between home pickup or self-drop.
- [ ] Delivery updates reflect correctly in order status.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add location-based tracking.

## Contribution Notes
Looking for design ideas for delivery time algorithm.
