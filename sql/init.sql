CREATE TABLE companies (
  id UUID PRIMARY KEY,
  name VARCHAR(15) UNIQUE NOT NULL,
  description TEXT,
  amount_of_employees INT NOT NULL,
  registered BOOLEAN NOT NULL,
  type VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password TEXT NOT NULL
);