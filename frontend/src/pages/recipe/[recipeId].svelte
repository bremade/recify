<script>
  import { params } from '@sveltech/routify'
  import Card from '@smui/card';
  import { Container, Row, Col } from 'sveltestrap';
  import Textfield, {Input, Textarea} from '@smui/textfield';
  import Icon from '@smui/textfield/icon/index';
  import Chip, {Set, Text} from '@smui/chips';
  import Button from '@smui/button';
  import Dialog, {Title, Content, Actions, InitialFocus} from '@smui/dialog';

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
    price: 1.5,
    description: "A delious soup",
    tags: ['Delicious', 'Soup', 'veggy']
  };

  const all_tags = ['meat', 'snack', 'low cal', 'alcoholic']
  
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

  var custom_tag = '';
  var tag_dialog;
  var custom_tag_value = '';
</script>

<div>
  <main id="recipe">
    <Card variant="outlined">
      <h1 class="mt-4">{recipe.title}</h1>
      <Container class="mt-4">
        <Row>
          <Col sm="12">
          <Textfield fullwidth textarea bind:value={recipe.description} label="Description" />
          </Col>
        </Row>
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
      </Container>
    </Card>
    <Card variant="outlined" class="mt-3">
      <Container class="mt-4 mb-4">
        <h4>Tags</h4>
        <a class="link" on:click={clearTags}>Clear all</a>
        <a class="link ml-2" on:click={addCustomTag}>Add</a>
        <hr />
        <Set chips={recipe.tags} let:chip input>
          <Chip><Text>{chip}</Text><Icon class="material-icons tag-icon" trailing tabindex="0" on:click={() => removeTag(chip)}>cancel</Icon></Chip>
        </Set>
        <hr />
        Choose tags:
        <Set chips={all_tags} let:chip input>
          <Chip on:click={() => addTag(chip)}><Text>{chip}</Text></Chip>
        </Set>
      </Container>
    </Card>
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
</div>
