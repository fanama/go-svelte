<script>
    export let name = "fanama"
</script>

<style>
    h1 {
        padding: 1em;
        float: right;
    }
</style>

<h1>
    {#if name=="fanama"}
        you are the owner
    {/if}
</h1>