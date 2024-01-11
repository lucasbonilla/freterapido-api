CREATE SCHEMA IF NOT EXISTS freterapidoapi;

-- Tabela 'carrier'
CREATE TABLE freterapidoapi.carrier (
    id_carrier INTEGER PRIMARY KEY,
    carrier_name VARCHAR(255) NOT NULL
);

-- Tabela 'quote'
CREATE TABLE freterapidoapi.quote (
    id_quote SERIAL PRIMARY KEY,
    id_carrier INTEGER REFERENCES freterapidoapi.carrier(id_carrier),
    price_quote FLOAT NOT NULL
);