CREATE TABLE "users"(
    "user_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "username" VARCHAR(255) NOT NULL UNIQUE, -- Ensure username is unique
    "password" VARCHAR(255) NOT NULL,
    PRIMARY KEY("user_id")
);


CREATE TABLE "subscription"(
    "sub_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "transaction_id" BIGINT NOT NULL,
    "type_id" BIGINT NOT NULL, -- Foreign key to subscription_types
    PRIMARY KEY("sub_id"),
    FOREIGN KEY("transaction_id") REFERENCES "Transactions"("transaction_id"),
    FOREIGN KEY("user_id") REFERENCES "users"("user_id"),
    FOREIGN KEY("type_id") REFERENCES "subscription_types"("type_id")
);

CREATE TABLE "subscription_types"(
    "type_id" BIGINT NOT NULL,
    "type_name" VARCHAR(255) NOT NULL,
    "cpu_limit" BIGINT,
    "memory_limit" BIGINT,
    "namespace_limit" BIGINT,
    PRIMARY KEY("type_id")
);

CREATE TABLE "transactions"(
    "transaction_id" BIGINT NOT NULL,
    "transaction_no" BIGINT NOT NULL UNIQUE,
    "transaction_date" DATE NOT NULL,
    "amount" BIGINT NOT NULL,
    "validity" DATE NOT NULL,
    PRIMARY KEY("transaction_id")
);

CREATE TABLE "namespace"(
    "namespace_id" BIGINT NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "namespace" BIGINT NOT NULL,
    "frontend_image" VARCHAR(255) NOT NULL,
    "backend_image" VARCHAR(255) NOT NULL,
    "active" BOOLEAN NOT NULL,
    "last_modified" DATE NOT NULL,
    PRIMARY KEY("namespace_id"),
    FOREIGN KEY("username") REFERENCES "users"("username") -- Ensure username consistency
);



