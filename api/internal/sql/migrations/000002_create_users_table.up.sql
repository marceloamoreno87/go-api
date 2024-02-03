CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Password: admin
INSERT INTO users (name, password, email) VALUES ('admin', '$2a$10$FcjdWT805.CjOEz9xc/P9eJojZ0.3SLlLRAgI/2ve6zPjgGY2jFsS', 'admin@admin.com.br');
