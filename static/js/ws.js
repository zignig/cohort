      url = 'ws://'+location.host+'/ws';
	
      c = new WebSocket(url);
      var interval ; 
      send = function(data){
        c.send(data)
      }

      c.onmessage = function(msg){
		var m = JSON.parse(msg.data);
		console.log(m);
		switch(m.class){
			case "loader":
				// load object
				loader = new THREE.ColladaLoader();
				loader.options.convertUpAxis = true;
				loader.load( 'asset/'+m.message.path, function ( collada ) {
					dae = collada.scene;
					dae.traverse( function ( child ) {
							if ( child instanceof THREE.SkinnedMesh ) {
								var animation = new THREE.Animation( child, child.geometry.animation );
								animation.play();
							}
						} );
					dae.position.x = m.message.Pos.X;
					dae.position.y = m.message.Pos.Y;
					dae.position.z = m.message.Pos.Z;
					//console.log(dae)
					dae.updateMatrix();
					scene.add(dae);
					} );
					break;
			case "floor":
				
				// floor
				// note: 4x4 checkboard pattern scaled so that each square is 25 by 25 pixels.
				floorTexture = new THREE.ImageUtils.loadTexture( 'static/images/dirt.jpg' );
				floorTexture.wrapS = floorTexture.wrapT = THREE.RepeatWrapping; 
				floorTexture.repeat.set( 10, 10 );
				// DoubleSide: render texture on both sides of mesh
				floorMaterial = new THREE.MeshBasicMaterial( { map: floorTexture, side: THREE.DoubleSide } );
				floorGeometry = new THREE.PlaneBufferGeometry(m.message.Size, m.message.Size, 1, 1);
				floor = new THREE.Mesh(floorGeometry, floorMaterial);
				floor.position.x = m.message.Pos.X
				floor.position.y = m.message.Pos.Y
				floor.position.z = m.message.Pos.Z


				floor.rotation.x = Math.PI / 2;
				console.log(floor);
				scene.add(floor);
			
			case "tile":
				var onProgress = function ( xhr ) {
					if ( xhr.lengthComputable ) {
						var percentComplete = xhr.loaded / xhr.total * 100;
						console.log( Math.round(percentComplete, 2) + '% downloaded' );
					}
				};
				var onError = function ( xhr ) {};
				var loader = new THREE.OBJMTLLoader();
				var path = 'asset/'+m.message.Ref+'/tiles/'+m.message.Name;
				loader.load( path+'.obj', path+'.mtl', function ( object ) {
					object.position.x = m.message.Pos.X;
					object.position.y = m.message.Pos.Y;
					object.position.z = m.message.Pos.Z;
					object.scale.x = 4
					object.scale.y = 4
					object.scale.z = 4
					
					scene.add( object );

				}, onProgress, onError );
		}
      }

      c.onopen = function(){
        interval = setInterval( 
          function(){ send(JSON.stringify({"class":"location","message":{"pos":controls.getObject().position,"rot":controls.getObject().quaternion,"uuid":controls.getObject().uuid}}))}
        , 10)
      }
      
	c.onlcose = function(){
		clearInterval(interval);
	}
