CREATE TABLE `carts`(
    id int PRIMARY KEY AUTO_INCREMENT,
    product_id int NOT NULL,
    user_id int NOT NULL,
    CONSTRAINT `fk__carts__product_id__products__id` FOREIGN KEY (product_id) REFERENCES `products`(id) ON DELETE CASCADE,
    CONSTRAINT `fk__carts__user_id__users__id` FOREIGN KEY (user_id) REFERENCES `users`(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
