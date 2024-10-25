CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "group" (
    "id" SERIAL PRIMARY KEY,
    "telegram_id" BIGINT NOT NULL UNIQUE,
    "user_id" INTEGER NOT NULL,
    "openai_model" VARCHAR(255) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

CREATE TABLE "topic" (
    "id" SERIAL PRIMARY KEY,
    "thread_id" BIGINT NOT NULL UNIQUE,
    "group_id" INT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "openai_model" VARCHAR(255) NOT NULL,

    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("group_id") REFERENCES "group" ("id")

);

CREATE TABLE "message" (
    "id" SERIAL PRIMARY KEY,
    "telegram_id" BIGINT NOT NULL UNIQUE,
    "text" TEXT NOT NULL,
    "topic_id" INTEGER NOT NULL,
    "sender" VARCHAR(255) NOT NULL,

    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("topic_id") REFERENCES "topic" ("id"),
    CONSTRAINT sender_check CHECK (sender IN ('user', 'assistant'))
);
