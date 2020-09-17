<script>
  import Filler from '../components/Filler.svelte';
  import Navbar from '../components/Navbar.svelte';
  import RecipeCard from '../components/RecipeCard.svelte';
  import { Container, Row, Col, Table } from 'sveltestrap';

  let topRecipes = ["Das ist ein unglaublich langer text um den umbruch zu testen", "Test2", "Der Text hier ist noch länger und geht über ganze 3 Zeilen statt nur 2 wuhuuu wuhuuuuu", "Test4", "Test5", "Test6", "Test7", "Test8"]
</script>

<Navbar/>
<main id="content">
  <Container class="contentContainer">
    <Row>
      <h2>Top Recipes</h2>
    </Row>
    <Row id="recipeCards">
      {#each topRecipes as topRecipe}
        <Col sm="12" md="4" xl="3" class="gridLayout">
          <RecipeCard recipeId={topRecipe}/>
        </Col>
      {/each}
    </Row>
    <Row>
      <Filler/>
    </Row>
  </Container>
</main>
