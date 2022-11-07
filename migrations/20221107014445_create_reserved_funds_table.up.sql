CREATE TABLE reserved_funds (
  id_user bigserial not null,
  id_service bigserial not null,
  id_order bigserial not null primary key,
  price bigint not null
);