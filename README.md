
# Gerador de Senhas CLI em Go

Uma ferramenta simples de linha de comando para gerar senhas aleatórias usando Golang.

## 🚀 Início Rápido

### 1. Compile o programa:


`go build password_generator.go` 

Isso irá gerar um executável chamado `password_generator` (ou `password_generator.exe` em sistemas Windows).

### 2. Execute o programa:

Para gerar uma senha com as configurações padrão:

`./password_generator` 

## ⚙️ Configurações

Você pode personalizar a senha gerada usando vários argumentos:

-   **-length**: Define o comprimento da senha. Padrão é `8`.
    
    
    `./password_generator -length=12` 
    
-   **-numbers**: Decide se a senha deve incluir números. Padrão é `true`.
    
    
    `./password_generator -numbers=false` 
    
-   **-symbols**: Decide se a senha deve incluir símbolos. Padrão é `true`.
    
    
    `./password_generator -symbols=false` 
    
-   **-uppercase**: Decide se a senha deve incluir letras maiúsculas. Padrão é `true`.
    
    
    `./password_generator -uppercase=false` 
    
-   **-lowercase**: Decide se a senha deve incluir letras minúsculas. Padrão é `true`.
    
    
    `./password_generator -lowercase=false` 
    

## 🔍 Exemplo

Para gerar uma senha de 12 caracteres, com números, símbolos, letras maiúsculas e minúsculas:


`./password_generator -length=12 -numbers=true -symbols=true -uppercase=true -lowercase=true` 

----------
