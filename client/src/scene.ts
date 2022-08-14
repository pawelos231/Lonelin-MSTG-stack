// @ts-nocheck
import * as THREE from 'three';

import { EffectComposer } from "/node_modules/three/examples/jsm/postprocessing/EffectComposer.js";
import { RenderPass } from "/node_modules/three/examples/jsm/postprocessing/RenderPass.js";
import { UnrealBloomPass } from "/node_modules/three/examples/jsm/postprocessing/UnrealBloomPass.js";


export const createScene = (el: any) => {
	const scene: THREE.Scene = new THREE.Scene();
	const geometry: THREE.TorusGeometry = new THREE.TorusGeometry(0.7, 0.2, 16, 100);
	
	const pointLight: THREE.PointLight = new THREE.PointLight(0xffffff, 0.1);
	pointLight.position.x = 2;
	pointLight.position.y = 3;
	pointLight.position.z = 10;
	pointLight.intensity = 0.07;
	let Particle: any;
	function addParticle() : {
		Particle: THREE.Mesh<THREE.SphereGeometry, THREE.MeshStandardMaterial>;
		material: THREE.MeshStandardMaterial;
	} {
	  let [x, y, z]: number = Array(3)
		.fill()
		.map(() => THREE.MathUtils.randFloat(-100, 80));
	
	  let geometry: THREE.SphereGeometry = new THREE.SphereGeometry(0.05, 24, 24);
	  let material: THREE.MeshStandardMaterial = new THREE.MeshStandardMaterial({
		color: 0xfdb813,
		emissive: "#FF0000",
		emissiveIntensity: 13,
	  });
	  Particle = new THREE.Mesh(geometry, material);
	  
	  if (Math.abs(x) > 100 || Math.abs(y) > 100 || Math.abs(z) > 100) {
		geometry = new THREE.SphereGeometry(0.2, 24, 24);
		Particle = new THREE.Mesh(geometry, material);
	  }
	  Particle.position.set(x /1.5, y / 1.1, z );
	
	  scene.add(Particle);
	  scene.add(pointLight);
	  return { Particle, material };
	}
	
	const material: THREE.PointsMaterial = new THREE.PointsMaterial({
	  size: 0.005,
	  color: "#0000ff",
	});

	interface itemsParticles {
		Particle: THREE.Mesh<THREE.SphereGeometry, THREE.MeshStandardMaterial>;
		material: THREE.MeshStandardMaterial;
	}
	var arrOfParticles: itemsParticles[] = [];

	Array(1350)
	  .fill()
	  .forEach((item: itemsParticles) => {
		arrOfParticles.push(addParticle());
	  });
	

	const sphere: THREE.Points<THREE.TorusGeometry, THREE.PointsMaterial> = new THREE.Points(geometry, material);
	
	const sizes: {
		width: number;
		height: number;
	} = {
	  width: window.innerWidth,
	  height: window.innerHeight,
	};
	

	window.addEventListener("resize", () => {
	  sizes.width = window.innerWidth;
	  sizes.height = window.innerHeight;
	
	  camera.aspect = sizes.width / sizes.height;
	  camera.updateProjectionMatrix();
	
	  renderer.setSize(sizes.width, sizes.height);
	  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
	});
	
	
	const camera: THREE.PerspectiveCamera = new THREE.PerspectiveCamera(
	  75,
	  sizes.width / sizes.height,
	  0.1,
	  500
	);
	camera.position.x = 0;
	camera.position.y = 0;
	camera.position.z = 0.5;
	scene.add(camera);
	
	const renderer: THREE.WebGLRenderer = new THREE.WebGLRenderer({
	  canvas: el,
	  alpha: true,
	});
	const renderScene: any = new RenderPass(scene, camera);
	const bloomPass: any = new UnrealBloomPass(
	  new THREE.Vector2(window.innerWidth, window.innerHeight),
	  1.5,
	  0.4,
	  0.85
	);
	bloomPass.threshold = 0.5;
	bloomPass.strength = 3; //intensity of glow
	bloomPass.radius = 0.5;
	const bloomComposer: any = new EffectComposer(renderer);
	bloomComposer.setSize(window.innerWidth, window.innerHeight);
	bloomComposer.renderToScreen = true;
	bloomComposer.addPass(renderScene);
	bloomComposer.addPass(bloomPass);
	
	renderer.setSize(sizes.width, sizes.height);
	renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
	
	const clock:THREE.Clock = new THREE.Clock();
	
	const tick = () => {
	  const elapsedTime: number = clock.getElapsedTime();
	
	  arrOfParticles.forEach((item: itemsParticles, i: number) => {
		item.Particle.position.z = item.Particle.position.z + 0.1055;
		item.material.emissive.r = 1 * Math.abs(item.Particle.position.z);
		item.material.emissive.g = 3.86 * Math.abs(1 / item.Particle.position.z);
		item.material.emissive.b = 3.86 * Math.abs(1 / item.Particle.position.z);

		if (item.Particle.position.z > 1.5) {
		  item.Particle.geometry.dispose();
		  item.Particle.material.dispose();
		  scene.remove(item.Particle);
		  arrOfParticles.splice(i, 1);
		  arrOfParticles.push(addParticle());
		}

	  });

	  sphere.rotation.y = 0.5 * elapsedTime;

	  renderer.render(scene, camera);
	  bloomComposer.render();
	  window.requestAnimationFrame(tick);
	};	
	tick();	
}



