ALTER TABLE transactions DROP FOREIGN KEY transactions_voucher_id_on_vouchers;
ALTER TABLE transactions DROP INDEX transactions_voucher_id_on_vouchers;
DROP TABLE IF EXISTS transactions;