<script>
    import { url } from "@sveltech/routify";
    import Button, {Label} from '@smui/button';
</script>

<div id="notfound">
    <div class="notfound">
        <div class="error404">
            <h3 class="notfoundh3">
            Oops! Page not found
            </h3>
            <h1 class="notfoundh1">
                <span>404</span>
            </h1>
        </div>
        <h2 class="notfoundh2">
            im sorry, but the page you requested was not found
        </h2>
        <Button class="notfoundbutton" href={$url("/")} variant="raised">
            <Label>Home Page</Label>
        </Button>
    </div>
</div>