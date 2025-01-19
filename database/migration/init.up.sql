
CREATE TABLE recipes (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         title VARCHAR(255) NOT NULL,
                         description TEXT,
                         cuisine VARCHAR(100),
                         photo_url TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP NULL DEFAULT NULL
);


CREATE TABLE ingredients (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             name VARCHAR(255) NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             deleted_at TIMESTAMP NULL DEFAULT NULL
);


CREATE TABLE recipe_ingredients (
                                    id INT AUTO_INCREMENT PRIMARY KEY,
                                    recipe_id INT NOT NULL,
                                    ingredient_id INT NOT NULL,
                                    quantity VARCHAR(100),
                                    unit VARCHAR(50),
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    deleted_at TIMESTAMP NULL DEFAULT NULL,
                                    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
                                    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE
);


CREATE TABLE instructions (
                              id INT AUTO_INCREMENT PRIMARY KEY,
                              recipe_id INT NOT NULL,
                              step_number INT NOT NULL,
                              description TEXT NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              deleted_at TIMESTAMP NULL DEFAULT NULL,
                              FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
);
