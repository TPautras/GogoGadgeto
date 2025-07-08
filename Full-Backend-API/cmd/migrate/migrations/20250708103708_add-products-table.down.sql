ALTER TABLE orders
DROP FOREIGN KEY `fk_products`;

ALTER TABLE orders
DROP COLUMN `productId`;

DROP TABLE IF EXISTS products;
