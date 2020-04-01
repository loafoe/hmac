CREATE TABLE alerts (
  id SERIAL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  payload JSONB
);
