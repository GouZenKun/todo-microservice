
CREATE TABLE IF NOT EXISTS todo_status (
  status_id SERIAL,
  status_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (status_id)
);

CREATE TABLE IF NOT EXISTS todos (
  todo_id SERIAL,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  status_id BIGINT UNSIGNED NOT NULL,
  due_date DATE,
  del_flag BOOLEAN NOT NULL DEFAULT FALSE,
  create_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (todo_id),
  FOREIGN KEY (status_id) REFERENCES todo_status(status_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;