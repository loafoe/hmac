CREATE TABLE alerts (
  id SERIAL,
  created_at TIMESTAMP WITH TIMEZONE NOT NULL DEFAULT now(),
  payload JSONB
);
