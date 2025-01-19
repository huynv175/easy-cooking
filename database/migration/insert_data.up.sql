-- Thêm dữ liệu vào bảng ingredients
INSERT INTO ingredients (name, created_at, updated_at) VALUES
                                                           ('Egg', NOW(), NOW()),
                                                           ('Flour', NOW(), NOW()),
                                                           ('Milk', NOW(), NOW()),
                                                           ('Butter', NOW(), NOW()),
                                                           ('Salt', NOW(), NOW());

-- Thêm dữ liệu vào bảng recipes
INSERT INTO recipes (title, description, cuisine, photo_url, created_at, updated_at) VALUES
                                                                                         ('Pancakes', 'Delicious fluffy pancakes', 'American', 'http://example.com/pancakes.jpg', NOW(), NOW()),
                                                                                         ('Scrambled Eggs', 'Soft and creamy scrambled eggs', 'American', 'http://example.com/scrambled_eggs.jpg', NOW(), NOW());

-- Thêm mối quan hệ giữa recipes và ingredients vào bảng recipe_ingredients
-- Lưu ý: Cần đảm bảo rằng bạn đã có id của các ingredients và recipes
-- Lấy id của các ingredients và recipes trước khi chèn vào bảng recipe_ingredients

-- Lấy id của các nguyên liệu
SET @egg_id = (SELECT id FROM ingredients WHERE name = 'Egg');
SET @flour_id = (SELECT id FROM ingredients WHERE name = 'Flour');
SET @milk_id = (SELECT id FROM ingredients WHERE name = 'Milk');
SET @butter_id = (SELECT id FROM ingredients WHERE name = 'Butter');
SET @salt_id = (SELECT id FROM ingredients WHERE name = 'Salt');

-- Lấy id của các recipes
SET @pancakes_id = (SELECT id FROM recipes WHERE title = 'Pancakes');
SET @scrambled_eggs_id = (SELECT id FROM recipes WHERE title = 'Scrambled Eggs');

-- Thêm vào bảng recipe_ingredients
INSERT INTO recipe_ingredients (recipe_id, ingredient_id) VALUES
                                                              (@pancakes_id, @egg_id),
                                                              (@pancakes_id, @flour_id),
                                                              (@pancakes_id, @milk_id),
                                                              (@pancakes_id, @butter_id),
                                                              (@scrambled_eggs_id, @egg_id),
                                                              (@scrambled_eggs_id, @butter_id);
