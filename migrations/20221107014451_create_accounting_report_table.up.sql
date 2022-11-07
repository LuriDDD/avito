CREATE TABLE accounting_report (
  id_service bigserial not null,
  month integer not null,
  year integer not null,
  funds bigint not null
);