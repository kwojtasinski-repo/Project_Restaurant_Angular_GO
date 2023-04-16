CREATE TABLE `order_products`(
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(200) NOT NULl,
    price decimal(14,4) NOT NULL,
    product_id int NOT NULL,
    order_id int NOT NULL,
    INDEX idx__order_products__name (name),
    CONSTRAINT `fk__order_products__product_id__products__id` FOREIGN KEY (product_id) REFERENCES `products`(id) ON DELETE NO ACTION,
    CONSTRAINT `fk__order_products__order_id__orders__id` FOREIGN KEY (order_id) REFERENCES `orders`(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
