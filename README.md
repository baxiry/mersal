
<div style="color:yellow">

# *مرسال*
## *mersal*
##### Instant messaging server, easy to use as a server or as a library

</div>    
how to use ?

create a websocket on browser and send message as json :

to subscribe in a topic (channel) send :
```json
{
   "event":"subscribe",
   "channel":"my-channel-id"
}
```
event must be : ```subscribe```, ```unsubscribe```, ```message```,

Later we will add events:  ```reseive```, and```seen``` in order to achieve  quality of service. "qos".

to send message to channel/topic:
```json
{
   "event" : "message",
   "channel" : "my-channel-123",
   "data" : "hi frends"
}
```

then all client subscribe with "my-channel-123" will be receive "hi frinds" message

