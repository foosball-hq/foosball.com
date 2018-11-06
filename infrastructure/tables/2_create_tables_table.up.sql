CREATE TABLE tables (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  description TEXT,
  location geography(POINT),
  address VARCHAR(255),
  city VARCHAR(255),
  state VARCHAR(50),
  zip VARCHAR(10), 
  phone VARCHAR(15)
);