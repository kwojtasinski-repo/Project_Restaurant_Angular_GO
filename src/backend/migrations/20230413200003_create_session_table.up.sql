CREATE TABLE `sessions` (
    session_id BINARY(16) PRIMARY KEY,
    user_id int,
    email varchar(255),
    role varchar(100),
    expiry DATETIME,
    INDEX idx_sessions_user_id (user_id)
);
