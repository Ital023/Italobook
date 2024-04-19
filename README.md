# Italobook - :warning::construction: EM DESENVOLVIMENTO
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
## :rocket: Sobre o projeto
Desenvolvi este módulo em GoLang para integrar-se à minha própria rede social, inspirada no Facebook. Por enquanto, o módulo oferece apenas funcionalidades básicas, como criação, leitura, atualização e exclusão (CRUD) de perfis de usuários. Utilizando as capacidades nativas da linguagem Go e pacotes como net/http e database/sql, construí uma base sólida para a gestão de usuários na minha rede social. Estou planejando expandir este módulo no futuro para incluir uma gama mais ampla de recursos específicos de redes sociais, como interações entre usuários, compartilhamento de conteúdo e notificações.

## :bookmark_tabs: Versões API
<details> 
<summary> 1.0 Version</summary>
  
  <ul>
    
  ## Funcionalidades
  - [x] CRUD básico
  - [x] Validações de JSON e atributos do usuário 
  - [x] Tratamento de erros
  - [x] Persistência no SGBD

  ## 🛣️ Rotas 
  
  ```
  -GET -> localhost:5000/usuarios?usuario=nome (pesquisa pelo nome ou nick do usuario) <br>
  -POST -> localhost:5000/usuarios (recebe o body do request) <br>
  -PUT -> localhost:5000/{usuarioID} (Apartir do ID é possível alterar as informações do usuário como nome, nick, email) <br>
  -DELETE -> localhost:5000/{usuarioID} (Apartir do ID, o usuário correspondente é apagado do DB)
 ```

  
## :arrows_counterclockwise: Dependências
```github.com/joho/godotenv```: go get github.com/joho/godotenv <a href="https://github.com/joho/godotenv">[how-to-install]</a><br>
```github.com/gorilla/mux```: go get github.com/gorilla/mux <a href="https://github.com/gorilla/mux">[how-to-install]</a><br>
```github.com/go-sql-driver/mysql```: go get github.com/go-sql-driver/mysql <a href="https://github.com/go-sql-driver/mysql">[how-to-install]</a><br>
```github.com/badoux/checkmail```: go get github.com/badoux/checkmail <a href="https://github.com/badoux/checkmail">[how-to-install]</a><br>

  </ul>
</details>

<details> 
<summary> 1.1 Version</summary>
  
  <ul>
    
  ## Funcionalidades
  - [x] CRUD básico
  - [x] Validações de JSON e atributos do usuário 
  - [x] Tratamento de erros
  - [x] Persistência no SGBD
  - [x] Autenticação JWT
  - [x] Middlewares
  - [x] Validação de ações do usuário quando autenticado   

  ## 🛣️ Rotas 
  
  ```
  /usuarios
  -GET -> localhost:5000/usuarios?usuario=nome (pesquisa pelo nome ou nick do usuario) <br>
  -POST -> localhost:5000/usuarios (recebe o body do request) <br>
  -PUT -> localhost:5000/{usuarioID} (Apartir do ID é possível alterar as informações do usuário como nome, nick, email) <br>
  -DELETE -> localhost:5000/{usuarioID} (Apartir do ID, o usuário correspondente é apagado do DB)

  /login
  -POST -> localhost:5000/login (Recebe pelo body request o email e senha para gerar o token de autenticação)
 ```

## :arrows_counterclockwise: Dependências
```github.com/joho/godotenv```: go get github.com/joho/godotenv <a href="https://github.com/joho/godotenv">[how-to-install]</a><br>
```github.com/gorilla/mux```: go get github.com/gorilla/mux <a href="https://github.com/gorilla/mux">[how-to-install]</a><br>
```github.com/go-sql-driver/mysql```: go get github.com/go-sql-driver/mysql <a href="https://github.com/go-sql-driver/mysql">[how-to-install]</a><br>
```github.com/badoux/checkmail```: go get github.com/badoux/checkmail <a href="https://github.com/badoux/checkmail">[how-to-install]</a><br>
```github.com/dgrijalva/jwt-go``` go get github.com/dgrijalva/jwt-go <a href="https://github.com/dgrijalva/jwt-go">[how-to-install]</a><br>
```golang.org/x/crypto``` go get golang.org/x/crypto <a href="https://golang.org/x/crypto">[how-to-install]</a><br>


  </ul>
</details>

<details> 
<summary> 1.2 Version (FINAL)</summary>
  
  <ul>
    
  ## Funcionalidades
  - [x] CRUD usuarios, publicações e login
  - [x] Validações de JSON e atributos do usuário, publicação e login
  - [x] Tratamento de erros
  - [x] Persistência no SGBD (MYSQL)
  - [x] Autenticação JWT
  - [x] Middlewares
  - [x] Validação de ações do usuário quando autenticado   

  ## 🛣️ Rotas 

  ### Usuario - >
  ```
  -GET -> localhost:5000/usuarios?usuario=nome (pesquisa pelo nome ou nick do usuario) 
  -POST -> localhost:5000/usuarios (recebe o body do request)
  -PUT -> localhost:5000/{usuarioID} (Apartir do ID é possível alterar as informações do usuário como nome, nick, email) 
  -DELETE -> localhost:5000/{usuarioID} (Apartir do ID, o usuário correspondente é apagado do DB)
 ```
### Login - >
  ```
  -POST -> localhost:5000/login (recebe o body do request (email, senha)) 
 ```
 ### Usuario - >
  ```
  -GET ->
{
    localhost:5000/publicacoes (através do user id do token, é procurado as publicações que estejam associados a rede do usuário)
    localhost:5000/publicacoes/{publicacaoId} (através do ID no parametro é feito a busca especifica da publicação)
    localhost:5000/usuarios/{usuarioId}/publicacoes (através do usuario id da url, pesquisa todas suas publicações em ordem de data)
  
}

  -POST ->
{
    localhost:5000/publicacoes (recebe o body do request (titulo, conteudo) )
    localhost:5000/publicacoes/{publicacaoId}/curtir (recebe o parametro da url e apartir do ID da publicacao é feito a curtida no post)
    localhost:5000//publicacoes/{publicacaoId}/descurtir (recebe o parametro da url e apartir do ID da publicacao é feito a descurtida no post)
}

  -PUT -> localhost:5000/publicacoes/{publicacaoId} (Apartir do ID é possível alterar as informações da publicação como titulo, conteudo) 
  -DELETE -> localhost:5000/publicacoes/{publicacaoId} (Apartir do ID, o usuário correspondente é apagado do DB)
 ```

  
## :arrows_counterclockwise: Dependências
```github.com/joho/godotenv```: go get github.com/joho/godotenv <a href="https://github.com/joho/godotenv">[how-to-install]</a><br>
```github.com/gorilla/mux```: go get github.com/gorilla/mux <a href="https://github.com/gorilla/mux">[how-to-install]</a><br>
```github.com/go-sql-driver/mysql```: go get github.com/go-sql-driver/mysql <a href="https://github.com/go-sql-driver/mysql">[how-to-install]</a><br>
```github.com/badoux/checkmail```: go get github.com/badoux/checkmail <a href="https://github.com/badoux/checkmail">[how-to-install]</a><br>

  </ul>
</details>



## 🔗 Link de download
```gh repo clone Ital023/Italobook``` <a href="https://github.com/Ital023/Italobook/archive/refs/heads/main.zip">[Click here]</a><br>

## 🤝 Colaboradores

Agradecemos às seguintes pessoas que contribuíram para este projeto:

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Ital023" title="Github do Ítalo Miranda">
        <img src="https://avatars.githubusercontent.com/u/113559117?v=4" width="100px;" alt="Foto do Ítalo Miranda no GitHub"/><br>
        <sub>
          <b>Ítalo Miranda</b>
        </sub>
      </a>
    </td>
  </tr>
</table>
