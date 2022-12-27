CREATE TABLE todo (
    id          SERIAL PRIMARY KEY,
    text        TEXT NOT NULL,
    done        BOOLEAN NOT NULL DEFAULT FALSE,
    createdAt   TIMESTAMP DEFAULT NOW()
)
