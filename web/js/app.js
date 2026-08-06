const logElement = document.getElementById("log");
const connectionElement = document.getElementById("connection");

const MAX_LOG = 10;

const logs = [];

document.querySelectorAll("[data-key]").forEach(button => {
    button.addEventListener("click", () => {
        sendKey(button.dataset.key);
    });
});

function addLog(message) {
    logs.unshift(message);
    if (logs.length > MAX_LOG) {
        logs.pop();
    }
    logElement.innerHTML = logs.join("<br>");
}

function setConnected(connected) {

    if (connected) {
        connectionElement.textContent = "🟢 Connected";
    } else {
        connectionElement.textContent = "🔴 Offline";
    }

}

async function sendKey(key) {
    try {

        const response = await fetch("/api/key", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                key: key
            })
        });
        const result = await response.json();
        if (response.ok && result.success) {
            setConnected(true);
            addLog(`✔ ${key}`);
        } else {
            setConnected(false);
            addLog(`✖ ${result.message}`);

        }
    } catch (err) {
        setConnected(false);
        addLog(`✖ ${err.message}`);

    }
}