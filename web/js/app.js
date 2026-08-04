console.log("TV Remote Started");

async function send(key){

    const r = await fetch("/api/key",{

        method:"POST",

        headers:{
            "Content-Type":"application/json"
        },

        body:JSON.stringify({

            key:key

        })

    });

    const json=await r.json();

    console.log(json);

}