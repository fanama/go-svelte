<script>

import {createEventDispatcher} from 'svelte'

export let name

const dispatch = createEventDispatcher()

let list = [
      {name:"hiro"},
      {name:"clark"}
    ]

    const handleSubmit = (e)=>{
      list.push({name})

      dispatch("alert",name)
      list = list
    }

</script>

<style lang="css">
    @import "./style.scss";

</style>

<div class="liste">
    <form on:submit|preventDefault={handleSubmit} >
        <input bind:value={name} type="text" />
        <button>add</button>
      </form>
    
      {#each list as element,i (element.name)}
            <li>{i}:{element.name}</li>
      {/each}
</div>

