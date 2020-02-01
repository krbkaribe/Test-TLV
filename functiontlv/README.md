# Test Integraciones TLV por Karla Rivera Bello

## Contexto
Función en GO que lee una cadena de caracteres en formato TLV y retorna 2 parámetros:
	- Un map[string]string con los campos TLV encontrados en la cadena.
	- Mensaje de error.

## Formato de los campos TLV
Cada campo TLV está compuesto por 3 partes:
	- Largo: los 2 primeros caracteres 2 caracteres e indican el largo del 'Valor'.
	- Tipo:  el Tipo tiene un largo de 3 caracteres, donde: 
		* El primer caracter indica el tipo de dato A: Alfanumérico o N: Numérico.
		* Los 2 siguientes caracteres indican la posición de inicio del 'Valor'.
	- Valor: es el valor del campo y su longitud está determinada por el 'Largo'.

## Especificaciones de uso
La función contiene un arreglo de bytes que representa una cadena de texto en el formato TLV descrito anteriormente.
El resultado será un map con los valores de los campos TLV y un mensaje de error.
Si la cadena se encuentra correcta, el mensaje de error se indica como "-":
Los escenarios comunes de error son:
	- Cadena vacía.
	- Campos que no cumplen con el tipo de dato especificado.

## Indicaciones del código
Repositorio Git: https://github.com/krbkaribe/Test-TLV.git
Link Travis CI:  https://travis-ci.com/krbkaribe/Test-TLV
Para compilar, debe seguir los siguientes pasos:
	- En una ventana de símbolo del sistema compile con la herramienta go: %USERPROFILE%\go\src\Test-TLV\functiontlv> go build
	- Ejecútelo para iniciar y ver el resultado: %USERPROFILE%\go\src\Test-TLV\functiontlv>functiontlv
