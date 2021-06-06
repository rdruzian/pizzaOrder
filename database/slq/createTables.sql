create table if not exists address(
  addressID SERIAL not null primary key,
  type varchar(15) not null,
  publicPlace varchar(100) not null,
  number integer not null,
  complement varchar(100)
);

create table if not exists customer(
    customerId SERIAL not null primary key,
    name varchar(150) not null,
    favorite integer,
    login varchar(50) not null,
    pass varchar(25) not null
);

create table if not exists order(
    orderID SERIAL not null primary key,
    customerID integer not null,
    flavor integer not null,
    quantity integer not null,
    details varchar(50),
    price decimal(4,2) not null,
    edge varchar(50),
    dateTime date not null,
    status integer not null,
    totalPrice decimal(5,2) not null
);

create table if not exists pizza(
    pizzaID SERIAL not null primary key,
    nameFlavor varchar(50) not null,
    price decimal(4,2) not null
);

create table if not exists customerAddress(
    customerID integer not null primary key references customer,
    addressID integer not null references address
);

create table if not exists orderCustomer(
    cusotmerID integer not null primary key references customer,
    orderID integer not null references order,
    total decimal(5,2) not null references order
);

create table if not exists ingredients(
    ingredientID SERIAL not null primary key,
    pizzaID integer not null references pizza
);