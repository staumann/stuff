UPDATE positions
SET
    amount = ?,
    description = ?,
    price = ?,
    discount = ?,
    bill = ?,
    type = ?
WHERE id = ?;