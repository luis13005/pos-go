package dto

type CreateProductDto struct {
	Nome  string  `json:"nome"`
	Preco float32 `json:"preco"`
}

type GetJWT struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type CreateUserDto struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
