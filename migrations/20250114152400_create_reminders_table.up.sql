CREATE TYPE reminder_target AS ENUM ('user', 'company');

--bun:split

CREATE TABLE reminders (
                       id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

                       author_id         VARCHAR(255) NOT NULL,
                       public_identifier VARCHAR(255) NOT NULL,
                       target            reminder_target NOT NULL,

                       content           TEXT NOT NULL,
                       updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                       expired_at        TIMESTAMP WITH TIME ZONE NOT NULL,

                       CONSTRAINT fk_reminders_author_id UNIQUE (author_id, target, public_identifier)
);

--bun:split

CREATE INDEX reminders_per_author ON reminders (author_id);
