-- CREATING AND USING DATABASE
CREATE DATABASE db_task;
USE db_task;

-- CREATING TABLES

DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
	task_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	task_name VARCHAR(100) NOT NULL,
	description VARCHAR(200) NOT NULL,
	status ENUM('Completed', 'In Progress', 'Pending') DEFAULT 'Pending',
	start_date DATE NOT NULL,
	end_date DATE NOT NULL
);
