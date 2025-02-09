-- Create users table
CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       email VARCHAR(50) UNIQUE NOT NULL,
                       password_hash VARCHAR(200) NOT NULL,
                       created_at TIMESTAMP NOT NULL
);

-- Create ENUM types
CREATE TYPE priority_enum AS ENUM ('Low', 'Medium', 'High');
CREATE TYPE status_enum AS ENUM ('Pending', 'Completed');

-- Create tasks table
CREATE TABLE tasks (
                       id UUID PRIMARY KEY,
                       user_id UUID,
                       title VARCHAR(200) NOT NULL,
                       description TEXT,
                       due_date TIMESTAMP,
                       priority priority_enum,
                       category TEXT,
                       status status_enum,
                       created_at TIMESTAMP NOT NULL ,
                       updated_at TIMESTAMP NOT NULL ,
                       FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Insert dummy data for testing
INSERT INTO users (id, email, password_hash,created_at) VALUES
                                                            ('af85cd26-89dd-4b83-a8ea-37ff2b11d8c5', 'ad1@test1.com','szedxfcgvhbdxcfgvhbj','2025-01-25 18:46:00'),
                                                            ('745efc53-fb83-4edd-9b2e-cc8bc21d1a54', 'ad2@test2.com','vghbjngvhbjnkghbjkjbh','2025-01-25 18:46:00');

INSERT INTO tasks (id, user_id, title, description, due_date, priority, category,status,created_at,updated_at) VALUES
                                                                                                                   ('112ae328-a2da-4560-8934-c4f63f376911', 'af85cd26-89dd-4b83-a8ea-37ff2b11d8c5','Go to gym','health conscious','2024-12-10 18:46:00','Low','GYM','Pending','2025-01-25 18:46:00','2025-01-25 18:46:00'),
                                                                                                                   ('222ae328-a2da-4560-8934-c4f63f376922', '745efc53-fb83-4edd-9b2e-cc8bc21d1a54','Study','good for career','2024-12-10 18:47:00','High','Career','Pending','2025-01-25 18:46:00','2025-01-25 18:46:00'),
                                                                                                                   ('332ae328-a2da-4560-8934-c4f63f376933', 'af85cd26-89dd-4b83-a8ea-37ff2b11d8c5','Eat diet','body maintenance','2024-12-10 18:48:00','Medium','GYM','Completed','2025-01-25 18:46:00','2025-01-25 18:46:00'),
                                                                                                                   ('442ae328-a2da-4560-8934-c4f63f376944', '745efc53-fb83-4edd-9b2e-cc8bc21d1a54','Get a job','prepare for interviews','2024-12-10 18:49:00','High','Career','Completed','2025-01-25 18:46:00','2025-01-25 18:46:00');

