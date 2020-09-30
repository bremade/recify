import HMR from '@sveltech/routify/hmr';
import App from './App.svelte';
import 'swiper/swiper-bundle.css';

const app = HMR(
	App, 
	{ target: document.body }, 
	'recify'
	)

export default app;
