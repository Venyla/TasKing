const usernameInput = document.querySelector('#username-input') as HTMLInputElement;
const usernameDisplay = document.querySelector('#username-display')
const usernameModal = document.querySelector('.username-modal')
const taskList = document.querySelector('.tasks');
const tasks = getTasks();
let username = localStorage.getItem("username");

function fillTaskList() {
    for (const task of tasks) {
        const html = `
            <li id="task-${task.id}" style="top: ${task.x}%; left: ${task.y}%; background-image: url(${task.image_url});" onclick="updateMyCount(${task.id}, this)">
                <img class="crown" src="../assets/crown-solid.svg" alt="A Crown"/>
                <div class="counter">
                    <span class="my-count">0</span>
                    <span class="their-count">0</span>
                </div>
                <div class="title">
                    ${task.title}
                </div>
            </li>`;

        taskList.insertAdjacentHTML('beforeend', html);
    }
}

function showUsernameOrPrompt() {
    if(username == null) {
        usernameModal.classList.remove("hidden");
    } else {
        usernameDisplay.innerHTML = username;
    }
}

function saveUsername(){
    username = usernameInput.value;
    usernameDisplay.innerHTML = username;
    usernameModal.classList.add("hidden");
    localStorage.setItem("username", username);
}

function initUglyPolling() {
    // setInterval(updateTheirCount,5000);
}

fillTaskList();
initUglyPolling();
showUsernameOrPrompt();

function updateMyCount(id: number, caller: Element) {
    const counter = caller.querySelector('.my-count');
    counter.innerHTML = String(+counter.innerHTML + 1);
    updateKingStatus(caller);
    postHistory();
}


function updateTheirCount() {
    for (const task of tasks){
        const random = Math.floor(Math.random() * (3 - 1 + 1)) + 1;
        const counter = document.querySelector(`#task-${task.id} .their-count`);
        counter.innerHTML = String(+counter.innerHTML + random);
        
        // getHistory(task.id);
    }
}

function updateKingStatus(element: Element) {
    const myCount = +element.querySelector('.my-count').innerHTML;
    const theirCount = +element.querySelector(`.their-count`).innerHTML;
    if(myCount > theirCount) {
        element.classList.add('king');
    } else {
        element.classList.remove('king');
    }
}


function getTasks(): Task[]  {

    return [
        {
           id: 1,
           x: 43,
           y: 51,
           title: 'Pet Elvis',
           image_url: 'https://scontent.fzrh3-1.fna.fbcdn.net/v/t39.30808-6/309430860_474589338041453_5993993981776996852_n.jpg?_nc_cat=108&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=7Tkot0fTAsoAX-FYOaX&_nc_ht=scontent.fzrh3-1.fna&oh=00_AfAsy-DH7-AJZmeKHnIMzUDN24XHC4BrAXHZ-SXf_ys_Kg&oe=643AC639'
        },
        {
            id: 2,
            x: 20,
            y: 35,
            title: 'Drink Beer',
            image_url: 'https://imageresizer.static9.net.au/mAbtmTO6BX05IdEILplNAgXv_Wc=/1200x675/https%3A%2F%2Fprod.static9.net.au%2Ffs%2Fff4238d6-65e7-4f73-afa2-537d3f64378e'
        }
    ]
}

function postHistory() {

}

interface Task {
    id: number,
    x: number,
    y: number,
    title: string,
    image_url: string,
}

interface TaskHistory {
    id: number,
    task_id: number,
    username_of_creator: string
}

interface Ranking {
    username_of_creator: string,
    rank: number,
}