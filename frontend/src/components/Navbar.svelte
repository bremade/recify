<script>
  import TopAppBar from '@smui/top-app-bar';
  import { Collapse, Container, Col, Row } from 'sveltestrap';
  import { Button, Card } from 'sveltestrap';
  import IconButton from '@smui/icon-button';
  import Sidebar from './Sidebar.svelte';
  import Searchbox from './Searchbox.svelte';

  let isOpen = false;
  let prominent = false;
  let dense = false;
  let secondaryColor = false;
  let drawerOpen = false;
  let mobile = false;

  function checkMobile(query) {
    if (query.matches) {
      console.log("mobile");
      mobile = true;
    } else {
      console.log("normal");
      mobile = false;
      isOpen = false;
    }
  }

  let query = window.matchMedia("(max-width: 850px)");
  checkMobile(query);
  query.addListener(checkMobile);

</script>

<Sidebar bind:drawerOpen/>
<TopAppBar variant="static" {prominent} {dense} color={secondaryColor ? 'secondary' : 'primary'}>
  <Container class="header-container">
    <Row>
      {#if !mobile}
        <Col xs="2" class="containerCenter">
          Recify
        </Col>
        <Col class="containerCenter">
          <Searchbox/>
        </Col>
        <Col xs="2" class="text-right">
          <IconButton class="material-icons input-button" on:click={() => drawerOpen = !drawerOpen}>menu</IconButton>
        </Col>
      {:else}
        <Col xs="auto" class="containerCenter">
          Recify
        </Col>
        <Col/>
        <Col xs="auto">
          <IconButton class="material-icons" aria-label="Open recipe search" on:click={() => isOpen = !isOpen}>search</IconButton>
          <IconButton class="material-icons" on:click={() => drawerOpen = !drawerOpen}>menu</IconButton>
        </Col>
      {/if}
    </Row>
  </Container>
</TopAppBar>
<Collapse {isOpen}>
  <div class="mobileSearch">
    <Searchbox bind:isOpen/>
  </div>
</Collapse>

