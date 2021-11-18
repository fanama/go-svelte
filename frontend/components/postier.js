const {API} = process.env

class Postier {
    Get(route){
        fetch(`${API}/${route}`).then(res=>{
            console.log(res.body)
        }).catch(err=>{
            console.log(err)
        })
    }

    Post(route,form){
        fetch(`${API}/${route}`, {
            method: "POST",
            body: JSON.stringify(form)
          }).then(res=>{
            console.log(res.body)
        }).catch(err=>{
            console.log(err)
            alert(`erreur de post at ${API}`)
        })
    }
}

export const postier = new Postier()