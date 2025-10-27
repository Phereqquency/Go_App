const input = document.getElementById("message");
const button = document.getElementById("send");
const chat = document.getElementById("chat");

button.onclick = async () => {
    const msg = input.value;
    if (!msg) return;

    const res = await fetch("/api/echo", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ message: msg })
    });

    const data = await res.json();
    chat.innerHTML = "";
    data.messages.forEach(m => {
        const li = document.createElement("li");
        li.textContent = m.message;
        chat.appendChild(li);
    });

    input.value = "";
};
