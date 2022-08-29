create table combustivel(
    id serial primary key,
    nome varchar
);

create table posto(
    id serial primary key,
    nome varchar
);

CREATE TABLE valor (
    id serial primary key,
    combustivel_id int DEFAULT NULL,
    posto_id int DEFAULT NULL,
    valor float DEFAULT NULL,    
    FOREIGN KEY (combustivel_id) REFERENCES combustivel (id),
    FOREIGN KEY (posto_id) REFERENCES posto (id)
);

INSERT INTO posto(nome) VALUES
('Posto 1'),
('Posto 2');

INSERT INTO combustivel(nome) VALUES
('Etanol'),
('Gasolina Comum'),
('Gasolina Adtivada'),
('Diesel');

INSERT INTO valor(combustivel_id,posto_id,valor) VALUES
(1,1,5.8),
(3,1,7.2),
(2,1,6.8),
(1,2,6.75);