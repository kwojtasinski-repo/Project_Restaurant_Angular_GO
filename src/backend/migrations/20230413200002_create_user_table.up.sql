CREATE TABLE `users` (
    id int PRIMARY KEY,
    email varchar(255),
    password varchar(500),
    role varchar(100),
    INDEX idx_users_email (email)
);
