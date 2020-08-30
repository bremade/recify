<script>
  import { params } from '@sveltech/routify'
  import Card from '@smui/card';
  import Button from '@smui/button';
  import Textfield from '@smui/textfield';
  import { Container, Row, Col, Table } from 'sveltestrap';
  import Navbar from '../../components/Navbar.svelte';
  import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
  import TimeInput from '../../components/Time.svelte';

  let recipe = {
    id: '',
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

  const steps = [
    { name: 'General', completed: false },
    { name: 'Time', completed: false },
    { name: 'Ingredients', completed: false },
    { name: 'Steps', completed: false },
    { name: 'Tags', completed: false },
    { name: 'Check and publish', completed: false }
  ];

  let current_step = 0;
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
            {/if}
            <!-- Time -->
            {#if current_step === 1}
              <TimeInput bind:time={recipe.time.worktime} label="Work time" />
              <TimeInput bind:time={recipe.time.cooktime} label="Cook time" />
              <TimeInput bind:time={recipe.time.resttime} label="Rest time" />
            {/if}
            <!-- Next-button -->
            <hr />
            {#if current_step < steps.length -1}
              <Button variant="raised" class="float-right" on:click={() => current_step += 1}>Next: {steps[current_step + 1].name}</Button>
            {/if}
          </Container>
        </Card>
        </Col>
      </Row>
    </Container>
  </main>
</div>
