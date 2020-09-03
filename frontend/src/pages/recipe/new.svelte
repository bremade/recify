<script>
  import { params } from '@sveltech/routify'
  import Card from '@smui/card';
  import Button from '@smui/button';
  import Textfield from '@smui/textfield';
  import { Container, Row, Col, Table } from 'sveltestrap';
  import Navbar from '../../components/Navbar.svelte';
  import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
  import TimeInput from '../../components/Time.svelte';
  import Select, {Option} from '@smui/select';

  const units = [' ', 'g', 'mg', 'kg', 'l', 'ml', 'cl', 'tsp', 'tbsp', 'grain'];

  const empty_ingredient = { name: '', amount: 0, unit: units[0] };
  const empty_step =  { description: '' };

  let recipe = {
    id: '',
    title: '',
    servings: 1,
    time: {
      cooktime: 0,
      worktime: 0,
      resttime: 0
    },
    ingredients: [{...empty_ingredient}],
    steps: [
      {...empty_step}
    ],
    price: 0,
    description: '',
    tags: []
  };

  const steps = [
    { name: 'General', completed: false },
    { name: 'Time', completed: false },
    { name: 'Ingredients', completed: false },
    { name: 'Steps', completed: false },
    { name: 'Tags', completed: false },
    { name: 'Check and publish', completed: false }
  ];

  let current_step = 0;

  function addIngredient() {
    recipe.ingredients = [...recipe.ingredients, {...empty_ingredient}];
  }

  function removeIngredient(index) {
    recipe.ingredients.splice(index, 1);
    recipe.ingredients = recipe.ingredients;
  }

  function addStep() {
    recipe.steps = [...recipe.steps, {...empty_step}];
  }

  function removeStep(index) {
    recipe.steps.splice(index, 1);
    recipe.steps = recipe.steps;
  }

  function publish() {
    // Convert all amounts to numbers
    recipe.ingredients = recipe.ingredients.map(i => { return {...i, amount: Number(i.amount)}; });
    // Send request
    fetch('/api/v1/recipe', {
      method: 'POST',
      body: JSON.stringify(recipe)
    }).then(result => {
      if (result.status === 201) {
        window.location.pathname = '/';
      }
    });
  }

</script>

<div>
  <Navbar />
  <main id="recipe">
    <Container>
      <Row class="mt-4">
        <Col md="4">
        <Card>
          <Container class="mt-2 p-3">
            <h4>Create Recipe</h4>
            <div>
              <List>
                {#each steps as step, index}
                  <Item on:click={() => current_step = index} activated={current_step === index} class="step-item">
                    <Text>{index + 1} - {step.name}</Text>
                  </Item>
                {/each}
              </List>
            </div>
          </Container>
        </Card>
        </Col>
        <Col md="8">
        <Card>
          <Container class="mt-2 p-3">
            <h4>{steps[current_step].name}</h4>
            <!-- General -->
            {#if current_step === 0}
              <Textfield variant="outlined" bind:value={recipe.title} label="Title" class="fullwidth" />
              <Textfield textarea variant="outlined" bind:value={recipe.description} label="Short description" class="fullwidth mt-2" />
              <hr />
              <Textfield variant="outlined" bind:value={recipe.servings} label="Servings" class="fullwidth" type="number" />
              <!-- Time -->
            {:else if current_step === 1}
              <TimeInput bind:time={recipe.time.worktime} label="Work time" />
              <TimeInput bind:time={recipe.time.cooktime} label="Cook time" />
              <TimeInput bind:time={recipe.time.resttime} label="Rest time" />
              <!-- Ingredients -->
            {:else if current_step === 2}
              <Table>
                <thead>
                  <tr>
                    <th>Ingredient</th><th>Amount</th><th>Unit</th><th></th>
                  </tr>
                </thead>
                {#each recipe.ingredients as ingredient, index}
                  <tr>
                    <td><Textfield bind:value={ingredient.name} class="fullwidth"/></td>
                    <td><Textfield bind:value={ingredient.amount} class="fullwidth"/></td>
                    <td>
                      <Select bind:value={ingredient.unit} label="Unit">
                        {#each units as unit}
                          <Option value={unit} selected={ingredient.unit === unit}>{unit}</Option>
                        {/each}
                      </Select>
                    </td>
                    <td>
                      {#if index > 0}
                        <Button on:click={() => removeIngredient(index)}>Remove</Button>
                      {/if}
                    </td>
                  </tr>
                {/each}
              </Table>
              <Button on:click={addIngredient}>Add ingredient</Button>
              <!-- Steps -->
            {:else if current_step === 3}
              {#each recipe.steps as step, index}
                <Container class="m-2 p-2" variant="outlined">
                  <Row class="steps">
                    <Col md="1">
                    <span class="number">{index + 1}</span>
                    </Col>
                    <Col md="11">
                    <Textfield bind:value={step.description} label={`Step ${index + 1}`} textarea class="w-100" />
                     {#if index > 0}
                        <Button on:click={() => removeStep(index)}>Remove</Button>
                      {/if}
                    </Col>
                  </Row>
                </Container>
              {/each}
              <Button on:click={addStep}>Add Step</Button>
            {:else if current_step === 4}
              <div>Tags</div>
            {:else if current_step === 5}
              <div>
                <Button on:click={publish}>Publish recipe</Button>
              </div>
            {/if}
            <!-- Next-button -->
            {#if current_step < steps.length -1}
              <hr />
              <Button variant="raised" class="float-right" on:click={() => current_step += 1}>Next: {steps[current_step + 1].name}</Button>
            {/if}
          </Container>
        </Card>
        </Col>
      </Row>
    </Container>
  </main>
</div>
