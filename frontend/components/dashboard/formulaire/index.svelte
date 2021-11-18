<script>
    import {postier} from "../../postier"

    let database,name,power

    const sendData = ()=>{
        const data = {name,power}
        postier.Post(`database/${database}`,data)
    }

</script>

<style lang="css">
    @import "./style.scss";

</style>

<main class="formulaire">
    <form on:submit|preventDefault={sendData}>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>database</label><input  bind:value={database} type="text"/>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>name</label><input  bind:value={name} type="text"/>
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label>power</label><input  bind:value={power} type="text"/>
        <button>envoyer</button>

    
    </form>
</main>