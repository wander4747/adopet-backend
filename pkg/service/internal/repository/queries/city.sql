-- name: by-state-id
SELECT id,
       name,
       state_id
FROM cities
WHERE state_id = ?
ORDER BY name ASC