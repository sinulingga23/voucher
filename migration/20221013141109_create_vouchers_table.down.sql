ALTER TABLE vouchers DROP FOREIGN KEY vouchers_brand_id_on_brands;
ALTER TABLE vouchers DROP INDEX vouchers_brand_id_on_brands;
DROP TABLE IF EXISTS vouchers;