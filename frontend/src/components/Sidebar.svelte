<script>
  import Drawer, {Content, Header, Title, Subtitle, Scrim} from '@smui/drawer';
  import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
  import H6 from '@smui/common/H6.svelte';
  import session from '../stores/session.js';

  session.updateSessionStatus();

  let drawer;
  export let drawerOpen = false;
  let active = 'Auth';
  function setActive(value) {
    active = value;
    drawerOpen = false;
  }

  const random_string = Math.random().toString(36).substring(7);
  const keycloak_login = `/auth/realms/recify/protocol/openid-connect/auth?client_id=recify_app&response_type=code&login=true&scope=openid+profile+email&redirect_uri=${window.location.protocol}//${window.location.host}/api/v1/auth/token`;
</script>

<Drawer variant="modal" anchor="right" bind:this={drawer} bind:open={drawerOpen}>
  <Header>
    <Title>Recify</Title>
    {#if $session.logged_in}
      <Subtitle>Logged in as <strong>{$session.username}</strong></Subtitle>
    {:else}
      <Subtitle>Welcome to our website.</Subtitle>
    {/if}
  </Header>
    <Content>
      <List>
        <Separator nav />
        <Subheader component={H6}>User</Subheader>

        {#if $session.logged_in}
          <Item href="javascript:void(0)" on:click={() => {window.location.href = '/recipe/new'; setActive('Create Recipe')}} activated={active === 'Create Recipe'}>
            <Graphic class="material-icons" aria-hidden="true">fastfood</Graphic>
            <Text>Create Recipe</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => setActive('Import Recipe')} activated={active === 'Import Recipe'}>
            <Graphic class="material-icons" aria-hidden="true">restaurant_menu</Graphic>
            <Text>Import Recipe</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => {window.location.href = '/logout'; setActive('Auth')}} activated={active === 'Auth'}>
            <Graphic class="material-icons" aria-hidden="true">exit_to_app</Graphic>
            <Text>Logout</Text>
          </Item>
        {:else}
          <Item on:click={() => window.location.href = keycloak_login} activated={active === 'Auth'}>
            <Graphic class="material-icons" aria-hidden="true">account_circle</Graphic>
            <Text>Login</Text>
          </Item>
        {/if}

        <Separator nav />
        <Subheader component={H6}>Recipes</Subheader>
        <Item href="javascript:void(0)" on:click={() => setActive('Bread')} activated={active === 'Bread'}>
            <Graphic class="material-icons" aria-hidden="true">restaurant</Graphic>
            <Text>Bread</Text>
        </Item>
        <Item href="javascript:void(0)" on:click={() => setActive('Fish')} activated={active === 'Fish'}>
            <Graphic class="material-icons" aria-hidden="true">restaurant</Graphic>
            <Text>Fish</Text>
        </Item>
        <Item href="javascript:void(0)" on:click={() => setActive('Noodles')} activated={active === 'Noodles'}>
            <Graphic class="material-icons" aria-hidden="true">restaurant</Graphic>
            <Text>Noodles</Text>
        </Item>

        <Separator nav />
        <Subheader component={H6}>Cookbooks</Subheader>
        <Item href="javascript:void(0)" on:click={() => setActive('Top 10')} activated={active === 'Top 10'}>
            <Graphic class="material-icons" aria-hidden="true">book</Graphic>
            <Text>Top 10</Text>
        </Item>
        <Item href="javascript:void(0)" on:click={() => setActive('Asia')} activated={active === 'Asia'}>
            <Graphic class="material-icons" aria-hidden="true">book</Graphic>
            <Text>Asia</Text>
        </Item>
        <Item href="javascript:void(0)" on:click={() => setActive('Vegetarian')} activated={active === 'Vegetarian'}>
            <Graphic class="material-icons" aria-hidden="true">book</Graphic>
            <Text>Vegetarian</Text>
        </Item>
    </List>
    </Content>
</Drawer>
<Scrim/>
