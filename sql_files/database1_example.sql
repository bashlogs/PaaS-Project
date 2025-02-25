-- User Table

INSERT INTO "users" ("user_id", "name", "email", "username", "password", "sub_id") VALUES
(1, 'Alice Johnson', 'alice@example.com', 'alicej', 'password123', 101),
(2, 'Bob Smith', 'bob@example.com', 'bobsmith', 'securepass', 102),
(3, 'Charlie Brown', 'charlie@example.com', 'charlieb', 'mypassword', 103);


-- Insert subscription types

INSERT INTO "subscription_types" ("type_id", "type_name", "cpu_limit", "memory_limit", "namespace_limit") VALUES
(1, 'Basic', 2, 4, 1),
(2, 'Premium', 4, 8, 2),
(3, 'Standard', 3, 6, 1);


-- Insert subscriptions with type_id

INSERT INTO "subscription" ("sub_id", "user_id", "transaction_id", "type_id") VALUES
(101, 1, 201, 1),
(102, 2, 202, 2),
(103, 3, 203, 3);

-- Transactions Table

INSERT INTO "Transactions" ("transaction_id", "transaction_no", "transaction_date", "amount", "validity") VALUES
(201, 1001, '2023-01-15', 1000, '2023-12-31'),
(202, 1002, '2023-02-20', 1500, '2024-02-19'),
(203, 1003, '2023-03-25', 2000, '2024-03-24');



-- Namespace Table

INSERT INTO "namespace" ("namespace_id", "username", "namespace", "frontend_image", "backend_image", "active", "last_modified") VALUES
(401, 'alicej', 501, 'frontend1.img', 'backend1.img', TRUE, '2023-01-01'),
(402, 'bobsmith', 502, 'frontend2.img', 'backend2.img', FALSE, '2023-02-01'),
(403, 'charlieb', 503, 'frontend3.img', 'backend3.img', TRUE, '2023-03-01');

