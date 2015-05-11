var camera, scene, renderer;
var geometry, material, mesh;
var controls;

var objects = [];

var raycaster;

var blocker = document.getElementById( 'blocker' );
var instructions = document.getElementById( 'instructions' );

init();
animate();

function init() {
	var dae;
	camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 1, 1000 );

	scene = new THREE.Scene();
	//scene.fog = new THREE.Fog( 0xFFFFFF, 0, 1024);

	var light = new THREE.HemisphereLight( 0xFFFFFF, 0xFFFFFF, 1.0 );
	light.position.set( 0.5, 10, 0.75 );
	scene.add( light );

	dirLight = new THREE.DirectionalLight( 0xffffff, 1 );
	dirLight.color.setHSL( 0.1, 1, 0.95 );
	dirLight.position.set( -1, 150, 1 );
	dirLight.position.multiplyScalar( 50 );
	scene.add( dirLight );

    controls = new THREE.OrbitControls( camera );
    controls.damping = 0.2;
    controls.addEventListener( 'change', render );

	raycaster = new THREE.Raycaster( new THREE.Vector3(), new THREE.Vector3( 0, - 1, 0 ), 0, 30 );

	//renderer = new THREE.WebGLRenderer({ antialias: true });
	renderer = new THREE.WebGLRenderer();
	renderer.setClearColor( 0xFFFFFF);
	renderer.setSize( window.innerWidth, window.innerHeight );

	document.body.appendChild( renderer.domElement );

	//

	window.addEventListener( 'resize', onWindowResize, false );

}

function onWindowResize() {

	camera.aspect = window.innerWidth / window.innerHeight;
	camera.updateProjectionMatrix();

	renderer.setSize( window.innerWidth, window.innerHeight );

}

function animate() {

	requestAnimationFrame( animate );

	//controls.isOnObject( false );

	//raycaster.ray.origin.copy( controls.getObject().position );
	//raycaster.ray.origin.y -= 10;

	//var intersections = raycaster.intersectObjects( objects );

	//if ( intersections.length > 0 ) {

	//	controls.isOnObject( true );

	//}

	//controls.update();

	renderer.render( scene, camera );

}

            function render() {

                                renderer.render( scene, camera );
                                                stats.update();

                                                            }

