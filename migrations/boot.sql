CREATE TABLE alerts (
  id SERIAL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  payload JSONB
);