      url = 'ws://localhost:8090/ws';
      c = new WebSocket(url);
      
      send = function(data){
        c.send(data)
      }

      c.onmessage = function(msg){
        //console.log(msg)
      }

      c.onopen = function(){
        setInterval( 
          function(){ send(JSON.stringify([controls.getObject().position,controls.getObject().quaternion]))}
        , 100 )
      }