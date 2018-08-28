
<!DOCTYPE html>
<!--
To change this license header, choose License Headers in Project Properties.
To change this template file, choose Tools | Templates
and open the template in the editor.
-->
<html>
    <head>
        <meta charset="UTF-8" name="viewport" content="width=device-width, initial-scale=1">
        <title>Puma Calendar</title>
        <link rel="stylesheet" href="./css/estiloDecimaPantalla.css">
    </head>
    <body>
        
        <div class="div1">
            <a href="TerceraPantalla.php">
                <img class="img1" src="./img/logoSistemas.jpg" alt="Logo Sistemas">
</a>
</div>

<div class="div2">
<h1>Facultad de Sistemas</h1>
</div>

<div class="div3">
<div class="div menu">
 <ul>
   <li><a href="TerceraPantalla.php">Inicio</a></li>
   <li><a href="SeptimaPantalla.php">Eventos</a></li>
   <li><a href="NovenaPantalla.php">Configuracion</a></li>
 </ul>
 </div>
</div>
        <form action="crear.php" method="POST"  NAME="form">
<div class="div6">
    <input class="nom" type="text" name="NombreDeEvento" placeholder="Ingresa Nombre del Evento" id="NombreDeEvento" required>
</div>

<div class="div7">
    <input class="fecha1" type="date" name="FechaDeEvento" id="FechaDeEvento" required>
</div>

<div class="div8">
    <img class="img2" src="./img/ImagenCalendario.png" alt="Imagen Calendario">
</div>

<div class="div9">
    <textarea class="a1" rows="15" cols="200" id="DescripcionDeEvento" name="DescripcionDeEvento" placeholder="DescripcionDeEvento" required></textarea>
</div>

<div class="div10">
<textarea class="a2" rows="15" cols="120" id="UbicacionDeEvento" name="UbicacionDeEvento" placeholder="Ubicación del Evento" required></textarea>
</div>

<div class="div15">
<p>Selecciona una Foto del Evento</p>
</div>


<div class="div14">
<input class="bot4" type="file" name="botonSubirFotoEvento" value="Ingresa Foto del Evento">
</div>

<div class="div13">
<input class="bot4" type="submit" name="botonConfirmarEvento" value="Confirmar Nuevo Evento">
</div>

</form>
    </body>
</html>