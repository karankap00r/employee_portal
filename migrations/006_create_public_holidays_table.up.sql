CREATE TABLE IF NOT EXISTS public_holidays
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    country      TEXT    NOT NULL,
    start_date   TEXT    NOT NULL,
    end_date     TEXT    NOT NULL,
    name         TEXT    NOT NULL,
    status       TEXT    NOT NULL,
    is_mandatory BOOLEAN NOT NULL,
    created_by   TEXT    NOT NULL,
    created_at   TEXT    NOT NULL,
    updated_at   TEXT    NOT NULL,
    UNIQUE (country, start_date, end_date)
);
    