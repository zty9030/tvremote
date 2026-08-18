const logElement = document.getElementById("log");
const connectionElement = document.getElementById("connection");

const MAX_LOG = 10;

const logs = [];

document.querySelectorAll("[data-key]").forEach(button => {
    const key = button.dataset.key;

    if (key === "VOL_UP" || key === "VOL_DOWN") {
        setupHoldButton(button, key);
    } else {
        button.addEventListener("click", () => {
            sendKey(key);
        });
    }
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

function setupHoldButton(button, key) {
    let pressed = false;

    const release = async event => {
        if (event) {
            event.preventDefault();
        }
        if (!pressed) {
            return;
        }
        pressed = false;
        button.classList.remove("pressed");
        await sendKeyUp(key);
    };

    button.addEventListener("pointerdown", async event => {
        event.preventDefault();

        if (pressed) {
            return;
        }

        pressed = true;

        button.setPointerCapture(event.pointerId);
        button.classList.add("pressed");

        await sendKeyDown(key);
    });

    button.addEventListener("pointerup", release);
    button.addEventListener("pointercancel", release);
    button.addEventListener("lostpointercapture", release);
}

async function sendKeyDown(key) {
    try {
        const response = await fetch("/api/key/down", {
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
            addLog(`↓ ${key}`);
        } else {
            setConnected(false);
            addLog(`✖ ${result.message}`);
        }
    } catch (err) {
        setConnected(false);
        addLog(`✖ ${err.message}`);
    }
}

async function sendKeyUp(key) {
    try {
        const response = await fetch("/api/key/up", {
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
            addLog(`↑ ${key}`);
        } else {
            setConnected(false);
            addLog(`✖ ${result.message}`);
        }
    } catch (err) {
        setConnected(false);
        addLog(`✖ ${err.message}`);
    }
}