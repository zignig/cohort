      url = 'ws://'+location.host+'/ws';
      c = new WebSocket(url);
      
      send = function(data){
        c.send(data)
      }

      c.onmessage = function(msg){
		//$("#output").append((new Date())+ " <== "+msg.data+"\n")
        console.log(msg.data);
		switch(msg.data["class"]){
			case "loader":
				// load object
				console.log(msg.data["message"])
				break;
		}
		var loader = new THREE.ColladaLoader();
		loader.options.convertUpAxis = true;
		loader.load( '/static/models/monkey.dae', function ( collada ) {
			dae = collada.scene;
			dae.traverse( function ( child ) {

					if ( child instanceof THREE.SkinnedMesh ) {

						var animation = new THREE.Animation( child, child.geometry.animation );
						animation.play();

					}

				} );
			//dae.scale.x = dae.scale.y = dae.scale.z = 0.02;
			dae.scale.x = dae.scale.y = dae.scale.z = 1;
			dae.position.z = -20
			dae.updateMatrix();
			scene.add(dae);
		} );

      }

      c.onopen = function(){
        setInterval( 
          function(){ send(JSON.stringify({"class":"location","message":{"pos":controls.getObject().position,"rot":controls.getObject().quaternion,"uuid":controls.getObject().uuid}}))}
        , 20)
      }