BEGIN;
USE logistic;

CREATE TABLE location (
    id bigint AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100),
    city VARCHAR(50),
    address VARCHAR(255)
);

CREATE TABLE recipient (
    id INT PRIMARY KEY,
    name VARCHAR(10),
    address VARCHAR(255),
    phone VARCHAR(10)
);

CREATE TABLE product (
     id INT PRIMARY KEY ,
     name VARCHAR(50)
);


CREATE TABLE `order` (
     id INT PRIMARY KEY,
     recipient_id INT,
     product_id INT

--      FOREIGN KEY (recipient_id) REFERENCES recipient(id),
--      FOREIGN KEY (product_id) REFERENCES product(id)
);

CREATE TABLE logistic_detail (
     id INT PRIMARY KEY,
     order_id INT,
     location_id INT,
     time DATETIME,
     status VARCHAR(30)

--      FOREIGN KEY (order_id) REFERENCES `order`(id),
--      FOREIGN KEY (location_id) REFERENCES location(id)
);

COMMIT;