<script>
  import { params } from '@sveltech/routify'
  import Card from '@smui/card';
  import { Container, Row, Col, Table } from 'sveltestrap';
  import Navbar from '../../components/Navbar.svelte';

  // Extract recipe-id from route
  const id = $params.recipeId;

  let recipe = {
    id: id,
    title: '',
    servings: 0,
    time: {
      cooktime: 0,
      worktime: 0,
      resttime: 0
    },
    ingredients: [],
    steps: [],
    price: 0,
    description: '',
    tags: []
  };

  let servings = recipe.servings;

  fetch('/api/v1/recipe/' + id)
  .then(resp => resp.json())
  .then(resp => {
    recipe = { ...recipe, ...resp };
    servings = recipe.servings;
  });

  function formatTime(minutes) {
    if (minutes < 60) {
      return `${minutes} min`;
    } else {
      const h = Math.floor(minutes / 60);
      const m = minutes % 60;

      if (m > 0) {
        return `${h} h ${m} min`;
      } else {
        return `${h} h`;
      }
    }
  }

  function formatPrice(price) {
    const eur = Math.floor(price);
    const cent = (price - eur) * 100;

    return `${eur}.${String(cent).padStart(2, '0')}`;
  }

</script>

<div>
  <Navbar />
  <main id="recipe">
    <Card>
      <Container class="mt-4 p-3">
        <Row>
          <Col sm="12">
          <h1>{recipe.title}</h1>
          </Col>
        </Row>
        <Row>
          <Col sm="12">
          <p>{recipe.description}</p>
          </Col>
        </Row>
      </Container>
      {#if recipe.picture }
        <img class="recipe-image" src={`/pictures/${recipe.picture}`}>
      {/if}
    </Card>
    <Row class="mt-4">
      <Col md={4}>
      <Card>
        <Container class="mt-2 p-3">
          <div class="ingredients">
            <h4>Details</h4>
            <h5 class="ml-3">A recipe for <input type="number" class="servings" bind:value={servings} /> persons</h5>
            <Table class="mt-3">
              <tr>
                <td>Price <small>(per serving)</small></td>
                <td>{formatPrice(recipe.price)} &euro;</td>
              </tr>
              <tr>
                <td>Work time</td>
                <td>{formatTime(recipe.time.worktime)}</td>
              </tr>
              <tr>
                <td>Cook time</td>
                <td>{formatTime(recipe.time.cooktime)}</td>
              </tr>
              <tr>
                <td>Rest time</td>
                <td>{formatTime(recipe.time.resttime)}</td>
              </tr>
            </Table>
          </div>
        </Container>
      </Card>
      <Card class="mt-4">
        <Container class="mt-2 p-3">
          <div class="ingredients">
            <h4>Ingredients</h4>
            <table class="ingredients">
              {#each recipe.ingredients as ingredient, index}
                <tr class={index % 2 == 0 ? 'highlighted' : ''}>
                  <td class="amount">{(ingredient.amount / recipe.servings) * servings} {ingredient.unit}</td>
                  <td class="ingredient">{ingredient.name}</td>
                  </tr>
              {/each}
            </table>
          </div>
        </Container>
      </Card>
        </Col>
        <Col md={8}>
        <Card>
          <Container class="mt-2 p-3">
            <div class="method">
              <h4>Method</h4>
              {#each recipe.steps as step, index}
                <Container class="m-2 p-2" variant="outlined">
                  <Row class="steps">
                    <Col md="1">
                    <span class="number">{index + 1}</span>
                    </Col>
                    <Col md="11">
                    <div class="step">
                      {step.description}
                    </div>
                    </Col>
                  </Row>
                </Container>
              {/each}
            </div>
          </Container>
        </Card>
        </Col>
    </Row>
  </main>
</div>
