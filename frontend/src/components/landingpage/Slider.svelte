<script>
  import Carousel from '@beyonk/svelte-carousel';
  import { Swiper, SwiperSlide } from 'svelte-swiper';
  import Tag from './Tag.svelte';

  export let tags;

  const options = {
    slidesPerView: 'auto',
    spaceBetween: 10,
    freeMode: true,
  };
</script>

<Swiper className="centered-auto" {options} id="test">
  {#each tags as tag}
    <SwiperSlide className="swiper-slide">
      <Tag tag={tag}/>
    </SwiperSlide>
  {/each}
</Swiper>

<!--
<style>
	.slider {
		font-family: sans-serif;
		scroll-snap-type: x mandatory;	
		display: flex;
		-webkit-overflow-scrolling: touch;
		overflow-x: scroll;
	}
	section {
    min-width: 100%;
    height: 50px;
		scroll-snap-align: start;
		text-align: center;
		position: relative;
	}
	h1	{
		position: absolute;
    font-size: 30px;
		text-align: center;
		color: black;
	}
</style>

<div class="slider">
  {#each sections as section}
    <section>
      <h1>{section}</h1>
    </section>
  {/each}
</div>


<style>
	.slide-content {
    width: 60%;
    text-align: center;
	}
</style>

<Carousel perPage="8" dots="false" controls="false">
  {#each tags as tag}
    <div class="slide-content" >
      <Tag tag={tag}/>
    </div>
  {/each}
</Carousel>
-->
