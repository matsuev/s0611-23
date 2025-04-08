// Use WebSocket transport endpoint.
const client = new Centrifuge('ws://localhost:8080/centrifugo/connection/websocket', {
   data: {
      username: "alex",
      password: "tuytutuytu"
   }
});

// Allocate Subscription to a channel.
const sub = client.newSubscription('news');

// React on `news` channel real-time publications.
sub.on('publication', function (ctx)
{
   console.log(ctx.data);
});

// Trigger subscribe process.
sub.subscribe();

// Trigger actual connection establishement.
client.connect();



