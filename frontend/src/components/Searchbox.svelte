<script>
  import { Container, Col, Row } from 'sveltestrap';
  import IconButton from '@smui/icon-button';

  export let isOpen = false;

  // Extract random recipe id
  //const id = 
  let recipeIds = [];
  let randomId = "";

  fetch('/api/v1/recipe')
  .then(resp => resp.json())
  .then(jsonData => {
    for (const key in jsonData) {
      recipeIds.push(jsonData[key].id);
    }
    let index = Math.floor(Math.random()*(recipeIds.length - 0) + 0);
    randomId = recipeIds[index];
  });
</script>

<div class="search-box">
  <input class="input-field" id="searchRecipe" autocomplete="off" maxlength="100" placeholder="e.g. Pizza, Quesadilla"/>
  <IconButton class="material-icons" aria-label="Open recipe search">search</IconButton>
  <IconButton class="material-icons" aria-label="Random recipe" href={"/recipe/" + randomId}>casino</IconButton>
  {#if isOpen}
    <IconButton class="material-icons" aria-label="Close search bar" on:click={() => isOpen = !isOpen}>close</IconButton>
  {/if}
</div>
