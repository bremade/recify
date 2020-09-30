<script>
  import { context, params } from '@sveltech/routify'
  import Card from '@smui/card';
  import { Container, Row, Col, Table } from 'sveltestrap';
  import Navbar from '../../components/Navbar.svelte';
  import formatter from '../../utils/formatter.js';

  // Extract recipe-id from route
  const id = $params.recipeId;

  let recipe = {
    id: id,
    title: '',
    servings: 1,
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

  // Constraint values for servings

  function servingsConstraint(_) {
    if (servings < 1) servings = 1;
    if (servings > 99) servings = 99;
  }

  $: servingsConstraint(servings);

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
            <h5 class="ml-3">A recipe for <input type="number" class="servings" bind:value={servings} /> person{servings > 1 ? 's' : ''}</h5>
            <Table class="mt-3">
              <tr>
                <td>Price <small>(per serving)</small></td>
                <td>{formatter.formatPrice(recipe.price)} &euro;</td>
              </tr>
              <tr>
                <td>Work time</td>
                <td>{formatter.formatTime(recipe.time.worktime)}</td>
              </tr>
              <tr>
                <td>Cook time</td>
                <td>{formatter.formatTime(recipe.time.cooktime)}</td>
              </tr>
              <tr>
                <td>Rest time</td>
                <td>{formatter.formatTime(recipe.time.resttime)}</td>
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
