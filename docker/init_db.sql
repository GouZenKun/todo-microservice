CREATE DATABASE IF NOT EXISTS db_todo;
USE db_todo;

CREATE TABLE IF NOT EXISTS todo_status (
  status_id SERIAL,
  status_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (status_id)
);

INSERT INTO todo_status (status_name)
VALUES ('Incomplete'),
       ('Pending'),
       ('Completed');

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

INSERT INTO todos (title, description, status_id, due_date, del_flag)
VALUES ('Todo 1', 'サンプルデータ1', 3, '2024-02-28', FALSE),
       ('Todo 2', 'サンプルデータ2', 1, '2024-04-30', FALSE),
       ('Todo 3', 'サンプルデータ3', 2, '2024-01-31', FALSE),
       ('Todo 4', 'サンプルデータ4', 3, '2023-12-31', FALSE),
       ('Todo 5', 'サンプルデータ5', 1, '2024-03-31', FALSE);
