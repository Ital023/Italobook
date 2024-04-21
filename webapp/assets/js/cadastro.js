$("#formulario-cadastro").on("submit", criarUsuario);


function criarUsuario(event) {
    event.preventDefault();
    console.log("Oi");

    if ( $("#senha").val() != $("#confirmarSenha").val() ) {
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: $("#senha").val(),
        }
    }).done(function(){
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro){
        console.log(erro);
        alert("Erro ao cadastrar o usuario!");
    });
}