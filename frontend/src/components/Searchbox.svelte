<script>
  import { Container, Col, Row } from 'sveltestrap';
  import IconButton from '@smui/icon-button';

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
  <form class="search-form">
    <Row>
      <Col>
        <input class="input-field" id="searchRecipe" autocomplete="off" maxlength="100" placeholder="e.g. Pizza, Quesadilla"/>
      </Col>
      <Col class="text-right input-buttons">
        <IconButton class="material-icons" aria-label="Open recipe search">search</IconButton>
        <IconButton class="material-icons" aria-label="Random recipe" href={"/recipe/" + randomId}>casino</IconButton>
      </Col>
   </Row>
  </form>
</div>
