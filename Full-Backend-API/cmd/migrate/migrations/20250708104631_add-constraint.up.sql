ALTER TABLE orders
ADD CONSTRAINT `fk_products` FOREIGN KEY (`productId`) REFERENCES products(`id`);