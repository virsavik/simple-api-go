alter table invoice_product drop constraint fk_invoice_product_invoices;
alter table invoice_product drop constraint fk_invoice_product_products;
alter table invoices drop constraint fk_invoice_user;
alter table users drop constraint fk_user_role;

drop table invoice_product;
drop table invoices;
drop table products;
drop table users;
drop table roles;