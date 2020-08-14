<script>
    import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
    import IconButton from '@smui/icon-button';

    let prominent = false;
    let dense = false;
    let secondaryColor = false;
    export let drawerOpen = false;
</script>

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