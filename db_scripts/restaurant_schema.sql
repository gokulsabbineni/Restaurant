CREATE DATABASE restaurant_db;
USE restaurant_db;

CREATE TABLE users (
  user_id INT PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  user_type ENUM('customer', 'staff') NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  contact_number VARCHAR(20)
);

CREATE TABLE tables (
  table_id INT PRIMARY KEY,
  table_number INT NOT NULL,
  capacity INT NOT NULL,
  location VARCHAR(50),
  description TEXT
);

CREATE TABLE reservations (
    reservation_id INT PRIMARY KEY,
    customer_id INT,
    table_id INT,
    reservation_date DATE NOT NULL,
    reservation_time TIME NOT NULL,
    party_size INT NOT NULL,
    status ENUM('Pending', 'Confirmed', 'Cancelled', 'Completed') NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES users(user_id),
    FOREIGN KEY (table_id) REFERENCES tables(table_id)
);

CREATE VIEW reservation_history AS
SELECT
  reservations.reservation_id,
  reservations.customer_id,
  reservations.table_id,
  reservations.reservation_date,
  reservations.reservation_time,
  reservations.party_size,
  reservations.status,
  users.username,
  users.first_name,
  users.last_name
FROM reservations
JOIN users ON reservations.customer_id = users.user_id;