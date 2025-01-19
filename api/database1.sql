CREATE TABLE "users" (
    "user_id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL
);

CREATE TABLE "subscription"(
    "sub_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "transaction_id" BIGINT NOT NULL
);
ALTER TABLE
    "subscription" ADD PRIMARY KEY("sub_id");
CREATE TABLE "Transactions"(
    "transaction_id" BIGINT NOT NULL,
    "date" DATE NOT NULL,
    "transaction_no" BIGINT NOT NULL,
    "amount" BIGINT NOT NULL,
    "model" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "Transactions" ADD PRIMARY KEY("transaction_id");
CREATE TABLE "containers"(
    "deployment_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "active" BOOLEAN NOT NULL,
    "last_modified" DATE NOT NULL
);
ALTER TABLE
    "containers" ADD PRIMARY KEY("deployment_id");
ALTER TABLE
    "subscription" ADD CONSTRAINT "subscription_transaction_id_foreign" FOREIGN KEY("transaction_id") REFERENCES "Transactions"("transaction_id");
ALTER TABLE
    "containers" ADD CONSTRAINT "containers_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "subscription" ADD CONSTRAINT "subscription_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");