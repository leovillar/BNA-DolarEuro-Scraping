# API y webhook scraping valor Dolar/Euro del BNA

![](https://github.com/leovillar/bna-dolareuro-scraping/actions/workflows/go.yml/badge.svg)

## Objetivo:
El objetivo de este proyecto es obtener el valor del dolar y euro del BNA y disponibilizarlo como REST API en GET /cotizacion y enviarlo a una webhook.

### API REST
### GET /cotizacion
Response 200:
```json
{
  "dolarCompra": 134.25,
  "dollarVenta": 142.25,
  "euroCompra": 135,
  "euroVenta": 143
}
```

### Webhook
Request:
```json
{
  "dolarCompra": 134.25,
  "dollarVenta": 142.25,
  "euroCompra": 135,
  "euroVenta": 143
}
```
Para habilitar el webhook se requiere la configuracion de la variable `WEBHOOK_ENDPOINT` en el archivo `.env`, y la
respectiva habilitacion de la variable `CRON_NOTIFICATION_ENABLED=true` en el mismo archivo.

El delay de la notificacion es configurable en la variable `CRON_NOTIFICATION` en el archivo `.env` y tiene el formato de cron de linux.

## Para correr con docker compose
Tener instalados `docker` y `docker-compose` o `docker compose`.

Bajar en tu pc local el docker-compose.yml y el .env en el mismo directorio.
Modificar según conveniencia el archivo .env, puerto donde corre la API, habilitar o no el webhook configurando el endpoint de webhook.

```...
   API_SERVER_PORT=4000
   CRON_NOTIFICATION_ENABLED=false
   #At minute 1 past every hour from 10 through 15 on every day-of-week from Monday through Friday.
   CRON_NOTIFICATION=1 10-15 * * 1-5
   WEBHOOK_ENDPOINT=http://localhost:8000
```

Luego de modificar el archivo .env, ejecutar el siguiente comando:
```
docker-compose up -d
```

### Funciona perfectamente bajo entornos linux, no lo probado en windows, pero estimo que no habría problemas.

