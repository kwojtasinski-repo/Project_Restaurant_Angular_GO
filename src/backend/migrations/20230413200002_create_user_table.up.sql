CREATE TABLE `users` (
    id int PRIMARY KEY AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    password varchar(500) NOT NULL,
    role varchar(100) NOT NULL,
    deleted boolean NOT NULL,
    INDEX idx__users__email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
