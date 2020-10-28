const mqtt = require('mqtt');

const client = mqtt.connect('tcp://localhost:1883',{
    username: 'myuser',
    password: 'mypassword'
  });

client.on('connect', () => {
    setInterval(() => {
        client.publish('lol', 'worksing');
        console.log("sdkjngk");

        client.publish('bar', 'working too :(');
    }, 500);

});