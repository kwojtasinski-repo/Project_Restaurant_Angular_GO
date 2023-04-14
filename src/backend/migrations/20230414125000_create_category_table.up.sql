CREATE TABLE `categories` (
    id int PRIMARY KEY,
    name varchar(255) NOT NULL,
    deleted boolean NOT NULL,
    INDEX idx__categories__name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
