var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var usernameInput = document.querySelector('#username-input');
var usernameDisplay = document.querySelector('#username-display');
var usernameModal = document.querySelector('.username-modal');
var taskList = document.querySelector('.tasks');
var username = localStorage.getItem("username");
var taskElements = new Map;
var ENDPOINT = 'http://127.0.0.1:8080';
function fillTaskList() {
    getTasks().then(function (tasks) {
        tasks.forEach(function (task) {
            var html = "\n                <li id=\"t-".concat(task.id, "\" style=\"top: ").concat(task.x, "%; left: ").concat(task.y, "%; background-image: url(").concat(task.image_url, ");\" onclick=\"updateMyCount('").concat(task.id, "')\" data-current-max=\"0\">\n                    <img class=\"crown\" src=\"../assets/crown-solid.svg\" alt=\"A Crown\"/>\n                    <div class=\"counter\">\n                        <span class=\"my-count\">0</span>\n                        <span class=\"their-count\">0</span>\n                    </div>\n                    <div class=\"title\">\n                        ").concat(task.title, "\n                    </div>\n                    <ol class=\"ranking\"></ol>\n                </li>");
            var element = htmlToElement(html);
            taskList.insertAdjacentElement('beforeend', element);
            taskElements[task.id] = element;
        });
    });
    updateRankings();
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
    // setInterval(updateRankings,5000);
}
fillTaskList();
initUglyPolling();
showUsernameOrPrompt();
function updateMyCount(id) {
    postHistory(id).then(function () {
        return updateRanking(id);
    });
}
function updateRankings() {
    getTasks().then(function (tasks) {
        tasks.forEach(function (task) {
            updateRanking(task.id);
        });
    });
}
function getRankings(id) {
    return __awaiter(this, void 0, void 0, function () {
        var response, endpointRankings, rankings;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, fetch("".concat(ENDPOINT, "/api/tasks/").concat(id, "/rankings"))];
                case 1:
                    response = _a.sent();
                    return [4 /*yield*/, response.json()];
                case 2:
                    endpointRankings = _a.sent();
                    rankings = Object.keys(endpointRankings).map(function (value) {
                        return {
                            username_of_creator: value,
                            amount: +endpointRankings[value]
                        };
                    });
                    rankings.sort(function (a, b) { return b.amount - a.amount; });
                    return [2 /*return*/, rankings];
            }
        });
    });
}
function updateRankingDisplay(id, myCount, maxCount, rankings) {
    var rankingElement = taskElements[id].querySelector('.ranking');
    rankingElement.innerHTML = '';
    rankings.slice(0, 3).forEach(function (t) {
        var html = "<li><span class=\"count\">".concat(t.amount, "</span>").concat(t.username_of_creator, "</li>");
        rankingElement.insertAdjacentHTML('beforeend', html);
    });
    taskElements[id].querySelector('.my-count').innerHTML = myCount;
    taskElements[id].dataset.currentMax = maxCount;
    var index = rankings.findIndex(function (r) { return r.username_of_creator == username; });
    taskElements[id].querySelector('.their-count').innerHTML = '#' + (index + 1);
}
function updateRanking(id) {
    return __awaiter(this, void 0, void 0, function () {
        var rankings, myRanking, maxCount, myCount;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, getRankings(id)];
                case 1:
                    rankings = _a.sent();
                    myRanking = rankings.find(function (value) { return value.username_of_creator === username; });
                    maxCount = Math.max.apply(Math, rankings.map(function (value) { return value.amount; }));
                    myCount = myRanking != undefined ? myRanking.amount : 0;
                    updateRankingDisplay(id, myCount, maxCount, rankings);
                    updateKingStatus(id, myCount, maxCount);
                    return [2 /*return*/];
            }
        });
    });
}
function updateKingStatus(id, myCount, maxCount) {
    if (myCount > 0 && myCount >= maxCount) {
        taskElements[id].classList.add('king');
    }
    else {
        taskElements[id].classList.remove('king');
    }
}
function getTasks() {
    return __awaiter(this, void 0, void 0, function () {
        var response, endpointTasks;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, fetch(ENDPOINT + '/api/tasks')];
                case 1:
                    response = _a.sent();
                    return [4 /*yield*/, response.json()];
                case 2:
                    endpointTasks = _a.sent();
                    return [2 /*return*/, endpointTasks.map(function (t) {
                            return {
                                id: t.id,
                                title: t.Title,
                                image_url: t.IconUrl,
                                x: t.XCoordinates,
                                y: t.YCoordinates
                            };
                        })];
            }
        });
    });
}
function htmlToElement(html) {
    var template = document.createElement('template');
    template.innerHTML = html.trim();
    return template.content.firstElementChild;
}
function postHistory(id) {
    return __awaiter(this, void 0, void 0, function () {
        var data, response;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    data = {
                        id: undefined,
                        TaskId: id,
                        CreatedBy: username
                    };
                    return [4 /*yield*/, fetch("".concat(ENDPOINT, "/api/history"), {
                            method: 'POST',
                            body: JSON.stringify(data)
                        })];
                case 1:
                    response = _a.sent();
                    return [2 /*return*/];
            }
        });
    });
}
