console.log("working");
let items = document.getElementsByTagName('li');

for (let i = 0; i < items.length; i++) {
    items[i].addEventListener('click', () => {
        this.classList.toggle('done');
    });
}