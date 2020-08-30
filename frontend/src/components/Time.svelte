<script>
  import Card from '@smui/card';
  import Textfield from '@smui/textfield';
  import { Container, Row, Col, Table } from 'sveltestrap';
  
  export let time = 0;
  export let label = '';

  let minutes = time % 60;
  let hours = Math.floor(time / 60);

  function calculateTime(_) {
    time = minutes + hours * 60;
  }

  $: calculateTime(hours);
  $: calculateTime(minutes);
  
</script>

<div>
  <Card variant="outlined" class="p-2 mt-2 mb-2">
    <h6>{label}</h6>
    <Container>
      <Row>
        <Col md="6">
        <Textfield label="Hours" bind:value={hours} class="fullwidth" type="number" />
        </Col>
        <Col md="6">
        <Textfield label="Minutes" bind:value={minutes} class="fullwidth" type="number" />
        </Col>
      </Row>
    </Container>
  </Card>
</div>
