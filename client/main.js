(function(){

    // Call api 

    setTimeout(()=>{
        fetch('http://localhost:3000')
        .then(response => response.json())
        .then(data =>{
          document.getElementById("loader").style.display="none"
         const div=document.getElementById("append");
      
         data.data.forEach(element => {
          div.innerHTML+=`<img src="${element.image}" alt="Crime 1" class="rounded-md">`
          
         });
      
      
      
        });

    },1000)  // intentional delay in loading

   



})()