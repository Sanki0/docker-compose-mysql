CREATE TABLE `students`
(
    id   bigint auto_increment,
    nombre varchar(255) NOT NULL,
    dni varchar(255) NOT NULL,
    direccion varchar(255) NOT NULL,
    fecha_nacimiento varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `students` (`nombre`, `dni`,`direccion`,`fecha_nacimiento`) VALUES 
       ('Jose', '12345678', 'Calle falsa 123', '2020-01-01'),
       ('Juan', '87654321', 'Calle verdadera 456', '2020-01-01'),
       ('Adrian','1234124', 'Valle', '2020-01-01'),
       ('John', '1234124', 'Valle', '2020-01-01'),
       ('Mary', '1234124', 'Valle', '2020-01-01');

CREATE TABLE `tipoCliente`
(
    idTipoCliente   bigint auto_increment,
    nombre varchar(255) NOT NULL,
    detalle varchar(255) NOT NULL,
    PRIMARY KEY (`idTipoCliente`)
);

INSERT INTO `tipoCliente` (`nombre`, `detalle`) VALUES 
       ('Cliente', 'Detalle para Cliente'),
       ('Proveedor', 'Detalle para Proveedor'),
       ('Empleado', 'Detalle para Empleado');

CREATE TABLE `cliente`
(
    idCliente   bigint auto_increment,
    nombre varchar(255) NOT NULL,
    dni char(8) NOT NULL,
    telefono varchar(255) NOT NULL,
    correo varchar(255) NOT NULL,
    estado char(1) NOT NULL,
    idTipoCliente bigint NOT NULL,
    PRIMARY KEY (`idCliente`),
    FOREIGN KEY (`idTipoCliente`) REFERENCES `tipoCliente` (`idTipoCliente`)
);

INSERT INTO `cliente` (`nombre`, `dni`,`telefono`,`correo`,`estado`,`idTipoCliente`) VALUES 
       ('sanki0', '12345678', '9886472', 'sebas@gmail.com', 'A', 1),
       ('sanki1', '98565354', '4586472', 'sanki1@gmail.com', 'A', 1),
       ('sanki2', '12676578', '9872471', 'sanki2@gmail.com', 'A', 2);
