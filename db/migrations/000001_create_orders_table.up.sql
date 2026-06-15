CREATE TABLE orders (
    order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id TEXT NOT NULL,
    reference_id TEXT,
    currency TEXT NOT NULL,
    status TEXT NOT NULL,
    total BIGINT NOT NULL CHECK (total > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
