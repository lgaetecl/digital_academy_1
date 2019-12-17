function add(numeros) {
    if (!numeros.length) return 0;
    let separador;
    let numToSum;
    if (numeros[0] === "/" && numeros[1] === "/") {
        separador = numeros[2];
        numToSum = numeros.split("\n")[1];

    } else {
        separador = ",";
        numToSum = numeros.split("\n").join(separador);
    }
    
    let numbers = numToSum.split(separador);
    let suma = 0;
    numbers.forEach(element => {
        if (parseInt(element) < 0) throw new Error("numeros negativos no soportados")
        suma += parseInt(element);
    });
    return suma;
}

function prueba1() {
    let r = add("");

    console.log("llamando con vacio", r, r === 0);
}

function prueba2() {
    let r = add("1");

    console.log("llamando con 1", r, r === 1);
}

function prueba3() {
    let r = add("1,2");

    console.log("llamando con 1,2", r, r === 3);
}

function prueba4() {
    let r = add("1,2,3,4,5");

    console.log("llamando con 1,2,3,4,5", r, r === 15);
}

function prueba5() {
    let r = add("1,2\n3");

    console.log("llamando con 1,2\\n3", r, r === 6);
}

function prueba6() {
    let r = add("1,2\n3,4\n5");

    console.log("llamando con 1,2\\n3,4\\n5", r, r === 15);
}

function prueba7() {
    try {
        let r = add("1,2,-3");
        console.log("Faloo la prueba 7")
    }
    catch(err) {
        console.log("llamando con 1,2,-3", err.message, err.message === "numeros negativos no soportados");
    }
}

function prueba8() {
    let r = add("//;\n1;2;3");

    console.log("llamando con //;\\n1;2;3", r, r === 6);
}

prueba1();
prueba2();
prueba3();
prueba4();
prueba5();
prueba6();
prueba7();
prueba8();