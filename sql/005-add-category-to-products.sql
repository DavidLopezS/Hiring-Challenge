ALTER TABLE products 
    ADD COLUMN category_id INTEGER NULL;

ALTER TABLE products
    ADD CONSTRAINT fk_product_category
        FOREIGN KEY (category_id)
        REFERENCES product_categories(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE;