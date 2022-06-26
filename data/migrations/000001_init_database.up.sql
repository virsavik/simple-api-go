create extension if not exists "uuid-ossp";

create table roles (
    name VARCHAR(100) PRIMARY KEY NOT NULL
);

create table users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

create table products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    amount INTEGER NOT NULL,
    slot INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

create table invoices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    storekeeper_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

create table invoice_product (
    invoice_id UUID NOT NULL,
    product_id UUID NOT NULL,
    amount INTEGER NOT NULL,
    unit_price NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

alter table users add constraint fk_user_role foreign key (role_name) references roles(name);

alter table invoices add constraint fk_invoice_user foreign key (storekeeper_id) references users (id);

alter table invoice_product add constraint fk_invoice_product_invoices foreign key (invoice_id) references invoices (id);
alter table invoice_product add constraint fk_invoice_product_products foreign key (product_id) references products (id);
alter table invoice_product add constraint pk_invoice_product primary key (invoice_id, product_id);