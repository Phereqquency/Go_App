document.addEventListener('DOMContentLoaded', () => {
const form = document.getElementById('echoForm')
const result = document.getElementById('result')


form.addEventListener('submit', async (e) => {
e.preventDefault()
const input = document.getElementById('msg')
const payload = { message: input.value }


try {
const res = await fetch('/api/echo', {
method: 'POST',
headers: { 'Content-Type': 'application/json' },
body: JSON.stringify(payload)
})


const data = await res.json()
if (res.ok) {
result.textContent = data.reply || JSON.stringify(data)
} else {
result.textContent = 'Ошибка: ' + (data.error || JSON.stringify(data))
}
} catch (err) {
result.textContent = 'Network error: ' + err.message
}
})
})