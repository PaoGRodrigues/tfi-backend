# Trabajo Final Integrador - Backend

## Configuración del bot de Telegram
Para crear y configurar el bot de Telegram por el cual recibiremos las notificaciones, se siguió la guía presente en [este link](https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1).
1. Instalar Telegram en el celular.
2. Buscar **BotFather** en Telegram.
3. Enviar el mensaje _/start_ para comenzar la conversación.
4. Enviar _/newbot_ para comenzar con la creación del bot.
5. Ingresar el nombre elegido. En este trabajo se usó **SecurityBot**.
6. Ingresar un nombre de usuario, terminado en _bot_. En este trabajo se usó **sec_home_bot**. Una vez ingresado el nombre de usuario, **BotFather** mostrará un mensaje donde se indica que el _bot_ ha sido creado y generará un _token_ que se usará para configurar en el _frontend_ de nuestra aplicación.
7. Una vez creado el _bot_ y configurado el _token_ en el _frontend_, comenzar una conversación con el mismo a través del usuario de Telegram del celular. Esta conversación es interpretada por el _bot_ y enviada al _backend_, lo que permitirá tener identificado el _ID_ del _chat_ al cual se le enviarán las notificaciones a partir de ese instante.