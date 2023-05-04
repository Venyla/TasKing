const usernameInput = document.querySelector('#username-input') as HTMLInputElement;
const usernameDisplay = document.querySelector('#username-display')
const usernameModal = document.querySelector('.username-modal')
const taskList = document.querySelector('.tasks');
let username = localStorage.getItem("username");

const taskElements = new Map<string, Element>;

const ENDPOINT = '/api'; // todo traefik forward to backend
// old endpoint
//const ENDPOINT = 'http://127.0.0.1:8080';

function fillTaskList() {
    getTasks().then(tasks => {
        tasks.forEach(task => {
            const html = `
                <li id="t-${task.id}" style="top: ${task.x}%; left: ${task.y}%; background-image: url(${task.image_url});" onclick="updateMyCount('${task.id}')" data-current-max="0">
                    <img class="crown" src="../assets/crown-solid.svg" alt="A Crown"/>
                    <div class="counter">
                        <span class="my-count">0</span>
                        <span class="their-count">0</span>
                    </div>
                    <div class="title">
                        ${task.title}
                    </div>
                    <ol class="ranking"></ol>
                </li>`;

            const element = htmlToElement(html);
            taskList.insertAdjacentElement('beforeend', element)
            taskElements[task.id] = element
        })
    })
    updateRankings()
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
    // setInterval(updateRankings,5000);
}

fillTaskList();
initUglyPolling();
showUsernameOrPrompt();

function updateMyCount(id: string) {
    postHistory(id).then(() =>
        updateRanking(id)
    );
}


function updateRankings() {
    getTasks().then(tasks => {
        tasks.forEach(task => {
            updateRanking(task.id)
        })
    });
}

async function getRankings(id: string) {
    const response = await fetch(`${ENDPOINT}/api/tasks/${id}/rankings`);
    const endpointRankings = await response.json();
    const rankings: Ranking[] = Object.keys(endpointRankings).map(value => {
        return {
            username_of_creator: value,
            amount: +endpointRankings[value]
        }
    })
    rankings.sort((a, b) => b.amount - a.amount);
    return rankings
}

function updateRankingDisplay(id: string, myCount: number, maxCount: number, rankings: Ranking[]) {
    const rankingElement = taskElements[id].querySelector('.ranking')
    rankingElement.innerHTML = '';
    rankings.slice(0, 3).forEach(t =>{
        const html = `<li><span class="count">${t.amount}</span>${t.username_of_creator}</li>`
        rankingElement.insertAdjacentHTML('beforeend', html)
    });
    taskElements[id].querySelector('.my-count').innerHTML = myCount;
    taskElements[id].dataset.currentMax = maxCount;

    const index = rankings.findIndex(r => r.username_of_creator == username);
    taskElements[id].querySelector('.their-count').innerHTML = '#' + (index + 1);
}

async function updateRanking(id: string) {

    const rankings = await getRankings(id);
    const myRanking = rankings.find(value => value.username_of_creator === username);

    const maxCount = Math.max(...rankings.map(value => value.amount));
    const myCount = myRanking != undefined ? myRanking.amount : 0;

    updateRankingDisplay(id, myCount, maxCount, rankings);
    updateKingStatus(id, myCount, maxCount);
}

function updateKingStatus(id: string, myCount: number, maxCount: number) {
    if(myCount > 0 && myCount >= maxCount) {
        taskElements[id].classList.add('king');
    } else {
        taskElements[id].classList.remove('king');
    }
}

async function getTasks(): Promise<Task[]> {
    const response = await fetch(ENDPOINT + '/api/tasks');
    return response.json();
}

function htmlToElement(html: string) {
    const template = document.createElement('template');
    template.innerHTML = html.trim();
    return  template.content.firstElementChild;
}

async function postHistory(id: string) {
    const data: TaskHistory = {
        task_id: id,
        username_of_creator: username,
    }
    const response = await fetch(`${ENDPOINT}/api/history`, {
        method: 'POST',
        body: JSON.stringify(data),
    });
}

interface Task {
    id: string,
    x: number,
    y: number,
    title: string,
    image_url: string,
}

interface TaskHistory {
    task_id: string,
    username_of_creator: string
}

interface Ranking {
    username_of_creator: string,
    amount: number,
}

