-- Script de inicialização PostgreSQL
-- Criação de tabelas para exemplos com Go

-- Tabela de usuários
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    idade INTEGER,
    ativo BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de produtos
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    preco DECIMAL(10,2) NOT NULL,
    categoria VARCHAR(50),
    estoque INTEGER DEFAULT 0,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de pedidos
CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER,
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pendente' CHECK (status IN ('pendente', 'processando', 'enviado', 'entregue', 'cancelado')),
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);

-- Função para atualizar timestamp automaticamente
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.atualizado_em = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger para atualizar timestamp na tabela usuarios
CREATE TRIGGER update_usuarios_timestamp
    BEFORE UPDATE ON usuarios
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

-- Inserindo dados de exemplo
INSERT INTO usuarios (nome, email, idade) VALUES
('João Silva', 'joao@email.com', 25),
('Maria Santos', 'maria@email.com', 30),
('Pedro Oliveira', 'pedro@email.com', 22),
('Ana Costa', 'ana@email.com', 28);

INSERT INTO produtos (nome, descricao, preco, categoria, estoque) VALUES
('Notebook Dell', 'Notebook para desenvolvimento', 2500.00, 'Eletrônicos', 10),
('Mouse Gamer', 'Mouse com LED RGB', 150.00, 'Periféricos', 25),
('Teclado Mecânico', 'Teclado para programadores', 300.00, 'Periféricos', 15),
('Monitor 24"', 'Monitor Full HD', 800.00, 'Eletrônicos', 8);

INSERT INTO pedidos (usuario_id, total, status) VALUES
(1, 2500.00, 'entregue'),
(2, 450.00, 'processando'),
(3, 150.00, 'pendente'),
(1, 800.00, 'enviado');