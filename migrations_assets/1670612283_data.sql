--migrate-up--
CREATE TABLE nada(
    id int auto_increment not null primary key
)

--migrate-down--
DROP TABLE nada