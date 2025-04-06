# Gerador de Senhas Seguras em Go

Este repositório contém um gerador de senhas seguras escrito em Go que cria senhas aleatórias com caracteres diversos e certas obrigatoriedades para aumentar a segurança.

## Funcionalidades

- Gera senhas aleatórias com comprimento personalizável
- Inclui automaticamente:
  - Pelo menos uma letra maiúscula
  - Pelo menos um número
- Caracteres opcionais incluem:
  - Letras minúsculas
  - Caracteres especiais (`!@#$%^&*()+?><:{}[]`)
- Embaralha os caracteres para maior aleatoriedade

## Como Usar

### Requisitos
- Go instalado (versão 1.16 ou superior recomendada)

### Instalação
1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/gerador-senha-go.git
   cd gerador-senha-go
   ```

2. Execute o programa:
   ```bash
   go run main.go
   ```

### Personalização
Por padrão, o programa gera uma senha com 5 caracteres. Para modificar:

1. Edite o arquivo `main.go` e altere o argumento na linha:
   ```go
   password := generatePassword(5) // Altere 5 para o tamanho desejado
   ```

2. Ou modifique o programa para receber entrada do usuário (sugestão de melhoria).

## Exemplo de Saída
```
xK9!j
```

## Estrutura do Código

### Função Principal
- `main()`: Ponto de entrada do programa, chama a função de geração de senha e imprime o resultado.

### Função de Geração de Senha
- `generatePassword(length int) string`: 
  - Recebe o comprimento desejado da senha
  - Define os conjuntos de caracteres disponíveis
  - Garante a inclusão de pelo menos uma letra maiúscula e um número
  - Preenche o restante da senha com caracteres aleatórios
  - Embaralha os caracteres antes de retornar a senha
