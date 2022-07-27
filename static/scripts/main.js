const inputField = document.getElementById('filterInput')

inputField.addEventListener('keyup', (e) => {
    const searchString = e.target.value;

    const trimming = searchString.trim();
    const searchTerms = trimming.split(' ');
    console.log(searchTerms)

    // const style = searchTerms.map((term) => {
    //     return `<span>${term}</span>`
    // })
    //inputField.append(style);

    // const style = searchTerms.map((term) => {
    //     const styled = term.classListadd("fw-semibold")
    //     return inputField.append(styled);
    // })
    for (value of inputField.value){
        value.classList.add("fw-semibold")
    }
 

})

function clearSearchBar(){
    inputField.value = "";
}