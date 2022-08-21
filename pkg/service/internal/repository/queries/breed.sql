-- name: by-animal-id
SELECT id,
       name
FROM breeds
WHERE animal_id = ?
ORDER BY name ASC