--migrate-up--
CREATE TABLE another(
    id int auto_increment not null primary key
)

--migrate-down--
DROP TABLE another