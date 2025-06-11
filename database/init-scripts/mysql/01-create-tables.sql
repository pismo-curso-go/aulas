-- Script de inicialização MySQL
-- Criação de tabelas para exemplos com Go

USE aula_go;

-- Configurar charset UTF-8 para a sessão
SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

-- Tabela de usuários
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    idade INT,
    ativo BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela de produtos
CREATE TABLE produtos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    preco DECIMAL(10,2) NOT NULL,
    categoria VARCHAR(50),
    estoque INT DEFAULT 0,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela de pedidos
CREATE TABLE pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT,
    total DECIMAL(10,2) NOT NULL,
    status ENUM('pendente', 'processando', 'enviado', 'entregue', 'cancelado') DEFAULT 'pendente',
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Inserindo dados de exemplo com acentuação
INSERT INTO usuarios (nome, email, idade) VALUES
('João Silva', 'joao@email.com', 25),
('Maria Santos', 'maria@email.com', 30),
('Pedro Oliveira', 'pedro@email.com', 22),
('Ana Costa', 'ana@email.com', 28),
('José da Silva', 'jose@email.com', 35),
('Lúcia Fernández', 'lucia@email.com', 27);

INSERT INTO produtos (nome, descricao, preco, categoria, estoque) VALUES
('Notebook Dell', 'Notebook para desenvolvimento', 2500.00, 'Eletrônicos', 10),
('Mouse Gamer RGB', 'Mouse gamer com iluminação LED', 150.00, 'Periféricos', 25),
('Teclado Mecânico', 'Teclado mecânico para programação', 300.00, 'Periféricos', 15),
('Monitor 24"', 'Monitor Full HD para trabalho', 800.00, 'Eletrônicos', 8),
('SSD 500GB', 'SSD NVMe para alta performance', 400.00, 'Armazenamento', 20),
('Câmera HD', 'Câmera para videoconferências', 200.00, 'Periféricos', 12);

INSERT INTO pedidos (usuario_id, total, status) VALUES
(1, 2500.00, 'entregue'),
(2, 450.00, 'processando'),
(3, 150.00, 'pendente'),
(1, 800.00, 'enviado'),
(5, 700.00, 'processando'),
(6, 200.00, 'entregue'),
(2, 400.00, 'pendente');