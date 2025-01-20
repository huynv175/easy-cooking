-- Thêm dữ liệu vào bảng ingredients
INSERT INTO ingredients (name, created_at, updated_at) VALUES
                                                           ('Egg', NOW(), NOW()),
                                                           ('Flour', NOW(), NOW()),
                                                           ('Milk', NOW(), NOW()),
                                                           ('Butter', NOW(), NOW()),
                                                           ('Salt', NOW(), NOW()),
                                                           ('Chicken', NOW(), NOW()),
                                                           ('Rice', NOW(), NOW()),
                                                           ('Tomato', NOW(), NOW()),
                                                           ('Onion', NOW(), NOW()),
                                                           ('Garlic', NOW(), NOW()),
                                                           ('Beef', NOW(), NOW()),
                                                           ('Pasta', NOW(), NOW()),
                                                           ('Cheese', NOW(), NOW()),
                                                           ('Olive Oil', NOW(), NOW()),
                                                           ('Pepper', NOW(), NOW());

-- Thêm dữ liệu vào bảng recipes
INSERT INTO recipes (title, description, cuisine, photo_url, created_at, updated_at) VALUES
                                                                                         ('Pancakes', 'Delicious fluffy pancakes', 'American', 'http://example.com/pancakes.jpg', NOW(), NOW()),
                                                                                         ('Scrambled Eggs', 'Soft and creamy scrambled eggs', 'American', 'http://example.com/scrambled_eggs.jpg', NOW(), NOW()),
                                                                                         ('Chicken Fried Rice', 'Classic Asian stir-fried rice with chicken', 'Asian', 'http://example.com/chicken_fried_rice.jpg', NOW(), NOW()),
                                                                                         ('Beef Stir Fry', 'Spicy beef stir-fry with vegetables', 'Asian', 'http://example.com/beef_stir_fry.jpg', NOW(), NOW()),
                                                                                         ('Pasta Carbonara', 'Creamy Italian pasta with eggs and cheese', 'Italian', 'http://example.com/pasta_carbonara.jpg', NOW(), NOW()),
                                                                                         ('Tomato Soup', 'Homemade creamy tomato soup', 'European', 'http://example.com/tomato_soup.jpg', NOW(), NOW()),
                                                                                         ('Grilled Chicken Salad', 'Fresh salad with grilled chicken', 'Mediterranean', 'http://example.com/chicken_salad.jpg', NOW(), NOW()),
                                                                                         ('Beef Burger', 'Classic American beef burger', 'American', 'http://example.com/beef_burger.jpg', NOW(), NOW()),
                                                                                         ('Vegetable Stir Fry', 'Healthy mixed vegetable stir fry', 'Asian', 'http://example.com/vegetable_stir_fry.jpg', NOW(), NOW()),
                                                                                         ('Cheese Omelette', 'Fluffy omelette with melted cheese', 'French', 'http://example.com/cheese_omelette.jpg', NOW(), NOW()),
                                                                                         ('Spaghetti Bolognese', 'Classic Italian pasta with meat sauce', 'Italian', 'http://example.com/spaghetti_bolognese.jpg', NOW(), NOW()),
                                                                                         ('Chicken Curry', 'Spicy Indian chicken curry', 'Indian', 'http://example.com/chicken_curry.jpg', NOW(), NOW());

-- Thêm mối quan hệ giữa recipes và ingredients
WITH ingredient_map AS (
    SELECT name, id FROM ingredients
),
     recipe_map AS (
         SELECT title, id FROM recipes
     )
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
    -- Pancakes
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Egg'), '2', 'pieces'),
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Flour'), '1', 'cup'),
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Milk'), '1', 'cup'),
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Butter'), '2', 'tablespoons'),

    -- Scrambled Eggs
    ((SELECT id FROM recipe_map WHERE title = 'Scrambled Eggs'), (SELECT id FROM ingredient_map WHERE name = 'Egg'), '3', 'pieces'),
    ((SELECT id FROM recipe_map WHERE title = 'Scrambled Eggs'), (SELECT id FROM ingredient_map WHERE name = 'Butter'), '1', 'tablespoon'),
    ((SELECT id FROM recipe_map WHERE title = 'Scrambled Eggs'), (SELECT id FROM ingredient_map WHERE name = 'Salt'), '1', 'pinch'),

    -- Chicken Fried Rice
    ((SELECT id FROM recipe_map WHERE title = 'Chicken Fried Rice'), (SELECT id FROM ingredient_map WHERE name = 'Chicken'), '200', 'grams'),
    ((SELECT id FROM recipe_map WHERE title = 'Chicken Fried Rice'), (SELECT id FROM ingredient_map WHERE name = 'Rice'), '2', 'cups'),
    ((SELECT id FROM recipe_map WHERE title = 'Chicken Fried Rice'), (SELECT id FROM ingredient_map WHERE name = 'Onion'), '1', 'piece'),

    -- Beef Stir Fry
    ((SELECT id FROM recipe_map WHERE title = 'Beef Stir Fry'), (SELECT id FROM ingredient_map WHERE name = 'Beef'), '300', 'grams'),
    ((SELECT id FROM recipe_map WHERE title = 'Beef Stir Fry'), (SELECT id FROM ingredient_map WHERE name = 'Onion'), '1', 'piece'),
    ((SELECT id FROM recipe_map WHERE title = 'Beef Stir Fry'), (SELECT id FROM ingredient_map WHERE name = 'Garlic'), '2', 'cloves'),

    -- Pasta Carbonara
    ((SELECT id FROM recipe_map WHERE title = 'Pasta Carbonara'), (SELECT id FROM ingredient_map WHERE name = 'Pasta'), '200', 'grams'),
    ((SELECT id FROM recipe_map WHERE title = 'Pasta Carbonara'), (SELECT id FROM ingredient_map WHERE name = 'Egg'), '2', 'pieces'),
    ((SELECT id FROM recipe_map WHERE title = 'Pasta Carbonara'), (SELECT id FROM ingredient_map WHERE name = 'Cheese'), '50', 'grams');
