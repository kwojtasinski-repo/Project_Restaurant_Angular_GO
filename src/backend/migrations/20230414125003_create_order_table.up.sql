CREATE TABLE `orders`(
    id int PRIMARY KEY,
    order_number varchar(300) NOT NULl,
    price decimal(14,4) NOT NULL,
    created DATETIME NOT NULL,
    modified DATETIME NULL,
    user_id int NOT NULL,
    INDEX idx__orders__order_number (order_number),
    CONSTRAINT `fk__orders__user_id__users__id` FOREIGN KEY (user_id) REFERENCES `users`(id) ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
