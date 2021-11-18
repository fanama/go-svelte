<script>

    import Headder from './utils/headder.svelte'
    import Liste from './dashboard/liste/liste.svelte'
    import Condition from './dashboard/condition/index.svelte'
    import Form from "./dashboard/formulaire/index.svelte"
    let name="fanama";

    const changeName = (e)=>{
      name=e.detail
    }
    

  </script>

  <style>
    h1 {
      text-align: center;
    }
  </style>
  

  <main>
    <Headder name="my svelte" />

    <h1>Hello {name}! {process.env.API} </h1>
    <Condition name={name} />
    
    
    <div class="list">

      <Liste name={name} on:alert={changeName} />
      <Form/>

    </div>
    

  </main>
