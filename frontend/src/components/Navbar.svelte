<script>
    import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
    import IconButton from '@smui/icon-button';
    import Sidebar from './Sidebar.svelte';
    let prominent = false;
    let dense = false;
    let secondaryColor = false;
    let drawerOpen = false;
</script>
<div>
    <Sidebar bind:drawerOpen/>
    <TopAppBar variant="static" {prominent} {dense} color={secondaryColor ? 'secondary' : 'primary'}>
        <Row>
            <Section>
            <IconButton class="material-icons" on:click={() => drawerOpen = !drawerOpen}>menu</IconButton>
            <Title>Recify</Title>
            </Section>
            <Section align="end" toolbar>
                <IconButton class="material-icons" aria-label="Share">share</IconButton>
                <IconButton class="material-icons" aria-label="Like">favorite</IconButton>
                <IconButton class="material-icons" aria-label="Bookmark this page">bookmark</IconButton>
            </Section>
        </Row>
    </TopAppBar>
</div>