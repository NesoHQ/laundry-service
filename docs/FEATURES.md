# Feature: User Authentication

## Problem Solved
Users and admins need a secure and easy way to access the system to manage laundry services and orders.

## Solution Overview
Implements secure registration and login using email or social accounts, storing hashed passwords and generating secure tokens.

## User Flow
1. User signs up using email or social login (Google/Facebook).
2. System validates credentials and stores securely.
3. User receives authentication token for future requests.

## Technical Details
- **Implementation**: Node.js + Express with JWT-based auth.
- **Dependencies**: bcrypt, jsonwebtoken, express-validator.
- **Non-Functional**: Encrypted passwords, HTTPS enforced.

## Acceptance Criteria
- [ ] Users can register/login successfully.
- [ ] Tokens are valid and expire correctly.
- [ ] Unauthorized users cannot access protected endpoints.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add social login integrations.

## Contribution Notes
Help wanted for OAuth integrations. See CONTRIBUTING.md.
---

# Feature: User Management

## Problem Solved
Admin needs to manage users and laundry shops efficiently from a central interface.

## Solution Overview
Provides admin tools to create, update, and delete users and shops, with role-based access control.

## User Flow
1. Admin logs in.
2. Admin views user/shop lists.
3. Admin performs CRUD operations.

## Technical Details
- **Implementation**: REST API routes with role-based middleware.
- **Dependencies**: express, prisma or mongoose.
- **Non-Functional**: Secure authorization; fast response (<2s).

## Acceptance Criteria
- [ ] Admin can create/update/delete users and shops.
- [ ] Unauthorized users can’t access admin endpoints.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add bulk import/export feature.

## Contribution Notes
Contributors can extend admin dashboards or add logging.
---

# Feature: Shop Management

## Problem Solved
Laundry shop owners need a way to manage their offered services and view users who ordered from them.

## Solution Overview
Allows authenticated shop accounts to define services, pricing, and manage customer lists.

## User Flow
1. Shop owner logs in.
2. Adds or updates available laundry services.
3. Views customer list and order details.

## Technical Details
- **Implementation**: CRUD endpoints for shop services and profiles.
- **Dependencies**: express, database ORM.
- **Non-Functional**: Scalable for multi-shop support.

## Acceptance Criteria
- [ ] Shop can list, update, and delete services.
- [ ] Shop can view associated users and orders.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add revenue analytics per shop.

## Contribution Notes
Looking for contributors for shop analytics.
---

# Feature: Product & Order Management

## Problem Solved
Users need to order multiple laundry items, track progress, and manage delivery efficiently.

## Solution Overview
Enables users to create, view, and manage laundry orders with real-time status updates.

## User Flow
1. User browses nearby laundry shops.
2. Adds items (with image) to an order.
3. Submits order and tracks progress until delivery.

## Technical Details
- **Implementation**: RESTful endpoints for orders and products.
- **Dependencies**: multer (for image upload), database ORM.
- **Non-Functional**: Orders processed within 2s.

## Acceptance Criteria
- [ ] Orders have unique IDs and statuses.
- [ ] Users can track and cancel eligible orders.
- [ ] Laundry can update status in real time.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add push notifications for status updates.

## Contribution Notes
Contributors can improve tracking or notifications.
---

# Feature: Payment System

## Problem Solved
Users need a convenient and secure way to pay for laundry services.

## Solution Overview
Provides both cash-on-delivery and mobile financial service (MFS) payment options.

## User Flow
1. User selects payment method (Cash or MFS).
2. Completes payment confirmation.
3. Order marked as “Paid” in the system.

## Technical Details
- **Implementation**: Payment gateway integration (Bkash/Nagad).
- **Dependencies**: payment SDKs or REST APIs.
- **Non-Functional**: Secure transaction validation.

## Acceptance Criteria
- [ ] Payment status updates correctly.
- [ ] System handles failed or pending payments.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Integrate automatic MFS payment confirmation.

## Contribution Notes
PRs welcome for payment gateway integration.
---

# Feature: Drop & Delivery Management

## Problem Solved
Users require flexibility in how they send and receive their laundry.

## Solution Overview
Supports home collection, self-drop, and scheduled delivery options.

## User Flow
1. User chooses pickup or drop option.
2. System schedules collection/delivery.
3. User receives notification upon completion.

## Technical Details
- **Implementation**: Scheduler service + route management.
- **Dependencies**: cron jobs, map APIs.
- **Non-Functional**: <2s scheduling response.

## Acceptance Criteria
- [ ] Users can select drop/delivery mode.
- [ ] Laundry can update collection/delivery status.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add location tracking on maps.

## Contribution Notes
Need help integrating Google Maps API.
---

# Feature: Review System

## Problem Solved
Users need to share feedback to improve service quality and trust.

## Solution Overview
Allows verified users to post reviews after completing orders.

## User Flow
1. User completes order.
2. Accesses order history and writes a review.
3. Review displayed on the laundry shop profile.

## Technical Details
- **Implementation**: Review API linked with order status.
- **Dependencies**: express, database ORM.
- **Non-Functional**: Prevent duplicate or fake reviews.

## Acceptance Criteria
- [ ] Only completed orders can post reviews.
- [ ] Reviews appear under correct shop.

## Status & Roadmap
- **Current**: Planned.
- **Version**: MVP v0.1.
- **Next**: Add review ratings and analytics.

## Contribution Notes
Help needed for moderation or spam detection.
---
