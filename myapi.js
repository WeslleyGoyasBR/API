function carregaDados(){
    fetch("http://localhost:8080/articles")
    .then(response => {
      response.json()
      .then(data => {
        console.log(data)
        let tb=document.getElementById("corpo-tabela");
        for( let i=0; i<data.length; i++) {
          let tr=document.createElement("TR");

          let tdTitle=document.createElement("TD");
          let title=data[i].Title;
          let tnTitle=document.createTextNode(title);
          tdTitle.appendChild(tnTitle);
          tr.appendChild(tdTitle);

          let tdDesc=document.createElement("TD");
          let desc=data[i].Desc;
          let tnDesc=document.createTextNode(desc);
          tdDesc.appendChild(tnDesc);
         tr.appendChild(tdDesc);

          let tdContent=document.createElement("TD");
          let content = data[i].Content;
          let tnContent=document.createTextNode(content);
          tdContent.appendChild(tnContent);
          tr.appendChild(tdContent);
          
          tb.appendChild(tr);
        }
      });
    })
}

