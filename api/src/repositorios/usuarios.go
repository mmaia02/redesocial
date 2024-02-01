package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// Cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Insere um usuário no DB
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
