USE italobook;


insert into usuarios (nome, nick, email, senha) 
values 
("Davi Cotting","davizinC","davithomeny@gmail.com","$2a$10$5b/fhHw8rOrtyz3PX.rA7.MniUatj64ldh9iUd.c0NXXzIfOH08I2"),
("Cristion","cristionmc","cristionmc@gmail.com","$2a$10$5b/fhHw8rOrtyz3PX.rA7.MniUatj64ldh9iUd.c0NXXzIfOH08I2"),
("Antonio Jose","joseantonio","antoniojose@gmail.com","$2a$10$5b/fhHw8rOrtyz3PX.rA7.MniUatj64ldh9iUd.c0NXXzIfOH08I2"),
("Fabiano Braga","fabiaanobraaga","fabianobraga@gmail.com","$2a$10$5b/fhHw8rOrtyz3PX.rA7.MniUatj64ldh9iUd.c0NXXzIfOH08I2");

insert into seguidores(usuario_id, seguidor_id)
values
(1,2)
(3,1)
(1,3)