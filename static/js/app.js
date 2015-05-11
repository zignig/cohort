			var container, stats;

			var camera, controls, scene, renderer;

			init();
			render();

			function animate() {

				requestAnimationFrame(animate);
				controls.update();

			}

			function init() {

				camera = new THREE.PerspectiveCamera( 60, window.innerWidth / window.innerHeight, 1, 1000 );
				camera.position.z = 10;
				camera.position.x = 10;

				controls = new THREE.OrbitControls( camera );
				controls.damping = 0.2;
				controls.addEventListener( 'change', render );

				scene = new THREE.Scene();
				//scene.fog = new THREE.FogExp2( 0xcccccc, 0.002 );

				// lights

				light = new THREE.DirectionalLight( 0xffffff );
				light.position.set( 1, 1, 1 );
				scene.add( light );

				light = new THREE.DirectionalLight( 0x002288 );
				light.position.set( -1, -1, -1 );
				scene.add( light );

				light = new THREE.AmbientLight( 0x222222 );
				scene.add( light );


				// renderer

				renderer = new THREE.WebGLRenderer( { antialias: false } );
				renderer.setClearColor( 0xAAAAAA );
				renderer.setPixelRatio( window.devicePixelRatio );
				renderer.setSize( window.innerWidth, window.innerHeight  - 20);

				container = document.getElementById( 'container' );
				container.appendChild( renderer.domElement );

				//

				window.addEventListener( 'resize', onWindowResize, false );

				animate();

			}

			function onWindowResize() {

				camera.aspect = window.innerWidth / window.innerHeight;
				camera.updateProjectionMatrix();

				renderer.setSize( window.innerWidth, window.innerHeight );

				render();

			}

			function render() {

				renderer.render( scene, camera );

			}
