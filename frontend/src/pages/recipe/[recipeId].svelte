<script>
  import { params } from '@sveltech/routify'
  import Card from '@smui/card';
  import { Container, Row, Col, Table } from 'sveltestrap';
  import Textfield, {Input, Textarea} from '@smui/textfield';
  import Icon from '@smui/textfield/icon/index';
  import Chip, {Set, Text} from '@smui/chips';
  import Button from '@smui/button';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';
  import Navbar from '../../components/Navbar.svelte';
  import Select, {Option} from '@smui/select';

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
      { amount: 1, unit: '', ingredient: 'Onion' }
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

  // All supported units
  const units = ['', 'g', 'kg', 'L', 'ml', 'tbsp', 'tsp', 'grain'];

  const all_tags = ['meat', 'snack', 'low cal', 'alcoholic'];
  
  for (let i = 0; i < 20; i++) {
    all_tags.push(`Item ${i}`);
  }

  function removeTag(tag) {
    recipe.tags = recipe.tags.filter(k => k !== tag);
  }
  
  function addTag(tag) {
    if (recipe.tags.includes(tag)) return;
    recipe.tags = [...recipe.tags, tag];
  }

  function clearTags() {
    recipe.tags = [];
  }

  function addCustomTag() {
    custom_tag_value = '';
    tag_dialog.open();
  }

  function formatTime(minutes) {
    if (minutes < 60) {
      return `${minutes} min`;
    } else {
      const h = Math.floor(minutes / 60);
      const m = minutes % 60;
      return `${h} h ${m} min`;
    }
  }
  
  function addIngredient() {
    recipe.ingredients = [{...new_ingredient}, ...recipe.ingredients];
  }

  function removeIngredient(name) {
    recipe.ingredients = recipe.ingredients.filter(i => i.ingredient !== name);
  }

  function removeStep(index) {
    recipe.steps.splice(index, 1);
    recipe.steps = [...recipe.steps];
  }

  function addStep() {
    recipe.steps = [...recipe.steps, { description: '' }];
  }

  function quick_add_ingredients() {
    const ingredients = quick_ingredients
      .split('\n')
      .map(line => /(\d+)([a-z]+)\s(.*)/.exec(line))
      .filter(match => match != null)
      .map(match => {
        return {
          ingredient: match[3],
          amount: match[1],
          unit: match[2]
        };
      });

    recipe.ingredients = [...ingredients, ...recipe.ingredients];
    
    quick_ingredients = '';
  }

  var edit = false;
  var custom_tag = '';
  var tag_dialog;
  var quick_ingredients_dialog;
  var custom_tag_value = '';
  var quick_ingredients = '';

  var new_ingredient = {
    ingredient: '',
    unit: 'g',
    amount: 1
  };
</script>

<div>
  <Navbar />
  <main id="recipe">
    <Card>
      <Container class="mt-4 p-3">
        <Row>
          <Col sm="12">
          {#if edit}
            <Textfield variant="outlined" bind:value={recipe.title} label="Title" />
          {:else}
            <h1>{recipe.title}</h1>
          {/if}
          </Col>
        </Row>
        {#if edit}
          <hr />
        {/if}
        <Row class="mt-2">
          <Col sm="12">
          {#if edit}
            <Textfield fullwidth textarea bind:value={recipe.description} label="Description" />
          {:else}
            <p>{recipe.description}</p>
          {/if}
          </Col>
        </Row>
        {#if edit}
          <hr />
          <Row class="mt-2">
            <Col sm="6">
            <Textfield variant="outlined" withTrailingIcon bind:value={recipe.servings} label="Servings" class="fullwidth" type="number">
              <Icon class="material-icons">people</Icon>
            </Textfield>
          </Col>
          <Col sm="6">
          <Textfield variant="outlined" withTrailingIcon bind:value={recipe.price} label="Price"  class="fullwidth">
            <Icon class="material-icons">euro_symbol</Icon>
          </Textfield>
          </Col>
          </Row>
          <hr />
          <Row class="mt-2">
            <Col sm="4">
            <Textfield variant="outlined" withTrailingIcon bind:value={recipe.time.cooktime} label="Cook time (minutes)" class="fullwidth">
              <Icon class="material-icons">schedule</Icon>
            </Textfield>
          </Col>
          <Col sm="4">
          <Textfield variant="outlined" withTrailingIcon bind:value={recipe.time.worktime} label="Work time (minutes)"  class="fullwidth">
            <Icon class="material-icons">schedule</Icon>
          </Textfield>
          </Col>
          <Col sm="4">
          <Textfield variant="outlined" withTrailingIcon bind:value={recipe.time.resttime} label="Rest time (minutes)"  class="fullwidth">
            <Icon class="material-icons">schedule</Icon>
          </Textfield>
          </Col>
          </Row>
        {:else}
          <p>
            <span class="m-2"><strong>{recipe.servings}</strong> Servings</span>&#x25CF;
            <span class="m-2"><strong>{recipe.price}</strong> &#x20AC</span>
          </p>
          <p>
            <span class="m-2"><em>Cook time:</em> <strong>{formatTime(recipe.time.cooktime)}</strong></span>&#x25CF;
            <span class="m-2"><em>Work time:</em> <strong>{formatTime(recipe.time.worktime)}</strong></span>&#x25CF;
            <span class="m-2"><em>Rest time:</em> <strong>{formatTime(recipe.time.resttime)}</strong></span>
          </p>
          <hr />
          <p>
            <Set chips={recipe.tags} let:chip input style="display: inline">
              <Chip><Text>{chip}</Text></Chip>
            </Set>
          </p>
        {/if}
      </Container>
    </Card>

    {#if !edit}
      <Row class="mt-4">
        <Col md={4}>
        <Card>
          <Container class="mt-2 p-3">
            <div class="ingredients">
              <h4>Ingredients</h4>
              <ul>
                {#each recipe.ingredients as ingredient}
                  <li>{ingredient.amount}{ingredient.unit} {ingredient.ingredient}</li>
                {/each}
              </ul>
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
                  <Row>
                    <Col md="1" class="step-number">
                    <h5>{index + 1}</h5>
                    </Col>
                    <Col md="11">
                    {step.description}
                    </Col>
                  </Row>
                </Container>
              {/each}
            </div>
          </Container>
        </Card>
        </Col>
      </Row>
    {/if}

    {#if edit}
      <Card class="mt-3">
        <Container class="mt-4 mb-4">
          <h4>Tags</h4>
          <a class="link" on:click={clearTags}>Clear all</a>
          <a class="link ml-2" on:click={addCustomTag}>Add</a>
          <hr />
          <Set chips={recipe.tags} let:chip input>
            <Chip>
              <Text>{chip}</Text>
              <Icon class="material-icons tag-icon" trailing tabindex="0" on:click={() => removeTag(chip)}>cancel</Icon>
            </Chip>
          </Set>
          <hr />
          Choose tags:
          <Set chips={all_tags} let:chip input>
            <Chip on:click={() => addTag(chip)}><Text>{chip}</Text></Chip>
          </Set>
        </Container>
      </Card>
      <Card class="mt-3">
        <Container class="mt-4 mb-4">
          <h4>Ingredients</h4>
          <a class="link" on:click={() => quick_ingredients_dialog.open()}>Quick add</a>
          <Table class="mt-3">
            <thead>
              <tr>
                <td>
                  <Textfield variant="outlined" bind:value={new_ingredient.ingredient} label="Ingredient" class="fullwidth"/>
                </td>
                <td>
                  <Textfield variant="outlined" bind:value={new_ingredient.amount} label="Amount" class="fullwidth"/>
                </td>
                <td>
                  <Select variant="outlined" bind:value={new_ingredient.unit} label="Unit" class="fullwidth">
                    {#each units as unit}
                      <Option value={unit} selected={new_ingredient.unit === unit}>{unit}</Option>
                    {/each}
                  </Select>
                </td>
                <td>
                  <Button on:click={addIngredient}>Add</Button>
                </td>
              </tr>
              {#each recipe.ingredients as ingredient}
                <tr>
                  <td>{ingredient.ingredient}</td>
                  <td>{ingredient.amount}</td>
                  <td>{ingredient.unit}</td>
                  <td><Button on:click={() => removeIngredient(ingredient.ingredient)}>Remove</Button>
                </tr>
              {/each}
            </thead>
          </Table>
        </Container>
      </Card>
      <Card class="mt-3">
        <Container class="mt-4 mb-4">
          <h4>Steps</h4>
          <a class="link" on:click={addStep}>Add step</a>
          <hr />
          {#each recipe.steps as step, index}
            <Card variant="outlined" class="p-2 mb-2">
              <span class="pt-2 pb-2">Step {index}:</span>
              <Textfield textarea variant="outlined" bind:value={step.description} class="fullwidth"/>
              <Button on:click={() => removeStep(index)}>Remove</Button>
            </Card>
          {/each}
        </Container>
      </Card>
    {/if}

    <Button on:click={() => edit = !edit}>Toggle edit mode</Button>
  </main>

  <!-- Create new tag dialog -->
  <Dialog bind:this={tag_dialog}>
    <Title id="event-title">Add Tag</Title>
    <Content id="event-content">
      Add custom tag:
      <br />
      <Textfield variant="outlined" bind:value={custom_tag_value} label="Tag" class="fullwidth mt-2"/>
    </Content>
    <Actions>
      <Button action="all" on:click={() => addTag(custom_tag_value)}>
        OK
      </Button>
    </Actions>
  </Dialog>

  <!-- Quick add ingredients dialog -->
  <Dialog bind:this={quick_ingredients_dialog}>
    <Title id="event-title">Quick add ingredients</Title>
    <Content id="event-content">
      Format:
      <pre><code>
          [amount][unit] [name]
      </code></pre>
      <br />
      <Textfield variant="outlined" textarea bind:value={quick_ingredients} class="fullwidth mt-2"/>
    </Content>
    <Actions>
      <Button action="all" on:click={quick_add_ingredients}>
        OK
      </Button>
    </Actions>
  </Dialog>
</div>
