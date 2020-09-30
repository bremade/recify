<script>
  import { Row, Col } from 'sveltestrap';
  import Star from "svelte-material-icons/Star.svelte";

  export let rating = 0;

  let maxStars = 5;
  if (0 > rating || rating > maxStars) {
    rating = 0;
  }

  export let size = "22px";
  let colorEmpty = "#cfcdcd";
  let colorFull = "#2196f3";

  let xs = 12/maxStars;
  console.log(xs);
</script>

<Row class="alignRight">
  {#each Array(rating) as _}
    <Col xs={xs} class="mobileRating">
      <Star color={colorFull} {size}/>
    </Col>
  {/each}
  {#each Array(maxStars - rating) as _}
  <Col xs={xs} class="mobileRating">
    <Star color={colorEmpty} {size}/>
  </Col>
  {/each}
</Row>
