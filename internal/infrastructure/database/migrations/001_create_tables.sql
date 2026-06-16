CREATE TABLE IF NOT EXISTS fragrances (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    family      VARCHAR(50) NOT NULL,
    intensity   VARCHAR(50) NOT NULL,
    description TEXT,
    top_notes   TEXT[],
    heart_notes TEXT[],
    base_notes  TEXT[],
    seasons     TEXT[],
    occasions   TEXT[],
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS products (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sku          VARCHAR(100) UNIQUE NOT NULL,
    slug         VARCHAR(255) UNIQUE NOT NULL,
    name         VARCHAR(255) NOT NULL,
    brand        VARCHAR(255) NOT NULL,
    description  TEXT,
    price        NUMERIC(10,2) NOT NULL,
    volume       INT NOT NULL,
    image_url    TEXT,
    stock        INT NOT NULL DEFAULT 0,
    fragrance_id UUID REFERENCES fragrances(id),
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);
