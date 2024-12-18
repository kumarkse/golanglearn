// alert('hello');
console.log("hello");

let score =50;
score =90
console.log(score)

const score2 = 50;
// score2=10;   error

let name = "hello"
let age= 30
rating = 4.5
x = true
y = null
z = undefined
let beta

console.log(typeof(beta))
console.log('myname'+age+'hello d')
console.log(`my age is ${age}`)

let str = 'hello world'

console.log(str.length)
console.log(str.substring(0,5))
console.log(str.toUpperCase())
console.log(str.toLowerCase())
console.log(str.split(''))

let numbers = new Array(1,2,5,5,2)
console.log(numbers)
let fruit = ["hello",52,true,undefined]
fruit[7]="end"
fruit.push("pushed")
fruit.unshift('0110')
fruit.pop()
console.log(fruit)

console.log(Array.isArray(fruit))

console.log(fruit.indexOf(undefined))


let person = {
    name : "fname",
    addr : "addresss"
}

console.log(person)
console.log(person.addr)

person.email = "ad@ii"

console.log(person);

let todos = [
    {
        id: 1,
        text : "hello1",
        isdone:true
    },
    {
        id: 2,
        text : "hello2",
        isdone:false
    },
    {
        id: 3,
        text : "hello3",
        isdone:true
    }
]

let jsonval = JSON.stringify(todos)
console.log(todos,jsonval)

for (let i=0;i<3;i++)
    console.log(i)

let i=0;
while(i<10)
{
    console.log(i);
    i++;
}

for (let j of todos){
    console.log(j)
}

todos.forEach(function (todo) {console.log(todo.text)});
todos.forEach((todo) => {console.log(todo.text)});


let result =  todos.map((todo) => { return todo.text});
console.log(result)


let result2 =  todos.filter((todo) => { return todo.id>1});
console.log(result2)


let result3 =  todos.filter((todo) => { return todo.id>1}).map((todo)=>{return todo.isdone})
console.log(result3)

x='10',y=55
if (x===10 || y<100)
    console.log(true) 
else
    console.log(false)

let color = (y==55)?"blue":"red"
console.log(color)

switch(color){
    case "red" : console.log(1);break;
    case "blue" : console.log(11);break;
    case "re" : console.log(111);break;
    default:console.log(12)
}


function addnum(num1 , num2){
    console.log(typeof(num2))
    return num1+num2;
}

console.log(addnum(1,2))
console.log(addnum())
console.log(addnum(1))
console.log(addnum(1,'2'))
console.log(typeof(addnum(1,'2')))

console.log(addnum('1',2))
console.log(typeof(addnum('1',2)))

let vari = (num1) => num1+2;

console.log(vari(2))


function Person(firstName, lastName, dob) {
    // Set object properties
    this.firstName = firstName;
    this.lastName = lastName;
    this.dob = new Date(dob); 
    this.getBirthYear = function(){
      return this.dob.getFullYear();
    }
    
}

Person.prototype.getFullName = function() {
    return `${this.firstName} ${this.lastName}`
}

const p1 = new Person("ab","bc","2-1-2014")
console.log(p1)

class Person2 {
    constructor(firstName, lastName, dob){
    
    this.firstName = firstName;
    this.lastName = lastName;
    this.dob = new Date(dob); 
    }

    getBirthYear = function(){
      return this.dob.getFullYear();
    }
}
const p2 = new Person("ab","bc","2-1-2014")
console.log(p2)

/****************************************************************************************************/

console.log(document.getElementById('my-form'))
console.log(document.querySelector('#my-form'))
console.log(document.querySelectorAll(".item"))

document.querySelectorAll(".item").forEach((item)=> console.log(item) )

document.querySelectorAll(".item").forEach((item)=> item.textContent="changes done" )

document.querySelectorAll(".item")[0].remove()

const btn = document.querySelector('.btn');

btn.style.background="red";

// btn.addEventListener('click',(e)=>{
//     e.preventDefault();
//     console.log("click")
//     document.querySelector('#my-form').style.background="red";
//     document.querySelector('body').classList.add("bg-dark");
//     document.querySelector('.items').lastElementChild.innerHTML="hello boiys"
//     document.querySelector('.items').lastElementChild.style.background="black"
// })


btn.addEventListener('mouseover',(e)=>{
    e.preventDefault();
    console.log("click")
    document.querySelector('#my-form').style.background="blue";
    document.querySelector('body').classList.add("bg-dark");
    document.querySelector('.items').lastElementChild.innerHTML="hello boyass"
    document.querySelector('.items').lastElementChild.style.background="gray"
})


btn.addEventListener('mouseout',(e)=>{
    e.preventDefault();
    console.log("click")
    document.querySelector('#my-form').style.background="blue";
    document.querySelector('body').classList.add("bg-dark");
    document.querySelector('.items').lastElementChild.innerHTML="hello boyass"
    document.querySelector('.items').lastElementChild.style.background="yellow"
})

const myForm = document.querySelector('#my-form');
const nameInput = document.querySelector('#name');
const emailInput = document.querySelector('#email');
const msg = document.querySelector('.msg');
const userList = document.querySelector('#users');

myForm.addEventListener('submit', onSubmit);

function onSubmit(e) {
  e.preventDefault();
  
  if(nameInput.value === '' || emailInput.value === '') {
    msg.classList.add('error');
    msg.innerHTML = 'Please enter all fields';
    setTimeout(() => msg.remove(), 3000);
  } else {

    const li = document.createElement('li');

    // li.appendChild(document.createTextNode(`${nameInput.value}: ${emailInput.value}`));

    li.innerHTML = `<strong>${nameInput.value}</strong>e: ${emailInput.value}`;

    userList.appendChild(li);
    nameInput.value = '';
    emailInput.value = '';
  }
}