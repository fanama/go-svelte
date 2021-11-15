<script>

    import Headder from './utils/headder.svelte'
    import Liste from './dashboard/liste/liste.svelte'

    let name="fanama";

    const changeName = (e)=>{
      name=e.detail
    }
    

  </script>
  

  <main>
    <Headder name="my svelte" />

    <h1>Hello {name}!</h1>
    
    
    <div class="list">

      { #if name=="fanama"}
        you are the owner!!!
      {/if}

      <Liste name={name} on:alert={changeName} />
      

    </div>
    

  </main>
