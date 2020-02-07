INSERT INTO category(description) VALUES('fruits');
INSERT INTO category(description) VALUES('books');
INSERT INTO category(description) VALUES('computadores');
INSERT INTO category(description) VALUES('carros');

INSERT INTO products(name, price, quantity, amount, category) 
VALUES('Laranja', 1.80, 100, (1.8 * 100), 1);
INSERT INTO products(name, price, quantity, amount, category) 
VALUES('Banana', 1.80, 80, (1.8 * 80), 1);
INSERT INTO products(name, price, quantity, amount, category) 
VALUES('Limao', 1.80, 40, (1.8 * 40), 1);

INSERT INTO products(name, price, quantity, amount, category) 
VALUES('Manifesto Comunista', 20, 50, (20 * 50), 2);
INSERT INTO products(name, price, quantity, amount, category) 
VALUES('Vigiar e Punir', 50, 20, (50 * 20), 2);

INSERT INTO products(name, price, quantity, amount, category)
VALUES('Notebook Gamer', 3000, 5, (3000 * 5), 3);
INSERT INTO products(name, price, quantity, amount, category)
VALUES('Mouse CCE Gamer', 200, 13, (200 * 13), 3);

INSERT INTO products(name, price, quantity, amount, category)
VALUES('Bat-Móvel', 6303000, 2, (6303000 * 2), 4);
INSERT INTO products(name, price, quantity, amount, category)
VALUES('Fusca', 5000, 3, (5000 * 3), 4);

