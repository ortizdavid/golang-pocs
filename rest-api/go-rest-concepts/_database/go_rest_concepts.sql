
DROP DATABASE IF EXISTS go_rest_concepts;
CREATE DATABASE go_rest_concepts;
USE go_rest_concepts;


DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
  task_id int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  task_name varchar(100) NOT NULL,
  description varchar(200) NOT NULL,
  start_date date DEFAULT NULL,
  end_date date DEFAULT NULL,
  attachment varchar(100) DEFAULT NULL,
  unique_id varchar(50) DEFAULT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (task_id),
  UNIQUE KEY unique_id (unique_id),
  KEY fk_task_user (user_id)
);


DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
  user_id int NOT NULL AUTO_INCREMENT,
  user_name varchar(100) NOT NULL,
  password varchar(150) NOT NULL,
  image varchar(100) DEFAULT NULL,
  active enum('Yes','No') DEFAULT 'Yes',
  unique_id varchar(50) DEFAULT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP,
  token varchar(150) DEFAULT NULL,
  PRIMARY KEY (user_id),
  UNIQUE KEY unique_id (unique_id)
);


DROP VIEW IF EXISTS view_task_data;
CREATE VIEW view_task_data AS 
SELECT 
    ta.task_id AS task_id,
    ta.unique_id AS unique_id,
    ta.task_name AS task_name,
    ta.description AS description,
    DATE_FORMAT(ta.start_date, '%Y-%m-%d') AS start_date,
    DATE_FORMAT(ta.end_date, '%Y-%m-%d') AS end_date,
    ta.attachment AS attachment,
    DATE_FORMAT(ta.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
    DATE_FORMAT(ta.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at,
    us.user_id AS user_id,
    us.user_name AS user_name
FROM 
    tasks ta 
    JOIN users us ON (us.user_id = ta.user_id) 
ORDER BY ta.created_at DESC;


DROP VIEW IF EXISTS view_user_data;
CREATE VIEW view_user_data AS 
SELECT 
    us.user_id AS user_id,
    us.unique_id AS unique_id,
    us.user_name AS user_name,
    us.password AS password,
    us.active AS active,
    us.image AS image,
    DATE_FORMAT(us.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
    DATE_FORMAT(us.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at
FROM 
    users us 
ORDER BY us.created_at DESC;

