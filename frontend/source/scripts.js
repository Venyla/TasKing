var usernameInput = document.querySelector('#username-input');
var usernameDisplay = document.querySelector('#username-display');
var usernameModal = document.querySelector('.username-modal');
var taskList = document.querySelector('.tasks');
var tasks = getTasks();
var username = localStorage.getItem("username");
function fillTaskList() {
    for (var _i = 0, tasks_1 = tasks; _i < tasks_1.length; _i++) {
        var task = tasks_1[_i];
        var html = "\n            <li id=\"task-".concat(task.id, "\" style=\"top: ").concat(task.x, "%; left: ").concat(task.y, "%; background-image: url(").concat(task.image_url, ");\" onclick=\"updateMyCount(").concat(task.id, ", this)\">\n                <img class=\"crown\" src=\"../assets/crown-solid.svg\" alt=\"A Crown\"/>\n                <div class=\"counter\">\n                    <span class=\"my-count\">0</span>\n                    <span class=\"their-count\">0</span>\n                </div>\n                <div class=\"title\">\n                    ").concat(task.title, "\n                </div>\n            </li>");
        taskList.insertAdjacentHTML('beforeend', html);
    }
}
function showUsernameOrPrompt() {
    if (username == null) {
        usernameModal.classList.remove("hidden");
    }
    else {
        usernameDisplay.innerHTML = username;
    }
}
function saveUsername() {
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
function updateMyCount(id, caller) {
    var counter = caller.querySelector('.my-count');
    counter.innerHTML = String(+counter.innerHTML + 1);
    updateKingStatus(caller);
    postHistory();
}
function updateTheirCount() {
    for (var _i = 0, tasks_2 = tasks; _i < tasks_2.length; _i++) {
        var task = tasks_2[_i];
        var random = Math.floor(Math.random() * (3 - 1 + 1)) + 1;
        var counter = document.querySelector("#task-".concat(task.id, " .their-count"));
        counter.innerHTML = String(+counter.innerHTML + random);
        // getHistory(task.id);
    }
}
function updateKingStatus(element) {
    var myCount = +element.querySelector('.my-count').innerHTML;
    var theirCount = +element.querySelector(".their-count").innerHTML;
    if (myCount > theirCount) {
        element.classList.add('king');
    }
    else {
        element.classList.remove('king');
    }
}
function getTasks() {
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
    ];
}
function postHistory() {
}
