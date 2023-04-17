CREATE TABLE `sessions` (
    session_id BINARY(16) PRIMARY KEY,
    user_id int NOT NULL,
    email varchar(255) NOT NULL,
    role varchar(100) NOT NULL,
    expiry DATETIME NOT NULL,
    INDEX idx__sessions__user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
