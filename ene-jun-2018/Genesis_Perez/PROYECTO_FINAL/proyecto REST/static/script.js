//console.log(document)
var cont = 1;
var aciertos = 0;
document.getElementById("r1").addEventListener("click",enviarp1);


 
function enviarp1(){
    var su_respuesta = document.getElementById("r1").value;
    var xhr = new XMLHttpRequest();
 
    xhr.open("POST", "/validar");
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(su_respuesta);
    //console.log(xhr.responseText);
    //console.log(xhr);
 
    xhr.onreadystatechange = function(){
        if(xhr.readyState == 4 && xhr.status == 200){
            var resp = JSON.parse(xhr.response)
            document.getElementById("respuesta").innerHTML = resp;
            if (document.getElementById("respuesta").textContent == "Respuesta Correcta!"){
            aciertos = aciertos + 1;
            }
            console.log(aciertos);
            document.getElementById("respuesta").style.visibility = "visible";
            document.getElementById("r1").disabled = true;
            document.getElementById("r2").disabled = true;
            document.getElementById("r3").disabled = true;
            document.getElementById("r4").disabled = true;
            //document.getElementById("respuesta").innerHTML = xhr.responseText;
        }
    }
}


document.getElementById("r2").addEventListener("click",enviarp2);
 
function enviarp2(){
    var su_respuesta = document.getElementById("r2").value;
    var xhr = new XMLHttpRequest();
 
    xhr.open("POST", "/validar");
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(su_respuesta);
    //console.log(xhr);
    //console.log(xhr);
 
    xhr.onreadystatechange = function(){
        if(xhr.readyState == 4 && xhr.status == 200){
            var resp = JSON.parse(xhr.response)
            document.getElementById("respuesta").innerHTML = resp;
             if (document.getElementById("respuesta").textContent == "Respuesta Correcta!"){
            aciertos = aciertos + 1;
            }
            console.log(aciertos);
            document.getElementById("respuesta").style.visibility = "visible";
            document.getElementById("r1").disabled = true;
            document.getElementById("r2").disabled = true;
            document.getElementById("r3").disabled = true;
            document.getElementById("r4").disabled = true;
            //document.getElementById("respuesta").innerHTML = xhr.responseText;
            
        }
    }
}

document.getElementById("r3").addEventListener("click",enviarp3);
 
function enviarp3(){
    var su_respuesta = document.getElementById("r3").value;
    var xhr = new XMLHttpRequest();
 
    xhr.open("POST", "/validar");
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(su_respuesta);
    //console.log(xhr);
    console.log(xhr);
 
    xhr.onreadystatechange = function(){
        if(xhr.readyState == 4 && xhr.status == 200){
            var resp = JSON.parse(xhr.response)
            document.getElementById("respuesta").innerHTML = resp;
            if (document.getElementById("respuesta").textContent == "Respuesta Correcta!"){
            aciertos = aciertos + 1;
            }
            console.log(aciertos);
            document.getElementById("respuesta").style.visibility = "visible";
            document.getElementById("r1").disabled = true;
            document.getElementById("r2").disabled = true;
            document.getElementById("r3").disabled = true;
            document.getElementById("r4").disabled = true;
            //document.getElementById("respuesta").innerHTML = xhr.responseText;
            
        }
    }
}

document.getElementById("r4").addEventListener("click",enviarp4);
 
function enviarp4(){
    var su_respuesta = document.getElementById("r4").value;
    var xhr = new XMLHttpRequest();
 
    xhr.open("POST", "/validar");
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(su_respuesta);
    //console.log(xhr);
    //console.log(xhr);
 
    xhr.onreadystatechange = function(){
        if(xhr.readyState == 4 && xhr.status == 200){
            var resp = JSON.parse(xhr.response)
            document.getElementById("respuesta").innerHTML = resp;
             if (document.getElementById("respuesta").textContent == "Respuesta Correcta!"){
            aciertos = aciertos + 1;
            }
            
            console.log(aciertos);
            document.getElementById("respuesta").style.visibility = "visible";
            document.getElementById("r1").disabled = true;
            document.getElementById("r2").disabled = true;
            document.getElementById("r3").disabled = true;
            document.getElementById("r4").disabled = true;
            //document.getElementById("respuesta").innerHTML = xhr.responseText;
            
        }
    }
}
document.getElementById("sig").addEventListener("click",siguiente);

    function siguiente() {
            document.getElementById("respuesta").style.visibility = "hidden";
            document.getElementById("r1").disabled = false;
            document.getElementById("r2").disabled = false;
            document.getElementById("r3").disabled = false;
            document.getElementById("r4").disabled = false;
    cont = cont + 1;
    if (cont == 2) {
    document.getElementById("Pregunta").innerHTML = "¿Cuál es el tercer planeta del sistema solar?";
    document.getElementById("r1").value = "Saturno";
    document.getElementById("r2").value = "Urano";
    document.getElementById("r3").value = "Tierra";
    document.getElementById("r4").value = "Mercurio";
    }
     if (cont == 3) {
    document.getElementById("Pregunta").innerHTML = "¿En qué año fue la independencia de Mexico?";
    document.getElementById("r1").value = "2018";
    document.getElementById("r2").value = "200 a.c.";
    document.getElementById("r3").value = "1910";
    document.getElementById("r4").value = "1810";
    }
     if (cont == 4) {
    document.getElementById("Pregunta").innerHTML = "¿Cual es el animal mas grande del mundo?";
    document.getElementById("r1").value = "Elefante";
    document.getElementById("r2").value = "Jirafa";
    document.getElementById("r3").value = "Ballena Azul";
    document.getElementById("r4").value = "Ballena gris";
    }
     if (cont == 5) {
    document.getElementById("Pregunta").innerHTML = "¿Cual es la capital de Uruguay?";
    document.getElementById("r1").value = "Montevideo";
    document.getElementById("r2").value = "Lima";
    document.getElementById("r3").value = "Bogotá";
    document.getElementById("r4").value = "Londres";
    }
    if (cont == 6) {
    
    //document.getElementById("Pregunta").innerHTML = "¿Cual es la capital de Uruguay?";
    document.getElementById("Pregunta").style.visibility = "hidden"; 
    document.getElementById("r1").style.visibility = "hidden"; 
    document.getElementById("r2").style.visibility = "hidden"; 
    document.getElementById("r3").style.visibility = "hidden"; 
    document.getElementById("r4").style.visibility = "hidden";
    document.getElementById("sig").style.visibility = "hidden";  
    document.getElementById("resultado").style.visibility = "visible";
    document.getElementById("respuesta").style.visibility = "hidden";
    }
if (aciertos > 3){
document.getElementById("resultado").innerHTML = "GANASTE!";
} else {
document.getElementById("resultado").innerHTML = "PERDISTE!";
}               
    }

