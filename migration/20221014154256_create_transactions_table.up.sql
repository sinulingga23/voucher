CREATE TABLE transactions (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    voucher_id VARCHAR(36) NULL,
    qtty INT NOT NULL,
    total_point INT NOT NULL,
    CONSTRAINT transactions_voucher_id_on_vouchers FOREIGN KEY (voucher_id) REFERENCES vouchers (id)
    ON UPDATE CASCADE ON DELETE NO ACTION
);