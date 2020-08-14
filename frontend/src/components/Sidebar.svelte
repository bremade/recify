<script>
    import Drawer, {AppContent, Content, Header, Title, Subtitle, Scrim} from '@smui/drawer';
    import Button, {Label} from '@smui/button';
    import List, {Item, Text, Graphic, Separator, Subheader} from '@smui/list';
    import H6 from '@smui/common/H6.svelte';
    let drawer;
    let drawerOpen;
    let active = 'Inbox';
    function setActive(value) {
      active = value;
      drawerOpen = false;
    }
  </script>

<div class="drawer-container">
    <Drawer variant="modal" bind:this={drawer} bind:open={drawerOpen}>
      <Header>
        <Title>Super Mail</Title>
        <Subtitle>It's the best fake mail app drawer.</Subtitle>
      </Header>
      <Content>
        <List>
          <Item href="javascript:void(0)" on:click={() => setActive('Inbox')} activated={active === 'Inbox'}>
            <Graphic class="material-icons" aria-hidden="true">inbox</Graphic>
            <Text>Inbox</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => setActive('Star')} activated={active === 'Star'}>
            <Graphic class="material-icons" aria-hidden="true">star</Graphic>
            <Text>Star</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => setActive('Sent Mail')} activated={active === 'Sent Mail'}>
            <Graphic class="material-icons" aria-hidden="true">send</Graphic>
            <Text>Sent Mail</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => setActive('Drafts')} activated={active === 'Drafts'}>
            <Graphic class="material-icons" aria-hidden="true">drafts</Graphic>
            <Text>Drafts</Text>
          </Item>

          <Separator nav />
          <Subheader component={H6}>Labels</Subheader>
          <Item href="javascript:void(0)" on:click={() => setActive('Family')} activated={active === 'Family'}>
            <Graphic class="material-icons" aria-hidden="true">bookmark</Graphic>
            <Text>Family</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => setActive('Friends')} activated={active === 'Friends'}>
            <Graphic class="material-icons" aria-hidden="true">bookmark</Graphic>
            <Text>Friends</Text>
          </Item>
          <Item href="javascript:void(0)" on:click={() => setActive('Work')} activated={active === 'Work'}>
            <Graphic class="material-icons" aria-hidden="true">bookmark</Graphic>
            <Text>Work</Text>
          </Item>
        </List>
      </Content>
    </Drawer>

    <AppContent class="app-content">
      <main class="main-content">
        <Button on:click={() => drawerOpen = !drawerOpen}><Label>Toggle Drawer</Label></Button>
        <br />
        <pre class="status">Active: {active}</pre>
        <div style="height: 700px;">&nbsp;</div>
        And some stuff at the bottom.
      </main>
    </AppContent>
  </div>

  <style>
    .drawer-container {
      position: relative;
      display: flex;
      height: 350px;
      max-width: 600px;
      border: 1px solid rgba(0,0,0,.1);
      overflow: hidden;
      z-index: 0;
    }
    * :global(.mdc-drawer--modal, .mdc-drawer-scrim) {
      /* This is not needed for a page-wide modal. */
      position: absolute;
    }
    * :global(.app-content) {
      flex: auto;
      overflow: auto;
      position: relative;
      flex-grow: 1;
    }
    .main-content {
      overflow: auto;
      padding: 16px;
      height: 100%;
      box-sizing: border-box;
    }
  </style>