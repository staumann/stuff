UPDATE bills
SET
    userId = ?,
    shopId = ?,
    discount = ?,
    timestamp = ?
WHERE id = ?;