# GolangPlacetoPay
Una  librería instalable que implementa los métodos necesarios para el uso del botón de pago del PlacetoPay

### Librería para facilitar del método de pago de PlactoPay
Instalar la librería con el siguiente comando

```
go get github.com/dmjacas/placetoplay 
```

### Instalación

Inicializar la librería utilizando las configuraciones del proyecto.

```
URLPayment (string) Url de pago de PlacetoPay
Secret (string) llave secreta de PlacetoPay
Login (string) Usuario de PlacetoPay
Charset (string) Codificación de la base de datos
Dialect (string) Dialecto de la base de datos (MySql, Postgres)
DBName (string) Nombre de la base de datos
DBPassword (string) Contrasena de la base de datos
DBUsername (string) Usuario de la base de datos
Expiration (int) tiempo de expiración de la solicitud de pago en minutos
```
```
placetopay.Config(URLPayment,Secret, Login, Charset, Dialect,DBName, DBPassword, DBUsername,Expiration)

```

## CreateRequest

Crear una nueva solicitud de pago

Para crear una nueva solicitud de pago se llama al método

```
buyer := &placetopay.Person{
    // Inicializar los valores
}
payment:= &placetopay.PaymentRequest{
    // Inicializar los valores
}
fields:= []*placetopay.NameValuePair{
    // Inicializar los valores
}
data := &placetopay.RedirectRequest{
		Buyer:      buyer,
		Locale:     "es_EC",
		IPAddres:   c.ClientIP(),
		UserAgent:  "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0",
		Payment:    payment,
		ReturnURL:  "http://localhost/redirection/",
		SkipResult: false,
		Fields:     fields,
}
response, err := placetopay.CreateRequest(data)

```
La variable ```response``` es del tipo placetopay.RedirectInformation


## GetRequestInformation 

Solicita la información de un pago realizado
 

```
requestID identificador de la solicitud de pago

response, err := placetopay.GetRequestInformation(requestID)

```
La variable ```response``` es del tipo placetopay.RedirectInformation

## ReversePaymemt 

Método para cancelar un pago
```
internalReference referencia interna del pago

response, err := placetopay.ReversePaymemt(internalReference)

```
La variable ```response``` es del tipo placetopay.ReverseResponse

