CREATE TYPE "users_role_enum" AS ENUM ('admin', 'client');

CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "role" users_role_enum DEFAULT 'client',
    "display_name" varchar NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX "user_email" ON "users" ("email");
