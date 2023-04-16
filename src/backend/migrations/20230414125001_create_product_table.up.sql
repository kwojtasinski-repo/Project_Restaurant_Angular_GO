CREATE TABLE `products` (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    description TEXT(5000),
    price decimal(14,4) NOT NULL,
    category_id int NOT NULL,
    deleted boolean NOT NULL,
    INDEX idx__products__name (name),
    CONSTRAINT `fk__products__category_id__categories__id` FOREIGN KEY (category_id) REFERENCES `categories`(id) ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
