<script>
  import Filler from '../components/Filler.svelte';
  import Navbar from '../components/Navbar.svelte';
  import RecipeCard from '../components/landingpage/RecipeCard.svelte';
  import Slider from '../components/landingpage/Slider.svelte';
  import { Container, Row, Col, Button } from 'sveltestrap';

  let topRecipes = ["Das ist ein unglaublich langer text um den umbruch zu testen", "Test2", "Der Text hier ist noch länger und geht über ganze 3 Zeilen statt nur 2 wuhuuu wuhuuuuu", "Test4", "Test5", "Test6", "Test7", "Test8"]
  let topCategories = [
    "Tag 1",
    "Tag 2",
    "Vegan",
    "Meat",
    "Vegan",
    "Meat",
    "Vegan",
    "Meat",
    "Vegan",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
    "Meat",
  ]
  let top2Categories = topCategories.slice(0, 2);
  let tmp = [
    "Test",
    "Test"
  ]
</script>

<Navbar/>
<main id="content">
  <Container class="contentContainer">
    <!-- Top 8 recipes of all categories -->
    <Row>
      <h2>Top Recipes</h2>
    </Row>
    <Row id="recipeCards" class="rowLayout">
      {#each topRecipes as topRecipe}
        <Col sm="12" md="4" xl="3" class="gridLayout">
          <RecipeCard recipeId={topRecipe}/>
        </Col>
      {/each}
      <Col sm="12" class="gridLayout" align="center">
        <Button outline class="moreButton">more</Button>
      </Col>
    </Row>

    <!-- List of Top 20 categories -->
    <Row>
      <h2>Top Categories</h2>
    </Row>
    <Row class="rowLayout">
      <Slider tags={topCategories}/>
    </Row>

    <!-- Top categories displayed with its top recipes -->
    {#each top2Categories as category}
      <Row>
        <h2>Good {category} recipes</h2>
      </Row>
      <Row id="recipeCards" class="rowLayout">
        {#each tmp as tmp}
          <Col sm="12" md="4" xl="3" class="gridLayout">
            <RecipeCard recipeId={tmp}/>
          </Col>
        {/each}
        <Col sm="12" class="gridLayout" align="center">
          <Button outline class="moreButton">more</Button>
        </Col>
      </Row>
    {/each}
  </Container>
</main>
