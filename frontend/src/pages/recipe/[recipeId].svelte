<script>
  import { params } from '@sveltech/routify'
  import Card from '@smui/card';
  import { Container, Row, Col, Table } from 'sveltestrap';
  import Navbar from '../../components/Navbar.svelte';

  // Extract recipe-id from route
  const id = params.recipeId;

  const recipe = {
    _id: id,
    title: 'Tomato soup',
    servings: 4,
    time: {
      cooktime: 160,
      worktime: 100,
      resttime: 100
    },
    ingredients: [
      { amount: 1, unit: 'L', ingredient: 'Milk' },
      { amount: 5, unit: '', ingredient: 'Tomatoes' },
      { amount: 1, unit: '', ingredient: 'Onion' },
      { amount: 2, unit: 'g', ingredient: 'Salt' },
      { amount: 500, unit: 'ml', ingredient: 'Water' }
    ],
    steps: [
      { description: 'Cut tomatoes' },
      { description: 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.' },
      { description: 'Cook tomatoes for at least 10 minutes' },
      { description: 'Serve with an onion' }
    ],
    price: 1.5,
    description: "A delious soup",
    tags: ['Delicious', 'Soup', 'veggy']
  };

  var servings = recipe.servings;

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
      <img class="recipe-image" src="/placeholder.webp">
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
                  <td class="amount">{(ingredient.amount / recipe.servings) * servings}{ingredient.unit}</td>
                  <td class="ingredient">{ingredient.ingredient}</td>
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
