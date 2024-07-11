-- Active: 1720611425800@@127.0.0.1@5432@db_employee_management@public
-- Create the position (Master_Department) table
CREATE TABLE IF NOT EXISTS master_department (
  department_id SERIAL PRIMARY KEY,
  department_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255),
  deleted_at TIMESTAMP
);

-- Create the Master_Position table
CREATE TABLE IF NOT EXISTS master_position (
  position_id SERIAL PRIMARY KEY,
  department_id INTEGER,
  position_name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255),
  deleted_at TIMESTAMP,
  FOREIGN KEY (department_id) REFERENCES master_department(department_id)
);

-- Create the Master_Location table
CREATE TABLE IF NOT EXISTS master_location (
  location_id SERIAL PRIMARY KEY,
  location_name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255),
  deleted_at TIMESTAMP
);

-- Create the Employee table
CREATE TABLE IF NOT EXISTS employees (
  employee_id SERIAL PRIMARY KEY,
  employee_code VARCHAR(255),
  employee_name VARCHAR(255),
  Password VARCHAR(255), -- Assume encryption handled by application
  department_id INTEGER,
  position_id INTEGER,
  superior INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255),
  deleted_at TIMESTAMP,
  FOREIGN KEY (department_id) REFERENCES master_department(department_id),
  FOREIGN KEY (position_id) REFERENCES master_position(position_id)
);

-- Create the Attendance table
CREATE TABLE IF NOT EXISTS attendance (
  attendance_id SERIAL PRIMARY KEY,
  employee_id INTEGER,
  location_id INTEGER,
  absent_in TIMESTAMP,
  absent_out TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255),
  deleted_at TIMESTAMP,
  FOREIGN KEY (employee_id) REFERENCES employee(employee_id),
  FOREIGN KEY (location_id) REFERENCES master_location(location_id)
);
