let greeting = document.querySelector("#greeting");
let colorChangeButton = document.querySelector("#color-change");
let body = document.querySelector("body");

colorChangeButton.addEventListener(
    "click",
    changeColorsRand.bind(this)
);

function changeColorsRand() {
    colors = ["red", "orange", "yellow", "green", "blue", "purple"];
    randColor = colors[Math.floor(Math.random() * colors.length)];
    randBGColor = colors[Math.floor(Math.random() * colors.length)];

    while (randColor === randBGColor){
        randBGColor = colors[Math.floor(Math.random() * colors.length)];
    }

    greeting.style.color = randColor;
    body.style.backgroundColor = randBGColor;
    colorChangeButton.style.color = randColor;
    colorChangeButton.style.backgroundColor = randBGColor;
    
}